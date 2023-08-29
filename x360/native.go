package x360

import (
	"github.com/openstadia/go-vigem"
	"syscall"
	"unsafe"
)

func TargetX360Alloc() vigem.TargetType {
	ret, _, _ := vigem.ProcTargetX360Alloc.Call()
	return ret
}

func TargetX360Update(client vigem.ClientType, target vigem.TargetType, report XUSB_REPORT) vigem.ErrorType {
	ret, _, _ := vigem.ProcTargetX360Update.Call(client, target, uintptr(unsafe.Pointer(&report)))
	return vigem.ErrorType(ret)
}

func TargetX360RegisterNotification(client vigem.ClientType, target vigem.TargetType, notification PVIGEM_X360_NOTIFICATION) vigem.ErrorType {
	ret, _, _ := vigem.ProcTargetX360RegisterNotification.Call(client, target, syscall.NewCallback(notification))
	return vigem.ErrorType(ret)
}

func TargetX360UnregisterNotification(target vigem.TargetType) {
	_, _, _ = vigem.ProcTargetX360UnregisterNotification.Call(target)
}

func TargetX360GetUserIndex(client vigem.ClientType, target vigem.TargetType, index *uint32) vigem.ErrorType {
	ret, _, _ := vigem.ProcTargetX360GetUserIndex.Call(client, target, uintptr(unsafe.Pointer(index)))
	return vigem.ErrorType(ret)
}
