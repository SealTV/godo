package data

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
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

type TodoCollection []Todo

func GetAllTodos(db *sql.DB) TodoCollection {
	rows, err := db.Query("SELECT * FROM todos")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	return parseTodoRows(rows)
}

func GetAllTodosForUser(db *sql.DB, user User) TodoCollection {
	rows, err := db.Query("SELECT * FROM todos WHERE user_id = $1", user.Id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	return parseTodoRows(rows)
}

func GetAllTodosForUserList(db *sql.DB, user User, list List) TodoCollection {
	rows, err := db.Query("SELECT * FROM todos WHERE user_id = $1 AND list_id = $2", user.Id, list.Id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	return parseTodoRows(rows)
}

func AddTodo(db *sql.DB, todo Todo) Todo {
	err := db.QueryRow(`INSERT
			INTO todos(title, description, list_id, is_active, user_id)
			VALUES($1, $2, $3, $4, $5)
			RETURNING id, date_create;`,
		todo.Title, todo.Description, todo.ListId, todo.IsActive, todo.UserId).Scan(&todo.Id, &todo.DateCreate)
	if err != nil {
		log.Fatal(err)
	}

	return todo
}

func UpdateTodo(db *sql.DB, todo Todo) (int64, error) {
	r, err := db.Exec(
		`UPDATE todos
				SET title = $2, description = $3, list_id = $4, is_active = $5, user_id = $6
				WHERE id = $1`,
		todo.Id, todo.Title, todo.Description, todo.ListId, todo.IsActive, todo.UserId)
	if err != nil {
		log.Fatal(err)
	}
	return r.RowsAffected()
}

func DeleteTodo(db *sql.DB, todo Todo) (int64, error) {
	r, err := db.Exec(`DELETE FROM todos WHERE id = $1`, todo.Id)
	if err != nil {
		log.Fatal(err)
	}
	return r.RowsAffected()
}

func parseTodoRows(rows *sql.Rows) TodoCollection {
	result := TodoCollection{}
	for rows.Next() {
		todo := Todo{}

		err := rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.ListId, &todo.IsActive, &todo.UserId, &todo.DateCreate)
		if err != nil {
			log.Fatal(err)
		}

		result = append(result, todo)
	}
	return result
}
