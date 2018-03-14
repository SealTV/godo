package data

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"bitbucket.org/SealTV/go-site/backend/model"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetAllUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	tests := []struct {
		name    string
		db      *pgConnector
		mock    sqlmock.Sqlmock
		want    model.UsersCollection
		wantErr bool
	}{
		{name: "1",
			db:   &pgConnector{db},
			mock: mock,
			want: []model.User{
				model.User{Id: 1, Email: "some1@email.com", Login: "SomeLogin1", Password: "Some pass", RegisterDate: time.Now()},
				model.User{Id: 2, Email: "some2@email.com", Login: "SomeLogin2", Password: "Some pass", RegisterDate: time.Now()},
			},
			wantErr: false,
		},
		{name: "2",
			db:   &pgConnector{db},
			mock: mock,
			want: []model.User{
				model.User{Id: 1, Email: "some1@email.com", Login: "SomeLogin1", Password: "Some pass", RegisterDate: time.Now()},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectQuery := tt.mock.ExpectQuery("SELECT (.+) FROM users")
			if tt.wantErr {
				expectQuery.WillReturnError(fmt.Errorf("Some error"))
			} else {
				rs := sqlmock.NewRows([]string{"id", "login", "password", "email", "register_date"})
				for _, user := range tt.want {
					rs = rs.AddRow(user.Id, user.Login, user.Password, user.Email, user.RegisterDate)
				}

				expectQuery.WillReturnRows(rs)
			}

			got, err := tt.db.GetAllUsers()

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expections: %s", err)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("postgresConnector.GetAllUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postgresConnector.GetAllUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUserById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type args struct {
		id int
	}
	tests := []struct {
		name    string
		db      *pgConnector
		mock    sqlmock.Sqlmock
		args    args
		want    model.User
		wantErr bool
	}{
		{name: "1",
			db:      &pgConnector{db},
			mock:    mock,
			args:    args{id: 1},
			want:    model.User{Id: 1, Email: "some1@email.com", Login: "SomeLogin1", Password: "Some pass", RegisterDate: time.Now()},
			wantErr: false,
		},
		{name: "2",
			db:      &pgConnector{db},
			mock:    mock,
			args:    args{id: 2},
			want:    model.User{Id: 2, Email: "some1@email.com", Login: "SomeLogin1", Password: "Some pass", RegisterDate: time.Now()},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectQuery := tt.mock.ExpectQuery("SELECT (.+) FROM users WHERE id = (.+) LIMIT 1").WithArgs(tt.args.id)
			if tt.wantErr {
				expectQuery.WillReturnError(fmt.Errorf("Some error"))
			} else {
				rs := sqlmock.NewRows([]string{"id", "login", "password", "email", "register_date"}).
					AddRow(tt.want.Id, tt.want.Login, tt.want.Password, tt.want.Email, tt.want.RegisterDate)
				expectQuery.WillReturnRows(rs)
			}

			got, err := tt.db.GetUserById(tt.args.id)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expections: %s", err)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("postgresConnector.GetUserById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postgresConnector.GetUserById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUserByLoginAndPassword(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type args struct {
		login    string
		password string
	}
	tests := []struct {
		name    string
		db      *pgConnector
		mock    sqlmock.Sqlmock
		args    args
		want    model.User
		wantErr bool
	}{
		{
			name:    "1",
			db:      &pgConnector{db},
			mock:    mock,
			args:    args{login: "SomeLogin1", password: "Some pass"},
			want:    model.User{Id: 1, Email: "some1@email.com", Login: "SomeLogin1", Password: "Some pass", RegisterDate: time.Now()},
			wantErr: false,
		},
		{
			name:    "2",
			db:      &pgConnector{db},
			mock:    mock,
			args:    args{login: "some1@email.com", password: "Some pass"},
			want:    model.User{Id: 1, Email: "some1@email.com", Login: "SomeLogin1", Password: "Some pass", RegisterDate: time.Now()},
			wantErr: false,
		},
		{
			name:    "3",
			db:      &pgConnector{db},
			mock:    mock,
			args:    args{login: "SomeLogin1", password: "Some pass"},
			want:    model.User{Id: 2, Email: "some1@email.com", Login: "SomeLogin1", Password: "Some pass", RegisterDate: time.Now()},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectQuery := tt.mock.ExpectQuery("SELECT (.+) FROM users WHERE ((.+)) AND password = (.+) LIMIT 1").
				WithArgs(tt.args.login, tt.args.password)

			if tt.wantErr {
				expectQuery.WillReturnError(fmt.Errorf("Some error"))
			} else {
				rs := sqlmock.NewRows([]string{"id", "login", "password", "email", "register_date"}).
					AddRow(tt.want.Id, tt.want.Login, tt.want.Password, tt.want.Email, tt.want.RegisterDate)
				expectQuery.WillReturnRows(rs)
			}

			got, err := tt.db.GetUserByLoginAndPassword(tt.args.login, tt.args.password)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expections: %s", err)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("postgresConnector.GetUserByLoginAndPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postgresConnector.GetUserByLoginAndPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type args struct {
		user model.User
	}
	tests := []struct {
		name    string
		db      *pgConnector
		mock    sqlmock.Sqlmock
		args    args
		want    model.User
		wantErr bool
	}{
		{
			name:    "1",
			db:      &pgConnector{db},
			mock:    mock,
			args:    args{user: model.User{Email: "some1@email.com", Login: "SomeLogin1", Password: "Some pass"}},
			want:    model.User{Id: 1, Email: "some1@email.com", Login: "SomeLogin1", Password: "Some pass", RegisterDate: time.Now()},
			wantErr: false,
		},
		{
			name:    "2",
			db:      &pgConnector{db},
			mock:    mock,
			args:    args{user: model.User{Email: "some1@email.com", Login: "SomeLogin1", Password: "Some pass"}},
			want:    model.User{Id: 1, Email: "some1@email.com", Login: "SomeLogin1", Password: "Some pass", RegisterDate: time.Now()},
			wantErr: false,
		},
		{
			name:    "3",
			db:      &pgConnector{db},
			mock:    mock,
			args:    args{user: model.User{Email: "some1@email.com", Login: "SomeLogin1", Password: "Some pass"}},
			want:    model.User{Id: 2, Email: "some1@email.com", Login: "SomeLogin1", Password: "Some pass", RegisterDate: time.Now()},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectQuery := tt.mock.ExpectQuery(`INSERT INTO users((.+)) VALUES ((.+)) RETURNING (.+)`).
				WithArgs(tt.args.user.Login, tt.args.user.Password, tt.args.user.Email)

			if tt.wantErr {
				expectQuery.WillReturnError(fmt.Errorf("Some error"))
			} else {
				rs := sqlmock.NewRows([]string{"id", "register_date"}).AddRow(tt.want.Id, tt.want.RegisterDate)
				expectQuery.WillReturnRows(rs)
			}

			got, err := tt.db.AddUser(tt.args.user)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expections: %s", err)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("postgresConnector.AddUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postgresConnector.AddUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type args struct {
		user model.User
	}
	tests := []struct {
		name    string
		db      *pgConnector
		mock    sqlmock.Sqlmock
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:    "1",
			db:      &pgConnector{db},
			mock:    mock,
			args:    args{user: model.User{Id: 1, Email: "some1@email.com", Login: "SomeLogin1", Password: "Some pass", RegisterDate: time.Now()}},
			want:    1,
			wantErr: false,
		},
		{
			name:    "2",
			db:      &pgConnector{db},
			mock:    mock,
			args:    args{user: model.User{Id: 1, Email: "some1@email.com", Login: "SomeLogin1", Password: "Some pass", RegisterDate: time.Now()}},
			want:    1,
			wantErr: false,
		},
		{
			name:    "3",
			db:      &pgConnector{db},
			mock:    mock,
			args:    args{user: model.User{Id: 2, Email: "some1@email.com", Login: "SomeLogin1", Password: "Some pass", RegisterDate: time.Now()}},
			want:    1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectExec := tt.mock.ExpectExec(`UPDATE users SET (.+) WHERE (.+)`).
				WithArgs(tt.args.user.Id, tt.args.user.Login, tt.args.user.Password, tt.args.user.Email)

			if tt.wantErr {
				expectExec.WillReturnError(fmt.Errorf("Some error"))
			} else {
				expectExec.WillReturnResult(sqlmock.NewResult(tt.want, tt.want))
			}

			got, err := tt.db.UpdateUser(tt.args.user)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expections: %s", err)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("postgresConnector.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("postgresConnector.UpdateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type args struct {
		user model.User
	}
	tests := []struct {
		name     string
		db       *pgConnector
		mock     sqlmock.Sqlmock
		args     args
		want     int64
		wantErr1 bool
		wantErr2 bool
		wantErr3 bool
		wantErr4 bool
	}{
		{
			name:     "1",
			db:       &pgConnector{db},
			mock:     mock,
			args:     args{user: model.User{Id: 1, Email: "some1@email.com", Login: "SomeLogin1", Password: "Some pass", RegisterDate: time.Now()}},
			want:     1,
			wantErr1: true,
			wantErr2: false,
			wantErr3: false,
			wantErr4: false,
		},
		{
			name:     "2",
			db:       &pgConnector{db},
			mock:     mock,
			args:     args{user: model.User{Id: 1, Email: "some1@email.com", Login: "SomeLogin1", Password: "Some pass", RegisterDate: time.Now()}},
			want:     1,
			wantErr1: false,
			wantErr2: true,
			wantErr3: false,
			wantErr4: false,
		},
		{
			name:     "3",
			db:       &pgConnector{db},
			mock:     mock,
			args:     args{user: model.User{Id: 2, Email: "some1@email.com", Login: "SomeLogin1", Password: "Some pass", RegisterDate: time.Now()}},
			want:     1,
			wantErr1: false,
			wantErr2: false,
			wantErr3: true,
			wantErr4: false,
		},
		{
			name:     "4",
			db:       &pgConnector{db},
			mock:     mock,
			args:     args{user: model.User{Id: 2, Email: "some1@email.com", Login: "SomeLogin1", Password: "Some pass", RegisterDate: time.Now()}},
			want:     1,
			wantErr1: false,
			wantErr2: false,
			wantErr3: false,
			wantErr4: true,
		},
		{
			name:     "5",
			db:       &pgConnector{db},
			mock:     mock,
			args:     args{user: model.User{Id: 1, Email: "some1@email.com", Login: "SomeLogin1", Password: "Some pass", RegisterDate: time.Now()}},
			want:     1,
			wantErr1: false,
			wantErr2: false,
			wantErr3: false,
			wantErr4: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock.ExpectBegin()

			expectExec := tt.mock.ExpectExec(`DELETE FROM todos WHERE (.+)`).WithArgs(tt.args.user.Id).WillReturnResult(sqlmock.NewResult(tt.want, tt.want))
			if tt.wantErr1 {
				expectExec.WillReturnError(fmt.Errorf("Some error"))
				tt.mock.ExpectRollback()
			} else {
				expectExec := tt.mock.ExpectExec("DELETE FROM lists WHERE (.+)").WithArgs(tt.args.user.Id).WillReturnResult(sqlmock.NewResult(tt.want, tt.want))
				if tt.wantErr2 {
					expectExec.WillReturnError(fmt.Errorf("Some error"))
					tt.mock.ExpectRollback()
				} else {
					expectExec := tt.mock.ExpectExec(`DELETE FROM users WHERE (.+)`).WithArgs(tt.args.user.Id).WillReturnResult(sqlmock.NewResult(tt.want, tt.want))
					if tt.wantErr3 {
						expectExec.WillReturnError(fmt.Errorf("Some error"))
						tt.mock.ExpectRollback()
					} else {
						expectExec.WillReturnResult(sqlmock.NewResult(tt.want, tt.want))
						expectCommit := tt.mock.ExpectCommit()
						if tt.wantErr4 {
							expectCommit.WillReturnError(fmt.Errorf("some error"))
						}
					}
				}
			}

			got, err := tt.db.DeleteUser(tt.args.user)
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expections: %s", err)
			}

			wantErr := tt.wantErr1 || tt.wantErr2 || tt.wantErr3 || tt.wantErr4
			if (err != nil) != wantErr {
				t.Errorf("postgresConnector.DeleteUser() error = %v, wantErr %v", err, wantErr)
				return
			}

			if !wantErr && got != tt.want {
				t.Errorf("postgresConnector.DeleteUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
