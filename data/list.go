package data

import (
	"database/sql"
	"log"
)

func (db *PostgresConnector) GetAllLists() (ListsCollection, error) {
	rows, err := db.Query("SELECT * FROM lists")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	return parseListsRows(rows)
}

func (db *PostgresConnector) GetAllListsForUser(user User) (ListsCollection, error) {
	return db.GetAllListsForUserId(user.Id)
}

func (db *PostgresConnector) GetAllListsForUserId(user int) (ListsCollection, error) {
	rows, err := db.Query("SELECT * FROM lists WHERE user_id = $1", user)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	return parseListsRows(rows)
}

func (db *PostgresConnector) GetListById(id int) (List, error) {
	var list List
	err := db.QueryRow(`SELECT * FROM lists WHERE id = $1`, id).Scan(&list.Id, &list.Name, &list.UserId)
	if err != nil {
		return list, err
	}

	return list, nil
}

func (db *PostgresConnector) AddList(list List) (List, error) {
	err := db.QueryRow(`INSERT
			INTO lists(name, user_id)
			VALUES($1, $2)
			RETURNING id;`,
		list.Name, list.UserId).Scan(&list.Id)
	if err != nil {
		return list, err
	}

	return list, nil
}

func (db *PostgresConnector) UpdateList(list List) (int64, error) {
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

func (db *PostgresConnector) DeleteList(list List) (int64, error) {
	return db.DeleteListById(list.Id)
}

func (db *PostgresConnector) DeleteListById(list int) (int64, error) {
	tx, err := db.Begin()
	r, err := tx.Exec(`DELETE FROM todos WHERE list_id = $1`, list)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	r, err = tx.Exec(`DELETE FROM lists WHERE id = $1`, list)
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

func parseListsRows(rows *sql.Rows) (ListsCollection, error) {
	result := ListsCollection{}
	for rows.Next() {
		list := List{}

		err := rows.Scan(&list.Id, &list.Name, &list.UserId)
		if err != nil {
			return nil, err
		}

		result = append(result, list)
	}
	return result, nil
}
