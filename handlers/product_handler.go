package handlers

import (
	"fmt"
	"strconv"

	"github.com/herbetyp/crud-product-api/internal/configs/logger"
	"github.com/herbetyp/crud-product-api/internal/interfaces"
	model "github.com/herbetyp/crud-product-api/models/product"
	"github.com/herbetyp/crud-product-api/services"
	service "github.com/herbetyp/crud-product-api/services"
)

const (
	PRODUCT_PREFIX = "product>>>"
)

type ProductHandler struct {
	repository interfaces.IProductRepository
}

func (h *ProductHandler) CreateProduct(data model.ProductDTO) (model.Product, error) {
	prod := model.NewProduct(data.Name, data.Price, data.Code, data.Qtd, data.Unity)

	p, err := h.repository.Create(*prod)

	if err != nil {
		return model.Product{}, fmt.Errorf("cannot create product: %v", err)
	}

	services.SetCache(PRODUCT_PREFIX+fmt.Sprintf("%d", p.ID), &p)

	return p, nil
}

func (h *ProductHandler) GetProduct(id uint) (model.Product, error) {
	var prod model.Product

	cacheKey := PRODUCT_PREFIX + fmt.Sprintf("%d", id)

	cached := services.GetCache(cacheKey, &prod)
	if cached == "" {
		p, err := h.repository.Get(id)
		if err != nil {
			logger.Error("error on get product from database: %v", err)
			return model.Product{}, err
		}
		services.SetCache(cacheKey, &p)
		prod = p
	}

	return prod, nil
}

func (h *ProductHandler) GetProducts() ([]model.Product, error) {
	var prods []model.Product

	cacheKey := PRODUCT_PREFIX + "all"

	cached := services.GetCache(cacheKey, &prods)
	if cached == "" {
		ps, err := h.repository.GetAll()
		if err != nil {
			logger.Error("error on get products from database: %v", err)
			return nil, err
		}
		services.SetCache(cacheKey, &ps)
		prods = ps
	}

	return prods, nil
}

func (h *ProductHandler) UpdateProduct(data model.ProductDTO) (model.Product, error) {
	prod := model.NewProductWithID(data.ID, data.Name, data.Price, data.Code, data.Qtd, data.Unity)

	p, err := h.repository.Update(*prod)

	if err != nil {
		return model.Product{}, fmt.Errorf("cannot update product: %v", err)
	}

	service.DeleteCache(strconv.FormatUint(uint64(p.ID), 10), PRODUCT_PREFIX, true)

	return p, nil
}

func (h *ProductHandler) DeleteProduct(id uint) (model.Product, error) {
	p, err := h.repository.Delete(id)

	if err != nil {
		return model.Product{}, fmt.Errorf("cannot delete product: %v", err)
	}

	service.DeleteCache(strconv.FormatUint(uint64(p.ID), 10), PRODUCT_PREFIX, true)

	return p, nil
}

func NewProductHandler(r interfaces.IProductRepository) *ProductHandler {
	return &ProductHandler{
		repository: r,
	}
}
