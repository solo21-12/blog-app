// Code generated by MockGen. DO NOT EDIT.
// Source: Domain/Interfaces/url_service_interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	gomock "github.com/golang/mock/gomock"
)

// MockURLService is a mock of URLService interface.
type MockURLService struct {
	ctrl     *gomock.Controller
	recorder *MockURLServiceMockRecorder
}

// MockURLServiceMockRecorder is the mock recorder for MockURLService.
type MockURLServiceMockRecorder struct {
	mock *MockURLService
}

// NewMockURLService creates a new mock instance.
func NewMockURLService(ctrl *gomock.Controller) *MockURLService {
	mock := &MockURLService{ctrl: ctrl}
	mock.recorder = &MockURLServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockURLService) EXPECT() *MockURLServiceMockRecorder {
	return m.recorder
}

// GenerateURL mocks base method.
func (m *MockURLService) GenerateURL(token, purpose string) (string, *models.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateURL", token, purpose)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*models.ErrorResponse)
	return ret0, ret1
}

// GenerateURL indicates an expected call of GenerateURL.
func (mr *MockURLServiceMockRecorder) GenerateURL(token, purpose interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateURL", reflect.TypeOf((*MockURLService)(nil).GenerateURL), token, purpose)
}

// GetURL mocks base method.
func (m *MockURLService) GetURL(short_url_code string) (*models.URL, *models.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetURL", short_url_code)
	ret0, _ := ret[0].(*models.URL)
	ret1, _ := ret[1].(*models.ErrorResponse)
	return ret0, ret1
}

// GetURL indicates an expected call of GetURL.
func (mr *MockURLServiceMockRecorder) GetURL(short_url_code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetURL", reflect.TypeOf((*MockURLService)(nil).GetURL), short_url_code)
}

// RemoveURL mocks base method.
func (m *MockURLService) RemoveURL(short_url_code string) *models.ErrorResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveURL", short_url_code)
	ret0, _ := ret[0].(*models.ErrorResponse)
	return ret0
}

// RemoveURL indicates an expected call of RemoveURL.
func (mr *MockURLServiceMockRecorder) RemoveURL(short_url_code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveURL", reflect.TypeOf((*MockURLService)(nil).RemoveURL), short_url_code)
}

// MockURLServiceRepository is a mock of URLServiceRepository interface.
type MockURLServiceRepository struct {
	ctrl     *gomock.Controller
	recorder *MockURLServiceRepositoryMockRecorder
}

// MockURLServiceRepositoryMockRecorder is the mock recorder for MockURLServiceRepository.
type MockURLServiceRepositoryMockRecorder struct {
	mock *MockURLServiceRepository
}

// NewMockURLServiceRepository creates a new mock instance.
func NewMockURLServiceRepository(ctrl *gomock.Controller) *MockURLServiceRepository {
	mock := &MockURLServiceRepository{ctrl: ctrl}
	mock.recorder = &MockURLServiceRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockURLServiceRepository) EXPECT() *MockURLServiceRepositoryMockRecorder {
	return m.recorder
}

// DeleteURL mocks base method.
func (m *MockURLServiceRepository) DeleteURL(id string, ctx context.Context) *models.ErrorResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteURL", id, ctx)
	ret0, _ := ret[0].(*models.ErrorResponse)
	return ret0
}

// DeleteURL indicates an expected call of DeleteURL.
func (mr *MockURLServiceRepositoryMockRecorder) DeleteURL(id, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteURL", reflect.TypeOf((*MockURLServiceRepository)(nil).DeleteURL), id, ctx)
}

// GetURL mocks base method.
func (m *MockURLServiceRepository) GetURL(short_url_code string, ctx context.Context) (*models.URL, *models.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetURL", short_url_code, ctx)
	ret0, _ := ret[0].(*models.URL)
	ret1, _ := ret[1].(*models.ErrorResponse)
	return ret0, ret1
}

// GetURL indicates an expected call of GetURL.
func (mr *MockURLServiceRepositoryMockRecorder) GetURL(short_url_code, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetURL", reflect.TypeOf((*MockURLServiceRepository)(nil).GetURL), short_url_code, ctx)
}

// SaveURL mocks base method.
func (m *MockURLServiceRepository) SaveURL(url models.URL, ctx context.Context) *models.ErrorResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveURL", url, ctx)
	ret0, _ := ret[0].(*models.ErrorResponse)
	return ret0
}

// SaveURL indicates an expected call of SaveURL.
func (mr *MockURLServiceRepositoryMockRecorder) SaveURL(url, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveURL", reflect.TypeOf((*MockURLServiceRepository)(nil).SaveURL), url, ctx)
}
