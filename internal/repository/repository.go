package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/len3fun/money-tracker/internal/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type MoneySource interface {
}

type Currency interface {
	CreateCurrency(item models.Currency) error
	GetAllCurrencies() ([]models.Currency, error)
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
