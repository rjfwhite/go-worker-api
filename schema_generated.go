package main

import "github.com/rjfwhite/go-worker-api/swig"

func ReadOption_Object_Coordinates(object swig.Schema_Object, field uint, index uint) *Coordinates {
	if swig.Schema_GetObjectCount(object, field) > 0 {
		result := ReadObject_Coordinates(object, field, index)
		return &result
	}
	return nil
}

func WriteOption_Object_Coordinates(object swig.Schema_Object, field uint, value *Coordinates) {
	if value != nil {
		WriteObject_Coordinates(object, field, *value)
	}
}

func ReadMap_Primitive_uint32_to_List_Object_WorkerRequirementSet (object swig.Schema_Object, field uint, index uint) map[uint][]WorkerRequirementSet  {
	count := swig.Schema_GetObjectCount(object, field)
	result := map[uint][]WorkerRequirementSet {}
	for i := uint(0); i < count; i++ {
		innerObject := swig.Schema_IndexObject(object, field, i)
		key := ReadPrimitive_uint32(innerObject, 1, 0)
		value := ReadList_Object_WorkerRequirementSet (innerObject, 2, 0)
		result[key] = value
	}
	return result
}

func WriteMap_Primitive_uint32_to_List_Object_WorkerRequirementSet (object swig.Schema_Object, field uint, value map[uint][]WorkerRequirementSet ) {
	for k, v := range(value) {
		innerObject := swig.Schema_AddObject(object, field)
		WritePrimitive_uint32(innerObject, 1, k)
		WriteList_Object_WorkerRequirementSet (innerObject, 2, v)
	}
}

type WorkerRequirementSet  struct {
	attribute_set []WorkerAttributeSet
}

func ReadObject_WorkerRequirementSet (object swig.Schema_Object, field uint, index uint) WorkerRequirementSet  {
	innerObject := swig.Schema_IndexObject(object, field, index)
	return WorkerRequirementSet  {
		attribute_set : ReadList_Object_WorkerAttributeSet (innerObject, 1, 0),
	}
}

func WriteObject_WorkerRequirementSet (object swig.Schema_Object, field uint, value WorkerRequirementSet ) {
	innerObject := swig.Schema_AddObject(object, field)
	WriteList_Object_WorkerAttributeSet (innerObject, 1, value.attribute_set)
}

type MetaData struct {
	EntityType string
}

type MetaDataUpdate struct {
	EntityType *string
}

func ReadComponent_MetaData(object swig.Schema_Object) MetaData {
	return MetaData {
		EntityType : ReadPrimitive_string(object, 1, 0),
	}
}

func WriteComponent_MetaData(object swig.Schema_Object, value MetaData) {
	WritePrimitive_string(object, 1, value.EntityType)
}

func ReadComponentUpdate_MetaData(object swig.Schema_Object) MetaDataUpdate {
	return MetaDataUpdate {
		EntityType : ReadOption_Primitive_string(object, 1, 0),
	}
}

func WriteComponentUpdate_MetaData(object swig.Schema_Object, value MetaDataUpdate) {
	WriteOption_Primitive_string(object, 1, value.EntityType)
}

type MetaDataAddedCallback func(entityId int64, data MetaData)
type MetaDataUpdatedCallback func(entityId int64, update MetaDataUpdate)

func (dispatcher *Dispatcher) OnMetaDataAdded(callback MetaDataAddedCallback) {
	component_id := uint(53)
	inner_callback := func(entity_id int64, component_data swig.Worker_ComponentData) {
		data_fields := swig.Schema_GetComponentDataFields(component_data.GetSchema_type())
		component := ReadComponent_MetaData(data_fields)
		callback(entity_id, component)
	}
	dispatcher.OnComponentAdded(component_id, inner_callback)
}

func (dispatcher *Dispatcher) OnMetaDataUpdated(callback MetaDataUpdatedCallback) {
	component_id := uint(53)
	inner_callback := func(entity_id int64, component_update swig.Worker_ComponentUpdate) {
		data_fields := swig.Schema_GetComponentUpdateFields(component_update.GetSchema_type())
		component := ReadComponentUpdate_MetaData(data_fields)
		callback(entity_id, component)
	}
	dispatcher.OnComponentUpdated(component_id, inner_callback)
}

func (dispatcher *Dispatcher) OnMetaDataAuthority(callback ComponentAuthorityCallback) {
	component_id := uint(53)
	dispatcher.OnComponentAuthority(component_id, callback)
}

func (dispatcher *Dispatcher) OnMetaDataRemoved(callback ComponentRemovedCallback) {
	component_id := uint(53)
	dispatcher.OnComponentRemoved(component_id, callback)
}

func (connection Connection) SendMetaDataUpdate(entity_id int64, value MetaDataUpdate) {
	component_id := uint(53)
	component_update := swig.Schema_CreateComponentUpdate(component_id)
	component_update_fields := swig.Schema_GetComponentUpdateFields(component_update)
	WriteComponentUpdate_MetaData(component_update_fields, value)
	connection.SendComponentUpdate(entity_id, component_id, component_update)
}

