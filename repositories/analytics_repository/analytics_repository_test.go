package analytics_repository

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/kelompok4-loyaltypointagent/backend/constant"
	"github.com/kelompok4-loyaltypointagent/backend/testhelper"
	"github.com/stretchr/testify/suite"
)

type analyticsRepositorySuite struct {
	suite.Suite
	mock       sqlmock.Sqlmock
	repository AnalyticsRepository
}

func (s *analyticsRepositorySuite) SetupSuite() {
	var conn *sql.DB
	conn, s.mock = testhelper.MockDB()
	gorm := testhelper.InitDB(conn)
	s.repository = NewAnalyticsRepository(gorm)
}

func TestAnalyticsRepositorySuite(t *testing.T) {
	suite.Run(t, new(analyticsRepositorySuite))
}

func (s *analyticsRepositorySuite) TestSalesCount() {
	query := regexp.QuoteMeta("SELECT COUNT(1) FROM transactions WHERE status=? AND type=? AND YEAR(created_at)=?")
	rows := sqlmock.NewRows([]string{"count"}).AddRow(10)
	s.mock.ExpectQuery(query).WithArgs(constant.TransactionStatusSuccess, constant.TransactionTypePurchase, 2022).WillReturnRows(rows)
	s.repository.SalesCount(2022)
}

func (s *analyticsRepositorySuite) TestIncome() {
	query := regexp.QuoteMeta("SELECT SUM(amount) FROM transactions WHERE status=? AND type=? AND YEAR(created_at)=?")
	rows := sqlmock.NewRows([]string{"count"}).AddRow(10)
	s.mock.ExpectQuery(query).WithArgs(constant.TransactionStatusSuccess, constant.TransactionTypePurchase, 2022).WillReturnRows(rows)
	s.repository.Income(2022)
}

func (s *analyticsRepositorySuite) TestTransactionsByMonth() {
	query := regexp.QuoteMeta("SELECT MONTH(created_at) AS month, COUNT(1) AS value FROM transactions WHERE status=? AND YEAR(created_at)=? GROUP BY MONTH(created_at) ORDER BY MONTH(created_at) ASC")
	rows := sqlmock.NewRows([]string{"month"}).AddRow(10)
	s.mock.ExpectQuery(query).WithArgs(constant.TransactionStatusSuccess, 2022).WillReturnRows(rows)
	s.repository.TransactionsByMonth(2022)
}

func (s *analyticsRepositorySuite) TestRecentTransactions() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("SELECT * FROM `transactions`")
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	s.mock.ExpectQuery(query).WillReturnRows(rows)

	tdID := uuid.NewString()
	query = regexp.QuoteMeta("SELECT * FROM `transaction_details`")
	rows = sqlmock.NewRows([]string{"id", "transaction_id"}).AddRow(tdID, id)
	s.mock.ExpectQuery(query).WithArgs(id).WillReturnRows(rows)

	s.repository.RecentTransactions(2022, 1)
}
