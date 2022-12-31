package model

type (
	ExpenseTable struct {
		ID     int      `db:"id"`
		Title  string   `db:"title"`
		Amount float64  `db:"amount"`
		Note   string   `db:"note"`
		Tags   []string `db:"tags"`
	}
	ExpenseRepository interface {
		Create(expense *ExpenseTable) error
	}
)
