package main

import (
	"time"
	"fmt"
	"math"
)

type EntityComponent struct {
	entity_id    int64
	component_id uint
}

func main() {
	authoritativeComponents := map[EntityComponent]bool{}

	connection := MakeConnection()
	dispatcher := MakeDispatcher()

	fmt.Println("Checking if connected")
	if connection.Connect() {
		fmt.Println("Connected!")
		connection.SendLog("mylogger", "Hello, World!")

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
			fmt.Printf("GOT POS %d\n", entity_id, update.Coords.X, update.Coords.Y, update.Coords.Z)
		})

		for connection.IsConnected() {

			dispatcher.dispatchOps(connection.ReadOps())

			for ec, value := range (authoritativeComponents) {
				if value {
					if ec.component_id == 54 {
						x := math.Sin(float64(time.Now().UnixNano()) / 1000000000.0) * 10.0
						newCoordinates := Coordinates{x, 2.0, 3.0}
						connection.SendPositionUpdate(ec.entity_id, PositionUpdate{&newCoordinates})
					}
				}
			}
		}

	} else {
		fmt.Println("Did not Connect!")
	}
}
