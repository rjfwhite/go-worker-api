package main

import "fmt"

func GenerateComponentUpdateType(t ComponentDefinition) string {
	output := ""
	output += fmt.Sprintf("type %sUpdate struct {\n", t.Type.Name)
	for _, f := range t.Fields {
		output += fmt.Sprintf("\t%s %s\n", f.Name, GoTypeFor(OptionType{f.Type}))
	}
	output += "}\n"
	return output
}

func GenerateReadComponentUpdateType(t ComponentDefinition) string {
	output := ""
	output += fmt.Sprintf("func ReadComponentUpdate_%s(object example.Schema_Object) %s {\n", t.Type.Name, t.Type.Name + "Update")
	output += fmt.Sprintf("\treturn %sUpdate {\n", t.Type.Name)
	for _, f := range t.Fields {
		output += fmt.Sprintf("\t\t%s : Read%s(object, %d, 0),\n", f.Name, MethodSuffixForType(OptionType{f.Type}), f.Id)
	}
	output += "\t}\n"
	output += "}\n"
	return output
}

func GenerateWriteComponentUpdateType(t ComponentDefinition) string {
	output := ""
	output += fmt.Sprintf("func WriteComponentUpdate_%s(object example.Schema_Object, value %s) {\n", t.Type.Name, t.Type.Name + "Update")
	for _, f := range t.Fields {
		output += fmt.Sprintf("\tWrite%s(object, %d, value.%s)\n", MethodSuffixForType(OptionType{f.Type}), f.Id, f.Name)
	}
	output += "}\n"
	return output
}