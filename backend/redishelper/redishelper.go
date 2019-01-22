package redishelper

import (
	"time"

	"github.com/go-redis/redis"
)

// RedisHelper - provides additional functions for redis
type RedisHelper struct {
	r *redis.Client
}

// NewHelper - create new redis helper
func NewHelper(addr string, db int) *RedisHelper {
	opt := redis.Options{
		Addr: addr,
		DB:   db,
	}
	r := redis.NewClient(&opt)
	if err := r.Ping().Err(); err != nil {
		panic(err)
	}

	return &RedisHelper{r}
}

// Close - close redis connection
func (rh *RedisHelper) Close() {
	rh.r.Close()
}

// SetTokenForUser - set user token in redis cash
func (rh *RedisHelper) SetTokenForUser(user string, token string, expiration time.Duration) error {
	return rh.r.Set(user, token, expiration).Err()
}

// GetTokenForUser - return user token from cash
func (rh *RedisHelper) GetTokenForUser(user string) (string, error) {
	return rh.r.Get(user).Result()
}
