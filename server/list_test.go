package server

import (
	"testing"

	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"

	"bitbucket.org/SealTV/go-site/data"
	"bitbucket.org/SealTV/go-site/model"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestServer_getList(t *testing.T) {
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
			s:       &Server{db: data.GetDefaultDBInstance()},
			args:    args{e, model.User{Id: 1, Login: "SealTV", Email: "seal@test.com", Password: "pass", RegisterDate: time.Now()}},
			wantErr: false,
		},
		{
			name:    "2",
			s:       &Server{db: data.GetDefaultDBInstance()},
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

			if assert.NoError(t, tt.s.getLists(c)) {
				assert.Equal(t, http.StatusOK, rec.Code)
				var result model.ListsCollection
				if err := json.Unmarshal(rec.Body.Bytes(), &result); err != nil {
					t.Error(fmt.Errorf("fail"))
				}

				if tt.wantErr {
					assert.Equal(t, 0, len(result))
				} else {
					assert.Equal(t, 1, len(result))
				}
			}
		})
	}
}

func TestServer_addList(t *testing.T) {
	e := echo.New()
	type args struct {
		e    *echo.Echo
		list model.List
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		wantErr bool
	}{
		{
			name: "1",
			s:    &Server{db: data.GetDefaultDBInstance()},
			args: args{e, model.List{
				Id:     1,
				Name:   "List",
				UserId: 1,
			}},
			wantErr: false,
		},
		{
			name: "2",
			s:    &Server{db: data.GetDefaultDBInstance()},
			args: args{e, model.List{
				Id:     -1,
				Name:   "List",
				UserId: -1,
			}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bytes, _ := json.Marshal(tt.args.list)
			req := httptest.NewRequest(echo.POST, "/", strings.NewReader(string(bytes)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := tt.args.e.NewContext(req, rec)

			if assert.NoError(t, tt.s.addList(c)) {
				if tt.wantErr {
					assert.Equal(t, http.StatusBadRequest, rec.Code)
				} else {
					assert.Equal(t, http.StatusCreated, rec.Code)
					var result model.List
					if err := json.Unmarshal(rec.Body.Bytes(), &result); err != nil {
						t.Error(fmt.Errorf("fail"))
					}

					assert.Equal(t, tt.args.list.UserId, result.UserId)
					assert.Equal(t, tt.args.list.Name, result.Name)
				}
			}
		})
	}
}

func TestServer_updateList(t *testing.T) {
	e := echo.New()
	type args struct {
		e    *echo.Echo
		list model.List
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		wantErr bool
	}{
		{
			name: "1",
			s:    &Server{db: data.GetDefaultDBInstance()},
			args: args{e, model.List{
				Id:     1,
				Name:   "List",
				UserId: 1,
			}},
			wantErr: false,
		},
		{
			name: "2",
			s:    &Server{db: data.GetDefaultDBInstance()},
			args: args{e, model.List{
				Id:     -1,
				Name:   "List",
				UserId: 1,
			}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bytes, _ := json.Marshal(tt.args.list)
			req := httptest.NewRequest(echo.POST, "/", strings.NewReader(string(bytes)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := tt.args.e.NewContext(req, rec)

			if assert.NoError(t, tt.s.updateList(c)) {
				assert.Equal(t, http.StatusOK, rec.Code)
				var result int
				if err := json.Unmarshal(rec.Body.Bytes(), &result); err != nil {
					t.Error(fmt.Errorf("fail"))
				}
				assert.Equal(t, 1, result)
			}
		})
	}
}

func TestServer_deleteList(t *testing.T) {
	e := echo.New()
	type args struct {
		e    *echo.Echo
		list model.List
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		wantErr bool
	}{
		{
			name: "1",
			s:    &Server{db: data.GetDefaultDBInstance()},
			args: args{e, model.List{
				Id:     1,
				Name:   "List",
				UserId: 1,
			}},
			wantErr: false,
		},
		{
			name: "2",
			s:    &Server{db: data.GetDefaultDBInstance()},
			args: args{e, model.List{
				Id:     -1,
				Name:   "List",
				UserId: 1,
			}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bytes, _ := json.Marshal(tt.args.list)
			req := httptest.NewRequest(echo.POST, "/", strings.NewReader(string(bytes)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := tt.args.e.NewContext(req, rec)

			if assert.NoError(t, tt.s.deleteList(c)) {
				assert.Equal(t, http.StatusOK, rec.Code)
				var result int
				if err := json.Unmarshal(rec.Body.Bytes(), &result); err != nil {
					t.Error(fmt.Errorf("fail"))
				}
				assert.Equal(t, 1, result)
			}
		})
	}
}
