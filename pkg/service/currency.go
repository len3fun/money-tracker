package service

import (
	moneytracker "github.com/len3fun/money-tracker"
	"github.com/len3fun/money-tracker/pkg/repository"
)

type CurrencyService struct {
	repo repository.Currency
}

func NewCurrencyService(repo repository.Currency) *CurrencyService {
	return &CurrencyService{repo: repo}
}

func (s *CurrencyService) CreateCurrency(item moneytracker.Currency) error {
	return s.repo.CreateCurrency(item)
}