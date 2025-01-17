package repositories

import (
	"github.com/herbetyp/crud-product-api/internal/database"
	model "github.com/herbetyp/crud-product-api/models/user"
)

type UserRepository struct {
}

func (r *UserRepository) Create(u model.User) (model.User, error) {
	db := database.GetDatabase()

	err := db.Model(&u).Create(&u).Error

	u.Password = ""
	u.UId = ""

	return u, err
}

func (r *UserRepository) Get(id uint, sensibleFilter bool) (model.User, error) {
	db := database.GetDatabase()

	var u model.User

	if sensibleFilter {
		err := db.Model(&u).Omit("password", "uid").First(&u, id).Error

		return u, err
	} else {
		err := db.Model(&u).First(&u, id).Error
		return u, err
	}
}

func (r *UserRepository) GetAll() ([]model.User, error) {
	db := database.GetDatabase()

	var u []model.User

	err := db.Model(&u).Omit("password", "uid").Find(&u).Error

	return u, err
}

func (r *UserRepository) UpdatePassw(u model.User) (model.User, error) {
	db := database.GetDatabase()

	user, err := r.Get(u.ID, true)

	if err != nil {
		return model.User{}, err
	}

	err = db.Model(&user).Omit("username", "email", "uid").Updates(model.User{Password: u.Password}).Error

	user.Password = ""

	return user, err
}

func (r *UserRepository) Delete(u model.User) (model.User, error) {
	db := database.GetDatabase()

	user, err := r.Get(u.ID, true)

	if err != nil {
		return model.User{}, err
	}

	err = db.Model(&user).Delete(&u).Error

	user.Password = ""
	user.UId = ""

	return user, err
}

func (r *UserRepository) Recovery(u model.User) (model.User, error) {
	db := database.GetDatabase()

	err := db.Unscoped().Model(&u).Where("id", u.ID).Update("deleted_at", nil).Error

	return u, err
}

func GetInfo(id uint) (model.User, error) {
	db := database.GetDatabase()

	var u model.User

	err := db.Model(u).Select("uid, email", "username").First(&u, id).Error

	return u, err
}
