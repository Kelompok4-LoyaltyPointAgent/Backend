// Code generated by MockGen. DO NOT EDIT.
// Source: ./services/product_service/product_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	payload "github.com/kelompok4-loyaltypointagent/backend/dto/payload"
	response "github.com/kelompok4-loyaltypointagent/backend/dto/response"
	models "github.com/kelompok4-loyaltypointagent/backend/models"
)

// MockProductService is a mock of ProductService interface.
type MockProductService struct {
	ctrl     *gomock.Controller
	recorder *MockProductServiceMockRecorder
}

// MockProductServiceMockRecorder is the mock recorder for MockProductService.
type MockProductServiceMockRecorder struct {
	mock *MockProductService
}

// NewMockProductService creates a new mock instance.
func NewMockProductService(ctrl *gomock.Controller) *MockProductService {
	mock := &MockProductService{ctrl: ctrl}
	mock.recorder = &MockProductServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductService) EXPECT() *MockProductServiceMockRecorder {
	return m.recorder
}

// CreateProductWithCredit mocks base method.
func (m *MockProductService) CreateProductWithCredit(payload payload.ProductWithCreditPayload) (*response.ProductWithCreditResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProductWithCredit", payload)
	ret0, _ := ret[0].(*response.ProductWithCreditResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProductWithCredit indicates an expected call of CreateProductWithCredit.
func (mr *MockProductServiceMockRecorder) CreateProductWithCredit(payload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProductWithCredit", reflect.TypeOf((*MockProductService)(nil).CreateProductWithCredit), payload)
}

// CreateProductWithPackages mocks base method.
func (m *MockProductService) CreateProductWithPackages(payload payload.ProductWithPackagesPayload) (*response.ProductWithPackagesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProductWithPackages", payload)
	ret0, _ := ret[0].(*response.ProductWithPackagesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProductWithPackages indicates an expected call of CreateProductWithPackages.
func (mr *MockProductServiceMockRecorder) CreateProductWithPackages(payload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProductWithPackages", reflect.TypeOf((*MockProductService)(nil).CreateProductWithPackages), payload)
}

// DeleteProductWithCredit mocks base method.
func (m *MockProductService) DeleteProductWithCredit(id any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProductWithCredit", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProductWithCredit indicates an expected call of DeleteProductWithCredit.
func (mr *MockProductServiceMockRecorder) DeleteProductWithCredit(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProductWithCredit", reflect.TypeOf((*MockProductService)(nil).DeleteProductWithCredit), id)
}

// DeleteProductWithPackages mocks base method.
func (m *MockProductService) DeleteProductWithPackages(id any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProductWithPackages", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProductWithPackages indicates an expected call of DeleteProductWithPackages.
func (mr *MockProductServiceMockRecorder) DeleteProductWithPackages(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProductWithPackages", reflect.TypeOf((*MockProductService)(nil).DeleteProductWithPackages), id)
}

// FindAllWithCredits mocks base method.
func (m *MockProductService) FindAllWithCredits() (*[]response.ProductWithCreditResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllWithCredits")
	ret0, _ := ret[0].(*[]response.ProductWithCreditResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllWithCredits indicates an expected call of FindAllWithCredits.
func (mr *MockProductServiceMockRecorder) FindAllWithCredits() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllWithCredits", reflect.TypeOf((*MockProductService)(nil).FindAllWithCredits))
}

// FindAllWithPackages mocks base method.
func (m *MockProductService) FindAllWithPackages() (*[]response.ProductWithPackagesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllWithPackages")
	ret0, _ := ret[0].(*[]response.ProductWithPackagesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllWithPackages indicates an expected call of FindAllWithPackages.
func (mr *MockProductServiceMockRecorder) FindAllWithPackages() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllWithPackages", reflect.TypeOf((*MockProductService)(nil).FindAllWithPackages))
}

// FindByIDWithCredit mocks base method.
func (m *MockProductService) FindByIDWithCredit(id any) (*response.ProductWithCreditResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByIDWithCredit", id)
	ret0, _ := ret[0].(*response.ProductWithCreditResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByIDWithCredit indicates an expected call of FindByIDWithCredit.
func (mr *MockProductServiceMockRecorder) FindByIDWithCredit(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByIDWithCredit", reflect.TypeOf((*MockProductService)(nil).FindByIDWithCredit), id)
}

// FindByIDWithPackages mocks base method.
func (m *MockProductService) FindByIDWithPackages(id any) (*response.ProductWithPackagesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByIDWithPackages", id)
	ret0, _ := ret[0].(*response.ProductWithPackagesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByIDWithPackages indicates an expected call of FindByIDWithPackages.
func (mr *MockProductServiceMockRecorder) FindByIDWithPackages(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByIDWithPackages", reflect.TypeOf((*MockProductService)(nil).FindByIDWithPackages), id)
}

// FindByProviderWithCredit mocks base method.
func (m *MockProductService) FindByProviderWithCredit(provider string) (*[]response.ProductWithCreditResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByProviderWithCredit", provider)
	ret0, _ := ret[0].(*[]response.ProductWithCreditResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByProviderWithCredit indicates an expected call of FindByProviderWithCredit.
func (mr *MockProductServiceMockRecorder) FindByProviderWithCredit(provider interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByProviderWithCredit", reflect.TypeOf((*MockProductService)(nil).FindByProviderWithCredit), provider)
}

// FindByProviderWithPackages mocks base method.
func (m *MockProductService) FindByProviderWithPackages(provider string) (*[]response.ProductWithPackagesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByProviderWithPackages", provider)
	ret0, _ := ret[0].(*[]response.ProductWithPackagesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByProviderWithPackages indicates an expected call of FindByProviderWithPackages.
func (mr *MockProductServiceMockRecorder) FindByProviderWithPackages(provider interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByProviderWithPackages", reflect.TypeOf((*MockProductService)(nil).FindByProviderWithPackages), provider)
}

// FindByRecommendedWithCredit mocks base method.
func (m *MockProductService) FindByRecommendedWithCredit() (*[]response.ProductWithCreditResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByRecommendedWithCredit")
	ret0, _ := ret[0].(*[]response.ProductWithCreditResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByRecommendedWithCredit indicates an expected call of FindByRecommendedWithCredit.
func (mr *MockProductServiceMockRecorder) FindByRecommendedWithCredit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByRecommendedWithCredit", reflect.TypeOf((*MockProductService)(nil).FindByRecommendedWithCredit))
}

// FindByRecommendedWithPackages mocks base method.
func (m *MockProductService) FindByRecommendedWithPackages() (*[]response.ProductWithPackagesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByRecommendedWithPackages")
	ret0, _ := ret[0].(*[]response.ProductWithPackagesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByRecommendedWithPackages indicates an expected call of FindByRecommendedWithPackages.
func (mr *MockProductServiceMockRecorder) FindByRecommendedWithPackages() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByRecommendedWithPackages", reflect.TypeOf((*MockProductService)(nil).FindByRecommendedWithPackages))
}

// FindIconByProvider mocks base method.
func (m *MockProductService) FindIconByProvider(provider string) (models.ProductPicture, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindIconByProvider", provider)
	ret0, _ := ret[0].(models.ProductPicture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindIconByProvider indicates an expected call of FindIconByProvider.
func (mr *MockProductServiceMockRecorder) FindIconByProvider(provider interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindIconByProvider", reflect.TypeOf((*MockProductService)(nil).FindIconByProvider), provider)
}

// UpdateProductWithCredit mocks base method.
func (m *MockProductService) UpdateProductWithCredit(payload payload.ProductWithCreditPayload, id any) (*response.ProductWithCreditResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProductWithCredit", payload, id)
	ret0, _ := ret[0].(*response.ProductWithCreditResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProductWithCredit indicates an expected call of UpdateProductWithCredit.
func (mr *MockProductServiceMockRecorder) UpdateProductWithCredit(payload, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProductWithCredit", reflect.TypeOf((*MockProductService)(nil).UpdateProductWithCredit), payload, id)
}

// UpdateProductWithPackages mocks base method.
func (m *MockProductService) UpdateProductWithPackages(payload payload.ProductWithPackagesPayload, id any) (*response.ProductWithPackagesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProductWithPackages", payload, id)
	ret0, _ := ret[0].(*response.ProductWithPackagesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProductWithPackages indicates an expected call of UpdateProductWithPackages.
func (mr *MockProductServiceMockRecorder) UpdateProductWithPackages(payload, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProductWithPackages", reflect.TypeOf((*MockProductService)(nil).UpdateProductWithPackages), payload, id)
}