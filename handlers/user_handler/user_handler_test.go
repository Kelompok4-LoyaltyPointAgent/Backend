package user_handler

import (
	"errors"
	"net/http"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/golang/mock/gomock"
	"github.com/kelompok4-loyaltypointagent/backend/dto/payload"
	"github.com/kelompok4-loyaltypointagent/backend/dto/response"
	"github.com/kelompok4-loyaltypointagent/backend/helper"
	"github.com/kelompok4-loyaltypointagent/backend/services/mocks"
	"github.com/kelompok4-loyaltypointagent/backend/testhelper"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type userHandlerSuite struct {
	suite.Suite

	ctrl    *gomock.Controller
	service *mocks.MockUserService
	handler UserHandler
}

func (s *userHandlerSuite) SetupSuite() {
	s.ctrl = gomock.NewController(s.T())
	s.service = mocks.NewMockUserService(s.ctrl)
	s.handler = NewUserHandler(s.service)
}

func (s *userHandlerSuite) TearDownSuite() {
	s.ctrl.Finish()
}

func TestUserHandlerSuite(t *testing.T) {
	suite.Run(t, new(userHandlerSuite))
}

func (s *userHandlerSuite) TestCreateUser() {
	request := testhelper.Request{
		Method:      http.MethodPost,
		Path:        "/api/v1/users",
		ContentType: echo.MIMEApplicationJSON,
	}

	testCases := []testhelper.HTTPTestCase{
		{
			Name:         "bad request",
			Request:      request,
			Body:         "invalid request",
			ExpectedCode: http.StatusBadRequest,
		},
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

func (s *userHandlerSuite) TestLogin() {
	request := testhelper.Request{
		Method:      http.MethodPost,
		Path:        "/api/v1/login",
		ContentType: echo.MIMEApplicationJSON,
	}

	testCases := []testhelper.HTTPTestCase{
		{
			Name:         "bad request",
			Request:      request,
			Body:         "invalid body",
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name:    "bad request",
			Request: request,
			Body: payload.LoginPayload{
				Email:    "user@example.com",
				Password: "password",
			},
			ExpectedCode: http.StatusBadRequest,
			ExpectedFunc: func() {
				s.service.EXPECT().Login(gomock.Any()).Return(response.LoginResponse{}, errors.New("error"))
			},
		},
		{
			Name:    "ok",
			Request: request,
			Body: payload.LoginPayload{
				Email:    "user@example.com",
				Password: "password",
			},
			ExpectedCode: http.StatusOK,
			ExpectedFunc: func() {
				s.service.EXPECT().Login(gomock.Any()).Return(response.LoginResponse{}, nil)
			},
		},
	}

	for _, testCase := range testCases {
		if testCase.ExpectedFunc != nil {
			testCase.ExpectedFunc()
		}

		s.T().Run(testCase.Name, func(t *testing.T) {
			ctx, rec := testhelper.NewContext(testCase)
			s.NoError(s.handler.Login(ctx))
			s.Equal(testCase.ExpectedCode, rec.Code)
		})
	}
}

func (s *userHandlerSuite) TestUpdateUser() {
	request := testhelper.Request{
		Method:      http.MethodPut,
		Path:        "/api/v1/users",
		ContentType: echo.MIMEApplicationJSON,
	}

	testCases := []testhelper.HTTPTestCase{
		{
			Name:         "bad request",
			Request:      request,
			Body:         "invalid body",
			ExpectedCode: http.StatusBadRequest,
		},
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
				Name:     "user",
				Email:    "user@example.com",
				Password: "password",
				Points:   0,
			},
			Token:        jwt.NewWithClaims(jwt.SigningMethodHS256, &helper.JWTCustomClaims{}),
			ExpectedCode: http.StatusInternalServerError,
			ExpectedFunc: func() {
				s.service.EXPECT().UpdateProfile(gomock.Any(), gomock.Any()).Return(response.UserResponse{}, errors.New("error"))
			},
		},
		{
			Name:    "ok",
			Request: request,
			Body: payload.UserPayload{
				Name:     "user",
				Email:    "user@example.com",
				Password: "password",
				Points:   0,
			},
			Token:        jwt.NewWithClaims(jwt.SigningMethodHS256, &helper.JWTCustomClaims{}),
			ExpectedCode: http.StatusOK,
			ExpectedFunc: func() {
				s.service.EXPECT().UpdateProfile(gomock.Any(), gomock.Any()).Return(response.UserResponse{}, nil)
			},
		},
	}

	for _, testCase := range testCases {
		if testCase.ExpectedFunc != nil {
			testCase.ExpectedFunc()
		}

		s.T().Run(testCase.Name, func(t *testing.T) {
			ctx, rec := testhelper.NewContext(testCase)
			s.NoError(s.handler.UpdateUser(ctx))
			s.Equal(testCase.ExpectedCode, rec.Code)
		})
	}
}

func (s *userHandlerSuite) TestChangePassword() {
	request := testhelper.Request{
		Method:      http.MethodPut,
		Path:        "/api/v1/users/change-password",
		ContentType: echo.MIMEApplicationJSON,
	}

	testCases := []testhelper.HTTPTestCase{
		{
			Name:         "bad request",
			Request:      request,
			Body:         "invalid body",
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name:         "bad request",
			Request:      request,
			Body:         payload.ChangePasswordPayload{},
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name:    "bad request",
			Request: request,
			Body: payload.ChangePasswordPayload{
				OldPassword:     "password",
				NewPassword:     "newpassword",
				ConfirmPassword: "newpassword",
			},
			Token:        jwt.NewWithClaims(jwt.SigningMethodHS256, &helper.JWTCustomClaims{}),
			ExpectedCode: http.StatusBadRequest,
			ExpectedFunc: func() {
				s.service.EXPECT().ChangePassword(gomock.Any(), gomock.Any()).Return(response.UserResponse{}, errors.New("error"))
			},
		},
		{
			Name:    "ok",
			Request: request,
			Body: payload.ChangePasswordPayload{
				OldPassword:     "password",
				NewPassword:     "newpassword",
				ConfirmPassword: "newpassword",
			},
			Token:        jwt.NewWithClaims(jwt.SigningMethodHS256, &helper.JWTCustomClaims{}),
			ExpectedCode: http.StatusOK,
			ExpectedFunc: func() {
				s.service.EXPECT().ChangePassword(gomock.Any(), gomock.Any()).Return(response.UserResponse{}, nil)
			},
		},
	}

	for _, testCase := range testCases {
		if testCase.ExpectedFunc != nil {
			testCase.ExpectedFunc()
		}

		s.T().Run(testCase.Name, func(t *testing.T) {
			ctx, rec := testhelper.NewContext(testCase)
			s.NoError(s.handler.ChangePassword(ctx))
			s.Equal(testCase.ExpectedCode, rec.Code)
		})
	}
}

func (s *userHandlerSuite) TestChangePasswordFromResetPassword() {
	request := testhelper.Request{
		Method:      http.MethodPut,
		Path:        "/api/v1/users/reset-password",
		ContentType: echo.MIMEApplicationJSON,
	}

	testCases := []testhelper.HTTPTestCase{
		{
			Name:         "bad request",
			Request:      request,
			Body:         "invalid body",
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name:         "bad request",
			Request:      request,
			Body:         payload.ChangePasswordFromResetPasswordPayload{},
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name:    "bad request",
			Request: request,
			Body: payload.ChangePasswordFromResetPasswordPayload{
				NewPassword:     "newpassword",
				ConfirmPassword: "newpassword",
			},
			Token:        jwt.NewWithClaims(jwt.SigningMethodHS256, &helper.JWTCustomClaims{}),
			ExpectedCode: http.StatusBadRequest,
			ExpectedFunc: func() {
				s.service.EXPECT().ChangePasswordFromResetPassword(gomock.Any(), gomock.Any()).Return(response.UserResponse{}, errors.New("error"))
			},
		},
		{
			Name:    "ok",
			Request: request,
			Body: payload.ChangePasswordFromResetPasswordPayload{
				NewPassword:     "newpassword",
				ConfirmPassword: "newpassword",
			},
			Token:        jwt.NewWithClaims(jwt.SigningMethodHS256, &helper.JWTCustomClaims{}),
			ExpectedCode: http.StatusOK,
			ExpectedFunc: func() {
				s.service.EXPECT().ChangePasswordFromResetPassword(gomock.Any(), gomock.Any()).Return(response.UserResponse{}, nil)
			},
		},
	}

	for _, testCase := range testCases {
		if testCase.ExpectedFunc != nil {
			testCase.ExpectedFunc()
		}

		s.T().Run(testCase.Name, func(t *testing.T) {
			ctx, rec := testhelper.NewContext(testCase)
			s.NoError(s.handler.ChangePasswordFromResetPassword(ctx))
			s.Equal(testCase.ExpectedCode, rec.Code)
		})
	}
}

func (s *userHandlerSuite) TestFindUserByID() {
	request := testhelper.Request{
		Method:      http.MethodGet,
		Path:        "/api/v1/users/:id",
		ContentType: echo.MIMEApplicationJSON,
	}

	testCases := []testhelper.HTTPTestCase{
		{
			Name:         "bad request",
			Request:      request,
			Token:        jwt.NewWithClaims(jwt.SigningMethodHS256, &helper.JWTCustomClaims{}),
			ExpectedCode: http.StatusBadRequest,
			ExpectedFunc: func() {
				s.service.EXPECT().FindByID(gomock.Any()).Return(response.UserResponse{}, errors.New("error"))
			},
		},
		{
			Name:         "ok",
			Request:      request,
			Token:        jwt.NewWithClaims(jwt.SigningMethodHS256, &helper.JWTCustomClaims{}),
			ExpectedCode: http.StatusOK,
			ExpectedFunc: func() {
				s.service.EXPECT().FindByID(gomock.Any()).Return(response.UserResponse{}, nil)
			},
		},
	}

	for _, testCase := range testCases {
		if testCase.ExpectedFunc != nil {
			testCase.ExpectedFunc()
		}

		s.T().Run(testCase.Name, func(t *testing.T) {
			ctx, rec := testhelper.NewContext(testCase)
			s.NoError(s.handler.FindUserByID(ctx))
			s.Equal(testCase.ExpectedCode, rec.Code)
		})
	}
}

func (s *userHandlerSuite) TestFindAllUser() {
	request := testhelper.Request{
		Method:      http.MethodGet,
		Path:        "/api/v1/users",
		ContentType: echo.MIMEApplicationJSON,
	}

	testCases := []testhelper.HTTPTestCase{
		{
			Name:         "bad request",
			Request:      request,
			Token:        jwt.NewWithClaims(jwt.SigningMethodHS256, &helper.JWTCustomClaims{}),
			ExpectedCode: http.StatusBadRequest,
			ExpectedFunc: func() {
				s.service.EXPECT().FindAll().Return([]response.UserResponse{}, errors.New("error"))
			},
		},
		{
			Name:         "ok",
			Request:      request,
			Token:        jwt.NewWithClaims(jwt.SigningMethodHS256, &helper.JWTCustomClaims{}),
			ExpectedCode: http.StatusOK,
			ExpectedFunc: func() {
				s.service.EXPECT().FindAll().Return([]response.UserResponse{}, nil)
			},
		},
	}

	for _, testCase := range testCases {
		if testCase.ExpectedFunc != nil {
			testCase.ExpectedFunc()
		}

		s.T().Run(testCase.Name, func(t *testing.T) {
			ctx, rec := testhelper.NewContext(testCase)
			s.NoError(s.handler.FindAllUser(ctx))
			s.Equal(testCase.ExpectedCode, rec.Code)
		})
	}
}

func (s *userHandlerSuite) TestFindUserByIDByAdmin() {
	request := testhelper.Request{
		Method:      http.MethodGet,
		Path:        "/api/v1/users/:id",
		ContentType: echo.MIMEApplicationJSON,
	}

	testCases := []testhelper.HTTPTestCase{
		{
			Name:         "bad request",
			Request:      request,
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name: "bad request",
			Request: testhelper.Request{
				Method:      request.Method,
				ContentType: request.ContentType,
				Path:        request.Path,
				PathParam: &testhelper.PathParam{
					Names:  []string{"id"},
					Values: []string{"1"},
				},
			},
			ExpectedCode: http.StatusBadRequest,
			ExpectedFunc: func() {
				s.service.EXPECT().FindByID(gomock.Any()).Return(response.UserResponse{}, errors.New("error"))
			},
		},
		{
			Name: "bad request",
			Request: testhelper.Request{
				Method:      request.Method,
				ContentType: request.ContentType,
				Path:        request.Path,
				PathParam: &testhelper.PathParam{
					Names:  []string{"id"},
					Values: []string{"1"},
				},
			},
			ExpectedCode: http.StatusOK,
			ExpectedFunc: func() {
				s.service.EXPECT().FindByID(gomock.Any()).Return(response.UserResponse{}, nil)
			},
		},
	}

	for _, testCase := range testCases {
		if testCase.ExpectedFunc != nil {
			testCase.ExpectedFunc()
		}

		s.T().Run(testCase.Name, func(t *testing.T) {
			ctx, rec := testhelper.NewContext(testCase)
			s.NoError(s.handler.FindUserByIDByAdmin(ctx))
			s.Equal(testCase.ExpectedCode, rec.Code)
		})
	}
}

func (s *userHandlerSuite) TestUpdateUserByAdmin() {
	request := testhelper.Request{
		Method:      http.MethodPut,
		Path:        "/api/v1/users/:id",
		ContentType: echo.MIMEApplicationJSON,
	}

	testCases := []testhelper.HTTPTestCase{
		{
			Name:         "bad request",
			Request:      request,
			Body:         "invalid request",
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name:         "bad request",
			Request:      request,
			Body:         payload.UserPayload{},
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name:    "bad request",
			Request: request,
			Body: payload.UserPayload{
				Name:     "user",
				Email:    "user@example.com",
				Password: "password",
				Points:   0,
			},
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name: "internal server error",
			Request: testhelper.Request{
				Method:      request.Method,
				ContentType: request.ContentType,
				Path:        request.Path,
				PathParam: &testhelper.PathParam{
					Names:  []string{"id"},
					Values: []string{"1"},
				},
			},
			Body: payload.UserPayload{
				Name:     "user",
				Email:    "user@example.com",
				Password: "password",
				Points:   0,
			},
			ExpectedFunc: func() {
				s.service.EXPECT().UpdateUserByAdmin(gomock.Any(), gomock.Any()).Return(response.UserResponse{}, errors.New("error"))
			},
			ExpectedCode: http.StatusInternalServerError,
		},
		{
			Name: "ok",
			Request: testhelper.Request{
				Method:      request.Method,
				ContentType: request.ContentType,
				Path:        request.Path,
				PathParam: &testhelper.PathParam{
					Names:  []string{"id"},
					Values: []string{"1"},
				},
			},
			Body: payload.UserPayload{
				Name:     "user",
				Email:    "user@example.com",
				Password: "password",
				Points:   0,
			},
			ExpectedFunc: func() {
				s.service.EXPECT().UpdateUserByAdmin(gomock.Any(), gomock.Any()).Return(response.UserResponse{}, nil)
			},
			ExpectedCode: http.StatusOK,
		},
	}

	for _, testCase := range testCases {
		if testCase.ExpectedFunc != nil {
			testCase.ExpectedFunc()
		}

		s.T().Run(testCase.Name, func(t *testing.T) {
			ctx, rec := testhelper.NewContext(testCase)
			s.NoError(s.handler.UpdateUserByAdmin(ctx))
			s.Equal(testCase.ExpectedCode, rec.Code)
		})
	}
}

func (s *userHandlerSuite) TestDeleteUserByAdmin() {
	request := testhelper.Request{
		Method:      http.MethodDelete,
		Path:        "/api/v1/users/:id",
		ContentType: echo.MIMEApplicationJSON,
	}

	testCases := []testhelper.HTTPTestCase{
		{
			Name:         "bad request",
			Request:      request,
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name: "internal server error",
			Request: testhelper.Request{
				Method:      request.Method,
				ContentType: request.ContentType,
				Path:        request.Path,
				PathParam: &testhelper.PathParam{
					Names:  []string{"id"},
					Values: []string{"1"},
				},
			},
			ExpectedFunc: func() {
				s.service.EXPECT().Delete(gomock.Any()).Return(response.UserResponse{}, errors.New("error"))
			},
			ExpectedCode: http.StatusInternalServerError,
		},
		{
			Name: "ok",
			Request: testhelper.Request{
				Method:      request.Method,
				ContentType: request.ContentType,
				Path:        request.Path,
				PathParam: &testhelper.PathParam{
					Names:  []string{"id"},
					Values: []string{"1"},
				},
			},
			ExpectedFunc: func() {
				s.service.EXPECT().Delete(gomock.Any()).Return(response.UserResponse{}, nil)
			},
			ExpectedCode: http.StatusOK,
		},
	}

	for _, testCase := range testCases {
		if testCase.ExpectedFunc != nil {
			testCase.ExpectedFunc()
		}

		s.T().Run(testCase.Name, func(t *testing.T) {
			ctx, rec := testhelper.NewContext(testCase)
			s.NoError(s.handler.DeleteUserByAdmin(ctx))
			s.Equal(testCase.ExpectedCode, rec.Code)
		})
	}
}

func (s *userHandlerSuite) TestCheckPassword() {
	request := testhelper.Request{
		Method:      http.MethodPost,
		Path:        "/api/v1/users/check-password",
		ContentType: echo.MIMEApplicationJSON,
	}

	testCases := []testhelper.HTTPTestCase{
		{
			Name:         "bad request",
			Request:      request,
			Body:         "invalid request",
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name:    "internal server error",
			Request: request,
			Body: payload.CheckPasswordPayload{
				CheckPassword: "password",
			},
			Token:        jwt.NewWithClaims(jwt.SigningMethodHS256, &helper.JWTCustomClaims{}),
			ExpectedCode: http.StatusInternalServerError,
			ExpectedFunc: func() {
				s.service.EXPECT().CheckPassword(gomock.Any(), gomock.Any()).Return(false, errors.New("error"))
			},
		},
		{
			Name:    "ok",
			Request: request,
			Body: payload.CheckPasswordPayload{
				CheckPassword: "password",
			},
			Token:        jwt.NewWithClaims(jwt.SigningMethodHS256, &helper.JWTCustomClaims{}),
			ExpectedCode: http.StatusOK,
			ExpectedFunc: func() {
				s.service.EXPECT().CheckPassword(gomock.Any(), gomock.Any()).Return(true, nil)
			},
		},
	}

	for _, testCase := range testCases {
		if testCase.ExpectedFunc != nil {
			testCase.ExpectedFunc()
		}

		s.T().Run(testCase.Name, func(t *testing.T) {
			ctx, rec := testhelper.NewContext(testCase)
			s.NoError(s.handler.CheckPassword(ctx))
			s.Equal(testCase.ExpectedCode, rec.Code)
		})
	}
}
