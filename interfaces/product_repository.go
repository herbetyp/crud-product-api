package interfaces

import (
	model "github.com/herbetyp/crud-product-api/models/product"
)

type IProductRepository interface {
	Create(p model.Product) (model.Product, error)
	Delete(id uint) (model.Product, error)
	Update(p model.Product) error
	Get(id uint) (model.Product, error)
	GetAll() ([]model.Product, error)
}