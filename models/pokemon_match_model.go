package models

import (
	"time"

	"gorm.io/gorm"
)

type MatchPokemon struct {
	ID        uint64    `json:"id" gorm:"column:id"`
	MatchId   uint64    `json:"matchId" gorm:"column:match_id"`
	Match     Match     `json:"match,omitempty" gorm:"foreignKey:MatchId;references:ID"`
	PokemonId uint64    `json:"pokemonId" gorm:"column:pokemon_id"`
	Position  uint64    `json:"position" gorm:"column:position"`
	Score     uint64    `json:"score" gorm:"column:score"`
	IsFraud   bool      `json:"isFraud" gorm:"column:is_fraud; DEFAULT:false"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
}

type MatchPokemonRepository struct {
	Database *gorm.DB
}

type ListMatchPokemonResponse struct {
	MatchPokemons []MatchPokemon `json:"match_pokemons"`
}

func NewMatchPokemon(d *gorm.DB) MatchPokemonRepository {
	return MatchPokemonRepository{
		Database: d,
	}
}

func (r MatchPokemonRepository) Save(w *MatchPokemon) (*MatchPokemon, error) {
	res := r.Database.Create(w)
	return w, res.Error
}

func (r MatchPokemonRepository) GetPokemonByMatchId(matchId int) (ListMatchPokemonResponse, error) {

	matches := []MatchPokemon{}

	results := r.Database.Where("match_id = ?", matchId).Order("score desc").Find(&matches)

	listTokensResponse := ListMatchPokemonResponse{
		MatchPokemons: matches,
	}

	return listTokensResponse, results.Error
}

func (r MatchPokemonRepository) GetPokemonsByDate(startDate time.Time, endDate time.Time) (ListMatchPokemonResponse, error) {

	matches := []MatchPokemon{}

	results := r.Database.Where("created_at >= ? and created_at <= ?", startDate, endDate).Order("score desc").Find(&matches)

	listTokensResponse := ListMatchPokemonResponse{
		MatchPokemons: matches,
	}

	return listTokensResponse, results.Error
}

func (r MatchPokemonRepository) GetAllPokemons() (ListMatchPokemonResponse, error) {

	matches := []MatchPokemon{}

	results := r.Database.Order("score desc").Find(&matches)

	listTokensResponse := ListMatchPokemonResponse{
		MatchPokemons: matches,
	}

	return listTokensResponse, results.Error
}
