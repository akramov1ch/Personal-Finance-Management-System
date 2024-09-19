package v1

import (
	"api-gateway/api/handler/models"
	"api-gateway/protos/budget"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateBudgetForCategory godoc
// @Summary Register a new Budget budget
// @Description Add a new expense budget to the system
// @Tags budget
// @Accept json
// @Produce json
// @Param userId path int true "User ID"
// @Param expense body models.CreateBudgetForCategoryReq true "Create Budget data"
// @Success 200 {object} models.CreateBudgetForCategoryResp
// @Failure 400 {object} map[string]string "Bad request error"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /budgets/{userId} [post]
func (h *handlerV1) CreateBudgetForCategory(c *gin.Context) {
	userId := c.Param("userId")
	req := models.CreateBudgetForCategoryReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := h.service.Budget().CreateBudget(c, &budget.CreateBudgetRequest{
		UserId:   userId,
		Category: req.Category,
		Amount:   req.Amount,
		Currency: req.Currency,
		Spent:    req.Spent,
		Date:     req.Date,
	})
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create budget: " + err.Error()})
		return
	}
	c.JSON(200, models.CreateBudgetForCategoryResp{
		Message: res.Message,
		BudgeId: res.BudgetId,
	})
}

// GetBudgets godoc
// @Summary Get Budgets
// @Description Add a new expense budget to the system
// @Tags budget
// @Accept json
// @Produce json
// @Param userId path int true "User ID"
// @Success 200 {object} []models.GetBudgetsResp
// @Failure 400 {object} map[string]string "Bad request error"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /budgets/{userId} [get]
func (h *handlerV1) GetBudgets(c *gin.Context) {
	userId := c.Param("userId")
	res, err := h.service.Budget().GetBudgets(c, &budget.GetBudgetsRequest{
		UserId: userId,
	})
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get budgets: " + err.Error()})
		return
	}
	budgets := []models.GetBudgetsResp{}
	for _, v := range res.Budgets {
		budgets = append(budgets, models.GetBudgetsResp{
			BudgetId: v.BudgetId,
			Category: v.Category,
			Amount:   v.Amount,
			Spent:    int(v.Spent),
			Currency: v.Currency,
		})
	}
	c.JSON(200, budgets)
}

// PutBudget godoc
// @Summary Put Budgets
// @Description Add a new expense budget to the system
// @Tags budget
// @Accept json
// @Produce json
// @Param budgetId path int true "Budget ID"
// @Param expense body models.PutBudgetReq true "Put Budget data"
// @Success 200 {object} models.PutBudgetResp
// @Failure 400 {object} map[string]string "Bad request error"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /budgets/{budgetId} [put]
func (h *handlerV1) PutBudget(c *gin.Context) {
	sid := c.Param("budgetId")
	id, err := strconv.Atoi(sid)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid budget id"})
		return
	}
	req := models.PutBudgetReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	res, err := h.service.Budget().PutBudget(c, &budget.PatchBudgetRequest{
		BudgetId: int32(id),
		Amount:   req.Amount,
	})
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update budget: " + err.Error()})
		return
	}
	c.JSON(200, models.PutBudgetResp{
		Message: res.Message,
	})
}
