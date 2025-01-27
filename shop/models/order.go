package models

type OrderItem struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type Order struct {
	ID         int         `json:"id"`
	CustomerID int         `json:"customer_id"`
	Items      []OrderItem `json:"items"`
	Total      float64     `json:"total"`
}
