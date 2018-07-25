package main

import (
	"fmt"
)

type SchemaType struct {
	Package string
	Name    string
	Fields  []SchemaField
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
	Name string
}

type ObjectType struct {
	Name string
}

type OptionType struct {
	Type interface{}
}

type ListType struct {
	Type interface{}
}

type MapType struct {
	KeyType interface{}
	ValueType interface{}
}

func GenerateReadType(t interface{}) string {
	switch t.(type) {
	case PrimitiveType:
		GenerateReadPrimitiveType(t.(PrimitiveType))
	}
	return ""
}

func GenerateReadPrimitiveType(t PrimitiveType) string {
	go_type := GoTypeFor(t.Name)
	function_family := primitive_type_to_function_family[t.Name]
	output := ""

	output += fmt.Sprintf("func ReadPrimitive_%s(object example.Schema_Object, field uint, index uint) %s {\n", t.Name, go_type)

	boolFix := ""
	if t.Name == "bool" {
		boolFix = " > 0"
	}

	output += fmt.Sprintf("\treturn example.Schema_Get%s(object, index)%s\n", function_family, boolFix)
	output += "}\n"

	return output
}

func GenerateWritePrimitiveType(t PrimitiveType) string {
	go_type := GoTypeFor(t)
	function_family := primitive_type_to_function_family[t.Name]
	output := ""

	output += fmt.Sprintf("func WritePrimitive_%s(object example.Schema_Object, field uint, value %s) {\n", t.Name, go_type)

	output += fmt.Sprintf("\texample.Schema_Add%s(object, index, value)\n", function_family)
	output += "}\n"

	return output
}

func GenerateReadListType(t ListType) string {
	output := ""
	output += fmt.Sprintf("func Read%s(object example.Schema_Object, field uint, index uint) %s {\n", MethodSuffixForType(t), GoTypeFor(t))
	output += fmt.Sprintf("\tcount := example.Schema_Get%sCount(object, field)\n", FunctionFamilyFor(t.Type))
	output += fmt.Sprintf("\tresult := %s{}\n", GoTypeFor(t))
	output += "\tfor i := uint(0); i < count; i++ {\n"
	output += fmt.Sprintf("\t\tresult = append(result, Read%s(object, field, i))\n", MethodSuffixForType(t.Type))
	output += "\t}\n"
	output += "\treturn result\n"
	output += "}\n"
	return output
}

func GenerateReadObjectType(t SchemaType) string {
	output := ""
	output += fmt.Sprintf("func ReadObject_%s(object example.Schema_Object, field uint, index uint) %s {\n", t.Name, t.Name)
	output += fmt.Sprintf("\tinnerObject := example.Schema_IndexObject(object, field, index)\n")
	output += fmt.Sprintf("\treturn %s {\n", t.Name)
	for _, f := range t.Fields {
		output += fmt.Sprintf("\t\t%s : Read%s(innerObject, %d, 0),\n", f.Name, MethodSuffixForType(f.Type), f.Id)
	}
	output += "\t}\n"
	output += "}\n"
	return output
}

func GenerateWriteObjectType(t SchemaType) string {
	output := ""
	output += fmt.Sprintf("func WriteObject_%s(object example.Schema_Object, field uint, value %s) {\n", t.Name, t.Name)
	output += "\texample.Schema_AddObject(object, field)\n"
	output += fmt.Sprintf("\tinnerObject := example.Schema_GetObject(object, field)\n") // will this hit a bug?
	for _, f := range t.Fields {
		output += fmt.Sprintf("\tWrite%s(innerObject, %d, value.%s)\n", MethodSuffixForType(f.Type), f.Id, f.Name)
	}
	output += "}\n"
	return output
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
	case ListType:
		return "[]" + GoTypeFor(t.(ListType).Type)
	case ObjectType:
		return t.(ObjectType).Name
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
	case ListType:
		return "Object"
	case ObjectType:
		return "Object"
	}
	return ""
}

