package changedata

import (
	"encoding/json"
	"log"
	"myapp/ml"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ChangePlanMl(c echo.Context) error {
	input := map[string]interface{}{
		"feature1": 10.5,
		"feature2": "text",
		"values":   []int{1, 2, 3},
	}

	result, err := ml.MLWork(input)
	if err != nil {
		log.Fatalf("Ошибка: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "No response from ml"})
	}

	jsonResult, _ := json.MarshalIndent(result, "", "  ")

	return c.JSON(http.StatusOK, map[string]string{
		"plan": string(jsonResult),
	})
}
