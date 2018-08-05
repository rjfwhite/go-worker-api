package main

import (
	"fmt"
	"io/ioutil"
)

type ComponentType struct {
	Package string
	Name    string
	Fields  []SchemaField
	Id      uint
}

type SchemaField struct {
	Name string
	Type interface{}
	Id   int
}

type PrimitiveType struct {
	Name string
}

type EnumType struct {
	Package string
	Name    string
	values  map[int]string
}

type ObjectType struct {
	Package string
	Name    string
	Fields  []SchemaField
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
	case EnumType:
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
	case ListType:
		return "List_" + MethodSuffixForType(t.(ListType).Type)
	case MapType:
		return "Map_" + MethodSuffixForType(t.(MapType).KeyType) + "_to_" + MethodSuffixForType(t.(MapType).ValueType)
	case OptionType:
		return "Option_" + MethodSuffixForType(t.(OptionType).Type)
	}
	return ""
}

func TraverseTypeDependencies(types []interface{}) []interface{} {
	found_types := map[string]interface{}{}
	exploration_queue := types

	for len(exploration_queue) > 0 {
		next_type := exploration_queue[0]
		exploration_queue = exploration_queue[1:]
		found_types[GoTypeFor(next_type)] = next_type
		switch next_type.(type) {
		case ComponentType:
			for _, field := range next_type.(ComponentType).Fields {
				exploration_queue = append(exploration_queue, field.Type)
				exploration_queue = append(exploration_queue, OptionType{field.Type})
			}

		case ObjectType:
			for _, field := range next_type.(ObjectType).Fields {
				exploration_queue = append(exploration_queue, field.Type)
			}

		case ListType:
			exploration_queue = append(exploration_queue, next_type.(ListType).Type)

		case MapType:
			exploration_queue = append(exploration_queue, next_type.(MapType).KeyType)
			exploration_queue = append(exploration_queue, next_type.(MapType).ValueType)

		case OptionType:
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
		case ComponentType:
			component_type := typ.(ComponentType)
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

		case EnumType:
			enum_type := typ.(EnumType)
			result += GenerateEnumType(enum_type)
			result += GenerateReadEnumType(enum_type)
			result += GenerateWriteEnumType(enum_type)

		case ObjectType:
			object_type := typ.(ObjectType)
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
	coordinatesFields := []SchemaField{{Name: "X", Type: PrimitiveType{Name: "double"}, Id: 1}, {Name: "Y", Type: PrimitiveType{Name: "double"}, Id: 2}, {Name: "Z", Type: PrimitiveType{Name: "double"}, Id: 3}}
	coordinatesType := ObjectType{Package: "", Name: "Coordinates", Fields: coordinatesFields}

	positionFields := []SchemaField{{Name: "Coords", Type: coordinatesType, Id: 1}}
	positionComponentType := ComponentType{Package: "", Name: "Position", Fields: positionFields, Id: 54}

	attributeSetType := ObjectType{Package: "", Name: "WorkerAttributeSet ", Fields: []SchemaField{{Name: "attribute", Type: ListType{Type: PrimitiveType{"string"}}, Id: 1}}}
	requirementSetType := ObjectType{Package: "", Name: "WorkerRequirementSet ", Fields: []SchemaField{{Name: "attribute_set", Type: ListType{attributeSetType}, Id: 1}}}
	aclFields := []SchemaField{
		{Name: "Read", Type: ListType{requirementSetType}, Id: 1},
		{Name: "Write", Type: MapType{KeyType: PrimitiveType{Name: "uint32"}, ValueType: ListType{requirementSetType}}, Id: 2},
	}
	aclComponentType := ComponentType{Package: "", Name: "EntityAcl", Fields: aclFields, Id: 50}

	testEnum := EnumType{Name: "Color", values: map[int]string{1: "Blue", 2: "Red"}}

	all_types := TraverseTypeDependencies([]interface{}{positionComponentType, aclComponentType, testEnum})

	TranslateFiles()

	ioutil.WriteFile("output.go.generated", []byte("package main\n\n" +GenerateMethodsForTypes(all_types)), 0644)
}
