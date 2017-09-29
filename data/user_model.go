package data

type UserModel struct {
	User,
	TodoLists []TodoListModel
}
type TodoListModel struct {
	Todo TodoCollection
}
