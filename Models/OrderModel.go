package Models

type Order struct {
	Id         uint   `json:"id"`
	ProductID  uint   `json:"product_id"`
	CustomerID uint   `json:"customer_id"`
	Quantity   int    `json:"quantity"`
	Status     string `json:"status"`
}
