// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/kalougata/bookkeeping/internal/controller"
	"github.com/kalougata/bookkeeping/internal/data"
	"github.com/kalougata/bookkeeping/internal/server"
	"github.com/kalougata/bookkeeping/internal/service"
	"github.com/kalougata/bookkeeping/pkg/config"
	"github.com/kalougata/bookkeeping/pkg/jwt"
)

// Injectors from wire.go:

func NewApp(conf *config.Config) (*server.Server, func(), error) {
	dataData, cleanup, err := data.NewData(conf)
	if err != nil {
		return nil, nil, err
	}
	jwtJWT := jwt.New(conf)
	userService := service.NewUserService(dataData, jwtJWT)
	authController := controller.NewAuthController(userService)
	tagService := service.NewTagService(dataData)
	tagController := controller.NewTagController(tagService)
	itemService := service.NewItemService(dataData)
	itemController := controller.NewItemController(itemService)
	app := server.NewHTTPServer(authController, tagController, itemController)
	serverServer := server.NewServer(app)
	return serverServer, func() {
		cleanup()
	}, nil
}
