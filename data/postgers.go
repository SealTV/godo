package data

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	DB_USER           = "postgres"
	DB_NAME           = "todo_database"
	DB_SSL_MODE       = "disable"
	DB_DEFAULT_SCHEMA = "main_schema"
)

func IniDB() *sql.DB {
	db, err := sql.Open("postgres", getConnectionSating())
	if err != nil {
		log.Fatal(err)
	}

	// если ошибок нет, но не можем подключиться к базе данных,
	// то так же выходим
	if db == nil {
		panic("db nil")
	}

	_, err = db.Exec(fmt.Sprintf("SET search_path TO %s;", DB_DEFAULT_SCHEMA))
	if err != nil {
		log.Fatal(err)
	}

	//todo := Todo{
	//	UserId:      1,
	//	IsActive:    true,
	//	Description: "some description",
	//	ListId:      1,
	//	Title:       "Som title",
	//}
	//
	//todo = AddTodo(db, todo)
	//todo.Title = "Another title 2"
	//UpdateTodo(db, todo)
	//DeleteTodo(db, todo)
	//todo.Id = 5
	//DeleteTodo(db, todo)
	//
	//userM, err := GetUserModel(db, 1)
	//if err != nil{
	//	log.Fatal(err)
	//}
	//fmt.Println(userM)
	return db
}

func getConnectionSating() string {
	dbInfo := fmt.Sprintf("user=%s dbname=%s sslmode=%s", DB_USER, DB_NAME, DB_SSL_MODE)
	return dbInfo
}
