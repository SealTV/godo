package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"bitbucket.org/SealTV/go-site/model"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	mockDB = &dbMock{
		users: map[int]model.User{},
	}
)

func TestServerRegister(t *testing.T) {
	userJSON := ``
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		wantErr bool
	}{
		{
			name:    "1",
			s:       &Server{db: mockDB},
			args:    args{c},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if assert.NoError(t, tt.s.register(tt.args.c)) {
				assert.Equal(t, http.StatusCreated, rec.Code)
				assert.Equal(t, userJSON, rec.Body.String())
			}
		})
	}
}

func TestServerLogin(t *testing.T) {
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
			if err := tt.s.login(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Server.login() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServerMainJwt(t *testing.T) {
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
			if err := tt.s.mainJwt(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Server.mainJwt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
