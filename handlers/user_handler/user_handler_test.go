package user_handler

import (
	"errors"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kelompok4-loyaltypointagent/backend/dto/payload"
	"github.com/kelompok4-loyaltypointagent/backend/dto/response"
	"github.com/kelompok4-loyaltypointagent/backend/services/mocks"
	"github.com/kelompok4-loyaltypointagent/backend/testhelper"
	"github.com/stretchr/testify/suite"
)

type userSuite struct {
	suite.Suite

	ctrl    *gomock.Controller
	service *mocks.MockUserService
	handler UserHandler
}

func (s *userSuite) SetupSuite() {
	s.ctrl = gomock.NewController(s.T())
	s.service = mocks.NewMockUserService(s.ctrl)
	s.handler = NewUserHandler(s.service)
}

func (s *userSuite) TearDownSuite() {
	s.ctrl.Finish()
}

func TestUserHandlerSuite(t *testing.T) {
	suite.Run(t, new(userSuite))
}

func (s *userSuite) TestCreateUser() {
	request := testhelper.Request{
		Method:      "post",
		URL:         "/api/v1/users",
		ContentType: "application/json",
	}

	testCases := []testhelper.HTTPTestCase{
		{
			Name:         "bad request",
			Request:      request,
			Body:         payload.UserPayload{},
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name:    "internal server error",
			Request: request,
			Body: payload.UserPayload{
				Name:     "Example User",
				Email:    "user@example.com",
				Password: "password",
				Points:   10,
			},
			ExpectedCode: http.StatusInternalServerError,
			ExpectedFunc: func() {
				s.service.EXPECT().Create(gomock.Any()).Return(response.UserResponse{}, errors.New("error"))
			},
		},
		{
			Name:    "ok",
			Request: request,
			Body: payload.UserPayload{
				Name:     "Example User",
				Email:    "user@example.com",
				Password: "password",
				Points:   10,
			},
			ExpectedCode: http.StatusOK,
			ExpectedFunc: func() {
				s.service.EXPECT().Create(gomock.Any()).Return(response.UserResponse{}, nil)
			},
		},
	}

	for _, testCase := range testCases {
		if testCase.ExpectedFunc != nil {
			testCase.ExpectedFunc()
		}

		s.T().Run(testCase.Name, func(t *testing.T) {
			ctx, rec := testhelper.NewContext(testCase)
			s.NoError(s.handler.CreateUser(ctx))
			s.Equal(testCase.ExpectedCode, rec.Code)
		})
	}
}
