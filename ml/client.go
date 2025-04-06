package ml

import (
	"bytes"
	"encoding/json"
	"fmt"
	"myapp/internal/db"
	"net/http"
	"time"
)

const (
	serverURL1 = "https://back.root.sx/ml/get_first_plan"
	timeout    = 30 * time.Second
)

// MLWork отправляет данные на ML сервер и возвращает результат
func MLWork(input db.User) (map[string]interface{}, error) {
	// Подготовка запроса
	request := map[string]interface{}{
		"пол":     input.Gender,
		"возраст": input.Age,
		"вес":     input.Weight,
		"рост":    input.Height,
		"цель":    input.Goal,
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("ошибка кодирования данных: %v", err)
	}

	// Создаем HTTP клиент с таймаутом
	client := &http.Client{
		Timeout: timeout,
	}

	// Отправляем POST запрос
	resp, err := client.Post(serverURL1, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("ошибка отправки запроса: %v", err)
	}
	defer resp.Body.Close()

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.StatusCode)
		return nil, fmt.Errorf("сервер вернул ошибку: %s", resp.Status)
	}

	// Декодируем JSON ответ
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println(result)
		fmt.Println(err)
		return nil, fmt.Errorf("ошибка декодирования ответа: %v", err)
	}

	return result, nil
}
