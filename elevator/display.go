package elevator

type Display struct {
	floorNo   int
	direction Direction
}

func newDisplay() *Display {
	return &Display{
		floorNo:   0,
		direction: NA,
	}
}
