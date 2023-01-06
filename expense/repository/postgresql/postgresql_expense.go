package postgresql

import (
	"context"

	"github.com/Kritsana135/assessment/domain"
	"gorm.io/gorm"
)

type expenseRepo struct {
	db *gorm.DB
}

func (e *expenseRepo) Create(ctx context.Context, expense *domain.ExpenseTable) error {
	db := e.db.WithContext(ctx)

	return db.Create(expense).Error
}

func NewExpenseRepo(db *gorm.DB) domain.ExpenseRepository {
	return &expenseRepo{
		db: db,
	}
}
