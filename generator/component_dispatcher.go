package main

import "fmt"

func GenerateComponentEventCallbacks(t ComponentType) string {
	output := ""
	output += fmt.Sprintf("type %sAddedCallback func(entity_id int64, data %s)\n", t.Name, GoTypeFor(t))
	output += fmt.Sprintf("type %sUpdatedCallback func(entity_id int64, update %s)\n", t.Name, GoTypeFor(t) + "Update")
	return output
}

// Creates a typed wrapper around the normal dispatcher method
// enabling syntax like 'dispatcher.OnFooAdded(callback)'
func GenerateAddComponentDispatcherMethod(t ComponentType) string {
	output := ""
	output += fmt.Sprintf("func (dispatcher *Dispatcher) On%sAdded(callback %sAddedCallback) {\n", t.Name, t.Name)
	output += fmt.Sprintf("\tcomponent_id := uint(%d)\n", t.Id)
	output += "\tinner_callback := func(entity_id int64, component_data swig.Worker_ComponentData) {\n"
	output += "\t\tdata_fields := swig.Schema_GetComponentDataFields(component_data.GetSchema_type())\n"
	output += fmt.Sprintf("\t\tcomponent := ReadComponent_%s(data_fields)\n", t.Name)
	output += "\t\tcallback(entity_id, component)\n"
	output += "\t}\n"
	output += "\tdispatcher.OnComponentAdded(component_id, inner_callback)\n"
	output += "}\n"
	return output
}

func GenerateUpdateComponentDispatcherMethod(t ComponentType) string {
	output := ""
	output += fmt.Sprintf("func (dispatcher *Dispatcher) On%sUpdated(callback %sUpdatedCallback) {\n", t.Name, t.Name)
	output += fmt.Sprintf("\tcomponent_id := uint(%d)\n", t.Id)
	output += "\tinner_callback := func(entity_id int64, component_update swig.Worker_ComponentUpdate) {\n"
	output += "\t\tdata_fields := swig.Schema_GetComponentUpdateFields(component_update.GetSchema_type())\n"
	output += fmt.Sprintf("\t\tcomponent := ReadComponentUpdate_%s(data_fields)\n", t.Name)
	output += "\t\tcallback(entity_id, component)\n"
	output += "\t}\n"
	output += "\tdispatcher.OnComponentUpdated(component_id, inner_callback)\n"
	output += "}\n"
	return output
}

func GenerateAuthorityComponentDispatcherMethod(t ComponentType) string {
	output := ""
	output += fmt.Sprintf("func (dispatcher *Dispatcher) On%sAuthority(callback ComponentAuthorityCallback) {\n", t.Name)
	output += fmt.Sprintf("\tcomponent_id := uint(%d)\n", t.Id)
	output += "\tdispatcher.OnComponentAuthority(component_id, callback)\n"
	output += "}\n"
	return output
}

func GenerateRemoveComponentDispatcherMethod(t ComponentType) string {
	output := ""
	output += fmt.Sprintf("func (dispatcher *Dispatcher) On%sRemoved(callback ComponentRemovedCallback) {\n", t.Name)
	output += fmt.Sprintf("\tcomponent_id := uint(%d)\n", t.Id)
	output += "\tdispatcher.OnComponentRemoved(component_id, callback)\n"
	output += "}\n"
	return output
}