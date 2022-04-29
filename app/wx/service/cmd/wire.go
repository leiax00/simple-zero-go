//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/simple-zero-go/app/wx/service/internal/conf"
	"github.com/simple-zero-go/app/wx/service/internal/server"
	"github.com/simple-zero-go/app/wx/service/internal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	// data.ProviderSet, biz.ProviderSet,
	panic(wire.Build(server.ProviderSet, service.ProviderSet, newApp))
}
