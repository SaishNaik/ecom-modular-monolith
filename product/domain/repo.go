package domain

import "context"

type (
	IProductRepo interface {
		GetAllProducts(context.Context) ([]*ProductModel, error)
	}
)
