package cfg

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	GoogleOauthConfig *oauth2.Config
	OauthStateString  = "pseudo-random"
)

func init() {
	// LOAD ENV VARIABLES
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	GoogleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/v1/auth/google/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}