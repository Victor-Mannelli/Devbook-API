package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	DBConnectionString = ""
	Port               = 0
	JwtSecret          []byte
)

func Load() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		Port = 5000
	}

	DBConnectionString = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	JwtSecret = []byte(os.Getenv("JWT_SECRET"))
}
