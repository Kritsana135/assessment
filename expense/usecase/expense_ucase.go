package usecase

import (
	"context"

	"github.com/Kritsana135/assessment/domain"
	"github.com/Kritsana135/assessment/domain/apperrors"
	"github.com/sirupsen/logrus"
)

type expenseUsecase struct {
	expenseRepo domain.ExpenseRepository
}

// CreateExpense implements domain.ExpenseUseCase
func (e *expenseUsecase) CreateExpense(ctx context.Context, req domain.CrateExpenseReq) (domain.ExpenseTable, error) {
	expense := domain.ExpenseTable{
		Title:  req.Title,
		Amount: req.Amount,
		Note:   req.Note,
		Tags:   req.Tags,
	}

	err := e.expenseRepo.Create(ctx, &expense)
	if err != nil {
		logrus.Error(err)
		return expense, apperrors.NewInternal()
	}

	return expense, err
}

func NewExpUsecase(expenseRepo domain.ExpenseRepository) domain.ExpenseUseCase {
	return &expenseUsecase{
		expenseRepo,
	}
}
