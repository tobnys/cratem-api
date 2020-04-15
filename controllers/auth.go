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
	fmt.Println("ENCODING", q.Encode())
	gothic.BeginAuthHandler(c.Writer, c.Request)
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
		fmt.Println("ERROR", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": user})
	return
}
