package postgresql

import (
	"github.com/Kritsana135/assessment/domain"
	"github.com/jmoiron/sqlx"
)

type expenseRepo struct {
	db *sqlx.DB
}

func (*expenseRepo) Create(expense *domain.ExpenseTable) error {
	panic("unimplemented")
}

func NewExpenseRepo(db *sqlx.DB) domain.ExpenseRepository {
	return &expenseRepo{
		db: db,
	}
}
