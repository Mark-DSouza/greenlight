package data

import (
	"database/sql"
	"errors"
)

var ErrRecordNotFound = errors.New("record not found")

type Model struct {
	Movie MovieModel
}

func NewModel(db *sql.DB) Model {
	return Model{
		Movie: MovieModel{
			DB: db,
		},
	}
}
