package cache

import (
	"encoding/json"

	"github.com/gomodule/redigo/redis"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type RedisCache struct {
	pool *redis.Pool
}

func NewRedisCache() *RedisCache {
	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}

	return &RedisCache{pool: pool}
}

func (rc *RedisCache) GetUserByID(userID string) (*User, error) {

	conn := rc.pool.Get()

	defer conn.Close()

	userData, err := redis.Bytes(conn.Do("GET", "user:"+userID))
	if err != nil {
		return nil, err
	}

	var user User
	err = json.Unmarshal(userData, &user)

	if err != nil {
		return nil, err
	}

}
