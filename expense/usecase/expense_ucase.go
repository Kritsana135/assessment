package usecase

import (
	"context"

	"github.com/Kritsana135/assessment/domain"
)

type expenseUsecase struct {
}

// CreateExpense implements domain.ExpenseUseCase
func (*expenseUsecase) CreateExpense(ctx context.Context, req domain.CrateExpenseReq) (domain.ExpenseTable, error) {
	panic("unimplemented")
}

func NewExpUsecase() domain.ExpenseUseCase {
	return &expenseUsecase{}
}
