package v1

import (
	"api-gateway/api/handler/models"
	"api-gateway/protos/income"
	"strconv"

	"github.com/gin-gonic/gin"
)

// RegisterIncome godoc
// @Summary Register a new income transaction
// @Description Add a new income transaction to the system
// @Tags transactions
// @Accept json
// @Produce json
// @Param userId path int true "User ID"
// @Param income body models.RegisterIncomeReq true "Register Income data"
// @Success 200 {object} models.RegisterIncomeResp
// @Failure 400 {object} map[string]string "Bad request error"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /transactions/income/{userId} [post]
func (h *handlerV1) RegisterIncome(c *gin.Context) {
	userId := c.Param("userId")
	id, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user id"})
		return
	}
	req := models.RegisterIncomeReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	res, err := h.service.IncomeAndExpenses().LogIncome(c, &income.IncomeRequest{
		Amount:   req.Amount,
		UserId:   int32(id),
		Currency: req.Currency,
		Category: req.Category,
		Date:     req.Date,
	})
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to log income: " + err.Error()})
		return
	}
	c.JSON(200, models.RegisterIncomeResp{
		Message:       res.Message,
		TransactionId: res.TransactionId,
	})
}

// RegisterExpense godoc
// @Summary Register a new expense transaction
// @Description Add a new expense transaction to the system
// @Tags transactions
// @Accept json
// @Produce json
// @Param userId path int true "User ID"
// @Param expense body models.RegisterIncomeReq true "Register Expense data"
// @Success 200 {object} models.RegisterIncomeReq
// @Failure 400 {object} map[string]string "Bad request error"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /transactions/expense/{userId} [post]
func (h *handlerV1) RegisterExpense(c *gin.Context) {
	userId := c.Param("userId")
	id, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user id"})
		return
	}
	req := models.RegisterIncomeReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	res, err := h.service.IncomeAndExpenses().LogExpense(c, &income.ExpenseRequest{
		Amount:   req.Amount,
		UserId:   int32(id),
		Currency: req.Currency,
		Category: req.Category,
		Date:     req.Date,
	})
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to log expense: " + err.Error()})
		return
	}
	c.JSON(200, models.RegisterIncomeResp{
		Message:       res.Message,
		TransactionId: res.TransactionId,
	})
}

// GetTransactions godoc
// @Summary Retrieve transactions
// @Description Get the list of all transactions
// @Tags transactions
// @Accept  json
// @Produce json
// @Param userId path int true "User ID"
// @Success 200 {object} []models.GetTransactionsRes "List of transactions"
// @Failure 400 {object} map[string]string "Bad request error"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /transactions/{userId} [get]
func (h *handlerV1) GetTransactions(c *gin.Context) {
	userId := c.Param("userId")
	res, err := h.service.IncomeAndExpenses().GetTransactions(c, &income.GetTransactionsRequest{
		UserId: userId,
	})
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get transactions: " + err.Error()})
		return
	}
	var transactions []models.GetTransactionsRes
	for _, v := range res.Transactions {
		transactions = append(transactions, models.GetTransactionsRes{
			TransactionId: v.TransactionId,
			Type:          v.Type,
			Amount:        v.Amount,
			Currency:      v.Currency,
			Category:      v.Category,
			Date:          v.Date,
		})
	}
	c.JSON(200, transactions)
}
