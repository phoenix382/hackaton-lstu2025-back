package db

import (
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
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

	DB.Migrator().DropTable(&User{}) // ПОТОМ УДАЛИТЬ!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

	if err := DB.AutoMigrate(&User{}); err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	if DB.Migrator().HasTable(&User{}) {
		log.Println("Table exists!")
	} else {
		log.Println("Table creation failed!")
	}

	return nil
}

// func runMigrations() error {
// 	// Реализация применения миграций (можно использовать библиотеку migrate)
// 	// Для простоты выполним SQL из файла
// 	sqlFile, err := os.ReadFile("./migrations/init.sql")
// 	if err != nil {
// 		return err
// 	}

// 	_, err = DB.Exec(string(sqlFile))
// 	return err
// }
