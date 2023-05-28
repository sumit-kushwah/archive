package models

type Product struct {
	ProductID   int     `json:"product_id"`
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	CategoryID  int     `json:"category_id"`
	Description string  `json:"description"`
}
