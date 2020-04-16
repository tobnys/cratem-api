package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func Login(c *gin.Context) {
	q := c.Request.URL.Query()
	q.Add("provider", "google")
	c.Request.URL.RawQuery = q.Encode()
	//gothic.BeginAuthHandler(c.Writer, c.Request)

	url, err := gothic.GetAuthURL(c.Writer, c.Request)
	if err != nil {
		fmt.Println("ERR", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"authUrl": url})
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
