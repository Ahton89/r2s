package r2b

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"os"
)

func (s *R2s) redisConnect(host string, db int) *redis.Client {
	c := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s", host),
		DB:   db,
	})

	_, err := c.Ping().Result()
	if err != nil {
		s.log.Errorf("Can not connect to redis %s", host)
		os.Exit(1)
	}

	return c
}

func (s *R2s) getHashKeys(hash string) ([]string, error) {
	exist := s.redisProd.Exists(hash).Val()
	if exist != 1 {
		return make([]string, 0), errors.New(fmt.Sprintf("hash %s can not found in production redis", hash))
	}
	return s.redisProd.HKeys(hash).Val(), nil
}

func (s *R2s) getHashValues(hash, key string) string {
	keyValue := s.redisProd.HGet(hash, key).Val()
	return keyValue
}

func (s *R2s) setHash(hash, key, value string) error {
	setHashError := s.redisSandbox.HSet(hash, key, value).Err()
	return setHashError
}
