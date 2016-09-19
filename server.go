package main

import (
	"github.com/SnippetsBucket/snicket/router"

	"github.com/labstack/echo/engine/standard"
)

func main() {
	router := router.Init()
	router.Run(standard.New(":8888"))
}
