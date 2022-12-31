package http_

import (
	"net/http"

	"github.com/Kritsana135/assessment/domain"
	"github.com/Kritsana135/assessment/domain/apperrors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ExpenseHandler struct {
	expUCase domain.ExpenseUseCase
}

func NewExpenseHandler(r *gin.RouterGroup, expUCase domain.ExpenseUseCase) {
	handler := &ExpenseHandler{
		expUCase,
	}

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

	res, err := h.expUCase.CreateExpense(c.Request.Context(), body)
	if err != nil {
		logrus.Error(err)
		c.JSON(apperrors.Status(err), domain.BaseResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *ExpenseHandler) GetExpenses(c *gin.Context) {}

func (h *ExpenseHandler) GetExpensesById(c *gin.Context) {}

func (h *ExpenseHandler) UpdateExpense(c *gin.Context) {}
