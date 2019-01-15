package data

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"bitbucket.org/SealTV/go-site/backend/model"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetUserModel(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type args struct {
		id int
	}
	tests := []struct {
		name     string
		db       *pgConnector
		mock     sqlmock.Sqlmock
		args     args
		want     model.UserModel
		wantErr1 bool
		wantErr2 bool
	}{
		{
			name: "1",
			db:   &pgConnector{db},
			mock: mock,
			args: args{id: 1},
			want: model.UserModel{
				User: model.User{ID: 1, Email: "some1@email.com", Login: "SomeLogin1", Password: "Some pass", RegisterDate: time.Now()},
				TodoLists: []model.TodoListModel{
					model.TodoListModel{
						List: model.List{ID: 1, Name: "Some name 1", UserID: 1},
						Todos: []model.Todo{
							model.Todo{
								ID:          1,
								Title:       "title",
								Description: "description",
								IsActive:    true,
								DateCreate:  time.Time{},
								UserID:      1,
								ListID:      1,
							},
							model.Todo{
								ID:          2,
								Title:       "title1",
								Description: "description1",
								IsActive:    true,
								DateCreate:  time.Time{},
								UserID:      1,
								ListID:      1,
							},
						},
					},
					model.TodoListModel{
						List: model.List{ID: 2, Name: "Some name 2", UserID: 1},
						Todos: []model.Todo{
							model.Todo{
								ID:          3,
								Title:       "title",
								Description: "description",
								IsActive:    true,
								DateCreate:  time.Time{},
								UserID:      1,
								ListID:      2,
							},
							model.Todo{
								ID:          4,
								Title:       "title1",
								Description: "description1",
								IsActive:    true,
								DateCreate:  time.Time{},
								UserID:      1,
								ListID:      2,
							},
						},
					},
				},
			},
			wantErr1: false,
			wantErr2: false,
		},
		{
			name: "2",
			db:   &pgConnector{db},
			mock: mock,
			args: args{id: 1},
			want: model.UserModel{
				User: model.User{ID: 1, Email: "some1@email.com", Login: "SomeLogin1", Password: "Some pass", RegisterDate: time.Now()},
				TodoLists: []model.TodoListModel{
					model.TodoListModel{
						List: model.List{ID: 1, Name: "Some name 1", UserID: 1},
						Todos: []model.Todo{
							model.Todo{
								ID:          1,
								Title:       "title",
								Description: "description",
								IsActive:    true,
								DateCreate:  time.Time{},
								UserID:      1,
								ListID:      1,
							},
							model.Todo{
								ID:          2,
								Title:       "title1",
								Description: "description1",
								IsActive:    true,
								DateCreate:  time.Time{},
								UserID:      1,
								ListID:      1,
							},
						},
					},
					model.TodoListModel{
						List: model.List{ID: 2, Name: "Some name 2", UserID: 1},
						Todos: []model.Todo{
							model.Todo{
								ID:          3,
								Title:       "title",
								Description: "description",
								IsActive:    true,
								DateCreate:  time.Time{},
								UserID:      1,
								ListID:      2,
							},
							model.Todo{
								ID:          4,
								Title:       "title1",
								Description: "description1",
								IsActive:    true,
								DateCreate:  time.Time{},
								UserID:      1,
								ListID:      2,
							},
						},
					},
				},
			},
			wantErr1: true,
			wantErr2: false,
		},
		{
			name: "3",
			db:   &pgConnector{db},
			mock: mock,
			args: args{id: 1},
			want: model.UserModel{
				User: model.User{ID: 1, Email: "some1@email.com", Login: "SomeLogin1", Password: "Some pass", RegisterDate: time.Now()},
				TodoLists: []model.TodoListModel{
					model.TodoListModel{
						List: model.List{ID: 1, Name: "Some name 1", UserID: 1},
						Todos: []model.Todo{
							model.Todo{
								ID:          1,
								Title:       "title",
								Description: "description",
								IsActive:    true,
								DateCreate:  time.Time{},
								UserID:      1,
								ListID:      1,
							},
							model.Todo{
								ID:          2,
								Title:       "title1",
								Description: "description1",
								IsActive:    true,
								DateCreate:  time.Time{},
								UserID:      1,
								ListID:      1,
							},
						},
					},
					model.TodoListModel{
						List: model.List{ID: 2, Name: "Some name 2", UserID: 1},
						Todos: []model.Todo{
							model.Todo{
								ID:          3,
								Title:       "title",
								Description: "description",
								IsActive:    true,
								DateCreate:  time.Time{},
								UserID:      1,
								ListID:      2,
							},
							model.Todo{
								ID:          4,
								Title:       "title1",
								Description: "description1",
								IsActive:    true,
								DateCreate:  time.Time{},
								UserID:      1,
								ListID:      2,
							},
						},
					},
				},
			},
			wantErr1: false,
			wantErr2: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectQuery := tt.mock.ExpectQuery("SELECT (.+) FROM users WHERE id = (.+) LIMIT 1")
			if tt.wantErr1 {
				expectQuery.WillReturnError(fmt.Errorf("Some error"))
			} else {
				rs := sqlmock.NewRows([]string{"id", "login", "password", "email", "register_date"}).
					AddRow(tt.want.ID, tt.want.Login, tt.want.Password, tt.want.Email, tt.want.RegisterDate)
				expectQuery.WillReturnRows(rs)

				expectQuery := tt.mock.ExpectQuery(`SELECT (.+) FROM lists AS list LEFT JOIN todos AS todo ON (.+) WHERE (.+)`).WithArgs(tt.args.id)
				if tt.wantErr2 {
					expectQuery.WillReturnError(fmt.Errorf("Some error"))
				} else {
					rs := sqlmock.NewRows([]string{"list.id", "list.name", "todo.id", "todo.title", "todo.description", "todo.is_active", "todo.date_create"})
					for _, list := range tt.want.TodoLists {
						for _, todo := range list.Todos {
							rs = rs.AddRow(list.ID, list.Name, todo.ID, todo.Title, todo.Description, todo.IsActive, todo.DateCreate)
						}
					}

					expectQuery.WillReturnRows(rs)
				}
			}

			got, err := tt.db.GetUserModel(tt.args.id)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expections: %s", err)
			}

			wantErr := tt.wantErr1 || tt.wantErr2
			if (err != nil) != wantErr {
				t.Errorf("postgresConnector.GetUserModel() error = %v, wantErr %v", err, wantErr)
				return
			}
			if !wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postgresConnector.GetUserModel() = %v, want %v", got, tt.want)
			}
		})
	}
}
