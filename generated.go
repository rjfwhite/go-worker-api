package main

import "github.com/rjfwhite/go-worker-api/example"

type Color uint

const (
	Blue Color = 1
	Red Color = 2
)
func ReadEnum_Color(object example.Schema_Object, field uint, index uint) Color {
	return Color(example.Schema_IndexEnum(object, field, index))
}
func WriteEnum_Color(object example.Schema_Object, field uint, value Color) {
	example.Schema_AddEnum(object, field, uint(value))
}
type Coordinates struct {
	X float64
	Y float64
	Z float64
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
func ReadMap_Primitive_uint32_to_List_Object_WorkerRequirementSet (object example.Schema_Object, field uint, index uint) map[uint][]WorkerRequirementSet  {
	count := example.Schema_GetObjectCount(object, field)
	result := map[uint][]WorkerRequirementSet {}
	for i := uint(0); i < count; i++ {
		innerObject := example.Schema_IndexObject(object, field, i)
		key := ReadPrimitive_uint32(innerObject, 1, 0)
		value := ReadList_Object_WorkerRequirementSet (innerObject, 2, 0)
		result[key] = value
	}
	return result
}
func WriteMap_Primitive_uint32_to_List_Object_WorkerRequirementSet (object example.Schema_Object, field uint, value map[uint][]WorkerRequirementSet ) {
	for k, v := range(value) {
		innerObject := example.Schema_AddObject(object, field)
		WritePrimitive_uint32(innerObject, 1, k)
		WriteList_Object_WorkerRequirementSet (innerObject, 2, v)
	}
}
func ReadList_Object_WorkerAttributeSet (object example.Schema_Object, field uint, index uint) []WorkerAttributeSet  {
	count := example.Schema_GetObjectCount(object, field)
	result := []WorkerAttributeSet {}
	for i := uint(0); i < count; i++ {
		result = append(result, ReadObject_WorkerAttributeSet (object, field, i))
	}
	return result
}
func WriteList_Object_WorkerAttributeSet (object example.Schema_Object, field uint, value []WorkerAttributeSet ) {
	for _, i := range(value) {
		WriteObject_WorkerAttributeSet (object, field, i)
	}
}
type WorkerAttributeSet  struct {
	attribute []string
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
type WorkerRequirementSet  struct {
	attribute_set []WorkerAttributeSet
}
func ReadObject_WorkerRequirementSet (object example.Schema_Object, field uint, index uint) WorkerRequirementSet  {
	innerObject := example.Schema_IndexObject(object, field, index)
	return WorkerRequirementSet  {
		attribute_set : ReadList_Object_WorkerAttributeSet (innerObject, 1, 0),
	}
}
func WriteObject_WorkerRequirementSet (object example.Schema_Object, field uint, value WorkerRequirementSet ) {
	innerObject := example.Schema_AddObject(object, field)
	WriteList_Object_WorkerAttributeSet (innerObject, 1, value.attribute_set)
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
func (dispatcher *Dispatcher) OnPositionAdded(callback PositionAddedCallback) {
	component_id := uint(54)
	inner_callback := func(entity_id int64, component_data example.Worker_ComponentData) {
		data_fields := example.Schema_GetComponentDataFields(component_data.GetSchema_type())
		component := ReadComponent_Position(data_fields)
		callback(entity_id, component)
	}
	dispatcher.OnComponentAdded(component_id, inner_callback)
}
func (dispatcher *Dispatcher) OnPositionUpdated(callback PositionUpdatedCallback) {
	component_id := uint(54)
	inner_callback := func(entity_id int64, component_update example.Worker_ComponentUpdate) {
		data_fields := example.Schema_GetComponentUpdateFields(component_update.GetSchema_type())
		component := ReadComponentUpdate_Position(data_fields)
		callback(entity_id, component)
	}
	dispatcher.OnComponentUpdated(component_id, inner_callback)
}
func (dispatcher *Dispatcher) OnPositionAuthority(callback ComponentAuthorityCallback) {
	component_id := uint(54)
	dispatcher.OnComponentAuthority(component_id, callback)
}
func (dispatcher *Dispatcher) OnPositionRemoved(callback ComponentRemovedCallback) {
	component_id := uint(54)
	dispatcher.OnComponentRemoved(component_id, callback)
}
func (connection Connection) SendPositionUpdate(entity_id int64, value PositionUpdate) {
	component_id := uint(54)
	component_update := example.Schema_CreateComponentUpdate(component_id)
	component_update_fields := example.Schema_GetComponentUpdateFields(component_update)
	WriteComponentUpdate_Position(component_update_fields, value)
	connection.SendComponentUpdate(entity_id, component_id, component_update)
}
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
func ReadList_Object_WorkerRequirementSet (object example.Schema_Object, field uint, index uint) []WorkerRequirementSet  {
	count := example.Schema_GetObjectCount(object, field)
	result := []WorkerRequirementSet {}
	for i := uint(0); i < count; i++ {
		result = append(result, ReadObject_WorkerRequirementSet (object, field, i))
	}
	return result
}
func WriteList_Object_WorkerRequirementSet (object example.Schema_Object, field uint, value []WorkerRequirementSet ) {
	for _, i := range(value) {
		WriteObject_WorkerRequirementSet (object, field, i)
	}
}
func ReadOption_List_Object_WorkerRequirementSet (object example.Schema_Object, field uint, index uint) *[]WorkerRequirementSet  {
	if example.Schema_GetObjectCount(object, field) > 0 {
		result := ReadList_Object_WorkerRequirementSet (object, field, index)
		return &result
	}
	return nil
}
func WriteOption_List_Object_WorkerRequirementSet (object example.Schema_Object, field uint, value *[]WorkerRequirementSet ) {
	if value != nil {
		WriteList_Object_WorkerRequirementSet (object, field, *value)
	}
}
type EntityAcl struct {
	Read []WorkerRequirementSet
	Write map[uint][]WorkerRequirementSet
}
type EntityAclUpdate struct {
	Read *[]WorkerRequirementSet
	Write *map[uint][]WorkerRequirementSet
}
func ReadComponent_EntityAcl(object example.Schema_Object) EntityAcl {
	return EntityAcl {
		Read : ReadList_Object_WorkerRequirementSet (object, 1, 0),
		Write : ReadMap_Primitive_uint32_to_List_Object_WorkerRequirementSet (object, 2, 0),
	}
}
func WriteComponent_EntityAcl(object example.Schema_Object, value EntityAcl) {
	WriteList_Object_WorkerRequirementSet (object, 1, value.Read)
	WriteMap_Primitive_uint32_to_List_Object_WorkerRequirementSet (object, 2, value.Write)
}
func ReadComponentUpdate_EntityAcl(object example.Schema_Object) EntityAclUpdate {
	return EntityAclUpdate {
		Read : ReadOption_List_Object_WorkerRequirementSet (object, 1, 0),
		Write : ReadOption_Map_Primitive_uint32_to_List_Object_WorkerRequirementSet (object, 2, 0),
	}
}
func WriteComponentUpdate_EntityAcl(object example.Schema_Object, value EntityAclUpdate) {
	WriteOption_List_Object_WorkerRequirementSet (object, 1, value.Read)
	WriteOption_Map_Primitive_uint32_to_List_Object_WorkerRequirementSet (object, 2, value.Write)
}
type EntityAclAddedCallback func(entity_id int64, data EntityAcl)
type EntityAclUpdatedCallback func(entity_id int64, update EntityAclUpdate)
func (dispatcher *Dispatcher) OnEntityAclAdded(callback EntityAclAddedCallback) {
	component_id := uint(50)
	inner_callback := func(entity_id int64, component_data example.Worker_ComponentData) {
		data_fields := example.Schema_GetComponentDataFields(component_data.GetSchema_type())
		component := ReadComponent_EntityAcl(data_fields)
		callback(entity_id, component)
	}
	dispatcher.OnComponentAdded(component_id, inner_callback)
}
func (dispatcher *Dispatcher) OnEntityAclUpdated(callback EntityAclUpdatedCallback) {
	component_id := uint(50)
	inner_callback := func(entity_id int64, component_update example.Worker_ComponentUpdate) {
		data_fields := example.Schema_GetComponentUpdateFields(component_update.GetSchema_type())
		component := ReadComponentUpdate_EntityAcl(data_fields)
		callback(entity_id, component)
	}
	dispatcher.OnComponentUpdated(component_id, inner_callback)
}
func (dispatcher *Dispatcher) OnEntityAclAuthority(callback ComponentAuthorityCallback) {
	component_id := uint(50)
	dispatcher.OnComponentAuthority(component_id, callback)
}
func (dispatcher *Dispatcher) OnEntityAclRemoved(callback ComponentRemovedCallback) {
	component_id := uint(50)
	dispatcher.OnComponentRemoved(component_id, callback)
}
func (connection Connection) SendEntityAclUpdate(entity_id int64, value EntityAclUpdate) {
	component_id := uint(50)
	component_update := example.Schema_CreateComponentUpdate(component_id)
	component_update_fields := example.Schema_GetComponentUpdateFields(component_update)
	WriteComponentUpdate_EntityAcl(component_update_fields, value)
	connection.SendComponentUpdate(entity_id, component_id, component_update)
}
func ReadOption_Map_Primitive_uint32_to_List_Object_WorkerRequirementSet (object example.Schema_Object, field uint, index uint) *map[uint][]WorkerRequirementSet  {
	if example.Schema_GetObjectCount(object, field) > 0 {
		result := ReadMap_Primitive_uint32_to_List_Object_WorkerRequirementSet (object, field, index)
		return &result
	}
	return nil
}
func WriteOption_Map_Primitive_uint32_to_List_Object_WorkerRequirementSet (object example.Schema_Object, field uint, value *map[uint][]WorkerRequirementSet ) {
	if value != nil {
		WriteMap_Primitive_uint32_to_List_Object_WorkerRequirementSet (object, field, *value)
	}
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
