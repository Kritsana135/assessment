package repositories

import (
	"github.com/Kritsana135/assessment/model"
	"github.com/jmoiron/sqlx"
)

type expenseRepo struct {
	db *sqlx.DB
}

func (*expenseRepo) Create(expense *model.ExpenseTable) error {
	panic("unimplemented")
}

func NewExpenseRepo(db *sqlx.DB) model.ExpenseRepository {
	return &expenseRepo{
		db: db,
	}
}
