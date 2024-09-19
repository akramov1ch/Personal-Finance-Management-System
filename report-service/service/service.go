package service

import (
	"context"
	"fmt"
	"log"
	"report-service/internal/client/budget"
	"report-service/internal/client/income"
	"report-service/protos/reporting"
	"strings"
)

type NotificationServiceServer struct {
	reporting.UnimplementedReportingServiceServer
}

func NewNotificationServer() *NotificationServiceServer {
	ns := &NotificationServiceServer{}
	return ns
}

func (ns *NotificationServiceServer) GetIncomeExpenseReport(ctx context.Context, req *reporting.GetIncomeExpenseReportRequest) (*reporting.IncomeExpenseReportResponse, error) {
	res, err := income.IncomeClinet(req.UserId)
	if err != nil {
		return nil, err
	}
	var daromad, zarar float64
	for _, exp := range res.Transactions {
		t := strings.ToLower(exp.Category)
		u := strings.ToLower(exp.Currency)
		s := exp.Amount
		if u == "usd" {
			s = exp.Amount * 12.600
		}
		if t == "daromad" {
			daromad += s
		}
		if t == "xarajat" {
			zarar += s
		}
	}
	som := daromad - zarar
	if som < 0 {
		return nil, fmt.Errorf("qarzga kirgansiz")
	}
	return &reporting.IncomeExpenseReportResponse{
		TotalIncome:   daromad,
		TotalExpenses: zarar,
		NetSavings:    som,
	}, nil
}

func (ns *NotificationServiceServer) GetSpendingByCategoryReport(ctx context.Context, req *reporting.GetSpendingByCategoryReportRequest) (*reporting.SpendingByCategoryReportResponse, error) {
	res, err := budget.BudgetClinet(req.UserId)
	if err != nil {
		return nil, err
	}
	var reportList []*reporting.CategorySpending

	for _, report := range res.Budgets {
		log.Printf("BudgetId: %s, Category: %s, Spent: %f\n", report.BudgetId, report.Category, report.Spent)

		reportList = append(reportList, &reporting.CategorySpending{
			Category:   report.Category,
			TotalSpent: report.Spent,
		})
	}

	return &reporting.SpendingByCategoryReportResponse{
		Spending: reportList,
	}, nil
}
