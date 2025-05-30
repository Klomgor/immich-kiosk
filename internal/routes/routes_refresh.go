package routes

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"

	"github.com/damongolding/immich-kiosk/internal/config"
	"github.com/damongolding/immich-kiosk/internal/utils"
)

// RefreshCheck endpoint to check if device requires a refresh
func RefreshCheck(baseConfig *config.Config) echo.HandlerFunc {
	return func(c echo.Context) error {

		kioskVersionHeader := c.Request().Header.Get("kiosk-version")
		kioskRefreshTimestampHeader := c.Request().Header.Get("kiosk-reload-timestamp")
		requestID := utils.ColorizeRequestID(c.Response().Header().Get(echo.HeaderXRequestID))

		// create a copy of the global config to use with this request
		requestConfig := *baseConfig

		if KioskVersion != kioskVersionHeader || kioskRefreshTimestampHeader != requestConfig.ReloadTimeStamp {
			c.Response().Header().Set("HX-Refresh", "true")
			return c.NoContent(http.StatusNoContent)
		}

		log.Debug(
			requestID,
			"method", c.Request().Method,
			"path", c.Request().URL.String(),
		)

		return c.NoContent(http.StatusNoContent)
	}
}
