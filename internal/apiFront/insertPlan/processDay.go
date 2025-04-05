package insertPlan

import (
	"myapp/internal/db"

	"gorm.io/gorm"
)

func processDay(dbConn *gorm.DB, planID, dayNumber uint, data DayData) (*db.Day, error) {
	day := &db.Day{}
	err := dbConn.FirstOrInit(day,
		"plan_id = ? AND day_week = ?",
		planID,
		dayNumber,
	).Error

	if err != nil {
		return nil, err
	}

	updates := map[string]interface{}{
		"goal_exercise": data.Workouts.Type,
		"calories_all":  data.Nutrition.DailyCalories,
	}

	if day.ID == 0 { // Новый день
		day.PlanID = planID
		day.DayWeek = dayNumber
		day.GoalExercise = data.Workouts.Type
		day.CaloriesAll = data.Nutrition.DailyCalories
		err = dbConn.Create(day).Error
	} else { // Обновление существующего
		err = dbConn.Model(day).Updates(updates).Error
	}

	return day, err
}
