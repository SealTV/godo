package data

import (
	"database/sql"
	_ "github.com/lib/pq"
	"time"
)

type Todo struct {
	Id          int       `db:"id" json:"id"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	IsActive    bool      `db:"is_active" json:"is_active"`
	DateCreate  time.Time `db:"date_create" json:"date_create"`
	ListId      int       `db:"list_id" json:"list_id"`
	UserId      int       `db:"user_id" json:"user_id"`
}

type TodoCollection struct {
	Todo []Todo `json:"items"`
}

func GetAllTodos(db *sql.DB) TodoCollection {
	s := "SELECT * FROM todos"
	rows, err := db.Query(s)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	return parseTodoRows(rows)
}

func GetAllTodosForUser(db *sql.DB, user User) TodoCollection {
	s := "SELECT * FROM todos WHERE user_id = $1"
	rows, err := db.Query(s, user.Id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	return parseTodoRows(rows)
}

func GetAllTodosForUserList(db *sql.DB, user User, list List) TodoCollection {
	s := "SELECT * FROM todos WHERE user_id = $1  AND list_id = $2"
	rows, err := db.Query(s, user.Id, list.Id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	return parseTodoRows(rows)
}

func parseTodoRows(rows *sql.Rows) TodoCollection {
	result := TodoCollection{}
	for rows.Next() {
		todo := Todo{}

		err := rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.ListId, &todo.IsActive, &todo.UserId, &todo.DateCreate)
		if err != nil {
			panic(err)
		}

		result.Todo = append(result.Todo, todo)
	}
	return result
}
