package main

import (
	"farmacare/configs"
	"farmacare/models"
	"fmt"
)

func init() {

	configs.Connect()

}

func main() {
	configs.DB.AutoMigrate(&models.Match{}, &models.MatchPokemon{})
	fmt.Println("? Migration complete")
}
