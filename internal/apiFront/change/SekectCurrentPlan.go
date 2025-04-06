package changedata

import (
	"myapp/internal/db"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// SelectCurrentPlan устанавливает активный план
func SelectCurrentPlan(c echo.Context) error {
	userID := c.Get("userID").(int)

	var req PlanRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	return db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&db.PlanWeek{}).
			Where("user_id = ?", userID).
			Update("current", false).Error; err != nil {
			return err
		}

		return tx.Model(&db.PlanWeek{}).
			Where("id = ? AND user_id = ?", req.PlanID, userID).
			Update("current", true).Error
	})
}
