package v1

import (
	"github.com/NajmiddinAbdulhakim/api-gateway/config"
	"github.com/NajmiddinAbdulhakim/api-gateway/pkg/logger"
	"github.com/NajmiddinAbdulhakim/api-gateway/services"
	"github.com/NajmiddinAbdulhakim/api-gateway/storage/repo"
)

type handlerV1 struct {
	log            	logger.Logger
	serviceManager 	services.IServiceManager
	cfg            	config.Config
	redisStorage 	repo.RedisRepoStorage
}

// HandlerV1Config ...
type HandlerV1Config struct {
	Logger         	logger.Logger
	ServiceManager 	services.IServiceManager
	Cfg            	config.Config
	Redis 			repo.RedisRepoStorage
}

// New ...
func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		log:            c.Logger,
		serviceManager: c.ServiceManager,
		cfg:            c.Cfg,
		redisStorage: 	c.Redis,
	}
}
