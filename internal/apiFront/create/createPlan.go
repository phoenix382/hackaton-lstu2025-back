package createplan

import (
	// planprocessing "myapp/internal/apiFront/planJson"
	"fmt"
	planprocessing "myapp/internal/apiFront/planJson"
	"myapp/internal/db"
	"myapp/ml"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreatePlan(c echo.Context) error {
	userID := c.Get("userID").(int)

	var user db.User
	if err := db.DB.First(&user, userID).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "database error" + err.Error()})
	}

	// Проверяем наличие активного плана
	var existingPlan db.PlanWeek
	err := db.DB.Where("user_id = ? AND current = ?", userID, true).First(&existingPlan).Error

	hasCurrent := true
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			hasCurrent = false
		} else {
			return c.JSON(http.StatusInternalServerError,
				map[string]string{"error": "database error: " + err.Error()})
		}
	}

	// Создаем новый план с инвертированным значением current
	newPlanWeek := db.PlanWeek{
		UserID:  uint(userID),
		Current: !hasCurrent, // Инвертируем статус
	}

	if err := db.DB.Create(&newPlanWeek).Error; err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"error": "plan creation failed: " + err.Error()})
	}

	planID := newPlanWeek.ID

	result, err := ml.MLWork(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "error ML SLOMALAS"})
	}

	err = planprocessing.ChangePlan(db.DB, planID, []byte(result))
	if err != nil {
		fmt.Println(result)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "database not save plan"})
	}

	return c.JSON(http.StatusOK, result)
}
