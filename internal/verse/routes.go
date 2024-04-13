package verse

import (
	"bible-verse-generator/internal/object"
	"net/http"
)

func (me *Controller) setupRoutes() {
	me.Routes = []object.Route{
		{
			Method:    http.MethodGet,
			HandlerFn: me.HandleGetAllVerses,
			Path:      "/verses",
		},
		{
			Method:    http.MethodGet,
			HandlerFn: me.HandleGetVerseById,
			Path:      "/verse/{id}",
		},
		{
			Method:    http.MethodGet,
			HandlerFn: me.HandleGetRandomVerse,
			Path:      "/verse/random",
		},
		{
			Method:    http.MethodPost,
			HandlerFn: me.HandleNewVerse,
			Path:      "/verse",
		},
		{
			Method:    http.MethodDelete,
			HandlerFn: me.HandleDeleteVerseById,
			Path:      "/verse/{id}",
		},
	}
}
