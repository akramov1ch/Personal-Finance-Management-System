package models

type ReportingRes struct {
	TotalIncome   int64 `json:"total_income"`
	TotalExpenses int64 `json:"total_expenses"`
	NetSavings    int64 `json:"net_savings"`
}

type CategorySpending struct {
	Category   string `json:"category"`
	TotalSpent int64  `json:"total_spent"`
}