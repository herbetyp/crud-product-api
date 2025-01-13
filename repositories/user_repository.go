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

	u.Password = ""
	u.UId = ""

	return u, err
}

func (r *UserRepository) Get(id uint) (model.User, error) {
	db := database.GetDatabase()

	var u model.User

	err := db.Model(&u).Omit("password", "uid").First(&u, id).Error

	return u, err
}

func (r *UserRepository) GetAll() ([]model.User, error) {
	db := database.GetDatabase()

	var u []model.User

	err := db.Model(&u).Omit("password", "uid").Find(&u).Error

	return u, err
}

func (r *UserRepository) UpdatePassw(u model.User) (model.User, error) {
	db := database.GetDatabase()

	user, err := r.Get(u.ID)

	if err != nil {
		return model.User{}, err
	}

	err = db.Model(&user).Omit("username", "email", "uid").Updates(model.User{Password: u.Password}).Error

	user.Password = ""
	
	return user, err
}

func (r *UserRepository) Delete(u model.User) (model.User, error) {
	db := database.GetDatabase()

	user, err := r.Get(u.ID)

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


func GetUID(id uint) (string, error) {
	db := database.GetDatabase()

	var u model.User

	err := db.Model(u).Select("uid").First(&u, id).Error

	return u.UId, err
}
