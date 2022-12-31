package domain

import "context"

type (
	ExpenseTable struct {
		ID     int      `json:"id" db:"id"`
		Title  string   `json:"title" db:"title"`
		Amount float64  `json:"amount" db:"amount"`
		Note   string   `json:"note" db:"note"`
		Tags   []string `json:"tags" db:"tags"`
	}
	ExpenseRepository interface {
		Create(expense *ExpenseTable) error
	}
	ExpenseUseCase interface {
		CreateExpense(ctx context.Context, req CrateExpenseReq) (ExpenseTable, error)
	}
	CrateExpenseReq struct {
		Title  string   `json:"title"`
		Amount float64  `json:"amount"`
		Note   string   `json:"note"`
		Tags   []string `json:"tags"`
	}
)
