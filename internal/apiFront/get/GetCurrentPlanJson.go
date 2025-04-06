package getdata

import (
	planprocessing "myapp/internal/apiFront/planJson"
	"myapp/internal/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetCurrentPlanJson(c echo.Context) error {
	userID := c.Get("userID").(int)

	var plan db.PlanWeek
	if err := db.DB.Where("user_id = ? AND current = ?", userID, true).First(&plan).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "not found plan in database. " + err.Error()})
	}

	planJson, err := planprocessing.BuildPlan(plan.ID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "failed buils plan. " + err.Error()})
	}

	return c.JSON(http.StatusOK, planJson)
}
