package repositories

import (
	"fmt"

	"github.com/herbetyp/crud-product-api/database"
	"github.com/herbetyp/crud-product-api/models"
)


func CreateUserRepository(user models.User) (string, error) {
	var username string
	db := database.ConnectDB()
	query, err := db.Prepare("INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING username")
	if err != nil {
		panic(err)
	}
	
	err = query.QueryRow(user.Username, user.Email, user.Password).Scan(&username)
	if err != nil {
		fmt.Println(err)
		return username, err
	}
	query.Close()
	return username, nil
}

