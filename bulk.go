package main

import "github.com/rjfwhite/go-worker-api/example"

func ReadPrimitive_sint32(object example.Schema_Object, field uint, index uint) int {
	return example.Schema_IndexSint32(object, field, index)
}

func WritePrimitive_sint32(object example.Schema_Object, field uint, index uint, value int) {
	example.Schema_AddSint32(object, field, value)
}

func ReadPrimitive_string(object example.Schema_Object, field uint, index uint) string {
	return string(*example.Schema_IndexBytes(object, field, index))
}

func WritePrimitive_string(object example.Schema_Object, field uint, index uint, value string) {
	v := byte(value[0])
	example.Schema_AddBytes(object, field, &v, uint(len(value)))
}

func ReadPrimitive_int32(object example.Schema_Object, field uint, index uint) int {
	return example.Schema_IndexInt32(object, field, index)
}

func WritePrimitive_int32(object example.Schema_Object, field uint, index uint, value int) {
	example.Schema_AddInt32(object, field, value)
}

func ReadPrimitive_uint64(object example.Schema_Object, field uint, index uint) uint64 {
	return example.Schema_IndexUint64(object, field, index)
}

func WritePrimitive_uint64(object example.Schema_Object, field uint, index uint, value uint64) {
	example.Schema_AddUint64(object, field, value)
}

func ReadPrimitive_int64(object example.Schema_Object, field uint, index uint) int64 {
	return example.Schema_IndexInt64(object, field, index)
}

func WritePrimitive_int64(object example.Schema_Object, field uint, index uint, value int64) {
	example.Schema_AddInt64(object, field, value)
}

func ReadPrimitive_double(object example.Schema_Object, field uint, index uint) float64 {
	return example.Schema_IndexDouble(object, field, index)
}

func WritePrimitive_double(object example.Schema_Object, field uint, index uint, value float64) {
	example.Schema_AddDouble(object, field, value)
}

func ReadPrimitive_sfixed64(object example.Schema_Object, field uint, index uint) int64 {
	return example.Schema_IndexSfixed64(object, field, index)
}

func WritePrimitive_sfixed64(object example.Schema_Object, field uint, index uint, value int64) {
	example.Schema_AddSfixed64(object, field, value)
}

func ReadPrimitive_bool(object example.Schema_Object, field uint, index uint) bool {
	return example.Schema_IndexBool(object, field, index) > 0
}

func WritePrimitive_bool(object example.Schema_Object, field uint, index uint, value bool) {
	byteValue := byte(0)
	if value {
		byteValue = byte(1)
	}
	example.Schema_AddBool(object, field, byteValue)
}

func ReadPrimitive_uint32(object example.Schema_Object, field uint, index uint) uint {
	return example.Schema_IndexUint32(object, field, index)
}

func WritePrimitive_uint32(object example.Schema_Object, field uint, index uint, value uint) {
	example.Schema_AddUint32(object, field, value)
}

func ReadPrimitive_sfixed32(object example.Schema_Object, field uint, index uint) int {
	return example.Schema_IndexSfixed32(object, field, index)
}

func WritePrimitive_sfixed32(object example.Schema_Object, field uint, index uint, value int) {
	example.Schema_AddSfixed32(object, field, value)
}

func ReadPrimitive_fixed64(object example.Schema_Object, field uint, index uint) uint64 {
	return example.Schema_IndexFixed64(object, field, index)
}

func WritePrimitive_fixed64(object example.Schema_Object, field uint, index uint, value uint64) {
	example.Schema_AddFixed64(object, field, value)
}

func ReadPrimitive_float(object example.Schema_Object, field uint, index uint) float32 {
	return example.Schema_IndexFloat(object, field, index)
}

func WritePrimitive_float(object example.Schema_Object, field uint, index uint, value float32) {
	example.Schema_AddFloat(object, field, value)
}

func ReadPrimitive_EntityId(object example.Schema_Object, field uint, index uint) int64 {
	return example.Schema_IndexEntityId(object, field, index)
}

func WritePrimitive_EntityId(object example.Schema_Object, field uint, index uint, value int64) {
	example.Schema_AddEntityId(object, field, value)
}

func ReadPrimitive_bytes(object example.Schema_Object, field uint, index uint) []byte {
	return []byte{*example.Schema_IndexBytes(object, field, index)}
}

func WritePrimitive_bytes(object example.Schema_Object, field uint, index uint, value []byte) {
	example.Schema_AddBytes(object, field, &value[0], uint(len(value)))
}

func ReadPrimitive_sint64(object example.Schema_Object, field uint, index uint) int64 {
	return example.Schema_IndexSint64(object, field, index)
}

func WritePrimitive_sint64(object example.Schema_Object, field uint, index uint, value int64) {
	example.Schema_AddSint64(object, field, value)
}

func ReadPrimitive_fixed32(object example.Schema_Object, field uint, index uint) uint {
	return example.Schema_IndexFixed32(object, field, index)
}

func WritePrimitive_fixed32(object example.Schema_Object, field uint, index uint, value uint) {
	example.Schema_AddFixed32(object, field, value)
}