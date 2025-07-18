package volk

/*
#cgo CFLAGS: -Ivolk
#cgo LDFLAGS: -ldl
#include "volk.h"
*/
import "C"
import (
	"errors"
	"unsafe"
)

func Initialize() error {
	if C.volkInitialize() != 0 {
		return nil
	}
	return errors.New("volk initialization failed")
}

func LoadInstance(instance uintptr) {
	C.volkLoadInstance((C.VkInstance)(unsafe.Pointer(instance)))
}

func LoadDevice(device uintptr) {
	C.volkLoadDevice((C.VkDevice)(unsafe.Pointer(device)))
}

func GetProcAddr() unsafe.Pointer {
	return unsafe.Pointer(C.volkGetInstanceProcAddr())
}
