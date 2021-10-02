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

func (r *CurrencyPostgres) CreateCurrency(item models.Currency) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, ticket) VALUES ($1, $2) RETURNING id", currenciesTable)
	row := r.db.QueryRow(query, item.Name, item.Ticket)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *CurrencyPostgres) GetAllCurrencies() ([]models.Currency, error) {
	// todo: fix null if there are no records in db
	var currencies []models.Currency
	query := fmt.Sprintf("SELECT id, name, ticket FROM %s", currenciesTable)

	err := r.db.Select(&currencies, query)
	return currencies, err
}
