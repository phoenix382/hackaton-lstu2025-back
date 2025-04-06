package planprocessing

import (
	"myapp/internal/db"

	"gorm.io/gorm"
)

func processDiet(dbConn *gorm.DB, dayID uint, nutrition NutritionData) error {
	// Удаляем старые записи питания
	if err := dbConn.Where("day_id = ?", dayID).
		Delete(&db.Diet{}).Error; err != nil {
		return err
	}

	// Создаем новые
	for _, meal := range nutrition.Meals {
		diet := db.Diet{
			DayID:    dayID,
			MealType: meal.MealType,
			Name:     meal.Dish,
			Colories: meal.Calories,
		}
		if err := dbConn.Create(&diet).Error; err != nil {
			return err
		}
	}
	return nil
}
