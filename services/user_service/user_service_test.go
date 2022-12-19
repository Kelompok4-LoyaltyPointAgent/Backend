package user_service

import (
	"errors"
	"log"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kelompok4-loyaltypointagent/backend/dto/payload"
	"github.com/kelompok4-loyaltypointagent/backend/dto/response"
	"github.com/kelompok4-loyaltypointagent/backend/models"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/mocks"
	"github.com/stretchr/testify/suite"
)

type userServiceSuite struct {
	suite.Suite
	ctrl       *gomock.Controller
	repository *mocks.MockUserRepository
	service    UserService
}

func (s *userServiceSuite) SetupSuite() {
	s.ctrl = gomock.NewController(s.T())
	s.repository = mocks.NewMockUserRepository(s.ctrl)
	s.service = NewUserService(s.repository)
}

func (s *userServiceSuite) TearDownSuite() {
	s.ctrl.Finish()
}

func TestUserServiceSuite(t *testing.T) {
	suite.Run(t, new(userServiceSuite))
}

func (s *userServiceSuite) TestFindAll() {

	testCase := []struct {
		Name           string
		ExpectedReturn []response.UserResponse
		ExpectedError  error
		MockReturn     []models.User
		MockError      error
	}{
		{
			Name: "Success",
			ExpectedReturn: []response.UserResponse{
				{

					Name:  "User 1",
					Email: "test@gmail.com",
				},
			},
			ExpectedError: nil,
			MockReturn: []models.User{
				{

					Name:  "User 1",
					Email: "test@gmail.com",
				},
			},
			MockError: nil,
		},
		{
			Name:           "Error",
			ExpectedReturn: nil,
			ExpectedError:  errors.New("Error"),
			MockReturn:     nil,
			MockError:      errors.New("Error"),
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {
			s.repository.EXPECT().FindAll(gomock.Any(), gomock.Any()).Return(tc.MockReturn, tc.MockError)
			result, err := s.service.FindAll("User")
			if err != nil {
				log.Println(err.Error())
			}
			if tc.ExpectedReturn != nil {
				s.Equal(tc.ExpectedReturn[0].Email, result[0].Email)
			}

			s.Equal(tc.ExpectedError, err)
		})
	}

}

func (s *userServiceSuite) TestCreate() {
	testCase := []struct {
		Name           string
		ExpectedReturn response.UserResponse
		ExpectedError  error
		MockReturn     models.User
		MockError      error
		MockPayload    payload.UserPayload
	}{
		{
			Name: "Success",
			ExpectedReturn: response.UserResponse{

				Name:  "User 1",
				Email: "test@gmail.com",
			},
			ExpectedError: nil,
			MockReturn: models.User{
				Name:  "User 1",
				Email: "test@gmail.com",
			},

			MockError: nil,
			MockPayload: payload.UserPayload{
				Name:     "User 1",
				Email:    "test@gmail.com",
				Password: "test",
			},
		},

		{
			Name:           "Error",
			ExpectedReturn: response.UserResponse{},
			ExpectedError:  errors.New("Error"),
			MockReturn:     models.User{},
			MockError:      errors.New("Error"),
			MockPayload: payload.UserPayload{
				Name:     "User 1",
				Email:    "test@gmail.com",
				Password: "test",
			},
		},
	}

	for _, tc := range testCase {
		s.Run(tc.Name, func() {

			s.repository.EXPECT().Create(gomock.Any()).Return(tc.MockReturn, tc.MockError)

			result, err := s.service.Create(tc.MockPayload)
			if err != nil {
				log.Println(err.Error())
			}
			
			s.Equal(tc.ExpectedReturn.Email, result.Email)

			s.Equal(tc.ExpectedError, err)
		})
	}

}
