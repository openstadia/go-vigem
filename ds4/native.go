package ds4

import (
	"github.com/openstadia/go-vigem"
	"unsafe"
)

func TargetDS4Alloc() vigem.TargetType {
	ret, _, _ := vigem.ProcTargetDS4Alloc.Call()
	return ret
}

func TargetDS4RegisterNotification(client vigem.ClientType, target vigem.TargetType, notification PVIGEM_DS4_NOTIFICATION) vigem.ErrorType {
	ret, _, _ := vigem.ProcTargetDS4RegisterNotification.Call(client, target, uintptr(unsafe.Pointer(&notification)))
	return vigem.ErrorType(ret)
}

func TargetDS4UnregisterNotification(target vigem.TargetType) {
	_, _, _ = vigem.ProcTargetDS4UnregisterNotification.Call(target)
}

func TargetDS4Update(client vigem.ClientType, target vigem.TargetType, report DS4_REPORT) vigem.ErrorType {
	ret, _, _ := vigem.ProcTargetDS4Update.Call(client, target, uintptr(unsafe.Pointer(&report)))
	return vigem.ErrorType(ret)
}

func TargetDS4UpdateEx(client vigem.ClientType, target vigem.TargetType, report DS4_REPORT_EX) vigem.ErrorType {
	ret, _, _ := vigem.ProcTargetDS4UpdateEx.Call(client, target, report)
	return vigem.ErrorType(ret)
}

func TargetDS4AwaitOutputReport(client vigem.ClientType, target vigem.TargetType, buffer *DS4_AWAIT_OUTPUT_BUFFER) vigem.ErrorType {
	ret, _, _ := vigem.ProcTargetDS4AwaitOutputReport.Call(client, target, uintptr(unsafe.Pointer(buffer)))
	return vigem.ErrorType(ret)
}

func TargetDS4AwaitOutputReportTimeout(client vigem.ClientType, target vigem.TargetType, milliseconds uint32, buffer *DS4_AWAIT_OUTPUT_BUFFER) vigem.ErrorType {
	ret, _, _ := vigem.ProcTargetDS4AwaitOutputReportTimeout.Call(client, target, uintptr(milliseconds), uintptr(unsafe.Pointer(buffer)))
	return vigem.ErrorType(ret)
}
