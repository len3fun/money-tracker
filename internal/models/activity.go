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

type Activity struct {
	UserId       int             `json:"user_id" db:"user_id"`
	SourceId     int             `json:"source_id" db:"source_id"`
	CategoryId   int             `json:"category_id" db:"category_id"`
	Type         string          `json:"type" db:"activity_type"`
	Change       decimal.Decimal `json:"change" db:"change"`
	Label        string          `json:"label" db:"label"`
	ActivityDate time.Time       `json:"activity_date" db:"activity_date"`
}
type ActivitiesOut struct {
	Type         string          `json:"type" db:"activity_type"`
	Change       decimal.Decimal `json:"change" db:"change"`
	Label        string          `json:"label" db:"label"`
	ActivityDate time.Time       `json:"activity_date" db:"activity_date"`
}

func (a *Activity) Validate() error {
	if a.Type == "" {
		return errors.New("field 'type' mustn't be empty")
	}
	if a.Type != incomeField && a.Type != expenseField {
		return errors.New("field 'type' must be equal to 'income' or 'expense'")
	}
	if a.Label == "" {
		return errors.New("field 'label' mustn't be empty")
	}
	if a.Change.LessThanOrEqual(decimal.Zero) {
		return errors.New("field 'change' must be greater than zero")
	}
	if a.SourceId == 0 {
		return errors.New("field 'source_id' mustn't be empty")
	}
	return nil
}
