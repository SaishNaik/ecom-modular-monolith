package usecases

import (
	"context"
	"github.com/SaishNaik/ecom-modular-monolith/product/domain"
)

type useCase struct {
	repo domain.IProductRepo
}

var _ IUseCase = (*useCase)(nil)

func NewUseCase(repo domain.IProductRepo) IUseCase {
	return &useCase{repo: repo}
}

func (uc *useCase) GetAllProducts(ctx context.Context) ([]*domain.ProductModel, error) {
	return uc.repo.GetAllProducts(ctx)
}
