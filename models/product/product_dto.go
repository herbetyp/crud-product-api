package product

type ProductDTO struct {
	Id    uint    `json:"id"`
	Name  string `json:"name"`
	Price float32 `json:"price"`
	Code  string `json:"code"`
	Qtd   float32    `json:"qtd"`
	Unity string `json:"unity"`
}