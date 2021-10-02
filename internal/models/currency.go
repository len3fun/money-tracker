package models

import (
	"errors"
)

type Currency struct {
	Id     int    `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	Ticket string `json:"ticket" db:"ticket"`
}

func (c *Currency) Validate() error {
	if c.Name == "" || c.Ticket == "" {
		return errors.New("invalid currency input, fields 'name' and 'ticket' are required")
	}
	return nil
}
