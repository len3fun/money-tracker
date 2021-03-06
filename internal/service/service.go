package service

import (
	"github.com/len3fun/money-tracker/internal/models"
	"github.com/len3fun/money-tracker/internal/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type MoneySource interface {
	GetAllSources(userId int) ([]models.SourceOut, error)
	CreateSource(source models.Source) error
}

type Currency interface {
	CreateCurrency(item models.Currency) (int, error)
	GetAllCurrencies() ([]models.Currency, error)
}

type Activity interface {
	GetAllActivities(userId int) ([]models.ActivitiesOut, error)
	CreateActivity(activity models.Activity) error
}

type Service struct {
	Authorization
	Currency
	Activity
	MoneySource
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Currency:      NewCurrencyService(repos.Currency),
		Activity:      NewActivityService(repos.Activity),
		MoneySource:   NewSource(repos.MoneySource),
	}
}
