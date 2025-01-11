package product

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"not null"`
	Price     float32        `json:"price" gorm:"not null"`
	Code      string         `json:"code" gorm:"index;unique;not null"`
	Qtd       float32        `json:"qtd" gorm:"not null"`
	Unity     string         `json:"unity" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
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
