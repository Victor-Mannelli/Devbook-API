package config

import (
	"api/src/utils"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	DBConnectionString = ""
	Port               = 0
)

func Load() {
	var err error

	err = godotenv.Load()
	utils.CheckError(err)

	Port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		Port = 5000
	}

	DBConnectionString = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
}
