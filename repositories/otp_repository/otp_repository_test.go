package otp_repository

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

type otpRepositorySuite struct {
	suite.Suite
	mock       sqlmock.Sqlmock
	repository OTPRepository
}

func (s *otpRepositorySuite) SetupSuite() {
	var conn *sql.DB
	conn, s.mock = testhelper.MockDB()
	gorm := testhelper.InitDB(conn)
	s.repository = NewOTPRepository(gorm)
}

func TestOTPRepositorySuite(t *testing.T) {
	suite.Run(t, new(otpRepositorySuite))
}

func (s *otpRepositorySuite) TestFindByID() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("SELECT * FROM `one_time_passwords`")
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	s.mock.ExpectQuery(query).WithArgs(id).WillReturnRows(rows)
	s.repository.FindByID(id)
}

func (s *otpRepositorySuite) TestFindByPin() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("SELECT * FROM `one_time_passwords`")
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	s.mock.ExpectQuery(query).WithArgs(id).WillReturnRows(rows)
	s.repository.FindByPin(id)
}

func (s *otpRepositorySuite) TestFindPinAndUserID() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("SELECT * FROM `one_time_passwords`")
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	s.mock.ExpectQuery(query).WithArgs("any", id).WillReturnRows(rows)
	s.repository.FindByPinAndUserID("any", id)
}

func (s *otpRepositorySuite) TestCreate() {
	query := regexp.QuoteMeta("INSERT INTO `one_time_passwords`")

	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	s.mock.ExpectClose()
	s.repository.Create(models.OTP{})

	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnError(errors.New("error"))
	s.mock.ExpectCommit()
	s.repository.Create(models.OTP{})
	s.mock.ExpectClose()
}

func (s *otpRepositorySuite) TestUpdate() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("UPDATE `one_time_passwords`")
	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	query = regexp.QuoteMeta("SELECT * FROM `one_time_passwords`")
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	s.mock.ExpectQuery(query).WithArgs(id).WillReturnRows(rows)

	s.repository.Update(models.OTP{}, id)
}

func (s *otpRepositorySuite) TestDelete() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("UPDATE `one_time_passwords`")
	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	s.repository.Delete(id)
}
