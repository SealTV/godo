package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"bitbucket.org/SealTV/go-site/data"
	"bitbucket.org/SealTV/go-site/model"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/stretchr/testify/assert"
)

func TestServerRegister(t *testing.T) {
	//Setup
	db := data.GetDefaultDBInstance()
	e := echo.New()
	type args struct {
		e *echo.Echo
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		wantErr bool
		user    model.User
	}{
		{
			name:    "1",
			s:       &Server{db: db},
			args:    args{e},
			wantErr: false,
			user: model.User{
				Id:           2,
				Login:        "Jon",
				Email:        "jon@mail.com",
				Password:     "pass",
				RegisterDate: time.Now(),
			},
		},
		{
			name:    "2",
			s:       &Server{db: db},
			args:    args{e},
			wantErr: true,
			user: model.User{
				Id:           2,
				Login:        "Jon",
				Email:        "jon@mail.com",
				Password:     "pass",
				RegisterDate: time.Now(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := make(url.Values)
			f.Set("name", tt.user.Login)
			f.Set("email", tt.user.Email)
			f.Set("password", tt.user.Password)
			req := httptest.NewRequest(echo.POST, "/", strings.NewReader(f.Encode()))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			req.Form = f
			rec := httptest.NewRecorder()
			c := tt.args.e.NewContext(req, rec)

			if assert.NoError(t, tt.s.register(c)) {
				if tt.wantErr {
					assert.Equal(t, http.StatusInternalServerError, rec.Code)
				} else {
					assert.Equal(t, http.StatusCreated, rec.Code)
					var result model.User

					if err := json.Unmarshal(rec.Body.Bytes(), &result); err != nil {
						t.Error(fmt.Errorf("fail"))
					}
					assert.Equal(t, tt.user.Id, result.Id)
					assert.Equal(t, tt.user.Login, result.Login)
					assert.Equal(t, tt.user.Email, result.Email)
					assert.Equal(t, tt.user.Password, result.Password)
				}
			}
		})
	}
}

func TestServerLogin(t *testing.T) {
	//Setup
	type args struct {
		e *echo.Echo
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		wantErr bool
		user    model.User
	}{
		{
			name:    "1",
			s:       &Server{db: data.GetDefaultDBInstance()},
			args:    args{echo.New()},
			wantErr: false,
			user: model.User{
				Id:           1,
				Login:        "SealTV",
				Email:        "seal@test.com",
				Password:     "pass",
				RegisterDate: time.Now(),
			},
		},
		{
			name:    "2",
			s:       &Server{db: data.GetDefaultDBInstance()},
			args:    args{echo.New()},
			wantErr: true,
			user: model.User{
				Id:           2,
				Login:        "Jonn",
				Email:        "jonn@mail.com",
				Password:     "passs",
				RegisterDate: time.Now(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := make(url.Values)
			q.Set("username", tt.user.Login)
			q.Set("password", tt.user.Password)
			req := httptest.NewRequest(echo.GET, "/?"+q.Encode(), nil)
			rec := httptest.NewRecorder()
			c := tt.args.e.NewContext(req, rec)
			c.SetPath("/users/:email")
			c.SetParamNames("email")
			c.SetParamValues(tt.user.Email)

			token, _ := createJwtToken(tt.user)
			// Assertions
			if assert.NoError(t, tt.s.login(c)) {
				if tt.wantErr {
					assert.Equal(t, http.StatusNotFound, rec.Code)
				} else {
					assert.Equal(t, http.StatusOK, rec.Code)

					result := auth{}

					if err := json.Unmarshal([]byte(rec.Body.String()), &result); err != nil {
						t.Error(fmt.Errorf("fail"))
					}

					assert.Equal(t, token, result.Token)
					assert.Equal(t, tt.user.Email, result.User.Email)
					assert.Equal(t, tt.user.Login, result.User.Login)
					assert.Equal(t, tt.user.Password, result.User.Password)
				}
			}
		})
	}
}

func TestServerMainJwt(t *testing.T) {
	//Setup
	type args struct {
		e *echo.Echo
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		wantErr bool
		user    model.User
	}{
		{
			name:    "1",
			s:       &Server{db: data.GetDefaultDBInstance()},
			args:    args{echo.New()},
			wantErr: false,
			user: model.User{
				Id:           1,
				Login:        "SealTV",
				Email:        "seal@test.com",
				Password:     "pass",
				RegisterDate: time.Now(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, _ := createJwtToken(tt.user)

			req := httptest.NewRequest(echo.GET, "/", nil)
			req.Header.Set(echo.HeaderAuthorization, middleware.DefaultJWTConfig.AuthScheme+" "+token)
			rec := httptest.NewRecorder()
			c := tt.args.e.NewContext(req, rec)

			jwt := middleware.DefaultJWTConfig
			jwt.SigningKey = []byte("mySecret")
			jwt.SigningMethod = "HS512"
			jwt.Claims = &jwtClaims{}
			h := middleware.JWTWithConfig(jwt)(tt.s.mainJwt)

			// Assertions
			if assert.NoError(t, h(c)) {
				if tt.wantErr {
					assert.Equal(t, http.StatusNotFound, rec.Code)
				} else {
					assert.Equal(t, http.StatusOK, rec.Code)
				}
			}
		})
	}
}
