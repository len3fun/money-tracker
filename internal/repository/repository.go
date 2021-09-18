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

type Activity interface {
	GetAllActivities(userId int) ([]models.ActivitiesOut, error)
}

type Currency interface {
	CreateCurrency(item models.Currency) error
	GetAllCurrencies() ([]models.Currency, error)
}

type Repository struct {
	Authorization
	Currency
	Activity
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Currency:      NewCurrencyPostgres(db),
		Activity:      NewActivityRepository(db),
	}
}