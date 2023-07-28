package main

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	"time"
)

// GetTodayBeginning
// Gets the beginning of the current day
func GetTodayBeginning() time.Time {
	today := time.Now()
	year, month, day := today.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, today.Location())
}

// GetTodayEnd
// Returns the end of the current day
func GetTodayEnd() time.Time {
	today := time.Now()
	year, month, day := today.Date()
	return time.Date(year, month, day, 23, 59, 59, 0, today.Location())
}

// GetCampaign
// Returns the current campaign
func GetCampaign(app core.App, user string) []*models.Record {

	records, err := app.Dao().FindRecordsByExpr("campaign", dbx.And(
		dbx.HashExp{"active": true, "user": user},
	))

	if err != nil {
		return nil
	}

	return records
}

func GetCurrentHabits(app core.App, campaignID string) []*models.Record {

	records, err := app.Dao().FindRecordsByExpr("habit_details", dbx.HashExp{"campaign": campaignID})

	if err != nil {
		return nil
	}
	return records
}
