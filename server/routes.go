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
		
		sms := router.Group("sms")
		{
			sms.POST("/send/")
		}
	}

	return router
}