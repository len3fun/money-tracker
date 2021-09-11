package money_tracker

import (
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
