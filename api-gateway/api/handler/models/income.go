package models

type RegisterIncomeReq struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
	Category string  `json:"category"`
	Date     string  `json:"date"`
}

type RegisterIncomeResp struct {
	Message       string `json:"message"`
	TransactionId string `json:"transaction_id"`
}

type GetTransactionsRes struct {
	TransactionId string  `json:"transaction_id"`
	Type          string  `json:"type"`
	Amount        float64 `json:"amount"`
	Currency      string  `json:"currency"`
	Category      string  `json:"category"`
	Date          string  `json:"date"`
}