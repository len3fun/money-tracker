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
	GetAllSources(userId int) ([]models.SourceOut, error)
	CreateSource(source models.Source) error
}

type Activity interface {
	GetAllActivities(userId int) ([]models.ActivitiesOut, error)
	CreateActivity(activity models.Activity) error
}

type Currency interface {
	CreateCurrency(item models.Currency) (int, error)
	GetAllCurrencies() ([]models.Currency, error)
}

type Repository struct {
	Authorization
	Currency
	Activity
	MoneySource
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Currency:      NewCurrencyPostgres(db),
		Activity:      NewActivityRepository(db),
		MoneySource:   NewSourcePostgres(db),
	}
}
