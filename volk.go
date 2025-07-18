package volk

/*
#cgo CFLAGS: -Ivolk
#cgo LDFLAGS: -ldl
#include "volk.h"
*/
import "C"
import "unsafe"

func Initialize() error {
	if C.volkInitialize() != 0 {
		return nil
	}
	return nil
}

func LoadInstance(instance uintptr) {
	C.volkLoadInstance((C.VkInstance)(unsafe.Pointer(instance)))
}

func GetProcAddr() unsafe.Pointer {
	return unsafe.Pointer(C.volkGetInstanceProcAddr())
}
