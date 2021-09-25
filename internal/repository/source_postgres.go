package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/len3fun/money-tracker/internal/models"
)

type SourcePostgres struct {
	db *sqlx.DB
}

func NewSourcePostgres(db *sqlx.DB) *SourcePostgres {
	return &SourcePostgres{db: db}
}

func (r *SourcePostgres) GetAllSources(userId int) ([]models.SourceOut, error) {
	sources := make([]models.SourceOut, 0)
	query := fmt.Sprintf("SELECT s.source_type, s.balance, c.ticket " +
		"FROM %s s LEFT JOIN %s c ON s.currency_id = c.id " +
		"WHERE s.user_id = $1", sourcesTable, currenciesTable)
	err := r.db.Select(&sources, query, userId)
	return sources, err
}

func (r *SourcePostgres) CreateSource(source models.Source) error {
	query := fmt.Sprintf("INSERT INTO %s(user_id, source_type, balance, currency_id) " +
		"VALUES ($1, $2, $3, $4)", sourcesTable)
	_, err := r.db.Exec(query, source.UserId, source.Type, source.Balance, source.CurrencyId)
	return err
}
