package main

import "fmt"

func GenerateComponentType(t ComponentDefinition) string {
	output := ""
	output += fmt.Sprintf("type %s struct {\n", t.Type.Name)
	for _, f := range t.Fields {
		output += fmt.Sprintf("\t%s %s\n", f.Name, GoTypeFor(f.Type))
	}
	output += "}\n"
	return output
}

func GenerateReadComponentType(t ComponentDefinition) string {
	output := ""
	output += fmt.Sprintf("func ReadComponent_%s(object example.Schema_Object) %s {\n", t.Type.Name, t.Type.Name)
	output += fmt.Sprintf("\treturn %s {\n", t.Type.Name)
	for _, f := range t.Fields {
		output += fmt.Sprintf("\t\t%s : Read%s(object, %d, 0),\n", f.Name, MethodSuffixForType(f.Type), f.Id)
	}
	output += "\t}\n"
	output += "}\n"
	return output
}

func GenerateWriteComponentType(t ComponentDefinition) string {
	output := ""
	output += fmt.Sprintf("func WriteComponent_%s(object example.Schema_Object, value %s) {\n", t.Type.Name, t.Type.Name)
	for _, f := range t.Fields {
		output += fmt.Sprintf("\tWrite%s(object, %d, value.%s)\n", MethodSuffixForType(f.Type), f.Id, f.Name)
	}
	output += "}\n"
	return output
}
