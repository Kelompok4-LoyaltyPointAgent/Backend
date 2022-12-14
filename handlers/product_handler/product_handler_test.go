package product_handler

import (
	"errors"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kelompok4-loyaltypointagent/backend/dto/response"
	"github.com/kelompok4-loyaltypointagent/backend/services/mocks"
	"github.com/kelompok4-loyaltypointagent/backend/testhelper"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type productHandlerSuite struct {
	suite.Suite

	ctrl    *gomock.Controller
	service *mocks.MockProductService
	handler ProductHandler
}

func (s *productHandlerSuite) SetupSuite() {
	s.ctrl = gomock.NewController(s.T())
	s.service = mocks.NewMockProductService(s.ctrl)
	s.handler = NewProductHandler(s.service)
}

func (s *productHandlerSuite) TearDownSuite() {
	s.ctrl.Finish()
}

func TestProductHandlerSuite(t *testing.T) {
	suite.Run(t, new(productHandlerSuite))
}

func (s *productHandlerSuite) TestGetProductsWithCredits() {
	request := testhelper.Request{
		Method:      http.MethodGet,
		Path:        "/api/v1/products/credits",
		ContentType: echo.MIMEApplicationJSON,
	}

	testCases := []testhelper.HTTPTestCase{
		{
			Name: "bad request",
			Request: testhelper.Request{
				Method:      request.Method,
				ContentType: request.ContentType,
				Path:        request.Path + "?provider=any&recommended=any",
			},
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name:         "internal server error",
			Request:      request,
			ExpectedCode: http.StatusInternalServerError,
			ExpectedFunc: func() {
				s.service.EXPECT().FindAllWithCredits().Return(&[]response.ProductWithCreditResponse{}, errors.New("error"))
			},
		},
		{
			Name:         "ok",
			Request:      request,
			ExpectedCode: http.StatusOK,
			ExpectedFunc: func() {
				s.service.EXPECT().FindAllWithCredits().Return(&[]response.ProductWithCreditResponse{}, nil)
			},
		},
	}

	for _, testCase := range testCases {
		if testCase.ExpectedFunc != nil {
			testCase.ExpectedFunc()
		}

		s.T().Run(testCase.Name, func(t *testing.T) {
			ctx, rec := testhelper.NewContext(testCase)
			s.NoError(s.handler.GetProductsWithCredits(ctx))
			s.Equal(testCase.ExpectedCode, rec.Code)
		})
	}
}

func (s *productHandlerSuite) TestGetProductWithCredit() {
	request := testhelper.Request{
		Method:      http.MethodGet,
		Path:        "/api/v1/products/credits/:id",
		ContentType: echo.MIMEApplicationJSON,
	}

	testCases := []testhelper.HTTPTestCase{
		{
			Name:         "internal server error",
			Request:      request,
			ExpectedCode: http.StatusInternalServerError,
			ExpectedFunc: func() {
				s.service.EXPECT().FindByIDWithCredit(gomock.Any()).Return(&response.ProductWithCreditResponse{}, errors.New("error"))
			},
		},
		{
			Name:         "ok",
			Request:      request,
			ExpectedCode: http.StatusOK,
			ExpectedFunc: func() {
				s.service.EXPECT().FindByIDWithCredit(gomock.Any()).Return(&response.ProductWithCreditResponse{}, nil)
			},
		},
	}

	for _, testCase := range testCases {
		if testCase.ExpectedFunc != nil {
			testCase.ExpectedFunc()
		}

		s.T().Run(testCase.Name, func(t *testing.T) {
			ctx, rec := testhelper.NewContext(testCase)
			s.NoError(s.handler.GetProductWithCredit(ctx))
			s.Equal(testCase.ExpectedCode, rec.Code)
		})
	}
}

func (s *productHandlerSuite) TestGetProductByProviderWithCredits() {
	request := testhelper.Request{
		Method:      http.MethodGet,
		Path:        "/api/v1/products/credits?provider=any",
		ContentType: echo.MIMEApplicationJSON,
	}

	testCases := []testhelper.HTTPTestCase{
		{
			Name:         "internal server error",
			Request:      request,
			ExpectedCode: http.StatusInternalServerError,
			ExpectedFunc: func() {
				s.service.EXPECT().FindByProviderWithCredit(gomock.Any()).Return(&[]response.ProductWithCreditResponse{}, errors.New("error"))
			},
		},
		{
			Name:         "ok",
			Request:      request,
			ExpectedCode: http.StatusOK,
			ExpectedFunc: func() {
				s.service.EXPECT().FindByProviderWithCredit(gomock.Any()).Return(&[]response.ProductWithCreditResponse{}, nil)
			},
		},
	}

	for _, testCase := range testCases {
		if testCase.ExpectedFunc != nil {
			testCase.ExpectedFunc()
		}

		s.T().Run(testCase.Name, func(t *testing.T) {
			ctx, rec := testhelper.NewContext(testCase)
			s.NoError(s.handler.GetProductByProviderWithCredits("any", ctx))
			s.Equal(testCase.ExpectedCode, rec.Code)
		})
	}
}

func (s *productHandlerSuite) TestGetProductByRecommendedWithCredits() {
	request := testhelper.Request{
		Method:      http.MethodGet,
		Path:        "/api/v1/products/credits?recommended=true",
		ContentType: echo.MIMEApplicationJSON,
	}

	testCases := []testhelper.HTTPTestCase{
		{
			Name:         "internal server error",
			Request:      request,
			ExpectedCode: http.StatusInternalServerError,
			ExpectedFunc: func() {
				s.service.EXPECT().FindByRecommendedWithCredit().Return(&[]response.ProductWithCreditResponse{}, errors.New("error"))
			},
		},
		{
			Name:         "ok",
			Request:      request,
			ExpectedCode: http.StatusOK,
			ExpectedFunc: func() {
				s.service.EXPECT().FindByRecommendedWithCredit().Return(&[]response.ProductWithCreditResponse{}, nil)
			},
		},
	}

	for _, testCase := range testCases {
		if testCase.ExpectedFunc != nil {
			testCase.ExpectedFunc()
		}

		s.T().Run(testCase.Name, func(t *testing.T) {
			ctx, rec := testhelper.NewContext(testCase)
			s.NoError(s.handler.GetProductByRecommendedWithCredits(ctx))
			s.Equal(testCase.ExpectedCode, rec.Code)
		})
	}
}

func (s *productHandlerSuite) TestCreateProductWithCredit() {
	request := testhelper.Request{
		Method:      http.MethodPost,
		Path:        "/api/v1/products/credits",
		ContentType: echo.MIMEApplicationForm,
	}

	testCases := []testhelper.HTTPTestCase{
		{
			Name:         "bad request",
			Request:      request,
			Body:         "invalid body",
			ExpectedCode: http.StatusBadRequest,
		},
	}

	for _, testCase := range testCases {
		if testCase.ExpectedFunc != nil {
			testCase.ExpectedFunc()
		}

		s.T().Run(testCase.Name, func(t *testing.T) {
			ctx, rec := testhelper.NewContext(testCase)
			s.NoError(s.handler.CreateProductWithCredit(ctx))
			s.Equal(testCase.ExpectedCode, rec.Code)
		})
	}
}

func (s *productHandlerSuite) TestUpdateProductWithCredit() {
	request := testhelper.Request{
		Method:      http.MethodPut,
		Path:        "/api/v1/products/credits",
		ContentType: echo.MIMEApplicationForm,
	}

	testCases := []testhelper.HTTPTestCase{
		{
			Name:         "bad request",
			Request:      request,
			Body:         "invalid body",
			ExpectedCode: http.StatusBadRequest,
		},
	}

	for _, testCase := range testCases {
		if testCase.ExpectedFunc != nil {
			testCase.ExpectedFunc()
		}

		s.T().Run(testCase.Name, func(t *testing.T) {
			ctx, rec := testhelper.NewContext(testCase)
			s.NoError(s.handler.UpdateProductWithCredit(ctx))
			s.Equal(testCase.ExpectedCode, rec.Code)
		})
	}
}

func (s *productHandlerSuite) TestDeleteProductWithCredit() {
	request := testhelper.Request{
		Method:      http.MethodDelete,
		Path:        "/api/v1/products/credits/:id",
		ContentType: echo.MIMEApplicationForm,
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
				s.service.EXPECT().DeleteProductWithCredit(gomock.Any()).Return(errors.New("error"))
			},
		},
		{
			Name:         "ok",
			Request:      request,
			ExpectedCode: http.StatusOK,
			ExpectedFunc: func() {
				s.service.EXPECT().DeleteProductWithCredit(gomock.Any()).Return(nil)
			},
		},
	}

	for _, testCase := range testCases {
		if testCase.ExpectedFunc != nil {
			testCase.ExpectedFunc()
		}

		s.T().Run(testCase.Name, func(t *testing.T) {
			ctx, rec := testhelper.NewContext(testCase)
			s.NoError(s.handler.DeleteProductWithCredit(ctx))
			s.Equal(testCase.ExpectedCode, rec.Code)
		})
	}
}

