package models

type Product struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Code      string `json:"code"`
	Qtd       float64 `json:"qtd"`
	Unity     string `json:"unity"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}
