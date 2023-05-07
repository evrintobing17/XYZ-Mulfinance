package models

type Limit struct {
	CustomerID int     `json:"customer_id"`
	Tenor      int     `json:"tenor"`
	Amount     float64 `json:"amount"`
}
