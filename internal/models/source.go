package models

import (
	"errors"
	"github.com/shopspring/decimal"
)

type Source struct {
	UserId  int             `json:"user_id"`
	Type    string          `json:"type"`
	Balance decimal.Decimal `json:"balance"`
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
	return nil
}
