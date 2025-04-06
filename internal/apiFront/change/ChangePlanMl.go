package changedata

import (
	"net/http"

	planprocessing "myapp/internal/apiFront/planJson"
	"myapp/internal/db"

	"github.com/labstack/echo/v4"
)

func ChangePlanMl(c echo.Context) error {
	var req PlanRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	planID := uint(req.PlanID)
	//
	// planprocessing.BuildPlan(planID)
	//
	var input string = `{
		"Понедельник": {
		  "тренировки": {
			"тип тренировки": "Силовая (грудь, спина)",
			"список упражнений": [
			  {
				"название": "Отжимания",
				"информация о выполнении": "4 подхода по 15-20 повторений"
			  },
			  {
				"название": "Подтягивания/Тяга гантелей",
				"информация о выполнении": "4 подхода по 10-12 повторений"
			  },
			  {
				"название": "Планка",
				"информация о выполнении": "3 подхода по 60 секунд"
			  }
			]
		  },
		  "питание": {
			"суточная калорийность, БЖУ": "2200 ккал (Б: 160г, Ж: 70г, У: 220г)",
			"приемы пищи": [
			  {
				"прием": "завтрак",
				"блюдо": "Овсянка с яйцом и овощами",
				"калории и БЖУ": "550 ккал (Б: 25г, Ж: 15г, У: 70г)"
			  },
			  {
				"прием": "обед",
				"блюдо": "Гречка с куриной грудкой и салатом",
				"калории и БЖУ": "700 ккал (Б: 50г, Ж: 20г, У: 80г)"
			  },
			  {
				"прием": "ужин",
				"блюдо": "Творог с огурцом и зеленью",
				"калории и БЖУ": "450 ккал (Б: 40г, Ж: 15г, У: 30г)"
			  },
			  {
				"прием": "перекус",
				"блюдо": "Яблоко и горсть миндаля",
				"калории и БЖУ": "300 ккал (Б: 10г, Ж: 20г, У: 25г)"
			  }
			]
		  }
		},
		"Вторник": {
		  "тренировки": {
			"тип тренировки": "Кардио",
			"список упражнений": [
			  {
				"название": "Бег/Ходьба",
				"информация о выполнении": "40 минут в среднем темпе"
			  },
			  {
				"название": "Прыжки на скакалке",
				"информация о выполнении": "5 подходов по 2 минуты"
			  }
			]
		  },
		  "питание": {
			"суточная калорийность, БЖУ": "2200 ккал (Б: 160г, Ж: 70г, У: 220г)",
			"приемы пищи": [
			  {
				"прием": "завтрак",
				"блюдо": "Омлет с овощами и тост из цельнозернового хлеба",
				"калории и БЖУ": "500 ккал (Б: 30г, Ж: 20г, У: 50г)"
			  },
			  {
				"прием": "обед",
				"блюдо": "Рис с тушеной индейкой и капустой",
				"калории и БЖУ": "750 ккал (Б: 45г, Ж: 25г, У: 85г)"
			  },
			  {
				"прием": "ужин",
				"блюдо": "Запеченная рыба с овощами",
				"калории и БЖУ": "450 ккал (Б: 40г, Ж: 15г, У: 35г)"
			  },
			  {
				"прием": "перекус",
				"блюдо": "Йогурт натуральный с ягодами",
				"калории и БЖУ": "300 ккал (Б: 15г, Ж: 10г, У: 35г)"
			  }
			]
		  }
		},
		"Среда": {
		  "тренировки": {
			"тип тренировки": "Отдых",
			"список упражнений": []
		  },
		  "питание": {
			"суточная калорийность, БЖУ": "2200 ккал (Б: 160г, Ж: 70г, У: 220г)",
			"приемы пищи": [
			  {
				"прием": "завтрак",
				"блюдо": "Творожная запеканка с фруктами",
				"калории и БЖУ": "500 ккал (Б: 35г, Ж: 15г, У: 60г)"
			  },
			  {
				"прием": "обед",
				"блюдо": "Суп куриный с овощами и гречневый хлебец",
				"калории и БЖУ": "650 ккал (Б: 40г, Ж: 20г, У: 70г)"
			  },
			  {
				"прием": "ужин",
				"блюдо": "Салат из тунца, яиц и овощей",
				"калории и БЖУ": "500 ккал (Б: 45г, Ж: 25г, У: 30г)"
			  },
			  {
				"прием": "перекус",
				"блюдо": "Банан и арахисовая паста (1 ст.л.)",
				"калории и БЖУ": "250 ккал (Б: 8г, Ж: 12г, У: 30г)"
			  }
			]
		  }
		},
		"Четверг": {
		  "тренировки": {
			"тип тренировки": "Силовая (ноги, плечи)",
			"список упражнений": [
			  {
				"название": "Приседания",
				"информация о выполнении": "4 подхода по 20 повторений"
			  },
			  {
				"название": "Выпады",
				"информация о выполнении": "3 подхода по 12 повторений на ногу"
			  },
			  {
				"название": "Подъем на носки",
				"информация о выполнении": "4 подхода по 25 повторений"
			  }
			]
		  },
		  "питание": {
			"суточная калорийность, БЖУ": "2200 ккал (Б: 160г, Ж: 70г, У: 220г)",
			"приемы пищи": [
			  {
				"прием": "завтрак",
				"блюдо": "Гречневая каша с молоком и орехами",
				"калории и БЖУ": "600 ккал (Б: 25г, Ж: 20г, У: 80г)"
			  },
			  {
				"прием": "обед",
				"блюдо": "Запеченная курица с картофелем и морковью",
				"калории и БЖУ": "700 ккал (Б: 50г, Ж: 25г, У: 75г)"
			  },
			  {
				"прием": "ужин",
				"блюдо": "Омлет с творогом и зеленью",
				"калории и БЖУ": "450 ккал (Б: 40г, Ж: 25г, У: 20г)"
			  },
			  {
				"прием": "перекус",
				"блюдо": "Протеиновый коктейль на молоке",
				"калории и БЖУ": "300 ккал (Б: 30г, Ж: 10г, У: 25г)"
			  }
			]
		  }
		},
		"Пятница": {
		  "тренировки": {
			"тип тренировки": "Кардио + пресс",
			"список упражнений": [
			  {
				"название": "Велосипед/Эллипс",
				"информация о выполнении": "30 минут"
			  },
			  {
				"название": "Скручивания",
				"информация о выполнении": "4 подхода по 20 повторений"
			  },
			  {
				"название": "Подъем ног лежа",
				"информация о выполнении": "3 подхода по 15 повторений"
			  }
			]
		  },
		  "питание": {
			"суточная калорийность, БЖУ": "2200 ккал (Б: 160г, Ж: 70г, У: 220г)",
			"приемы пищи": [
			  {
				"прием": "завтрак",
				"блюдо": "Рисовая каша с изюмом и орехами",
				"калории и БЖУ": "550 ккал (Б: 20г, Ж: 15г, У: 80г)"
			  },
			  {
				"прием": "обед",
				"блюдо": "Индейка тушеная с булгуром и салатом",
				"калории и БЖУ": "750 ккал (Б: 55г, Ж: 25г, У: 75г)"
			  },
			  {
				"прием": "ужин",
				"блюдо": "Кефир с отрубями и ягодами",
				"калории и БЖУ": "400 ккал (Б: 25г, Ж: 10г, У: 50г)"
			  },
			  {
				"прием": "перекус",
				"блюдо": "Тост с авокадо и яйцом",
				"калории и БЖУ": "300 ккал (Б: 15г, Ж: 20г, У: 25г)"
			  }
			]
		  }
		},
		"Суббота": {
		  "тренировки": {
			"тип тренировки": "Силовая (руки, пресс)",
			"список упражнений": [
			  {
				"название": "Отжимания узким хватом",
				"информация о выполнении": "4 подхода по 15 повторений"
			  },
			  {
				"название": "Обратные отжимания от стула",
				"информация о выполнении": "3 подхода по 12 повторений"
			  },
			  {
				"название": "Планка с подъемом рук",
				"информация о выполнении": "3 подхода по 40 секунд"
			  }
			]
		  },
		  "питание": {
			"суточная калорийность, БЖУ": "2200 ккал (Б: 160г, Ж: 70г, У: 220г)",
			"приемы пищи": [
			  {
				"прием": "завтрак",
				"блюдо": "Сырники из творога с медом",
				"калории и БЖУ": "600 ккал (Б: 30г, Ж: 20г, У: 70г)"
			  },
			  {
				"прием": "обед",
				"блюдо": "Говядина тушеная с овощами и рисом",
				"калории и БЖУ": "800 ккал (Б: 50г, Ж: 30г, У: 80г)"
			  },
			  {
				"прием": "ужин",
				"блюдо": "Салат из свеклы с куриной грудкой",
				"калории и БЖУ": "450 ккал (Б: 35г, Ж: 15г, У: 40г)"
			  },
			  {
				"прием": "перекус",
				"блюдо": "Горький шоколад (30г) и кефир",
				"калории и БЖУ": "250 ккал (Б: 10г, Ж: 12г, У: 25г)"
			  }
			]
		  }
		},
		"Воскресенье": {
		  "тренировки": {
			"тип тренировки": "Отдых",
			"список упражнений": []
		  },
		  "питание": {
			"суточная калорийность, БЖУ": "2200 ккал (Б: 160г, Ж: 70г, У: 220г)",
			"приемы пищи": [
			  {
				"прием": "завтрак",
				"блюдо": "Овсянка с бананом и орехами",
				"калории и БЖУ": "550 ккал (Б: 20г, Ж: 15г, У: 80г)"
			  },
			  {
				"прием": "обед",
				"блюдо": "Запеченный лосось с картофелем и спаржей",
				"калории и БЖУ": "750 ккал (Б: 50г, Ж: 35г, У: 60г)"
			  },
			  {
				"прием": "ужин",
				"блюдо": "Творог с огурцом и зеленью",
				"калории и БЖУ": "450 ккал (Б: 40г, Ж: 15г, У: 30г)"
			  },
			  {
				"прием": "перекус",
				"блюдо": "Фруктовый салат (яблоко, груша, йогурт)",
				"калории и БЖУ": "300 ккал (Б: 10г, Ж: 5г, У: 60г)"
			  }
			]
		  }
		}
	}`

	// result, err := ml.MLWork(input)
	// if err != nil {
	// 	log.Fatalf("Ошибка: %v", err)
	// 	return c.JSON(http.StatusInternalServerError, map[string]string{"error": "No response from ml"})
	// }

	// jsonResult, _ := json.MarshalIndent(result, "", "  ")

	jsonResult := input

	planprocessing.ImportPlan(db.DB, planID, []byte(jsonResult))

	return c.JSON(http.StatusOK, map[string]string{
		"plan": string(jsonResult),
	})

}
