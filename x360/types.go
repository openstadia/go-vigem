package x360

import "github.com/openstadia/go-vigem"

type XUSB_REPORT struct {
	wButtons      uint16
	bLeftTrigger  uint8
	bRightTrigger uint8
	sThumbLX      int16
	sThumbLY      int16
	sThumbRX      int16
	sThumbRY      int16
}

type PVIGEM_X360_NOTIFICATION func(
	client vigem.ClientType,
	target vigem.TargetType,
	largeMotor uint8,
	smallMotor uint8,
	ledNumber uint8,
	userData vigem.UserDataType) uintptr
