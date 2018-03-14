package data

import (
	"database/sql"
	"fmt"
	"reflect"
	"testing"
	"time"

	"bitbucket.org/SealTV/go-site/backend/model"
	_ "github.com/lib/pq"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetAllTodos(t *testing.T) {
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
		name          string
		fields        fields
		want          model.TodoCollection
		wantErr       bool
		wantErrInRows bool
	}{
		{
			"1",
			fields{db, mock},
			[]model.Todo{
				model.Todo{1, "title", "description", true, time.Now(), 1, 1},
			},
			false,
			false,
		},
		{
			"2",
			fields{db, mock},
			[]model.Todo{
				model.Todo{1, "title", "description", true, time.Now(), 2, 2},
				model.Todo{2, "title1", "description2", true, time.Now(), 2, 2},
			},
			false,
			false,
		},
		{
			"3",
			fields{db, mock},
			[]model.Todo{
				model.Todo{1, "title", "description", true, time.Now(), 2, 2},
				model.Todo{2, "title1", "description2", true, time.Now(), 2, 2},
			},
			true,
			false,
		},
		{
			"4",
			fields{db, mock},
			[]model.Todo{
				model.Todo{1, "title", "description", true, time.Now(), 2, 2},
				model.Todo{2, "title1", "description2", true, time.Now(), 2, 2},
			},
			false,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &pgConnector{
				DB: tt.fields.DB,
			}

			expectQuery := mock.ExpectQuery("SELECT (.+) FROM todos")

			if tt.wantErr {
				expectQuery.WillReturnError(fmt.Errorf("Some error"))

			} else {
				rs := sqlmock.NewRows([]string{"id", "title", "description", "list_id", "is_active", "user_id", "date_create"})

				for _, todo := range tt.want {
					rs = rs.AddRow(todo.Id, todo.Title, todo.Description, todo.ListId, todo.IsActive, todo.UserId, todo.DateCreate)
				}

				if tt.wantErrInRows {
					rs.RowError(1, fmt.Errorf("Some error in raw"))
				}

				expectQuery.WillReturnRows(rs)
			}

			got, err := db.GetAllTodos()

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expections: %s", err)
			}

			if (err != nil) != tt.wantErr && (err != nil) != tt.wantErrInRows {
				t.Errorf("postgresConnector.GetAllTodosForUserList() error = %v, wantErr %v, wantErrInRows %v", err, tt.wantErr, tt.wantErrInRows)
				return
			}

			if !tt.wantErr && !tt.wantErrInRows && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postgresConnector.GetAllTodosForUserList() = %v, want %v, err = %v", got, tt.want, err)
			}
		})
	}
}

