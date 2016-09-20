package api

import (
	"fmt"

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
		title := c.FormValue("snippetTitle")
		content := c.FormValue("snippetContent")

		// s := new(model.Snippet)
		// c.Bind(&s)

		tx := c.Get("Tx").(*dbr.Tx)
		// snippet := model.CreateSnippet(s.SnippetTitle, s.SnippetText)
		snippet := model.CreateSnippet(title, content)

		if err := model.Save(tx, title, content); err != nil {
			fmt.Println("Save Error")
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusInternalServerError)
		}
		fmt.Println("Json Success")

		return c.JSON(fasthttp.StatusCreated, snippet)
	}
}
