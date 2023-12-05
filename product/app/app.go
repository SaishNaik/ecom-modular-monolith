package app

import (
	"github.com/SaishNaik/ecom-modular-monolith/product/config"
	"github.com/SaishNaik/ecom-modular-monolith/product/usecases"
)

type ProductApp struct {
	cfg *config.Config
	uc  usecases.IUseCase
}

func NewApp(cfg *config.Config, uc usecases.IUseCase) *ProductApp {
	return &ProductApp{
		cfg: cfg,
		uc:  uc,
	}
}
