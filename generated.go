package main

import (
	"github.com/rjfwhite/go-worker-api/example"
	"fmt"
)

type Coordinates struct {
	X float64
	Y float64
	Z float64
}

type Position struct {
	Coords Coordinates
}

type PositionUpdate struct {
	Coords *Coordinates
}
type CoordinatesUpdate struct {
	X *float64
	Y *float64
	Z *float64
}

func ReadPositionUpdate(input example.Schema_Object) PositionUpdate {
	result := PositionUpdate{}
	if example.Schema_GetObjectCount(input, 1) > 0 {
		value := ReadCoordinates(example.Schema_GetObject(input, 1))
		result.Coords = &value
	}
	return result
}

func ReadCoordinatesUpdate(input example.Schema_Object) CoordinatesUpdate {
	result := CoordinatesUpdate{}

	{
		value := 3
		fmt.Print(value)
	}

	{
		value := 4
		fmt.Print(value)
	}

	if example.Schema_GetDoubleCount(input, 1) > 0 {
		value := example.Schema_GetDouble(input, 1)
		result.X = &value
	}
	if example.Schema_GetDoubleCount(input, 2) > 0 {
		value := example.Schema_GetDouble(input, 2)
		result.Y = &value
	}
	if example.Schema_GetDoubleCount(input, 3) > 0 {
		value := example.Schema_GetDouble(input, 3)
		result.Z = &value
	}
	return result
}

func WritePositionUpdate(output example.Schema_Object, update PositionUpdate) {
	if update.Coords != nil {
		example.Schema_AddObject(output, 1)
		WriteCoordinates(example.Schema_GetObject(output, 1), *update.Coords)
	}
}

func WriteCoordinatesUpdate(output example.Schema_Object, update CoordinatesUpdate) {
	if update.X != nil {
		example.Schema_AddDouble(output, 1, *update.X)
	}
	if update.Y != nil {
		example.Schema_AddDouble(output, 2, *update.Y)
	}
	if update.Z != nil {
		example.Schema_AddDouble(output, 3, *update.Z)
	}
}

func ApplyPositionUpdate(data Position, update PositionUpdate) {
	if update.Coords != nil {
		data.Coords = *update.Coords
	}
}

func ReadCoordinates(input example.Schema_Object) Coordinates {
	result := Coordinates{}
	{
		value := example.Schema_GetDouble(input, 1)
		result.X = value
	}
	{
		value := example.Schema_GetDouble(input, 2)
		result.Y = value
	}
	{
		value := example.Schema_GetDouble(input, 3)
		result.Z = value
	}
	return result
}

func ReadListy(input example.Schema_Object) Coordinates {

	count := example.Schema_GetObjectCount(input, 1)
	for  i := uint(0); i < count; i++ {

	}

	result := Coordinates{}
	{
		value := example.Schema_GetDouble(input, 1)
		result.X = value
	}
	{
		value := example.Schema_GetDouble(input, 2)
		result.Y = value
	}
	{
		value := example.Schema_GetDouble(input, 3)
		result.Z = value
	}
	return result
}

func WriteCoordinates(output example.Schema_Object, value Coordinates) {
	example.Schema_AddDouble(output, 1, value.X)
	example.Schema_AddDouble(output, 2, value.Y)
	example.Schema_AddDouble(output, 3, value.Z)
}

func ReadPosition(input example.Schema_Object) Position {
	return Position{
		Coords: ReadCoordinates(example.Schema_GetObject(input, 1)),
	}
}

func WritePosition(output example.Schema_Object, value Position) {
	example.Schema_AddObject(output, 1)
	WriteCoordinates(example.Schema_GetObject(output, 1), value.Coords)
}
