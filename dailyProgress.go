package main

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

type dataDTO struct {
	Completion float64
}

func progressCalculation(app core.App, c echo.Context, datString string) (float64, error) {
	var complete float64
	var total float64

	app.DB().
		Select("count(*)").
		From("goals").
		Row(&complete)

	app.DB().
		Select("count(*)").
		From("goals").
		Where(dbx.NewExp("goal_completion=1")).
		Row(&complete)

	return complete / total, nil
}

func DailyProgress(app core.App) {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.AddRoute(echo.Route{
			Method: http.MethodGet,
			Path:   "/progress/day/:date-string",
			Handler: func(c echo.Context) error {
				w, err := progressCalculation(app, c, "test")

				data := dataDTO{
					Completion: w,
				}

				if err != nil {
					return c.JSON(http.StatusExpectationFailed, w)
				}
				//s1 := fmt.Sprintf("%v", w)
				//return c.String(200, s1)
				return c.JSON(http.StatusOK, data)
			},
			Middlewares: []echo.MiddlewareFunc{
				apis.ActivityLogger(app),
			},
		})

		return nil
	})
}
