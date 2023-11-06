//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/kalougata/bookkeeping/internal/controller"
	"github.com/kalougata/bookkeeping/internal/data"
	"github.com/kalougata/bookkeeping/internal/server"
	"github.com/kalougata/bookkeeping/internal/service"
	"github.com/kalougata/bookkeeping/pkg/config"
	"github.com/kalougata/bookkeeping/pkg/jwt"
)

func NewApp(conf *config.Config) (*server.Server, func(), error) {
	panic(wire.Build(
		data.NewData,
		jwt.New,
		service.NewUserService,
		service.NewTagService,
		service.NewItemService,
		controller.NewAuthController,
		controller.NewTagController,
		controller.NewItemController,
		server.NewHTTPServer,
		server.NewServer,
	))
}
