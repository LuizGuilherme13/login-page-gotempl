package initializers

import (
	"log"
	"os"

	"github.com/LuizGuilherme13/unodatabase/pkg/unodatabase/models"
	"github.com/joho/godotenv"
)

var DB models.DBConn

func LoadDB() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Erro ao ler o arquivo .env")
	}

	DB = models.DBConn{
		Driver:   os.Getenv("DB_DRIVER"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_DBNAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}
}
