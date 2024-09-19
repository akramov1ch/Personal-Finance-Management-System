package model

type Transaction struct {
    TransactionID string
    UserID        int
    Type          string
    Amount        float64
    Currency      string
    Category      string
    Date          string
}
