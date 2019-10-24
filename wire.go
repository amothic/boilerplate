//+build wireinject

package main

import (
	"github.com/amothic/boilerplate/adapter/controller"
	"github.com/amothic/boilerplate/adapter/persistence"
	"github.com/amothic/boilerplate/driver"
	"github.com/amothic/boilerplate/usecase"
	"github.com/google/wire"
)

var usecaseSet = wire.NewSet(usecase.NewUserInteractor)
var controllerSet = wire.NewSet(controller.NewUserController)
var persistenceSet = wire.NewSet(persistence.NewUserRepository)

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
