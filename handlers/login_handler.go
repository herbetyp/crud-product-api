package handlers

import "github.com/herbetyp/crud-product-api/repositories"

func LoginHandler() {
	repositories.UsersLoginRepository()
}
