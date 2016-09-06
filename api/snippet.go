package api

import (
	"net/http"

	"github.com/labstack/echo"
)

func TestApi() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "TestApi page")
	}
}
