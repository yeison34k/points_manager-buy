package domain

type Buy struct {
	ID          string  `json:"id"`
	BuyId       string  `json:"buyId"`
	User        string  `json:"user"`
	ProductName string  `json:"productName"`
	Price       float32 `json:"price"`
	Points      int     `json:"points"`
	CreateDate  string
}
