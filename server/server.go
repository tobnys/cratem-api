package server

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
)

func init() {
	// LOAD ENV VARIABLES
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// DECLARE PROVIDERS
	goth.UseProviders(
		//twitter.New(os.Getenv("TWITTER_KEY"), os.Getenv("TWITTER_SECRET"), "http://localhost:3000/auth/twitter/callback"),
		//facebook.New(os.Getenv("FACEBOOK_KEY"), os.Getenv("FACEBOOK_SECRET"), "http://localhost:3000/auth/facebook/callback"),
		google.New(os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("GOOGLE_CLIENT_SECRET"), "http://localhost:8080/v1/auth/google/callback"),
	)
}

func Initialize() {
	r := Router()
	r.Use(cors.Default())
	r.Run("localhost:8080")
}
