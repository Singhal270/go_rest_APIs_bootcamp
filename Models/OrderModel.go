package Models

type Order struct {
	Id         uint   `json:"id"`
	ProductID  string `json:"product_id"`
	CustomerID string `json:"customer_id"`
	Quantity   int    `json:"quantity"`
	Status     string `json:"status"`
}
