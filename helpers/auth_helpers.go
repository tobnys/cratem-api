package helpers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tobnys/cratem-api/cfg"
	"golang.org/x/oauth2"
)

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

func GetUserInfo(state string, code string) ([]byte, error) {
	if state != cfg.OauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}
	token, err := cfg.GoogleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}
	return contents, nil
}

func GenerateStateOauthCookie(c *gin.Context) {
	cookie, err := c.Cookie("gin_cookie")

	if err != nil {
		cookie = "NotSet"
		c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
	}

	fmt.Printf("Cookie value: %s \n", cookie)
}
