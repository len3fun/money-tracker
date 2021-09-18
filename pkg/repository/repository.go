package repository

import (
	"github.com/jmoiron/sqlx"
	moneytracker "github.com/len3fun/money-tracker"
)

type Authorization interface {
	CreateUser(user moneytracker.User) (int, error)
	GetUser(username, password string) (moneytracker.User, error)
}

type MoneySource interface {
}

type Currency interface {
	CreateCurrency(item moneytracker.Currency) error
	GetAllCurrencies() ([]moneytracker.Currency, error)
}

type Repository struct {
	Authorization
	Currency
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Currency:      NewCurrencyPostgres(db),
	}
}
