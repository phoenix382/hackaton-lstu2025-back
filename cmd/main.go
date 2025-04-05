package main

import (
	"myapp/internal/db"
	"myapp/internal/handler"

	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Загрузка переменных окружения (можно использовать godotenv)
	// Пример для теста:
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := db.Init(); err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}
	defer db.DB.Close()

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Public routes
	e.POST("/api/register", handler.Register)
	e.POST("/api/login", handler.Login)

	e.GET("/api/hello", handler.Hello)
	e.GET("/api/add", handler.AddNumbers)      // GET с параметрами ?a=5&b=3
	e.POST("/api/add", handler.AddNumbersJSON) // POST с JSON телом

	// Protected routes
	api := e.Group("/api")
	api.Use(handler.JWTMiddleware)
	{
		api.GET("/projects", handler.GetProjects)
		// api.POST("/projects", handler.CreateProject)
	}

	// Start server
	e.Logger.Fatal(e.Start(":4173"))
}
