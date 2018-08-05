package main

import "fmt"

func GenerateComponentUpdateType(t ComponentType) string {
	output := ""
	output += fmt.Sprintf("type %sUpdate struct {\n", t.Name)
	for _, f := range t.Fields {
		output += fmt.Sprintf("\t%s %s\n", f.Name, GoTypeFor(OptionType{f.Type}))
	}
	output += "}\n"
	return output
}

func GenerateReadComponentUpdateType(t ComponentType) string {
	output := ""
	output += fmt.Sprintf("func ReadComponentUpdate_%s(object example.Schema_Object) %s {\n", t.Name, t.Name + "Update")
	output += fmt.Sprintf("\treturn %sUpdate {\n", t.Name)
	for _, f := range t.Fields {
		output += fmt.Sprintf("\t\t%s : Read%s(object, %d, 0),\n", f.Name, MethodSuffixForType(OptionType{f.Type}), f.Id)
	}
	output += "\t}\n"
	output += "}\n"
	return output
}

func GenerateWriteComponentUpdateType(t ComponentType) string {
	output := ""
	output += fmt.Sprintf("func WriteComponentUpdate_%s(object example.Schema_Object, value %s) {\n", t.Name, t.Name + "Update")
	for _, f := range t.Fields {
		output += fmt.Sprintf("\tWrite%s(object, %d, value.%s)\n", MethodSuffixForType(OptionType{f.Type}), f.Id, f.Name)
	}
	output += "}\n"
	return output
}