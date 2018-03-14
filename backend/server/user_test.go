package server

import (
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"encoding/json"
	"fmt"
	"net/http"

	"bitbucket.org/SealTV/go-site/backend/data"
	"bitbucket.org/SealTV/go-site/backend/model"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestServer_getUser(t *testing.T) {
	e := echo.New()
	type args struct {
		e    *echo.Echo
		user model.User
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		wantErr bool
	}{
		{
			name:    "1",
			s:       &Server{db: data.New(data.Config{UserDebugDB: true})},
			args:    args{e, model.User{Id: 1, Login: "SealTV", Email: "seal@test.com", Password: "pass", RegisterDate: time.Now()}},
			wantErr: false,
		},
		{
			name:    "2",
			s:       &Server{db: data.New(data.Config{UserDebugDB: true})},
			args:    args{e, model.User{Id: -2, Login: "Empty", Email: "emty@test.com", Password: "passEmpty", RegisterDate: time.Now()}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bytes, _ := json.Marshal(tt.args.user)
			req := httptest.NewRequest(echo.POST, "/", strings.NewReader(string(bytes)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := tt.args.e.NewContext(req, rec)

			if assert.NoError(t, tt.s.getUser(c)) {
				if tt.wantErr {
					assert.Equal(t, http.StatusBadRequest, rec.Code)
				} else {
					assert.Equal(t, http.StatusOK, rec.Code)
					var result model.User

					if err := json.Unmarshal(rec.Body.Bytes(), &result); err != nil {
						t.Error(fmt.Errorf("fail"))
					}
					assert.Equal(t, tt.args.user.Id, result.Id)
					assert.Equal(t, tt.args.user.Login, result.Login)
					assert.Equal(t, tt.args.user.Email, result.Email)
					assert.Equal(t, tt.args.user.Password, result.Password)
				}
			}
		})
	}
}

func TestServer_getUserModel(t *testing.T) {
	e := echo.New()
	type args struct {
		e    *echo.Echo
		user model.User
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		wantErr bool
	}{
		{
			name:    "1",
			s:       &Server{db: data.New(data.Config{UserDebugDB: true})},
			args:    args{e, model.User{Id: 1, Login: "SealTV", Email: "seal@test.com", Password: "pass", RegisterDate: time.Now()}},
			wantErr: false,
		},
		{
			name:    "2",
			s:       &Server{db: data.New(data.Config{UserDebugDB: true})},
			args:    args{e, model.User{Id: -2, Login: "SealTVV", Email: "seal@test.com", Password: "pass", RegisterDate: time.Now()}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bytes, _ := json.Marshal(tt.args.user)
			req := httptest.NewRequest(echo.POST, "/", strings.NewReader(string(bytes)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := tt.args.e.NewContext(req, rec)

			if assert.NoError(t, tt.s.getUserModel(c)) {
				if tt.wantErr {
					assert.Equal(t, http.StatusBadRequest, rec.Code)
				} else {
					assert.Equal(t, http.StatusOK, rec.Code)
					var result model.UserModel

					if err := json.Unmarshal(rec.Body.Bytes(), &result); err != nil {
						t.Error(fmt.Errorf("fail"))
					}
					assert.Equal(t, tt.args.user.Id, result.Id)
					assert.Equal(t, tt.args.user.Login, result.Login)
					assert.Equal(t, tt.args.user.Email, result.Email)
					assert.Equal(t, tt.args.user.Password, result.Password)
					assert.Equal(t, 1, len(result.TodoLists))
					assert.Equal(t, 1, len(result.TodoLists[0].Todos))
				}
			}
		})
	}
}

func TestServer_updateUser(t *testing.T) {
	e := echo.New()
	type args struct {
		e    *echo.Echo
		user model.User
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		wantErr bool
	}{
		{
			name:    "1",
			s:       &Server{db: data.New(data.Config{UserDebugDB: true})},
			args:    args{e, model.User{Id: 1, Login: "SealTV", Email: "seal@test.com", Password: "pass", RegisterDate: time.Now()}},
			wantErr: false,
		},
		{
			name:    "2",
			s:       &Server{db: data.New(data.Config{UserDebugDB: true})},
			args:    args{e, model.User{Id: -2, Login: "SealTVV", Email: "seal@test.com", Password: "pass", RegisterDate: time.Now()}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bytes, _ := json.Marshal(tt.args.user)
			req := httptest.NewRequest(echo.POST, "/", strings.NewReader(string(bytes)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := tt.args.e.NewContext(req, rec)

			if assert.NoError(t, tt.s.updateUser(c)) {
				if tt.wantErr {
					assert.Equal(t, http.StatusBadRequest, rec.Code)
				} else {
					assert.Equal(t, http.StatusOK, rec.Code)
					var result int
					if err := json.Unmarshal(rec.Body.Bytes(), &result); err != nil {
						t.Error(fmt.Errorf("fail"))
					}
					assert.Equal(t, 1, result)
				}
			}
		})
	}
}

func TestServer_deleteUser(t *testing.T) {
	e := echo.New()
	type args struct {
		e    *echo.Echo
		user model.User
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		wantErr bool
	}{
		{
			name:    "1",
			s:       &Server{db: data.New(data.Config{UserDebugDB: true})},
			args:    args{e, model.User{Id: 1, Login: "SealTV", Email: "seal@test.com", Password: "pass", RegisterDate: time.Now()}},
			wantErr: false,
		},
		{
			name:    "2",
			s:       &Server{db: data.New(data.Config{UserDebugDB: true})},
			args:    args{e, model.User{Id: -2, Login: "SealTVV", Email: "seal@test.com", Password: "pass", RegisterDate: time.Now()}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bytes, _ := json.Marshal(tt.args.user)
			req := httptest.NewRequest(echo.POST, "/", strings.NewReader(string(bytes)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := tt.args.e.NewContext(req, rec)

			if assert.NoError(t, tt.s.deleteUser(c)) {
				if tt.wantErr {
					assert.Equal(t, http.StatusBadRequest, rec.Code)
				} else {
					assert.Equal(t, http.StatusOK, rec.Code)
					var result int

					if err := json.Unmarshal(rec.Body.Bytes(), &result); err != nil {
						t.Error(fmt.Errorf("fail"))
					}
					assert.Equal(t, 1, result)
				}
			}
		})
	}
}
