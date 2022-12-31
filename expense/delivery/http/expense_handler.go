package http

import "github.com/gin-gonic/gin"

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

func (h *ExpenseHandler) CreateExpense(c *gin.Context) {}

func (h *ExpenseHandler) GetExpenses(c *gin.Context) {}

func (h *ExpenseHandler) GetExpensesById(c *gin.Context) {}

func (h *ExpenseHandler) UpdateExpense(c *gin.Context) {}
