package forgot_password_repository

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
	s.repository.FindByID(id)
}

func (s *forgotPasswordRepositorySuite) TestFindByAccessKey() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("SELECT * FROM `forgot_passwords`")
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	s.mock.ExpectQuery(query).WithArgs(id).WillReturnRows(rows)
	s.repository.FindByAccessKey(id)
}

func (s *forgotPasswordRepositorySuite) TestFindAccessKeyAndUserID() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("SELECT * FROM `forgot_passwords`")
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	s.mock.ExpectQuery(query).WithArgs("any", id).WillReturnRows(rows)
	s.repository.FindByAccessKeyAndUserID("any", id)
}

func (s *forgotPasswordRepositorySuite) TestCreate() {
	query := regexp.QuoteMeta("INSERT INTO `forgot_passwords`")

	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	s.mock.ExpectClose()
	s.repository.Create(models.ForgotPassword{})

	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnError(errors.New("error"))
	s.mock.ExpectCommit()
	s.repository.Create(models.ForgotPassword{})
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

	s.repository.Update(models.ForgotPassword{}, id)
}

func (s *forgotPasswordRepositorySuite) TestDelete() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("UPDATE `forgot_passwords`")
	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	s.repository.Delete(id)
}
