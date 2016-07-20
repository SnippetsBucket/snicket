package main

import (
    "snicket/handler"

    "github.com/labstack/echo"
    "github.com/labstack/echo/engine/standard"
    "github.com/labstack/echo/middleware"
)

func main() {
    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.GET("/", handler.HomePage())
    e.Run(standard.New(":1323"))
}
