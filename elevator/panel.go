package elevator

import (
	"fmt"
)

type Panel struct {
}

type IElevatorPanel interface {
	EnableOrDisableFloorButton(int, bool) (bool, error)
}

type ElevatorPanel struct {
	*Panel
	floorButtons []IFloorButton
	openButton   *OpenButton
	closeButton  *CloseButton
	elevatorCar  IElevatorCar
}

func newInnerPanel(elevatorCar IElevatorCar, minFloor, maxFloor int) *ElevatorPanel {
	floorButtons := make([]IFloorButton, maxFloor-minFloor+1)
	for i := minFloor; i <= maxFloor; i++ {
		floorButtons = append(floorButtons, &FloorButton{
			Button:  &Button{},
			floorNo: i,
		})
	}
	return &ElevatorPanel{
		Panel:        &Panel{},
		floorButtons: floorButtons,
		openButton:   &OpenButton{},
		closeButton:  &CloseButton{},
		elevatorCar:  elevatorCar,
	}
}

func (p *ElevatorPanel) EnableOrDisableFloorButton(floorNo int, flag bool) (bool, error) {
	for _, b := range p.floorButtons {
		if b.GetFloorNo() == floorNo {
			b.SetEnable(flag)
			return true, nil
		}
	}

	return false, fmt.Errorf("elevator panel does not support floor no %d", floorNo)
}

type IDirectionRequestHandler interface {
	HandleDirectionRequest(Direction)
}

type FloorPanel struct {
	*Panel
	floorNo           int
	upButton          *UpButton
	downButton        *DownButton
	dirRequestHandler IFloorDirectionRequestHandler
}

func newFloorPanel(floorNo int, dirReqHandler IFloorDirectionRequestHandler) *FloorPanel {
	floorPanel := &FloorPanel{
		Panel:             &Panel{},
		floorNo:           floorNo,
		dirRequestHandler: dirReqHandler,
	}

	floorPanel.downButton = newDownButton(floorPanel)
	floorPanel.upButton = newUpButton(floorPanel)

	return floorPanel
}

func (f *FloorPanel) HandleDirectionRequest(d Direction) {
	f.dirRequestHandler.HandleDirectionFromFloorRequest(f.floorNo, d)
}
