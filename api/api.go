package api

import (
	v1 "head_app/api/handlers/v1"
	log "head_app/pkg/logger"
	"head_app/storage"
	"head_app/storage/redis"

	"github.com/gin-gonic/gin"
)

type Options struct {
	Storage storage.StorageI
	Log     log.Log
	Cache   redis.RedisRepoI
}

func Api(opt Options) *gin.Engine {

	h := v1.NewHandler(v1.Handlers{Storage: opt.Storage, Log: opt.Log,Cache: opt.Cache})

	engine := gin.Default()

	api := engine.Group("/api")

	api.GET("/ping",h.Ping)

	own := api.Group("/own")
	{
		own.POST("/log-in")
		own.POST("/log-out")
		own.POST("/category", h.CreateCategory)
		own.PUT("/category/:id")
		own.DELETE("/category/:id")
	}

	vw := api.Group("/vw")
	{
		vw.POST("/log-out")
		vw.POST("/comment/:article_id")
	}

	pb := api.Group("/pb")
	{
		pb.POST("/check-user",h.CheckUser)
		pb.POST("/check-otp/:id",h.CheckUser)
		pb.POST("/sign-up")
		pb.POST("/log-in")
		pb.GET("/categories",h.GetCategoriesList)
		pb.GET("/categories/:id",h.GetCategory)
		pb.GET("/sub-categories")
		pb.GET("/sub-categories/:id")
		pb.GET("/articles")
		pb.GET("/articles/:id")
	}

	return engine
}
