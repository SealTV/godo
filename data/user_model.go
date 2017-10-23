package data

import (
	"bitbucket.org/SealTV/go-site/model"
)

func (db *postgresConnector) GetUserModel(id int) (model.UserModel, error) {
	var userModel model.UserModel
	var user model.User
	user, err := db.GetUserById(id)
	if err != nil {
		return userModel, err
	}
	userModel.User = user
	rows, err := db.Query(`SELECT list.id, list.name, todo.id, todo.title, todo.description, todo.is_active, todo.date_create FROM lists AS list LEFT JOIN todos AS todo ON list.id = todo.list_id WHERE list.user_id = $1`, id)
	if err != nil {
		return userModel, err
	}

	lists := make(map[int]model.TodoListModel)
	for rows.Next() {
		var listID int
		var listName string
		var todo model.Todo
		todo.UserId = user.Id

		err := rows.Scan(&listID, &listName, &todo.Id, &todo.Title, &todo.Description, &todo.IsActive, &todo.DateCreate)
		if err != nil {
			return userModel, err
		}
		todo.ListId = listID

		list, ok := lists[listID]
		if !ok {
			list.List = model.List{
				Id:     listID,
				Name:   listName,
				UserId: user.Id,
			}
		}

		list.Todos = append(list.Todos, todo)
		lists[listID] = list
	}

	for _, v := range lists {
		userModel.TodoLists = append(userModel.TodoLists, v)
	}

	for i := 0; i < len(userModel.TodoLists)-1; i++ {
		for j := i + 1; j < len(userModel.TodoLists); j++ {
			if userModel.TodoLists[i].Id > userModel.TodoLists[j].Id {
				userModel.TodoLists[i], userModel.TodoLists[j] = userModel.TodoLists[j], userModel.TodoLists[i]
			}
		}
	}

	return userModel, nil
}
