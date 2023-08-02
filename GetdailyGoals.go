package main

import (
	"fmt"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	"net/http"
)

func GetDailyGoals(app core.App) {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/base/dailygoals/", func(c echo.Context) error {
			//Get auth user
			user, _ := c.Get(apis.ContextAuthRecordKey).(*models.Record)

			if user == nil {
				return c.JSON(http.StatusForbidden, "Only Authenticated users can access this endpoint")
			}

			currentCampaign := GetCampaign(app, user.GetId())

			currentHabits := GetCurrentHabits(app, currentCampaign[0].GetId())

			var currentHabitIds []interface{}

			for _, element := range currentHabits {
				currentHabitIds = append(currentHabitIds, element.GetId())
			}

			currentHabitCompletion := GetDailyHabitCompletion(app, currentHabitIds)

			return c.JSON(
				http.StatusOK,
				currentHabitCompletion,
			)
		})

		return nil
	})
}

func TestEndpoint(app core.App) {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/base/testpath/", func(c echo.Context) error {
			//Get auth user
			user, _ := c.Get(apis.ContextAuthRecordKey).(*models.Record)

			currentCampaign := GetCampaign(app, user.GetId())
			fmt.Print(currentCampaign[0].Get("name"))

			currentHabits := GetCurrentHabits(app, currentCampaign[0].GetId())
			var currentHabitIds []interface{}

			for _, element := range currentHabits {
				currentHabitIds = append(currentHabitIds, element.GetId())
			}
			fmt.Print(currentHabitIds)

			currentHabitCompletion := GetDailyHabitCompletion(app, currentHabitIds)
			fmt.Print(currentHabitCompletion)

			return nil
		})
		return nil
	})
}
