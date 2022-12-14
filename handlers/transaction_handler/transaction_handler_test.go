package transaction_handler

import (
	"errors"
	"net/http"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/golang/mock/gomock"
	"github.com/kelompok4-loyaltypointagent/backend/constant"
	"github.com/kelompok4-loyaltypointagent/backend/dto/payload"
	"github.com/kelompok4-loyaltypointagent/backend/dto/response"
	"github.com/kelompok4-loyaltypointagent/backend/helper"
	"github.com/kelompok4-loyaltypointagent/backend/services/mocks"
	"github.com/kelompok4-loyaltypointagent/backend/testhelper"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type transactionHandlerSuite struct {
	suite.Suite

	ctrl    *gomock.Controller
	service *mocks.MockTransactionService
	handler TransactionHandler
}

func (s *transactionHandlerSuite) SetupSuite() {
	s.ctrl = gomock.NewController(s.T())
	s.service = mocks.NewMockTransactionService(s.ctrl)
	s.handler = NewTransactionHandler(s.service)
}

func (s *transactionHandlerSuite) TearDownSuite() {
	s.ctrl.Finish()
}

func TestTransactionHandlerSuite(t *testing.T) {
	suite.Run(t, new(transactionHandlerSuite))
}

func (s *transactionHandlerSuite) TestGetTransactions() {
	request := testhelper.Request{
		Method:      http.MethodGet,
		Path:        "/api/v1/transactions",
		ContentType: echo.MIMEApplicationJSON,
	}

	testCases := []testhelper.HTTPTestCase{
		{
			Name: "bad request",
			Request: testhelper.Request{
				Method:      request.Method,
				ContentType: request.ContentType,
				Path:        request.Path + "?type=any",
			},
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name:         "internal server error",
			Request:      request,
			Token:        jwt.NewWithClaims(jwt.SigningMethodHS256, helper.JWTCustomClaims{}),
			ExpectedCode: http.StatusInternalServerError,
			ExpectedFunc: func() {
				s.service.EXPECT().FindAllDetail(gomock.Any(), gomock.Any()).Return(&[]response.TransactionResponse{}, errors.New("error"))
			},
		},
		{
			Name:         "ok",
			Request:      request,
			Token:        jwt.NewWithClaims(jwt.SigningMethodHS256, helper.JWTCustomClaims{}),
			ExpectedCode: http.StatusOK,
			ExpectedFunc: func() {
				s.service.EXPECT().FindAllDetail(gomock.Any(), gomock.Any()).Return(&[]response.TransactionResponse{}, nil)
			},
		},
	}

	for _, testCase := range testCases {
		if testCase.ExpectedFunc != nil {
			testCase.ExpectedFunc()
		}

		s.T().Run(testCase.Name, func(t *testing.T) {
			ctx, rec := testhelper.NewContext(testCase)
			s.NoError(s.handler.GetTransactions(ctx))
			s.Equal(testCase.ExpectedCode, rec.Code)
		})
	}
}

func (s *transactionHandlerSuite) TestGetTransaction() {
	request := testhelper.Request{
		Method:      http.MethodGet,
		Path:        "/api/v1/transactions/:id",
		ContentType: echo.MIMEApplicationJSON,
		PathParam: &testhelper.PathParam{
			Names:  []string{"id"},
			Values: []string{"uuid"},
		},
	}

	testCases := []testhelper.HTTPTestCase{
		{
			Name:         "not found",
			Request:      request,
			Token:        jwt.NewWithClaims(jwt.SigningMethodHS256, helper.JWTCustomClaims{}),
			ExpectedCode: http.StatusNotFound,
			ExpectedFunc: func() {
				s.service.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(&response.TransactionResponse{}, errors.New("error"))
			},
		},
		{
			Name:         "unauthorized",
			Request:      request,
			Token:        jwt.NewWithClaims(jwt.SigningMethodHS256, helper.JWTCustomClaims{}),
			ExpectedCode: http.StatusUnauthorized,
			ExpectedFunc: func() {
				s.service.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(&response.TransactionResponse{}, errors.New("forbidden"))
			},
		},
		{
			Name:         "ok",
			Request:      request,
			Token:        jwt.NewWithClaims(jwt.SigningMethodHS256, helper.JWTCustomClaims{}),
			ExpectedCode: http.StatusOK,
			ExpectedFunc: func() {
				s.service.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(&response.TransactionResponse{}, nil)
			},
		},
	}

	for _, testCase := range testCases {
		if testCase.ExpectedFunc != nil {
			testCase.ExpectedFunc()
		}

		s.T().Run(testCase.Name, func(t *testing.T) {
			ctx, rec := testhelper.NewContext(testCase)
			s.NoError(s.handler.GetTransaction(ctx))
			s.Equal(testCase.ExpectedCode, rec.Code)
		})
	}
}

func (s *transactionHandlerSuite) TestCreateTransaction() {
	request := testhelper.Request{
		Method:      http.MethodPost,
		Path:        "/api/v1/transactions",
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
			Body:         payload.TransactionPayload{},
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name:    "internal server error",
			Request: request,
			Body: payload.TransactionPayload{
				UserID:    "uuid",
				ProductID: "uuid",
				Amount:    1,
				Method:    "gopay",
				Number:    "080123456789",
				Email:     "user@example.com",
				Status:    constant.TransactionStatusPending,
				Type:      constant.TransactionTypePurchase,
			},
			Token:        jwt.NewWithClaims(jwt.SigningMethodHS256, helper.JWTCustomClaims{}),
			ExpectedCode: http.StatusInternalServerError,
			ExpectedFunc: func() {
				s.service.EXPECT().Create(gomock.Any(), gomock.Any()).Return(&response.TransactionResponse{}, errors.New("error"))
			},
		},
		{
			Name:    "ok",
			Request: request,
			Body: payload.TransactionPayload{
				UserID:    "uuid",
				ProductID: "uuid",
				Amount:    1,
				Method:    "gopay",
				Number:    "080123456789",
				Email:     "user@example.com",
				Status:    constant.TransactionStatusPending,
				Type:      constant.TransactionTypePurchase,
			},
			Token:        jwt.NewWithClaims(jwt.SigningMethodHS256, helper.JWTCustomClaims{}),
			ExpectedCode: http.StatusOK,
			ExpectedFunc: func() {
				s.service.EXPECT().Create(gomock.Any(), gomock.Any()).Return(&response.TransactionResponse{}, nil)
			},
		},
	}

	for _, testCase := range testCases {
		if testCase.ExpectedFunc != nil {
			testCase.ExpectedFunc()
		}

		s.T().Run(testCase.Name, func(t *testing.T) {
			ctx, rec := testhelper.NewContext(testCase)
			s.NoError(s.handler.CreateTransaction(ctx))
			s.Equal(testCase.ExpectedCode, rec.Code)
		})
	}
}

