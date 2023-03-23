package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type appConfig struct {
	DBHost     string
	DBPort     int
	DBName     string
	DBUserName string
	dbPassword string
}

func readConfig() *appConfig {
	err := godotenv.Load(".env")
	result := appConfig{}

	if res == nil {
		log.Fatal("error connecting to database")
	}

	return res
}

func readConfig() *DatabaseConfig {
	err := godotenv.Load(".env")
	res := DatabaseConfig{}
	if err != nil {
		log.Print("failed to load env files")
	}

	result.DBHost = os.Getenv("DBHost")
	result.DBName = os.Getenv("DBName")
	result.dbPassword = os.Getenv("DBPassword")
	result.DBPort, err = strconv.Atoi(os.Getenv("DBPort"))
	if err != nil {
		log.Println("Error from converting DBPort")
	}
	result.DBUserName = os.Getenv("DBUserName")
	return &result
}

func InitConfig() *appConfig {
	result := readConfig()
	if result == nil {
		return nil
	}
	return result

}
