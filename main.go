package main

import (
	"example/user/hello/Endpoints"
	"github.com/pocketbase/pocketbase"
	"log"
)

func main() {
	app := pocketbase.New()

	Endpoints.GetDailyHabits(app)
	Endpoints.GetActiveCampaign(app)
	Endpoints.ToggleCampaignActiveFlag(app)
	Endpoints.ToggleHabitCompletion(app)

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
