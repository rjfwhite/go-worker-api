package main

import "github.com/rjfwhite/go-worker-api/example"

func ReadPrimitive_sint32(object example.Schema_Object, index uint) int {
	return example.Schema_GetSint32(object, index)
}

func WritePrimitive_sint32(object example.Schema_Object, index uint, value int) {
	example.Schema_AddSint32(object, index, value)
}

func ReadPrimitive_string(object example.Schema_Object, index uint) string {
	return string(*example.Schema_GetBytes(object, index))
}

func WritePrimitive_string(object example.Schema_Object, index uint, value string) {
	v := byte(value[0])
	example.Schema_AddBytes(object, index, &v, uint(len(value)))
}

func ReadPrimitive_int32(object example.Schema_Object, index uint) int {
	return example.Schema_GetInt32(object, index)
}

func WritePrimitive_int32(object example.Schema_Object, index uint, value int) {
	example.Schema_AddInt32(object, index, value)
}

func ReadPrimitive_uint64(object example.Schema_Object, index uint) uint64 {
	return example.Schema_GetUint64(object, index)
}

func WritePrimitive_uint64(object example.Schema_Object, index uint, value uint64) {
	example.Schema_AddUint64(object, index, value)
}

func ReadPrimitive_int64(object example.Schema_Object, index uint) int64 {
	return example.Schema_GetInt64(object, index)
}

func WritePrimitive_int64(object example.Schema_Object, index uint, value int64) {
	example.Schema_AddInt64(object, index, value)
}

func ReadPrimitive_double(object example.Schema_Object, index uint) float64 {
	return example.Schema_GetDouble(object, index)
}

func WritePrimitive_double(object example.Schema_Object, index uint, value float64) {
	example.Schema_AddDouble(object, index, value)
}

func ReadPrimitive_sfixed64(object example.Schema_Object, index uint) int64 {
	return example.Schema_GetSfixed64(object, index)
}

func WritePrimitive_sfixed64(object example.Schema_Object, index uint, value int64) {
	example.Schema_AddSfixed64(object, index, value)
}

func ReadPrimitive_bool(object example.Schema_Object, index uint) bool {
	return example.Schema_GetBool(object, index) > 0
}

func WritePrimitive_bool(object example.Schema_Object, index uint, value bool) {
	byteValue := byte(0)
	if value {
		byteValue = byte(1)
	}
	example.Schema_AddBool(object, index, byteValue)
}

func ReadPrimitive_uint32(object example.Schema_Object, index uint) uint {
	return example.Schema_GetUint32(object, index)
}

func WritePrimitive_uint32(object example.Schema_Object, index uint, value uint) {
	example.Schema_AddUint32(object, index, value)
}

func ReadPrimitive_sfixed32(object example.Schema_Object, index uint) int {
	return example.Schema_GetSfixed32(object, index)
}

func WritePrimitive_sfixed32(object example.Schema_Object, index uint, value int) {
	example.Schema_AddSfixed32(object, index, value)
}

func ReadPrimitive_fixed64(object example.Schema_Object, index uint) uint64 {
	return example.Schema_GetFixed64(object, index)
}

func WritePrimitive_fixed64(object example.Schema_Object, index uint, value uint64) {
	example.Schema_AddFixed64(object, index, value)
}

func ReadPrimitive_float(object example.Schema_Object, index uint) float32 {
	return example.Schema_GetFloat(object, index)
}

func WritePrimitive_float(object example.Schema_Object, index uint, value float32) {
	example.Schema_AddFloat(object, index, value)
}

func ReadPrimitive_EntityId(object example.Schema_Object, index uint) int64 {
	return example.Schema_GetEntityId(object, index)
}

func WritePrimitive_EntityId(object example.Schema_Object, index uint, value int64) {
	example.Schema_AddEntityId(object, index, value)
}

func ReadPrimitive_bytes(object example.Schema_Object, index uint) []byte {
	return []byte{*example.Schema_GetBytes(object, index)}
}

func WritePrimitive_bytes(object example.Schema_Object, index uint, value []byte) {
	example.Schema_AddBytes(object, index, &value[0], uint(len(value)))
}

func ReadPrimitive_sint64(object example.Schema_Object, index uint) int64 {
	return example.Schema_GetSint64(object, index)
}

func WritePrimitive_sint64(object example.Schema_Object, index uint, value int64) {
	example.Schema_AddSint64(object, index, value)
}

func ReadPrimitive_fixed32(object example.Schema_Object, index uint) uint {
	return example.Schema_GetFixed32(object, index)
}

func WritePrimitive_fixed32(object example.Schema_Object, index uint, value uint) {
	example.Schema_AddFixed32(object, index, value)
}