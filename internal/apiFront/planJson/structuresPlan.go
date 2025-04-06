package planprocessing

type (
	WeekSchedule map[string]DayData

	DayData struct {
		Workouts  WorkoutData   `json:"тренировки"`
		Nutrition NutritionData `json:"питание"`
	}

	WorkoutData struct {
		Type          string          `json:"тип тренировки"`
		ExercisesList []ExerciseInput `json:"список упражнений"`
	}

	ExerciseInput struct {
		Name      string `json:"название"`
		Execution string `json:"информация о выполнении"`
	}

	NutritionData struct {
		DailyCalories string      `json:"суточная калорийность. БЖУ"`
		Meals         []MealInput `json:"приемы пищи"`
	}

	MealInput struct {
		MealType string `json:"прием"`
		Dish     string `json:"блюдо"`
		Calories string `json:"калории и БЖУ"`
	}
)
