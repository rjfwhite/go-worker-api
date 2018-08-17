package main

import (
	"github.com/rjfwhite/go-worker-api/example"
	"math/rand"
	"time"
	"fmt"
)

type Connection struct {
	workerType      string
	innerConnection example.Worker_Connection
}

func (connection *Connection) Connect(host string, port uint16) bool {
	params := example.Worker_DefaultConnectionParameters()
	params.SetWorker_type(connection.workerType)

	params.SetDefault_component_vtable(example.NewWorker_ComponentVtable())

	rand.Seed(int64(time.Now().UnixNano()))
	workerId := fmt.Sprintf("%s-%d", connection.workerType, rand.Int())

	future := example.Worker_ConnectAsync(host, port, workerId, params)

	timeout := uint(1000)
	connection.innerConnection = example.Worker_ConnectionFuture_Get(future, &timeout)
	example.Worker_ConnectionFuture_Destroy(future)
	return connection.IsConnected()
}

func (connection Connection) IsConnected() bool {
	return example.Worker_Connection_IsConnected(connection.innerConnection) > 0
}

func (connection Connection) ReadOps(timeout uint) []example.Worker_Op {
	ops := example.Worker_Connection_GetOpList(connection.innerConnection, timeout)
	count := ops.GetOp_count()
	result := []example.Worker_Op{}
	for i := uint(0); i < count; i++ {
		result = append(result, example.Worker_OpList_GetSpecificOp(ops, i))
	}
	//example.Worker_OpList_Destroy(ops)
	return result
}

func (connection Connection) SendLog(logger string, message string) {
	logMessage := example.NewWorker_LogMessage()
	logMessage.SetLevel(4)
	logMessage.SetEntity_id(nil)
	logMessage.SetLogger_name(logger)
	logMessage.SetMessage(message)
	example.Worker_Connection_SendLogMessage(connection.innerConnection, logMessage)
	example.DeleteWorker_LogMessage(logMessage)
}

func (connection Connection) SendComponentUpdate(entity_id int64, component_id uint, component_update example.Schema_ComponentUpdate) {
	workerComponentupdate := example.NewWorker_ComponentUpdate()
	workerComponentupdate.SetComponent_id(component_id)
	workerComponentupdate.SetSchema_type(component_update)
	example.Worker_Connection_SendComponentUpdate(connection.innerConnection, entity_id, workerComponentupdate)
}

func MakeConnection(workerType string) *Connection {
	connection := Connection{workerType:workerType}
	return &connection
}
