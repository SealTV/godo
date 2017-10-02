package data

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func (db *PostgresConnector) GetAllTodos() (TodoCollection, error) {
	rows, err := db.Query("SELECT * FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return parseTodoRows(rows)
}

func (db *PostgresConnector) GetAllTodosForUser(user User) (TodoCollection, error) {
	rows, err := db.Query("SELECT * FROM todos WHERE user_id = $1", user.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return parseTodoRows(rows)
}

func (db *PostgresConnector) GetAllTodosForUserList(user User, list List) (TodoCollection, error) {
	rows, err := db.Query("SELECT * FROM todos WHERE user_id = $1 AND list_id = $2", user.Id, list.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return parseTodoRows(rows)
}

func (db *PostgresConnector) AddTodo(todo Todo) (Todo, error) {
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

func (db *PostgresConnector) UpdateTodo(todo Todo) (int64, error) {
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

func (db *PostgresConnector) DeleteTodo(todo Todo) (int64, error) {
	return db.DeleteTodoById(todo.Id)
}

func (db *PostgresConnector) DeleteTodoById(id int) (int64, error) {
	r, err := db.Exec(`DELETE FROM todos WHERE id = $1`, id)
	if err != nil {
		log.Fatal(err)
	}
	return r.RowsAffected()
}

func parseTodoRows(rows *sql.Rows) (TodoCollection, error) {
	result := TodoCollection{}
	for rows.Next() {
		todo := Todo{}

		err := rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.ListId, &todo.IsActive, &todo.UserId, &todo.DateCreate)
		if err != nil {
			return nil, err
		}

		result = append(result, todo)
	}
	return result, nil
}
