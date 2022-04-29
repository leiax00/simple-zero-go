package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	v1 "github.com/simple-zero-go/api/wx/service/v1"
	"github.com/simple-zero-go/app/wx/service/internal/conf"
	"github.com/simple-zero-go/app/wx/service/internal/service"
	_ "github.com/simple-zero-go/pkg/encoding"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, wx *service.WxSysService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterWxSysHTTPServer(srv, wx)
	return srv
}
