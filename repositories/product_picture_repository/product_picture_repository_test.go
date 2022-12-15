package product_picture_repository

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

type productPictureRepositorySuite struct {
	suite.Suite
	mock       sqlmock.Sqlmock
	repository ProductPictureRepository
}

func (s *productPictureRepositorySuite) SetupSuite() {
	var conn *sql.DB
	conn, s.mock = testhelper.MockDB()
	gorm := testhelper.InitDB(conn)
	s.repository = NewProductPictureRepository(gorm)
}

func TestProductPictureRepositorySuite(t *testing.T) {
	suite.Run(t, new(productPictureRepositorySuite))
}

func (s *productPictureRepositorySuite) TestFindByID() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("SELECT * FROM `product_pictures`")
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	s.mock.ExpectQuery(query).WithArgs(id).WillReturnRows(rows)
	s.repository.FindByID(id)
}

func (s *productPictureRepositorySuite) TestFindByName() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("SELECT * FROM `product_pictures`")
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	s.mock.ExpectQuery(query).WithArgs("user").WillReturnRows(rows)
	s.repository.FindByName("user")
}

func (s *productPictureRepositorySuite) TestCreate() {
	query := regexp.QuoteMeta("INSERT INTO `product_pictures`")

	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	s.mock.ExpectClose()
	s.repository.Create(models.ProductPicture{})

	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnError(errors.New("error"))
	s.mock.ExpectCommit()
	s.repository.Create(models.ProductPicture{})
	s.mock.ExpectClose()
}

func (s *productPictureRepositorySuite) TestUpdate() {
	s.mock.ExpectBegin()
	query := regexp.QuoteMeta("UPDATE `product_pictures`")
	s.mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	id := uuid.NewString()
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	query = regexp.QuoteMeta("SELECT * FROM `product_pictures`")
	s.mock.ExpectQuery(query).WithArgs(id).WillReturnRows(rows)

	s.repository.Update(models.ProductPicture{}, id)
}

func (s *productPictureRepositorySuite) TestDelete() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("UPDATE `product_pictures`")
	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	s.repository.Delete(id)
}
