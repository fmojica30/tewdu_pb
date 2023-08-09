package Endpoints

import (
	"example/user/hello/Collections"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	"net/http"
)

func GetDailyHabits(app core.App) {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET(
			"/base/dailygoals/",
			func(c echo.Context) error {
				//Get auth user
				user, _ := c.Get(apis.ContextAuthRecordKey).(*models.Record)
				if user == nil {
					return c.JSON(http.StatusForbidden, "Only Authenticated users can access this endpoint")
				}

				currentCampaign, err := Collections.GetActiveCampaign(app, user.GetId())
				if err != nil {
					return c.JSON(http.StatusInternalServerError, err.Error())
				}

				currentHabits, err := Collections.GetCurrentHabits(app, currentCampaign.GetId())
				if err != nil {
					return c.JSON(http.StatusInternalServerError, err.Error())
				}

				var currentHabitIds []interface{}

				for _, element := range currentHabits {
					currentHabitIds = append(currentHabitIds, element.GetId())
				}

				currentHabitCompletion, err := Collections.GetDailyHabitDetails(app, currentHabitIds)

				if err != nil {
					return c.JSON(http.StatusInternalServerError, err.Error())
				}

				return c.JSON(
					http.StatusOK,
					currentHabitCompletion,
				)
			},
		)

		return nil
	})
}
