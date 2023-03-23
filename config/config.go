package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Name     string
}

func InitConfig() *DatabaseConfig {
	res := readConfig()

	if res == nil {
		log.Fatal("error connecting to database")
	}

	return res
}

func readConfig() *DatabaseConfig {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Tidak bisa baca konfigurasi")
		return nil
	}

	port, err := strconv.Atoi(os.Getenv("Port"))
	if err != nil {
		fmt.Println("Nilai port tidak valid")
		return nil
	}

	return &DatabaseConfig{
		Host:     os.Getenv("Host"),
		Port:     port,
		Username: os.Getenv("Username"),
		Password: os.Getenv("Password"),
		Name:     os.Getenv("Name"),
	}
}
