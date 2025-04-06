package ml

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"myapp/internal/db"
	"net/http"
	"time"
)

const (
	serverURL1 = "https://back.root.sx/ml/get_first_plan"
	timeout    = 30 * time.Second
)

// MLWork отправляет данные на ML сервер и возвращает результат
func MLWork(input db.User) (string, error) {
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
		return "", fmt.Errorf("ошибка кодирования данных: %v", err)
	}

	// Создаем HTTP клиент с таймаутом
	client := &http.Client{
		Timeout: timeout,
	}

	// Отправляем POST запрос
	resp, err := client.Post(serverURL1, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("ошибка отправки запроса: %v", err)
	}
	defer resp.Body.Close()

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("сервер вернул ошибку: %s", resp.Status)
	}

	// Читаем тело ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("ошибка чтения тела ответа: %v", err)
	}

	// Закрываем тело ответа
	defer resp.Body.Close()
	ans := string(body)

	// Конвертируем в строку и возвращаем
	return ans[1 : len(ans)-1], nil
}
