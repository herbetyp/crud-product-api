package repository

import (
	"database/sql"
	"fmt"

	"github.com/herbetyp/crud-product-api/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (p *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id, product_name, price FROM product"
	rows, err := p.connection.Query(query)
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

func (p *ProductRepository) CreateProduct(product model.Product) (int, error) {
	var id int
	query, err := p.connection.Prepare("INSERT INTO product (product_name, price) VALUES ($1, $2) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return id, nil
}


func (p *ProductRepository) GetProductById(id int) (*model.Product, error) {
	query, err :=  p.connection.Prepare("SELECT * FROM product WHERE id = $1")
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

func (p *ProductRepository) UpdateProduct(id int, product model.Product) (int, error) {
	query, err := p.connection.Prepare("UPDATE product SET product_name = $1, price = $2 WHERE id = $3 RETURNING id")
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