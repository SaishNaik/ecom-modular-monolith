package domain

import (
	"github.com/google/uuid"
	"time"
)

type ProductModel struct {
	ID        string    `json:"id"`
	SKU       uuid.UUID `json:"sku" validate:"required"`
	Name      string    `json:"name" validate:"required"`
	Price     float64   `json:"price" validate:"required,gte=1,lte=5000"`
	Quantity  int64     `json:"quantity" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	Image     string    `json:"image" validate:"required"`
}
