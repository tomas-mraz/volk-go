package volk

/*
#cgo CFLAGS: -Ivolk
#cgo LDFLAGS: -Lvolk -lvolk -ldl
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

func LoadInstance(instance unsafe.Pointer) {
	C.volkLoadInstance((C.VkInstance)(instance))
}

func LoadDevice(device unsafe.Pointer) {
	C.volkLoadDevice((C.VkDevice)(device))
}