type Coordinates struct {
	X float64
	Y float64
	Z float64
}

func ReadObject_Coordinates(object swig.Schema_Object, field uint, index uint) Coordinates {
	innerObject := swig.Schema_IndexObject(object, field, index)
	return Coordinates {
		X : ReadPrimitive_double(innerObject, 1, 0),
		Y : ReadPrimitive_double(innerObject, 2, 0),
		Z : ReadPrimitive_double(innerObject, 3, 0),
	}
}

func WriteObject_Coordinates(object swig.Schema_Object, field uint, value Coordinates) {
	innerObject := swig.Schema_AddObject(object, field)
	WritePrimitive_double(innerObject, 1, value.X)
	WritePrimitive_double(innerObject, 2, value.Y)
	WritePrimitive_double(innerObject, 3, value.Z)
}

type Color uint

const (
	Blue Color = 1
	Red Color = 2
)

func ReadEnum_Color(object swig.Schema_Object, field uint, index uint) Color {
	return Color(swig.Schema_IndexEnum(object, field, index))
}

func WriteEnum_Color(object swig.Schema_Object, field uint, value Color) {
	swig.Schema_AddEnum(object, field, uint(value))
}

func ReadList_Object_WorkerRequirementSet (object swig.Schema_Object, field uint, index uint) []WorkerRequirementSet  {
	count := swig.Schema_GetObjectCount(object, field)
	result := []WorkerRequirementSet {}
	for i := uint(0); i < count; i++ {
		result = append(result, ReadObject_WorkerRequirementSet (object, field, i))
	}
	return result
}

func WriteList_Object_WorkerRequirementSet (object swig.Schema_Object, field uint, value []WorkerRequirementSet ) {
	for _, i := range(value) {
		WriteObject_WorkerRequirementSet (object, field, i)
	}
}

func ReadOption_List_Object_WorkerRequirementSet (object swig.Schema_Object, field uint, index uint) *[]WorkerRequirementSet  {
	if swig.Schema_GetObjectCount(object, field) > 0 {
		result := ReadList_Object_WorkerRequirementSet (object, field, index)
		return &result
	}
	return nil
}

func WriteOption_List_Object_WorkerRequirementSet (object swig.Schema_Object, field uint, value *[]WorkerRequirementSet ) {
	if value != nil {
		WriteList_Object_WorkerRequirementSet (object, field, *value)
	}
}

type WorkerAttributeSet  struct {
	attribute []string
}

func ReadObject_WorkerAttributeSet (object swig.Schema_Object, field uint, index uint) WorkerAttributeSet  {
	innerObject := swig.Schema_IndexObject(object, field, index)
	return WorkerAttributeSet  {
		attribute : ReadList_Primitive_string(innerObject, 1, 0),
	}
}

func WriteObject_WorkerAttributeSet (object swig.Schema_Object, field uint, value WorkerAttributeSet ) {
	innerObject := swig.Schema_AddObject(object, field)
	WriteList_Primitive_string(innerObject, 1, value.attribute)
}

type Position struct {
	Coords Coordinates
}

type PositionUpdate struct {
	Coords *Coordinates
}

func ReadComponent_Position(object swig.Schema_Object) Position {
	return Position {
		Coords : ReadObject_Coordinates(object, 1, 0),
	}
}

func WriteComponent_Position(object swig.Schema_Object, value Position) {
	WriteObject_Coordinates(object, 1, value.Coords)
}

func ReadComponentUpdate_Position(object swig.Schema_Object) PositionUpdate {
	return PositionUpdate {
		Coords : ReadOption_Object_Coordinates(object, 1, 0),
	}
}

func WriteComponentUpdate_Position(object swig.Schema_Object, value PositionUpdate) {
	WriteOption_Object_Coordinates(object, 1, value.Coords)
}

type PositionAddedCallback func(entityId int64, data Position)
type PositionUpdatedCallback func(entityId int64, update PositionUpdate)

func (dispatcher *Dispatcher) OnPositionAdded(callback PositionAddedCallback) {
	component_id := uint(54)
	inner_callback := func(entity_id int64, component_data swig.Worker_ComponentData) {
		data_fields := swig.Schema_GetComponentDataFields(component_data.GetSchema_type())
		component := ReadComponent_Position(data_fields)
		callback(entity_id, component)
	}
	dispatcher.OnComponentAdded(component_id, inner_callback)
}

