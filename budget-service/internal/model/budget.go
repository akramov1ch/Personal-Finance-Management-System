package model

type Budget struct {
    BudgetID string
    UserID        int
    Amount        float64
    Currency      string
    Category      string
    Spent         float64
    Date          string
}
