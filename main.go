package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
)

func main() {
	app := pocketbase.New()

	DateStringProcess(app)
	DailyProgress(app)

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
