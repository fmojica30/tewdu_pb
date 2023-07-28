package main

import (
	"github.com/pocketbase/pocketbase"
	"log"
)

func main() {
	app := pocketbase.New()

	DateStringProcess(app)
	DailyProgress(app)
	GetDailyGoals(app)
	TestEndpoint(app)

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
