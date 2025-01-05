package repository

import (
	"database/sql"
	"fmt"

	"github.com/herbetyp/crud-product-api/model"
	"github.com/herbetyp/crud-product-api/database"
)

var connection_db = database.ConnectDB()


func GetProductsRepository() ([]model.Product, error) {
	query := "SELECT id, product_name, price FROM product"
	rows, err := connection_db.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err := rows.Scan(&productObj.ID, &productObj.Name, &productObj.Price)
		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}
		productList = append(productList, productObj)
	}

	defer rows.Close()
	return productList, nil
}

func CreateProductRepository(product model.Product) (int, error) {
	var id int
	query, err := connection_db.Prepare("INSERT INTO product (product_name, price) VALUES ($1, $2) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	query.Close()
	return id, nil
}


func GetProductByIdRepository(id int) (*model.Product, error) {
	query, err :=  connection_db.Prepare("SELECT * FROM product WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var product model.Product

	err = query.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		fmt.Println(err)
		return nil, err
	}
	query.Close()
	return &product, nil
}

func UpdateProductRepository(id int, product model.Product) (int, error) {
	query, err := connection_db.Prepare("UPDATE product SET product_name = $1, price = $2 WHERE id = $3 RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	_, err = query.Exec(product.Name, product.Price, id)
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

func DeleteProductRepository(id int) (*model.Product, error) {
	query, err := connection_db.Prepare("DELETE FROM product WHERE id = $1 RETURNING id")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var product model.Product

	err = query.QueryRow(id).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		fmt.Println(err)
		return nil, err
	}
	query.Close()
	return &product, nil
}