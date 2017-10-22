package server

import (
	"reflect"
	"testing"

	"bitbucket.org/SealTV/go-site/data"
	"github.com/labstack/echo"
)

func TestGetTodos(t *testing.T) {
	type args struct {
		db data.DBConnector
	}
	tests := []struct {
		name string
		args args
		want echo.HandlerFunc
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTodos(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTodos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddTodo(t *testing.T) {
	type args struct {
		db data.DBConnector
	}
	tests := []struct {
		name string
		args args
		want echo.HandlerFunc
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddTodo(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddTodo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateTodo(t *testing.T) {
	type args struct {
		db data.DBConnector
	}
	tests := []struct {
		name string
		args args
		want echo.HandlerFunc
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpdateTodo(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateTodo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteTodo(t *testing.T) {
	type args struct {
		db data.DBConnector
	}
	tests := []struct {
		name string
		args args
		want echo.HandlerFunc
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteTodo(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteTodo() = %v, want %v", got, tt.want)
			}
		})
	}
}
