package main

import "fmt"

/*
func (dispatcher *Dispatcher) OnPositionAdded(callback PositionAddedCallback) {
	component_id := uint(54)
	innerCallback := func(entity_id int64, component_data example.Worker_ComponentData) {
		dataFields := example.Schema_GetComponentDataFields(component_data.GetSchema_type())
		component := ReadComponent_Position(dataFields)
		callback(entity_id, component)
	}
	dispatcher.ComponentAddedCallbacks[component_id] = []ComponentAddedCallback{innerCallback}
}
}*/

func GenerateAddComponentDispatcherMethod(t ComponentType) string {
	output := ""
	output += fmt.Sprintf("func (dispatcher *Dispatcher) On%sAdded(callback %sAddedCallback) {\n", t.Name, t.Name)
	output += fmt.Sprintf("\tcomponent_id := uint(%d)\n", t.Id)
	output += "\tinnerCallback := func(entity_id int64, component_data example.Worker_ComponentData) {\n"
	output += "\t\tdataFields := example.Schema_GetComponentDataFields(component_data.GetSchema_type())\n"
	output += fmt.Sprintf("\t\tcomponent := ReadComponent_%s(dataFields)\n", t.Name)
	output += "\t\tcallback(entity_id, component)\n"
	output += "\t}\n"
	output += "\tdispatcher.OnComponentAdded(component_id, innerCallback)\n"
	output += "}\n"
	return output
}