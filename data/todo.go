package data

import (
	"bitbucket.org/SealTV/go-site/model"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func (db *postgresConnector) GetAllTodos() (model.TodoCollection, error) {
	rows, err := db.Query("SELECT * FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return parseTodoRows(rows)
}

func (db *postgresConnector) GetAllTodosForUser(user model.User) (model.TodoCollection, error) {
	rows, err := db.Query("SELECT * FROM todos WHERE user_id = $1", user.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return parseTodoRows(rows)
}

func (db *postgresConnector) GetAllTodosForUserList(user model.User, list model.List) (model.TodoCollection, error) {
	rows, err := db.Query("SELECT * FROM todos WHERE user_id = $1 AND list_id = $2", user.Id, list.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return parseTodoRows(rows)
}

func (db *postgresConnector) AddTodo(todo model.Todo) (model.Todo, error) {
	err := db.QueryRow(`INSERT
			INTO todos(title, description, list_id, is_active, user_id)
			VALUES($1, $2, $3, $4, $5)
			RETURNING id, date_create;`,
		todo.Title, todo.Description, todo.ListId, todo.IsActive, todo.UserId).Scan(&todo.Id, &todo.DateCreate)
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (db *postgresConnector) UpdateTodo(todo model.Todo) (int64, error) {
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

func (db *postgresConnector) DeleteTodo(todo model.Todo) (int64, error) {
	return db.DeleteTodoById(todo.Id)
}

func (db *postgresConnector) DeleteTodoById(id int) (int64, error) {
	r, err := db.Exec(`DELETE FROM todos WHERE id = $1`, id)
	if err != nil {
		log.Fatal(err)
	}
	return r.RowsAffected()
}

func parseTodoRows(rows *sql.Rows) (model.TodoCollection, error) {
	result := model.TodoCollection{}
	for rows.Next() {
		todo := model.Todo{}

		err := rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.ListId, &todo.IsActive, &todo.UserId, &todo.DateCreate)
		if err != nil {
			return nil, err
		}

		result = append(result, todo)
	}
	return result, nil
}
