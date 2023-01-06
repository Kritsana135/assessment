package postgresql

import (
	"context"

	"github.com/Kritsana135/assessment/domain"
	"gorm.io/gorm"
)

type expenseRepo struct {
	db *gorm.DB
}

// GetExpenses implements domain.ExpenseRepository
func (e *expenseRepo) GetExpenses(ctx context.Context) ([]domain.ExpenseTable, error) {
	db := e.db.WithContext(ctx)

	var expenses []domain.ExpenseTable
	err := db.Table("expenses").Find(&expenses).Error

	return expenses, err
}

// UpdateExpense implements domain.ExpenseRepository
func (e *expenseRepo) UpdateExpense(ctx context.Context, id uint64, expense *domain.ExpenseTable) error {
	db := e.db.WithContext(ctx)

	tx := db.Table("expenses").Where("id = ?", id).Updates(expense)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// GetExpensesById implements domain.ExpenseRepository
func (e *expenseRepo) GetExpensesById(ctx context.Context, id uint64) (domain.ExpenseTable, error) {
	db := e.db.WithContext(ctx)

	var expense domain.ExpenseTable
	err := db.Table("expenses").Where("id = ?", id).First(&expense).Error

	return expense, err
}

func (e *expenseRepo) Create(ctx context.Context, expense *domain.ExpenseTable) error {
	db := e.db.WithContext(ctx)

	return db.Table("expenses").Create(expense).Error
}

func NewExpenseRepo(db *gorm.DB) domain.ExpenseRepository {
	return &expenseRepo{
		db: db,
	}
}
