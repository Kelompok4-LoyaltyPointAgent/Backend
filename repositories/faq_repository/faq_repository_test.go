package faq_repository

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

type faqRepositorySuite struct {
	suite.Suite
	mock       sqlmock.Sqlmock
	repository FAQRepository
}

func (s *faqRepositorySuite) SetupSuite() {
	var conn *sql.DB
	conn, s.mock = testhelper.MockDB()
	gorm := testhelper.InitDB(conn)
	s.repository = NewFAQRepository(gorm)
}

func TestFAQRepositorySuite(t *testing.T) {
	suite.Run(t, new(faqRepositorySuite))
}

func (s *faqRepositorySuite) TestFindAll() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("SELECT * FROM `frequently_asked_questions`")
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	s.mock.ExpectQuery(query).WillReturnRows(rows)
	if _, err := s.repository.FindAll(nil, nil); err != nil {
		log.Println(err)
	}
}

func (s *faqRepositorySuite) TestFindByID() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("SELECT * FROM `frequently_asked_questions`")
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	s.mock.ExpectQuery(query).WithArgs(id).WillReturnRows(rows)
	if _, err := s.repository.FindByID(id); err != nil {
		log.Println(err)
	}
}

func (s *faqRepositorySuite) TestCreate() {
	query := regexp.QuoteMeta("INSERT INTO `frequently_asked_questions`")

	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	s.mock.ExpectClose()
	if _, err := s.repository.Create(models.FAQ{}); err != nil {
		log.Println(err)
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnError(errors.New("error"))
	s.mock.ExpectCommit()
	if _, err := s.repository.Create(models.FAQ{}); err != nil {
		log.Println(err)
	}
	s.mock.ExpectClose()
}

func (s *faqRepositorySuite) TestUpdate() {
	s.mock.ExpectBegin()
	query := regexp.QuoteMeta("UPDATE `frequently_asked_questions`")
	s.mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	id := uuid.NewString()
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	query = regexp.QuoteMeta("SELECT * FROM `frequently_asked_questions`")
	s.mock.ExpectQuery(query).WithArgs(id).WillReturnRows(rows)

	if _, err := s.repository.Update(models.FAQ{}, id); err != nil {
		log.Println(err)
	}
}

func (s *faqRepositorySuite) TestDelete() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("UPDATE `frequently_asked_questions`")
	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	if err := s.repository.Delete(id); err != nil {
		log.Println(err)
	}
}
