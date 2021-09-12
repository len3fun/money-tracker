package service

import (
	moneytracker "github.com/len3fun/money-tracker"
	"github.com/len3fun/money-tracker/pkg/repository"
)

type Authorization interface {
	CreateUser(user moneytracker.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type MoneySource interface {

}

type Service struct {
	Authorization
	MoneySource
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
