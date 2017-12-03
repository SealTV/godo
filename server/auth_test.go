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

	"bitbucket.org/SealTV/go-site/model"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	mockDB = &dbMock{
		users: map[int]model.User{
			1: {
				Id:           1,
				Login:        "SealTV",
				Email:        "seal@test.com",
				Password:     "pass",
				RegisterDate: time.Now(),
			},
		},
		lists: map[int]model.List{
			1: {
				Id:     1,
				Name:   "List",
				UserId: 1,
			},
		},
		todos: map[int]model.Todo{
			1: {
				Id:          1,
				Title:       "todo1",
				Description: "Todo desc",
				ListId:      1,
				UserId:      1,
			},
		},
	}
)

func TestServerRegister(t *testing.T) {
	e := echo.New()
	user := model.User{
		Id:           2,
		Login:        "Jon",
		Email:        "jon@mail.com",
		Password:     "pass",
		RegisterDate: time.Now(),
	}
	f := make(url.Values)
	f.Set("name", user.Login)
	f.Set("email", user.Email)
	f.Set("password", user.Password)
	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Form = f
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
				var result model.User

				if err := json.Unmarshal(rec.Body.Bytes(), &result); err != nil {
					t.Error(fmt.Errorf("fail"))
				}
				assert.Equal(t, user.Id, result.Id)
				assert.Equal(t, user.Login, result.Login)
				assert.Equal(t, user.Email, result.Email)
				assert.Equal(t, user.Password, result.Password)
			}
		})
	}
}

func TestServerLogin(t *testing.T) {
	//Setup
	e := echo.New()
	user := model.User{
		Id:           1,
		Login:        "SealTV",
		Email:        "seal@test.com",
		Password:     "pass",
		RegisterDate: time.Now(),
	}
	q := make(url.Values)
	q.Set("username", user.Login)
	q.Set("password", user.Password)
	req := httptest.NewRequest(echo.GET, "/?" + q.Encode(), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:email")
	c.SetParamNames("email")
	c.SetParamValues("jon@labstack.com")
	h := &Server{db: mockDB}

	token := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiU2VhbFRWIiwiYWRtaW4iOmZhbHNlLCJleHAiOjE1MTI0MjAzMTMsImp0aSI6IjEifQ.6W2Z2xNRcXeRMlPobWWa885_509HY8iMWPr7f2XQbVEKmvXoi2s2svOiVsf6Elzk_STkAW2xG_k-JFmiTQH1FA"
	// Assertions
	if assert.NoError(t, h.login(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		result := struct {
			message string `json:"message"`
			token string `json:"token"`
		}{}

		ttt := rec.Body.String()
		fmt.Println(ttt)
		if err := json.Unmarshal([]byte(ttt), &result); err != nil {
			t.Error(fmt.Errorf("fail"))
		}

		assert.Equal(t, token, result.token)
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
