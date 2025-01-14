package handlers

import (
	"fmt"

	"github.com/herbetyp/crud-product-api/internal/interfaces"
	model "github.com/herbetyp/crud-product-api/models/product"
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

	return p, nil
}

func (h *ProductHandler) GetProduct(id uint) (model.Product, error) {
	p, err := h.repository.Get(id)

	if err != nil {
		return model.Product{}, fmt.Errorf("cannot find product: %v", err)
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