func MethodSuffixForType(t interface{}) string {
	switch t.(type) {
	case PrimitiveType:
		return "Primitive_" + t.(PrimitiveType).Name
	case ObjectType:
		return "Object_" + t.(ObjectType).Name
	case ListType:
		return "List_" + MethodSuffixForType(t.(ListType).Type)
	}
	return ""
}

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





func GenerateStruct(t SchemaType) string {
	output := ""
	output += fmt.Sprintf("type %s struct {\n", t.Name)

	for _, f := range t.Fields {
		output += fmt.Sprintf("\t%s %s\n", f.Name, GoTypeFor(f.Type))
	}

	output += "}\n"
	return output
}

func GenerateUpdateStruct(t SchemaType) string {
	output := ""
	output += fmt.Sprintf("type %sUpdate struct {\n", t.Name)

	for _, f := range t.Fields {
		output += fmt.Sprintf("\t%s *%s\n", f.Name, f.Type)
	}

	output += "}\n"
	return output
}

func GenerateApplyUpdate(t SchemaType) string {
	output := ""
	output += fmt.Sprintf("func Apply%sUpdate(data %s, update %sUpdate)  {\n", t.Name, t.Name, t.Name)

	for _, f := range t.Fields {
		output += fmt.Sprintf("\tif update.%s != nil {\n", f.Name)
		output += fmt.Sprintf("\t\tdata.%s = *update.%s\n", f.Name, f.Name)
		output += "\t}\n"
	}

	output += "}\n"
	return output
}

func GenerateUpdateWriter(t SchemaType) string {
	output := ""
	output += fmt.Sprintf("func Write%sUpdate(output example.Schema_Object, update %sUpdate) {\n", t.Name, t.Name)
	for _, f := range t.Fields {
		output += fmt.Sprintf("\tif update.%s != nil {\n", f.Name)
		output += GenerateFieldWriter("\t\t", f, "*update."+f.Name)
		output += "\t}\n"
	}
	output += "}\n"
	return output
}

func GenerateFieldUpdateReader(field SchemaField) string {
	output := ""
	if field.Type == "float64" {
		output += fmt.Sprintf("\tif example.Schema_GetDoubleCount(input, %d) > 0 {\n", field.Id)
	} else {
		output += fmt.Sprintf("\tif example.Schema_GetObjectCount(input, %d) > 0 {\n", field.Id)
	}
	output += GenerateFieldReader("\t\tvalue := ", field)
	output += fmt.Sprintf("\t\tresult.%s = &value\n", field.Name)
	output += "\t}\n"
	return output
}

func GenerateFieldReader(prefix string, field SchemaField) string {
	if field.Type == "float64" {
		return fmt.Sprintf("%sexample.Schema_GetDouble(input, %d)\n", prefix, field.Id)
	} else {
		return fmt.Sprintf("%sRead%s(example.Schema_GetObject(input, %d))\n", prefix, field.Type, field.Id)
	}
}

func GenerateFieldWriter(prefix string, field SchemaField, value string) string {
	if field.Type == "float64" {
		return fmt.Sprintf("%sexample.Schema_AddDouble(output, %d, %s)\n", prefix, field.Id, value)
	} else {
		a := fmt.Sprintf("%sexample.Schema_AddObject(output, %d)\n", prefix, field.Id)
		b := fmt.Sprintf("%sWrite%s(example.Schema_GetObject(output, %d), %s)\n", prefix, field.Type, field.Id, value)
		return a + b
	}
}

func GenerateUpdateReader(t SchemaType) string {
	output := ""
	output += fmt.Sprintf("func Read%sUpdate(input example.Schema_Object) %sUpdate {\n", t.Name, t.Name)
	output += fmt.Sprintf("\tresult := %sUpdate{}\n", t.Name)

	for _, f := range t.Fields {
		output += GenerateFieldUpdateReader(f)
	}
	output += "\treturn result\n"
	output += "}\n"
	return output
}

