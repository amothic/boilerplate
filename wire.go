//+build wireinject

package main

import (
	"github.com/amothic/boilerplate/adapter"
	"github.com/amothic/boilerplate/controller"
	"github.com/amothic/boilerplate/driver"
	"github.com/amothic/boilerplate/usecase"
	"github.com/google/wire"
)

var usecaseSet = wire.NewSet(usecase.NewUserInteractor)
var controllerSet = wire.NewSet(controller.NewUserController)
var persistenceSet = wire.NewSet(adapter.NewUserRepository)

func InitializeServer() (*driver.Server, error) {
	wire.Build(
		driver.NewServer,
		driver.OpenSQL,
		usecaseSet,
		controllerSet,
		persistenceSet,
	)
	return &driver.Server{}, nil
}
