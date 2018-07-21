package main

import (
     "fmt"
     "github.com/rjfwhite/go-worker-api/example"
     "math/rand"
     "time"
)

func main() {
     params := example.Worker_DefaultConnectionParameters()
     params.SetWorker_type("Managed")

     vtable := example.NewWorker_ComponentVtable()
     params.SetDefault_component_vtable(vtable)

     rand.Seed(int64(time.Now().UnixNano()))

     workerId := fmt.Sprintf("External%d", rand.Int())
     fmt.Println("WorkerId " + workerId)

     future := example.Worker_ConnectAsync("localhost", 7777, workerId, params)

     timeout := uint(5000)
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
               fmt.Println("Reading Ops")
               ops := example.Worker_Connection_GetOpList(connection, timeout)
               count := ops.GetOp_count()
               op := ops.GetOps()
               fmt.Println(op.GetOp_type())
               example.Worker_OpList_Destroy(ops)
               fmt.Printf("Finished Reading %d Ops\n", count)
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
