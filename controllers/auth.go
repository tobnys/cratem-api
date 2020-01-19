package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Send(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "User founded!"})
	return
}