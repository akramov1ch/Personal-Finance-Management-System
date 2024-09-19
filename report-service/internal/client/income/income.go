package income

import (
	"context"
	"fmt"
	"report-service/internal/config"
	"report-service/protos/income"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func IncomeClinet(req string) (*income.GetTransactionsResponse, error) {
	cfg := config.Load()
	port := fmt.Sprintf("%s:%s", cfg.IncomeAndExpensesServer.Http.Host, cfg.IncomeAndExpensesServer.Http.Port)

	conn, err := grpc.NewClient(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := income.NewIncomeExpensesServiceClient(conn)

	res, err := client.GetTransactions(context.Background(), &income.GetTransactionsRequest{UserId: req})
	if err != nil {
		return nil, err
	}

	return res, nil
}
