package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func progressCalculation(app core.App, c echo.Context, datString string) float64 {
	var complete float64
	var total float64

	app.DB().
		Select("count(*)").
		From("goals").
		Row(&total)

	app.DB().
		Select("count(*)").
		From("goals").
		Where(dbx.NewExp("goal_completion=1")).
		Row(&complete)

	return complete / total
}

func DailyProgress(app core.App) {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.AddRoute(echo.Route{
			Method: http.MethodGet,
			Path:   "/progress/day/:date-string",
			Handler: func(c echo.Context) error {
				w := progressCalculation(app, c, "test")
				s1 := fmt.Sprintf("%v", w)
				return c.String(200, s1)
			},
			Middlewares: []echo.MiddlewareFunc{
				apis.ActivityLogger(app),
			},
		})

		return nil
	})
}
