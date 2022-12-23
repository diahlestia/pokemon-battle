package routes

import (
	"farmacare/controllers"

	"github.com/gin-gonic/gin"
)

func MatchPokemonRoute(router *gin.Engine) {

	router.GET("/pokemon/count", controllers.GetAllPokemonsFromApi())
	router.POST("/pokemon/match/create", controllers.CreateMatchPokemon())
	router.PUT("/pokemon/match/start", controllers.StartMatchPokemon())
	router.GET("/pokemon/match", controllers.GetMatchPokemons())
	router.GET("/pokemon/matches", controllers.GetAllPokemons())
	router.PUT("/pokemon/discualification", controllers.MatchDiscualification())
}
