package service

import "github.com/len3fun/money-tracker/pkg/repository"

type Authorization interface {

}

type MoneySource interface {

}

type Service struct {
	Authorization
	MoneySource
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
