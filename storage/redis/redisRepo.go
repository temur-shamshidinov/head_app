package redis

import (
	"context"
	log "head_app/pkg/logger"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/saidamir98/udevs_pkg/logger"
)

type RedisRepoI interface {
	Exists(ctx context.Context, key string) (bool, error)
	Set(ctx context.Context, key, value string, exp int) error
	Get(ctx context.Context, key string) (string, error)
	GetDel(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, key string) (string, error)
}

type redisRepo struct {
	cli *redis.Client
	log log.Log
}

func NewRedisRepo(cli *redis.Client, log log.Log) RedisRepoI {
	return &redisRepo{cli: cli, log: log}
}

func (r *redisRepo) Exists(ctx context.Context, key string) (bool, error) {

	isExists, err := r.cli.Do(ctx, "EXISTS", key).Result()

	defer r.cli.Close()
	if err != nil {
		r.log.Error("error on checking exists", logger.Error(err))
		return false, nil
	}

	exists, _ := isExists.(int)

	return exists == 1, nil
}

func (r *redisRepo) Set(ctx context.Context, key, value string, exp int) error {

	_, err := r.cli.SetEX(ctx, key, value, time.Second*time.Duration(exp)).Result()
	defer r.cli.Close()
	if err != nil {
		r.log.Error("error on setting to cache", logger.Error(err))
		return err
	}
	return nil
}

func (r *redisRepo) Get(ctx context.Context, key string) (string, error) {

	return "", nil
}

func (r *redisRepo) GetDel(ctx context.Context, key string) (string, error) {

	anyData, err := r.cli.GetDel(ctx, key).Result()

	if err != nil {
		r.log.Error("error on setting to cache", logger.Error(err))
		return "", err
	}

	defer r.cli.Close()
	return anyData, nil
}

func (r *redisRepo) Del(ctx context.Context, key string) (string, error) {
	return "", nil
}
