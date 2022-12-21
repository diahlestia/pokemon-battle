package controllers

import (
	"encoding/json"
	"farmacare/configs"
	"farmacare/models"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type iMatchPokemonRepositorySave interface {
	Save(*models.MatchPokemon) (*models.MatchPokemon, error)
}

type matchPokemonServiceCreate struct {
	iMatchPokemonRepositorySave
}

type NewMatchPokemon struct {
	MatchId uint64 `json:"matchId"`
}

type UpdateMatchPokemon struct {
	MatchId  uint64 `json:"matchId"`
	Position uint64 `json:"position"`
}

type species struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type ResponsePokeApi struct {
	Count    uint64        `json:"count"`
	Next     string        `json:"next"`
	Previous string        `json:"previous"`
	Results  []interface{} `json:"results"`
}

type stats struct {
	BaseStat uint64 `json:"base_stat"`
}

type ResponsePokeApiDetail struct {
	Abilities              []interface{} `json:"abilities"`
	BaseExperience         uint64        `json:"base_experience"`
	Forms                  []interface{} `json:"forms"`
	GameIndices            []interface{} `json:"game_indices"`
	Height                 uint64        `json:"height"`
	HeldItems              []interface{} `json:"held_items"`
	Id                     uint64        `json:"id"`
	IsDefault              bool          `json:"is_default"`
	LocationAreaEncounters string        `json:"location_area_encounters"`
	Moves                  []interface{} `json:"moves"`
	Name                   string        `json:"name"`
	Order                  uint64        `json:"order"`
	PastYypes              []interface{} `json:"past_types"`
	Species                species       `json:"species"`
	Stats                  []stats       `json:"stats"`
	Types                  []interface{} `json:"types"`
	Weight                 uint64        `json:"weight"`
}

type GroceryUpdate struct {
	Position uint64 `json:"position"`
}

func CreateMatchPokemon() gin.HandlerFunc {
	return func(c *gin.Context) {

		var match NewMatchPokemon

		pokemonIds := getPokemonNumbers()

		if err := c.ShouldBindJSON(&match); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "Data": err})
			return
		}

		for i := 0; i < len(pokemonIds); i++ {

			stats := getStatsPokemonId(pokemonIds[i])

			newMatch := models.MatchPokemon{
				MatchId:   match.MatchId,
				PokemonId: pokemonIds[i],
				Score:     stats,
			}

			database := configs.Connect()
			matchPokemonRepository := models.NewMatchPokemon(database)
			matchPokemonService := matchPokemonServiceCreate{matchPokemonRepository}
			res, err := matchPokemonService.Save(&newMatch)

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "Data": err})

				return
			}

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "Data": err})

			} else {
				c.JSON(http.StatusCreated, gin.H{"message": "success", "Data": res})
			}
		}

	}
}

func StartMatchPokemon() gin.HandlerFunc {
	return func(c *gin.Context) {

		var updateMatch UpdateMatchPokemon
		var matchPokemon models.MatchPokemon

		if err := c.ShouldBindJSON(&updateMatch); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "Data": err})
			return
		}

		database := configs.Connect()
		matchPokemonRepository := models.NewMatchPokemon(database)

		results, err := matchPokemonRepository.GetPokemonByMatchId(int(updateMatch.MatchId))

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "Data": err})
			return
		}

		var position uint64 = 0

		updatedMatchPokemon := models.MatchPokemon{}

		for i, element := range results.MatchPokemons {
			position = uint64(i + 1)
			updatedMatchPokemon = models.MatchPokemon{
				Position: uint64(position),
			}

			if err := database.Model(&matchPokemon).Where("id = ?", element.ID).Updates(updatedMatchPokemon).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

		}

		c.JSON(http.StatusOK, matchPokemon)

	}
}

func getPokemonNumbers() []uint64 {

	min := 1
	max := getTotalPokemon()

	var pokemonIds []uint64

	// generate 5 random numbers
	for i := 0; i < max; i++ {

		generateRandom := uint64(rand.Intn(max-min) + min)
		isValid := validationPokemonId(generateRandom)

		if isValid && len(pokemonIds) <= 5 {
			pokemonIds = append(pokemonIds, generateRandom)
			if len(pokemonIds) == 5 {
				break
			}
		}
	}
	return pokemonIds

}

func getTotalPokemon() int {

	var responseAPI ResponsePokeApi

	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon?limit=10000")

	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err)
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&responseAPI); err != nil {
		fmt.Print(err)
	}

	return int(responseAPI.Count)

}

func validationPokemonId(id uint64) bool {
	var responseAPI ResponsePokeApiDetail

	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d", id)

	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return false
	}

	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&responseAPI); err != nil {
		fmt.Print(err)
		return false
	} else {
		return true
	}
}

func getStatsPokemonId(id uint64) uint64 {
	var responseAPI ResponsePokeApiDetail

	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d", id)

	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		fmt.Print(err)
	}

	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&responseAPI); err != nil {
		fmt.Print(err)
	}

	var score uint64 = 0

	for _, stat := range responseAPI.Stats {
		score += stat.BaseStat
	}
	return score
}
