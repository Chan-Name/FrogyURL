package storage

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

func configNew() (string, error) {
	err := godotenv.Load("C:/GoChan/.env")
	if err != nil {
		fmt.Print(err)
		return "", err
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	str := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		dbUser, dbPassword, dbName)

	slog.Info("Config create")
	return str, err
}
