package store

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

type RedisStore struct {
	Pool   *redis.Pool
	Prefix string
}

func dialWithDB(network string, address string, password string, DB string) (redis.Conn, error) {
	c, err := redis.Dial(network, address)
	if err != nil {
		return nil, err
	}
	if password != "" {
		if _, err := c.Do("AUTH", password); err != nil {
			c.Close()
			return nil, err
		}
	}
	if _, err := c.Do("SELECT", DB); err != nil {
		c.Close()
		return nil, err
	}
	return c, err
}

func NewRedisStore(size int, network string, address string, password string, DB string, prefix string) (Store, error) {
	pool := &redis.Pool{
		MaxIdle:     size,
		IdleTimeout: 240 * time.Second,
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		Dial: func() (redis.Conn, error) {
			return dialWithDB(network, address, password, DB)
		},
	}

	return &RedisStore{Pool: pool, Prefix: prefix}, nil
}

func (s *RedisStore) Close() error {
	return s.Pool.Close()
}

func (s *RedisStore) ping() (bool, error) {
	conn := s.Pool.Get()
	defer conn.Close()
	data, err := conn.Do("PING")
	if err != nil || data == nil {
		return false, err
	}
	return data == "PONG", nil
}

func (s *RedisStore) Get(key string) (string, error) {
	conn := s.Pool.Get()
	defer conn.Close()

	return redis.String(conn.Do("GET", fmt.Sprintf("%s%s", s.Prefix, key)))
}
