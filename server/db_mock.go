package server

import "bitbucket.org/SealTV/go-site/model"
import "fmt"
import "time"

type dbMock struct {
	users map[int]model.User
	lists map[int]model.List
	todos map[int]model.Todo
}

func (db *dbMock) GetAllUsers() (model.UsersCollection, error) {
	result := make([]model.User, len(db.users))
	for _, user := range db.users {
		result = append(result, user)
	}
	return result, nil
}

func (db *dbMock) GetUserById(id int) (model.User, error) {
	if user, b := db.users[id]; b {
		return user, nil
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
	var maxId int
	for id, _ := range db.users {
		if id > maxId {
			maxId = id
		}
	}

	maxId++
	user.Id = maxId
	user.RegisterDate = time.Now()
	db.users[maxId] = user
	return user, nil
}

func (db *dbMock) UpdateUser(user model.User) (int64, error) {
	db.users[user.Id] = user
	return 1, nil
}

func (db *dbMock) DeleteUser(user model.User) (int64, error) {
	delete(db.users, user.Id)

	return 1, nil
}

func (db *dbMock) DeleteUserById(user int) (int64, error) {
	delete(db.users, user)

	return 1, nil
}

func (db *dbMock) GetAllLists() (model.ListsCollection, error) {
	result := make([]model.List, len(db.lists))
	for _, list := range db.lists {
		result = append(result, list)
	}
	return result, nil
}

func (db *dbMock) GetAllListsForUser(user model.User) (model.ListsCollection, error) {
	result := make([]model.List, len(db.lists))
	for _, list := range db.lists {
		if list.UserId == user.Id {
			result = append(result, list)
		}
	}
	return result, nil
}

func (db *dbMock) GetAllListsForUserId(user int) (model.ListsCollection, error) {
	result := make([]model.List, len(db.lists))
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
	var maxId int
	for id, _ := range db.lists {
		if id > maxId {
			maxId = id
		}
	}

	maxId++
	list.Id = maxId
	db.lists[maxId] = list
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
	for _, todo := range db.todos {
		result = append(result, todo)
	}
	return result, nil
}

func (db *dbMock) GetAllTodosForUser(user model.User) (model.TodoCollection, error) {
	result := make([]model.Todo, len(db.todos))
	for _, todo := range db.todos {
		if todo.UserId == user.Id {
			result = append(result, todo)
		}
	}
	return result, nil
}

func (db *dbMock) GetAllTodosForUserList(user model.User, list model.List) (model.TodoCollection, error) {
	result := make([]model.Todo, len(db.todos))
	for _, todo := range db.todos {
		if todo.UserId == user.Id && todo.ListId == list.Id {
			result = append(result, todo)
		}
	}
	return result, nil
}

func (db *dbMock) AddTodo(todo model.Todo) (model.Todo, error) {
	var maxId int
	for id, _ := range db.todos {
		if id > maxId {
			maxId = id
		}
	}

	maxId++
	todo.Id = maxId
	db.todos[maxId] = todo
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
