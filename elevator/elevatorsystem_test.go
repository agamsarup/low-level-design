package elevator

import "testing"

func TestElevatorSystem(t *testing.T) {
	system := GetElevatorSystemInstance()

	err := system.AddFloor(0)
	if err != nil {
		t.Error(err)
	}

	err = system.AddFloor(1)
	if err != nil {
		t.Error(err)
	}

	err = system.AddFloor(-1)
	if err != nil {
		t.Error(err)
	}

	err = system.AddElevatorCar(1, -1, 1)
	if err != nil {
		t.Error(err)
	}

}
