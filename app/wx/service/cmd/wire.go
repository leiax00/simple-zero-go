//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/simple-zero-go/app/wx/service/internal/biz"
	"github.com/simple-zero-go/app/wx/service/internal/conf"
	"github.com/simple-zero-go/app/wx/service/internal/data"
	"github.com/simple-zero-go/app/wx/service/internal/server"
	"github.com/simple-zero-go/app/wx/service/internal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.WxConf, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, biz.ProviderSet, data.ProviderSet, service.ProviderSet, newApp))
}
