package repositories

import (
	"github.com/herbetyp/crud-product-api/internal/database"
	model "github.com/herbetyp/crud-product-api/models/product"
)

type ProductRepository struct {
}

func (r *ProductRepository) Create(p model.Product) (model.Product, error) {
	db := database.GetDatabase()

	err := db.Model(&p).Create(&p).Error

	return p, err
}

func (r *ProductRepository) Get(id uint) (model.Product, error) {
	db := database.GetDatabase()

	var p model.Product

	err := db.Model(&p).First(&p, id).Error

	return p, err
}

func (r *ProductRepository) GetAll() ([]model.Product, error) {
	db := database.GetDatabase()

	var p []model.Product

	err := db.Model(&p).Find(&p).Error

	return p, err
}

func (r *ProductRepository) Update(p model.Product) (model.Product, error) {
	db := database.GetDatabase()

	err := db.Model(&p).Save(&p).Error

	return p, err
}

func (r *ProductRepository) Delete(id uint) (model.Product, error) {
	db := database.GetDatabase()

	var p model.Product

	if err := db.First(&p, id).Error; err != nil {
		return model.Product{}, err
	}

	err := db.Delete(&p).Error
	if err != nil {
		return model.Product{}, err
	}

	return p, nil
}
