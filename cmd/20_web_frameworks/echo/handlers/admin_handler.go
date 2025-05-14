package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AdminDashboardHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Admin Dashboard")
}

func AdminSettingsHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Admin Settings")
}
