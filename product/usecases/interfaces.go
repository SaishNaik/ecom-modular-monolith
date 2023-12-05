package usecases

import (
	"context"
	"github.com/SaishNaik/ecom-modular-monolith/product/domain"
)

type (
	IUseCase interface {
		GetAllProducts(context.Context) ([]*domain.ProductModel, error)
	}
)
