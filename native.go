package vigem

import (
	"syscall"
)

type ClientType = uintptr

type TargetType = uintptr

type TargetAddResultType = uintptr

type UserDataType = uintptr

type ErrorType = uint32

type TargetTypeType = uint32

var (
	dll                                   = syscall.MustLoadDLL("vigemclient.dll")
	ProcAlloc                             = dll.MustFindProc("vigem_alloc")
	ProcFree                              = dll.MustFindProc("vigem_free")
	ProcConnect                           = dll.MustFindProc("vigem_connect")
	ProcDisconnect                        = dll.MustFindProc("vigem_disconnect")
	ProcTargetX360Alloc                   = dll.MustFindProc("vigem_target_x360_alloc")
	ProcTargetDS4Alloc                    = dll.MustFindProc("vigem_target_ds4_alloc")
	ProcTargetFree                        = dll.MustFindProc("vigem_target_free")
	ProcTargetAdd                         = dll.MustFindProc("vigem_target_add")
	ProcTargetAddAsync                    = dll.MustFindProc("vigem_target_add_async")
	ProcTargetRemove                      = dll.MustFindProc("vigem_target_remove")
	ProcTargetX360RegisterNotification    = dll.MustFindProc("vigem_target_x360_register_notification")
	ProcTargetDS4RegisterNotification     = dll.MustFindProc("vigem_target_ds4_register_notification")
	ProcTargetX360UnregisterNotification  = dll.MustFindProc("vigem_target_x360_unregister_notification")
	ProcTargetDS4UnregisterNotification   = dll.MustFindProc("vigem_target_ds4_unregister_notification")
	ProcTargetSetVid                      = dll.MustFindProc("vigem_target_set_vid")
	ProcTargetSetPid                      = dll.MustFindProc("vigem_target_set_pid")
	ProcTargetGetVid                      = dll.MustFindProc("vigem_target_get_vid")
	ProcTargetGetPid                      = dll.MustFindProc("vigem_target_get_pid")
	ProcTargetX360Update                  = dll.MustFindProc("vigem_target_x360_update")
	ProcTargetDS4Update                   = dll.MustFindProc("vigem_target_ds4_update")
	ProcTargetDS4UpdateEx                 = dll.MustFindProc("vigem_target_ds4_update_ex")
	ProcTargetGetIndex                    = dll.MustFindProc("vigem_target_get_index")
	ProcTargetGetType                     = dll.MustFindProc("vigem_target_get_type")
	ProcTargetIsAttached                  = dll.MustFindProc("vigem_target_is_attached")
	ProcTargetX360GetUserIndex            = dll.MustFindProc("vigem_target_x360_get_user_index")
	ProcTargetDS4AwaitOutputReport        = dll.MustFindProc("vigem_target_ds4_await_output_report")
	ProcTargetDS4AwaitOutputReportTimeout = dll.MustFindProc("vigem_target_ds4_await_output_report_timeout")
)

func Alloc() ClientType {
	ret, _, _ := ProcAlloc.Call()
	return ret
}

func Free(client ClientType) {
	_, _, _ = ProcFree.Call(client)
}

func Connect(client ClientType) ErrorType {
	ret, _, _ := ProcConnect.Call(client)
	return ErrorType(ret)
}

func Disconnect(client ClientType) {
	_, _, _ = ProcDisconnect.Call(client)
}

func TargetFree(target TargetType) {
	_, _, _ = ProcTargetFree.Call(target)
}

func TargetAdd(client ClientType, target TargetType) ErrorType {
	ret, _, _ := ProcTargetAdd.Call(client, target)
	return ErrorType(ret)
}

func TargetAddAsync(client ClientType, target TargetType, result TargetAddResultType) ErrorType {
	ret, _, _ := ProcTargetAddAsync.Call(client, target, result)
	return ErrorType(ret)
}

func TargetRemove(client ClientType, target TargetType) ErrorType {
	ret, _, _ := ProcTargetRemove.Call(client, target)
	return ErrorType(ret)
}

func TargetSetVid(target TargetType, vid uint16) {
	_, _, _ = ProcTargetSetVid.Call(target, uintptr(vid))
}

func TargetSetPid(target TargetType, pid uint16) {
	_, _, _ = ProcTargetSetPid.Call(target, uintptr(pid))
}

func TargetGetVid(target TargetType) uint16 {
	ret, _, _ := ProcTargetGetVid.Call(target)
	return uint16(ret)
}

func TargetGetPid(target TargetType) uint16 {
	ret, _, _ := ProcTargetGetPid.Call(target)
	return uint16(ret)
}

func TargetGetIndex(target TargetType) uint {
	ret, _, _ := ProcTargetGetIndex.Call(target)
	return uint(ret)
}

func TargetGetType(target TargetType) TargetTypeType {
	ret, _, _ := ProcTargetGetType.Call(target)
	return TargetTypeType(ret)
}

func TargetIsAttached(target TargetType) bool {
	ret, _, _ := ProcTargetIsAttached.Call(target)
	return ret != 0
}

const (
	VIGEM_ERROR_NONE                        ErrorType = 0x20000000
	VIGEM_ERROR_BUS_NOT_FOUND               ErrorType = 0xE0000001
	VIGEM_ERROR_NO_FREE_SLOT                ErrorType = 0xE0000002
	VIGEM_ERROR_INVALID_TARGET              ErrorType = 0xE0000003
	VIGEM_ERROR_REMOVAL_FAILED              ErrorType = 0xE0000004
	VIGEM_ERROR_ALREADY_CONNECTED           ErrorType = 0xE0000005
	VIGEM_ERROR_TARGET_UNINITIALIZED        ErrorType = 0xE0000006
	VIGEM_ERROR_TARGET_NOT_PLUGGED_IN       ErrorType = 0xE0000007
	VIGEM_ERROR_BUS_VERSION_MISMATCH        ErrorType = 0xE0000008
	VIGEM_ERROR_BUS_ACCESS_FAILED           ErrorType = 0xE0000009
	VIGEM_ERROR_CALLBACK_ALREADY_REGISTERED ErrorType = 0xE0000010
	VIGEM_ERROR_CALLBACK_NOT_FOUND          ErrorType = 0xE0000011
	VIGEM_ERROR_BUS_ALREADY_CONNECTED       ErrorType = 0xE0000012
	VIGEM_ERROR_BUS_INVALID_HANDLE          ErrorType = 0xE0000013
	VIGEM_ERROR_XUSB_USERINDEX_OUT_OF_RANGE ErrorType = 0xE0000014
	VIGEM_ERROR_INVALID_PARAMETER           ErrorType = 0xE0000015
	VIGEM_ERROR_NOT_SUPPORTED               ErrorType = 0xE0000016
	VIGEM_ERROR_WINAPI                      ErrorType = 0xE0000017
	VIGEM_ERROR_TIMED_OUT                   ErrorType = 0xE0000018
	VIGEM_ERROR_IS_DISPOSING                ErrorType = 0xE0000019
)

const (
	Xbox360Wired    TargetTypeType = 0
	DualShock4Wired TargetTypeType = 2
)
