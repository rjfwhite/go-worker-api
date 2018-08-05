package main

import (
	"github.com/rjfwhite/go-worker-api/example"
	"math/rand"
	"time"
	"fmt"
)

type Connection struct {
	inner_connection example.Worker_Connection
}

func (connection *Connection) Connect() bool {
	params := example.Worker_DefaultConnectionParameters()
	params.SetWorker_type("Managed")

	vtable := example.NewWorker_ComponentVtable()
	params.SetDefault_component_vtable(vtable)

	rand.Seed(int64(time.Now().UnixNano()))

	workerId := fmt.Sprintf("Managed%d", rand.Int())
	fmt.Println("WorkerId " + workerId)

	future := example.Worker_ConnectAsync("localhost", 7777, workerId, params)

	timeout := uint(1000)
	connection.inner_connection = example.Worker_ConnectionFuture_Get(future, &timeout)
	example.Worker_ConnectionFuture_Destroy(future)
	return connection.IsConnected()
}

func (connection Connection) IsConnected() bool {
	return example.Worker_Connection_IsConnected(connection.inner_connection) > 0
}

func (connection Connection) ReadOps() []example.Worker_Op {
	timeout := uint(0)
	ops := example.Worker_Connection_GetOpList(connection.inner_connection, timeout)
	count := ops.GetOp_count()
	result := []example.Worker_Op{}
	for i := uint(0); i < count; i++ {
		result = append(result, example.Worker_OpList_GetSpecificOp(ops, i))
	}
	example.Worker_OpList_Destroy(ops)
	return result
}

func (connection Connection) SendLog(logger string, message string) {
	logMessage := example.NewWorker_LogMessage()
	logMessage.SetLevel(4)
	logMessage.SetEntity_id(nil)
	logMessage.SetLogger_name(logger)
	logMessage.SetMessage(message)
	example.Worker_Connection_SendLogMessage(connection.inner_connection, logMessage)
	example.DeleteWorker_LogMessage(logMessage)
}

func (connection Connection) SendComponentUpdate(entity_id int64, component_id uint, component_update example.Schema_ComponentUpdate) {
	//componentUpdate := example.Schema_CreateComponentUpdate(54)
	//componentUpdateFields := example.Schema_GetComponentUpdateFields(componentUpdate)
	//
	//newCoordinates := Coordinates{x, y, z}
	//
	//WriteComponentUpdate_Position(componentUpdateFields, PositionUpdate{&newCoordinates})

	worker_componentUpdate := example.NewWorker_ComponentUpdate()
	worker_componentUpdate.SetComponent_id(component_id)
	worker_componentUpdate.SetSchema_type(component_update)
	example.Worker_Connection_SendComponentUpdate(connection.inner_connection, entity_id, worker_componentUpdate)
}

func MakeConnection() *Connection {
	connection := Connection{}
	return &connection
}
