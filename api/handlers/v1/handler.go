package v1

import (
	log "head_app/pkg/logger"
	"head_app/pkg/token"
	"head_app/storage"
	"head_app/storage/redis"

	"github.com/gin-gonic/gin"
)

type handlers struct {
	storage storage.StorageI
	log     log.Log
	cache   redis.RedisRepoI
}

type Handlers struct {
	Storage storage.StorageI
	Log     log.Log
	Cache   redis.RedisRepoI
}

func NewHandler(h Handlers) handlers {
	return handlers{h.Storage, h.Log, h.Cache}
}

func (h *handlers) Ping(ctx *gin.Context) {
	ctx.JSON(200, map[string]string{"meesage": "pong"})
}

func Auth(ctx *gin.Context) *token.Claim {

	tokenString := ctx.GetHeader("authorization")

	if tokenString == "" {
		ctx.JSON(401, gin.H{"error": "authorization token not provided"})
		ctx.Abort()
		return nil
	}

	claim, err := token.ParseJWT(tokenString)
	if err != nil {
		ctx.JSON(401, gin.H{"error": err.Error()})
		ctx.Abort()
		return nil
	}

	return claim
}
