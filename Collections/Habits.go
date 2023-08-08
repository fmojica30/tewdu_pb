package Collections

import (
	"errors"
	"example/user/hello/Utils"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

// GetCurrentHabits Get a list of the habits in the current campaign
func GetCurrentHabits(app core.App, campaignID string) ([]*models.Record, error) {
	records, err := app.Dao().FindRecordsByExpr("habit", dbx.HashExp{"campaign": campaignID})

	if err != nil {
		return nil, err
	}
	return records, nil
}

// GetDailyHabitDetails Gets daily habit detail information
func GetDailyHabitDetails(app core.App, currentHabits []interface{}) ([]*models.Record, error) {
	records, err := app.Dao().FindRecordsByExpr("habit_completion",
		dbx.And(
			dbx.In("habit", currentHabits...),
			dbx.Between("date", Utils.GetJuly272023Beginning(), Utils.GetJuly272023End()),
		),
	)

	if err != nil {
		return nil, err
	}

	return records, nil
}

func ToggleHabitCompletion(app core.App, habitID string) error {
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
		return errors.New("message: error in toggle habit completion")
	}

	return nil
}
