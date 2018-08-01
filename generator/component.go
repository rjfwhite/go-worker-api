package main

import "fmt"

type PositionAddCallback func (entity_id int64, position string)
type PositionUpdateCallback func (entity_id int64, position string)
type PositionRemovedCallback func (entity_id int64)

func GenerateComponentEventCallbacks(t ComponentType) string {
	output := ""
	output += fmt.Sprintf("type %sAddedCallback func(entity_id int64, data %s)\n", t.Name, GoTypeFor(ObjectType{t.Name}))
	output += fmt.Sprintf("type %sUpdatedCallback func(entity_id int64, update %s)\n", t.Name, GoTypeFor(ObjectType{t.Name}) + "Update")
	output += fmt.Sprintf("type %sRemovedCallback func(entity_id int64)\n", t.Name)
	return output
}

func GenerateComponentType(t ComponentType) string {
	output := ""
	output += fmt.Sprintf("type %s struct {\n", t.Name)
	for _, f := range t.Data.Fields {
		output += fmt.Sprintf("\t%s %s\n", f.Name, GoTypeFor(f.Type))
	}
	output += "}\n"
	return output
}

func GenerateComponentUpdateType(t ComponentType) string {
	output := ""
	output += fmt.Sprintf("type %sUpdate struct {\n", t.Name)
	for _, f := range t.Data.Fields {
		output += fmt.Sprintf("\t%s %s\n", f.Name, GoTypeFor(OptionType{f.Type}))
	}
	output += "}\n"
	return output
}

func GenerateReadComponentType(t ComponentType) string {
	output := ""
	output += fmt.Sprintf("func ReadComponent_%s(object example.Schema_Object) %s {\n", t.Name, t.Name)
	output += fmt.Sprintf("\treturn %s {\n", t.Name)
	for _, f := range t.Data.Fields {
		output += fmt.Sprintf("\t\t%s : Read%s(object, %d, 0),\n", f.Name, MethodSuffixForType(f.Type), f.Id)
	}
	output += "\t}\n"
	output += "}\n"
	return output
}

func GenerateWriteComponentType(t ComponentType) string {
	output := ""
	output += fmt.Sprintf("func WriteComponent_%s(object example.Schema_Object, value %s) {\n", t.Name, t.Name)
	for _, f := range t.Data.Fields {
		output += fmt.Sprintf("\tWrite%s(object, %d, value.%s)\n", MethodSuffixForType(f.Type), f.Id, f.Name)
	}
	output += "}\n"
	return output
}