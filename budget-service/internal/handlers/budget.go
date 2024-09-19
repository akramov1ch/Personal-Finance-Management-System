package handlers

import (
	"context"
	"strconv"

	pb "income-expense-service/grpc"
	"income-expense-service/internal/model"
	"income-expense-service/internal/service"
	"income-expense-service/methods"
)

type BudgetHandler struct {
	pb.BudgetServiceServer
	Service *service.BudgetService
}

func (h *BudgetHandler) CreateBudget(ctx context.Context, req *pb.CreateBudgetRequest) (*pb.CreateBudgetResponse, error) {
	userid, err := strconv.Atoi(req.UserId)
	if err != nil {
		return nil, err
	}
	tx := model.Budget{
		Amount:   req.Amount,
		UserID:   userid,
		Currency: req.Currency,
		Category: req.Category,
		Spent:    req.Spent,
		Date:     req.Date,
	}

	con := methods.ConnectUser()

	resp, err := con.GetUser(ctx, &pb.GetUserReq{UserId: int32(userid)})
	if err != nil {
		return nil, err
	}
	budgetid, err := h.Service.CreateBudget(tx, resp)
	if err != nil {
		return nil, err
	}

	return &pb.CreateBudgetResponse{
		Message:  "Income logged successfully",
		BudgetId: budgetid,
	}, nil
}

func (h *BudgetHandler) PutBudget(ctx context.Context, req *pb.PatchBudgetRequest) (*pb.PatchBudgetResponse, error) {
	// currentBudget, err := h.Service.GetBudgetsByID(budgetID)
	// if err != nil {
	// 	return nil, err
	// }

	// for _, v := range currentBudget {
	// 	tx := model.Budget{
	// 		BudgetID: req.BudgetId,
	// 		UserID:   v.UserID,
	// 	}

	// 	if req.Amount != nil {
	// 		tx.Amount = float64(req.Amount.Value)
	// 	} else {
	// 		tx.Amount = v.Amount
	// 	}

	// 	if req.Currency != "" {
	// 		tx.Currency = req.Currency
	// 	} else {
	// 		tx.Currency = v.Currency
	// 	}

	// 	if req.Category != "" {
	// 		tx.Category = req.Category
	// 	} else {
	// 		tx.Category = v.Category
	// 	}

	// 	if req.Spent != nil {
	// 		tx.Spent = float64(req.Spent.Value)
	// 	} else {
	// 		tx.Spent = v.Spent
	// 	}

	// 	if req.Date != "" {
	// 		tx.Date = req.Date
	// 	} else {
	// 		tx.Date = v.Date
	// 	}

	// 	con := methods.ConnectUser()

	// 	resp, err := con.GetUser(ctx, &pb.GetUserReq{UserId: int32(tx.UserID)})
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// if err := h.Service.PatchBudget(tx, resp); err != nil {
	// 	return nil, err
	// }
	// }
	err := h.Service.PutBudget(int(req.BudgetId), int(req.Amount))
	if err != nil {
		return nil, err
	}

	return &pb.PatchBudgetResponse{
		Message: "Budget updated successfully",
	}, nil
}

func (h *BudgetHandler) GetBudgets(ctx context.Context, req *pb.GetBudgetsRequest) (*pb.GetBudgetsResponse, error) {
	userID := req.GetUserId()

	userId, err := strconv.Atoi(userID)

	con := methods.ConnectUser()

	_, err = con.GetUser(ctx, &pb.GetUserReq{UserId: int32(userId)})
	if err != nil {
		return nil, err
	}

	budgets, err := h.Service.GetBudgets(userID)
	if err != nil {
		return nil, err
	}

	var pbBudgets []*pb.Budget
	for _, Bd := range budgets {
		pbBd := &pb.Budget{
			BudgetId: Bd.BudgetID,
			Amount:   Bd.Amount,
			Currency: Bd.Currency,
			Category: Bd.Category,
			Spent:    Bd.Spent,
			Date:     Bd.Date,
		}
		pbBudgets = append(pbBudgets, pbBd)
	}
	return &pb.GetBudgetsResponse{
		Budgets: pbBudgets,
	}, nil
}
