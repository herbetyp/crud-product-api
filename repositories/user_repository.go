package repositories

import (
	"github.com/herbetyp/crud-product-api/database"
	model "github.com/herbetyp/crud-product-api/models/user"
)

type UserRepository struct {
}

func (r *UserRepository) Create(u model.User) (model.User, error) {
	db := database.GetDatabase()

	err := db.Create(&u).Error

	return u, err
}

func (r *UserRepository) Get(id uint) (model.User, error) {
	db := database.GetDatabase()

	var u model.User

	err := db.First(&u, id).Error

	return u, err
}

func (r *UserRepository) GetAll() ([]model.User, error) {
	db := database.GetDatabase()

	var u []model.User

	err := db.Find(&u).Error

	return u, err
}

func (r *UserRepository) Update(id uint) (model.User, error) {
	db := database.GetDatabase()

	u, err := r.Get(id)

	if err != nil {
		return model.User{}, err
	}

	err = db.Save(&u).Error

	return u, err
}

func (r *UserRepository) Delete(id uint) (model.User, error) {
	db := database.GetDatabase()

	u, err := r.Get(id)

	if err != nil {
		return model.User{}, err
	}

	err = db.Delete(&u).Error

	return u, err
}
