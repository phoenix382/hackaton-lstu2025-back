package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"

	"myapp/internal/db"
)

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Регистрация
func Register(c echo.Context) error {
	var req AuthRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request format"})
	}

	// Валидация email и пароля
	if req.Email == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "email and password are required"})
	}

	// Проверка формата email
	if !isValidEmail(req.Email) {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid email format"})
	}

	// Проверка сложности пароля
	if len(req.Password) < 8 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "password must be at least 8 characters long"})
	}

	// Хеширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to process password"})
	}

	// Сохранение в БД
	var userID int
	err = db.DB.QueryRow(
		"INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING id",
		req.Email,
		string(hashedPassword),
	).Scan(&userID) // Сохраняем ID в переменную

	if err != nil {
		// Проверяем является ли ошибка нарушением уникальности email
		if strings.Contains(err.Error(), "users_email_key") || strings.Contains(err.Error(), "unique constraint") {
			return c.JSON(http.StatusConflict, map[string]string{"error": "user with this email already exists"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "registration failed"})
	}

	// Генерация JWT токена
	tokenString, err := GenerateTokenJWT(userID)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid credentials"})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": tokenString,
	})
}

// Вспомогательная функция для проверки email
func isValidEmail(email string) bool {
	// Простая проверка формата email, можно заменить на более сложную
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}

// Вход
func Login(c echo.Context) error {
	var req AuthRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	var user struct {
		ID           int
		PasswordHash string
	}

	err := db.DB.QueryRow(
		"SELECT id, password_hash FROM users WHERE email = $1",
		req.Email,
	).Scan(&user.ID, &user.PasswordHash)

	if err == sql.ErrNoRows {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid credentials"})
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "database error"})
	}

	// Проверка пароля
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid credentials"})
	}

	// Генерация JWT токена
	tokenString, err := GenerateTokenJWT(user.ID)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid credentials"})
	}

	return c.JSON(http.StatusOK, map[string]string{
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
