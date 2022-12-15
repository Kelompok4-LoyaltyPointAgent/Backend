package user_repository

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

type userRepositorySuite struct {
	suite.Suite
	mock       sqlmock.Sqlmock
	repository UserRepository
}

func (s *userRepositorySuite) SetupSuite() {
	var conn *sql.DB
	conn, s.mock = testhelper.MockDB()
	gorm := testhelper.InitDB(conn)
	s.repository = NewUserRepository(gorm)
}

func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(userRepositorySuite))
}

func (s *userRepositorySuite) TestFindByID() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("SELECT * FROM `users`")
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	s.mock.ExpectQuery(query).WithArgs(id).WillReturnRows(rows)
	if _, err := s.repository.FindByID(id); err != nil {
		log.Println(err)
	}
}

func (s *userRepositorySuite) TestCreate() {
	query := regexp.QuoteMeta("INSERT INTO `users`")

	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	s.mock.ExpectClose()
	if _, err := s.repository.Create(models.User{}); err != nil {
		log.Println(err)
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnError(errors.New("error"))
	s.mock.ExpectCommit()
	if _, err := s.repository.Create(models.User{}); err != nil {
		log.Println(err)
	}
	s.mock.ExpectClose()
}

func (s *userRepositorySuite) TestFindAll() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("SELECT * FROM `users`")
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	s.mock.ExpectQuery(query).WillReturnRows(rows)
	if _, err := s.repository.FindAll("", ""); err != nil {
		log.Println(err)
	}
}

func (s *userRepositorySuite) TestUpdate() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("UPDATE `users`")
	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	s.mock.ExpectClose()
	if _, err := s.repository.Update(models.User{}, id); err != nil {
		log.Println(err)
	}
}

func (s *userRepositorySuite) TestDelete() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("UPDATE `users`")
	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	if _, err := s.repository.Delete(id); err != nil {
		log.Println(err)
	}
}

func (s *userRepositorySuite) TestFindByEmail() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("SELECT * FROM `users`")
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	s.mock.ExpectQuery(query).WillReturnRows(rows)
	if _, err := s.repository.FindByEmail("user@example.com"); err != nil {
		log.Println(err)
	}
}
