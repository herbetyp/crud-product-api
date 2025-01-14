package handlers

import (
	"fmt"

	"github.com/herbetyp/crud-product-api/internal/interfaces"
	model "github.com/herbetyp/crud-product-api/models/user"
	"github.com/herbetyp/crud-product-api/services"
)

type UserHandler struct {
	repository interfaces.IUserRepository
}

func (h *UserHandler) CreateUser(data model.UserDTO) (model.User, error) {
	user := model.NewUser(data.Username, data.Email, data.Password)

	user.Password = services.SHA512Encoder(user.Password)

	u, err := h.repository.Create(*user)

	if err != nil {
		return model.User{}, fmt.Errorf("cannot create user: %v", err)
	}

	return u, nil
}

func (h *UserHandler) GetUser(id uint) (model.User, error) {
	u, err := h.repository.Get(id)

	if err != nil {
		return model.User{}, fmt.Errorf("cannot find user: %v", err)
	}

	return u, nil
}

func (h *UserHandler) GetUsers() ([]model.User, error) {
	us, err := h.repository.GetAll()

	if err != nil {
		return nil, fmt.Errorf("cannot find users: %v", err)
	}

	return us, nil
}

func (h *UserHandler) UpdateUser(data model.UserDTO) (model.User, error) {
	user := model.NewUserWithID(data.ID, data.Username, data.Email, data.Password)

	user.Password = services.SHA512Encoder(user.Password)

	u, err := h.repository.UpdatePassw(*user)

	if err != nil {
		return model.User{}, fmt.Errorf("cannot update user: %v", err)
	}

	return u, nil
}

func (h *UserHandler) DeleteUser(data model.UserDTO) (model.User, error) {
	user := model.NewUserWithID(data.ID, data.Username, data.Email, data.Password)

	u, err := h.repository.Delete(*user)

	if err != nil {
		return model.User{}, fmt.Errorf("cannot delete user: %v", err)
	}

	return u, nil
}

func (h *UserHandler) RecoveryUser(data model.UserDTO) (model.User, error) {
	user := model.NewUserWithID(data.ID, data.Username, data.Email, data.Password)

	u, err := h.repository.Recovery(*user)

	if err != nil {
		return model.User{}, fmt.Errorf("cannot recovery user: %v", err)
	}

	return u, nil
}

func NewUserHandler(r interfaces.IUserRepository) *UserHandler {
	return &UserHandler{
		repository: r,
	}
}
