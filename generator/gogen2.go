package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

type ComponentDefinition struct {
	Type    ComponentType
	Fields  []FieldDefinition
	Id      uint
}

type FieldDefinition struct {
	Name string
	Type interface{}
	Id   int
}

type EnumDefinition struct {
	Type    EnumType
	values  map[int]string
}

type ObjectDefinition struct {
	Type ObjectType
	Fields  []FieldDefinition
}

type EnumType struct {
	Name string
}

type ComponentType struct {
	Name string
}

type ObjectType struct {
	Name string
}

type PrimitiveType struct {
	Name string
}

type OptionType struct {
	Type interface{}
}

type ListType struct {
	Type interface{}
}

type MapType struct {
	KeyType   interface{}
	ValueType interface{}
}

func GoTypeFor(t interface{}) string {
	var primitive_type_to_go_type = map[string]string{
		"int32":    "int",
		"int64":    "int64",
		"uint32":   "uint",
		"uint64":   "uint64",
		"sint32":   "int",
		"sint64":   "int64",
		"fixed32":  "uint",
		"fixed64":  "uint64",
		"sfixed32": "int",
		"sfixed64": "int64",
		"bool":     "bool",
		"float":    "float32",
		"double":   "float64",
		"string":   "string",
		"EntityId": "int64",
		"bytes":    "[]byte",
	}

	switch t.(type) {
	case PrimitiveType:
		return primitive_type_to_go_type[t.(PrimitiveType).Name]
	case EnumType:
		return t.(EnumType).Name
	case ListType:
		return "[]" + GoTypeFor(t.(ListType).Type)
	case ObjectType:
		return t.(ObjectType).Name
	case ComponentType:
		return t.(ComponentType).Name
	case MapType:
		return fmt.Sprintf("map[%s]%s", GoTypeFor(t.(MapType).KeyType), GoTypeFor(t.(MapType).ValueType))
	case OptionType:
		return "*" + GoTypeFor(t.(OptionType).Type)
	}
	return ""
}

func FunctionFamilyFor(t interface{}) string {
	var primitive_type_to_function_family = map[string]string{
		"int32":    "Int32",
		"int64":    "Int64",
		"uint32":   "Uint32",
		"uint64":   "Uint64",
		"sint32":   "Sint32",
		"sint64":   "Sint64",
		"fixed32":  "Fixed32",
		"fixed64":  "Fixed64",
		"sfixed32": "Sfixed32",
		"sfixed64": "Sfixed64",
		"bool":     "Bool",
		"float":    "Float",
		"double":   "Double",
		"string":   "Bytes",
		"EntityId": "EntityId",
		"bytes":    "Bytes",
	}

	switch t.(type) {
	case PrimitiveType:
		return primitive_type_to_function_family[t.(PrimitiveType).Name]
	case EnumDefinition:
		return "Enum"
	case ListType:
		return "Object"
	case ObjectType:
		return "Object"
	case MapType:
		return "Object"
	case OptionType:
		return FunctionFamilyFor(t.(OptionType).Type)
	}
	log.Fatal("Unknown Schema Type", t)
	return ""
}

func MethodSuffixForType(t interface{}) string {
	switch t.(type) {
	case PrimitiveType:
		return "Primitive_" + t.(PrimitiveType).Name
	case EnumType:
		return "Enum_" + t.(EnumType).Name
	case ObjectType:
		return "Object_" + t.(ObjectType).Name
	case ComponentType:
		return "Object_" + t.(ComponentType).Name
	case ListType:
		return "List_" + MethodSuffixForType(t.(ListType).Type)
	case MapType:
		return "Map_" + MethodSuffixForType(t.(MapType).KeyType) + "_to_" + MethodSuffixForType(t.(MapType).ValueType)
	case OptionType:
		return "Option_" + MethodSuffixForType(t.(OptionType).Type)
	}
	log.Fatal("Unknown Schema Type", t)
	return ""
}

