package packages_repository

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

type packagesRepositorySuite struct {
	suite.Suite
	mock       sqlmock.Sqlmock
	repository PackagesRepository
}

func (s *packagesRepositorySuite) SetupSuite() {
	var conn *sql.DB
	conn, s.mock = testhelper.MockDB()
	gorm := testhelper.InitDB(conn)
	s.repository = NewPackagesRepository(gorm)
}

func TestPackagesRepositorySuite(t *testing.T) {
	suite.Run(t, new(packagesRepositorySuite))
}

func (s *packagesRepositorySuite) TestFindAll() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("SELECT * FROM `packages`")
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	s.mock.ExpectQuery(query).WillReturnRows(rows)
	if _, err := s.repository.FindAll(); err != nil {
		log.Println(err)
	}
}

func (s *packagesRepositorySuite) TestCreate() {
	query := regexp.QuoteMeta("INSERT INTO `packages`")

	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	s.mock.ExpectClose()
	if _, err := s.repository.Create(models.Packages{}); err != nil {
		log.Println(err)
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnError(errors.New("error"))
	s.mock.ExpectCommit()
	if _, err := s.repository.Create(models.Packages{}); err != nil {
		log.Println(err)
	}
	s.mock.ExpectClose()
}

func (s *packagesRepositorySuite) TestFindByProductID() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("SELECT * FROM `packages`")
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	s.mock.ExpectQuery(query).WithArgs(id).WillReturnRows(rows)
	if _, err := s.repository.FindByProductID(id); err != nil {
		log.Println(err)
	}
}

func (s *packagesRepositorySuite) TestFindByProvider() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("SELECT * FROM `packages`")
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	s.mock.ExpectQuery(query).WillReturnRows(rows)
	if _, err := s.repository.FindByProvider("any"); err != nil {
		log.Println(err)
	}
}

func (s *packagesRepositorySuite) TestFindByRecommended() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("SELECT * FROM `packages`")
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	s.mock.ExpectQuery(query).WillReturnRows(rows)
	if _, err := s.repository.FindByRecommended(); err != nil {
		log.Println(err)
	}
}

func (s *packagesRepositorySuite) TestUpdateByProductID() {
	s.mock.ExpectBegin()
	query := regexp.QuoteMeta("UPDATE `packages`")
	s.mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	id := uuid.NewString()
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	query = regexp.QuoteMeta("SELECT * FROM `packages`")
	s.mock.ExpectQuery(query).WithArgs(id).WillReturnRows(rows)

	if _, err := s.repository.UpdateByProductID(models.Packages{}, id); err != nil {
		log.Println(err)
	}
}

func (s *packagesRepositorySuite) TestDeleteByProductID() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("UPDATE `packages`")
	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	if err := s.repository.DeleteByProductID(id); err != nil {
		log.Println(err)
	}
}
