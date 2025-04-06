package getdata

import (
	planprocessing "myapp/internal/apiFront/planJson"
	"myapp/internal/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PlanRequest struct {
    PlanID int `json:"planID"`
}

func GetPlanJson(c echo.Context) error {
	var req PlanRequest
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
    }
    
    planID := req.PlanID

	var plan db.PlanWeek
	if err := db.DB.Where("plan_id = ?", planID).First(&plan).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "not found plan in database. " + err.Error()})
	}

	planJson, err := planprocessing.BuildPlan(plan.ID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "failed buils plan. " + err.Error()})
	}

	return c.JSON(http.StatusOK, planJson)
}
