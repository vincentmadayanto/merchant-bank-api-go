package models

type Transaction struct {
	CustomerID string  `json:"customer_id"`
	Amount     float64 `json:"amount"`
}
