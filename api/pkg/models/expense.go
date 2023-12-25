package models

import (
	"accounting/pkg/db"
	"time"
)

type Expense struct {
	Id        int
	Name      string
	Amount    float64
	UserId    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GetAllExpenseByUser(userId int) ([]*Expense, error) {
	var e = make([]*Expense, 0)
	err := db.Db.Select(&e, "SELECT * FROM expense WHERE userId = $1", userId)
	return e, err
}

func GetExpenseById(id int) (*Expense, error) {
	e := &Expense{}
	err := db.Db.Get(e, "SELECT * FROM expense WHERE id = $1", id)
	return e, err
}

func (e *Expense) Create() error {
	_, err := db.Db.Exec("INSERT INTO expense(name, amount, userId, createdAt, updatedAt) VALUES ($1, $2, $3, $4, $5)",
		e.Name, e.Amount, e.UserId, time.Now(), time.Now())
	return err
}

func (e *Expense) Update() error {
	_, err := db.Db.Exec("UPDATE expense SET name = $1, amount = $2, updatedAt = $3 WHERE id = $4",
		e.Name, e.Amount, time.Now(), e.Id)
	return err
}

func DeleteExpense(id int) error {
	_, err := db.Db.Exec("DELETE FROM expense WHERE id = $1", id)
	return err
}
