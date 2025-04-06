package ml

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"myapp/internal/db"
	"net/http"
	"time"
)

const (
	serverURL1   = "https://back.root.sx/ml/get_first_plan"
	serverURL2   = "https://back.root.sx/ml/update_plan"
	callbackPort = ":8080" // Добавьте нужный порт для callback
	timeout      = 30 * time.Second
)

// MLWork отправляет данные на ML сервер и возвращает результат
func MLWork(input db.User) (map[string]interface{}, error) {
	// Создаем канал для передачи ответа
	responseChan := make(chan map[string]interface{}, 1)

	// Генерируем уникальный токен для этого запроса
	token := generateToken()

	// Запуск сервера для приема ответа
	go func() {
		http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}

			var data struct {
				Token  string                 `json:"token"`
				Result map[string]interface{} `json:"result"`
			}

			if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			if data.Token == token {
				responseChan <- data.Result
			}
			w.WriteHeader(http.StatusOK)
		})

		if err := http.ListenAndServe(callbackPort, nil); err != nil {
			fmt.Printf("Ошибка запуска сервера: %v\n", err)
		}
	}()

	// Подготовка и отправка запроса
	request := map[string]interface{}{
		"пол":          input.Gender,
		"возраст":      input.Age,
		"вес":          input.Weight,
		"рост":         input.Height,
		"цель":         input.Goal,
		"token":        token,                                                     // Добавляем токен в запрос
		"callback_url": "http://your-server-address" + callbackPort + "/callback", // Укажите ваш адрес
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("ошибка кодирования данных: %v", err)
	}

	// Выбираем нужный URL (пример - используем serverURL1)
	serverURL := serverURL1
	resp, err := http.Post(serverURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("ошибка отправки: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("сервер вернул статус: %d", resp.StatusCode)
	}

	// Ожидание ответа через callback
	select {
	case result := <-responseChan:
		return result, nil
	case <-time.After(timeout):
		return nil, fmt.Errorf("таймаут ожидания ответа")
	}
}

// Вспомогательная функция для генерации токена
func generateToken() string {
	rand.Seed(time.Now().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 16) // Увеличим длину токена для надежности
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}
