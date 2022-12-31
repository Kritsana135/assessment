package http_

import (
	"net/http"

	"github.com/Kritsana135/assessment/domain"
	"github.com/gin-gonic/gin"
)

type ExpenseHandler struct {
}

func NewExpenseHandler(r *gin.RouterGroup) {
	handler := &ExpenseHandler{}

	er := r.Group("/expenses")

	er.POST("", handler.CreateExpense)
	er.GET("/:id", handler.GetExpensesById)
	er.PUT("/:id", handler.UpdateExpense)
	er.GET("", handler.GetExpenses)
}

func (h *ExpenseHandler) CreateExpense(c *gin.Context) {
	var body domain.CrateExpenseReq

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, domain.BaseResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (h *ExpenseHandler) GetExpenses(c *gin.Context) {}

func (h *ExpenseHandler) GetExpensesById(c *gin.Context) {}

func (h *ExpenseHandler) UpdateExpense(c *gin.Context) {}
