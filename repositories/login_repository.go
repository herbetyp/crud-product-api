package repositories

import (
	"fmt"

	"github.com/herbetyp/crud-product-api/database"
	"github.com/herbetyp/crud-product-api/models"
)

func UsersLoginRepository(l *models.Login) (*models.UserModel, error) {
	db := database.ConnectDB()
	query, err := db.Prepare("SELECT * FROM users WHERE email = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var user models.UserModel
	err = query.QueryRow(l.Email).Scan(&user.ID, &user.Username, &user.Email,
		&user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	query.Close()
	return &user, nil
}
