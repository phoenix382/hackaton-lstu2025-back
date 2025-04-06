package db

import (
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() error {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return err
	}

	// Проверка соединения с БД
	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("database ping failed: %w", err)
	}

	DB.Migrator().DropTable(&User{}, &PlanWeek{}, &Day{}, &Diet{}, &Exercise{}) // ПОТОМ УДАЛИТЬ!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

	if err := DB.AutoMigrate(&User{}, &PlanWeek{}, &Day{}, &Diet{}, &Exercise{}); err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	user, err := AddUser("111", "1@1.1", "123123123", "1", "1")
	if err != nil {
		// Обработка ошибки
		log.Printf("Failed to add user: %v", err)
	} else {
		log.Printf("Added user with ID: %d", user.ID)
	}

	return nil
}

func AddUser(name, email, password, goal, gender string) (*User, error) {
	// Хеширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Создание пользователя
	user := User{
		Name:         name,
		Email:        email,
		PasswordHash: string(hashedPassword),
		Goal:         goal,
		Gender:       gender,
	}

	// Сохранение в базе данных
	result := DB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
