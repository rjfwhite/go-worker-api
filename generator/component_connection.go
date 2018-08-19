package main

import "fmt"

func GenerateUpdateComponentConnectionMethod(t ComponentType) string {
	output := ""
	output += fmt.Sprintf("func (connection Connection) Send%sUpdate(entity_id int64, value %sUpdate) {\n", t.Name, t.Name)
	output += fmt.Sprintf("\tcomponent_id := uint(%d)\n", t.Id)
	output += "\tcomponent_update := swig.Schema_CreateComponentUpdate(component_id)\n"
	output += "\tcomponent_update_fields := swig.Schema_GetComponentUpdateFields(component_update)\n"
	output += fmt.Sprintf("\tWriteComponentUpdate_%s(component_update_fields, value)\n", t.Name)
	output += "\tconnection.SendComponentUpdate(entity_id, component_id, component_update)\n"
	output += "}\n"
	return output
}
