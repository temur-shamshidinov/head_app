package db

import (
	"context"

	"head_app/config"
	log "head_app/pkg/logger"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/saidamir98/udevs_pkg/logger"
)

func RedisAddr(host string, port int) string {

	return host + ":" + strconv.Itoa(port)
}
func ConnRedis(log log.Log, ctx context.Context, cfg config.RedisConfig) (*redis.Client, error) {

	redisCli := redis.NewClient(&redis.Options{

		Addr: RedisAddr(cfg.Host, cfg.Port),
		DB:   cfg.DBIndex,
	})

	_, err := redisCli.Ping(ctx).Result()
	if err != nil {
		log.Error("error on connecting with redis", logger.Error(err))
		return nil, err
	}

	return redisCli, nil
}
