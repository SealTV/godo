package data

import (
	"database/sql"
	"reflect"
	"testing"

	"bitbucket.org/SealTV/go-site/model"
)

func Test_postgresConnector_GetAllUsers(t *testing.T) {
	tests := []struct {
		name    string
		db      *postgresConnector
		want    model.UsersCollection
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.GetAllUsers()
			if (err != nil) != tt.wantErr {
				t.Errorf("postgresConnector.GetAllUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postgresConnector.GetAllUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postgresConnector_GetUserById(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		db      *postgresConnector
		args    args
		want    model.User
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.GetUserById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("postgresConnector.GetUserById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postgresConnector.GetUserById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postgresConnector_GetUserByLoginAndPassword(t *testing.T) {
	type args struct {
		login    string
		password string
	}
	tests := []struct {
		name    string
		db      *postgresConnector
		args    args
		want    model.User
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.GetUserByLoginAndPassword(tt.args.login, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("postgresConnector.GetUserByLoginAndPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postgresConnector.GetUserByLoginAndPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postgresConnector_AddUser(t *testing.T) {
	type args struct {
		user model.User
	}
	tests := []struct {
		name    string
		db      *postgresConnector
		args    args
		want    model.User
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.AddUser(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("postgresConnector.AddUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postgresConnector.AddUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postgresConnector_UpdateUser(t *testing.T) {
	type args struct {
		user model.User
	}
	tests := []struct {
		name    string
		db      *postgresConnector
		args    args
		want    int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.UpdateUser(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("postgresConnector.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("postgresConnector.UpdateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postgresConnector_DeleteUser(t *testing.T) {
	type args struct {
		user model.User
	}
	tests := []struct {
		name    string
		db      *postgresConnector
		args    args
		want    int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.DeleteUser(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("postgresConnector.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("postgresConnector.DeleteUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postgresConnector_DeleteUserById(t *testing.T) {
	type args struct {
		user int
	}
	tests := []struct {
		name    string
		db      *postgresConnector
		args    args
		want    int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.DeleteUserById(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("postgresConnector.DeleteUserById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("postgresConnector.DeleteUserById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseUserRows(t *testing.T) {
	type args struct {
		rows *sql.Rows
	}
	tests := []struct {
		name    string
		args    args
		want    model.UsersCollection
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseUserRows(tt.args.rows)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseUserRows() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseUserRows() = %v, want %v", got, tt.want)
			}
		})
	}
}
