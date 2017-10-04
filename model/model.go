package model

import "time"

type UserModel struct {
	User
	TodoLists []TodoListModel
}

type TodoListModel struct {
	List
	Todos TodoCollection
}

type User struct {
	Id           int       `db:"id" json:"id"`
	Login        string    `db:"login" json:"login"`
	Password     string    `db:"password" json:"password"`
	Email        string    `db:"email" json:"email"`
	RegisterDate time.Time `db:"register_date" json:"register_date"`
}

type List struct {
	Id     int    `db:"id" json:"id"`
	Name   string `db:"name" json:"name"`
	UserId int    `db:"user_id" json:"user_id"`
}

type Todo struct {
	Id          int       `db:"id" json:"id"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	IsActive    bool      `db:"is_active" json:"is_active"`
	DateCreate  time.Time `db:"date_create" json:"date_create"`
	ListId      int       `db:"list_id" json:"list_id"`
	UserId      int       `db:"user_id" json:"user_id"`
}

type UsersCollection []User
type ListsCollection []List
type TodoCollection []Todo
