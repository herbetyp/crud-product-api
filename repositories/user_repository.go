package repositories

import (
	"github.com/herbetyp/crud-product-api/database"
	model "github.com/herbetyp/crud-product-api/models/user"
)

type UserRepository struct {
}

func (r *UserRepository) Create(u model.User) (model.User, error) {
	db := database.GetDatabase()

	err := db.Model(&u).Create(&u).Error

	return u, err
}

func (r *UserRepository) Get(id uint) (model.User, error) {
	db := database.GetDatabase()

	var u model.User

	err := db.Model(&u).First(&u, id).Error

	return u, err
}

func (r *UserRepository) GetAll() ([]model.User, error) {
	db := database.GetDatabase()

	var u []model.User

	err := db.Model(&u).Find(&u).Error

	return u, err
}

func (r *UserRepository) UpdatePassw(u model.User) (model.User, error) {
	db := database.GetDatabase()

	user, err := r.Get(u.ID)

	if err != nil {
		return model.User{}, err
	}

	err = db.Model(&user).Updates(model.User{Password: u.Password}).Error
	
	return user, err
}

func (r *UserRepository) Delete(u model.User) (model.User, error) {
	db := database.GetDatabase()

	user, err := r.Get(u.ID)

	if err != nil {
		return model.User{}, err
	}

	err = db.Model(&user).Delete(&u).Error

	return user, err
}
