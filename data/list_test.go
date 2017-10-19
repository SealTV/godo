package data

import (
	"database/sql"
	"reflect"
	"testing"

	"bitbucket.org/SealTV/go-site/model"
)

func Test_postgresConnector_GetAllLists(t *testing.T) {
	tests := []struct {
		name    string
		db      *postgresConnector
		want    model.ListsCollection
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.GetAllLists()
			if (err != nil) != tt.wantErr {
				t.Errorf("postgresConnector.GetAllLists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postgresConnector.GetAllLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postgresConnector_GetAllListsForUser(t *testing.T) {
	type args struct {
		user model.User
	}
	tests := []struct {
		name    string
		db      *postgresConnector
		args    args
		want    model.ListsCollection
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.GetAllListsForUser(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("postgresConnector.GetAllListsForUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postgresConnector.GetAllListsForUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postgresConnector_GetAllListsForUserId(t *testing.T) {
	type args struct {
		user int
	}
	tests := []struct {
		name    string
		db      *postgresConnector
		args    args
		want    model.ListsCollection
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.GetAllListsForUserId(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("postgresConnector.GetAllListsForUserId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postgresConnector.GetAllListsForUserId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postgresConnector_GetListById(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		db      *postgresConnector
		args    args
		want    model.List
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.GetListById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("postgresConnector.GetListById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postgresConnector.GetListById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postgresConnector_AddList(t *testing.T) {
	type args struct {
		list model.List
	}
	tests := []struct {
		name    string
		db      *postgresConnector
		args    args
		want    model.List
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.AddList(tt.args.list)
			if (err != nil) != tt.wantErr {
				t.Errorf("postgresConnector.AddList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postgresConnector.AddList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postgresConnector_UpdateList(t *testing.T) {
	type args struct {
		list model.List
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
			got, err := tt.db.UpdateList(tt.args.list)
			if (err != nil) != tt.wantErr {
				t.Errorf("postgresConnector.UpdateList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("postgresConnector.UpdateList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postgresConnector_DeleteList(t *testing.T) {
	type args struct {
		list model.List
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
			got, err := tt.db.DeleteList(tt.args.list)
			if (err != nil) != tt.wantErr {
				t.Errorf("postgresConnector.DeleteList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("postgresConnector.DeleteList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postgresConnector_DeleteListById(t *testing.T) {
	type args struct {
		list int
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
			got, err := tt.db.DeleteListById(tt.args.list)
			if (err != nil) != tt.wantErr {
				t.Errorf("postgresConnector.DeleteListById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("postgresConnector.DeleteListById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseListsRows(t *testing.T) {
	type args struct {
		rows *sql.Rows
	}
	tests := []struct {
		name    string
		args    args
		want    model.ListsCollection
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseListsRows(tt.args.rows)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseListsRows() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseListsRows() = %v, want %v", got, tt.want)
			}
		})
	}
}
