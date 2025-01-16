package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/herbetyp/crud-product-api/internal/interfaces"
	model "github.com/herbetyp/crud-product-api/models/product"
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

	bytes, _ := json.Marshal(p)
	service.SetCache(fmt.Sprintf("%d", p.ID), string(bytes))

	return p, nil
}

func (h *ProductHandler) GetProduct(id uint) (model.Product, error) {
	var p model.Product
	cacheKey := PRODUCT_PREFIX + fmt.Sprintf("%d", id)

	if cachedData := service.GetCache(cacheKey); cachedData != "" {
		err := json.Unmarshal([]byte(cachedData), &p)
		if err != nil {
			return model.Product{}, nil
		}
	} else {
		p, err := h.repository.Get(id)
		if err != nil {
			return model.Product{}, fmt.Errorf("cannot find product: %v", err)
		}
		cacheValue, _ := json.Marshal(p)
		service.SetCache(cacheKey, string(cacheValue))
	}

	return p, nil
}

func (h *ProductHandler) GetProducts() ([]model.Product, error) {
	ps, err := h.repository.GetAll()

	if err != nil {
		return nil, fmt.Errorf("cannot find products: %v", err)
	}

	return ps, nil
}

func (h *ProductHandler) UpdateProduct(data model.ProductDTO) (model.Product, error) {
	prod := model.NewProductWithID(data.ID, data.Name, data.Price, data.Code, data.Qtd, data.Unity)

	p, err := h.repository.Update(*prod)

	if err != nil {
		return model.Product{}, fmt.Errorf("cannot update product: %v", err)
	}

	service.DeleteCache(fmt.Sprintf("%d", p.ID))

	return p, nil
}

func (h *ProductHandler) DeleteProduct(data model.ProductDTO, id uint) (model.Product, error) {
	prod := model.NewProductWithID(data.ID, data.Name, data.Price, data.Code, data.Qtd, data.Unity)

	p, err := h.repository.Delete(*prod)

	if err != nil {
		return model.Product{}, fmt.Errorf("cannot delete product: %v", err)
	}

	return p, nil
}

func NewProductHandler(r interfaces.IProductRepository) *ProductHandler {
	return &ProductHandler{
		repository: r,
	}
}
