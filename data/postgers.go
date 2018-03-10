package data

import (
	"database/sql"
	"fmt"
	"log"

	"bitbucket.org/SealTV/go-site/model"
	// init postgres sql lib
	_ "github.com/lib/pq"
)

type Config struct {
	UserDebugDB bool   `json:"use_debug_db"`
	Host        string `json:"host"`
	Port        int    `json:"port"`
	User        string `json:"user"`
	Password    string `json:"password"`
	DBName      string `json:"db_name"`
	Scheme      string `json:"scheme"`
	SSLMode     string `json:"ssl_mode"`
}

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

type pgConnector struct {
	*sql.DB
}

func New(c Config) DBConnector {
	if c.UserDebugDB {
		log.Println("Use db mock")
		return initMock()
	}
	log.Println("Use databese", c.DBName, c.Host, c.Port)
	return initDB(c)
}

func initDB(c Config) DBConnector {
	log.Println("DB connection string:", c.getConnectionString())
	db, err := sql.Open("postgres", c.getConnectionString())
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	if _, err = db.Exec(fmt.Sprintf("SET search_path TO %s;", c.Scheme)); err != nil {
		log.Fatal(err)
	}

	var connector DBConnector
	connector = &pgConnector{db}
	return connector
}

func (c *Config) getConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Host,
		c.Port,
		c.User,
		c.Password,
		c.DBName,
		c.SSLMode)
}
