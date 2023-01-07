package postgresql_test

import (
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

func TestExpenseRepoSuite(t *testing.T) {
	suite.Run(t, new(ExpenseRepoSuite))
}
