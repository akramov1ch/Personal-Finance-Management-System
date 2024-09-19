package v1

import (
	"api-gateway/api/handler/models"
	"api-gateway/protos/reporting"

	"github.com/gin-gonic/gin"
)

// GetReporting godoc
// @Summary get reporting
// @Description Add a new expense reporting to the system
// @Tags reporting
// @Accept json
// @Produce json
// @Param userId path int true "User ID"
// @Success 200 {object} models.ReportingRes
// @Failure 400 {object} map[string]string "Bad request error"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /reports/income-expense/{userId} [get]
func (h *handlerV1) GetRepoting(c *gin.Context) {
	id := c.Param("userId")
	res, err := h.service.Reporting().GetIncomeExpenseReport(c, &reporting.GetIncomeExpenseReportRequest{
		UserId: id,
	})
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get reporting: " + err.Error()})
		return
	}

	c.JSON(200, models.ReportingRes{
		TotalIncome:   int64(res.TotalIncome),
		TotalExpenses: int64(res.TotalExpenses),
		NetSavings:    int64(res.NetSavings),
	})
}

// GetReportingByCategory godoc
// @Summary get reporting By Category
// @Description Add a new expense reporting to the system
// @Tags reporting
// @Accept json
// @Produce json
// @Param userId path int true "User ID"
// @Success 200 {object} []models.CategorySpending
// @Failure 400 {object} map[string]string "Bad request error"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /reports/spending-by-category/{userId} [get]
func (h *handlerV1) GetRepotingByCategory(c *gin.Context) {
	id := c.Param("userId")
	res, err := h.service.Reporting().GetSpendingByCategoryReport(c, &reporting.GetSpendingByCategoryReportRequest{
		UserId: id,
	})
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get reporting: " + err.Error()})
		return
	}
	categories := []models.CategorySpending{}
	for _, v := range res.Spending {
		categories = append(categories, models.CategorySpending{
			Category:   v.Category,
			TotalSpent: int64(v.TotalSpent),
		})
	}
	c.JSON(200, categories)
}
