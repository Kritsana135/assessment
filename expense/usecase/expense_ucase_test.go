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
	"gorm.io/gorm"
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
		mockExpenseRepo.On("GetExpenses", ctx).Return([]domain.ExpenseTable{
			{ID: 1, Title: "test", Amount: 100, Note: "test"},
		}, nil)
		expUCase := usecase.NewExpUsecase(mockExpenseRepo)

		expense, err := expUCase.GetExpenses(ctx)

		assert.NoError(t, err)
		assert.Equal(t, int(1), expense[0].ID)
	})
}

func TestCreateExpense(t *testing.T) {
	ctx := context.Background()
	t.Run("ce_1: error when create", func(t *testing.T) {
		mockExpenseRepo := mocks.NewExpenseRepository(t)
		mockExpenseRepo.On("Create", ctx, &domain.ExpenseTable{Title: "test", Amount: 100, Note: "test"}).Return(errors.New("error"))
		expUCase := usecase.NewExpUsecase(mockExpenseRepo)

		_, err := expUCase.CreateExpense(ctx, domain.CreateExpenseReq{Title: "test", Amount: 100, Note: "test"})

		assert.Error(t, err)
		assert.Equal(t, apperrors.NewInternal().Message, err.Error())
	})

	t.Run("ce_2: success", func(t *testing.T) {
		mockExpenseRepo := mocks.NewExpenseRepository(t)
		mockExpenseRepo.On("Create", ctx, &domain.ExpenseTable{Title: "test", Amount: 100, Note: "test"}).Return(nil)
		expUCase := usecase.NewExpUsecase(mockExpenseRepo)

		_, err := expUCase.CreateExpense(ctx, domain.CreateExpenseReq{Title: "test", Amount: 100, Note: "test"})

		assert.NoError(t, err)
	})
}

func TestGetExpensesById(t *testing.T) {
	ctx := context.Background()
	t.Run("geu_1: error when get", func(t *testing.T) {
		mockExpenseRepo := mocks.NewExpenseRepository(t)
		mockExpenseRepo.On("GetExpensesById", ctx, uint64(1)).Return(domain.ExpenseTable{}, errors.New("error"))
		expUCase := usecase.NewExpUsecase(mockExpenseRepo)

		_, err := expUCase.GetExpensesById(ctx, 1)

		assert.Error(t, err)
		assert.Equal(t, apperrors.NewInternal().Message, err.Error())
	})

	t.Run("geu_2: error not found", func(t *testing.T) {
		mockExpenseRepo := mocks.NewExpenseRepository(t)
		mockExpenseRepo.On("GetExpensesById", ctx, uint64(1)).Return(domain.ExpenseTable{}, gorm.ErrRecordNotFound)
		expUCase := usecase.NewExpUsecase(mockExpenseRepo)

		_, err := expUCase.GetExpensesById(ctx, 1)

		assert.Error(t, err)
		assert.Equal(t, 404, apperrors.Status(err))
	})

	t.Run("geu_3: success", func(t *testing.T) {
		mockExpenseRepo := mocks.NewExpenseRepository(t)
		mockExpenseRepo.On("GetExpensesById", ctx, uint64(1)).Return(domain.ExpenseTable{}, nil)
		expUCase := usecase.NewExpUsecase(mockExpenseRepo)

		_, err := expUCase.GetExpensesById(ctx, 1)

		assert.NoError(t, err)
	})
}

func TestUpdateExpense(t *testing.T) {
	ctx := context.Background()
	t.Run("ue_1: error when update", func(t *testing.T) {
		mockExpenseRepo := mocks.NewExpenseRepository(t)
		mockExpenseRepo.On("UpdateExpense", ctx, uint64(1), &domain.ExpenseTable{Title: "test", Amount: 100, Note: "test", ID: 1}).
			Return(errors.New("error"))
		expUCase := usecase.NewExpUsecase(mockExpenseRepo)

		_, err := expUCase.UpdateExpense(ctx, 1, domain.UpdateExpenseReq{Title: "test", Amount: 100, Note: "test"})

		assert.Error(t, err)
		assert.Equal(t, apperrors.NewInternal().Message, err.Error())
	})

	t.Run("ue_2: error not found", func(t *testing.T) {
		mockExpenseRepo := mocks.NewExpenseRepository(t)
		mockExpenseRepo.On("UpdateExpense", ctx, uint64(1), &domain.ExpenseTable{ID: 1, Title: "test", Amount: 100, Note: "test"}).
			Return(gorm.ErrRecordNotFound)
		expUCase := usecase.NewExpUsecase(mockExpenseRepo)

		_, err := expUCase.UpdateExpense(ctx, 1, domain.UpdateExpenseReq{Title: "test", Amount: 100, Note: "test"})

		assert.Error(t, err)
		assert.Equal(t, 404, apperrors.Status(err))
	})

	t.Run("ue_3: success", func(t *testing.T) {
		mockExpenseRepo := mocks.NewExpenseRepository(t)
		mockExpenseRepo.On("UpdateExpense", ctx, uint64(1), &domain.ExpenseTable{ID: 1, Title: "test", Amount: 100, Note: "test"}).
			Return(nil)
		expUCase := usecase.NewExpUsecase(mockExpenseRepo)

		_, err := expUCase.UpdateExpense(ctx, 1, domain.UpdateExpenseReq{Title: "test", Amount: 100, Note: "test"})

		assert.NoError(t, err)
	})
}
