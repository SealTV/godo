package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/SealTV/godo/model"

	"github.com/SealTV/godo/data"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestService_authenticator(t *testing.T) {
	s := New(data.New(data.Config{UserDebugDB: true}), Config{
		SecretKey: "secret",
		Host:      "localhost",
		Port:      3000,
	})

	type args struct {
		username string
		password string
		c        *gin.Context
	}
	type want struct {
		username string
		ok       bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{"1", args{"SealTV", "pass", nil}, want{"SealTV", true}},
		{"2", args{"seal@test.com", "pass", nil}, want{"seal@test.com", true}},
		{"3", args{"seal", "pass", nil}, want{"seal", false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			username, ok := s.authenticator(tt.args.username, tt.args.password, tt.args.c)
			if username != tt.want.username {
				t.Errorf("Service.authenticator() got = %v, want %v", username, tt.want.username)
			}
			if ok != tt.want.ok {
				t.Errorf("Service.authenticator() got1 = %v, want %v", ok, tt.want.ok)
			}
		})
	}
}

func TestService_payloadFunc(t *testing.T) {
	s := New(data.New(data.Config{UserDebugDB: true}), Config{
		SecretKey: "secret",
		Host:      "localhost",
		Port:      3000,
	})

	tests := []struct {
		name string
		args string
		want map[string]interface{}
	}{
		{"1", "SealTV", map[string]interface{}{
			"userId": 1,
			"login":  "SealTV",
			"email":  "seal@test.com",
		}},
		{"2", "test", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt)
			if got := s.payloadFunc(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.payloadFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_authorizator(t *testing.T) {
	s := New(data.New(data.Config{UserDebugDB: true}), Config{
		SecretKey: "secret",
		Host:      "localhost",
		Port:      3000,
	})

	type args struct {
		userID string
		c      *gin.Context
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"SealTV", nil}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := s.authorizator(tt.args.userID, tt.args.c); got != tt.want {
				t.Errorf("Service.authorizator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_register(t *testing.T) {
	s := New(data.New(data.Config{UserDebugDB: true}), Config{
		SecretKey: "secret",
		Host:      "localhost",
		Port:      3000,
	})

	type args struct {
		c *http.Request
	}
	type want struct {
		code int
		err  bool
	}
	tests := []struct {
		name string
		args interface{}
		want want
	}{
		{"1", &registerData{"sealTest", "seal@test.t", "pass"}, want{http.StatusOK, false}},
		{"2", &registerData{"sealTest", "seal@test.t", "pass"}, want{http.StatusBadRequest, true}},
		{"3", "some string here", want{http.StatusBadRequest, true}},
		{"4", nil, want{http.StatusBadRequest, true}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var data []byte
			if tt.args != nil {
				data, _ = json.Marshal(tt.args)
			}
			req, _ := http.NewRequest("POST", "/auth/register", bytes.NewBuffer(data))
			w := httptest.NewRecorder()
			s.router.ServeHTTP(w, req)

			assert.Equal(t, tt.want.code, w.Code)
			if !tt.want.err {
				u := model.User{}
				_ = json.Unmarshal(w.Body.Bytes(), &u)
				r := tt.args.(*registerData)
				assert.Equal(t, r.Email, u.Email)
				assert.Equal(t, r.Username, u.Login)
				assert.Equal(t, r.Password, u.Password)

			}
		})
	}
}

func TestService_delete(t *testing.T) {
	s := New(data.New(data.Config{UserDebugDB: true}), Config{
		SecretKey: "secret",
		Host:      "localhost",
		Port:      3000,
	})

	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s.delete(tt.args.c)
		})
	}
}

func TestService_logout(t *testing.T) {
	s := New(data.New(data.Config{UserDebugDB: true}), Config{
		SecretKey: "secret",
		Host:      "localhost",
		Port:      3000,
	})

	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			s.logout(tt.args.c)
		})
	}
}
