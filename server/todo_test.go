package server

import (
	"testing"

	"github.com/labstack/echo"
)

func TestServer_getTodos(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.getTodos(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Server.getTodos() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_addTodo(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.addTodo(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Server.addTodo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_updateTodo(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.updateTodo(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Server.updateTodo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_deleteTodo(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.deleteTodo(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Server.deleteTodo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
