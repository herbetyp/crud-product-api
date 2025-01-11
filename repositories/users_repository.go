package repositories

import (
	"database/sql"
	"fmt"

	"github.com/herbetyp/crud-product-api/database"
	"github.com/herbetyp/crud-product-api/models"
)

func CreateUserRepository(u models.UserModel) (string, error) {
	var username string
	db := database.ConnectDB()
	query, err := db.Prepare("INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING username")
	if err != nil {
		panic(err)
	}

	err = query.QueryRow(u.Username, u.Email, u.Password).Scan(&username)
	if err != nil {
		fmt.Println(err)
		return username, err
	}
	query.Close()
	return username, nil
}

func GetUserByIdRepository(id int) (*models.UserModel, error) {
	db := database.ConnectDB()
	query, err := db.Prepare("SELECT * FROM users WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var user models.UserModel
	err = query.QueryRow(id).Scan(&user.ID, &user.Username, &user.Email,
		&user.Password, &user.UId, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	query.Close()
	return &user, nil
}

func UpdateUserPassRepository(id int, p string) (int, error) {
	db := database.ConnectDB()
	query, err := db.Prepare("UPDATE users SET password = $1 WHERE id = $2 RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(p, id).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		fmt.Println(err)
		return 0, err
	}
	query.Close()
	return id, nil
}

func DeleteUserRepository(id int) (int, error) {
	db := database.ConnectDB()
	query, err := db.Prepare("DELETE FROM users WHERE id = $1 RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(id).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		fmt.Println(err)
		return 0, err
	}

	return id, nil
}
