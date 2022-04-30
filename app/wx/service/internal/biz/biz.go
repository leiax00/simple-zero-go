package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/go-resty/resty/v2"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewBiz, NewWxSysUseCase)

type Biz struct {
	log *log.Helper
	hc  *resty.Client
	rc  *redis.Client
}

func NewBiz(hc *resty.Client, rc *redis.Client, logger log.Logger) *Biz {
	return &Biz{
		log: log.NewHelper(log.With(logger, "module", "useCase")),
		hc:  hc,
		rc:  rc,
	}
}
