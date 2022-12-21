package controllers

import (
	"farmacare/configs"
	"farmacare/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type iMatchRepositorySave interface {
	Save(*models.Match) (*models.Match, error)
}

type matchServiceCreate struct {
	iMatchRepositorySave
}

type NewMatch struct {
	Name string `json:"name"`
}

func Create() gin.HandlerFunc {
	return func(c *gin.Context) {

		var match NewMatch

		if err := c.ShouldBindJSON(&match); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "Data": err})
			return
		}

		newMatch := models.Match{
			Name: match.Name,
		}

		database := configs.Connect()
		matchRepository := models.New(database)
		matchService := matchServiceCreate{matchRepository}
		res, err := matchService.Save(&newMatch)

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

func GetMatchPokemon() gin.HandlerFunc {
	return func(c *gin.Context) {

		database := configs.Connect()
		matchPokemonRepository := models.New(database)
		matchId := c.Query("matchId")
		isFinished := c.Query("isFinished")

		intVar, err := strconv.Atoi(matchId)
		boolVar, err := strconv.ParseBool(isFinished)

		results, err := matchPokemonRepository.GetByMatchId(intVar, boolVar)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "Data": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "success", "Data": results})
	}
}
