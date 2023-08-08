package Collections

import (
	"errors"
	"fmt"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

// GetActiveCampaign
// Returns the current campaign using the user
func GetActiveCampaign(app core.App, userID string) (*models.Record, error) {
	records, err := app.Dao().FindRecordsByExpr("campaign", dbx.And(
		dbx.HashExp{"active": true, "user": userID},
	))

	if len(records) > 1 {
		return nil, errors.New("GetActiveCampaign: Too many active campaigns")
	}

	if err != nil {
		return nil, err
	}

	return records[0], nil
}

// GetAllCampaigns
//
//	Gets all campaigns
func GetAllCampaigns(app core.App, userID string) ([]*models.Record, error) {
	records, err := app.Dao().FindRecordsByExpr("campaign", dbx.HashExp{"user": userID})

	if err != nil {
		return nil, err
	}

	return records, err
}

// ToggleCampaignFlag
// Changes the campaign flag to what it is not
func ToggleCampaignFlag(app core.App, campaignID string) error {
	record, err := app.Dao().FindRecordById("campaign", campaignID)

	if err != nil {
		return err
	}

	user := record.Get("user")

	fmt.Println(user)
	err = campaignValidations(app, user)

	if err != nil {
		return err
	}

	if record.Get("active") == true {
		record.Set("active", false)
	} else {
		record.Set("active", true)
	}

	if err := app.Dao().SaveRecord(record); err != nil {
		return err
	}

	return nil
}

func campaignValidations(app core.App, userID any) error {
	if records, err := app.Dao().FindRecordsByExpr("campaign", dbx.HashExp{"user": userID, "active": true}); err != nil {
		return err
	} else if len(records) > 1 {
		return errors.New("too many active campaigns")
	}

	return nil
}
