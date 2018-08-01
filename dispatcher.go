package main

import (
	"github.com/rjfwhite/go-worker-api/example"
)

type EntityAddedCallback func(entity_id int64)
type EntityRemovedCallback func(entity_id int64)

type ComponentAddedCallback func(entity_id int64, component_data example.Worker_ComponentData)
type ComponentUpdatedCallback func(entity_id int64, component_update example.Worker_ComponentUpdate)
type ComponentRemovedCallback func(entity_id int64)
type ComponentAuthorityCallback func(entity_id int64, is_authoritative bool)

type WORKER_OP_TYPE int

const (
	WORKER_OP_TYPE_DISCONNECT                  WORKER_OP_TYPE = 1
	WORKER_OP_TYPE_FLAG_UPDATE                 WORKER_OP_TYPE = 2
	WORKER_OP_TYPE_LOG_MESSAGE                 WORKER_OP_TYPE = 3
	WORKER_OP_TYPE_METRICS                     WORKER_OP_TYPE = 4
	WORKER_OP_TYPE_CRITICAL_SECTION            WORKER_OP_TYPE = 5
	WORKER_OP_TYPE_ADD_ENTITY                  WORKER_OP_TYPE = 6
	WORKER_OP_TYPE_REMOVE_ENTITY               WORKER_OP_TYPE = 7
	WORKER_OP_TYPE_RESERVE_ENTITY_ID_RESPONSE  WORKER_OP_TYPE = 8
	WORKER_OP_TYPE_RESERVE_ENTITY_IDS_RESPONSE WORKER_OP_TYPE = 9
	WORKER_OP_TYPE_CREATE_ENTITY_RESPONSE      WORKER_OP_TYPE = 10
	WORKER_OP_TYPE_DELETE_ENTITY_RESPONSE      WORKER_OP_TYPE = 11
	WORKER_OP_TYPE_ENTITY_QUERY_RESPONSE       WORKER_OP_TYPE = 12
	WORKER_OP_TYPE_ADD_COMPONENT               WORKER_OP_TYPE = 13
	WORKER_OP_TYPE_REMOVE_COMPONENT            WORKER_OP_TYPE = 14
	WORKER_OP_TYPE_AUTHORITY_CHANGE            WORKER_OP_TYPE = 15
	WORKER_OP_TYPE_COMPONENT_UPDATE            WORKER_OP_TYPE = 16
	WORKER_OP_TYPE_COMMAND_REQUEST             WORKER_OP_TYPE = 17
	WORKER_OP_TYPE_COMMAND_RESPONSE            WORKER_OP_TYPE = 18
)

type Dispatcher struct {
	EntityAddedCallbacks   []EntityAddedCallback
	EntityRemovedCallbacks []EntityRemovedCallback

	ComponentAddedCallbacks     map[int][]ComponentAddedCallback
	ComponentUpdatedCallbacks   map[int][]ComponentUpdatedCallback
	ComponentRemovedCallbacks   map[int][]ComponentRemovedCallback
	ComponentAuthorityCallbacks map[int][]ComponentAuthorityCallback
}

func (dispatcher Dispatcher) dispatch(op example.Worker_Op) {
	opType := WORKER_OP_TYPE(op.GetOp_type())
	switch opType {
	case WORKER_OP_TYPE_ADD_ENTITY:
		specificOp := op.GetAdd_entity()
		for _, callback := range dispatcher.EntityAddedCallbacks {
			callback(specificOp.GetEntity_id())
		}

	case WORKER_OP_TYPE_REMOVE_ENTITY:
		specificOp := op.GetRemove_entity()
		for _, callback := range dispatcher.EntityRemovedCallbacks {
			callback(specificOp.GetEntity_id())
		}

	case WORKER_OP_TYPE_ADD_COMPONENT:

	case WORKER_OP_TYPE_AUTHORITY_CHANGE:

	case WORKER_OP_TYPE_METRICS:
	case WORKER_OP_TYPE_LOG_MESSAGE:
	default:
	}
}
