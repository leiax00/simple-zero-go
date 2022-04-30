package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewWxSysRepo)

// Data .
type Data struct {
	log *log.Helper
	rc  *redis.Client
}

// NewData .
func NewData(c *redis.Client, logger log.Logger) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data"))
	cleanup := func() {
		l.Info("closing the data resources")
		defer func(c *redis.Client) {
			_ = c.Close()
		}(c)
	}
	return &Data{
		log: l,
		rc:  c,
	}, cleanup, nil
}
