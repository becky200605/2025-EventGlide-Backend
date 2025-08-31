//go:build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/raiki02/EG/internal/controller"
	"github.com/raiki02/EG/internal/dao"
	"github.com/raiki02/EG/internal/ioc"
	"github.com/raiki02/EG/internal/middleware"
	"github.com/raiki02/EG/internal/router"
	"github.com/raiki02/EG/internal/server"
	"github.com/raiki02/EG/internal/service"
)

func InitApp() *server.Server {
	panic(wire.Build(
		ioc.Provider,
		middleware.Provider,
		dao.Provider,
		router.Provider,
		controller.Provider,
		service.Provider,
		server.Provider,
	))
}
