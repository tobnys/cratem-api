package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tobnys/cratem-api/cfg"
	"github.com/tobnys/cratem-api/helpers"
)

func Login(c *gin.Context) {
	//helpers.GenerateStateOauthCookie(c)
	//fmt.Println("THIS COOKIE", oauthState)
	url := cfg.GoogleOauthConfig.AuthCodeURL("pseudo-random")
	c.JSON(200, gin.H{"authUrl": url})
	return
}

func Logout(c *gin.Context) {
	//gothic.Logout(c.Writer, c.Request)
}

func Callback(c *gin.Context) {
	state := c.Request.FormValue("state")

	if state != "pseudo-random" {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", "pseudo-random", state)
		http.Redirect(c.Writer, c.Request, "/", http.StatusTemporaryRedirect)
		return
	}

	fmt.Println("STATE PARAM", c.Request.FormValue("state"))

	token, err := helpers.GetUserToken(c.Request.FormValue("state"), c.Request.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(c.Writer, c.Request, "/", http.StatusTemporaryRedirect)
		return
	}

	fmt.Println("TOKENNNNNNNNNNNNNNNNNNN", token)

	// Create user in DB
	// Redirect user with token here
	http.Redirect(c.Writer, c.Request, "http://localhost:3000", http.StatusPermanentRedirect)
	//fmt.Fprintf(c.Writer, "Content: %s\n", content)
	return
}
