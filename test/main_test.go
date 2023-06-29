package main

import (
	"net/http"
	s "strings"
	"testing"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tests"
)

func dateStringProcess(app core.App) {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.AddRoute(echo.Route{
			Method: http.MethodGet,
			Path:   "/test/:week",
			Handler: func(c echo.Context) error {
				w := process(c.PathParam("week"))
				return c.String(200, w)
			},
			Middlewares: []echo.MiddlewareFunc{
				apis.ActivityLogger(app),
			},
		})

		return nil
	})
}

func process(dateString string) string {
	return s.Replace(dateString, "e", "3", -1)
}

const data_dir = "./pb_data"

func TestEndP(t *testing.T) {

	setupTestApp := func() (*tests.TestApp, error) {
		testApp, err := tests.NewTestApp(data_dir)
		if err != nil {
			return nil, err
		}

		dateStringProcess(testApp)

		return testApp, nil
	}

	scenarios := []tests.ApiScenario{
		{
			Name:            "Basic get test",
			Method:          http.MethodGet,
			Url:             "/test/week",
			ExpectedStatus:  200,
			ExpectedContent: []string{"w33k"},
			TestAppFactory:  setupTestApp,
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}
