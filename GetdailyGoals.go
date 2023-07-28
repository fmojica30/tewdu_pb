package main

import (
	"fmt"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
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

			//queryString := "SELECT goal_name FROM goals WHERE user_id = " + user.GetId()
			//queryString := "SELECT goal_name FROM goals"

			err := app.DB().
				Select("*").
				From("habit_completion").
				Where(dbx.HashExp{"user": user.GetId()}).
				AndWhere(dbx.Between("date", GetTodayBeginning(), GetTodayEnd()))

			if err != nil {
				fmt.Printf("Error on query")
			}

			return c.JSON(
				http.StatusOK,
				"test",
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

			fmt.Print(currentHabits[0])

			return nil
		})
		return nil
	})
}
