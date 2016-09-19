package api

import (
	"github.com/Sirupsen/logrus"
	"github.com/SnippetsBucket/snicket/model"
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
)

func TestApi() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(fasthttp.StatusCreated, "TestApi page")
	}
}

func Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		// title := c.FormValue("snippetTitle")
		// content := c.FormValue("snippetContent")

		s := new(model.Snippet)
		c.Bind(&s)

		tx := c.Get("Tx").(*dbr.Tx)

		snippet := model.CreateSnippet(s.SnippetTitle, s.SnippetText)

		if err := snippet.Save(tx); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusInternalServerError)
		}

		return c.JSON(fasthttp.StatusCreated, snippet)
	}
}
