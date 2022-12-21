package routes

import (
	"farmacare/controllers"

	"github.com/gin-gonic/gin"
)

func MatchRoute(router *gin.Engine) {

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
		})
	})
	router.POST("/match", controllers.Create())
	router.GET("/match", controllers.GetMatchPokemon())
}
