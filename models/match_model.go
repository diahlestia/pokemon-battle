package models

import (
	"time"

	"gorm.io/gorm"
)

type Match struct {
	ID         uint64    `json:"id" gorm:"column:id"`
	Name       string    `json:"name" gorm:"column:name"`
	IsFinished bool      `json:"isFinished" gorm:"column:is_finished; DEFAULT:false"`
	CreatedAt  time.Time `json:"createdAt" gorm:"column:created_at"`
}

type MatchResponse struct {
	ID         uint64    `json:"id" gorm:"column:id"`
	Name       string    `json:"name" gorm:"column:name"`
	IsFinished bool      `json:"isFinished" gorm:"column:is_finished; DEFAULT:false"`
	CreatedAt  time.Time `json:"createdAt" gorm:"column:created_at"`
}

type Repository struct {
	Database *gorm.DB
}

type MetaResponse struct {
	Page        int `json:"page"`
	Limit       int `json:"limit"`
	TotalRecord int `json:"totalRecords"`
	TotalPages  int `json:"totalPages"`
}

type ListMatchResponse struct {
	Matches []Match `json:"matches"`
}

func New(d *gorm.DB) Repository {
	return Repository{
		Database: d,
	}
}

func (r Repository) Save(w *Match) (*Match, error) {
	res := r.Database.Create(w)
	return w, res.Error
}

func (r Repository) GetByMatchId(matchId int, isFinised bool) (ListMatchResponse, error) {

	matches := []Match{}

	results := r.Database.Where("id = ? AND is_finished = ?", matchId, isFinised).Order("created_at desc").Find(&matches)

	listTokensResponse := ListMatchResponse{
		Matches: matches,
	}

	return listTokensResponse, results.Error
}
