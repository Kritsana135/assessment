package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/Kritsana135/assessment/domain"
	"github.com/Kritsana135/assessment/domain/apperrors"
	"github.com/Kritsana135/assessment/domain/mocks"
	"github.com/Kritsana135/assessment/expense/usecase"
	"github.com/stretchr/testify/assert"
)

func TestGetExpenses(t *testing.T) {
	ctx := context.Background()
	t.Run("geu_1: error when get", func(t *testing.T) {
		mockExpenseRepo := mocks.NewExpenseRepository(t)
		mockExpenseRepo.On("GetExpenses", ctx).Return([]domain.ExpenseTable{}, errors.New("error"))
		expUCase := usecase.NewExpUsecase(mockExpenseRepo)

		_, err := expUCase.GetExpenses(ctx)

		assert.Error(t, err)
		assert.Equal(t, apperrors.NewInternal().Message, err.Error())
	})

	t.Run("geu_2: success", func(t *testing.T) {
		mockExpenseRepo := mocks.NewExpenseRepository(t)
		mockExpenseRepo.On("GetExpenses", ctx).Return([]domain.ExpenseTable{}, nil)
		expUCase := usecase.NewExpUsecase(mockExpenseRepo)

		_, err := expUCase.GetExpenses(ctx)

		assert.NoError(t, err)
	})
}