func (dispatcher *Dispatcher) OnPositionUpdated(callback PositionUpdatedCallback) {
	component_id := uint(54)
	inner_callback := func(entity_id int64, component_update swig.Worker_ComponentUpdate) {
		data_fields := swig.Schema_GetComponentUpdateFields(component_update.GetSchema_type())
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
	component_update := swig.Schema_CreateComponentUpdate(component_id)
	component_update_fields := swig.Schema_GetComponentUpdateFields(component_update)
	WriteComponentUpdate_Position(component_update_fields, value)
	connection.SendComponentUpdate(entity_id, component_id, component_update)
}

type EntityAcl struct {
	Read []WorkerRequirementSet
	Write map[uint][]WorkerRequirementSet
}

type EntityAclUpdate struct {
	Read *[]WorkerRequirementSet
	Write *map[uint][]WorkerRequirementSet
}

func ReadComponent_EntityAcl(object swig.Schema_Object) EntityAcl {
	return EntityAcl {
		Read : ReadList_Object_WorkerRequirementSet (object, 1, 0),
		Write : ReadMap_Primitive_uint32_to_List_Object_WorkerRequirementSet (object, 2, 0),
	}
}

func WriteComponent_EntityAcl(object swig.Schema_Object, value EntityAcl) {
	WriteList_Object_WorkerRequirementSet (object, 1, value.Read)
	WriteMap_Primitive_uint32_to_List_Object_WorkerRequirementSet (object, 2, value.Write)
}

func ReadComponentUpdate_EntityAcl(object swig.Schema_Object) EntityAclUpdate {
	return EntityAclUpdate {
		Read : ReadOption_List_Object_WorkerRequirementSet (object, 1, 0),
		Write : ReadOption_Map_Primitive_uint32_to_List_Object_WorkerRequirementSet (object, 2, 0),
	}
}

func WriteComponentUpdate_EntityAcl(object swig.Schema_Object, value EntityAclUpdate) {
	WriteOption_List_Object_WorkerRequirementSet (object, 1, value.Read)
	WriteOption_Map_Primitive_uint32_to_List_Object_WorkerRequirementSet (object, 2, value.Write)
}

type EntityAclAddedCallback func(entityId int64, data EntityAcl)
type EntityAclUpdatedCallback func(entityId int64, update EntityAclUpdate)

func (dispatcher *Dispatcher) OnEntityAclAdded(callback EntityAclAddedCallback) {
	component_id := uint(50)
	inner_callback := func(entity_id int64, component_data swig.Worker_ComponentData) {
		data_fields := swig.Schema_GetComponentDataFields(component_data.GetSchema_type())
		component := ReadComponent_EntityAcl(data_fields)
		callback(entity_id, component)
	}
	dispatcher.OnComponentAdded(component_id, inner_callback)
}

func (dispatcher *Dispatcher) OnEntityAclUpdated(callback EntityAclUpdatedCallback) {
	component_id := uint(50)
	inner_callback := func(entity_id int64, component_update swig.Worker_ComponentUpdate) {
		data_fields := swig.Schema_GetComponentUpdateFields(component_update.GetSchema_type())
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
	component_update := swig.Schema_CreateComponentUpdate(component_id)
	component_update_fields := swig.Schema_GetComponentUpdateFields(component_update)
	WriteComponentUpdate_EntityAcl(component_update_fields, value)
	connection.SendComponentUpdate(entity_id, component_id, component_update)
}

func ReadList_Object_WorkerAttributeSet (object swig.Schema_Object, field uint, index uint) []WorkerAttributeSet  {
	count := swig.Schema_GetObjectCount(object, field)
	result := []WorkerAttributeSet {}
	for i := uint(0); i < count; i++ {
		result = append(result, ReadObject_WorkerAttributeSet (object, field, i))
	}
	return result
}

func WriteList_Object_WorkerAttributeSet (object swig.Schema_Object, field uint, value []WorkerAttributeSet ) {
	for _, i := range(value) {
		WriteObject_WorkerAttributeSet (object, field, i)
	}
}

func ReadList_Primitive_string(object swig.Schema_Object, field uint, index uint) []string {
	count := swig.Schema_GetBytesCount(object, field)
	result := []string{}
	for i := uint(0); i < count; i++ {
		result = append(result, ReadPrimitive_string(object, field, i))
	}
	return result
}

func WriteList_Primitive_string(object swig.Schema_Object, field uint, value []string) {
	for _, i := range(value) {
		WritePrimitive_string(object, field, i)
	}
}

func ReadOption_Map_Primitive_uint32_to_List_Object_WorkerRequirementSet (object swig.Schema_Object, field uint, index uint) *map[uint][]WorkerRequirementSet  {
	if swig.Schema_GetObjectCount(object, field) > 0 {
		result := ReadMap_Primitive_uint32_to_List_Object_WorkerRequirementSet (object, field, index)
		return &result
	}
	return nil
}

func WriteOption_Map_Primitive_uint32_to_List_Object_WorkerRequirementSet (object swig.Schema_Object, field uint, value *map[uint][]WorkerRequirementSet ) {
	if value != nil {
		WriteMap_Primitive_uint32_to_List_Object_WorkerRequirementSet (object, field, *value)
	}
}

func ReadOption_Primitive_string(object swig.Schema_Object, field uint, index uint) *string {
	if swig.Schema_GetBytesCount(object, field) > 0 {
		result := ReadPrimitive_string(object, field, index)
		return &result
	}
	return nil
}

func WriteOption_Primitive_string(object swig.Schema_Object, field uint, value *string) {
	if value != nil {
		WritePrimitive_string(object, field, *value)
	}
}

