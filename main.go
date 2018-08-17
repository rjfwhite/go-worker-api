package main

import (
	tm "github.com/buger/goterm"
	"time"
	"fmt"
	"math"
	"github.com/eiannone/keyboard"
)

var (
	connection          *Connection
	positionAuthorities       = make(map[int64]bool)
	positions                 = make(map[int64]Coordinates)
	playerEntityId      int64 = 0
	x                         = 0
	y                         = 0
)

func main() {
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	connection = MakeConnection("Managed")
	dispatcher := MakeDispatcher()

	fmt.Println("Checking if connected")
	if connection.Connect("localhost", 7777) {
		fmt.Println("Connected!")
		connection.SendLog("mylogger", "Hello, World!")

		dispatcher.OnEntityAdded(func(entityId int64) {
			fmt.Printf("ENTITY ADDED %d", entityId)
		})

		dispatcher.OnPositionAdded(func(entityId int64, data Position) {
			positions[entityId] = data.Coords
			fmt.Printf("GOT POS %d\n", entityId, data.Coords.X, data.Coords.Y, data.Coords.Z)
		})

		dispatcher.OnEntityAclAdded(func(entityId int64, data EntityAcl) {
			fmt.Printf("GOT ACL %d\n", entityId, data.Read, data.Write)
		})

		dispatcher.OnPositionAuthority(func(entityId int64, isAuthoritative bool) {
			playerEntityId = entityId
			positionAuthorities[entityId] = isAuthoritative
		})

		dispatcher.OnPositionUpdated(func(entityId int64, update PositionUpdate) {
			if update.Coords != nil {
				positions[entityId] = *update.Coords
			}
		})

		for connection.IsConnected() {
			dispatcher.dispatchOps(connection.ReadOps(0))
			for entityId, isAuthoritative := range positionAuthorities {
				if isAuthoritative && entityId != playerEntityId {
					x := 5.0 + math.Sin(float64(time.Now().UnixNano())/1000000000.0)*5.0
					newCoordinates := Coordinates{x, 2.0, 3.0}
					connection.SendPositionUpdate(entityId, PositionUpdate{&newCoordinates})
				}
			}
			applyUpdates()
			renderScreen()
			time.Sleep(time.Millisecond * 100)
		}

	} else {
		fmt.Println("Did not Connect!")
	}
}

func applyUpdates() {

}

func renderScreen() {
	tm.Clear()
	for _, position := range positions {
		tm.MoveCursor(int(position.X), int(position.Z))
		tm.Println(tm.Color("🌳", tm.GREEN))
	}

	_, key, _ := keyboard.GetKeyAsync()

	switch key {
	case keyboard.KeyArrowUp:
		y -= 1
	case keyboard.KeyArrowDown:
		y += 1
	case keyboard.KeyArrowRight:
		x += 1
	case keyboard.KeyArrowLeft:
		x -= 1
	}

	if playerEntityId != 0 {
		coordinates := Coordinates{float64(x), 0, float64(-y)}
		connection.SendPositionUpdate(playerEntityId, PositionUpdate{&coordinates})
	}

	tm.MoveCursor(int(x), int(y))
	tm.MoveCursor(0, 20)
	tm.Println(tm.Background("SpatialOS Golang Worker                                        ", tm.RED))
	//tm.Println(tm.Color("🚶", tm.BLUE))

	tm.Flush()
}
