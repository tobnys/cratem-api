package cfg

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	HOST string
	PORT string
)

func init() {
	// LOAD ENV
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set ENV after initialize
	HOST = os.Getenv("HOST")
	PORT = os.Getenv("PORT")
}
