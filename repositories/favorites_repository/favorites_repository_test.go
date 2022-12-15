package favorites_repository

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

type favoritesRepositorySuite struct {
	suite.Suite
	mock       sqlmock.Sqlmock
	repository FavoritesRepository
}

func (s *favoritesRepositorySuite) SetupSuite() {
	var conn *sql.DB
	conn, s.mock = testhelper.MockDB()
	gorm := testhelper.InitDB(conn)
	s.repository = NewFavoritesRepository(gorm)
}

func TestFavoritesRepositorySuite(t *testing.T) {
	suite.Run(t, new(favoritesRepositorySuite))
}

func (s *favoritesRepositorySuite) TestFindAll() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("SELECT * FROM `favorites`")
	rows := sqlmock.NewRows([]string{"user_id"}).AddRow(id)
	s.mock.ExpectQuery(query).WillReturnRows(rows)

	s.repository.FindAll(nil, nil)
}

func (s *favoritesRepositorySuite) TestCreate() {
	query := regexp.QuoteMeta("INSERT INTO `favorites`")

	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	s.mock.ExpectClose()
	s.repository.Create(models.Favorites{})

	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnError(errors.New("error"))
	s.mock.ExpectCommit()
	s.repository.Create(models.Favorites{})
	s.mock.ExpectClose()
}

func (s *favoritesRepositorySuite) TestDelete() {
	query := regexp.QuoteMeta("DELETE FROM `favorites`")
	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	s.repository.Delete("user_id", "product_id")
}
