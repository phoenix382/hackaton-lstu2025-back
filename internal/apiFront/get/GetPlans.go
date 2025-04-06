package getdata

import (
	"myapp/internal/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetPlans(c echo.Context) error {
	userID := c.Get("userID").(int)

	var plans []db.PlanWeek
	if err := db.DB.Where("user_id = ?", userID).Find(&plans).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "database error. " + err.Error()})
	}
	var result []uint
	for _, p := range plans {
		result = append(result, p.ID)
	}

	return c.JSON(http.StatusOK, map[string][]uint{"plans": result})
}
