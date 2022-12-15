package product_repository

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

type productRepositorySuite struct {
	suite.Suite
	mock       sqlmock.Sqlmock
	repository ProductRepository
}

func (s *productRepositorySuite) SetupSuite() {
	var conn *sql.DB
	conn, s.mock = testhelper.MockDB()
	gorm := testhelper.InitDB(conn)
	s.repository = NewProductRepository(gorm)
}

func TestProductRepositorySuite(t *testing.T) {
	suite.Run(t, new(productRepositorySuite))
}

func (s *productRepositorySuite) TestFindAll() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("SELECT * FROM `products`")
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	s.mock.ExpectQuery(query).WillReturnRows(rows)
	s.repository.FindAll()
}

func (s *productRepositorySuite) TestFindByID() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("SELECT * FROM `products`")
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	s.mock.ExpectQuery(query).WithArgs(id).WillReturnRows(rows)
	s.repository.FindByID(id)
}

func (s *productRepositorySuite) TestCreate() {
	query := regexp.QuoteMeta("INSERT INTO `products`")

	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	s.mock.ExpectClose()
	s.repository.Create(models.Product{})

	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnError(errors.New("error"))
	s.mock.ExpectCommit()
	s.repository.Create(models.Product{})
	s.mock.ExpectClose()
}

func (s *productRepositorySuite) TestUpdate() {
	s.mock.ExpectBegin()
	query := regexp.QuoteMeta("UPDATE `products`")
	s.mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	id := uuid.NewString()
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	query = regexp.QuoteMeta("SELECT * FROM `products`")
	s.mock.ExpectQuery(query).WithArgs(id).WillReturnRows(rows)

	s.repository.Update(models.Product{}, id)
}

func (s *productRepositorySuite) TestDelete() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("UPDATE `products`")
	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	s.repository.DeleteByID(id)
}

func (s *productRepositorySuite) TestSetBooleanRecommended() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("UPDATE `products`")
	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	s.repository.SetBooleanRecommended(id, true)
}
