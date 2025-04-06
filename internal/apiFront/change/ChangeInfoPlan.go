package changedata

import (
	"myapp/internal/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PlanRequestInfo struct {
	PlanID      int    `json:"planID"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

func ChangeInfoPlan(c echo.Context) error {
	userID := c.Get("userID").(int)

	var req PlanRequestInfo
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	return db.DB.Model(&db.PlanWeek{}).
		Where("id = ? AND user_id = ?", req.PlanID, userID).
		Updates(map[string]interface{}{
			"description": req.Description,
			"name":        req.Name,
		}).Error
}
