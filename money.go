package money_tracker

import (
	"errors"
	"github.com/shopspring/decimal"
	"time"
)

type Sources struct {
	Id      int             `json:"id"`
	UserId  int             `json:"user_id"`
	Type    string          `json:"type"`
	Balance decimal.Decimal `json:"balance"`
}

type Categories struct {
	Id     int    `json:"id"`
	UserId int    `json:"user_id"`
	Name   string `json:"name"`
}

type Activities struct {
	Id           int             `json:"id"`
	UserId       int             `json:"user_id"`
	SourceId     int             `json:"source_id"`
	CategoryId   int             `json:"category_id"`
	Type         string          `json:"type"`
	Change       decimal.Decimal `json:"change"`
	Label        string          `json:"label"`
	ActivityDate time.Time       `json:"activity_date"`
}

type Currency struct {
	Name   string `json:"name" db:"name"`
	Ticket string `json:"ticket" db:"ticket"`
}

func (c *Currency) Validate() error {
	if c.Name == "" || c.Ticket == "" {
		return errors.New("invalid currency input, fields 'name' and 'ticket' are required")
	}
	return nil
}