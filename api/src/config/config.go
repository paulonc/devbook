package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DBConnectionString string
	Port               int
	SecretKey          []byte
)

func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		log.Printf("Invalid or missing API_PORT, defaulting to 9000: %v", err)
		Port = 9000
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	if dbUser == "" || dbPassword == "" || dbName == "" || dbHost == "" || dbPort == "" {
		log.Fatal("Database connection information is missing")
	}

	DBConnectionString = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
	if len(SecretKey) == 0 {
		log.Fatal("SECRET_KEY is missing or empty")
	}
}
