package main

import (
	"github.com/rjfwhite/go-worker-api/swig"
	"math/rand"
	"time"
	"fmt"
)

type Connection struct {
	workerType      string
	innerConnection swig.Worker_Connection
}

func (connection *Connection) Connect(host string, port uint16) bool {
	params := swig.Worker_DefaultConnectionParameters()
	params.SetWorker_type(connection.workerType)

	params.SetDefault_component_vtable(swig.NewWorker_ComponentVtable())

	rand.Seed(int64(time.Now().UnixNano()))
	workerId := fmt.Sprintf("%s-%d", connection.workerType, rand.Int())

	future := swig.Worker_ConnectAsync(host, port, workerId, params)

	timeout := uint(1000)
	connection.innerConnection = swig.Worker_ConnectionFuture_Get(future, &timeout)
	swig.Worker_ConnectionFuture_Destroy(future)
	return connection.IsConnected()
}

func (connection Connection) IsConnected() bool {
	return swig.Worker_Connection_IsConnected(connection.innerConnection) > 0
}

func (connection Connection) ReadOps(timeout uint) []swig.Worker_Op {
	ops := swig.Worker_Connection_GetOpList(connection.innerConnection, timeout)
	count := ops.GetOp_count()
	result := []swig.Worker_Op{}
	for i := uint(0); i < count; i++ {
		result = append(result, swig.Worker_OpList_GetSpecificOp(ops, i))
	}
	//example.Worker_OpList_Destroy(ops)
	return result
}

func (connection Connection) SendLog(logger string, message string) {
	logMessage := swig.NewWorker_LogMessage()
	logMessage.SetLevel(4)
	logMessage.SetEntity_id(nil)
	logMessage.SetLogger_name(logger)
	logMessage.SetMessage(message)
	swig.Worker_Connection_SendLogMessage(connection.innerConnection, logMessage)
	swig.DeleteWorker_LogMessage(logMessage)
}

func (connection Connection) SendComponentUpdate(entity_id int64, component_id uint, component_update swig.Schema_ComponentUpdate) {
	workerComponentupdate := swig.NewWorker_ComponentUpdate()
	workerComponentupdate.SetComponent_id(component_id)
	workerComponentupdate.SetSchema_type(component_update)
	swig.Worker_Connection_SendComponentUpdate(connection.innerConnection, entity_id, workerComponentupdate)
}

func MakeConnection(workerType string) *Connection {
	connection := Connection{workerType:workerType}
	return &connection
}
