package dto

type Request struct {
	Ids []int `json:"id"`
}

type Response struct {
	TotalPrice float64 `json:"total_price"`
}
