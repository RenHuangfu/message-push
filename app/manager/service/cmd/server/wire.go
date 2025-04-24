//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"message-push/app/manager/service/internal/conf"
	"message-push/app/manager/service/internal/di"
)

// wireApp init kratos application.
func wireApp(*conf.Bootstrap, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		di.ServerProviderSet,
		di.ServiceProviderSet,
		di.UsecaseProviderSet,
		di.DataProviderSet,
		newApp,
	))
}
