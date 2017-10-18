package data

import (
	"database/sql"
	"reflect"
	"testing"
	"time"

	"bitbucket.org/SealTV/go-site/model"
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/lib/pq"
)

func Test_postgresConnector_GetAllTodos(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	connector := new(postgresConnector)
	connector.DB = db

	type fields struct {
		DB *sql.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    model.TodoCollection
		wantErr bool
	}{
	// TODO: Add test cases.
	}

	mock.ExpectationsWereMet()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &postgresConnector{
				DB: tt.fields.DB,
			}
			got, err := db.GetAllTodos()
			if (err != nil) != tt.wantErr {
				t.Errorf("postgresConnector.GetAllTodos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postgresConnector.GetAllTodos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postgresConnector_GetAllTodosForUser(t *testing.T) {
	type fields struct {
		DB *sql.DB
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
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &postgresConnector{
				DB: tt.fields.DB,
			}
			got, err := db.GetAllTodosForUser(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("postgresConnector.GetAllTodosForUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postgresConnector.GetAllTodosForUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postgresConnector_GetAllTodosForUserList(t *testing.T) {
	type fields struct {
		DB *sql.DB
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
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &postgresConnector{
				DB: tt.fields.DB,
			}
			got, err := db.GetAllTodosForUserList(tt.args.user, tt.args.list)
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
		DB *sql.DB
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
			model.Todo{0, "title", "description", true, time.Time{}, 1, 1,},
			model.Todo{1, "title", "description", true, time.Now(), 1, 1,},
			false,
		},
		{
			"2",
			fields{db, mock},
			model.Todo{0, "title1", "description1", true, time.Time{}, 2, 2,},
			model.Todo{1, "title1", "description1", true, time.Now(), 2, 2,},
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
		DB *sql.DB
		Mock sqlmock.Sqlmock
	}

	tests := []struct {
		name    string
		fields  fields
		args	model.Todo
		want    model.Todo
		wantErr bool
	}{
		{
			"1",
			fields{db, mock},
			model.Todo{1, "title2", "description3", true, time.Time{}, 1, 1,},
			model.Todo{1, "title2", "description3", true, time.Now(), 1, 1,},
			false,
		},
		{
			"2",
			fields{db, mock},
			model.Todo{1, "title11", "description11", true, time.Now(), 2, 2,},
			model.Todo{1, "title11", "description11", true, time.Now(), 2, 2,},
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

			if res != 1{
				t.Errorf("expected affected rows to be 1, but got %d instead", res)
			}
		})
	}
}

func Test_postgresConnector_DeleteTodo(t *testing.T) {
	type fields struct {
		DB *sql.DB
	}
	type args struct {
		todo model.Todo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &postgresConnector{
				DB: tt.fields.DB,
			}
			got, err := db.DeleteTodo(tt.args.todo)
			if (err != nil) != tt.wantErr {
				t.Errorf("postgresConnector.DeleteTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("postgresConnector.DeleteTodo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postgresConnector_DeleteTodoById(t *testing.T) {
	type fields struct {
		DB *sql.DB
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &postgresConnector{
				DB: tt.fields.DB,
			}
			got, err := db.DeleteTodoById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("postgresConnector.DeleteTodoById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("postgresConnector.DeleteTodoById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseTodoRows(t *testing.T) {
	type args struct {
		rows *sql.Rows
	}
	tests := []struct {
		name    string
		args    args
		want    model.TodoCollection
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseTodoRows(tt.args.rows)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseTodoRows() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseTodoRows() = %v, want %v", got, tt.want)
			}
		})
	}
}
