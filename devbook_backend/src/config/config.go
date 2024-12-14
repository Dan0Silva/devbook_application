package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	StringConnectionDB = ""
	Port               = ""
	SecretKey          = ""
)

func LoadEnvironment() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port = os.Getenv("API_PORT")

	StringConnectionDB = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_ADDRESS"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	SecretKey = os.Getenv("TOKEN_SECRET")

	fmt.Printf("  Environment loaded successfully\n")
}
