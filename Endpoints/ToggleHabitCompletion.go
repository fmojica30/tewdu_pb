package Endpoints

import (
	"example/user/hello/Collections"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
	"net/http"
)

func ToggleHabitCompletion(app core.App) {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.PUT("/base/togglehabitcompletion/:habitCompletionID", func(c echo.Context) error {
			habitCompletionID := c.PathParam("habitCompletionID")

			if err := Collections.ToggleHabitCompletion(app, habitCompletionID); err != nil {
				return err
			}

			return c.JSON(http.StatusOK, "Habit completion updated")

		})

		return nil
	})
}
