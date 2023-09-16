package x360

import "github.com/openstadia/go-vigem"

func x360DefaultReport() XUSB_REPORT {
	return XUSB_REPORT{
		wButtons:      0,
		bLeftTrigger:  0,
		bRightTrigger: 0,
		sThumbLX:      0,
		sThumbLY:      0,
		sThumbRX:      0,
		sThumbRY:      0,
	}
}

type Gamepad struct {
	*vigem.GamepadBase
	report       XUSB_REPORT
	notification PVIGEM_X360_NOTIFICATION
}

func NewGamepad(client vigem.Client) *Gamepad {
	nativeHandle := TargetX360Alloc()

	gamepad := vigem.NewGamepad(client, nativeHandle)

	x360Gamepad := Gamepad{
		GamepadBase: gamepad,
		report:      x360DefaultReport(),
	}

	return &x360Gamepad
}

func (x *Gamepad) Reset() {
	x.report = x360DefaultReport()
}

func (x *Gamepad) PressButton(button uint16) {
	x.report.wButtons |= button
}

func (x *Gamepad) ReleaseButton(button uint16) {
	x.report.wButtons &= ^button
}

func (x *Gamepad) LeftTrigger(value uint8) {
	x.report.bLeftTrigger = value
}

func (x *Gamepad) RightTrigger(value uint8) {
	x.report.bRightTrigger = value
}

func (x *Gamepad) LeftJoystick(xValue int16, yValue int16) {
	x.report.sThumbLX = xValue
	x.report.sThumbLY = yValue
}

func (x *Gamepad) RightJoystick(xValue int16, yValue int16) {
	x.report.sThumbRX = xValue
	x.report.sThumbRY = yValue
}

func (x *Gamepad) Update() {
	TargetX360Update(x.Client.NativeHandle(), x.NativeHandle, x.report)
}

func (x *Gamepad) TargetAlloc() vigem.TargetType {
	return TargetX360Alloc()
}

func (x *Gamepad) RegisterNotification(notification PVIGEM_X360_NOTIFICATION) {
	x.notification = notification
	TargetX360RegisterNotification(x.Client.NativeHandle(), x.NativeHandle, notification)
}

func (x *Gamepad) UnregisterNotification() {
	TargetX360UnregisterNotification(x.NativeHandle)
	x.notification = nil
}

type XUSB_BUTTON = uint16

const (
	XUSB_GAMEPAD_DPAD_UP        XUSB_BUTTON = 0x0001
	XUSB_GAMEPAD_DPAD_DOWN      XUSB_BUTTON = 0x0002
	XUSB_GAMEPAD_DPAD_LEFT      XUSB_BUTTON = 0x0004
	XUSB_GAMEPAD_DPAD_RIGHT     XUSB_BUTTON = 0x0008
	XUSB_GAMEPAD_START          XUSB_BUTTON = 0x0010
	XUSB_GAMEPAD_BACK           XUSB_BUTTON = 0x0020
	XUSB_GAMEPAD_LEFT_THUMB     XUSB_BUTTON = 0x0040
	XUSB_GAMEPAD_RIGHT_THUMB    XUSB_BUTTON = 0x0080
	XUSB_GAMEPAD_LEFT_SHOULDER  XUSB_BUTTON = 0x0100
	XUSB_GAMEPAD_RIGHT_SHOULDER XUSB_BUTTON = 0x0200
	XUSB_GAMEPAD_GUIDE          XUSB_BUTTON = 0x0400
	XUSB_GAMEPAD_A              XUSB_BUTTON = 0x1000
	XUSB_GAMEPAD_B              XUSB_BUTTON = 0x2000
	XUSB_GAMEPAD_X              XUSB_BUTTON = 0x4000
	XUSB_GAMEPAD_Y              XUSB_BUTTON = 0x8000
)
