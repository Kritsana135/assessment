package http__test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Kritsana135/assessment/domain"
	"github.com/Kritsana135/assessment/domain/apperrors"
	"github.com/Kritsana135/assessment/domain/mocks"
	"github.com/Kritsana135/assessment/expense/delivery/http_"
	"github.com/Kritsana135/assessment/misc"
	"github.com/stretchr/testify/assert"
)

func TestCreateExpense(t *testing.T) {
	t.Run("ce_1:the request body failed because the required field is missing.", func(t *testing.T) {
		bodies := []string{
			`{
				"amount": 79,
				"note": "night market promotion discount 10 bath",
				"tags": ["food", "beverage"]
			}`,
			`{
				"title": "night market",
				"note": "night market promotion discount 10 bath",
				"tags": ["food", "beverage"]
			}`,
		}

		for _, body := range bodies {
			w := httptest.NewRecorder()
			ctx := misc.GetTestGinContext(w)
			ctx.Request.Header.Set("Content-Type", "application/json")

			ctx.Request.Body = io.NopCloser(bytes.NewBuffer([]byte(body)))

			handler := http_.ExpenseHandler{}

			handler.CreateExpense(ctx)

			assert.Equal(t, http.StatusBadRequest, w.Code)
		}
	})

	t.Run("ce_2:should get error when create", func(t *testing.T) {
		body := `{
				  "title": "night market",
				  "amount": 79,
				  "note": "night market promotion discount 10 bath",
				  "tags": ["food", "beverage"]
			     }`

		w := httptest.NewRecorder()
		ctx := misc.GetTestGinContext(w)
		ctx.Request.Header.Set("Content-Type", "application/json")

		ctx.Request.Body = io.NopCloser(bytes.NewBuffer([]byte(body)))

		mockExpenseUseCase := mocks.NewExpenseUseCase(t)
		parseBody := domain.CrateExpenseReq{
			Title:  "night market",
			Amount: 79,
			Note:   "night market promotion discount 10 bath",
			Tags:   []string{"food", "beverage"},
		}
		mockExpenseUseCase.On("CreateExpense", ctx.Request.Context(), parseBody).
			Return(domain.ExpenseTable{}, apperrors.NewInternal())
		handler := http_.ExpenseHandler{
			ExpUCase: mockExpenseUseCase,
		}

		handler.CreateExpense(ctx)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("ce_3:should get 201 status", func(t *testing.T) {
		body := `{
				  "title": "night market",
				  "amount": 79,
				  "note": "night market promotion discount 10 bath",
				  "tags": ["food", "beverage"]
			     }`

		w := httptest.NewRecorder()
		ctx := misc.GetTestGinContext(w)
		ctx.Request.Header.Set("Content-Type", "application/json")

		ctx.Request.Body = io.NopCloser(bytes.NewBuffer([]byte(body)))

		mockExpenseUseCase := mocks.NewExpenseUseCase(t)
		parseBody := domain.CrateExpenseReq{
			Title:  "night market",
			Amount: 79,
			Note:   "night market promotion discount 10 bath",
			Tags:   []string{"food", "beverage"},
		}
		mockExpenseUseCase.On("CreateExpense", ctx.Request.Context(), parseBody).
			Return(domain.ExpenseTable{Title: "night market"}, nil)
		handler := http_.ExpenseHandler{
			ExpUCase: mockExpenseUseCase,
		}

		handler.CreateExpense(ctx)

		var res domain.ExpenseTable
		json.NewDecoder(w.Body).Decode(&res)

		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Equal(t, "night market", res.Title)
	})
}