func TestGetAllTodosForUser(t *testing.T) {
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
		{
			"third",
			fields{db, mock},
			args{model.User{2, "login", "pass", "email@mail.com", time.Now()}},
			[]model.Todo{
				model.Todo{1, "title", "description", true, time.Now(), 2, 2},
				model.Todo{2, "title1", "description2", true, time.Now(), 2, 2},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &pgConnector{
				DB: tt.fields.DB,
			}

			expectQuery := mock.ExpectQuery("SELECT (.+) FROM todos WHERE (.+)").WithArgs(tt.args.user.Id)
			if tt.wantErr {
				expectQuery.WillReturnError(fmt.Errorf("some error"))
			} else {

				rs := sqlmock.
					NewRows([]string{"id", "title", "description", "list_id", "is_active", "user_id", "date_create"})

				for _, todo := range tt.want {
					rs = rs.AddRow(todo.Id, todo.Title, todo.Description, todo.ListId, todo.IsActive, todo.UserId, todo.DateCreate)
				}
				expectQuery.WillReturnRows(rs)
			}

			got, err := db.GetAllTodosForUser(tt.args.user)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expections: %s", err)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("postgresConnector.GetAllTodosForUserList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postgresConnector.GetAllTodosForUserList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllTodosForUserList(t *testing.T) {
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
		{
			"third",
			fields{db, mock},
			args{model.User{2, "login", "pass", "email@mail.com", time.Now()}, model.List{2, "test_list", 1}},
			[]model.Todo{
				model.Todo{1, "title", "description", true, time.Now(), 2, 2},
				model.Todo{2, "title1", "description2", true, time.Now(), 2, 2},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &pgConnector{
				DB: tt.fields.DB,
			}
			expectQuery := mock.ExpectQuery("SELECT (.+) FROM todos WHERE (.+)").WithArgs(tt.args.user.Id, tt.args.list.Id)

			if tt.wantErr {
				expectQuery.WillReturnError(fmt.Errorf("Some error"))
			} else {

				rs := sqlmock.NewRows([]string{"id", "title", "description", "list_id", "is_active", "user_id", "date_create"})
				for _, todo := range tt.want {
					rs = rs.AddRow(todo.Id, todo.Title, todo.Description, todo.ListId, todo.IsActive, todo.UserId, todo.DateCreate)
				}

				expectQuery.WillReturnRows(rs)
			}

			got, err := db.GetAllTodosForUserList(tt.args.user, tt.args.list)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expections: %s", err)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("postgresConnector.GetAllTodosForUserList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postgresConnector.GetAllTodosForUserList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddTodo(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	connector := new(pgConnector)
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
		{
			"3",
			fields{db, mock},
			model.Todo{0, "title1", "description1", true, time.Time{}, 2, 2},
			model.Todo{1, "title1", "description1", true, time.Now(), 2, 2},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			excpectQuery := mock.ExpectQuery("INSERT INTO todos(.+) RETURNING id, date_create").
				WithArgs(tt.args.Title, tt.args.Description, tt.args.ListId, tt.args.IsActive, tt.args.UserId)

			if tt.wantErr {
				excpectQuery.WillReturnError(fmt.Errorf("Some error"))
			} else {
				rs := sqlmock.NewRows([]string{"id", "date_create"}).AddRow(tt.want.Id, tt.want.DateCreate)
				excpectQuery.WillReturnRows(rs)
			}

			db := &pgConnector{
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
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postgresConnector.AddTodo() = %v, want %v", got, tt)
			}
		})
	}
}

func TestUpdateTodo(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

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
		{
			"3",
			fields{db, mock},
			model.Todo{1, "title11", "description11", true, time.Now(), 2, 2},
			model.Todo{1, "title11", "description11", true, time.Now(), 2, 2},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &pgConnector{
				DB: tt.fields.DB,
			}
			expectExec := mock.ExpectExec(`UPDATE todos`)

			if tt.wantErr {
				expectExec.WillReturnError(fmt.Errorf("Some error 12"))
			} else {
				expectExec.WillReturnResult(sqlmock.NewResult(1, 1))
			}

			res, err := db.UpdateTodo(tt.args)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expections: %s", err)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("postgresConnector.AddTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && res != 1 {
				t.Errorf("expected affected rows to be 1, but got %d instead", res)
			}
		})
	}
}

func TestDeleteTodo(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

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
		{
			"3",
			fields{db, mock},
			model.Todo{1, "title11", "description11", true, time.Now(), 2, 2},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &pgConnector{
				DB: tt.fields.DB,
			}

			expectExec := mock.ExpectExec(`DELETE FROM todos WHERE`)
			if tt.wantErr {
				expectExec.WillReturnError(fmt.Errorf("Some error"))
			} else {
				expectExec.WillReturnResult(sqlmock.NewResult(1, 1))
			}

			res, err := db.DeleteTodo(tt.args)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expections: %s", err)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("postgresConnector.AddTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && res != 1 {
				t.Errorf("expected affected rows to be 1, but got %d instead", res)
			}
		})
	}
}

func TestDeleteTodoById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

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
		{"3", fields{db, mock}, 2, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &pgConnector{
				DB: tt.fields.DB,
			}

			expectExec := mock.ExpectExec(`DELETE FROM todos WHERE`)
			if tt.wantErr {
				expectExec.WillReturnError(fmt.Errorf("some error"))
			} else {
				expectExec.WillReturnResult(sqlmock.NewResult(1, 1))
			}
			res, err := db.DeleteTodoById(tt.args)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expections: %s", err)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("postgresConnector.AddTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && res != 1 {
				t.Errorf("expected affected rows to be 1, but got %d instead", res)
			}
		})
	}
}
