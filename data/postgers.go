package data

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	dbUser          = "postgres"
	dbName          = "todo_database"
	dbSslMode       = "disable"
	dbDefaultSchema = "main_schema"
)

type DBConnector interface {
	GetUserModel(id int) (UserModel, error)

	//User section
	GetAllUsers() (UsersCollection, error)
	GetUserById(id int) (User, error)
	GetUserByLoginAndPassword(login, password string) (User, error)
	AddUser(user User) (User, error)
	UpdateUser(user User) (int64, error)
	DeleteUser(user User) (int64, error)
	DeleteUserById(user int) (int64, error)

	//List section
	GetAllLists() (ListsCollection, error)
	GetAllListsForUser(user User) (ListsCollection, error)
	GetAllListsForUserId(user int) (ListsCollection, error)
	GetListById(id int) (List, error)
	AddList(list List) (List, error)
	UpdateList(list List) (int64, error)
	DeleteList(list List) (int64, error)
	DeleteListById(list int) (int64, error)

	//User section
	GetAllTodos() (TodoCollection, error)
	GetAllTodosForUser(user User) (TodoCollection, error)
	GetAllTodosForUserList(user User, list List) (TodoCollection, error)
	AddTodo(todo Todo) (Todo, error)
	UpdateTodo(todo Todo) (int64, error)
	DeleteTodo(todo Todo) (int64, error)
	DeleteTodoById(id int) (int64, error)
}

type PostgresConnector struct {
	*sql.DB
}

func IniDB() DBConnector {
	db, err := sql.Open("postgres", getConnectionSating())
	if err != nil {
		log.Fatal(err)
	}

	// если ошибок нет, но не можем подключиться к базе данных,
	// то так же выходим
	if db == nil {
		panic("db nil")
	}

	_, err = db.Exec(fmt.Sprintf("SET search_path TO %s;", dbDefaultSchema))
	if err != nil {
		log.Fatal(err)
	}

	var connector DBConnector
	connector = &PostgresConnector{db}
	return connector
}

func getConnectionSating() string {
	dbInfo := fmt.Sprintf("user=%s dbname=%s sslmode=%s", dbUser, dbName, dbSslMode)
	return dbInfo
}
