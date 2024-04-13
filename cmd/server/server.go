package main

import (
	"bible-verse-generator/internal/db"
	"bible-verse-generator/internal/object"
	"bible-verse-generator/internal/verse"
	"log"
	"net/http"

	"github.com/caarlos0/env/v10"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func apiRouter(r chi.Router) {
}

func main() {
	godotenv.Load()
	conf := &object.Config{}
	err := env.ParseWithOptions(conf, env.Options{
		RequiredIfNoDef: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	db, err := db.New(conf)
	if err != nil {
		log.Fatal(err)
	}
	verseRepo := verse.NewDBRepo(db)
	verseController := verse.NewController(verseRepo)

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.Dir("template")).ServeHTTP(w, r)
	})
	r.Route("/api", func(r chi.Router) {
		for _, route := range verseController.Routes {
			r.MethodFunc(route.Method, route.Path, route.HandlerFn)
		}
	})

	log.Println("Listening on :8080")
	http.ListenAndServe(":8080", r)
}
