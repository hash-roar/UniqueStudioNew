package dbhandlers

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxActive:   800,
		MaxIdle:     20,
		IdleTimeout: time.Second * 100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func GetRedisConn() redis.Conn {
	return pool.Get()
}
