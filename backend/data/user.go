package data

import (
	"database/sql"

	"bitbucket.org/SealTV/go-site/backend/model"
)

func (db *pgConnector) GetAllUsers() (model.UsersCollection, error) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	return parseUserRows(rows)
}

func (db *pgConnector) GetUserById(id int) (model.User, error) {
	var user model.User
	err := db.QueryRow(`SELECT * FROM users WHERE id = $1 LIMIT 1`, id).
		Scan(&user.ID, &user.Login, &user.Password, &user.Email, &user.RegisterDate)

	if err != nil {
		return user, err
	}
	return user, nil
}

func (db *pgConnector) GetUserByLogin(login string) (model.User, error) {
	var user model.User
	err := db.QueryRow(`SELECT * FROM users WHERE (login = $1 OR email = $1) LIMIT 1`, login).
		Scan(&user.ID, &user.Login, &user.Password, &user.Email, &user.RegisterDate)

	if err != nil {
		return user, err
	}
	return user, nil
}

func (db *pgConnector) GetUserByLoginAndPassword(login, password string) (model.User, error) {
	var user model.User
	err := db.QueryRow(`SELECT * FROM users WHERE (login = $1 OR email = $1) AND password = $2 LIMIT 1`, login, password).
		Scan(&user.ID, &user.Login, &user.Password, &user.Email, &user.RegisterDate)

	if err != nil {
		return user, err
	}
	return user, nil
}

func (db *pgConnector) AddUser(user model.User) (model.User, error) {
	err := db.QueryRow(`INSERT INTO users(login, password, email) VALUES ($1, $2, $3) RETURNING id, register_date;`,
		user.Login, user.Password, user.Email).Scan(&user.ID, &user.RegisterDate)

	if err != nil {
		return user, err
	}
	return user, nil
}

func (db *pgConnector) UpdateUser(user model.User) (int64, error) {
	r, err := db.Exec(`UPDATE users SET login = $2, password = $3, email = $4 WHERE id = $1`,
		user.ID, user.Login, user.Password, user.Email)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

func (db *pgConnector) DeleteUser(user model.User) (int64, error) {
	return db.DeleteUserById(user.ID)
}

func (db *pgConnector) DeleteUserById(user int) (int64, error) {
	tx, err := db.Begin()
	r, err := tx.Exec(`DELETE FROM todos WHERE user_id = $1`, user)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	r, err = tx.Exec(`DELETE FROM lists WHERE user_id = $1`, user)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	r, err = db.Exec(`DELETE FROM users WHERE id = $1`, user)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return r.RowsAffected()
}

func parseUserRows(rows *sql.Rows) (model.UsersCollection, error) {
	result := model.UsersCollection{}
	for rows.Next() {
		user := model.User{}

		err := rows.Scan(&user.ID, &user.Login, &user.Password, &user.Email, &user.RegisterDate)
		if err != nil {
			return nil, err
		}

		result = append(result, user)
	}
	return result, nil
}
