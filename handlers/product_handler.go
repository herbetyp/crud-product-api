package handlers

import (
	"github.com/herbetyp/crud-product-api/models"
	"github.com/herbetyp/crud-product-api/repositories"
)


func GetProductsHandler() ([]models.Product, error) {
	return repositories.GetProductsRepository()
}

func CreateProductHandler(p models.Product) (int, error) {
	newProductId, err := repositories.CreateProductRepository(p)
	if err != nil {
		return 0, err
	}

	return newProductId, nil
}

func GetProductByIdHandler(id int) (*models.Product, error) {
	getProduct, err := repositories.GetProductByIdRepository(id)
	if err != nil {
		return &models.Product{}, err
	}

	return getProduct, nil
}

func UpdateProductHandler(id int, p models.Product) (int, error) {
	updatedProductId, err := repositories.UpdateProductRepository(id, p)
	if err != nil {
		return 0, err
	}

	return updatedProductId, nil
}

func DeleteProductHandler(id int) (int, error) {
	deletedProductId, err := repositories.DeleteProductRepository(id)
	if err != nil {
		return 0, err
	}

	return deletedProductId, nil
}
