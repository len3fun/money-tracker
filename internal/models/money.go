package models

import (
	"errors"
	"github.com/shopspring/decimal"
	"time"
)

const (
	incomeField  = "income"
	expenseField = "expense"
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

type Activity struct {
	UserId       int             `json:"user_id" db:"user_id"`
	SourceId     int             `json:"source_id" db:"source_id"`
	CategoryId   int             `json:"category_id" db:"category_id"`
	Type         string          `json:"type" db:"activity_type"`
	Change       decimal.Decimal `json:"change" db:"change"`
	Label        string          `json:"label" db:"label"`
	ActivityDate time.Time       `json:"activity_date" db:"activity_date"`
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

type ActivitiesOut struct {
	Type         string          `json:"type" db:"activity_type"`
	Change       decimal.Decimal `json:"change" db:"change"`
	Label        string          `json:"label" db:"label"`
	ActivityDate time.Time       `json:"activity_date" db:"activity_date"`
}

func (a *Activity) Validate() error {
	if a.Type == "" {
		return errors.New("field 'type' shouldn't be empty")
	}
	if a.Type != incomeField && a.Type != expenseField {
		return errors.New("field 'type' should be equal to 'income' or 'expense'")
	}
	if a.Label == "" {
		return errors.New("field 'label' shouldn't be empty")
	}
	if a.Change.LessThanOrEqual(decimal.Zero) {
		return errors.New("field 'change' should be greater than zero")
	}
	return nil
}
