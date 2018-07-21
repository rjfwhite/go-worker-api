package example

import "fmt"

func MyTypeConvertToUpdate(data MyType) MyTypeUpdate {
	return MyTypeUpdate {
		Age : &data.Age,
		Name : &data.Name,
	}
}
func ApplyMyTypeUpdate(data MyType, update MyTypeUpdate)  {
	if update.Age != nil {
		data.Age = *update.Age
	}
	if update.Name != nil {
		data.Name = *update.Name
	}
}

type MyTypeComponentUpdateHandler interface {

}

type MyType struct {
	Age int
	Name string
}
type MyTypeUpdate struct {
	//Begin Fields
	Age *int
	Name *string
	//End Fields
	//Begin Events
	ThingHappened []string
	//End Events
}

type SchemaType struct {
	Package string
	Name string
	Fields []SchemaField
	Events []SchemaEvent
}

type SchemaField struct {
	Name string
	Type string
	Id int
}

type SchemaEvent struct {
	Name string
	Type string
}

func SchemaTypeToGoData(t SchemaType) string {
	output := ""
	output += fmt.Sprintf("type %s struct {\n", t.Name)

	for _, f := range t.Fields {
		output += fmt.Sprintf("\t%s %s\n", f.Name, f.Type)
	}

	output += "}\n"
	return output
}

func SchemaTypeToGoUpdate(t SchemaType) string {
	output := ""
	output += fmt.Sprintf("type %sUpdate struct {\n", t.Name)

	for _, f := range t.Fields {
		output += fmt.Sprintf("\t%s *%s\n", f.Name, f.Type)
	}

	for _, e := range t.Events {
		output += fmt.Sprintf("\t%s []%s\n", e.Name, e.Type)
	}

	output += "}\n"
	return output
}


func GenerateApplyUpdate(t SchemaType) string {
	output := ""
	output += fmt.Sprintf("func %sApplyUpdate(data %s, update %sUpdate)  {\n", t.Name, t.Name, t.Name)

	for _, f := range t.Fields {
		output += fmt.Sprintf("\tif update.%s != nil {\n", f.Name)
		output += fmt.Sprintf("\t\tdata.%s = *update.%s\n", f.Name, f.Name)
		output += "\t}\n"
	}

	output += "}\n"
	return output
}

func GenerateConvertToUpdate(t SchemaType) string {
	output := ""
	output += fmt.Sprintf("func %sConvertToUpdate(data %s) %sUpdate {\n", t.Name, t.Name, t.Name)

	output += fmt.Sprintf("\treturn %sUpdate {\n", t.Name)

	for _, f := range t.Fields {
		output += fmt.Sprintf("\t\t%s : &data.%s,\n", f.Name, f.Name)
	}

	output += "\t}\n"
	output += "}\n"
	return output
}

func test() {

	events := []SchemaEvent{SchemaEvent{Name:"ThingHappened", Type:"string"}}

	fields := []SchemaField{SchemaField{Name:"Age", Type:"int"}, SchemaField{Name:"Name", Type:"string"}}

	typ := SchemaType{Package:"", Name:"MyType", Fields:fields, Events:events}


	fmt.Print(GenerateApplyUpdate(typ))
	fmt.Print(GenerateConvertToUpdate(typ))
	fmt.Print(SchemaTypeToGoData(typ))
	fmt.Print(SchemaTypeToGoUpdate(typ))
}