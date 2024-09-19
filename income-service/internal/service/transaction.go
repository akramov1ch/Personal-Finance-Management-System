package service

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	pb "income-expense-service/grpc"
	"income-expense-service/internal/model"
	"income-expense-service/kafka"

	"github.com/go-redis/redis/v8"
)

type TransactionService struct {
	db          *sql.DB
	RedisClient *redis.Client
}

func NewTransactionService(db *sql.DB, redisClient *redis.Client) *TransactionService {
	return &TransactionService{db: db, RedisClient: redisClient}
}

func (s *TransactionService) LogIncome(tx model.Transaction) (string, error) {
	query := "INSERT INTO transactions (user_id, type, amount, currency, category, date) VALUES ($5,'income', $1, $2, $3, $4) returning transaction_id"
	var transactionId string
	_ = s.db.QueryRow(query, tx.Amount, tx.Currency, tx.Category, tx.Date, tx.UserID).Scan(&transactionId)
	return transactionId, nil
}

func (s *TransactionService) LogExpense(tx model.Transaction, resp *pb.GetUserResp) (string, error) {
	expenseLimitStr := os.Getenv("EXPENSE_LIMIT")
	expenseLimit, err := strconv.ParseFloat(expenseLimitStr, 64)
	if err != nil {
		return "", fmt.Errorf("failed to parse EXPENSE_LIMIT: %v", err)
	}

	var email string
	for _, v := range resp.Response {
		email = v.Email
		break
	}

	if tx.Amount > expenseLimit {
		message := fmt.Sprintf("Expense of %f exceeds the limit of %f", tx.Amount, expenseLimit)
		if err := kafka.SendNotification(email, "Xarajat oshishi", message); err != nil {
			log.Printf("Failed to send Kafka notification: %v", err)
			return "", fmt.Errorf("failed to send notification: %v", err)
		}
	}

	query := `
		INSERT INTO transactions (user_id, type, amount, currency, category, date)
		VALUES ($1, 'expense', $2, $3, $4, $5)
		RETURNING transaction_id`
	var transactionId string
	err = s.db.QueryRow(query, tx.UserID, tx.Amount, tx.Currency, tx.Category, tx.Date).Scan(&transactionId)
	if err != nil {
		return "", fmt.Errorf("failed to insert transaction: %v", err)
	}

	return transactionId, nil
}

func (s *TransactionService) GetTransactions(userID string) ([]model.Transaction, error) {
	rows, err := s.db.Query("SELECT transaction_id, type, amount, currency, category, date FROM transactions WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []model.Transaction
	for rows.Next() {
		var tx model.Transaction
		if err := rows.Scan(&tx.TransactionID, &tx.Type, &tx.Amount, &tx.Currency, &tx.Category, &tx.Date); err != nil {
			return nil, err
		}
		transactions = append(transactions, tx)
	}
	return transactions, nil
}
