package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

type ComponentType struct {
	Package string
	Name    string
	Fields  []SchemaField
	Id      uint
}

type TypeRef struct {
	Name string
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
	var primitiveTypeToGoType = map[string]string{
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
		return primitiveTypeToGoType[t.(PrimitiveType).Name]
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
	var primitiveTypeToFunctionFamily = map[string]string{
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
		return primitiveTypeToFunctionFamily[t.(PrimitiveType).Name]
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
	log.Fatal("Could not find function family for ", t)
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
	log.Fatal("Could not find method suffix for ", t)
	return ""
}

func TraverseTypeDependencies(types []interface{}) []interface{} {
	foundTypes := map[string]interface{}{}
	explorationQueue := types

	for len(explorationQueue) > 0 {
		nextType := explorationQueue[0]
		explorationQueue = explorationQueue[1:]
		foundTypes[GoTypeFor(nextType)] = nextType
		switch nextType.(type) {
		case ComponentType:
			for _, field := range nextType.(ComponentType).Fields {
				explorationQueue = append(explorationQueue, field.Type)
				explorationQueue = append(explorationQueue, OptionType{field.Type})
			}

		case ObjectType:
			for _, field := range nextType.(ObjectType).Fields {
				explorationQueue = append(explorationQueue, field.Type)
			}

		case ListType:
			explorationQueue = append(explorationQueue, nextType.(ListType).Type)

		case MapType:
			explorationQueue = append(explorationQueue, nextType.(MapType).KeyType)
			explorationQueue = append(explorationQueue, nextType.(MapType).ValueType)

		case OptionType:
			explorationQueue = append(explorationQueue, nextType.(OptionType).Type)
		}
	}

	result := []interface{}{}
	for _, foundType := range foundTypes {
		result = append(result, foundType)
	}

	return result
}

func GenerateMethodsForTypes(types []interface{}) string {
	result := ""
	for _, typ := range types {
		switch typ.(type) {
		case ComponentType:
			componentType := typ.(ComponentType)
			result += GenerateComponentType(componentType)
			result += GenerateComponentUpdateType(componentType)
			result += GenerateReadComponentType(componentType)
			result += GenerateWriteComponentType(componentType)
			result += GenerateReadComponentUpdateType(componentType)
			result += GenerateWriteComponentUpdateType(componentType)
			result += GenerateComponentEventCallbacks(componentType)
			result += GenerateAddComponentDispatcherMethod(componentType)
			result += GenerateUpdateComponentDispatcherMethod(componentType)
			result += GenerateAuthorityComponentDispatcherMethod(componentType)
			result += GenerateRemoveComponentDispatcherMethod(componentType)
			result += GenerateUpdateComponentConnectionMethod(componentType)

		case EnumType:
			enumType := typ.(EnumType)
			result += GenerateEnumType(enumType)
			result += GenerateReadEnumType(enumType)
			result += GenerateWriteEnumType(enumType)

		case ObjectType:
			objectType := typ.(ObjectType)
			result += GenerateObjectType(objectType)
			result += GenerateReadObjectType(objectType)
			result += GenerateWriteObjectType(objectType)

		case ListType:
			listType := typ.(ListType)
			result += GenerateReadListType(listType)
			result += GenerateWriteListType(listType)

		case MapType:
			mapType := typ.(MapType)
			result += GenerateReadMapType(mapType)
			result += GenerateWriteMapType(mapType)

		case OptionType:
			optionType := typ.(OptionType)
			result += GenerateReadOptionType(optionType)
			result += GenerateWriteOptionType(optionType)

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

	metaDataFields := []SchemaField{{Name: "EntityType", Type: PrimitiveType{Name:"string"}, Id: 1}}
	metaDataComponentType := ComponentType{Package: "", Name: "MetaData", Fields: metaDataFields, Id: 53}

	attributeSetType := ObjectType{Package: "", Name: "WorkerAttributeSet ", Fields: []SchemaField{{Name: "attribute", Type: ListType{Type: PrimitiveType{"string"}}, Id: 1}}}
	requirementSetType := ObjectType{Package: "", Name: "WorkerRequirementSet ", Fields: []SchemaField{{Name: "attribute_set", Type: ListType{attributeSetType}, Id: 1}}}
	aclFields := []SchemaField{
		{Name: "Read", Type: ListType{requirementSetType}, Id: 1},
		{Name: "Write", Type: MapType{KeyType: PrimitiveType{Name: "uint32"}, ValueType: ListType{requirementSetType}}, Id: 2},
	}
	aclComponentType := ComponentType{Package: "", Name: "EntityAcl", Fields: aclFields, Id: 50}

	testEnum := EnumType{Name: "Color", values: map[int]string{1: "Blue", 2: "Red"}}

	allTypes := TraverseTypeDependencies([]interface{}{positionComponentType, aclComponentType, testEnum, metaDataComponentType})

	ioutil.WriteFile("output.go.generated", []byte("package main\n\n" +GenerateMethodsForTypes(allTypes)), 0644)
}
