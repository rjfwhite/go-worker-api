package main

import "github.com/rjfwhite/go-worker-api/example"

type WorkerAttributeSet struct {
	attribute []string
}
type WorkerRequirementSet struct {
	attribute_set []WorkerAttributeSet
}
type Coordinates struct {
	X float64
	Y float64
	Z float64
}
type Position struct {
	Coords Coordinates
}

func ReadObject_WorkerAttributeSet(object example.Schema_Object, field uint, index uint) WorkerAttributeSet {
	innerObject := example.Schema_IndexObject(object, field, index)
	return WorkerAttributeSet{
		attribute: ReadList_Primitive_string(innerObject, 1, 0),
	}
}

func ReadObject_WorkerRequirementSet(object example.Schema_Object, field uint, index uint) WorkerRequirementSet {
	innerObject := example.Schema_IndexObject(object, field, index)
	return WorkerRequirementSet{
		attribute_set: ReadList_Object_WorkerAttributeSet(innerObject, 1, 0),
	}
}

func ReadObject_Coordinates(object example.Schema_Object, field uint, index uint) Coordinates {
	innerObject := example.Schema_IndexObject(object, field, index)
	return Coordinates{
		X: ReadPrimitive_double(innerObject, 1, 0),
		Y: ReadPrimitive_double(innerObject, 2, 0),
		Z: ReadPrimitive_double(innerObject, 3, 0),
	}
}

func WriteObject_Coordinates(object example.Schema_Object, field uint, value Coordinates) {
	example.Schema_AddObject(object, field)
	innerObject := example.Schema_GetObject(object, field)
	WritePrimitive_double(innerObject, 1, value.X)
	WritePrimitive_double(innerObject, 2, value.Y)
	WritePrimitive_double(innerObject, 3, value.Z)
}

func ReadObject_Position(object example.Schema_Object, field uint, index uint) Position {
	innerObject := example.Schema_IndexObject(object, field, index)
	return Position{
		Coords: ReadObject_Coordinates(innerObject, 1, 0),
	}
}

func WriteObject_Position(object example.Schema_Object, field uint, value Position) {
	example.Schema_AddObject(object, field)
	innerObject := example.Schema_GetObject(object, field)
	WriteObject_Coordinates(innerObject, 1, value.Coords)
}

func ReadList_Primitive_string(object example.Schema_Object, field uint, index uint) []string {
	count := example.Schema_GetBytesCount(object, field)
	result := []string{}
	for i := uint(0); i < count; i++ {
		result = append(result, ReadPrimitive_string(object, field, i))
	}
	return result
}

func ReadList_Object_WorkerAttributeSet(object example.Schema_Object, field uint, index uint) []WorkerAttributeSet {
	count := example.Schema_GetObjectCount(object, field)
	result := []WorkerAttributeSet{}
	for i := uint(0); i < count; i++ {
		result = append(result, ReadObject_WorkerAttributeSet(object, field, i))
	}
	return result
}

//Needs generating
func ReadMap_Primitive_uint32_Object_WorkerAttributeSet(object example.Schema_Object, field uint, index uint) map[uint][]WorkerAttributeSet {
	count := example.Schema_GetObjectCount(object, field)
	result := map[uint][]WorkerAttributeSet{}
	for i := uint(0); i < count; i++ {
		innerObject := example.Schema_IndexObject(object, field, i)
		key := ReadPrimitive_uint32(innerObject, 1, 0)
		value := ReadList_Object_WorkerAttributeSet(innerObject, 2, 0)
		result[key] = value
	}
	return result
}
