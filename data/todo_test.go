package data

import (
	"database/sql"
	"reflect"
	"testing"
	"time"

	"bitbucket.org/SealTV/go-site/model"
	_ "github.com/lib/pq"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func Test_postgresConnector_GetAllTodos(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type fields struct {
		DB   *sql.DB
		mock sqlmock.Sqlmock
	}

	tests := []struct {
		name    string
		fields  fields
		want    model.TodoCollection
		wantErr bool
	}{
		{
			"first",
			fields{db, mock},
			[]model.Todo{
				model.Todo{1, "title", "description", true, time.Now(), 1, 1},
			},
			false,
		},
		{
			"second",
			fields{db, mock},
			[]model.Todo{
				model.Todo{1, "title", "description", true, time.Now(), 2, 2},
				model.Todo{2, "title1", "description2", true, time.Now(), 2, 2},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &postgresConnector{
				DB: tt.fields.DB,
			}

			rs := sqlmock.
				NewRows([]string{"id", "title", "description", "list_id", "is_active", "user_id", "date_create"})

			for _, todo := range tt.want {
				rs = rs.AddRow(todo.Id, todo.Title, todo.Description, todo.ListId, todo.IsActive, todo.UserId, todo.DateCreate)
			}

			mock.ExpectQuery("SELECT (.+) FROM todos").WillReturnRows(rs)

			got, err := db.GetAllTodos()

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expections: %s", err)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("postgresConnector.GetAllTodosForUserList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postgresConnector.GetAllTodosForUserList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postgresConnector_GetAllTodosForUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type fields struct {
		DB   *sql.DB
		mock sqlmock.Sqlmock
	}
	type args struct {
		user model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.TodoCollection
		wantErr bool
	}{
		{
			"first",
			fields{db, mock},
			args{model.User{1, "login", "pass", "email@mail.com", time.Now()}},
			[]model.Todo{
				model.Todo{1, "title", "description", true, time.Now(), 1, 1},
			},
			false,
		},
		{
			"second",
			fields{db, mock},
			args{model.User{2, "login", "pass", "email@mail.com", time.Now()}},
			[]model.Todo{
				model.Todo{1, "title", "description", true, time.Now(), 2, 2},
				model.Todo{2, "title1", "description2", true, time.Now(), 2, 2},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &postgresConnector{
				DB: tt.fields.DB,
			}

			rs := sqlmock.
				NewRows([]string{"id", "title", "description", "list_id", "is_active", "user_id", "date_create"})

			for _, todo := range tt.want {
				rs = rs.AddRow(todo.Id, todo.Title, todo.Description, todo.ListId, todo.IsActive, todo.UserId, todo.DateCreate)
			}

			mock.ExpectQuery("SELECT (.+) FROM todos WHERE (.+)").WillReturnRows(rs).WithArgs(tt.args.user.Id)

			got, err := db.GetAllTodosForUser(tt.args.user)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expections: %s", err)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("postgresConnector.GetAllTodosForUserList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postgresConnector.GetAllTodosForUserList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postgresConnector_GetAllTodosForUserList(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type fields struct {
		DB   *sql.DB
		mock sqlmock.Sqlmock
	}
	type args struct {
		user model.User
		list model.List
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.TodoCollection
		wantErr bool
	}{
		{
			"first",
			fields{db, mock},
			args{model.User{1, "login", "pass", "email@mail.com", time.Now()}, model.List{1, "test_list", 1}},
			[]model.Todo{
				model.Todo{1, "title", "description", true, time.Now(), 1, 1},
			},
			false,
		},
		{
			"second",
			fields{db, mock},
			args{model.User{2, "login", "pass", "email@mail.com", time.Now()}, model.List{2, "test_list", 1}},
			[]model.Todo{
				model.Todo{1, "title", "description", true, time.Now(), 2, 2},
				model.Todo{2, "title1", "description2", true, time.Now(), 2, 2},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &postgresConnector{
				DB: tt.fields.DB,
			}

			rs := sqlmock.
				NewRows([]string{"id", "title", "description", "list_id", "is_active", "user_id", "date_create"})

			for _, todo := range tt.want {
				rs = rs.AddRow(todo.Id, todo.Title, todo.Description, todo.ListId, todo.IsActive, todo.UserId, todo.DateCreate)
			}

			mock.ExpectQuery("SELECT (.+) FROM todos WHERE (.+)").WillReturnRows(rs).WithArgs(tt.args.user.Id, tt.args.list.Id)

			got, err := db.GetAllTodosForUserList(tt.args.user, tt.args.list)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expections: %s", err)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("postgresConnector.GetAllTodosForUserList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postgresConnector.GetAllTodosForUserList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postgresConnector_AddTodo(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	connector := new(postgresConnector)
	connector.DB = db

	type fields struct {
		DB   *sql.DB
		Mock sqlmock.Sqlmock
	}

	tests := []struct {
		name    string
		fields  fields
		args    model.Todo
		want    model.Todo
		wantErr bool
	}{
		{
			"1",
			fields{db, mock},
			model.Todo{0, "title", "description", true, time.Time{}, 1, 1},
			model.Todo{1, "title", "description", true, time.Now(), 1, 1},
			false,
		},
		{
			"2",
			fields{db, mock},
			model.Todo{0, "title1", "description1", true, time.Time{}, 2, 2},
			model.Todo{1, "title1", "description1", true, time.Now(), 2, 2},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := sqlmock.NewRows([]string{"id", "date_create"}).AddRow(tt.want.Id, tt.want.DateCreate)
			mock.ExpectQuery("INSERT INTO todos(.+) RETURNING id, date_create").
				WithArgs(tt.args.Title, tt.args.Description, tt.args.ListId, tt.args.IsActive, tt.args.UserId).
				WillReturnRows(rs)

			db := &postgresConnector{
				DB: tt.fields.DB,
			}
			got, err := db.AddTodo(tt.args)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expections: %s", err)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("postgresConnector.AddTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postgresConnector.AddTodo() = %v, want %v", got, tt)
			}
		})
	}
}

func Test_postgresConnector_UpdateTodo(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	connector := new(postgresConnector)
	connector.DB = db

	type fields struct {
		DB   *sql.DB
		Mock sqlmock.Sqlmock
	}

	tests := []struct {
		name    string
		fields  fields
		args    model.Todo
		want    model.Todo
		wantErr bool
	}{
		{
			"1",
			fields{db, mock},
			model.Todo{1, "title2", "description3", true, time.Time{}, 1, 1},
			model.Todo{1, "title2", "description3", true, time.Now(), 1, 1},
			false,
		},
		{
			"2",
			fields{db, mock},
			model.Todo{1, "title11", "description11", true, time.Now(), 2, 2},
			model.Todo{1, "title11", "description11", true, time.Now(), 2, 2},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &postgresConnector{
				DB: tt.fields.DB,
			}

			mock.ExpectExec(`UPDATE todos`).
				WillReturnResult(sqlmock.NewResult(1, 1))

			res, err := db.UpdateTodo(tt.args)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expections: %s", err)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("postgresConnector.AddTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				t.Errorf("error '%s' was not expected, while inserting a row", err)
			}

			if res != 1 {
				t.Errorf("expected affected rows to be 1, but got %d instead", res)
			}
		})
	}
}

func Test_postgresConnector_DeleteTodo(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	connector := new(postgresConnector)
	connector.DB = db

	type fields struct {
		DB   *sql.DB
		Mock sqlmock.Sqlmock
	}

	tests := []struct {
		name    string
		fields  fields
		args    model.Todo
		wantErr bool
	}{
		{
			"1",
			fields{db, mock},
			model.Todo{1, "title2", "description3", true, time.Time{}, 1, 1},
			false,
		},
		{
			"2",
			fields{db, mock},
			model.Todo{1, "title11", "description11", true, time.Now(), 2, 2},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &postgresConnector{
				DB: tt.fields.DB,
			}

			mock.ExpectExec(`DELETE FROM todos WHERE`).
				WillReturnResult(sqlmock.NewResult(1, 1))

			res, err := db.DeleteTodo(tt.args)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expections: %s", err)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("postgresConnector.AddTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				t.Errorf("error '%s' was not expected, while inserting a row", err)
			}

			if res != 1 {
				t.Errorf("expected affected rows to be 1, but got %d instead", res)
			}
		})
	}
}

func Test_postgresConnector_DeleteTodoById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	connector := new(postgresConnector)
	connector.DB = db

	type fields struct {
		DB   *sql.DB
		Mock sqlmock.Sqlmock
	}

	tests := []struct {
		name    string
		fields  fields
		args    int
		wantErr bool
	}{
		{"1", fields{db, mock}, 1, false},
		{"2", fields{db, mock}, 2, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &postgresConnector{
				DB: tt.fields.DB,
			}

			mock.ExpectExec(`DELETE FROM todos WHERE`).
				WillReturnResult(sqlmock.NewResult(1, 1))

			res, err := db.DeleteTodoById(tt.args)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expections: %s", err)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("postgresConnector.AddTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				t.Errorf("error '%s' was not expected, while inserting a row", err)
			}

			if res != 1 {
				t.Errorf("expected affected rows to be 1, but got %d instead", res)
			}
		})
	}
}
