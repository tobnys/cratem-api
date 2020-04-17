package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	"github.com/tobnys/cratem-api/cfg"
)

func Login(c *gin.Context) {
	url := cfg.GoogleOauthConfig.AuthCodeURL(cfg.OauthStateString)
	c.Redirect(http.StatusTemporaryRedirect, url)
	return
}

func Logout(c *gin.Context) {
	gothic.Logout(c.Writer, c.Request)
}

func Callback(c *gin.Context) {
	q := c.Request.URL.Query()
	q.Add("provider", "google")
	c.Request.URL.RawQuery = q.Encode()

	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		fmt.Println("ERROR CALLBACK", err)
		return
	}
	fmt.Println("USER", user)

	c.Redirect(http.StatusFound, "http://localhost:3000")
	return
}
