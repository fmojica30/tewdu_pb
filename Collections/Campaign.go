package Collections

import (
	"errors"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

// GetActiveCampaign
// Returns the current campaign using the user
func GetActiveCampaign(app core.App, userID string) ([]*models.Record, error) {
	records, err := app.Dao().FindRecordsByExpr("campaign", dbx.And(
		dbx.HashExp{"active": true, "user": userID},
	))

	if err != nil {
		return nil, err
	}

	return records, nil
}

// GetAllCampaigns Gets all campaigns
func GetAllCampaigns(app core.App, userID string) ([]*models.Record, error) {
	records, err := app.Dao().FindRecordsByExpr("campaign", dbx.HashExp{"user": userID})

	if err != nil {
		return nil, err
	}

	return records, err
}

// ToggleCampaignFlag Changes the campaign flag to what it is not
func ToggleCampaignFlag(app core.App, campaignID string) error {
	record, err := app.Dao().FindRecordById("campaign", campaignID)

	if err != nil {
		return err
	}

	if record.Get("active") == true {
		record.Set("active", false)
	} else {
		record.Set("active", true)
	}

	return nil
}

// CampaignValidations Validates if there is only one campaign active at a time
func CampaignValidations(app core.App, userID string) error {
	activeCampaignRecords, _ := GetActiveCampaign(app, userID)

	if len(activeCampaignRecords) != 1 {
		return errors.New("CampaignValidations: Number of active campaigns is too much for this user")
	}

	return nil
}
