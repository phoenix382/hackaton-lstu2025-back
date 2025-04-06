package planprocessing

import (
	"myapp/internal/db"

	"gorm.io/gorm"
)

func processExercises(dbConn *gorm.DB, dayID uint, workout WorkoutData) error {
	// Удаляем старые упражнения
	if err := dbConn.Where("day_id = ?", dayID).
		Delete(&db.Exercise{}).Error; err != nil {
		return err
	}

	// Создаем новые
	for _, ex := range workout.ExercisesList {
		exercise := db.Exercise{
			DayID: dayID,
			Name:  ex.Name,
			Info:  ex.Execution,
		}
		if err := dbConn.Create(&exercise).Error; err != nil {
			return err
		}
	}
	return nil
}
