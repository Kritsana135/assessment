// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/Kritsana135/assessment/domain"
	mock "github.com/stretchr/testify/mock"
)

// ExpenseUseCase is an autogenerated mock type for the ExpenseUseCase type
type ExpenseUseCase struct {
	mock.Mock
}

// CreateExpense provides a mock function with given fields: ctx, req
func (_m *ExpenseUseCase) CreateExpense(ctx context.Context, req domain.CreateExpenseReq) (domain.ExpenseTable, error) {
	ret := _m.Called(ctx, req)

	var r0 domain.ExpenseTable
	if rf, ok := ret.Get(0).(func(context.Context, domain.CreateExpenseReq) domain.ExpenseTable); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(domain.ExpenseTable)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.CreateExpenseReq) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetExpenses provides a mock function with given fields: ctx
func (_m *ExpenseUseCase) GetExpenses(ctx context.Context) ([]domain.ExpenseTable, error) {
	ret := _m.Called(ctx)

	var r0 []domain.ExpenseTable
	if rf, ok := ret.Get(0).(func(context.Context) []domain.ExpenseTable); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.ExpenseTable)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetExpensesById provides a mock function with given fields: ctx, id
func (_m *ExpenseUseCase) GetExpensesById(ctx context.Context, id uint64) (domain.ExpenseTable, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.ExpenseTable
	if rf, ok := ret.Get(0).(func(context.Context, uint64) domain.ExpenseTable); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.ExpenseTable)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateExpense provides a mock function with given fields: ctx, id, req
func (_m *ExpenseUseCase) UpdateExpense(ctx context.Context, id uint64, req domain.UpdateExpenseReq) (domain.ExpenseTable, error) {
	ret := _m.Called(ctx, id, req)

	var r0 domain.ExpenseTable
	if rf, ok := ret.Get(0).(func(context.Context, uint64, domain.UpdateExpenseReq) domain.ExpenseTable); ok {
		r0 = rf(ctx, id, req)
	} else {
		r0 = ret.Get(0).(domain.ExpenseTable)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64, domain.UpdateExpenseReq) error); ok {
		r1 = rf(ctx, id, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewExpenseUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewExpenseUseCase creates a new instance of ExpenseUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewExpenseUseCase(t mockConstructorTestingTNewExpenseUseCase) *ExpenseUseCase {
	mock := &ExpenseUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
