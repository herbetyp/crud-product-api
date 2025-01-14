package repositories

import (
	"time"

	"github.com/herbetyp/crud-product-api/internal/database"
	model "github.com/herbetyp/crud-product-api/models/user"
)

type LoginRepository struct {
}

func (r *LoginRepository) GetLogin(email string) (model.User, error) {
	db := database.GetDatabase()

	var u model.User

	err := db.Where(map[string]interface{}{"email": email}).Find(&u).First(&u).Error

	u.LastLogin = time.Now().Local()

	db.Omit("updated_at").Save(&u)

	return u, err
}
