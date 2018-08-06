package main

import (
	"fmt"
)

func GenerateEnumType(t EnumDefinition) string {
	output := ""
	output += fmt.Sprintf("type %s uint\n\n", t.Type.Name)
	output += "const (\n"
	for number, value := range (t.values) {
		output += fmt.Sprintf("\t%s %s = %d\n", value, t.Type.Name, number)
	}
	output += ")\n"
	return output
}

func GenerateReadEnumType(t EnumDefinition) string {
	output := ""
	output += fmt.Sprintf("func Read%s(object example.Schema_Object, field uint, index uint) %s {\n", MethodSuffixForType(t.Type), GoTypeFor(t.Type))
	output += fmt.Sprintf("\treturn %s(example.Schema_IndexEnum(object, field, index))\n", t.Type.Name)
	output += "}\n"
	return output
}

func GenerateWriteEnumType(t EnumDefinition) string {
	output := ""
	output += fmt.Sprintf("func Write%s(object example.Schema_Object, field uint, value %s) {\n", MethodSuffixForType(t.Type), GoTypeFor(t.Type))
	output += "\texample.Schema_AddEnum(object, field, uint(value))\n"
	output += "}\n"
	return output
}
