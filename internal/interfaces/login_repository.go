package interfaces

import (
	model "github.com/herbetyp/crud-product-api/models/user"
)

type ILoginRepository interface {
	GetLogin(email string) (model.User, error)
}
