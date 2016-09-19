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

func Home(c echo.Context) error {
	return c.Render(http.StatusOK, "home", "Hello")
}

func Snippet(c echo.Context) error {
	return c.Render(http.StatusOK, "snippet", "Hello")
}

func SnippetCreate(c echo.Context) error {
	return c.Render(http.StatusOK, "snippet_create", "Hello")
}
