package repositories

import (
	"fmt"

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

func (r *ProductRepository) GetAll() ([]model.Product, error) {
	db := database.GetDatabase()

	var p []model.Product

	err := db.Find(&p).Error

	return p, err
}

func (r *ProductRepository) Get(id uint) (model.Product, error) {
	db := database.GetDatabase()

	var p model.Product

	err := db.First(&p, id).Error

	return p, err
}

func (r *ProductRepository) Update(p model.Product) error {
	db := database.GetDatabase()

	err := db.Save(&p).Error

	return err
}

func (r *ProductRepository) Delete(id uint) (model.Product, error) {
	db := database.GetDatabase()
	n, err := r.Get(id)
	if err != nil {
		return model.Product{}, err
	}

	err = db.Delete(&n).Error
	if err != nil {
		return model.Product{}, fmt.Errorf("cannot delete file: %v", err)
	}

	return n, nil
}
