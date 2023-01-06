package domain

import (
	"context"

	"github.com/lib/pq"
)

type (
	ExpenseTable struct {
		ID     int            `json:"id" gorm:"column:id"`
		Title  string         `json:"title" gorm:"column:title"`
		Amount float64        `json:"amount" gorm:"column:amount; type:float"`
		Note   string         `json:"note" gorm:"column:note"`
		Tags   pq.StringArray `json:"tags" gorm:"column:tags; type:text[]"`
	}
	ExpenseRepository interface {
		Create(ctx context.Context, expense *ExpenseTable) error
	}
	ExpenseUseCase interface {
		CreateExpense(ctx context.Context, req CrateExpenseReq) (ExpenseTable, error)
	}
	CrateExpenseReq struct {
		Title  string   `json:"title" binding:"required"`
		Amount float64  `json:"amount" binding:"required"`
		Note   string   `json:"note"`
		Tags   []string `json:"tags"`
	}
)

func (u ExpenseTable) TableName() string {
	return "expenses"
}
