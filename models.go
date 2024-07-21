package main

func init() {
	recordStore = make(map[string]TransactionDetail)
	recordPoints = make(map[string]int64)
}

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type TransactionDetail struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Total        string `json:"total"`
	Items        []Item `json:"items"`
}

type IdResponse struct {
	ID string `json:"id"`
}

type PointsResponse struct {
	Points int64 `json:"points"`
}
