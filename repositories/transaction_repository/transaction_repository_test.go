package transaction_repository

import (
	"database/sql"
	"errors"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/kelompok4-loyaltypointagent/backend/models"
	"github.com/kelompok4-loyaltypointagent/backend/testhelper"
	"github.com/stretchr/testify/suite"
)

type transactionRepositorySuite struct {
	suite.Suite
	mock       sqlmock.Sqlmock
	repository TransactionRepository
}

func (s *transactionRepositorySuite) SetupSuite() {
	var conn *sql.DB
	conn, s.mock = testhelper.MockDB()
	gorm := testhelper.InitDB(conn)
	s.repository = NewTransactionRepository(gorm)
}

func TestTransactionRepositorySuite(t *testing.T) {
	suite.Run(t, new(transactionRepositorySuite))
}

func (s *transactionRepositorySuite) TestFindAll() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("SELECT * FROM `transactions`")
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	s.mock.ExpectQuery(query).WillReturnRows(rows)
	if _, err := s.repository.FindAll(nil, nil); err != nil {
		log.Println(err)
	}
}

func (s *transactionRepositorySuite) TestFindByID() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("SELECT * FROM `transactions`")
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	s.mock.ExpectQuery(query).WithArgs(id).WillReturnRows(rows)

	tdID := uuid.NewString()
	query = regexp.QuoteMeta("SELECT * FROM `transaction_details`")
	rows = sqlmock.NewRows([]string{"id", "transaction_id"}).AddRow(tdID, id)
	s.mock.ExpectQuery(query).WithArgs(id).WillReturnRows(rows)

	if _, err := s.repository.FindByID(id); err != nil {
		log.Println(err)
	}
}

func (s *transactionRepositorySuite) TestCreate() {
	query := regexp.QuoteMeta("INSERT INTO `transactions`")

	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	s.mock.ExpectClose()
	if _, err := s.repository.Create(models.Transaction{}); err != nil {
		log.Println(err)
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnError(errors.New("error"))
	s.mock.ExpectCommit()
	if _, err := s.repository.Create(models.Transaction{}); err != nil {
		log.Println(err)
	}
	s.mock.ExpectClose()
}

func (s *transactionRepositorySuite) TestUpdate() {
	s.mock.ExpectBegin()
	query := regexp.QuoteMeta("UPDATE `transactions`")
	s.mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	id := uuid.NewString()
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	query = regexp.QuoteMeta("SELECT * FROM `transactions`")
	s.mock.ExpectQuery(query).WithArgs(id).WillReturnRows(rows)

	tdID := uuid.NewString()
	query = regexp.QuoteMeta("SELECT * FROM `transaction_details`")
	rows = sqlmock.NewRows([]string{"id", "transaction_id"}).AddRow(tdID, id)
	s.mock.ExpectQuery(query).WithArgs(id).WillReturnRows(rows)

	if _, err := s.repository.Update(models.Transaction{}, id); err != nil {
		log.Println(err)
	}
}

func (s *transactionRepositorySuite) TestDelete() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("UPDATE `transactions`")
	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	if err := s.repository.Delete(id); err != nil {
		log.Println(err)
	}
}

func (s *transactionRepositorySuite) TestCreateDetail() {
	query := regexp.QuoteMeta("INSERT INTO `transaction_details`")

	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	s.mock.ExpectClose()
	if _, err := s.repository.CreateDetail(models.TransactionDetail{}); err != nil {
		log.Println(err)
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnError(errors.New("error"))
	s.mock.ExpectCommit()
	if _, err := s.repository.CreateDetail(models.TransactionDetail{}); err != nil {
		log.Println(err)
	}
	s.mock.ExpectClose()
}

func (s *transactionRepositorySuite) TestDeleteDetailByTransactionID() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("UPDATE `transaction_details`")
	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	if err := s.repository.DeleteDetailByTransactionID(id); err != nil {
		log.Println(err)
	}
}
