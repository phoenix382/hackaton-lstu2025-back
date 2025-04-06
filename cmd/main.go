package main

import (
	changedata "myapp/internal/apiFront/change"
	getdata "myapp/internal/apiFront/get"
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

	e := echo.New()
	e.Debug = true // Включаем режим отладки

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Public routes
	e.POST("/api/register", handler.Register)
	e.POST("/api/login", handler.Login)

	// Test
	e.POST("/api/add", handler.AddNumbersJSON) // POST с JSON телом
	e.GET("/api/hello", handler.Hello)
	e.GET("/api/add", handler.AddNumbers) // GET с параметрами ?a=5&b=3

	// Protected routes
	api := e.Group("/api")
	api.Use(handler.JWTMiddleware)
	{
		api.GET("/getId", handler.GetUserId)

		api.POST("/getUserInfo", getdata.GetUserInfo)
		api.POST("/getPlans", getdata.GetPlans)
		api.POST("/getCurrentPlan", getdata.GetCurrentPlan)
		api.POST("/getCurrentPlanJson", getdata.GetCurrentPlanJson)
		api.POST("/getPlanJson", getdata.GetPlanJson)

		api.POST("/changeUserInfo", changedata.ChangeUserInfo)
		api.POST("/ChangePlanMl", changedata.ChangePlanMl)
	}

	// Start server
	e.Logger.Fatal(e.Start(":4173"))
}
