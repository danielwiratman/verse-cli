package main

import (
	"flag"
	"log"
	"net/http"
	"verse-cli/internal/db"
	"verse-cli/internal/object"
	"verse-cli/internal/ui"
	"verse-cli/internal/verse"

	"github.com/caarlos0/env/v10"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

var ServerPort *string

func parseFlags() {
	ServerPort = flag.String("port", "8080", "port to run server on")
	flag.Parse()
}

func main() {
	parseFlags()

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

	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		verseRepo := verse.NewDBRepo(db)
		verseController := verse.NewController(verseRepo)
		verseController.Route(r)
	})

	r.Route("/", func(r chi.Router) {
		uiController := ui.NewController()
		uiController.Route(r)
	})

	log.Printf("Listening on :%s\n", *ServerPort)
	http.ListenAndServe(":"+*ServerPort, r)
}
