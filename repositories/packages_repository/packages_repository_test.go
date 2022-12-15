package packages_repository

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
	s.repository.FindAll()
}

func (s *packagesRepositorySuite) TestCreate() {
	query := regexp.QuoteMeta("INSERT INTO `packages`")

	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	s.mock.ExpectClose()
	s.repository.Create(models.Packages{})

	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnError(errors.New("error"))
	s.mock.ExpectCommit()
	s.repository.Create(models.Packages{})
	s.mock.ExpectClose()
}

func (s *packagesRepositorySuite) TestFindByProductID() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("SELECT * FROM `packages`")
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	s.mock.ExpectQuery(query).WithArgs(id).WillReturnRows(rows)
	s.repository.FindByProductID(id)
}

func (s *packagesRepositorySuite) TestFindByProvider() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("SELECT * FROM `packages`")
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	s.mock.ExpectQuery(query).WillReturnRows(rows)
	s.repository.FindByProvider("any")
}

func (s *packagesRepositorySuite) TestFindByRecommended() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("SELECT * FROM `packages`")
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	s.mock.ExpectQuery(query).WillReturnRows(rows)
	s.repository.FindByRecommended()
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

	s.repository.UpdateByProductID(models.Packages{}, id)
}

func (s *packagesRepositorySuite) TestDeleteByProductID() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("UPDATE `packages`")
	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	s.repository.DeleteByProductID(id)
}
