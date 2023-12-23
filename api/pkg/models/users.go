package models

import (
	"accounting/pkg/db"
	"time"
)

type Users struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
	Password  string
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UsersRegister struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type UsersLogin struct {
	Email    string
	Password string
}

func (u *UsersRegister) Create() error {
	_, err := db.Db.Exec("INSERT INTO users(firstName, lastName, email, password, active, createdAt, updatedAt) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		u.FirstName, u.LastName, u.Email, u.Password, true, time.Now(), time.Now())

	return err
}

func (u *Users) GetByEmail(email string) error {
	err := db.Db.Get(u, "SELECT * FROM users WHERE email = $1", email)
	return err
}
