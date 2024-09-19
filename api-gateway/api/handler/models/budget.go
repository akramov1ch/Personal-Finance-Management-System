package models

type CreateBudgetForCategoryReq struct {
	Category string  `json:"category"`
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
	Spent    float64 `json:"spent"`
	Date     string  `json:"date"`
}

type CreateBudgetForCategoryResp struct {
	Message string `json:"message"`
	BudgeId string `json:"budge_id"`
}

type GetBudgetsResp struct {
	BudgetId string  `json:"budget_id"`
	Category string  `json:"category"`
	Amount   float64 `json:"amount"`
	Spent    int     `json:"spent"`
	Currency string  `json:"currency"`
}

type PutBudgetReq struct {
	Amount int64 `json:"amount"`
}
type PutBudgetResp struct {
	Message string `json:"message"`
}
