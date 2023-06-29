package elevator

type IButton interface {
	Press()
	IsPressed() bool
	SetEnable(bool)
	IsEnabled() bool
}

type Button struct {
	isPressed bool
	isEnabled bool
}

func (b *Button) IsPressed() bool {
	return b.isPressed
}

func (b *Button) Press() {
	if b.IsEnabled() {
		if b.IsPressed() {
			b.isPressed = false
		} else {
			b.isPressed = true
		}
	}
}

func (b *Button) SetEnable(enable bool) {
	b.isEnabled = enable
	if !enable {
		b.isPressed = false
	}
}

func (b *Button) IsEnabled() bool {
	return b.isEnabled
}

type OpenButton struct {
	*Button
}

type CloseButton struct {
	*Button
}

type IFloorButton interface {
	IButton
	GetFloorNo() int
}

type FloorButton struct {
	*Button
	floorNo int
}

func (b *FloorButton) GetFloorNo() int {
	return b.floorNo
}

type UpButton struct {
	*Button
	dirReqHandler IDirectionRequestHandler
}

func newUpButton(dirReqHandler IDirectionRequestHandler) *UpButton {
	return &UpButton{
		Button: &Button{
			isPressed: false,
			isEnabled: true,
		},
		dirReqHandler: dirReqHandler,
	}
}

type DownButton struct {
	*Button
	dirReqHandler IDirectionRequestHandler
}

func newDownButton(dirReqHandler IDirectionRequestHandler) *DownButton {
	return &DownButton{
		Button: &Button{
			isPressed: false,
			isEnabled: true,
		},
		dirReqHandler: dirReqHandler,
	}
}

func (b *DownButton) Press() {
	if b.IsEnabled() && !b.IsPressed() {
		b.isPressed = true
		b.dirReqHandler.HandleDirectionRequest(Down)
	}
}
