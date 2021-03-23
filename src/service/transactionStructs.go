package service

type Transaction struct {
	transactionIDs string
	Amount         float64 `json:"amount"`
	Type           string  `json:"type"`
	ParentID       float32 `json:"parent_id"`
}

type ErrorResponse struct {
	Status       string `json:"status"`
	ErrorMessage string `json:"errorMessage"`
}
