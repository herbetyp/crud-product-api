package product

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Price     float32        `json:"price"`
	Code      string         `json:"code"`
	Qtd       float32        `json:"qtd"`
	Unity     string         `json:"unity"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted"`
}

func NewProduct(name string, price float32, code string, qtd float32, unity string) *Product {
	return &Product{
		Name:  name,
		Price: price,
		Code:  code,
		Qtd:   qtd,
		Unity: unity,
	}
}

func NewProductWithID(id uint, name string, price float32, code string, qtd float32, unity string) *Product {
	return &Product{
		ID:    id,
		Name:  name,
		Price: price,
		Code:  code,
		Qtd:   qtd,
		Unity: unity,
	}
}
