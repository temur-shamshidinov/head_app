package v1

import (
	log "head_app/pkg/logger"
	"head_app/storage"
	"head_app/storage/redis"

	"github.com/gin-gonic/gin"
)

type handlers struct{
	storage  storage.StorageI
	log 	 log.Log
	cache   redis.RedisRepoI
}

type Handlers struct{
	Storage  storage.StorageI
	Log 	 log.Log
	Cache   redis.RedisRepoI
}

func NewHandler(h Handlers) handlers{
	return handlers{h.Storage, h.Log,h.Cache}
}


func (h *handlers) Ping(ctx *gin.Context) {
	ctx.JSON(200,map[string]string{"meesage":"pong"})
}

