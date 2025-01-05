package handlers

import (
	"github.com/herbetyp/crud-product-api/model"
	"github.com/herbetyp/crud-product-api/repository"
)


func GetProductsHandler() ([]model.Product, error) {
	return repository.GetProductsRepository()
}

func CreateProductHandler(product model.Product) (model.Product, error) {
	productId, err := repository.CreateProductRepository(product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = productId
	return product, nil
}

func GetProductByIdHandler(id int) (*model.Product, error) {
	retriveProduct, err := repository.GetProductByIdRepository(id)
	if err != nil {
		return nil, err
	}

	return retriveProduct, nil
}

func UpdateProductHandler(id int, product model.Product) (model.Product, error) {
	updatedProductId, err := repository.UpdateProductRepository(id, product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = updatedProductId
	return product, nil
}

func DeleteProductHandler(id int) (*model.Product, error) {
	deletedProduct, err := repository.DeleteProductRepository(id)
	if err != nil {
		return nil, err
	}

	return deletedProduct, nil
}
