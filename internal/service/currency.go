package service

import (
	"github.com/len3fun/money-tracker/internal/models"
	"github.com/len3fun/money-tracker/internal/repository"
)

type CurrencyService struct {
	repo repository.Currency
}

func NewCurrencyService(repo repository.Currency) *CurrencyService {
	return &CurrencyService{repo: repo}
}

func (s *CurrencyService) CreateCurrency(item models.Currency) (int, error) {
	return s.repo.CreateCurrency(item)
}

func (s *CurrencyService) GetAllCurrencies() ([]models.Currency, error) {
	return s.repo.GetAllCurrencies()
}