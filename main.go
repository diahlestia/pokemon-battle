package main

import (
	"farmacare/configs"
	routes "farmacare/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// database connect
	configs.Connect()

	// routes
	routes.MatchRoute(router)
	routes.MatchPokemonRoute(router)

	router.Run()
}
