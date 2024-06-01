package main

import (
	"os"

	"github.com/Kokosalah45/share-my-notes/internal/database"
)

func main() {
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_HOST")
	dbPassword := os.Getenv("DB_PASSWORD")

	config := database.NewDBConfig(dbUser, dbName, dbHost, dbPassword)
	x, r := database.PostgresDB.InitDB(config)
}    