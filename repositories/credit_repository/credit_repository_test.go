package credit_repository

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
	if _, err := s.repository.FindAll(); err != nil {
		log.Println(err)
	}
}

func (s *creditRepositorySuite) TestCreate() {
	query := regexp.QuoteMeta("INSERT INTO `credits`")

	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	s.mock.ExpectClose()
	if _, err := s.repository.Create(models.Credit{}); err != nil {
		log.Println(err)
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnError(errors.New("error"))
	s.mock.ExpectCommit()
	if _, err := s.repository.Create(models.Credit{}); err != nil {
		log.Println(err)
	}
	s.mock.ExpectClose()
}

func (s *creditRepositorySuite) TestFindByProductID() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("SELECT * FROM `credits`")
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	s.mock.ExpectQuery(query).WithArgs(id).WillReturnRows(rows)
	if _, err := s.repository.FindByProductID(id); err != nil {
		log.Println(err)
	}
}

func (s *creditRepositorySuite) TestFindByProvider() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("SELECT * FROM `credits`")
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	s.mock.ExpectQuery(query).WillReturnRows(rows)
	if _, err := s.repository.FindByProvider("any"); err != nil {
		log.Println(err)
	}
}

func (s *creditRepositorySuite) TestFindByRecommended() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("SELECT * FROM `credits`")
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	s.mock.ExpectQuery(query).WillReturnRows(rows)
	if _, err := s.repository.FindByRecommended(); err != nil {
		log.Println(err)
	}
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

	if _, err := s.repository.UpdateByProductID(models.Credit{}, id); err != nil {
		log.Println(err)
	}
}

func (s *creditRepositorySuite) TestDeleteByProductID() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("UPDATE `credits`")
	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	if err := s.repository.DeleteByProductID(id); err != nil {
		log.Println(err)
	}
}
