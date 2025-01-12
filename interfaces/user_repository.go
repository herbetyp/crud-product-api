package interfaces

import (
	model "github.com/herbetyp/crud-product-api/models/user"
)

type IUserRepository interface {
	Create(u model.User) (model.User, error)
	Delete(u model.User) (model.User, error)
	UpdatePassw(u model.User) (model.User, error)
	Get(id uint) (model.User, error)
	GetAll() ([]model.User, error)
}
