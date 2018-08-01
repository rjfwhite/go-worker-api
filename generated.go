package main

import "github.com/rjfwhite/go-worker-api/example"

type WorkerAttributeSet  struct {
	attribute []string
}
type WorkerRequirementSet  struct {
	attribute_set []WorkerAttributeSet
}
type Coordinates struct {
	X float64
	Y float64
	Z float64
}
func ReadObject_WorkerAttributeSet (object example.Schema_Object, field uint, index uint) WorkerAttributeSet  {
	innerObject := example.Schema_IndexObject(object, field, index)
	return WorkerAttributeSet  {
		attribute : ReadList_Primitive_string(innerObject, 1, 0),
	}
}

func WriteObject_WorkerAttributeSet (object example.Schema_Object, field uint, value WorkerAttributeSet ) {
	innerObject := example.Schema_AddObject(object, field)
	WriteList_Primitive_string(innerObject, 1, value.attribute)
}

func ReadObject_WorkerRequirementSet (object example.Schema_Object, field uint, index uint) WorkerRequirementSet  {
	innerObject := example.Schema_IndexObject(object, field, index)
	return WorkerRequirementSet  {
		attribute_set : ReadList_Object_WorkerAttributeSet(innerObject, 1, 0),
	}
}

func WriteObject_WorkerRequirementSet (object example.Schema_Object, field uint, value WorkerRequirementSet ) {
	innerObject := example.Schema_AddObject(object, field)
	WriteList_Object_WorkerAttributeSet(innerObject, 1, value.attribute_set)
}

func ReadObject_Coordinates(object example.Schema_Object, field uint, index uint) Coordinates {
	innerObject := example.Schema_IndexObject(object, field, index)
	return Coordinates {
		X : ReadPrimitive_double(innerObject, 1, 0),
		Y : ReadPrimitive_double(innerObject, 2, 0),
		Z : ReadPrimitive_double(innerObject, 3, 0),
	}
}

func WriteObject_Coordinates(object example.Schema_Object, field uint, value Coordinates) {
	innerObject := example.Schema_AddObject(object, field)
	WritePrimitive_double(innerObject, 1, value.X)
	WritePrimitive_double(innerObject, 2, value.Y)
	WritePrimitive_double(innerObject, 3, value.Z)
}

func ReadObject_Position(object example.Schema_Object, field uint, index uint) Position {
	innerObject := example.Schema_IndexObject(object, field, index)
	return Position {
		Coords : ReadObject_Coordinates(innerObject, 1, 0),
	}
}

func WriteObject_Position(object example.Schema_Object, field uint, value Position) {
	innerObject := example.Schema_AddObject(object, field)
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

func WriteList_Primitive_string(object example.Schema_Object, field uint, value []string) {
	for _, i := range(value) {
		WritePrimitive_string(object, field, i)
	}
}

func ReadList_Object_WorkerAttributeSet(object example.Schema_Object, field uint, index uint) []WorkerAttributeSet {
	count := example.Schema_GetObjectCount(object, field)
	result := []WorkerAttributeSet{}
	for i := uint(0); i < count; i++ {
		result = append(result, ReadObject_WorkerAttributeSet(object, field, i))
	}
	return result
}

func WriteList_Object_WorkerAttributeSet(object example.Schema_Object, field uint, value []WorkerAttributeSet) {
	for _, i := range(value) {
		WriteObject_WorkerAttributeSet(object, field, i)
	}
}

func ReadMap_Primitive_uint32_to_List_Object_WorkerAttributeSet(object example.Schema_Object, field uint, index uint) map[uint][]WorkerAttributeSet {
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

func WriteMap_Primitive_uint32_to_List_Object_WorkerAttributeSet(object example.Schema_Object, field uint, index uint, value map[uint][]WorkerAttributeSet) {
	for k, v := range(value) {
		innerObject := example.Schema_AddObject(object, field)
		WritePrimitive_uint32(innerObject, 1, k)
		WriteList_Object_WorkerAttributeSet(innerObject, 2, v)
	}
}

func ReadOption_Object_WorkerAttributeSet(object example.Schema_Object, field uint, index uint) *WorkerAttributeSet {
	if example.Schema_GetObjectCount(object, field) > 0 {
		result := ReadObject_WorkerAttributeSet(object, field, index)
		return &result
	}
	return nil
}

func WriteOption_Object_WorkerAttributeSet(object example.Schema_Object, field uint, value *WorkerAttributeSet) {
	if value != nil {
		WriteObject_WorkerAttributeSet(object, field, *value)
	}
}

type Color uint

const (
	Red Color = 2
	Blue Color = 1
)

func ReadEnum_Color(object example.Schema_Object, field uint, index uint) Color {
	return Color(example.Schema_IndexEnum(object, field, index))
}

func WriteEnum_Color(object example.Schema_Object, field uint, value Color) {
	example.Schema_AddEnum(object, field, uint(value))
}

func ReadList_Enum_Color(object example.Schema_Object, field uint, index uint) []Color {
	count := example.Schema_GetEnumCount(object, field)
	result := []Color{}
	for i := uint(0); i < count; i++ {
		result = append(result, ReadEnum_Color(object, field, i))
	}
	return result
}

func WriteList_Enum_Color(object example.Schema_Object, field uint, value []Color) {
	for _, i := range(value) {
		WriteEnum_Color(object, field, i)
	}
}

type Position struct {
	Coords Coordinates
}

type PositionUpdate struct {
	Coords *Coordinates
}

func ReadComponent_Position(object example.Schema_Object) Position {
	return Position {
		Coords : ReadObject_Coordinates(object, 1, 0),
	}
}

func WriteComponent_Position(object example.Schema_Object, value Position) {
	WriteObject_Coordinates(object, 1, value.Coords)
}

func ReadComponentUpdate_Position(object example.Schema_Object) PositionUpdate {
	return PositionUpdate {
		Coords : ReadOption_Object_Coordinates(object, 1, 0),
	}
}

func WriteComponentUpdate_Position(object example.Schema_Object, value PositionUpdate) {
	WriteOption_Object_Coordinates(object, 1, value.Coords)
}

type PositionAddedCallback func(entity_id int64, data Position)
type PositionUpdatedCallback func(entity_id int64, update PositionUpdate)
type PositionRemovedCallback func(entity_id int64)

func ReadOption_Object_Coordinates(object example.Schema_Object, field uint, index uint) *Coordinates {
	if example.Schema_GetObjectCount(object, field) > 0 {
		result := ReadObject_Coordinates(object, field, index)
		return &result
	}
	return nil
}

func WriteOption_Object_Coordinates(object example.Schema_Object, field uint, value *Coordinates) {
	if value != nil {
		WriteObject_Coordinates(object, field, *value)
	}
}