func (s *transactionHandlerSuite) TestUpdateTransaction() {
	request := testhelper.Request{
		Method:      http.MethodPut,
		Path:        "/api/v1/transactions/:id",
		ContentType: echo.MIMEApplicationJSON,
		PathParam: &testhelper.PathParam{
			Names:  []string{"id"},
			Values: []string{"uuid"},
		},
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
			Body:         payload.TransactionPayload{},
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name:    "internal server error",
			Request: request,
			Body: payload.TransactionPayload{
				UserID:    "uuid",
				ProductID: "uuid",
				Amount:    1,
				Method:    "gopay",
				Number:    "080123456789",
				Email:     "user@example.com",
				Status:    constant.TransactionStatusPending,
				Type:      constant.TransactionTypePurchase,
			},
			ExpectedCode: http.StatusInternalServerError,
			ExpectedFunc: func() {
				s.service.EXPECT().Update(gomock.Any(), gomock.Any()).Return(&response.TransactionResponse{}, errors.New("error"))
			},
		},
		{
			Name:    "ok",
			Request: request,
			Body: payload.TransactionPayload{
				UserID:    "uuid",
				ProductID: "uuid",
				Amount:    1,
				Method:    "gopay",
				Number:    "080123456789",
				Email:     "user@example.com",
				Status:    constant.TransactionStatusPending,
				Type:      constant.TransactionTypePurchase,
			},
			ExpectedCode: http.StatusOK,
			ExpectedFunc: func() {
				s.service.EXPECT().Update(gomock.Any(), gomock.Any()).Return(&response.TransactionResponse{}, nil)
			},
		},
	}

	for _, testCase := range testCases {
		if testCase.ExpectedFunc != nil {
			testCase.ExpectedFunc()
		}

		s.T().Run(testCase.Name, func(t *testing.T) {
			ctx, rec := testhelper.NewContext(testCase)
			s.NoError(s.handler.UpdateTransaction(ctx))
			s.Equal(testCase.ExpectedCode, rec.Code)
		})
	}
}

