package util

import (
	"encoding/json"
	"net/http"
)

type PageInfo struct {
	Total int `json:"total"`
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type MessageData struct {
	Message string `json:"message"`
}

type Response struct {
	Data  any      `json:"data,omitempty"`
	Error string   `json:"error,omitempty"`
	Page  PageInfo `json:"page,omitempty"`
	Code  int      `json:"code"`
}

func OKResponse(w http.ResponseWriter, data any) {
	SendSuccessResponse(w, http.StatusOK, data)
}

func OKResponseWithPage(w http.ResponseWriter, data any, page PageInfo) {
	SendSuccessResponseWithPage(w, http.StatusOK, data, page)
}

func CreatedResponse(w http.ResponseWriter, data any) {
	SendSuccessResponse(w, http.StatusCreated, data)
}

func BadRequestResponse(w http.ResponseWriter, message string) {
	SendErrorResponse(w, http.StatusBadRequest, message)
}

func ForbiddenResponse(w http.ResponseWriter, message string) {
	SendErrorResponse(w, http.StatusForbidden, message)
}

func NotFoundResponse(w http.ResponseWriter, message string) {
	SendErrorResponse(w, http.StatusNotFound, message)
}

func InternalErrorResponse(w http.ResponseWriter, message string) {
	SendErrorResponse(w, http.StatusInternalServerError, message)
}

func SendErrorResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	res := Response{
		Code:  code,
		Error: message,
	}
	json.NewEncoder(w).Encode(res)
}

func SendSuccessResponse(w http.ResponseWriter, code int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	res := Response{
		Code: code,
		Data: data,
	}
	json.NewEncoder(w).Encode(res)
}

func SendSuccessResponseWithPage(w http.ResponseWriter, code int, data any, page PageInfo) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	res := Response{
		Code: code,
		Data: data,
		Page: page,
	}
	json.NewEncoder(w).Encode(res)
}
