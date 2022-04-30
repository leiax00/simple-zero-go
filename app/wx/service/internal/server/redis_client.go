package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/simple-zero-go/app/wx/service/internal/conf"
)

func NewRedisClient(opts *conf.Data, logger log.Logger) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     opts.Redis.Addr,
		Password: opts.Redis.Password,
		DB:       int(opts.Redis.Db),
	})
}
