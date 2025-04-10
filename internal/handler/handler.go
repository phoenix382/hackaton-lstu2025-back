package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Hello(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Hello from Hackathon Server!",
	})
}

// Для GET-запроса с параметрами
func AddNumbers(c echo.Context) error {
	a := c.QueryParam("a")
	b := c.QueryParam("b")

	// Конвертация в числа и сложение
	// Добавь обработку ошибок самостоятельно!
	numA, _ := strconv.Atoi(a)
	numB, _ := strconv.Atoi(b)

	return c.JSON(http.StatusOK, map[string]int{
		"result": numA + numB,
	})
}

// ИЛИ для POST с JSON телом
type NumbersRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

func AddNumbersJSON(c echo.Context) error {
	var req NumbersRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}

	return c.JSON(http.StatusOK, map[string]int{
		"result": req.A + req.B,
	})
}

func GetUserId(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]int{
		"result": c.Get("userID").(int),
	})
}