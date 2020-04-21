package controllers

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tobnys/cratem-api/cfg"
	"github.com/tobnys/cratem-api/helpers"
)

func AuthValidate(c *gin.Context) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	bodyString := buf.String()

	validToken := helpers.ValidateToken(bodyString)
	if validToken {
		c.JSON(200, gin.H{})
		return
	} else {
		c.JSON(404, gin.H{})
	}
}

func Login(c *gin.Context) {
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

	user, err := helpers.GetUserInfo(c.Request.FormValue("state"), c.Request.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(c.Writer, c.Request, "/", http.StatusTemporaryRedirect)
		return
	}

	fmt.Printf("%+v \n", user)

	fmt.Println("TEST", user.ID)

	// Create user in DB IF NOT EXIST (CHECK user.ID)
	// Connect OAUTHID (user.ID) to new user table

	// Create cookie for user
	helpers.GenerateStateOauthCookie(c, user)

	// Redirect user with token here
	http.Redirect(c.Writer, c.Request, "http://localhost:3000/main", http.StatusPermanentRedirect)
	return
}
