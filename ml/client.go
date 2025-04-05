package ml

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

const (
	serverURL    = "https://back.root.sx/ml"
	callbackPort = ""
	timeout      = 30 * time.Second
)

// MLWork отправляет данные на ML сервер и возвращает результат
func MLWork(input map[string]interface{}) (map[string]interface{}, error) {
	// Генерация уникального токена
	token := generateToken()
	responses := make(map[string]map[string]interface{})
	// callbackURL := "http://localhost" + callbackPort
	callbackURL := "https://back.root.sx/ml"

	// Запуск сервера для приема ответа
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" {
				return
			}

			var data struct {
				Token  string
				Result map[string]interface{}
			}
			json.NewDecoder(r.Body).Decode(&data)

			if data.Token != "" {
				responses[data.Token] = data.Result
			}
			w.WriteHeader(200)
		})
		http.ListenAndServe(callbackPort, nil)
	}()

	// Отправка запроса
	request := map[string]interface{}{
		"token":        token,
		"payload":      input,
		"callback_url": callbackURL,
	}

	jsonData, _ := json.Marshal(request)
	resp, err := http.Post(serverURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("ошибка отправки: %v", err)
	}
	defer resp.Body.Close()

	// Ожидание ответа
	start := time.Now()
	for time.Since(start) < timeout {
		if result, exists := responses[token]; exists {
			return result, nil
		}
		time.Sleep(500 * time.Millisecond)
	}

	return nil, fmt.Errorf("таймаут ожидания")
}

// Вспомогательная функция для генерации токена
func generateToken() string {
	rand.Seed(time.Now().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 8)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}
