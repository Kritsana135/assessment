package http_

import (
	"net/http"
	"strconv"

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
	{
		er.POST("", handler.CreateExpense)
		er.GET("/:id", handler.GetExpensesById)
		er.PUT("/:id", handler.UpdateExpense)
		er.GET("", handler.GetExpenses)
	}

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

	c.JSON(http.StatusCreated, res)
}

func (h *ExpenseHandler) GetExpenses(ctx *gin.Context) {

}

func (h *ExpenseHandler) GetExpensesById(ctx *gin.Context) {
	id := ctx.Param("id")
	uId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusBadRequest, domain.BaseResponse{Message: "invalid id"})
		return
	}

	res, err := h.expUCase.GetExpenses(ctx.Request.Context(), uId)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(apperrors.Status(err), domain.BaseResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *ExpenseHandler) UpdateExpense(ctx *gin.Context) {
	id := ctx.Param("id")
	uId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusBadRequest, domain.BaseResponse{Message: "invalid id"})
		return
	}

	var body domain.UpdateExpenseReq
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, domain.BaseResponse{Message: err.Error()})
		return
	}

	res, err := h.expUCase.UpdateExpense(ctx.Request.Context(), uId, body)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(apperrors.Status(err), domain.BaseResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
