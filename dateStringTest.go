package main

import (
	"net/http"
	s "strings"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func Process(dateString string) string {
	return s.Replace(dateString, "e", "3", -1)
}

func DateStringProcess(app core.App) {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.AddRoute(echo.Route{
			Method: http.MethodGet,
			Path:   "/test/:week",
			Handler: func(c echo.Context) error {
				w := Process(c.PathParam("week"))
				return c.String(200, w)
			},
			Middlewares: []echo.MiddlewareFunc{
				apis.ActivityLogger(app),
			},
		})

		return nil
	})
}
