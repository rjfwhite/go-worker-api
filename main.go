package main

import (
	"github.com/rjfwhite/go-worker-api/example"
	"math/rand"
	"time"
	"fmt"
	"math"
)


type EntityComponent struct {
	entity_id int64
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

		for example.Worker_Connection_IsConnected(connection) > 0 {


			//fmt.Println("Reading Ops")
			ops := example.Worker_Connection_GetOpList(connection, uint(100))
			count := ops.GetOp_count()

			for i := 0; i < int(count); i++ {
				op := example.Worker_OpList_GetSpecificOp(ops, uint(i))
				opType := WORKER_OP_TYPE(op.GetOp_type())

				switch opType {
				case WORKER_OP_TYPE_METRICS:
					//fmt.Println("GOT METRICS")

				case WORKER_OP_TYPE_ADD_ENTITY:
					addEntity := op.GetAdd_entity()
					fmt.Printf("GOT ADD ENTITY: %d \n", addEntity.GetEntity_id())

				case WORKER_OP_TYPE_ADD_COMPONENT:
					addComponent := op.GetAdd_component()
					component_id := addComponent.GetData().GetComponent_id()
					entity_id := addComponent.GetEntity_id()
					fmt.Printf("GOT ADD COMPONENT %d for ENTITY %d\n", component_id, entity_id)
					if component_id == 54 {
						fields := example.Schema_GetComponentDataFields(addComponent.GetData().GetSchema_type())
						position := ReadComponent_Position(fields)
						fmt.Println("GOT POSITION ", position.Coords.X, position.Coords.Y, position.Coords.Z)
					} else if component_id == 50 {
						fmt.Println("GOT ACL")
						fields := example.Schema_GetComponentDataFields(addComponent.GetData().GetSchema_type())
						read := ReadList_Object_WorkerAttributeSet(fields, 1, 0)
						write := ReadMap_Primitive_uint32_to_List_Object_WorkerAttributeSet(fields, 2, 0)


						fmt.Printf("GOT ACL READ :%s - WRITE:%s", read, write)
					}

				case WORKER_OP_TYPE_LOG_MESSAGE:
					fmt.Println("GOT LOG")

				case WORKER_OP_TYPE_AUTHORITY_CHANGE:
					//fmt.Println("Authority Change")
					authorityChange := op.GetAuthority_change()
					entity_id := authorityChange.GetEntity_id()
					component_id := authorityChange.GetComponent_id()
					authoritative := (authorityChange.GetAuthority() > 0)
					authoritativeComponents[EntityComponent{entity_id:entity_id, component_id:component_id}] = authoritative

				default:

				}


				//fmt.Printf("%d:", op.GetOp_type())
			}

			for ec, value := range(authoritativeComponents) {
				if value {
					if ec.component_id == 54 {
						sendPositionUpdate(connection, ec.entity_id, math.Sin(float64(time.Now().UnixNano()) / 1000000000.0) * 10, 2, 3)
					}
				}
			}



			example.Worker_OpList_Destroy(ops)
			//fmt.Printf("Finished Reading %d Ops\n", count)
		}



		//fmt.Println("CONNECTED!")

	} else {
		fmt.Println("Did not Connect!")
		component := example.Schema_CreateComponentUpdate(32)
		fields := example.Schema_GetComponentUpdateFields(component)

		example.Schema_AddInt32(fields, 1, 32)
		val := example.Schema_GetInt32(fields, 1)

		fmt.Println(val)
	}

	//  Schema_AddFloat(component, 2, 5.0)
	// object := Schema_AddObject(component, 3)
	// Schema_AddInt32(object, 1, 1337)
	//data := []byte{1,2,3}
	//Schema_AddBytes(object, 2, &data[0], 3)

	//     op_list := Worker_Connection_GetOpList(connection, 0)

	//Worker_OpList_GetOp(op_list, 1)

	//fields := Schema_GetComponentUpdateFields(component)
	//
	//if Schema_GetInt32Count(fields, 1) > 0 {
	//     data := Schema_GetInt32(fields, 1)
	//     fmt.Print(data)
	//}

}



func sendPositionUpdate(connection example.Worker_Connection, entity_id int64, x float64, y float64, z float64) {
	//fmt.Printf("Sending update %f %f %f\n", x, y, z)

	componentUpdate := example.Schema_CreateComponentUpdate(54)
	componentUpdateFields := example.Schema_GetComponentUpdateFields(componentUpdate)

	newCoordinates := Coordinates{x, y, z}

	WriteComponentUpdate_Position(componentUpdateFields, PositionUpdate{&newCoordinates})

	workerComponentUpdate := example.NewWorker_ComponentUpdate()
	workerComponentUpdate.SetComponent_id(54)
	workerComponentUpdate.SetSchema_type(componentUpdate)
	example.Worker_Connection_SendComponentUpdate(connection, entity_id, workerComponentUpdate)
}
