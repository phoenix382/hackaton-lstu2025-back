package handler

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"myapp/internal/db"
)

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegistrationRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Goal     string `json:"goal"`
	Gender   string `json:"gender"`
}

// Регистрация
func Register(c echo.Context) error {
	var req RegistrationRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request format"})
	}

	if errStr := isValidAuthRequest(req); errStr != "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": errStr})
	}

	// Хеширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to process hashed password"})
	}

	// Создание пользователя через GORM
	user := db.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		Goal:         req.Goal,
		Gender:       req.Gender,
	}

	result := db.DB.Create(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return c.JSON(http.StatusConflict, map[string]string{"error": "user already exists"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "registration failed"})
	}

	tokenString, err := GenerateTokenJWT(int(user.ID))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid credentials"})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": tokenString,
	})
}

func isValidAuthRequest(req RegistrationRequest) string {
	if !isValidEmail(req.Email) {
		return "invalid email format"
	}

	if !isValidPassword(req.Password) {
		return "invalid password format"
	}

	if !isValidName(req.Name) {
		return "invalid name format"
	}

	if !isValidGoal(req.Goal) {
		return "invalid goal format"
	}

	if !isValidGender(req.Gender) {
		return "invalid gender format"
	}

	return ""
}

func isValidGoal(goal string) bool {
	return goal != ""
}

func isValidGender(gender string) bool {
	return gender != ""
}

func isValidName(name string) bool {
	return name != ""
}

func isValidEmail(email string) bool {
	return strings.Contains(email, "@") && strings.Contains(email, ".") && len(email) > 4
}

func isValidPassword(password string) bool {
	return len(password) > 8
}

func Login(c echo.Context) error {
	var req AuthRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	var user db.User
	result := db.DB.Where("email = ?", req.Email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid credentials"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "database error"})
	}

	// Проверка пароля
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid credentials"})
	}

	tokenString, err := GenerateTokenJWT(int(user.ID))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid credentials"})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"userID": user.ID,
		"token": tokenString,
	})
}

func GenerateTokenJWT(userID int) (string, error) {
	// Генерация JWT токена
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": strconv.Itoa(userID),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "failed to generate token", err
	}

	return tokenString, nil
}

// Middleware для проверки JWT
func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "missing token"})
		}

		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID, _ := strconv.Atoi(claims["sub"].(string))
			c.Set("userID", userID)
			return next(c)
		}

		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token"})
	}
}
