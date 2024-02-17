package db

import (
	"time"
	"userdata-redis-example-go/config"

	"github.com/gomodule/redigo/redis"
)

type RedisDB struct {
	pool *redis.Pool
}

func NewRedisDB(redisConfig *config.RediConfig) *RedisDB {
	return &RedisDB{
		pool: &redis.Pool{
			MaxIdle:     10,
			IdleTimeout: 240 * time.Second,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", redisConfig.Address)
				if err != nil {
					return nil, err
				}

				if redisConfig.Password != "" {
					if _, err := c.Do("AUTH", redisConfig.Password); err != nil {
						c.Close()
						return nil, err
					}
				}

				if redisConfig.DB != 0 {
					if _, err := c.Do("SELECT", redisConfig.DB); err != nil {
						c.Close()
						return nil, err
					}

				}
				return c, nil
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {

				_, err := c.Do("PING")
				return err
			},
		},
	}

}

func (rdb *RedisDB) Close() {
	rdb.pool.Close()
}

func (rdb *RedisDB) GetRedisConn() redis.Conn {
	return rdb.pool.Get()
}
