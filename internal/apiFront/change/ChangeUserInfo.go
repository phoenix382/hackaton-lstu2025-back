package changedata

import (
	"myapp/internal/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ChangeUserInfo(c echo.Context) error {
	userID := c.Get("userID").(int)

	var req db.User
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input json format"})
	}

	// Создаем карту для обновления, исключая чувствительные поля
	updates := map[string]interface{}{
		"name":   req.Name,
		"gender": req.Gender,
		"email":  req.Email,
		"age":    req.Age,
		"height": req.Height,
		"weight": req.Weight,
		"goal":   req.Goal,
	}

	return db.DB.Model(&db.User{}).Where("id = ?", userID).Updates(updates).Error
}
