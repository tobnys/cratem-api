package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/tobnys/cratem-api/cfg"
	"golang.org/x/oauth2"
)

type GoogleUserReturn struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
}

func ValidateToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		fmt.Println("ERROR PARSING JWT", err)
	}

	fmt.Println("TOKEN", token)

	/*
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println(claims["authorize"], claims["ID"], claims["exp"])
		} else {
			fmt.Println(err)
		}
	*/

	return token.Valid
}

func GetUserToken(state string, code string) (*oauth2.Token, error) {
	if state != cfg.OauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}
	token, err := cfg.GoogleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}
	return token, nil
}

func GetUserInfo(state string, code string) (GoogleUserReturn, error) {
	var googleUserReturn GoogleUserReturn
	if state != cfg.OauthStateString {
		return googleUserReturn, fmt.Errorf("invalid oauth state")
	}
	token, err := cfg.GoogleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return googleUserReturn, fmt.Errorf("code exchange failed: %w", err)
	}
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return googleUserReturn, fmt.Errorf("failed getting user info: %w", err)
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&googleUserReturn)
	if err != nil {
		return googleUserReturn, fmt.Errorf("failed decoding: %w", err)
	}

	return googleUserReturn, nil
}

func GenerateStateOauthCookie(c *gin.Context, user GoogleUserReturn) {
	cookie, err := c.Cookie("Bearer")
	if err != nil {
		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)

		claims["authorize"] = true
		claims["ID"] = user.ID
		claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

		tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
		if err != nil {
			fmt.Errorf("failed JWT: %w", err)
		}

		// New
		cookie = "NotSet"
		c.SetCookie("Bearer", tokenString, 3600, "/", cfg.HOST, false, true)
	}

	fmt.Printf("Cookie value: %s \n", cookie)
}
