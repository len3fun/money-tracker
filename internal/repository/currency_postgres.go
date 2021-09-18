package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/len3fun/money-tracker/internal/models"
)

type CurrencyPostgres struct {
	db *sqlx.DB
}

func NewCurrencyPostgres(db *sqlx.DB) *CurrencyPostgres {
	return &CurrencyPostgres{db: db}
}

func (r *CurrencyPostgres) CreateCurrency(item models.Currency) error {
	query := fmt.Sprintf("INSERT INTO %s (name, ticket) VALUES ($1, $2) RETURNING id", currenciesTable)
	_, err := r.db.Exec(query, item.Name, item.Ticket)
	return err
}

func (r *CurrencyPostgres) GetAllCurrencies() ([]models.Currency, error) {
	var currencies []models.Currency
	query := fmt.Sprintf("SELECT name, ticket FROM %s", currenciesTable)

	err := r.db.Select(&currencies, query)
	return currencies, err
}
