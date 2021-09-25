package models

import (
	"errors"
	"github.com/shopspring/decimal"
)

type Source struct {
	UserId     int             `json:"user_id"`
	Type       string          `json:"type"`
	Balance    decimal.Decimal `json:"balance"`
	CurrencyId int             `json:"currency_id"`
}

type SourceOut struct {
	Type    string          `json:"type" db:"source_type"`
	Balance decimal.Decimal `json:"balance" db:"balance"`
	Ticket  string          `json:"currency_id" db:"ticket"`
}

func (s *Source) Validate() error {
	if s.Type == "" {
		return errors.New("'type' field shouldn't be empty")
	}
	if s.CurrencyId <= 0 {
		return errors.New("'currency_id' field should be more than 0")
	}
	return nil
}
