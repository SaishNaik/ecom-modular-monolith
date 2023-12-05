//go:build wireinject
// +build wireinject

package app

import (
	"github.com/SaishNaik/ecom-modular-monolith/product/config"
	"github.com/SaishNaik/ecom-modular-monolith/product/repo/mongorepo"
	"github.com/SaishNaik/ecom-modular-monolith/product/usecases"
	"github.com/google/wire"
)

func InitProductApp(cfg *config.Config, mongoUrl string) (*ProductApp, error) {
	panic(wire.Build(NewApp, usecases.NewUseCase, mongorepo.NewMongoDBRepo))
}
