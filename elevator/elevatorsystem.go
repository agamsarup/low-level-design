package elevator

import (
	"errors"
	"sync"
)

type IDestFloorRequestHandler interface {
	HandleDestFloorRequest(elevatorID int, floorNo int)
}

type IFloorDirectionRequestHandler interface {
	HandleDirectionFromFloorRequest(int, Direction)
}

type IElevatorSystem interface {
	AddFloor(int) error
	AddElevatorCar(int, int, int) error
	IDestFloorRequestHandler
	IFloorDirectionRequestHandler
}

var lock sync.Mutex
var elevatorSystem *ElevatorSystem

func GetElevatorSystemInstance() *ElevatorSystem {
	if elevatorSystem == nil {
		lock.Lock()
		defer lock.Unlock()
		if elevatorSystem == nil {
			elevatorSystem = &ElevatorSystem{}
		}
	}

	return elevatorSystem
}

type ElevatorSystem struct {
	floorPanels []*FloorPanel
	elevators   []*ElevatorCar
}

func (e *ElevatorSystem) HandleDestFloorRequest(elevatorId, destFloorNo int) {
	//TODO
}

func (e *ElevatorSystem) HandleDirectionFromFloorRequest(fromFloor int, direction Direction) {
	//TODO
}

func (e *ElevatorSystem) AddElevatorCar(id int, minFloor, maxFloor int) error {
	for _, c := range e.elevators {
		if c.GetID() == id {
			return errors.New("elevator already added")
		}
	}

	e.elevators = append(e.elevators, newElevatorCar(id, minFloor, maxFloor, e))
	return nil
}

func (e *ElevatorSystem) AddFloor(floorNo int) error {
	for _, f := range e.floorPanels {
		if f.floorNo == floorNo {
			return errors.New("floor already added")
		}
	}

	e.floorPanels = append(e.floorPanels, newFloorPanel(floorNo, e))
	return nil
}
