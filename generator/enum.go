package main

import (
	"fmt"
)

func GenerateEnumType(t EnumType) string {
	output := ""
	output += fmt.Sprintf("type %s uint\n\n", t.Name)
	output += "const (\n"
	for number, value := range (t.values) {
		output += fmt.Sprintf("\t%s %s = %d\n", value, t.Name, number)
	}
	output += ")\n"
	return output
}

func GenerateReadEnumType(t EnumType) string {
	output := ""
	output += fmt.Sprintf("func Read%s(object example.Schema_Object, field uint, index uint) %s {\n", MethodSuffixForType(t), GoTypeFor(t))
	output += fmt.Sprintf("\treturn %s(example.Schema_IndexEnum(object, field, index))\n", t.Name)
	output += "}\n"
	return output
}

func GenerateWriteEnumType(t EnumType) string {
	output := ""
	output += fmt.Sprintf("func Write%s(object example.Schema_Object, field uint, value %s) {\n", MethodSuffixForType(t), GoTypeFor(t))
	output += "\texample.Schema_AddEnum(object, field, uint(value))\n"
	output += "}\n"
	return output
}