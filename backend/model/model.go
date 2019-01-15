package model

import (
	"time"
)

type (
	UserModel struct {
		User      `json:"user"`
		TodoLists []TodoListModel `json:"todoLists"`
	}

	TodoListModel struct {
		List  `json:"list"`
		Todos TodoCollection `json:"todos"`
	}

	User struct {
		ID           int       `db:"id" json:"id"`
		Login        string    `db:"login" json:"login"`
		Password     string    `db:"password" json:"password"`
		Email        string    `db:"email" json:"email"`
		RegisterDate time.Time `db:"register_date" json:"register_date"`
	}

	List struct {
		ID     int    `db:"id" json:"id"`
		Name   string `db:"name" json:"name"`
		UserID int    `db:"user_id" json:"user_id"`
	}

	Todo struct {
		ID          int       `db:"id" json:"id"`
		Title       string    `db:"title" json:"title" form:"title" binding:"required"`
		Description string    `db:"description" json:"description" form:"description" binding:"required"`
		IsActive    bool      `db:"is_active" json:"is_active" form:"is_active"`
		DateCreate  time.Time `db:"date_create" json:"date_create"`
		ListID      int       `db:"list_id" json:"list_id" form:"list_id" binding:"required"`
		UserID      int       `db:"user_id" json:"user_id" form:"user_id" binding:"required"`
	}

	UsersCollection []User
	ListsCollection []List
	TodoCollection  []Todo
)
