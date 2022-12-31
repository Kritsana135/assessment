package domain

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
	ExpenseUseCase interface {
	}
	CrateExpenseReq struct {
		Title  string   `json:"title"`
		Amount float64  `json:"amount"`
		Note   string   `json:"note"`
		Tags   []string `json:"tags"`
	}
)
