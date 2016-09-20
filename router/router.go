package router

import (
	"html/template"
	"io"

	"github.com/SnippetsBucket/snicket/api"
	"github.com/SnippetsBucket/snicket/db"
	"github.com/SnippetsBucket/snicket/handler"
	sbMw "github.com/SnippetsBucket/snicket/middleware"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Init() *echo.Echo {
	t := &Template{
		// templates: template.Must(template.ParseGlob("public/views/*.html")),
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e := echo.New()
	e.Static("/assets", "public")

	e.Use(mw.Logger())
	e.Use(mw.Recover())
	e.Use(mw.Gzip())
	e.Use(mw.CORSWithConfig(mw.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	}))

	e.SetRenderer(t)

	// set custome middleware
	e.Use(sbMw.TransactionHandler(db.Init()))

	// view
	e.GET("/", handler.Home)
	e.GET("/snippet", handler.Snippet)
	e.GET("/create", handler.SnippetCreate)

	// api
	v1 := e.Group("/api/v1")
	{
		v1.POST("/snippet/create", api.Create())
		v1.GET("/test", api.TestApi())
	}
	return e
}
