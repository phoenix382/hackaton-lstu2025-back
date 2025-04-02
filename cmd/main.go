package main

import (
	"myapp/internal/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/hello", handler.Hello)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
