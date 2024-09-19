package handlers

import (
	"context"
	"strconv"

	pb "income-expense-service/grpc"
	"income-expense-service/internal/model"
	"income-expense-service/internal/service"
	"income-expense-service/methods"
)

type TransactionHandler struct {
	pb.UnimplementedIncomeExpensesServiceServer
	Service *service.TransactionService
}

func (h *TransactionHandler) LogIncome(ctx context.Context, req *pb.IncomeRequest) (*pb.IncomeResponse, error) {
	tx := model.Transaction{
		Type:     "income",
		Amount:   req.Amount,
		UserID:   int(req.UserId),
		Currency: req.Currency,
		Category: req.Category,
		Date:     req.Date,
	}
	con := methods.ConnectUser()

	_, err := con.GetUser(ctx, &pb.GetUserReq{UserId: int32(tx.UserID)})
	if err != nil {
		return nil, err
	}
	t, err := h.Service.LogIncome(tx)
	if err != nil {
		return nil, err
	}

	return &pb.IncomeResponse{
		Message:       "Income logged successfully",
		TransactionId: t,
	}, nil
}

func (h *TransactionHandler) LogExpense(ctx context.Context, req *pb.ExpenseRequest) (*pb.ExpenseResponse, error) {
	tx := model.Transaction{
		Type:     "expense",
		Amount:   req.Amount,
		UserID:   int(req.UserId),
		Currency: req.Currency,
		Category: req.Category,
		Date:     req.Date,
	}
	con := methods.ConnectUser()

	resp, err := con.GetUser(ctx, &pb.GetUserReq{UserId: int32(tx.UserID)})
	if err != nil {
		return nil, err
	}

	// conn := methods.ConnectBudget()
	// sid := strconv.Itoa(int(req.UserId))
	// rsp, err := conn.GetBudgets(ctx, &pb.GetBudgetsRequest{UserId: sid})
	// var bID string
	// for _, v := range rsp.Budgets {
	// 	if v.Category == tx.Category {
	// 		bID = v.BudgetId
	// 	}
	// }

	// rps, err := conn.PutBudget(ctx, &pb.PatchBudgetRequest{BudgetId: req.UserId})
	// log.Println(err)
	// if err != nil {
	// 	return nil, err
	// }

	t, err := h.Service.LogExpense(tx, resp)
	if err != nil {
		return nil, err
	}

	return &pb.ExpenseResponse{
		Message:       "Xarajat muvaffaqiyatli kiritildi",
		TransactionId: t,
	}, nil
}

func (h *TransactionHandler) GetTransactions(ctx context.Context, req *pb.GetTransactionsRequest) (*pb.GetTransactionsResponse, error) {
	userID := req.GetUserId()

	userId, err := strconv.Atoi(userID)
	if err != nil {
		return nil, err
	}
	con := methods.ConnectUser()

	_, err = con.GetUser(ctx, &pb.GetUserReq{UserId: int32(userId)})
	if err != nil {
		return nil, err
	}

	transactions, err := h.Service.GetTransactions(userID)
	if err != nil {
		return nil, err
	}

	var pbTransactions []*pb.Transaction
	for _, tx := range transactions {
		pbTx := &pb.Transaction{
			TransactionId: tx.TransactionID,
			Type:          tx.Type,
			Amount:        tx.Amount,
			Currency:      tx.Currency,
			Category:      tx.Category,
			Date:          tx.Date,
		}
		pbTransactions = append(pbTransactions, pbTx)
	}

	return &pb.GetTransactionsResponse{
		Transactions: pbTransactions,
	}, nil
}
