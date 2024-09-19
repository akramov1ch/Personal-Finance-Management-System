package budget

import (
	"context"
	"fmt"
	"report-service/internal/config"
	"report-service/protos/budget"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func BudgetClinet(req string) (*budget.GetBudgetsResponse, error) {
	cfg := config.Load()
	port := fmt.Sprintf("%s:%s", cfg.BudgetServer.Http.Host, cfg.BudgetServer.Http.Port)

	conn, err := grpc.NewClient(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := budget.NewBudgetServiceClient(conn)

	res, err := client.GetBudgets(context.Background(), &budget.GetBudgetsRequest{UserId: req})
	if err != nil {
		return nil, err
	}

	return res, nil
}
