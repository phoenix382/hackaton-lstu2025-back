package changedata

import (
	"myapp/internal/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PlanRequest struct {
	PlanID int `json:"planID"`
}

func CopyPlan(c echo.Context) error {
	userID := c.Get("userID").(int)
	var req PlanRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	
	// Начинаем транзакцию
	tx := db.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Создаем новый план недели для пользователя
	newPlan := db.PlanWeek{
		UserID: uint(userID),
	}
	if err := tx.Create(&newPlan).Error; err != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create plan: " + err.Error()})
	}

	// Получаем дни из исходного плана
	var days []db.Day
	if err := tx.Where("plan_id = ?", req.PlanID).Find(&days).Error; err != nil {
		tx.Rollback()
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "database error: " + err.Error()})
	}

	dayMap := make(map[uint]uint) // для отображения старых dayID -> новые dayID

	// Копируем дни
	for _, day := range days {
		newDay := db.Day{
			PlanID:   newPlan.ID,
			DayWeek:  day.DayWeek,
			Goal:     day.Goal,
			Calories: day.Calories,
		}
		if err := tx.Create(&newDay).Error; err != nil {
			tx.Rollback()
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create day: " + err.Error()})
		}
		dayMap[day.ID] = newDay.ID
	}

	// Копируем диеты для каждого дня
	for _, day := range days {
		var diets []db.Diet
		if err := tx.Where("day_id = ?", day.ID).Find(&diets).Error; err != nil {
			tx.Rollback()
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to get diets: " + err.Error()})
		}

		for _, diet := range diets {
			newDiet := db.Diet{
				DayID:     dayMap[day.ID],
				MealType:  diet.MealType,
				Name:      diet.Name,
				Structure: diet.Structure,
				Colories:  diet.Colories,
			}
			if err := tx.Create(&newDiet).Error; err != nil {
				tx.Rollback()
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create diet: " + err.Error()})
			}
		}
	}

	// Копируем упражнения для каждого дня
	for _, day := range days {
		var exercises []db.Exercise
		if err := tx.Where("day_id = ?", day.ID).Find(&exercises).Error; err != nil {
			tx.Rollback()
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to get exercises: " + err.Error()})
		}

		for _, exercise := range exercises {
			newExercise := db.Exercise{
				DayID: dayMap[day.ID],
				Name:  exercise.Name,
				Info:  exercise.Info,
				Done:  false, // сбрасываем статус выполнения
			}
			if err := tx.Create(&newExercise).Error; err != nil {
				tx.Rollback()
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create exercise: " + err.Error()})
			}
		}
	}

	// Фиксируем транзакцию
	if err := tx.Commit().Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "transaction commit failed: " + err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]uint{"plan_id": newPlan.ID})
}
