package main

import (
	"github.com/rjfwhite/go-worker-api/example"
	"math/rand"
	"time"
	"fmt"
	"math"
)

type EntityComponent struct {
	entity_id    int64
	component_id uint
}

func main() {
	authoritativeComponents := make(map[EntityComponent]bool)

	params := example.Worker_DefaultConnectionParameters()
	params.SetWorker_type("Managed")

	vtable := example.NewWorker_ComponentVtable()
	params.SetDefault_component_vtable(vtable)

	rand.Seed(int64(time.Now().UnixNano()))

	workerId := fmt.Sprintf("Managed%d", rand.Int())
	fmt.Println("WorkerId " + workerId)

	future := example.Worker_ConnectAsync("localhost", 7777, workerId, params)

	timeout := uint(1000)
	connection := example.Worker_ConnectionFuture_Get(future, &timeout)
	example.Worker_ConnectionFuture_Destroy(future)

	fmt.Println("Checking if connected")
	if example.Worker_Connection_IsConnected(connection) > 0 {
		fmt.Println("Connected!")

		fmt.Println("Sending Welcome Log")
		logMessage := example.NewWorker_LogMessage()
		logMessage.SetLevel(4)
		logMessage.SetEntity_id(nil)
		logMessage.SetLogger_name("mylogger")
		logMessage.SetMessage("Hello, World!")
		example.Worker_Connection_SendLogMessage(connection, logMessage)
		example.DeleteWorker_LogMessage(logMessage)
		fmt.Println("Sent Welcome Log")

		dispatcher := Dispatcher{}

		dispatcher.Init()

		dispatcher.OnEntityAdded(func(entity_id int64) {
			fmt.Printf("ENTITY ADDED %d", entity_id)
		})

		dispatcher.OnPositionAdded(func(entity_id int64, data Position) {
			fmt.Printf("GOT POS %d\n", entity_id, data.Coords.X, data.Coords.Y, data.Coords.Z)
		})

		dispatcher.OnEntityAclAdded(func(entity_id int64, data EntityAcl) {
			fmt.Printf("GOT ACL %d\n", entity_id, data.Read, data.Write)
		})

		dispatcher.OnPositionAuthority(func(entity_id int64, is_authoritative bool) {
			authoritativeComponents[EntityComponent{entity_id: entity_id, component_id: 54}] = is_authoritative
		})

		dispatcher.OnPositionUpdated(func(entity_id int64, update PositionUpdate) {
			fmt.Printf("GOT POSUP %d\n", entity_id, update.Coords.X, update.Coords.Y, update.Coords.Z)
		})

		for example.Worker_Connection_IsConnected(connection) > 0 {
			ops := example.Worker_Connection_GetOpList(connection, uint(100))
			dispatcher.dispatchOps(ops)
			example.Worker_OpList_Destroy(ops)
			for ec, value := range (authoritativeComponents) {
				if value {
					if ec.component_id == 54 {
						sendPositionUpdate(connection, ec.entity_id, math.Sin(float64(time.Now().UnixNano())/1000000000.0)*10, 2, 3)
					}
				}
			}
		}

	} else {
		fmt.Println("Did not Connect!")
	}
}

func sendPositionUpdate(connection example.Worker_Connection, entity_id int64, x float64, y float64, z float64) {

	componentUpdate := example.Schema_CreateComponentUpdate(54)
	componentUpdateFields := example.Schema_GetComponentUpdateFields(componentUpdate)

	newCoordinates := Coordinates{x, y, z}

	WriteComponentUpdate_Position(componentUpdateFields, PositionUpdate{&newCoordinates})

	workerComponentUpdate := example.NewWorker_ComponentUpdate()
	workerComponentUpdate.SetComponent_id(54)
	workerComponentUpdate.SetSchema_type(componentUpdate)
	example.Worker_Connection_SendComponentUpdate(connection, entity_id, workerComponentUpdate)
}
