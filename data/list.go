package data

import (
	"database/sql"
	"log"
)

type List struct {
	Id     int    `db:"id"`
	Name   string `db:"name"`
	UserId int    `db:"user_id"`
}
type ListsCollection struct {
	lists []List `json:"lists"`
}

func GetAllLists(db *sql.DB) ListsCollection {
	rows, err := db.Query("SELECT * FROM lists")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	return parseListsRows(rows)
}

func GetAllListsForUser(db *sql.DB, user User) ListsCollection {
	rows, err := db.Query("SELECT * FROM lists WHERE user_id = $1", user.Id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	return parseListsRows(rows)
}

func AddList(db *sql.DB, list List) List {
	err := db.QueryRow(`INSERT
			INTO lists(name, user_id)
			VALUES($1, $2)
			RETURNING id;`,
		list.Name, list.UserId).Scan(&list.Id)
	if err != nil {
		log.Fatal(err)
	}

	return list
}

func UpdateList(db *sql.DB, list List) (int64, error) {
	r, err := db.Exec(
		`UPDATE lists
				SET name = $2, user_id = $3
				WHERE id = $1`,
		list.Id, list.Name, list.UserId)
	if err != nil {
		log.Fatal(err)
	}
	return r.RowsAffected()
}

func DeleteList(db *sql.DB, list List) (int64, error) {
	r, err := db.Exec(`DELETE FROM lists WHERE id = $1`, list.Id)
	if err != nil {
		log.Fatal(err)
	}
	return r.RowsAffected()
}

func parseListsRows(rows *sql.Rows) ListsCollection {
	result := ListsCollection{}
	for rows.Next() {
		list := List{}

		err := rows.Scan(&list.Id, &list.Name, &list.UserId)
		if err != nil {
			log.Fatal(err)
		}

		result.lists = append(result.lists, list)
	}
	return result
}
