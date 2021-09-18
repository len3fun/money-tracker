package service

import (
	"github.com/len3fun/money-tracker/internal/models"
	"github.com/len3fun/money-tracker/internal/repository"
)

type Source struct {
	repo repository.MoneySource
}

func NewSource(repo repository.MoneySource) *Source {
	return &Source{repo: repo}
}

func (s *Source) GetAllSources(userId int) ([]models.SourceOut, error) {
	return s.repo.GetAllSources(userId)
}


