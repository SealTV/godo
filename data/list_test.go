package data

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"bitbucket.org/SealTV/go-site/model"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetAllLists(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	tests := []struct {
		name    string
		db      *postgresConnector
		mock    sqlmock.Sqlmock
		want    model.ListsCollection
		wantErr bool
	}{
		{name: "1",
			db:   &postgresConnector{db},
			mock: mock,
			want: []model.List{
				model.List{Id: 1, Name: "Some name 1", UserId: 1},
				model.List{Id: 2, Name: "Some name 2", UserId: 1},
			},
			wantErr: false,
		},
		{name: "2",
			db:   &postgresConnector{db},
			mock: mock,
			want: []model.List{
				model.List{Id: 1, Name: "Some name 1", UserId: 1},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectQuery := tt.mock.ExpectQuery("SELECT (.+) FROM lists")
			if tt.wantErr {
				expectQuery.WillReturnError(fmt.Errorf("Some error"))
			} else {
				rs := sqlmock.NewRows([]string{"id", "name", "user_id"})
				for _, list := range tt.want {
					rs = rs.AddRow(list.Id, list.Name, list.UserId)
				}

				expectQuery.WillReturnRows(rs)
			}

			got, err := tt.db.GetAllLists()
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expections: %s", err)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("postgresConnector.GetAllLists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postgresConnector.GetAllLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllListsForUser(t *testing.T) {
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
		db      *postgresConnector
		mock    sqlmock.Sqlmock
		args    args
		want    model.ListsCollection
		wantErr bool
	}{
		{name: "1",
			db:   &postgresConnector{db},
			mock: mock,
			args: args{user: model.User{Id: 1, Email: "some1@email.com", Login: "SomeLogin1", Password: "Some pass", RegisterDate: time.Now()}},
			want: []model.List{
				model.List{Id: 1, Name: "Some name 1", UserId: 1},
				model.List{Id: 2, Name: "Some name 2", UserId: 1},
			},
			wantErr: false,
		},
		{name: "2",
			db:   &postgresConnector{db},
			mock: mock,
			args: args{user: model.User{Id: 1, Email: "some1@email.com", Login: "SomeLogin1", Password: "Some pass", RegisterDate: time.Now()}},
			want: []model.List{
				model.List{Id: 1, Name: "Some name 1", UserId: 1},
				model.List{Id: 2, Name: "Some name 2", UserId: 1},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectQuery := tt.mock.ExpectQuery("SELECT  (.+) FROM lists WHERE (.+)").WithArgs(tt.args.user.Id)
			if tt.wantErr {
				expectQuery.WillReturnError(fmt.Errorf("Some error"))
			} else {
				rs := sqlmock.NewRows([]string{"id", "name", "user_id"})
				for _, list := range tt.want {
					rs = rs.AddRow(list.Id, list.Name, list.UserId)
				}

				expectQuery.WillReturnRows(rs)
			}

			got, err := tt.db.GetAllListsForUser(tt.args.user)
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expections: %s", err)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("postgresConnector.GetAllListsForUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postgresConnector.GetAllListsForUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllListsForUserId(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type args struct {
		user int
	}
	tests := []struct {
		name    string
		db      *postgresConnector
		mock    sqlmock.Sqlmock
		args    args
		want    model.ListsCollection
		wantErr bool
	}{
		{name: "1",
			db:   &postgresConnector{db},
			mock: mock,
			args: args{user: 1},
			want: []model.List{
				model.List{Id: 1, Name: "Some name 1", UserId: 1},
				model.List{Id: 2, Name: "Some name 2", UserId: 1},
			},
			wantErr: false,
		},
		{name: "2",
			db:   &postgresConnector{db},
			mock: mock,
			args: args{user: 1},
			want: []model.List{
				model.List{Id: 1, Name: "Some name 1", UserId: 1},
				model.List{Id: 2, Name: "Some name 2", UserId: 1},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectQuery := tt.mock.ExpectQuery("SELECT  (.+) FROM lists WHERE (.+)").WithArgs(tt.args.user)
			if tt.wantErr {
				expectQuery.WillReturnError(fmt.Errorf("Some error"))
			} else {
				rs := sqlmock.NewRows([]string{"id", "name", "user_id"})
				for _, list := range tt.want {
					rs = rs.AddRow(list.Id, list.Name, list.UserId)
				}

				expectQuery.WillReturnRows(rs)
			}

			got, err := tt.db.GetAllListsForUserId(tt.args.user)
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expections: %s", err)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("postgresConnector.GetAllListsForUserId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postgresConnector.GetAllListsForUserId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetListById(t *testing.T) {
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
		db      *postgresConnector
		mock    sqlmock.Sqlmock
		args    args
		want    model.List
		wantErr bool
	}{
		{name: "1",
			db:      &postgresConnector{db},
			mock:    mock,
			args:    args{id: 1},
			want:    model.List{Id: 1, Name: "Some name 1", UserId: 1},
			wantErr: false,
		},
		{name: "2",
			db:      &postgresConnector{db},
			mock:    mock,
			args:    args{id: 1},
			want:    model.List{Id: 1, Name: "Some name 1", UserId: 1},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectQuery := tt.mock.ExpectQuery("SELECT (.+) FROM lists WHERE").WithArgs(tt.args.id)
			if tt.wantErr {
				expectQuery.WillReturnError(fmt.Errorf("Some error"))
			} else {
				rs := sqlmock.NewRows([]string{"id", "name", "user_id"}).AddRow(tt.want.Id, tt.want.Name, tt.want.UserId)

				expectQuery.WillReturnRows(rs)
			}

			got, err := tt.db.GetListById(tt.args.id)
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expections: %s", err)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("postgresConnector.GetListById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postgresConnector.GetListById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddList(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type args struct {
		list model.List
	}
	tests := []struct {
		name    string
		db      *postgresConnector
		mock    sqlmock.Sqlmock
		args    args
		want    model.List
		wantErr bool
	}{
		{
			name:    "1",
			db:      &postgresConnector{db},
			mock:    mock,
			args:    args{list: model.List{Id: 1, Name: "Some list", UserId: 1}},
			want:    model.List{Id: 1, Name: "Some list", UserId: 1},
			wantErr: false,
		},
		{
			name:    "2",
			db:      &postgresConnector{db},
			mock:    mock,
			args:    args{list: model.List{Id: 1, Name: "Some list", UserId: 1}},
			want:    model.List{Id: 1, Name: "Some list", UserId: 1},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectQuery := tt.mock.ExpectQuery("INSERT INTO lists((.+)) VALUES((.+)) RETURNING (.+)").
				WithArgs(tt.args.list.Name, tt.args.list.UserId)

			if tt.wantErr {
				expectQuery.WillReturnError(fmt.Errorf("Some error"))
			} else {
				rs := sqlmock.NewRows([]string{"id"}).AddRow(tt.want.Id)
				expectQuery.WillReturnRows(rs)
			}

			got, err := tt.db.AddList(tt.args.list)
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expections: %s", err)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("postgresConnector.AddList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postgresConnector.AddList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateList(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type args struct {
		list model.List
	}
	tests := []struct {
		name    string
		db      *postgresConnector
		mock    sqlmock.Sqlmock
		args    args
		want    int64
		wantErr bool
	}{
		{name: "1",
			db:      &postgresConnector{db},
			mock:    mock,
			args:    args{list: model.List{Id: 1, Name: "Some list", UserId: 1}},
			want:    1,
			wantErr: false,
		},
		{name: "2",
			db:      &postgresConnector{db},
			mock:    mock,
			args:    args{list: model.List{Id: 1, Name: "Some list", UserId: 1}},
			want:    1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectExec := tt.mock.ExpectExec("UPDATE lists SET (.+) WHERE (.+)").WithArgs(tt.args.list.Id, tt.args.list.Name, tt.args.list.UserId)
			if tt.wantErr {
				expectExec.WillReturnError(fmt.Errorf("Some error"))
			} else {
				expectExec.WillReturnResult(sqlmock.NewResult(tt.want, tt.want))
			}

			got, err := tt.db.UpdateList(tt.args.list)
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expections: %s", err)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("postgresConnector.UpdateList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("postgresConnector.UpdateList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteList(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type args struct {
		list model.List
	}
	tests := []struct {
		name     string
		db       *postgresConnector
		mock     sqlmock.Sqlmock
		args     args
		want     int64
		wantErr1 bool
		wantErr2 bool
		wantErr3 bool
	}{
		{name: "1",
			db:       &postgresConnector{db},
			mock:     mock,
			args:     args{list: model.List{Id: 1, Name: "Some name 1", UserId: 1}},
			want:     1,
			wantErr1: false,
			wantErr2: false,
			wantErr3: false,
		},
		{name: "2",
			db:       &postgresConnector{db},
			mock:     mock,
			args:     args{list: model.List{Id: 1, Name: "Some name 1", UserId: 1}},
			want:     1,
			wantErr1: true,
			wantErr2: false,
			wantErr3: false,
		},
		{name: "3",
			db:       &postgresConnector{db},
			mock:     mock,
			args:     args{list: model.List{Id: 1, Name: "Some name 1", UserId: 1}},
			want:     1,
			wantErr1: false,
			wantErr2: true,
			wantErr3: false,
		},
		{name: "4",
			db:       &postgresConnector{db},
			mock:     mock,
			args:     args{list: model.List{Id: 1, Name: "Some name 1", UserId: 1}},
			want:     1,
			wantErr1: false,
			wantErr2: false,
			wantErr3: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock.ExpectBegin()
			expectExec := tt.mock.ExpectExec(`DELETE FROM todos WHERE (.+)`).WithArgs(tt.args.list.Id).WillReturnResult(sqlmock.NewResult(tt.want, tt.want))
			if tt.wantErr1 {
				expectExec.WillReturnError(fmt.Errorf("Some error"))
				tt.mock.ExpectRollback()
			} else {
				expectExec = tt.mock.ExpectExec("DELETE FROM lists WHERE (.+)").WithArgs(tt.args.list.Id)
				if tt.wantErr2 {
					expectExec.WillReturnError(fmt.Errorf("Some error"))
					tt.mock.ExpectRollback()
				} else {
					expectExec.WillReturnResult(sqlmock.NewResult(tt.want, tt.want))
					expectCommit := tt.mock.ExpectCommit()
					if tt.wantErr3 {
						expectCommit.WillReturnError(fmt.Errorf("some error"))
					}
				}
			}

			got, err := tt.db.DeleteList(tt.args.list)
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expections: %s", err)
			}

			wantErr := tt.wantErr1 || tt.wantErr2 || tt.wantErr3
			if (err != nil) != wantErr {
				t.Errorf("postgresConnector.DeleteList() error = %v, wantErr %v", err, wantErr)
				return
			}

			if !wantErr && got != tt.want {
				t.Errorf("postgresConnector.DeleteList() = %v, want %v", got, tt.want)
			}
		})
	}
}
