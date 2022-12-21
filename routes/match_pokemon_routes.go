package routes

import (
	"farmacare/controllers"

	"github.com/gin-gonic/gin"
)

func MatchPokemonRoute(router *gin.Engine) {

	router.POST("/pokemon", controllers.CreateMatchPokemon())
	router.PUT("/pokemon", controllers.StartMatchPokemon())
}
