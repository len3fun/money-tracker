package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {

}

type MoneySource interface {

}

type Repository struct {
	Authorization
	MoneySource
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}