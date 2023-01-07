package postgresql_test

import (
	"context"
	"testing"

	"github.com/Kritsana135/assessment/domain"
	"github.com/Kritsana135/assessment/expense/repository/postgresql"
	"github.com/Kritsana135/assessment/misc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type ExpenseRepoSuite struct {
	suite.Suite

	db            *gorm.DB
	cleanupDocker func()
	expenseRepo   domain.ExpenseRepository
}

func (suite *ExpenseRepoSuite) SetupSuite() {
	suite.db, suite.cleanupDocker = misc.SetupGormWithDocker()
}

func (suite *ExpenseRepoSuite) TearDownSuite() {
	suite.cleanupDocker()
}

func (suite *ExpenseRepoSuite) SetupTest() {
	suite.expenseRepo = postgresql.NewExpenseRepo(suite.db)
	err := suite.db.AutoMigrate(&domain.ExpenseTable{})
	assert.NoError(suite.T(), err)
}

func (suite *ExpenseRepoSuite) TearDownTest() {
	err := suite.db.Migrator().DropTable(&domain.ExpenseTable{})
	assert.NoError(suite.T(), err)
}

func (suite *ExpenseRepoSuite) TestGetExpenses() {
	ctx := context.TODO()
	expenses := []domain.ExpenseTable{
		{
			Amount: 100,
		},
		{
			Amount: 200,
		},
		{
			Amount: 300,
		},
	}

	for _, expense := range expenses {
		err := suite.expenseRepo.Create(ctx, &expense)
		assert.NoError(suite.T(), err)
	}

	expensesFromRepo, err := suite.expenseRepo.GetExpenses(ctx)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), len(expenses), len(expensesFromRepo))
}

func (suite *ExpenseRepoSuite) TestUpdateExpense() {
	ctx := context.TODO()
	expense := domain.ExpenseTable{
		Amount: 100,
	}

	err := suite.expenseRepo.Create(ctx, &expense)
	assert.NoError(suite.T(), err)

	expense.Amount = 200
	err = suite.expenseRepo.UpdateExpense(ctx, uint64(expense.ID), &expense)
	assert.NoError(suite.T(), err)

	expenseFromRepo, err := suite.expenseRepo.GetExpensesById(ctx, uint64(expense.ID))
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expense.Amount, expenseFromRepo.Amount)

	// not found expense
	err = suite.expenseRepo.UpdateExpense(ctx, 999, &expense)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), gorm.ErrRecordNotFound, err)
}

func TestExpenseRepoSuite(t *testing.T) {
	suite.Run(t, new(ExpenseRepoSuite))
}
