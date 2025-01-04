package usecase

import (
	"github.com/herbetyp/crud-product-api/model"
	"github.com/herbetyp/crud-product-api/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (p *ProductUsecase) GetProducts() ([]model.Product, error) {
	return p.repository.GetProducts()
}

func (p *ProductUsecase) CreateProduct(product model.Product) (model.Product, error) {
	productId, err := p.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = productId
	return product, nil
}

func (p *ProductUsecase) GetProductById(id int) (*model.Product, error) {
	product, err := p.repository.GetProductById(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductUsecase) UpdateProduct(id int, product model.Product) (model.Product, error) {
	updatedProductId, err := p.repository.UpdateProduct(id, product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = updatedProductId
	return product, nil
}
