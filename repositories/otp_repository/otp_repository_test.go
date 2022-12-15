package otp_repository

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
	if _, err := s.repository.FindByID(id); err != nil {
		log.Println(err)
	}
}

func (s *otpRepositorySuite) TestFindByPin() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("SELECT * FROM `one_time_passwords`")
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	s.mock.ExpectQuery(query).WithArgs(id).WillReturnRows(rows)
	if _, err := s.repository.FindByPin(id); err != nil {
		log.Println(err)
	}
}

func (s *otpRepositorySuite) TestFindPinAndUserID() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("SELECT * FROM `one_time_passwords`")
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	s.mock.ExpectQuery(query).WithArgs("any", id).WillReturnRows(rows)
	if _, err := s.repository.FindByPinAndUserID("any", id); err != nil {
		log.Println(err)
	}
}

func (s *otpRepositorySuite) TestCreate() {
	query := regexp.QuoteMeta("INSERT INTO `one_time_passwords`")

	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	s.mock.ExpectClose()
	if _, err := s.repository.Create(models.OTP{}); err != nil {
		log.Println(err)
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnError(errors.New("error"))
	s.mock.ExpectCommit()
	if _, err := s.repository.Create(models.OTP{}); err != nil {
		log.Println(err)
	}
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

	if _, err := s.repository.Update(models.OTP{}, id); err != nil {
		log.Println(err)
	}
}

func (s *otpRepositorySuite) TestDelete() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("UPDATE `one_time_passwords`")
	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	if err := s.repository.Delete(id); err != nil {
		log.Println(err)
	}
}
