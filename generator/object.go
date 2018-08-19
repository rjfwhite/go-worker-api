package main

import "fmt"

func GenerateObjectType(t ObjectType) string {
	output := ""
	output += fmt.Sprintf("type %s struct {\n", t.Name)
	for _, f := range t.Fields {
		output += fmt.Sprintf("\t%s %s\n", f.Name, GoTypeFor(f.Type))
	}
	output += "}\n\n"
	return output
}

func GenerateReadObjectType(t ObjectType) string {
	output := ""
	output += fmt.Sprintf("func ReadObject_%s(object swig.Schema_Object, field uint, index uint) %s {\n", t.Name, t.Name)
	output += fmt.Sprintf("\tinnerObject := swig.Schema_IndexObject(object, field, index)\n")
	output += fmt.Sprintf("\treturn %s {\n", t.Name)
	for _, f := range t.Fields {
		output += fmt.Sprintf("\t\t%s : Read%s(innerObject, %d, 0),\n", f.Name, MethodSuffixForType(f.Type), f.Id)
	}
	output += "\t}\n"
	output += "}\n\n"
	return output
}

func GenerateWriteObjectType(t ObjectType) string {
	output := ""
	output += fmt.Sprintf("func WriteObject_%s(object swig.Schema_Object, field uint, value %s) {\n", t.Name, t.Name)
	output += "\tinnerObject := swig.Schema_AddObject(object, field)\n"
	for _, f := range t.Fields {
		output += fmt.Sprintf("\tWrite%s(innerObject, %d, value.%s)\n", MethodSuffixForType(f.Type), f.Id, f.Name)
	}
	output += "}\n\n"
	return output
}