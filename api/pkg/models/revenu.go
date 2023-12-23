package models

import "time"

type Revenu struct {
	Id        int
	Name      string
	Amount    float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
