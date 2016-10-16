package handler

import (
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
)

type PageInfo struct {
	Title string
	Js    []string
}

func JSONHTTPErrorHandler(err error, c echo.Context) {
	code := fasthttp.StatusInternalServerError
	msg := "Internal Server Error"
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message
	}
	if !c.Response().Committed() {
		c.JSON(code, map[string]interface{}{
			"statusCode": code,
			"message":    msg,
		})
	}
}

func Home(c echo.Context) error {
	pageInfo := PageInfo{
		"Home",
		[]string{"create.bundle.js"},
	}
	return c.Render(fasthttp.StatusCreated, "home", pageInfo)
}

func Snippet(c echo.Context) error {
	pageInfo := PageInfo{
		"Snippet",
		[]string{"create.bundle.js"},
	}
	return c.Render(fasthttp.StatusCreated, "snippet", pageInfo)
}

func SnippetCreate(c echo.Context) error {
	pageInfo := PageInfo{
		"Create a new snippet",
		[]string{"create.bundle.js"},
	}
	return c.Render(fasthttp.StatusCreated, "snippet_create", pageInfo)
}
