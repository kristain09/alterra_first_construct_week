//file untuk konfigurasi database secara umum bagi aplikasi.

package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct{
	Host			string
	Port			int
	Username 	string
	Password 	string
	Name 			string
}

func InitDatabase() *DatabaseConfig {
	res := readConfig()

	if res == nil {
		log.Fatal("error connecting to database")
	}

	return res
}

func readConfig() *DatabaseConfig {
	err := godotenv.Load(".env")
	res := DatabaseConfig{}
	if err != nil {
		log.Println("Tidak bisa baca konfigurasi")
		return nil
	}
	res.Username = os.Getenv("Username")
	res.Password = os.Getenv("Password")
	res.Host = os.Getenv("Host")
	cnv, err := strconv.Atoi(os.Getenv("Port"))
	if err != nil {
		log.Println("Nilai port tidak valid")
		return nil
	}
	res.Port = cnv
	res.Name = os.Getenv("Name")

	return &res
}