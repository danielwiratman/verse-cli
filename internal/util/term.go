package util

import (
	"syscall"
	"unsafe"
)

func GetTermWidth() int {
	var dimensions [4]uint16
	syscall.Syscall(syscall.SYS_IOCTL, uintptr(0), syscall.TIOCGWINSZ, uintptr(unsafe.Pointer(&dimensions)))
	return int(dimensions[1])
}
