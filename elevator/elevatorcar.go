package elevator

import "math"

type ElevatorState int

const (
	Idle ElevatorState = iota
	GoingUp
	GoingDown
)

type ElevatorDoor struct {
}

type IElevatorCar interface {
	GetID() int
	EnableOrDisableFloor(int, bool) (bool, error)
	ForwardFloorRequest(int)
	// move
	// stop
	// open door
	// close door
}

type ElevatorCar struct {
	id        int
	door      ElevatorDoor
	state     ElevatorState
	direction Direction
	curFloor  int
	display   *Display
	panel     *ElevatorPanel
	system    IElevatorSystem
}

func newElevatorCar(id int, minFloor, maxFloor int, system IElevatorSystem) *ElevatorCar {
	elevatorCar := &ElevatorCar{
		id:        id,
		door:      ElevatorDoor{},
		state:     Idle,
		direction: NA,
		curFloor:  math.MinInt,
		display:   newDisplay(),
		system:    system,
	}
	elevatorCar.panel = newInnerPanel(elevatorCar, minFloor, maxFloor)

	return elevatorCar
}

func (c *ElevatorCar) ForwardFloorRequest(floorNo int) {
	c.system.HandleDestFloorRequest(c.id, floorNo)
}

func (c *ElevatorCar) GetID() int {
	return c.id
}

func (c *ElevatorCar) EnableOrDisableFloor(f int, flag bool) (bool, error) {
	return c.panel.EnableOrDisableFloorButton(f, flag)
}
