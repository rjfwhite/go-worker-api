package main

import (
	"github.com/rjfwhite/go-worker-api/swig"
	"unsafe"
)

func ReadPrimitive_sint32(object swig.Schema_Object, field uint, index uint) int {
	return swig.Schema_IndexSint32(object, field, index)
}

func WritePrimitive_sint32(object swig.Schema_Object, field uint, value int) {
	swig.Schema_AddSint32(object, field, value)
}

func ReadPrimitive_string(object swig.Schema_Object, field uint, index uint) string {
	length := swig.Schema_IndexBytesLength(object, field, index)
	unsafePtr := unsafe.Pointer(swig.Schema_IndexBytes(object, field, index))
	bytes := (*(*[4096]byte)(unsafePtr))[0:length]
	return string(bytes)
}

func WritePrimitive_string(object swig.Schema_Object, field uint, value string) {
	bytes := []byte(value)
	swig.Schema_AddBytes(object, field, &(bytes[0]), uint(len(bytes)))
}

func ReadPrimitive_int32(object swig.Schema_Object, field uint, index uint) int {
	return swig.Schema_IndexInt32(object, field, index)
}

func WritePrimitive_int32(object swig.Schema_Object, field uint, value int) {
	swig.Schema_AddInt32(object, field, value)
}

func ReadPrimitive_uint64(object swig.Schema_Object, field uint, index uint) uint64 {
	return swig.Schema_IndexUint64(object, field, index)
}

func WritePrimitive_uint64(object swig.Schema_Object, field uint, value uint64) {
	swig.Schema_AddUint64(object, field, value)
}

func ReadPrimitive_int64(object swig.Schema_Object, field uint, index uint) int64 {
	return swig.Schema_IndexInt64(object, field, index)
}

func WritePrimitive_int64(object swig.Schema_Object, field uint, value int64) {
	swig.Schema_AddInt64(object, field, value)
}

func ReadPrimitive_double(object swig.Schema_Object, field uint, index uint) float64 {
	return swig.Schema_IndexDouble(object, field, index)
}

func WritePrimitive_double(object swig.Schema_Object, field uint, value float64) {
	swig.Schema_AddDouble(object, field, value)
}

func ReadPrimitive_sfixed64(object swig.Schema_Object, field uint, index uint) int64 {
	return swig.Schema_IndexSfixed64(object, field, index)
}

func WritePrimitive_sfixed64(object swig.Schema_Object, field uint, value int64) {
	swig.Schema_AddSfixed64(object, field, value)
}

func ReadPrimitive_bool(object swig.Schema_Object, field uint, index uint) bool {
	return swig.Schema_IndexBool(object, field, index) > 0
}

func WritePrimitive_bool(object swig.Schema_Object, field uint, value bool) {
	byteValue := byte(0)
	if value {
		byteValue = byte(1)
	}
	swig.Schema_AddBool(object, field, byteValue)
}

func ReadPrimitive_uint32(object swig.Schema_Object, field uint, index uint) uint {
	return swig.Schema_IndexUint32(object, field, index)
}

func WritePrimitive_uint32(object swig.Schema_Object, field uint, value uint) {
	swig.Schema_AddUint32(object, field, value)
}

func ReadPrimitive_sfixed32(object swig.Schema_Object, field uint, index uint) int {
	return swig.Schema_IndexSfixed32(object, field, index)
}

func WritePrimitive_sfixed32(object swig.Schema_Object, field uint, value int) {
	swig.Schema_AddSfixed32(object, field, value)
}

func ReadPrimitive_fixed64(object swig.Schema_Object, field uint, index uint) uint64 {
	return swig.Schema_IndexFixed64(object, field, index)
}

func WritePrimitive_fixed64(object swig.Schema_Object, field uint, value uint64) {
	swig.Schema_AddFixed64(object, field, value)
}

func ReadPrimitive_float(object swig.Schema_Object, field uint, index uint) float32 {
	return swig.Schema_IndexFloat(object, field, index)
}

func WritePrimitive_float(object swig.Schema_Object, field uint, value float32) {
	swig.Schema_AddFloat(object, field, value)
}

func ReadPrimitive_EntityId(object swig.Schema_Object, field uint, index uint) int64 {
	return swig.Schema_IndexEntityId(object, field, index)
}

func WritePrimitive_EntityId(object swig.Schema_Object, field uint, value int64) {
	swig.Schema_AddEntityId(object, field, value)
}

func ReadPrimitive_bytes(object swig.Schema_Object, field uint, index uint) []byte {
	length := swig.Schema_IndexBytesLength(object, field, index)
	unsafePtr := unsafe.Pointer(swig.Schema_IndexBytes(object, field, index))
	return (*(*[4096]byte)(unsafePtr))[0:length]
}

func WritePrimitive_bytes(object swig.Schema_Object, field uint, value[]byte) {
	swig.Schema_AddBytes(object, field, &(value[0]), uint(len(value)))
}

func ReadPrimitive_sint64(object swig.Schema_Object, field uint, index uint) int64 {
	return swig.Schema_IndexSint64(object, field, index)
}

func WritePrimitive_sint64(object swig.Schema_Object, field uint, value int64) {
	swig.Schema_AddSint64(object, field, value)
}

func ReadPrimitive_fixed32(object swig.Schema_Object, field uint, index uint) uint {
	return swig.Schema_IndexFixed32(object, field, index)
}

func WritePrimitive_fixed32(object swig.Schema_Object, field uint, value uint) {
	swig.Schema_AddFixed32(object, field, value)
}