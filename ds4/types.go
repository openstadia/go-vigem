package ds4

import "github.com/openstadia/go-vigem"

type DS4_REPORT struct {
	bThumbLX  uint8
	bThumbLY  uint8
	bThumbRX  uint8
	bThumbRY  uint8
	wButtons  uint16
	bSpecial  uint8
	bTriggerL uint8
	bTriggerR uint8
}

type DS4_REPORT_EX = uintptr

type DS4_LIGHTBAR_COLOR struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

type PVIGEM_DS4_NOTIFICATION func(
	client vigem.ClientType,
	target vigem.TargetType,
	largeMotor uint8,
	smallMotor uint8,
	ledNumber uint8,
	userData vigem.UserDataType)

type DS4_AWAIT_OUTPUT_BUFFER struct {
}
