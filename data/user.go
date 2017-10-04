package data

import (
	"bitbucket.org/SealTV/go-site/model"
	"database/sql"
	"log"
)

func (db *postgresConnector) GetAllUsers() (model.UsersCollection, error) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	return parseUserRows(rows)
}

func (db *postgresConnector) GetUserById(id int) (model.User, error) {
	var user model.User
	err := db.QueryRow(`SELECT * FROM users WHERE id = $1 LIMIT 1`, id).
		Scan(&user.Id, &user.Login, &user.Password, &user.Email, &user.RegisterDate)

	if err != nil {
		return user, err
	}
	return user, nil
}

func (db *postgresConnector) GetUserByLoginAndPassword(login, password string) (model.User, error) {
	var user model.User
	err := db.QueryRow(`SELECT * FROM users WHERE (login = $1 OR email = $1) AND password = $2 LIMIT 1`, login, password).
		Scan(&user.Id, &user.Login, &user.Password, &user.Email, &user.RegisterDate)

	if err != nil {
		return user, err
	}
	return user, nil
}

func (db *postgresConnector) AddUser(user model.User) (model.User, error) {
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

func (db *postgresConnector) UpdateUser(user model.User) (int64, error) {
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

func (db *postgresConnector) DeleteUser(user model.User) (int64, error) {
	return db.DeleteUserById(user.Id)
}

func (db *postgresConnector) DeleteUserById(user int) (int64, error) {
	r, err := db.Exec(`DELETE FROM users WHERE id = $1`, user)
	if err != nil {
		log.Fatal(err)
	}
	return r.RowsAffected()
}

func parseUserRows(rows *sql.Rows) (model.UsersCollection, error) {
	result := model.UsersCollection{}
	for rows.Next() {
		user := model.User{}

		err := rows.Scan(&user.Id, &user.Login, &user.Password, &user.Email, &user.RegisterDate)
		if err != nil {
			return nil, err
		}

		result = append(result, user)
	}
	return result, nil
}