func (s *productHandlerSuite) TestGetProductsWithPackages() {
	request := testhelper.Request{
		Method:      http.MethodGet,
		Path:        "/api/v1/products/packages",
		ContentType: echo.MIMEApplicationJSON,
	}

	testCases := []testhelper.HTTPTestCase{
		{
			Name: "bad request",
			Request: testhelper.Request{
				Method:      request.Method,
				ContentType: request.ContentType,
				Path:        request.Path + "?provider=any&recommended=any",
			},
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name:         "internal server error",
			Request:      request,
			ExpectedCode: http.StatusInternalServerError,
			ExpectedFunc: func() {
				s.service.EXPECT().FindAllWithPackages().Return(&[]response.ProductWithPackagesResponse{}, errors.New("error"))
			},
		},
		{
			Name:         "ok",
			Request:      request,
			ExpectedCode: http.StatusOK,
			ExpectedFunc: func() {
				s.service.EXPECT().FindAllWithPackages().Return(&[]response.ProductWithPackagesResponse{}, nil)
			},
		},
	}

	for _, testCase := range testCases {
		if testCase.ExpectedFunc != nil {
			testCase.ExpectedFunc()
		}

		s.T().Run(testCase.Name, func(t *testing.T) {
			ctx, rec := testhelper.NewContext(testCase)
			s.NoError(s.handler.GetProductsWithPackages(ctx))
			s.Equal(testCase.ExpectedCode, rec.Code)
		})
	}
}

