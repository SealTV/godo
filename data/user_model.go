package data

func (db *PostgresConnector) GetUserModel(id int) (UserModel, error) {
	var userModel UserModel
	var user User
	user, err := db.GetUserById(id)
	if err != nil {
		return userModel, err
	}
	userModel.User = user
	rows, err := db.Query(`SELECT list.id, list.name,
		todo.id, todo.title, todo.description, todo.is_active, todo.date_create
		FROM lists AS list
		LEFT JOIN todos AS todo ON list.id = todo.list_id
		WHERE list.user_id = $1`, id)
	if err != nil {
		return userModel, err
	}

	lists := make(map[int]TodoListModel)
	for rows.Next() {
		var listId int
		var listName string
		var todo Todo
		todo.UserId = user.Id

		err := rows.Scan(&listId, &listName, &todo.Id, &todo.Title, &todo.Description, &todo.IsActive, &todo.DateCreate)
		if err != nil {
			return userModel, err
		}
		todo.ListId = listId

		list, ok := lists[listId]
		if !ok {
			list.List = List{
				Id:     listId,
				Name:   listName,
				UserId: user.Id,
			}
		}

		list.Todos = append(list.Todos, todo)
		lists[listId] = list
	}

	for _, v := range lists {
		userModel.TodoLists = append(userModel.TodoLists, v)
	}

	return userModel, nil
}
