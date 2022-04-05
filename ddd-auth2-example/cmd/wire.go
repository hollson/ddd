//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
)

//go:generate wire
var providerSet = wire.NewSet(
	conf.NewViper,
	conf.NewAppConfigCfg,
	conf.NewLoggerCfg,
	conf.NewRedisConfig,
	conf.NewMongoConfig,
	log.NewLogger,
	redis.NewRedis,
	mongo.NewMongo,
	repository.NewRepository,
	aggregate.NewFactory,
	service.NewService,
	adpter.NewSrv,
)

func NewApp() (*adpter.Server, error) {
	panic(wire.Build(providerSet))
}
