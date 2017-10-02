package data

import (
	"database/sql"
	"log"
)

func (db *PostgresConnector) GetAllUsers() (UsersCollection, error) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	return parseUserRows(rows)
}

func (db *PostgresConnector) GetUserById(id int) (User, error) {
	var user User
	err := db.QueryRow(`SELECT * FROM users WHERE id = $1 LIMIT 1`, id).
		Scan(&user.Id, &user.Login, &user.Password, &user.Email, &user.RegisterDate)

	if err != nil {
		return user, err
	}
	return user, nil
}

func (db *PostgresConnector) GetUserByLoginAndPassword(login, password string) (User, error) {
	var user User
	err := db.QueryRow(`SELECT * FROM users WHERE (login = $1 OR email = $1) AND password = $2 LIMIT 1`, login, password).
		Scan(&user.Id, &user.Login, &user.Password, &user.Email, &user.RegisterDate)

	if err != nil {
		return user, err
	}
	return user, nil
}

func (db *PostgresConnector) AddUser(user User) (User, error) {
	err := db.QueryRow(`INSERT
			INTO users(login, password, email)
			VALUES ($1, $2, $3)
			RETURNING id, register_date;`,
		user.Login, user.Password, user.Email).Scan(&user.Id, &user.RegisterDate)

	if err != nil {
		return user, err
	}
	return user, nil
}

func (db *PostgresConnector) UpdateUser(user User) (int64, error) {
	r, err := db.Exec(
		`UPDATE users
				SET login = $2, password = $3, email = $4
				WHERE id = $1`,
		user.Id, user.Login, user.Password, user.Email)
	if err != nil {
		log.Fatal(err)
	}
	return r.RowsAffected()
}

func (db *PostgresConnector) DeleteUser(user User) (int64, error) {
	return db.DeleteUserById(user.Id)
}

func (db *PostgresConnector) DeleteUserById(user int) (int64, error) {
	r, err := db.Exec(`DELETE FROM users WHERE id = $1`, user)
	if err != nil {
		log.Fatal(err)
	}
	return r.RowsAffected()
}

func parseUserRows(rows *sql.Rows) (UsersCollection, error) {
	result := UsersCollection{}
	for rows.Next() {
		user := User{}

		err := rows.Scan(&user.Id, &user.Login, &user.Password, &user.Email, &user.RegisterDate)
		if err != nil {
			return nil, err
		}

		result = append(result, user)
	}
	return result, nil
}
