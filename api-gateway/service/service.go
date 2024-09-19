package service

import (
	"api-gateway/internal/config"

	budgetpb "api-gateway/protos/budget"
	incomepb "api-gateway/protos/income"
	reportingpb "api-gateway/protos/reporting"
	userpb "api-gateway/protos/user"

	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IClients interface {
	Budget() budgetpb.BudgetServiceClient
	IncomeAndExpenses() incomepb.IncomeExpensesServiceClient
	User() userpb.UserServiceClient
	Reporting() reportingpb.ReportingServiceClient
}

type ServiceManager struct {
	BudgetService            budgetpb.BudgetServiceClient
	IncomeAndExpensesService incomepb.IncomeExpensesServiceClient
	userService              userpb.UserServiceClient
	ReportingService         reportingpb.ReportingServiceClient
}

func New(cfg config.Config) (*ServiceManager, error) {

	connUser, err := grpc.NewClient(
		fmt.Sprintf("%s:%s", cfg.UserServer.Http.Host, cfg.UserServer.Http.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to dial user service: %v", err)
	}

	connReporting, err := grpc.NewClient(
		fmt.Sprintf("%s:%s", cfg.ReportingServer.Http.Host, cfg.ReportingServer.Http.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to dial reporting service: %v", err)
	}

	connincome, err := grpc.NewClient(
		fmt.Sprintf("%s:%s", cfg.IncomeAndExpensesServer.Http.Host, cfg.IncomeAndExpensesServer.Http.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to dial icome service: %v", err)
	}

	connBudget, err := grpc.NewClient(
		fmt.Sprintf("%s:%s", cfg.BudgetServer.Http.Host, cfg.BudgetServer.Http.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to dial budget service: %v", err)
	}

	return &ServiceManager{
		userService:              userpb.NewUserServiceClient(connUser),
		IncomeAndExpensesService: incomepb.NewIncomeExpensesServiceClient(connincome),
		ReportingService:         reportingpb.NewReportingServiceClient(connReporting),
		BudgetService:            budgetpb.NewBudgetServiceClient(connBudget),
	}, nil
}

func (s *ServiceManager) User() userpb.UserServiceClient {
	return s.userService
}

func (s *ServiceManager) IncomeAndExpenses() incomepb.IncomeExpensesServiceClient {
	return s.IncomeAndExpensesService
}

func (s *ServiceManager) Reporting() reportingpb.ReportingServiceClient {
	return s.ReportingService
}

func (s *ServiceManager) Budget() budgetpb.BudgetServiceClient {
	return s.BudgetService
}