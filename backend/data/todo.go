package data

import (
	"database/sql"

	"bitbucket.org/SealTV/go-site/backend/model"
)

func (db *pgConnector) GetAllTodos() (model.TodoCollection, error) {
	rows, err := db.Query("SELECT * FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return parseTodoRows(rows)
}

func (db *pgConnector) GetAllTodosForUser(user model.User) (model.TodoCollection, error) {
	rows, err := db.Query("SELECT * FROM todos WHERE user_id = $1", user.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return parseTodoRows(rows)
}

func (db *pgConnector) GetAllTodosForUserList(user model.User, list model.List) (model.TodoCollection, error) {
	rows, err := db.Query("SELECT * FROM todos WHERE user_id = $1 AND list_id = $2", user.Id, list.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return parseTodoRows(rows)
}

func (db *pgConnector) AddTodo(todo model.Todo) (model.Todo, error) {
	err := db.QueryRow(`INSERT INTO todos(title, description, list_id, is_active, user_id)
			VALUES($1, $2, $3, $4, $5)
			RETURNING id, date_create;`,
		todo.Title, todo.Description, todo.ListId, todo.IsActive, todo.UserId).Scan(&todo.Id, &todo.DateCreate)
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (db *pgConnector) UpdateTodo(todo model.Todo) (int64, error) {
	r, err := db.Exec(
		`UPDATE todos
				SET title = $2, description = $3, list_id = $4, is_active = $5, user_id = $6
				WHERE id = $1`,
		todo.Id, todo.Title, todo.Description, todo.ListId, todo.IsActive, todo.UserId)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

func (db *pgConnector) DeleteTodo(todo model.Todo) (int64, error) {
	return db.DeleteTodoById(todo.Id)
}

func (db *pgConnector) DeleteTodoById(id int) (int64, error) {
	r, err := db.Exec(`DELETE FROM todos WHERE id = $1`, id)
	if err != nil {
		return 0, err
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

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return result, nil
}
