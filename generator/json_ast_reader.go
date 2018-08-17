package main

import (
	"path/filepath"
	"os"
	"fmt"
	"strings"
	"github.com/Jeffail/gabs"
	"io/ioutil"
)


func EnumerateComponentsToDataTypes(jsons []*gabs.Container) map[string]string {
	result := map[string]string{}

	for _, json := range jsons {
		components, _ := json.Path("componentDefinitions").Children()
		for _, component := range components {
			component_name := component.Path("qualifiedName").String()
			component_data_type := component.Path("dataDefinition.userType").String()
			result[component_name] = component_data_type
		}
	}

	return result
}

type ParsedType struct {
	Name string
	Fields []SchemaField
	Id int
}

func ParseField(json gabs.Container) SchemaField {
	field := SchemaField{}
	field.Name = json.Path("name").Data().(string)
	field.Id = int(json.Path("number").Data().(float64))
	field.Type = ParseType(json)
	return field
}

func ParseType(json gabs.Container) interface{} {
	if json.Exists("singularType") {
		singular_json := json.Path("singularType")
		if singular_json.Exists("userType") {
			return TypeRef{singular_json.Path("userType").Data().(string)}
		} else if singular_json.Exists("builtInType") {
			return PrimitiveType{singular_json.Path("builtInType").Data().(string)}
		} else {
			return TypeRef{"UNKNOWN"}
		}
	}
	return nil
}

func EnumerateTypeDefinitions(type_json gabs.Container, type_list *map[string]ParsedType) {

	parsed_type := ParsedType{}
	field_definitions, _ := type_json.Path("fieldDefinitions").Children()
	fields := []SchemaField{}
	for _, field_definition := range(field_definitions) {
		fields = append(fields, ParseField(*field_definition))
	}
	parsed_type.Name = type_json.Path("qualifiedName").String()
	parsed_type.Fields = fields

	(*type_list)[parsed_type.Name] = parsed_type

	type_definitions, _ := type_json.Path("typeDefinitions").Children()
	for _, type_definition := range type_definitions {
		EnumerateTypeDefinitions(*type_definition, type_list)
	}
}

type SchemaJsons struct {
	Files []*gabs.Container
}

func EnumerateTypeDefinitionsInJson(schema SchemaJsons) map[string]ParsedType {
	result := map[string]ParsedType{}
	for _, json := range schema.Files {
		type_definitions, _ := json.Path("typeDefinitions").Children()
		for _, type_definition := range type_definitions {
			EnumerateTypeDefinitions(*type_definition, &result)
		}
	}
	return result
}

func TranslateFiles() {
	searchDir := "schema_json"

	jsons := []*gabs.Container{}
	filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".json") {
			data, _ := ioutil.ReadFile(path)
			json, _ := gabs.ParseJSON(data)
			jsons = append(jsons, json)
		}
		return nil
	})

	schema := SchemaJsons{jsons}

	type_name_to_type := EnumerateTypeDefinitionsInJson(schema)
	component_name_to_type_name := EnumerateComponentsToDataTypes(jsons)

	for _, defined_type := range type_name_to_type {
		fmt.Println(defined_type)
	}

	for component, data := range component_name_to_type_name {
		fmt.Printf("%s : %s\n", component, data)
	}
}
