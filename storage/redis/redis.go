package redis

import (
	"github.com/NajmiddinAbdulhakim/api-gateway/storage/repo"
	redis "github.com/gomodule/redigo/redis"
)

type redisRepo struct {
	rConn *redis.Pool
}

func NewRedisRepo(rds *redis.Pool) repo.RedisRepoStorage {
	return &redisRepo{rConn: rds}
}

func (r *redisRepo) Set(key, value string) (err error) {
	conn := r.rConn.Get()
	defer conn.Close()

	_, err = conn.Do("SET",key, value)
	return
}

func (r *redisRepo) SetWithTTL(key, value string, sec int64) (err error) {
	conn := r.rConn.Get()
	defer conn.Close()

	_, err = conn.Do("SETEX", key, sec, value)
	return

}

func (r *redisRepo) Get(key string) (interface{} ,error) {
	conn := r.rConn.Get()
	defer conn.Close()

	return conn.Do("GET", key)

}
