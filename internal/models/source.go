package models

import "github.com/shopspring/decimal"

type Source struct {
	Id      int             `json:"id"`
	UserId  int             `json:"user_id"`
	Type    string          `json:"type"`
	Balance decimal.Decimal `json:"balance"`
}