func (s *transactionHandlerSuite) TestDeleteTransaction() {
	request := testhelper.Request{
		Method:      http.MethodDelete,
		Path:        "/api/v1/transactions/:id",
		ContentType: echo.MIMEApplicationJSON,
		PathParam: &testhelper.PathParam{
			Names:  []string{"id"},
			Values: []string{"uuid"},
		},
	}

	testCases := []testhelper.HTTPTestCase{
		{
			Name:         "internal server error",
			Request:      request,
			ExpectedCode: http.StatusInternalServerError,
			ExpectedFunc: func() {
				s.service.EXPECT().Delete(gomock.Any()).Return(errors.New("error"))
			},
		},
		{
			Name:         "ok",
			Request:      request,
			ExpectedCode: http.StatusOK,
			ExpectedFunc: func() {
				s.service.EXPECT().Delete(gomock.Any()).Return(nil)
			},
		},
	}

	for _, testCase := range testCases {
		if testCase.ExpectedFunc != nil {
			testCase.ExpectedFunc()
		}

		s.T().Run(testCase.Name, func(t *testing.T) {
			ctx, rec := testhelper.NewContext(testCase)
			s.NoError(s.handler.DeleteTransaction(ctx))
			s.Equal(testCase.ExpectedCode, rec.Code)
		})
	}
}

func (s *transactionHandlerSuite) TestCancelTransaction() {
	request := testhelper.Request{
		Method:      http.MethodPost,
		Path:        "/api/v1/transactions/cancel/:id",
		ContentType: echo.MIMEApplicationJSON,
		PathParam: &testhelper.PathParam{
			Names:  []string{"id"},
			Values: []string{"uuid"},
		},
	}

	testCases := []testhelper.HTTPTestCase{
		{
			Name:         "internal server error",
			Request:      request,
			ExpectedCode: http.StatusInternalServerError,
			ExpectedFunc: func() {
				s.service.EXPECT().Cancel(gomock.Any()).Return(&response.TransactionResponse{}, errors.New("error"))
			},
		},
		{
			Name:         "ok",
			Request:      request,
			ExpectedCode: http.StatusOK,
			ExpectedFunc: func() {
				s.service.EXPECT().Cancel(gomock.Any()).Return(&response.TransactionResponse{}, nil)
			},
		},
	}

	for _, testCase := range testCases {
		if testCase.ExpectedFunc != nil {
			testCase.ExpectedFunc()
		}

		s.T().Run(testCase.Name, func(t *testing.T) {
			ctx, rec := testhelper.NewContext(testCase)
			s.NoError(s.handler.CancelTransaction(ctx))
			s.Equal(testCase.ExpectedCode, rec.Code)
		})
	}
}

func (s *transactionHandlerSuite) TestTransactionWebhook() {
	request := testhelper.Request{
		Method:      http.MethodPost,
		Path:        "/api/v1/transactions/webhook",
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
			Name:         "internal server error",
			Request:      request,
			Body:         map[string]any{"transaction_status": "success"},
			ExpectedCode: http.StatusInternalServerError,
			ExpectedFunc: func() {
				s.service.EXPECT().CallbackXendit(gomock.Any()).Return(false, errors.New("error"))
			},
		},
		{
			Name:         "ok",
			Request:      request,
			Body:         map[string]any{"transaction_status": "success"},
			ExpectedCode: http.StatusOK,
			ExpectedFunc: func() {
				s.service.EXPECT().CallbackXendit(gomock.Any()).Return(true, nil)
			},
		},
	}

	for _, testCase := range testCases {
		if testCase.ExpectedFunc != nil {
			testCase.ExpectedFunc()
		}

		s.T().Run(testCase.Name, func(t *testing.T) {
			ctx, rec := testhelper.NewContext(testCase)
			s.NoError(s.handler.TransactionWebhook(ctx))
			s.Equal(testCase.ExpectedCode, rec.Code)
		})
	}
}
