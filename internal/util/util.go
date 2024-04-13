package util

import (
	"encoding/json"
	"net/http"
	"syscall"
	"unsafe"
)

func GetTermWidth() int {
	var dimensions [4]uint16
	syscall.Syscall(syscall.SYS_IOCTL, uintptr(0), syscall.TIOCGWINSZ, uintptr(unsafe.Pointer(&dimensions)))
	return int(dimensions[1])
}

type GenericResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func SendErrorResponse(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	if message == "" {
		message = http.StatusText(code)
	}
	res := ErrorResponse{Error: message}
	json.NewEncoder(w).Encode(res)
}

func SendGenericResponse(w http.ResponseWriter) {
	res := GenericResponse{Message: "success"}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
