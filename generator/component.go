package main

import "fmt"

func GenerateComponentType(t ComponentType) string {
	output := ""
	output += fmt.Sprintf("type %s struct {\n", t.Name)
	for _, f := range t.Fields {
		output += fmt.Sprintf("\t%s %s\n", f.Name, GoTypeFor(f.Type))
	}
	output += "}\n\n"
	return output
}

func GenerateReadComponentType(t ComponentType) string {
	output := ""
	output += fmt.Sprintf("func ReadComponent_%s(object swig.Schema_Object) %s {\n", t.Name, t.Name)
	output += fmt.Sprintf("\treturn %s {\n", t.Name)
	for _, f := range t.Fields {
		output += fmt.Sprintf("\t\t%s : Read%s(object, %d, 0),\n", f.Name, MethodSuffixForType(f.Type), f.Id)
	}
	output += "\t}\n"
	output += "}\n\n"
	return output
}

func GenerateWriteComponentType(t ComponentType) string {
	output := ""
	output += fmt.Sprintf("func WriteComponent_%s(object swig.Schema_Object, value %s) {\n", t.Name, t.Name)
	for _, f := range t.Fields {
		output += fmt.Sprintf("\tWrite%s(object, %d, value.%s)\n", MethodSuffixForType(f.Type), f.Id, f.Name)
	}
	output += "}\n\n"
	return output
}
