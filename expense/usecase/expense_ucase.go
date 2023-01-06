package usecase

import (
	"context"
	"fmt"

	"github.com/Kritsana135/assessment/domain"
	"github.com/Kritsana135/assessment/domain/apperrors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type expenseUsecase struct {
	expenseRepo domain.ExpenseRepository
}

// GetExpenses implements domain.ExpenseUseCase
func (e *expenseUsecase) GetExpenses(ctx context.Context) ([]domain.ExpenseTable, error) {
	expenses, err := e.expenseRepo.GetExpenses(ctx)
	if err != nil {
		logrus.Error(err)
		return nil, apperrors.NewInternal()
	}

	return expenses, nil
}

// UpdateExpense implements domain.ExpenseUseCase
func (e *expenseUsecase) UpdateExpense(ctx context.Context, id uint64, req domain.UpdateExpenseReq) (domain.ExpenseTable, error) {
	expense := domain.ExpenseTable{
		ID:     int(id),
		Title:  req.Title,
		Amount: req.Amount,
		Note:   req.Note,
		Tags:   req.Tags,
	}

	err := e.expenseRepo.UpdateExpense(ctx, id, &expense)
	if err != nil {
		logrus.Error(err)
		if err == gorm.ErrRecordNotFound {
			return expense, apperrors.NewNotFound("expense", fmt.Sprint(id))
		}
		return expense, apperrors.NewInternal()
	}

	return expense, err
}

// GetExpensesById implements domain.ExpenseUseCase
func (e *expenseUsecase) GetExpensesById(ctx context.Context, id uint64) (domain.ExpenseTable, error) {
	expense, err := e.expenseRepo.GetExpensesById(ctx, id)
	if err != nil {
		logrus.Error(err)
		if err == gorm.ErrRecordNotFound {
			return expense, apperrors.NewNotFound("expense", fmt.Sprint(id))
		}
		return expense, apperrors.NewInternal()
	}

	return expense, err
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
