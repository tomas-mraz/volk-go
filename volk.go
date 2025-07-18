package volk

/*
#cgo CFLAGS: -Ivolk
#cgo LDFLAGS: -Lvolk -lvolk -ldl
#include "volk.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func Initialize() error {
	result := C.volkInitialize()
	if result != 0 {
		return fmt.Errorf("volkInitialize failed with error code: %d", int(result))
	}
	return nil
}

func LoadInstance(instance unsafe.Pointer) {
	C.volkLoadInstance((C.VkInstance)(instance))
}

func LoadDevice(device unsafe.Pointer) {
	C.volkLoadDevice((C.VkDevice)(device))
}
