package main



import (
     "fmt"
     "github.com/rjfwhite/go-worker-api/example"
)

func main() {
//     params := Worker_DefaultConnectionParameters()
//     connection := Worker_ConnectAsync("hostname", 1337, "myworker", params)
     component := example.Schema_CreateComponentUpdate(32)
     fields := example.Schema_GetComponentUpdateFields(component)

     example.Schema_AddInt32(fields, 1, 32)
     val := example.Schema_GetInt32(fields, 1)

     fmt.Println(val)


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