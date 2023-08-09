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

func ToggleHabitCompletion(app core.App, habitCompletionID string) error {
	record, err := app.Dao().FindRecordById("habit_completion", habitCompletionID)
	if err != nil {
		return err
	}

	completionFlag := record.Get("complete")

	habitRecord, err := app.Dao().FindRecordById("habit", record.GetString("habit"))
	if err != nil {
		return err
	}

	campaignRecord, err := app.Dao().FindRecordById("campaign", habitRecord.GetString("campaign"))
	if err != nil {
		return err
	}

	if completionFlag == true {
		record.Set("complete", false)
		if err = decreaseXP(app, campaignRecord.GetId()); err != nil {
			return err
		}
	} else if completionFlag == false {
		record.Set("complete", true)
		if err = increaseXP(app, campaignRecord.GetId()); err != nil {
			return err
		}
	} else {
		return errors.New("message: error in toggle habit completion")
	}

	if err := app.Dao().SaveRecord(record); err != nil {
		return err
	}

	return nil
}
