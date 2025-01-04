package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/herbetyp/crud-product-api/model"
	"github.com/herbetyp/crud-product-api/usecase"
	"strconv"
)

type productController struct {
	ProductUsecase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		ProductUsecase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products, err := p.ProductUsecase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

func (p *productController) CreateProduct(ctx *gin.Context) {
	var product model.Product
	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	insertedProduct, err := p.ProductUsecase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (p *productController) GetProductById(ctx *gin.Context) {
	id := ctx.Param("id")
	if (id == "") {
		response := model.Response{Message: "Missing product ID"}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productID, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{Message: "Invalid product ID. Only integer are allowed"}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	product, err := p.ProductUsecase.GetProductById(productID)
	if product == nil {
		response := model.Response{Message: "Product not found"}
		ctx.JSON(http.StatusNotFound, response)
		return
	}
	ctx.JSON(http.StatusOK, product)
}