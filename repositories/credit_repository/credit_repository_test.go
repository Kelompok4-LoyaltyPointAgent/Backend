package credit_repository

import (
	"database/sql"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/kelompok4-loyaltypointagent/backend/models"
	"github.com/kelompok4-loyaltypointagent/backend/testhelper"
	"github.com/stretchr/testify/suite"
)

type creditRepositorySuite struct {
	suite.Suite
	mock       sqlmock.Sqlmock
	repository CreditRepository
}

func (s *creditRepositorySuite) SetupSuite() {
	var conn *sql.DB
	conn, s.mock = testhelper.MockDB()
	gorm := testhelper.InitDB(conn)
	s.repository = NewCreditRepository(gorm)
}

func TestCreditRepositorySuite(t *testing.T) {
	suite.Run(t, new(creditRepositorySuite))
}

func (s *creditRepositorySuite) TestFindAll() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("SELECT * FROM `credits`")
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	s.mock.ExpectQuery(query).WillReturnRows(rows)
	s.repository.FindAll()
}

func (s *creditRepositorySuite) TestCreate() {
	query := regexp.QuoteMeta("INSERT INTO `credits`")

	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	s.mock.ExpectClose()
	s.repository.Create(models.Credit{})

	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnError(errors.New("error"))
	s.mock.ExpectCommit()
	s.repository.Create(models.Credit{})
	s.mock.ExpectClose()
}

func (s *creditRepositorySuite) TestFindByProductID() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("SELECT * FROM `credits`")
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	s.mock.ExpectQuery(query).WithArgs(id).WillReturnRows(rows)
	s.repository.FindByProductID(id)
}

func (s *creditRepositorySuite) TestFindByProvider() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("SELECT * FROM `credits`")
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	s.mock.ExpectQuery(query).WillReturnRows(rows)
	s.repository.FindByProvider("any")
}

func (s *creditRepositorySuite) TestFindByRecommended() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("SELECT * FROM `credits`")
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	s.mock.ExpectQuery(query).WillReturnRows(rows)
	s.repository.FindByRecommended()
}

func (s *creditRepositorySuite) TestUpdateByProductID() {
	s.mock.ExpectBegin()
	query := regexp.QuoteMeta("UPDATE `credits`")
	s.mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	id := uuid.NewString()
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	query = regexp.QuoteMeta("SELECT * FROM `credits`")
	s.mock.ExpectQuery(query).WithArgs(id).WillReturnRows(rows)

	s.repository.UpdateByProductID(models.Credit{}, id)
}

func (s *creditRepositorySuite) TestDeleteByProductID() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("UPDATE `credits`")
	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	s.repository.DeleteByProductID(id)
}