func GenerateReader(t SchemaType) string {
	output := ""
	output += fmt.Sprintf("func Read%s(input example.Schema_Object) %s {\n", t.Name, t.Name)
	output += fmt.Sprintf("\treturn %s {\n", t.Name)
	for _, f := range t.Fields {
		output += GenerateFieldReader(fmt.Sprintf("\t\t%s : ", f.Name), f)
	}
	output += "\t}\n"
	output += "}\n"
	return output
}

func GenerateWriter(t SchemaType) string {
	output := ""
	output += fmt.Sprintf("func Write%s(output example.Schema_Object, value %s) {\n", t.Name, t.Name)
	for _, f := range t.Fields {
		output += GenerateFieldWriter("\t", f, "value."+f.Name)
	}
	output += "}\n"
	return output
}

func main() {
	positionFields := []SchemaField{{Name: "Coords", Type: ObjectType{Name:"Coordinates"}, Id: 1}}
	positionType := SchemaType{Package: "", Name: "Position", Fields: positionFields}
	coordinatesFields := []SchemaField{{Name: "X", Type: PrimitiveType{Name:"double"}, Id: 1}, {Name: "Y", Type: PrimitiveType{Name:"double"}, Id: 2}, {Name: "Z", Type: PrimitiveType{Name:"double"}, Id: 3}}
	coordinatesType := SchemaType{Package: "", Name: "Coordinates", Fields: coordinatesFields}


	fmt.Println("package main")
	//for k, _ := range(primitive_type_to_function_family) {
	//	fmt.Println(GenerateReadPrimitiveType(PrimitiveType{Name:k}))
	//	fmt.Println(GenerateWritePrimitiveType(PrimitiveType{Name:k}))
	//}

	attributeSetType := SchemaType{Package:"", Name:"WorkerAttributeSet ", Fields:[]SchemaField{{Name: "attribute", Type: ListType{Type:PrimitiveType{"string"}}, Id: 1}}}
	requirementSetType := SchemaType{Package:"", Name:"WorkerRequirementSet ", Fields:[]SchemaField{{Name: "attribute_set", Type: ListType{Type:ObjectType{"WorkerAttributeSet"}}, Id: 1}}}


	fmt.Print(GenerateStruct(attributeSetType))
	fmt.Print(GenerateStruct(requirementSetType))
	fmt.Print(GenerateStruct(coordinatesType))
	fmt.Print(GenerateStruct(positionType))
	fmt.Println(GenerateReadObjectType(attributeSetType))
	fmt.Println(GenerateReadObjectType(requirementSetType))
	fmt.Println(GenerateReadObjectType(coordinatesType))
	fmt.Println(GenerateWriteObjectType(coordinatesType))
	fmt.Println(GenerateReadObjectType(positionType))
	fmt.Println(GenerateWriteObjectType(positionType))
	fmt.Println(GenerateReadListType(ListType{Type:PrimitiveType{Name:"string"}}))
	fmt.Println(GenerateReadListType(ListType{Type:ObjectType{"WorkerAttributeSet"}}))


	//fmt.Println("package main")
	//fmt.Print(GenerateStruct(coordinatesType))
	//fmt.Println()
	//fmt.Print(GenerateStruct(positionType))
	//fmt.Println()
	//fmt.Print(GenerateUpdateStruct(positionType))
	//fmt.Print(GenerateUpdateStruct(coordinatesType))
	//fmt.Println()
	//fmt.Print(GenerateUpdateReader(positionType))
	//fmt.Print(GenerateUpdateReader(coordinatesType))
	//fmt.Print(GenerateUpdateWriter(positionType))
	//fmt.Print(GenerateUpdateWriter(coordinatesType))
	//fmt.Println()
	//fmt.Print(GenerateApplyUpdate(positionType))
	//fmt.Println()
	//fmt.Print(GenerateReader(coordinatesType))
	//fmt.Println()
	//fmt.Print(GenerateWriter(coordinatesType))
	//fmt.Println()
	//fmt.Print(GenerateReader(positionType))
	//fmt.Println()
	//fmt.Print(GenerateWriter(positionType))
}
