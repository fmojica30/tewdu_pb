package Endpoints

import (
	"example/user/hello/Collections"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
	"net/http"
)

func ToggleCampaignActiveFlag(app core.App) {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.PUT(
			"/base/togglecampaign/:campaignID",
			func(c echo.Context) error {
				campaignID := c.PathParam("campaignID")

				err := Collections.ToggleCampaignFlag(app, campaignID)

				if err != nil {
					return c.JSON(http.StatusInternalServerError, err)
				}

				return c.JSON(http.StatusOK, "campaign updated")

			},
		)

		return nil
	})
}
