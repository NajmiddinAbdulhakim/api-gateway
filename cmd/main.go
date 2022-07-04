package main

import (
	"fmt"

	"github.com/NajmiddinAbdulhakim/api-gateway/api"
	"github.com/NajmiddinAbdulhakim/api-gateway/config"
	"github.com/NajmiddinAbdulhakim/api-gateway/pkg/logger"
	"github.com/NajmiddinAbdulhakim/api-gateway/services"
	rds "github.com/NajmiddinAbdulhakim/api-gateway/storage/redis"
	_ "github.com/NajmiddinAbdulhakim/api-gateway/storage/repo"
	"github.com/gomodule/redigo/redis"
)

func main() {
	// var redisRepo repo.RedisRepoStorage
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "api_gateway")

	pool := redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort))
			if err != nil {
				panic(err)
			}
			return c, err
		},
	}

	redisRepo := rds.NewRedisRepo(&pool)

	serviceManager, err := services.NewServiceManager(&cfg)
	if err != nil {
		log.Error("gRPC dial error", logger.Error(err))
	}

	server := api.New(api.Option{
		Conf:           cfg,
		Logger:         log,
		ServiceManager: serviceManager,
		RedisRepo:      redisRepo,
	})

	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}
}
