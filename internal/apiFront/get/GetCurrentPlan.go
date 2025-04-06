package getdata

import (
	"myapp/internal/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetCurrentPlan(c echo.Context) error {
	userID := c.Get("userID").(int)

	var plan db.PlanWeek
	if err := db.DB.Where("user_id = ? and current = ?", userID, true).First(&plan).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "not found plan in database. " + err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]uint{"plan": plan.ID})
}
