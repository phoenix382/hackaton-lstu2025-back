package planprocessing

type (
	WeekSchedule map[string]DayData

	DayData struct {
		Workouts  WorkoutData   `json:"тренировки"`
		Nutrition NutritionData `json:"питание"`
	}

	WorkoutData struct {
		Type          string          `json:"тип_тренировки"`
		ExercisesList []ExerciseInput `json:"упражнения"`
	}

	ExerciseInput struct {
		Name      string `json:"название"`
		Execution string `json:"информация_о_выполнении"`
	}

	NutritionData struct {
		DailyCalories string      `json:"суточная_калорийность"`
		Meals         []MealInput `json:"приемы_пищи"`
	}

	MealInput struct {
		MealType string `json:"прием"`
		Dish     string `json:"блюдо"`
		Calories string `json:"калории_и_БЖУ"`
	}
)
