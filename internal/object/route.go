package object

import (
	"net/http"
)

type Route struct {
	Method    string
	HandlerFn http.HandlerFunc
	Path      string
}
