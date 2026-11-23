package domain

type ProductForm struct {
	Description string  `json:"description"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}
