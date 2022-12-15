package feedback_repository

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

type feedbackRepositorySuite struct {
	suite.Suite
	mock       sqlmock.Sqlmock
	repository FeedbackRepository
}

func (s *feedbackRepositorySuite) SetupSuite() {
	var conn *sql.DB
	conn, s.mock = testhelper.MockDB()
	gorm := testhelper.InitDB(conn)
	s.repository = NewFeedbackRepository(gorm)
}

func TestFeedbackRepositorySuite(t *testing.T) {
	suite.Run(t, new(feedbackRepositorySuite))
}

func (s *feedbackRepositorySuite) TestFindByID() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("SELECT * FROM `feedbacks`")
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	s.mock.ExpectQuery(query).WithArgs(id).WillReturnRows(rows)
	s.repository.FindByID(id)
}

func (s *feedbackRepositorySuite) TestCreate() {
	query := regexp.QuoteMeta("INSERT INTO `feedbacks`")

	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	s.mock.ExpectClose()
	s.repository.Create(models.Feedbacks{})

	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).WillReturnError(errors.New("error"))
	s.mock.ExpectCommit()
	s.repository.Create(models.Feedbacks{})
	s.mock.ExpectClose()
}

func (s *feedbackRepositorySuite) TestFindAll() {
	id := uuid.NewString()
	query := regexp.QuoteMeta("SELECT * FROM `feedbacks`")
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	s.mock.ExpectQuery(query).WillReturnRows(rows)
	s.repository.FindAll()
}
