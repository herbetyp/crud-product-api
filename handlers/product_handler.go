package handlers

import (
	"fmt"

	model "github.com/herbetyp/crud-product-api/models/product"
	"github.com/herbetyp/crud-product-api/interfaces"
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

func (h *ProductHandler) GetProductById(id uint) (model.Product, error) {
	p, err := h.repository.Get(id)

	if err != nil {
		return model.Product{}, fmt.Errorf("cannot find product: %v", err)
	}

	return p, nil
}

func (h *ProductHandler) GetProducts() ([]model.Product, error) {
	p, err := h.repository.GetAll()

	if err != nil {
		return nil, fmt.Errorf("cannot find products: %v", err)
	}

	return p, nil
}

func (h *ProductHandler) UpdateProduct(data model.ProductDTO) error {
	prod := model.NewProductWithID(data.Id, data.Name, data.Price, data.Code, data.Qtd, data.Unity)

	err := h.repository.Update(*prod)

	if err != nil {
		return fmt.Errorf("cannot update product: %v", err)
	}	

	return nil
}

func (h *ProductHandler) DeleteProduct(id uint) (model.Product, error) {
	p, err := h.repository.Delete(id)

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
