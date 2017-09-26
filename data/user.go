package data

import "time"

type User struct {
	Id           int       `db:"id"`
	Login        string    `db:"login"`
	Password     string    `db:"password"`
	Email        string    `db:"email"`
	RegisterDate time.Time `db:"register_date"`
}
