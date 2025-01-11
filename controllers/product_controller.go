package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/herbetyp/crud-product-api/handlers"
	model "github.com/herbetyp/crud-product-api/models/product"
	repository "github.com/herbetyp/crud-product-api/repositories"
)

func CreateProduct(c *gin.Context) {
	var dto model.ProductDTO

	err := c.BindJSON(&dto)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	repo := &repository.ProductRepository{}
	handler := handlers.NewProductHandler(repo)

	result, err := handler.CreateProduct(dto)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(200, result)
}

func GetProduct(c *gin.Context) {
	id := c.Param("product_id")
	if id == "" {
		c.JSON(400, "Missing product id")
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	repo := &repository.ProductRepository{}
	handler := handlers.NewProductHandler(repo)

	result, err := handler.GetProduct(uint(productId))

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, result)
}

func GetProducts(c *gin.Context) {
	repo := &repository.ProductRepository{}
	handler := handlers.NewProductHandler(repo)

	result, err := handler.GetProducts()

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, result)
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("product_id")

	if id == "" {
		c.JSON(400, "Missing product id")
		return
	}

	_, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	var dto model.ProductDTO

	err = c.BindJSON(&dto)

	if err != nil {
		c.JSON(400, "Invalid request payload")
		return
	}

	repo := &repository.ProductRepository{}
	handler := handlers.NewProductHandler(repo)

	result, err := handler.UpdateProduct(dto)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, result)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("product_id")
	if id == "" {
		c.JSON(400, "Missing product id")
		return
	}

	_, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	var dto model.ProductDTO

	repo := &repository.ProductRepository{}
	handler := handlers.NewProductHandler(repo)

	result, err := handler.DeleteProduct(dto)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(200, result)
}
