package service

import (
	"database/sql"
	"fmt"
	"log"

	pb "income-expense-service/grpc"
	"income-expense-service/internal/model"
	"income-expense-service/kafka"

	"github.com/go-redis/redis/v8"
	kf "github.com/segmentio/kafka-go"
)

type BudgetService struct {
	db          *sql.DB
	redisClient *redis.Client
	producer    *kf.Writer
}

func NewBudgetService(db *sql.DB, redisClient *redis.Client) *BudgetService {
	return &BudgetService{db: db, redisClient: redisClient}
}

func (s *BudgetService) CreateBudget(tx model.Budget, resp *pb.GetUserResp) (string, error) {
	var email string
	for _, v := range resp.Response {
		email = v.Email
	}
	if tx.Spent > tx.Amount {
		message := fmt.Sprintf("Spent of %f exceeds the limit of %f", tx.Spent, tx.Amount)
		if err := kafka.SendNotification(email, "Xarajat oshishi", message); err != nil {
			log.Println("Failed to send Kafka notification:", err)
		}
	}
	query := "INSERT INTO budgets (user_id, amount, currency, category, spent,  date) VALUES ($1, $2, $3, $4, $5, $6) returning budget_id"
	rows, err := s.db.Query(query, tx.UserID, tx.Amount, tx.Currency, tx.Category, tx.Spent, tx.Date)
	if err != nil {
		return "", err
	}
	var id string
	if rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return "", err
		}
	}
	return id, err
}

func (s *BudgetService) GetBudgets(userID string) ([]model.Budget, error) {
	rows, err := s.db.Query("SELECT budget_id, user_id, amount, currency, category, spent, date FROM budgets WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var budgets []model.Budget
	for rows.Next() {
		var tx model.Budget
		if err := rows.Scan(&tx.BudgetID, &tx.UserID, &tx.Amount, &tx.Currency, &tx.Category, &tx.Spent, &tx.Date); err != nil {
			return nil, err
		}
		budgets = append(budgets, tx)
	}
	return budgets, nil
}

func (s *BudgetService) PutBudget(budgetId, amount int) error {
	_, err := s.db.Exec("update budgets SET amount = $1 WHERE budget_id = $2", amount, budgetId)
	if err != nil {
		return err
	}
	return nil
}

// func (s *BudgetService) GetBudgetsByID(budgetid int) ([]model.Budget, error) {
// 	rows, err := s.db.Query("SELECT budget_id, user_id, amount, currency, category, spent, date FROM budgets WHERE budget_id = $1", budgetid)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var budgets []model.Budget
// 	for rows.Next() {
// 		var tx model.Budget
// 		if err := rows.Scan(&tx.BudgetID, &tx.UserID, &tx.Amount, &tx.Currency, &tx.Category, tx.Spent, &tx.Date); err != nil {
// 			return nil, err
// 		}
// 		budgets = append(budgets, tx)
// 	}
// 	return budgets, nil
// }

// func (s *BudgetService) PatchBudget(tx model.Budget, resp *pb.GetUserResp) error {
// 	var email string
// 	for _, v := range resp.Response {
// 		email = v.Email
// 	}
// 	if tx.Spent > tx.Amount {
// 		message := fmt.Sprintf("Spent of %f exceeds the limit of %f", tx.Spent, tx.Amount)
// 		if err := kafka.SendNotification(email, "Xarajat oshishi", message); err != nil {
// 			log.Println("Failed to send Kafka notification:", err)
// 		}
// 	}
// 	query := "UPDATE budgets SET"
// 	args := []interface{}{}
// 	argCount := 1

// 	if tx.Amount != 0 {
// 		query += fmt.Sprintf(" amount = $%d,", argCount)
// 		args = append(args, tx.Amount)
// 		argCount++
// 	}

// 	if tx.Currency != "" {
// 		query += fmt.Sprintf(" currency = $%d,", argCount)
// 		args = append(args, tx.Currency)
// 		argCount++
// 	}

// 	if tx.Category != "" {
// 		query += fmt.Sprintf(" category = $%d,", argCount)
// 		args = append(args, tx.Category)
// 		argCount++
// 	}

// 	if tx.Spent != 0 {
// 		query += fmt.Sprintf(" spent = $%d,", argCount)
// 		args = append(args, tx.Spent)
// 		argCount++
// 	}

// 	if tx.Date != "" {
// 		query += fmt.Sprintf(" date = $%d,", argCount)
// 		args = append(args, tx.Date)
// 		argCount++
// 	}

// 	query = query[:len(query)-1]

// 	query += fmt.Sprintf(" WHERE budget_id = $%d", argCount)
// 	args = append(args, tx.BudgetID)

// 	_, err := s.db.Exec(query, args...)
// 	return err
// }