func (s *productHandlerSuite) TestGetProductWithPackage() {
	request := testhelper.Request{
		Method:      http.MethodGet,
		Path:        "/api/v1/products/packages/:id",
		ContentType: echo.MIMEApplicationJSON,
	}

	testCases := []testhelper.HTTPTestCase{
		{
			Name:         "internal server error",
			Request:      request,
			ExpectedCode: http.StatusInternalServerError,
			ExpectedFunc: func() {
				s.service.EXPECT().FindByIDWithPackages(gomock.Any()).Return(&response.ProductWithPackagesResponse{}, errors.New("error"))
			},
		},
		{
			Name:         "ok",
			Request:      request,
			ExpectedCode: http.StatusOK,
			ExpectedFunc: func() {
				s.service.EXPECT().FindByIDWithPackages(gomock.Any()).Return(&response.ProductWithPackagesResponse{}, nil)
			},
		},
	}

	for _, testCase := range testCases {
		if testCase.ExpectedFunc != nil {
			testCase.ExpectedFunc()
		}

		s.T().Run(testCase.Name, func(t *testing.T) {
			ctx, rec := testhelper.NewContext(testCase)
			s.NoError(s.handler.GetProductWithPackage(ctx))
			s.Equal(testCase.ExpectedCode, rec.Code)
		})
	}
}

func (s *productHandlerSuite) TestGetProductByProviderWithPackages() {
	request := testhelper.Request{
		Method:      http.MethodGet,
		Path:        "/api/v1/products/packages?provider=any",
		ContentType: echo.MIMEApplicationJSON,
	}

	testCases := []testhelper.HTTPTestCase{
		{
			Name:         "internal server error",
			Request:      request,
			ExpectedCode: http.StatusInternalServerError,
			ExpectedFunc: func() {
				s.service.EXPECT().FindByProviderWithPackages(gomock.Any()).Return(&[]response.ProductWithPackagesResponse{}, errors.New("error"))
			},
		},
		{
			Name:         "ok",
			Request:      request,
			ExpectedCode: http.StatusOK,
			ExpectedFunc: func() {
				s.service.EXPECT().FindByProviderWithPackages(gomock.Any()).Return(&[]response.ProductWithPackagesResponse{}, nil)
			},
		},
	}

	for _, testCase := range testCases {
		if testCase.ExpectedFunc != nil {
			testCase.ExpectedFunc()
		}

		s.T().Run(testCase.Name, func(t *testing.T) {
			ctx, rec := testhelper.NewContext(testCase)
			s.NoError(s.handler.GetProductByProviderWithPackages("any", ctx))
			s.Equal(testCase.ExpectedCode, rec.Code)
		})
	}
}

