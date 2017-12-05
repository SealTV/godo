package server

import (
	"testing"

	"github.com/labstack/echo"
)

func TestServer_getUser(t *testing.T) {
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
			if err := tt.s.getUser(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Server.getUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_getUserModel(t *testing.T) {
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
			if err := tt.s.getUserModel(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Server.getUserModel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_updateUser(t *testing.T) {
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
			if err := tt.s.updateUser(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Server.updateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_deleteUser(t *testing.T) {
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
			if err := tt.s.deleteUser(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Server.deleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