func TraverseTypeDependencies(types []interface{}) []interface{} {
	found_types := map[string]interface{}{}
	exploration_queue := types

	for len(exploration_queue) > 0 {
		next_type := exploration_queue[0]
		exploration_queue = exploration_queue[1:]

		switch next_type.(type) {
		case ComponentDefinition:
			found_types[GoTypeFor(next_type.(ComponentDefinition).Type)] = next_type.(ComponentDefinition)
			for _, field := range next_type.(ComponentDefinition).Fields {
				exploration_queue = append(exploration_queue, field.Type)
				exploration_queue = append(exploration_queue, OptionType{field.Type})
			}

		case ObjectDefinition:
			found_types[GoTypeFor(next_type.(ObjectDefinition).Type)] = next_type.(ObjectDefinition)
			for _, field := range next_type.(ObjectDefinition).Fields {
				exploration_queue = append(exploration_queue, field.Type)
			}

		case ListType:
			found_types[GoTypeFor(next_type)] = next_type
			exploration_queue = append(exploration_queue, next_type.(ListType).Type)

		case MapType:
			found_types[GoTypeFor(next_type)] = next_type
			exploration_queue = append(exploration_queue, next_type.(MapType).KeyType)
			exploration_queue = append(exploration_queue, next_type.(MapType).ValueType)

		case OptionType:
			found_types[GoTypeFor(next_type)] = next_type
			exploration_queue = append(exploration_queue, next_type.(OptionType).Type)

		}
	}

	result := []interface{}{}
	for _, found_type := range (found_types) {
		result = append(result, found_type)
	}

	return result
}

func GenerateMethodsForTypes(types []interface{}) string {
	result := ""
	for _, typ := range types {
		switch typ.(type) {
		case ComponentDefinition:
			component_type := typ.(ComponentDefinition)
			result += GenerateComponentType(component_type)
			result += GenerateComponentUpdateType(component_type)
			result += GenerateReadComponentType(component_type)
			result += GenerateWriteComponentType(component_type)
			result += GenerateReadComponentUpdateType(component_type)
			result += GenerateWriteComponentUpdateType(component_type)
			result += GenerateComponentEventCallbacks(component_type)
			result += GenerateAddComponentDispatcherMethod(component_type)
			result += GenerateUpdateComponentDispatcherMethod(component_type)
			result += GenerateAuthorityComponentDispatcherMethod(component_type)
			result += GenerateRemoveComponentDispatcherMethod(component_type)
			result += GenerateUpdateComponentConnectionMethod(component_type)

		case EnumDefinition:
			enum_type := typ.(EnumDefinition)
			result += GenerateEnumType(enum_type)
			result += GenerateReadEnumType(enum_type)
			result += GenerateWriteEnumType(enum_type)

		case ObjectDefinition:
			object_type := typ.(ObjectDefinition)
			result += GenerateObjectType(object_type)
			result += GenerateReadObjectType(object_type)
			result += GenerateWriteObjectType(object_type)

		case ListType:
			list_type := typ.(ListType)
			result += GenerateReadListType(list_type)
			result += GenerateWriteListType(list_type)

		case MapType:
			map_type := typ.(MapType)
			result += GenerateReadMapType(map_type)
			result += GenerateWriteMapType(map_type)

		case OptionType:
			option_type := typ.(OptionType)
			result += GenerateReadOptionType(option_type)
			result += GenerateWriteOptionType(option_type)

		default:

		}
	}
	return result
}

func main() {
	coordinatesFields := []FieldDefinition{{Name: "X", Type: PrimitiveType{Name: "double"}, Id: 1}, {Name: "Y", Type: PrimitiveType{Name: "double"}, Id: 2}, {Name: "Z", Type: PrimitiveType{Name: "double"}, Id: 3}}
	coordinatesDefinition := ObjectDefinition{Type:ObjectType{"Coordinates"}, Fields: coordinatesFields}

	positionFields := []FieldDefinition{{Name: "Coords", Type: coordinatesDefinition.Type, Id: 1}}
	positionComponentDefinition := ComponentDefinition{Type:ComponentType{"Position"}, Fields: positionFields, Id: 54}

	attributeSetDefinition := ObjectDefinition{Type:ObjectType{"WorkerAttributeSet"}, Fields: []FieldDefinition{{Name: "attribute", Type: ListType{Type: PrimitiveType{"string"}}, Id: 1}}}
	requirementSetDefinition := ObjectDefinition{Type:ObjectType{"WorkerRequirementSet"}, Fields: []FieldDefinition{{Name: "attribute_set", Type: ListType{attributeSetDefinition.Type}, Id: 1}}}
	aclFields := []FieldDefinition{
		{Name: "Read", Type: ListType{requirementSetDefinition.Type}, Id: 1},
		{Name: "Write", Type: MapType{KeyType: PrimitiveType{Name: "uint32"}, ValueType: ListType{requirementSetDefinition.Type}}, Id: 2},
	}
	aclComponentType := ComponentDefinition{Type:ComponentType{"EntityAcl"}, Fields: aclFields, Id: 50}

	testEnum := EnumDefinition{Type:EnumType{"Color"}, values: map[int]string{1: "Blue", 2: "Red"}}

	all_types := TraverseTypeDependencies([]interface{}{positionComponentDefinition, aclComponentType, testEnum})

	TranslateFiles()

	ioutil.WriteFile("output.go.generated", []byte("package main\n\n" +GenerateMethodsForTypes(all_types)), 0644)
}
