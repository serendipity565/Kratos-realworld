// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"Kratos-realworld/internal/biz"
	"Kratos-realworld/internal/conf"
	"Kratos-realworld/internal/data"
	"Kratos-realworld/internal/server"
	"Kratos-realworld/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewDB(confData)
	dataData, cleanup, err := data.NewData(confData, logger, db)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	profileRepo := data.NewProfileRepo(dataData, logger)
	userUsecase := biz.NewUserUsecase(userRepo, profileRepo, logger)
	realWorldService := service.NewRealWorldService(userUsecase)
	grpcServer := server.NewGRPCServer(confServer, realWorldService, logger)
	httpServer := server.NewHTTPServer(confServer, realWorldService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
