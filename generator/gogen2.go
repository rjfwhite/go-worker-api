package main

import "fmt"

type SchemaType struct {
	Package string
	Name    string
	Fields  []SchemaField
}

type SchemaField struct {
	Name string
	Type string
	Id   int
}

func GenerateStruct(t SchemaType) string {
	output := ""
	output += fmt.Sprintf("type %s struct {\n", t.Name)

	for _, f := range t.Fields {
		output += fmt.Sprintf("\t%s %s\n", f.Name, f.Type)
	}

	output += "}\n"
	return output
}

func GenerateUpdateStruct(t SchemaType) string {
	output := ""
	output += fmt.Sprintf("type %sUpdate struct {\n", t.Name)

	for _, f := range t.Fields {
		output += fmt.Sprintf("\t%s *%s\n", f.Name, f.Type)
	}

	output += "}\n"
	return output
}

func GenerateWriter(t SchemaType) string {
	output := ""
	output += fmt.Sprintf("func Write%s(output example.Schema_Object, value %s) {\n", t.Name, t.Name)
	for _, f := range t.Fields {
		if f.Type == "float64" {
			output += fmt.Sprintf("\texample.Schema_AddDouble(output, %d, value.%s)\n", f.Id, f.Name)
		} else {
			output += fmt.Sprintf("\texample.Schema_AddObject(output, %d)\n", f.Id)
			output += fmt.Sprintf("\tWrite%s(example.Schema_GetObject(output, %d), value.%s)\n", f.Type, f.Id, f.Name)
		}
	}
	output += "}\n"
	return output
}

func GenerateReader(t SchemaType) string {
	output := ""
	output += fmt.Sprintf("func Read%s(input example.Schema_Object) %s {\n", t.Name, t.Name)
	output += fmt.Sprintf("\treturn %s {\n", t.Name)
	for _, f := range t.Fields {
		if f.Type == "float64" {
			output += fmt.Sprintf("\t\t%s : example.Schema_GetDouble(input, %d),\n", f.Name, f.Id)
		} else {
			output += fmt.Sprintf("\t\t%s : Read%s(example.Schema_GetObject(input, %d)),\n", f.Name, f.Type, f.Id)
		}
	}
	output += "\t}\n"
	output += "}\n"
	return output
}

func main() {
	positionFields := []SchemaField{{Name: "Coords", Type: "Coordinates", Id: 1}}
	positionType := SchemaType{Package: "", Name: "Position", Fields: positionFields}

	coordinatesFields := []SchemaField{{Name: "X", Type: "float64", Id: 1}, {Name: "Y", Type: "float64", Id: 2}, {Name: "Z", Type: "float64", Id: 3}}
	coordinatesType := SchemaType{Package: "", Name: "Coordinates", Fields: coordinatesFields}

	fmt.Print(GenerateStruct(coordinatesType))
	fmt.Print(GenerateStruct(positionType))
	fmt.Print(GenerateUpdateStruct(positionType))
	fmt.Print(GenerateReader(coordinatesType))
	fmt.Print(GenerateWriter(coordinatesType))
	fmt.Print(GenerateReader(positionType))
	fmt.Print(GenerateWriter(positionType))
}