func (s *productHandlerSuite) TestGetProductByRecommendedWithPackages() {
	request := testhelper.Request{
		Method:      http.MethodGet,
		Path:        "/api/v1/products/packages?recommended=true",
		ContentType: echo.MIMEApplicationJSON,
	}

	testCases := []testhelper.HTTPTestCase{
		{
			Name:         "internal server error",
			Request:      request,
			ExpectedCode: http.StatusInternalServerError,
			ExpectedFunc: func() {
				s.service.EXPECT().FindByRecommendedWithPackages().Return(&[]response.ProductWithPackagesResponse{}, errors.New("error"))
			},
		},
		{
			Name:         "ok",
			Request:      request,
			ExpectedCode: http.StatusOK,
			ExpectedFunc: func() {
				s.service.EXPECT().FindByRecommendedWithPackages().Return(&[]response.ProductWithPackagesResponse{}, nil)
			},
		},
	}

	for _, testCase := range testCases {
		if testCase.ExpectedFunc != nil {
			testCase.ExpectedFunc()
		}

		s.T().Run(testCase.Name, func(t *testing.T) {
			ctx, rec := testhelper.NewContext(testCase)
			s.NoError(s.handler.GetProductByRecommendedWithPackages(ctx))
			s.Equal(testCase.ExpectedCode, rec.Code)
		})
	}
}

func (s *productHandlerSuite) TestCreateProductWithPackage() {
	request := testhelper.Request{
		Method:      http.MethodPost,
		Path:        "/api/v1/products/packages",
		ContentType: echo.MIMEApplicationForm,
	}

	testCases := []testhelper.HTTPTestCase{
		{
			Name:         "bad request",
			Request:      request,
			Body:         "invalid body",
			ExpectedCode: http.StatusBadRequest,
		},
	}

	for _, testCase := range testCases {
		if testCase.ExpectedFunc != nil {
			testCase.ExpectedFunc()
		}

		s.T().Run(testCase.Name, func(t *testing.T) {
			ctx, rec := testhelper.NewContext(testCase)
			s.NoError(s.handler.CreateProductWithPackage(ctx))
			s.Equal(testCase.ExpectedCode, rec.Code)
		})
	}
}

func (s *productHandlerSuite) TestUpdateProductWithPackage() {
	request := testhelper.Request{
		Method:      http.MethodPut,
		Path:        "/api/v1/products/packages",
		ContentType: echo.MIMEApplicationForm,
	}

	testCases := []testhelper.HTTPTestCase{
		{
			Name:         "bad request",
			Request:      request,
			Body:         "invalid body",
			ExpectedCode: http.StatusBadRequest,
		},
	}

	for _, testCase := range testCases {
		if testCase.ExpectedFunc != nil {
			testCase.ExpectedFunc()
		}

		s.T().Run(testCase.Name, func(t *testing.T) {
			ctx, rec := testhelper.NewContext(testCase)
			s.NoError(s.handler.UpdateProductWithPackage(ctx))
			s.Equal(testCase.ExpectedCode, rec.Code)
		})
	}
}

func (s *productHandlerSuite) TestDeleteProductWithPackage() {
	request := testhelper.Request{
		Method:      http.MethodDelete,
		Path:        "/api/v1/products/packages/:id",
		ContentType: echo.MIMEApplicationForm,
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
				s.service.EXPECT().DeleteProductWithPackages(gomock.Any()).Return(errors.New("error"))
			},
		},
		{
			Name:         "ok",
			Request:      request,
			ExpectedCode: http.StatusOK,
			ExpectedFunc: func() {
				s.service.EXPECT().DeleteProductWithPackages(gomock.Any()).Return(nil)
			},
		},
	}

	for _, testCase := range testCases {
		if testCase.ExpectedFunc != nil {
			testCase.ExpectedFunc()
		}

		s.T().Run(testCase.Name, func(t *testing.T) {
			ctx, rec := testhelper.NewContext(testCase)
			s.NoError(s.handler.DeleteProductWithPackage(ctx))
			s.Equal(testCase.ExpectedCode, rec.Code)
		})
	}
}
