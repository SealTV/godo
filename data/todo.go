package data

import "time"

type Todo struct {
	Id          int       `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	IsActive    bool      `db:"is_active"`
	DateCreate  time.Time `db:"date_create"`
	ListId      int       `db:"list_id"`
	UserId      int       `db:"user_id"`
}
