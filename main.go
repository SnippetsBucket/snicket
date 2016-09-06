package main

import (
	"github.com/labstack/echo/engine/standard"

	"snicket/router"
)

func main() {
	router := router.Init()
	router.Run(standard.New(":8888"))
}
