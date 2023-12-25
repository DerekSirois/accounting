package models

import (
	"accounting/pkg/db"
	"time"
)

type Revenu struct {
	Id        int
	Name      string
	Amount    float64
	UserId    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GetAllRevenuByUser(userId int) ([]*Revenu, error) {
	var r = make([]*Revenu, 0)
	err := db.Db.Select(&r, "SELECT * FROM revenu WHERE userId = $1", userId)
	return r, err
}

func GetRevenuById(id int) (*Revenu, error) {
	r := &Revenu{}
	err := db.Db.Get(r, "SELECT * FROM revenu WHERE id = $1", id)
	return r, err
}

func (r *Revenu) Create() error {
	_, err := db.Db.Exec("INSERT INTO revenu(name, amount, userId, createdAt, updatedAt) VALUES ($1, $2, $3, $4, $5)",
		r.Name, r.Amount, r.UserId, time.Now(), time.Now())
	return err
}

func (r *Revenu) Update() error {
	_, err := db.Db.Exec("UPDATE revenu SET name = $1, amount = $2, updatedAt = $3 WHERE id = $4",
		r.Name, r.Amount, time.Now(), r.Id)
	return err
}

func Delete(id int) error {
	_, err := db.Db.Exec("DELETE FROM revenu WHERE id = $1", id)
	return err
}
