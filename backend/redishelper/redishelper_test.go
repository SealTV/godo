package redishelper

import (
	"testing"
	"time"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
)

func TestRedisHelper_SetTokenForUser(t *testing.T) {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer s.Close()

	type args struct {
		user       string
		token      string
		expiration time.Duration
	}
	tests := []struct {
		name    string
		rh      *RedisHelper
		args    args
		wantErr bool
	}{
		{
			name: "1",
			rh:   NewHelper(s.Addr(), 0),
			args: args{
				user:       "user_1",
				token:      "token_1",
				expiration: 1 * time.Millisecond,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.rh.SetTokenForUser(tt.args.user, tt.args.token, tt.args.expiration); (err != nil) != tt.wantErr {
				t.Errorf("RedisHelper.SetTokenForUser() error = %v, wantErr %v", err, tt.wantErr)
			}

			if token, err := s.Get(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("RedisHelper.SetTokenForUser() error = %v, wantErr %v", err, tt.wantErr)
			} else if token != tt.args.token {
				t.Errorf("RedisHelper.SetTokenForUser() token = %v, want token = %v", token, tt.args.token)
			}
		})
	}
}

func TestRedisHelper_GetTokenForUser(t *testing.T) {
	type fields struct {
		r *redis.Client
	}
	type args struct {
		user string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rh := &RedisHelper{
				r: tt.fields.r,
			}
			got, err := rh.GetTokenForUser(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("RedisHelper.GetTokenForUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RedisHelper.GetTokenForUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
