package insertPlan

import (
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

func ImportPlan(db *gorm.DB, planWeekID uint, jsonData []byte) error {
	var schedule WeekSchedule
	if err := json.Unmarshal(jsonData, &schedule); err != nil {
		return fmt.Errorf("ошибка парсинга JSON: %v", err)
	}

	dayMapping := map[string]uint{
		"Понедельник": 1, "Вторник": 2, "Среда": 3,
		"Четверг": 4, "Пятница": 5, "Суббота": 6,
		"Воскресенье": 7,
	}

	for russianDay, data := range schedule {
		dayNumber, ok := dayMapping[russianDay]
		if !ok {
			return fmt.Errorf("неизвестный день: %s", russianDay)
		}

		// Обработка дня
		day, err := processDay(db, planWeekID, dayNumber, data)
		if err != nil {
			return err
		}

		// Обработка упражнений
		if err := processExercises(db, day.ID, data.Workouts); err != nil {
			return err
		}

		// Обработка питания
		if err := processDiet(db, day.ID, data.Nutrition); err != nil {
			return err
		}
	}

	return nil
}
