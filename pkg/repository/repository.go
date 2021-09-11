package repository

import (
	"github.com/jmoiron/sqlx"
	moneytracker "github.com/len3fun/money-tracker"
)

type Authorization interface {
	CreateUser(user moneytracker.User) (int, error)
}

type MoneySource interface {

}

type Repository struct {
	Authorization
	MoneySource
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}