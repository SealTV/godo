package data

type List struct {
	Id     int    `db:"id"`
	Name   string `db:"Name"`
	UserId int    `db:"user_id"`
}
