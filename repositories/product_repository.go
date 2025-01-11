package repositories

import (
	"github.com/herbetyp/crud-product-api/database"
	model "github.com/herbetyp/crud-product-api/models/product"
)

type ProductRepository struct {
}

func (r *ProductRepository) Create(p model.Product) (model.Product, error) {
	db := database.GetDatabase()
	
	err := db.Create(&p).Error

	return p, err
}

func (r *ProductRepository) Get(id uint) (model.Product, error) {
	db := database.GetDatabase()
	
	var p model.Product
	
	err := db.First(&p, id).Error
	
	return p, err
}

func (r *ProductRepository) GetAll() ([]model.Product, error) {
	db := database.GetDatabase()

	var p []model.Product

	err := db.Find(&p).Error

	return p, err
}

func (r *ProductRepository) Update(id uint) (model.Product, error) {
	db := database.GetDatabase()

	p, err := r.Get(id)

	if err != nil {
		return model.Product{}, err
	}

	err = db.Save(&p).Error

	return p, err
}

func (r *ProductRepository) Delete(id uint) (model.Product, error) {
	db := database.GetDatabase()
	
	p, err := r.Get(id)

	if err != nil {
		return model.Product{}, err
	}

	err = db.Delete(&p).Error

	return p, err
}
