package main

import (
	"context"
	"fmt"
	"head_app/api"
	"head_app/config"
	"head_app/pkg/db"
	log "head_app/pkg/logger"
	"head_app/storage"
	"head_app/storage/redis"

	"github.com/saidamir98/udevs_pkg/logger"
)

var ctx = context.Background()

func main() {

	cfg := config.Load()

	log := log.NewLogger(cfg.GeneralConfig)

	pgxdb, err := db.ConnectToDb(cfg.PgConfig)
	if err != nil {
		log.Error("error on connecting with database", logger.Error(err))
		return
	}

	log.Debug("successfully connected with database")

	fmt.Println(pgxdb)

	redisCli, err := db.ConnRedis(log, ctx, cfg.RedisConfig)
	if err != nil {
		log.Error("error on connecting with database", logger.Error(err))
		return
	}

	log.Debug("successfully connected with redis")

	cache := redis.NewRedisRepo(redisCli, log)

	storage := storage.NewStorage(pgxdb, log)

	engine := api.Api(api.Options{
		Storage: storage,
		Log:     log,
		Cache:   cache,
	})

	log.Debug("server is running on", logger.String("port", cfg.GeneralConfig.HTTPPort))

	engine.Run(cfg.GeneralConfig.HTTPPort)
}

// func main () {

// 	tokens, err := token.GenerateJWT(models.Claim{UserID: "1001",UserRole: "beyafandi"})

// 	if err != nil {
// 		return
// 	}

// 	fmt.Println(tokens)

// 	time.Sleep(6 * time.Second)

// 	claims, err := token.ParseJWT(tokens)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Println(claims)
// }
