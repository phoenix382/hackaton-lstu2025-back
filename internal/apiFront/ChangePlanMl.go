// package apifront

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"myapp/ml"
// )

// func ChangePlanMl() {
// 	input := map[string]interface{}{
// 		"feature1": 10.5,
// 		"feature2": "text",
// 		"values":   []int{1, 2, 3},
// 	}

// 	fmt.Println("Отправка данных на ML обработку...")
// 	result, err := ml.MLWork(input)
// 	if err != nil {
// 		log.Fatalf("Ошибка: %v", err)
// 	}

// 	// Красивый вывод результата
// 	jsonResult, _ := json.MarshalIndent(result, "", "  ")
// 	fmt.Println("Результат ML обработки:")
// 	fmt.Println(string(jsonResult))
// }
