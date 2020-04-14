package server

import (
	"github.com/tobnys/cratem-api/controllers"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("v1")
	{
		v1.GET("/", controllers.Index)
		
		auth := router.Group("auth")
		{
			auth.GET("/login", controllers.Login)
			auth.GET("/callback", controllers.Callback)
		}
	}

	return router
}