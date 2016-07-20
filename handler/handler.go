package handler

import (
    "net/http"

    "github.com/labstack/echo"
)

func HomePage() echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.String(http.StatusOK, "This page is top page")
    }
}
