package data

import (
	"database/sql"
	"fmt"
	"log"

	"bitbucket.org/SealTV/go-site/model"
	_ "github.com/lib/pq"
)

const (
	dbUser          = "postgres"
	dbName          = "todo_database"
	dbSslMode       = "disable"
	dbDefaultSchema = "main_schema"
)

//UserTable provide methods for table queries
type UserTable interface {
	GetAllUsers() (model.UsersCollection, error)
	GetUserById(id int) (model.User, error)
	GetUserByLoginAndPassword(login, password string) (model.User, error)
	AddUser(user model.User) (model.User, error)
	UpdateUser(user model.User) (int64, error)
	DeleteUser(user model.User) (int64, error)
	DeleteUserById(user int) (int64, error)
}

//ListTable provide methods for table queries
type ListTable interface {
	GetAllLists() (model.ListsCollection, error)
	GetAllListsForUser(user model.User) (model.ListsCollection, error)
	GetAllListsForUserId(user int) (model.ListsCollection, error)
	GetListById(id int) (model.List, error)
	AddList(list model.List) (model.List, error)
	UpdateList(list model.List) (int64, error)
	DeleteList(list model.List) (int64, error)
	DeleteListById(list int) (int64, error)
}

//TodoTable provide methods for table queries
type TodoTable interface {
	GetAllTodos() (model.TodoCollection, error)
	GetAllTodosForUser(user model.User) (model.TodoCollection, error)
	GetAllTodosForUserList(user model.User, list model.List) (model.TodoCollection, error)
	AddTodo(todo model.Todo) (model.Todo, error)
	UpdateTodo(todo model.Todo) (int64, error)
	DeleteTodo(todo model.Todo) (int64, error)
	DeleteTodoById(id int) (int64, error)
}

//DBConnector provide methods for datase
type DBConnector interface {
	GetUserModel(id int) (model.UserModel, error)
	//User section
	UserTable
	//List section
	ListTable
	//Todos section
	TodoTable
}

type postgresConnector struct {
	*sql.DB
}

//InitDB create connector for db
func InitDB() DBConnector {
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
	connector = &postgresConnector{db}
	return connector
}

func getConnectionSating() string {
	dbInfo := fmt.Sprintf("user=%s dbname=%s sslmode=%s", dbUser, dbName, dbSslMode)
	return dbInfo
}
