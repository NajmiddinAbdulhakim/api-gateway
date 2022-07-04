package api

import (
	v1 "github.com/NajmiddinAbdulhakim/api-gateway/api/handlers/v1"
	"github.com/NajmiddinAbdulhakim/api-gateway/config"
	"github.com/NajmiddinAbdulhakim/api-gateway/pkg/logger"
	"github.com/NajmiddinAbdulhakim/api-gateway/services"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	swaggerFiles	"github.com/swaggo/files" // swagger embed files
	_ "github.com/NajmiddinAbdulhakim/api-gateway/api/docs"
	"github.com/NajmiddinAbdulhakim/api-gateway/storage/repo"
)

// Option ...
type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
	RedisRepo 	repo.RedisRepoStorage
}

// New ...
func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
		Redis: option.RedisRepo,
	})

	api := router.Group("/v1")
	api.POST("/users", handlerV1.CreateUser)
	api.GET("/users/:id", handlerV1.GetUser)
	api.GET("/users", handlerV1.GetListUsers)
	api.PUT("/users/:id", handlerV1.UpdateUser)
	api.DELETE("/users/:id", handlerV1.DeleteUser)
	api.GET("/users/login", handlerV1.LoginUser)

	api.GET("/allposts", handlerV1.GetAllPosts)
	api.PUT("/post/:id", handlerV1.UpdatePost)
	
	
	url := ginSwagger.URL("seagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,url))
	router.Run()
	return router
}
