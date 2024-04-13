package main

import (
	"fmt"
	"log"
	"strings"
	"time"
	"verse-cli/internal/db"
	"verse-cli/internal/object"
	"verse-cli/internal/util"
	"verse-cli/internal/verse"

	"github.com/apatters/go-wordwrap"
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

const MaxTermWidth = 80

func outputCLI(verse *verse.Verse, termWidth int) {
	fmt.Println()
	fmt.Println(strings.Repeat("=", termWidth))
	fmt.Println()
	fmt.Println(wordwrap.IndentWithWrap(termWidth-4, "    ", false, verse.Address))
	fmt.Println(wordwrap.IndentWithWrap(termWidth-4, "    ", false, verse.Content))
	fmt.Println()
	_, currentGMTsec := time.Now().Zone()
	currentGMThour := currentGMTsec / 3600
	formattedTime := fmt.Sprintf(
		"created %s",
		verse.Created.Add(time.Duration(currentGMThour)*time.Hour).Format("2006/01/02 15:04"),
	)
	fmt.Printf("%s%s\n", strings.Repeat(" ", termWidth-4-len(formattedTime)), formattedTime)
	fmt.Println()
	fmt.Println(strings.Repeat("=", termWidth))
	fmt.Println()
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
	verseService := verse.NewService(verseRepo)

	verse, err := verseService.GetRandomVerse()
	if err != nil {
		log.Fatal(err)
	}

	termWidth := util.GetTermWidth()
	if termWidth > MaxTermWidth {
		termWidth = MaxTermWidth
	}

	outputCLI(verse, termWidth)
}
