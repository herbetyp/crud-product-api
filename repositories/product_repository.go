package repositories

import (
	"database/sql"
	"fmt"

	"github.com/herbetyp/crud-product-api/database"
	"github.com/herbetyp/crud-product-api/models"
)

func GetProductsRepository() ([]models.Product, error) {
	db := database.ConnectDB()
	query := "SELECT * FROM products"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return []models.Product{}, err
	}

	var productList []models.Product
	var productObj models.Product

	for rows.Next() {
		err := rows.Scan(&productObj.ID, &productObj.Name, &productObj.Price, &productObj.CreatedAt, &productObj.UpdatedAt, &productObj.Code, &productObj.Qtd, &productObj.Unity)
		if err != nil {
			fmt.Println(err)
			return []models.Product{}, err
		}
		productList = append(productList, productObj)
	}

	defer rows.Close()
	return productList, nil
}

func CreateProductRepository(product models.Product) (int, error) {
	var id int
	db := database.ConnectDB()
	query, err := db.Prepare("INSERT INTO products (name, price, code, qtd, unity) VALUES ($1, $2, $3, $4, $5) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price, product.Code, product.Qtd, product.Unity).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	query.Close()
	return id, nil
}

func GetProductByIdRepository(id int) (*models.Product, error) {
	db := database.ConnectDB()
	query, err := db.Prepare("SELECT * FROM products WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var product models.Product

	err = query.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Code, &product.Qtd, &product.Unity, &product.CreatedAt, &product.UpdatedAt)
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

func UpdateProductRepository(id int, product models.Product) (int, error) {
	db := database.ConnectDB()
	query, err := db.Prepare("UPDATE products SET name = $1, price = $2 WHERE id = $3")
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

func DeleteProductRepository(id int) (int, error) {
	db := database.ConnectDB()
	query, err := db.Prepare("DELETE FROM products WHERE id = $1 RETURNING id")
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
	query.Close()
	return id, nil
}
