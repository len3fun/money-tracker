package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	moneytracker "github.com/len3fun/money-tracker"
)

type CurrencyPostgres struct {
	db *sqlx.DB
}

func NewCurrencyPostgres(db *sqlx.DB) *CurrencyPostgres {
	return &CurrencyPostgres{db: db}
}

func (r *CurrencyPostgres) CreateCurrency(item moneytracker.Currency) error {
	query := fmt.Sprintf("INSERT INTO %s (name, ticket) VALUES ($1, $2) RETURNING id", currenciesTable)
	_, err := r.db.Exec(query, item.Name, item.Ticket)
	return err
}
