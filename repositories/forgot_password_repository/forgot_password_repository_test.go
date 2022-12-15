package forgot_password_repository

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

type forgotPasswordRepositorySuite struct {
	suite.Suite
	mock       sqlmock.Sqlmock
	repository ForgotPasswordRepository
}

func (s *forgotPasswordRepositorySuite) SetupSuite() {
	var conn *sql.DB
	conn, s.mock = testhelper.MockDB()
	gorm := testhelper.InitDB(conn)
	s.repository = NewForgotPasswordRepository(gorm)
}

func TestForgotPasswordRepositorySuite(t *testing.T) {
	suite.Run(t, new(forgotPasswordRepositorySuite))
}

func (s *forgotPasswordRepositorySuite) TestFindByID() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("SELECT * FROM `forgot_passwords`")
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	s.mock.ExpectQuery(query).WithArgs(id).WillReturnRows(rows)
	if _, err := s.repository.FindByID(id); err != nil {
		log.Println(err)
	}
}

func (s *forgotPasswordRepositorySuite) TestFindByAccessKey() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("SELECT * FROM `forgot_passwords`")
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	s.mock.ExpectQuery(query).WithArgs(id).WillReturnRows(rows)
	if _, err := s.repository.FindByAccessKey(id); err != nil {
		log.Println(err)
	}
}

func (s *forgotPasswordRepositorySuite) TestFindAccessKeyAndUserID() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("SELECT * FROM `forgot_passwords`")
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	s.mock.ExpectQuery(query).WithArgs("any", id).WillReturnRows(rows)
	if _, err := s.repository.FindByAccessKeyAndUserID("any", id); err != nil {
		log.Println(err)
	}
}

func (s *forgotPasswordRepositorySuite) TestCreate() {
	query := regexp.QuoteMeta("INSERT INTO `forgot_passwords`")

	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	s.mock.ExpectClose()
	if _, err := s.repository.Create(models.ForgotPassword{}); err != nil {
		log.Println(err)
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnError(errors.New("error"))
	s.mock.ExpectCommit()
	if _, err := s.repository.Create(models.ForgotPassword{}); err != nil {
		log.Println(err)
	}
	s.mock.ExpectClose()
}

func (s *forgotPasswordRepositorySuite) TestUpdate() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("UPDATE `forgot_passwords`")
	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	query = regexp.QuoteMeta("SELECT * FROM `forgot_passwords`")
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	s.mock.ExpectQuery(query).WithArgs(id).WillReturnRows(rows)

	if _, err := s.repository.Update(models.ForgotPassword{}, id); err != nil {
		log.Println(err)
	}
}

func (s *forgotPasswordRepositorySuite) TestDelete() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("UPDATE `forgot_passwords`")
	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	if err := s.repository.Delete(id); err != nil {
		log.Println(err)
	}
}
