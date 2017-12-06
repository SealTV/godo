package server

import (
	"testing"

	"github.com/labstack/echo"
)

func TestServer_getList(t *testing.T) {
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
			if err := tt.s.getList(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Server.getList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_addList(t *testing.T) {
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
			if err := tt.s.addList(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Server.addList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_updateList(t *testing.T) {
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
			if err := tt.s.updateList(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Server.updateList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_deleteList(t *testing.T) {
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
			if err := tt.s.deleteList(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Server.deleteList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
