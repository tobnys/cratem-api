package server

import (
	"github.com/gin-gonic/gin"
	"github.com/tobnys/cratem-api/controllers"
	"github.com/tobnys/cratem-api/helpers"
)

func Router() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("v1")
	{
		v1.GET("/", controllers.Index)
		auth := v1.Group("auth")
		{
			google := auth.Group("google")
			{
				google.GET("/login", controllers.Login)
				google.GET("/logout", controllers.Logout)
				google.GET("/callback", controllers.Callback)
			}
		}

		v1.GET("/cookie", func(c *gin.Context) {
			helpers.GenerateStateOauthCookie(c)
		})
	}

	return router
}
