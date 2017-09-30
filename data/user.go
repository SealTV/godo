package data

import (
	"database/sql"
	"log"
	"time"
)

type User struct {
	Id           int       `db:"id"`
	Login        string    `db:"login"`
	Password     string    `db:"password"`
	Email        string    `db:"email"`
	RegisterDate time.Time `db:"register_date"`
}

type UsersCollection struct {
	Users []User `json:"users"`
}

func GetAllUsers(db *sql.DB) UsersCollection {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}

	return parseUserRows(rows)
}

func GetUserById(db *sql.DB, id int) (User, error) {
	var user User
	err := db.QueryRow(`SELECT * FROM users WHERE id = $1 LIMIT 1`, id).
		Scan(&user.Id, &user.Login, &user.Password, &user.Email, &user.RegisterDate)

	if err != nil {
		return user, err
	}
	return user, nil
}

func GetUserByLoginAndPassword(db *sql.DB, login, password string) (User, error) {
	var user User
	err := db.QueryRow(`SELECT * FROM users WHERE (login = $1 OR email = $1) AND password = $2 LIMIT 1`, login, password).
		Scan(&user.Id, &user.Login, &user.Password, &user.Email, &user.RegisterDate)

	if err != nil {
		return user, err
	}
	return user, nil
}

func AddUser(db *sql.DB, user User) (User, error) {
	err := db.QueryRow(`INSERT
			INTO users(login, password, email)
			VALUES ($1, $2, $3)
			RETURNING id, register_date;`).Scan(&user.Id, &user.RegisterDate)

	if err != nil {
		return user, err
	}
	return user, nil
}

func UpdateUser(db *sql.DB, user User) (int64, error) {
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

func DeleteUser(db *sql.DB, user User) (int64, error) {
	r, err := db.Exec(`DELETE FROM users WHERE id = $1`, user.Id)
	if err != nil {
		log.Fatal(err)
	}
	return r.RowsAffected()
}

func parseUserRows(rows *sql.Rows) UsersCollection {
	result := UsersCollection{}
	for rows.Next() {
		user := User{}

		err := rows.Scan(&user.Id, &user.Login, &user.Password, &user.Email, &user.RegisterDate)
		if err != nil {
			log.Fatal(err)
		}

		result.Users = append(result.Users, user)
	}
	return result
}
