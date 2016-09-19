package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func Home(c echo.Context) error {
	return c.Render(http.StatusOK, "home", "Hello")
}

func Snippet(c echo.Context) error {
	return c.Render(http.StatusOK, "snippet", "Hello")
}

func SnippetCreate(c echo.Context) error {
	return c.Render(http.StatusOK, "snippet_create", "Hello")
}
