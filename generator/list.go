package main

import "fmt"

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

func GenerateWriteListType(t ListType) string {
	output := ""
	output += fmt.Sprintf("func Write%s(object example.Schema_Object, field uint, value %s) {\n", MethodSuffixForType(t), GoTypeFor(t))
	output += "\tfor _, i := range(value) {\n"
	output += fmt.Sprintf("\t\tWrite%s(object, field, i)\n", MethodSuffixForType(t.Type))
	output += "\t}\n"
	output += "}\n"
	return output
}
