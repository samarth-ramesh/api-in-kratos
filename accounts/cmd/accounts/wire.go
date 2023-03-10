//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"accountsapi/accounts/internal/biz"
	"accountsapi/accounts/internal/conf"
	"accountsapi/accounts/internal/data"
	"accountsapi/accounts/internal/server"
	"accountsapi/accounts/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet1, data.ProviderSet2, biz.ProviderSet3, service.ProviderSet, newApp))
}
