package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	// TODO: randomize it
	OauthStateString = "pseudo-random"
)

func Login(c *gin.Context) {
	urlCode := GoogleOauthConfig.AuthCodeURL(oauthStateString)
	c.Redirect(urlCode, "http://google.com")
	return
}

func Callback(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "User founded!"})
	return
}