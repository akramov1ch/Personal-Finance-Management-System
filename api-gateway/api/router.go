package api

import (
	v1 "api-gateway/api/handler/v1"
	"api-gateway/service"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Options struct {
	Service service.IClients
}

func Router(opt Options) *gin.Engine {
	r := gin.Default()

	h := v1.NewHandlerV1(&v1.HandlerConfig{Service: opt.Service})
	// add swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{
		v1.POST("/transactions/income/:userId", h.RegisterIncome)
		v1.POST("/transactions/expense/:userId", h.RegisterExpense)
		v1.GET("/transactions/:userId", h.GetTransactions)

		v1.POST("/budgets/:userId", h.CreateBudgetForCategory)
		v1.GET("/budgets/:userId", h.GetBudgets)
		v1.PUT("/budgets/:budgetId", h.PutBudget)

		v1.POST("/users", h.Register)
		v1.POST("/users/verify", h.VerifyUser)
		v1.POST("/users/login", h.Login)
		v1.GET("/users/profile/:userId", h.GetProfile)

		v1.GET("/reports/income-expense/:userId", h.GetRepoting)
		v1.GET("/reports/spending-by-category/:userId", h.GetRepotingByCategory)
	}

	return r
}
