package main

import "fmt"

func GenerateReadOptionType(t OptionType) string {
	output := ""
	output += fmt.Sprintf("func Read%s(object swig.Schema_Object, field uint, index uint) %s {\n", MethodSuffixForType(t), GoTypeFor(t))
	output += fmt.Sprintf("\tif swig.Schema_Get%sCount(object, field) > 0 {\n", FunctionFamilyFor(t.Type))
	output += fmt.Sprintf("\t\tresult := Read%s(object, field, index)\n", MethodSuffixForType(t.Type))
	output += "\t\treturn &result\n"
	output += "\t}\n"
	output += "\treturn nil\n"
	output += "}\n\n"
	return output
}

func GenerateWriteOptionType(t OptionType) string {
	output := ""
	output += fmt.Sprintf("func Write%s(object swig.Schema_Object, field uint, value %s) {\n", MethodSuffixForType(t), GoTypeFor(t))
	output += "\tif value != nil {\n"
	output += fmt.Sprintf("\t\tWrite%s(object, field, *value)\n", MethodSuffixForType(t.Type))
	output += "\t}\n"
	output += "}\n\n"
	return output
}
