package main

import "fmt"

func GenerateReadMapType(t MapType) string {
	output := ""
	output += fmt.Sprintf("func Read%s(object example.Schema_Object, field uint, index uint) %s {\n", MethodSuffixForType(t), GoTypeFor(t))
	output += "\tcount := example.Schema_GetObjectCount(object, field)\n"
	output += fmt.Sprintf("\tresult := %s{}\n", GoTypeFor(t))
	output += "\tfor i := uint(0); i < count; i++ {\n"
	output += "\t\tinnerObject := example.Schema_IndexObject(object, field, i)\n"
	output += fmt.Sprintf("\t\tkey := Read%s(innerObject, 1, 0)\n", MethodSuffixForType(t.KeyType))
	output += fmt.Sprintf("\t\tvalue := Read%s(innerObject, 2, 0)\n", MethodSuffixForType(t.ValueType))
	output += "\t\tresult[key] = value\n"
	output += "\t}\n"
	output += "\treturn result\n"
	output += "}\n"
	return output
}

func GenerateWriteMapType(t MapType) string {
	output := ""
	output += fmt.Sprintf("func Write%s(object example.Schema_Object, field uint, value %s) {\n", MethodSuffixForType(t), GoTypeFor(t))
	output += "\tfor k, v := range(value) {\n"
	output += "\t\tinnerObject := example.Schema_AddObject(object, field)\n"
	output += fmt.Sprintf("\t\tWrite%s(innerObject, 1, k)\n", MethodSuffixForType(t.KeyType))
	output += fmt.Sprintf("\t\tWrite%s(innerObject, 2, v)\n", MethodSuffixForType(t.ValueType))
	output += "\t}\n"
	output += "}\n"
	return output
}
