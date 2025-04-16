package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Model struct {
	Movie MovieModel
	User  UserModel
}

func NewModel(db *sql.DB) Model {
	return Model{
		Movie: MovieModel{DB: db},
		User:  UserModel{DB: db},
	}
}
