package handlers

import (
	"github.com/herbetyp/crud-product-api/model"
	"github.com/herbetyp/crud-product-api/repository"
)

type ProductHandler struct {
	repository repository.ProductRepository
}

func NewProductHandler(repo repository.ProductRepository) ProductHandler {
	return ProductHandler{
		repository: repo,
	}
}

func (p *ProductHandler) GetProducts() ([]model.Product, error) {
	return p.repository.GetProducts()
}

func (p *ProductHandler) CreateProduct(product model.Product) (model.Product, error) {
	productId, err := p.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = productId
	return product, nil
}

func (p *ProductHandler) GetProductById(id int) (*model.Product, error) {
	retriveProduct, err := p.repository.GetProductById(id)
	if err != nil {
		return nil, err
	}

	return retriveProduct, nil
}

func (p *ProductHandler) UpdateProduct(id int, product model.Product) (model.Product, error) {
	updatedProductId, err := p.repository.UpdateProduct(id, product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = updatedProductId
	return product, nil
}

func (p *ProductHandler) DeleteProduct(id int) (*model.Product, error) {
	deletedProduct, err := p.repository.DeleteProduct(id)
	if err != nil {
		return nil, err
	}

	return deletedProduct, nil
}
