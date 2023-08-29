package main

import (
	"fmt"
	"github.com/openstadia/go-vigem"
	X360 "github.com/openstadia/go-vigem/x360"
	"time"
)

func callback(
	client vigem.ClientType,
	target vigem.TargetType,
	largeMotor uint8,
	smallMotor uint8,
	ledNumber uint8,
	userData vigem.UserDataType) uintptr {
	fmt.Printf("Callback %d %d %d\n", largeMotor, smallMotor, ledNumber)
	return 0
}

func main() {
	client := vigem.NewClient()
	x360 := X360.NewGamepad(client)
	x360.Connect()

	x360.RegisterNotification(callback)

	for i := 0; i < 5; i++ {
		fmt.Println("Press")
		x360.PressButton(X360.XUSB_GAMEPAD_A)
		x360.Update()
		time.Sleep(5 * time.Second)

		fmt.Println("Release")
		x360.ReleaseButton(0x1000)
		x360.Update()
		time.Sleep(5 * time.Second)
	}

	x360.UnregisterNotification()
	x360.Disconnect()
	x360.Release()

	client.Release()
}
