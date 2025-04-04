package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

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
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		return fmt.Errorf("database ping failed: %v", err)
	}

	// Применяем миграции
	if err := runMigrations(); err != nil {
		return err
	}

	return nil
}

func runMigrations() error {
	// Реализация применения миграций (можно использовать библиотеку migrate)
	// Для простоты выполним SQL из файла
	sqlFile, err := os.ReadFile("./migrations/init.sql")
	if err != nil {
		return err
	}

	_, err = DB.Exec(string(sqlFile))
	return err
}
