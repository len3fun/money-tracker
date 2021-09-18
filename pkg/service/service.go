package service

import (
	moneytracker "github.com/len3fun/money-tracker"
	"github.com/len3fun/money-tracker/pkg/repository"
)

type Authorization interface {
	CreateUser(user moneytracker.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type MoneySource interface {

}

type Currency interface {
	CreateCurrency(item moneytracker.Currency) error
	GetAllCurrencies() ([]moneytracker.Currency, error)
}

type Service struct {
	Authorization
	Currency
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Currency: NewCurrencyService(repos.Currency),
	}
}
