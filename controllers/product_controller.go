package controllers

import (
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/herbetyp/crud-product-api/handlers"
	"github.com/herbetyp/crud-product-api/models"
	"github.com/herbetyp/crud-product-api/repositories"
)


func CreateProductController(ctx *gin.Context) {
	var product models.Product
	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	newProductId, err := handlers.CreateProductHandler(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var newProduct *models.Product
	newProduct, _ = repositories.GetProductByIdRepository(newProductId)
	ctx.JSON(http.StatusCreated, newProduct)
}

func GetProductsController(ctx *gin.Context) {
	products, err := handlers.GetProductsHandler()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

func GetProductByIdController(ctx *gin.Context) {
	id := ctx.Param("product_id")
	if (id == "") {
		response := models.Response{Message: "Missing product ID"}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productID, err := strconv.Atoi(id)
	if err != nil {
		response := models.Response{Message: "Invalid product ID. Only integer are allowed"}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	product, err := handlers.GetProductByIdHandler(productID)
	if (product.ID != 0 && err != nil) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if (product.ID == 0) {
		response := models.Response{Message: "Product not found"}
		ctx.JSON(http.StatusNotFound, response)
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func UpdateProductController(ctx *gin.Context) {
	id := ctx.Param("product_id")
	if (id == "") {
		response := models.Response{Message: "Missing product ID"}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productID, err := strconv.Atoi(id)
	if err != nil {
		response := models.Response{Message: "Invalid product ID. Only integer are allowed"}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var product models.Product
	err = ctx.BindJSON(&product)
	if err != nil {
		response := models.Response{Message: "Invalid request payload"}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	updatedProductId, err := handlers.UpdateProductHandler(productID, product)
	if (updatedProductId != 0 && err != nil) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if (updatedProductId == 0) {
		response := models.Response{Message: "Product not found"}
		ctx.JSON(http.StatusNotFound, response)
		return
	}
	var updatedProduct *models.Product
	updatedProduct, _ = repositories.GetProductByIdRepository(updatedProductId)
	ctx.JSON(http.StatusOK, updatedProduct)
}

func DeleteProductController(ctx *gin.Context) {
	id := ctx.Param("product_id")
	if (id == "") {
		response := models.Response{Message: "Missing product ID"}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productID, err := strconv.Atoi(id)
	if err != nil {
		response := models.Response{Message: "Invalid product ID. Only integer are allowed"}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	deletedProductId, err := handlers.DeleteProductHandler(productID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if (deletedProductId == 0) {
		response := models.Response{Message: "Product not found"}
		ctx.JSON(http.StatusNotFound, response)
		return
	}
	response := models.Response{Message: "Product deleted successfully"}
	ctx.JSON(http.StatusOK, response)
}