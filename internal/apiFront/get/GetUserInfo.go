package getdata

import (
	"myapp/internal/db"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func ptrUintToString(p *uint) string {
	if p == nil {
		return ""
	}
	return strconv.FormatUint(uint64(*p), 10)
}

func GetUserInfo(c echo.Context) error {
	userID := c.Get("userID").(int)

	var user db.User
	if err := db.DB.First(&user, userID).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "database error"})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"name":   user.Name,
		"gender": user.Gender,
		"email":  user.Email,
		"age":    ptrUintToString(user.Age),
		"height": ptrUintToString(user.Height),
		"weight": ptrUintToString(user.Weight),
		"goal":   user.Goal,
	})
}
