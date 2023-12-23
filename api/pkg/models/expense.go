package models

import "time"

type Expense struct {
	Id        int
	Name      string
	Amount    float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
