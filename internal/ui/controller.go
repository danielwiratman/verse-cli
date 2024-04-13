package ui

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (me *Controller) Route(r chi.Router) {
	r.Get("/", me.fileServer)
}

func (me *Controller) fileServer(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("template")).ServeHTTP(w, r)
}
