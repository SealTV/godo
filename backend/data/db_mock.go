package data

import (
	"fmt"
	"log"
	"time"

	"bitbucket.org/SealTV/go-site/backend/model"
)

type dbMock struct {
	users map[int]model.User
	lists map[int]model.List
	todos map[int]model.Todo
}

func initMock() *dbMock {
	return &dbMock{
		users: map[int]model.User{
			1: {
				Id:           1,
				Login:        "SealTV",
				Email:        "seal@test.com",
				Password:     "pass",
				RegisterDate: time.Now(),
			},
		},
		lists: map[int]model.List{
			1: {
				Id:     1,
				Name:   "List",
				UserId: 1,
			},
		},
		todos: map[int]model.Todo{
			1: {
				Id:          1,
				Title:       "todo1",
				Description: "Todo desc",
				ListId:      1,
				UserId:      1,
			},
		},
	}
}

func (db *dbMock) GetUserModel(id int) (model.UserModel, error) {
	user, err := db.GetUserById(id)
	if err != nil {
		return model.UserModel{}, err
	}

	var todoLists []model.TodoListModel
	if lists, err := db.GetAllListsForUser(user); err == nil {
		for _, l := range lists {
			todos, err := db.GetAllTodosForUserList(user, l)
			if err != nil {
				return model.UserModel{}, err
			}
			todoLists = append(todoLists, model.TodoListModel{
				List:  l,
				Todos: todos,
			})
		}
	} else {
		return model.UserModel{}, err
	}

	result := model.UserModel{
		User:      user,
		TodoLists: todoLists,
	}
	return result, nil
}

func (db *dbMock) GetAllUsers() (model.UsersCollection, error) {
	result := make([]model.User, 0)
	for _, user := range db.users {
		result = append(result, user)
	}
	return result, nil
}

func (db *dbMock) GetUserById(id int) (model.User, error) {
	for _, user := range db.users {
		if user.Id == id {
			return user, nil
		}
	}

	return model.User{}, fmt.Errorf("User not found")
}

func (db *dbMock) GetUserByLogin(login string) (model.User, error) {
	for _, user := range db.users {
		log.Println("u: ", user)
		if user.Login == login || user.Email == login {
			return user, nil
		}
	}

	return model.User{}, fmt.Errorf("User not found")
}

func (db *dbMock) GetUserByLoginAndPassword(login, password string) (model.User, error) {
	for _, user := range db.users {
		if (user.Email == login || user.Login == login) && user.Password == password {
			return user, nil
		}
	}

	return model.User{}, fmt.Errorf("User not found")
}

func (db *dbMock) AddUser(user model.User) (model.User, error) {
	var maxID int
	for id, u := range db.users {
		if id > maxID {
			maxID = id
		}

		if u.Email == user.Email || u.Login == user.Login {
			return user, fmt.Errorf("Email or login alreaedy exist")
		}
	}

	maxID++
	user.Id = maxID
	user.RegisterDate = time.Now()
	db.users[maxID] = user
	return user, nil
}

func (db *dbMock) UpdateUser(user model.User) (int64, error) {
	for i, u := range db.users {
		if u.Id == user.Id {
			db.users[i] = user
			return 1, nil
		}
	}

	return -1, fmt.Errorf("User %d not found", user.Id)
}

func (db *dbMock) DeleteUser(user model.User) (int64, error) {
	for i, u := range db.users {
		if u.Id == user.Id {
			delete(db.users, i)
			return 1, nil
		}
	}

	return -1, fmt.Errorf("User %d not found", user.Id)
}

func (db *dbMock) DeleteUserById(user int) (int64, error) {
	for i, u := range db.users {
		if u.Id == user {
			delete(db.users, i)
			return 1, nil
		}
	}

	return -1, fmt.Errorf("User %d not found", user)
}

func (db *dbMock) GetAllLists() (model.ListsCollection, error) {
	result := make([]model.List, 0)
	for _, list := range db.lists {
		result = append(result, list)
	}
	return result, nil
}

func (db *dbMock) GetAllListsForUser(user model.User) (model.ListsCollection, error) {
	result := make([]model.List, 0)
	for _, list := range db.lists {
		if list.UserId == user.Id {
			result = append(result, list)
		}
	}
	return result, nil
}

func (db *dbMock) GetAllListsForUserId(user int) (model.ListsCollection, error) {
	result := make([]model.List, 0)
	for _, list := range db.lists {
		if list.UserId == user {
			result = append(result, list)
		}
	}
	return result, nil
}

func (db *dbMock) GetListById(id int) (model.List, error) {
	if list, b := db.lists[id]; b {
		return list, nil
	}

	return model.List{}, nil
}

func (db *dbMock) AddList(list model.List) (model.List, error) {
	_, err := db.GetUserById(list.UserId)
	if err != nil {
		return list, err
	}

	var maxID int
	for id := range db.lists {
		if id > maxID {
			maxID = id
		}
	}

	maxID++
	list.Id = maxID
	db.lists[maxID] = list
	return list, nil
}

func (db *dbMock) UpdateList(list model.List) (int64, error) {
	db.lists[list.Id] = list
	return 1, nil
}

func (db *dbMock) DeleteList(list model.List) (int64, error) {
	delete(db.lists, list.Id)
	return 1, nil
}

func (db *dbMock) DeleteListById(list int) (int64, error) {
	delete(db.lists, list)
	return 1, nil
}

func (db *dbMock) GetAllTodos() (model.TodoCollection, error) {
	result := make([]model.Todo, len(db.todos))
	for i, todo := range db.todos {
		result[i] = todo
	}
	return result, nil
}

func (db *dbMock) GetAllTodosForUser(user model.User) (model.TodoCollection, error) {
	result := make([]model.Todo, 0)
	for _, todo := range db.todos {
		if todo.UserId == user.Id {
			result = append(result, todo)
		}
	}
	return result, nil
}

func (db *dbMock) GetAllTodosForUserList(user model.User, list model.List) (model.TodoCollection, error) {
	result := make([]model.Todo, 0)
	for _, todo := range db.todos {
		if todo.UserId == user.Id && todo.ListId == list.Id {
			result = append(result, todo)
		}
	}
	return result, nil
}

func (db *dbMock) AddTodo(todo model.Todo) (model.Todo, error) {
	_, err := db.GetUserById(todo.UserId)
	if err != nil {
		return todo, err
	}

	_, err = db.GetListById(todo.ListId)
	if err != nil {
		return todo, err
	}

	var maxID int
	for id := range db.todos {
		if id > maxID {
			maxID = id
		}
	}

	maxID++
	todo.Id = maxID
	db.todos[maxID] = todo
	return todo, nil
}

func (db *dbMock) UpdateTodo(todo model.Todo) (int64, error) {
	db.todos[todo.Id] = todo
	return 1, nil
}

func (db *dbMock) DeleteTodo(todo model.Todo) (int64, error) {
	delete(db.todos, todo.Id)
	return 1, nil
}

func (db *dbMock) DeleteTodoById(id int) (int64, error) {
	delete(db.todos, id)
	return 1, nil
}
