package Collections

import (
	"errors"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

// GetCurrentHabits Get a list of the habits in the current campaign
func GetCurrentHabits(app core.App, campaignID string) ([]*models.Record, error) {
	records, err := app.Dao().FindRecordsByExpr("habit_details", dbx.HashExp{"campaign": campaignID})

	if err != nil {
		return nil, err
	}
	return records, nil
}

// GetDailyHabitCompletion Gets daily habit detail information
func GetDailyHabitCompletion(app core.App, currentHabits []interface{}) ([]*models.Record, error) {
	records, err := app.Dao().FindRecordsByExpr("habit_completion",
		dbx.And(
			dbx.In("habit", currentHabits...),
			dbx.Between("date", GetJuly272023Beginning(), GetJuly272023End()),
		),
	)

	if err != nil {
		return nil, err
	}

	return records, nil
}

func ToggleHabitCompletion(app core.App, habitID string, userID string) error {
	record, err := app.Dao().FindRecordById("habit_completion", habitID)

	if err != nil {
		return err
	}

	completionFlag := record.Get("completion")

	if completionFlag == true {
		record.Set("completion", false)
	} else if completionFlag == false {
		record.Set("completion", true)
	} else {
		return errors.New("Message: error in Toggle Habit Completion")
	}

	//Add in level calculation here

	return nil
}
