// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "github.com/aait.backend.g5.main/backend/Domain/Models"
)

// PasswordUsecase is an autogenerated mock type for the PasswordUsecase type
type PasswordUsecase struct {
	mock.Mock
}

// GenerateResetURL provides a mock function with given fields: ctx, email
func (_m *PasswordUsecase) GenerateResetURL(ctx context.Context, email string) (string, *models.ErrorResponse) {
	ret := _m.Called(ctx, email)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 *models.ErrorResponse
	if rf, ok := ret.Get(1).(func(context.Context, string) *models.ErrorResponse); ok {
		r1 = rf(ctx, email)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*models.ErrorResponse)
		}
	}

	return r0, r1
}

// SendResetEmail provides a mock function with given fields: ctx, email, resetURL
func (_m *PasswordUsecase) SendResetEmail(ctx context.Context, email string, resetURL string) *models.ErrorResponse {
	ret := _m.Called(ctx, email, resetURL)

	var r0 *models.ErrorResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *models.ErrorResponse); ok {
		r0 = rf(ctx, email, resetURL)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ErrorResponse)
		}
	}

	return r0
}

// SetNewUserPassword provides a mock function with given fields: ctx, shortURlCode, password
func (_m *PasswordUsecase) SetNewUserPassword(ctx context.Context, shortURlCode string, password string) *models.ErrorResponse {
	ret := _m.Called(ctx, shortURlCode, password)

	var r0 *models.ErrorResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *models.ErrorResponse); ok {
		r0 = rf(ctx, shortURlCode, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ErrorResponse)
		}
	}

	return r0
}

// SetUpdateUserPassword provides a mock function with given fields: ctx, shortURlCode, password
func (_m *PasswordUsecase) SetUpdateUserPassword(ctx context.Context, shortURlCode string, password string) *models.ErrorResponse {
	ret := _m.Called(ctx, shortURlCode, password)

	var r0 *models.ErrorResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *models.ErrorResponse); ok {
		r0 = rf(ctx, shortURlCode, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ErrorResponse)
		}
	}

	return r0
}

type mockConstructorTestingTNewPasswordUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewPasswordUsecase creates a new instance of PasswordUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPasswordUsecase(t mockConstructorTestingTNewPasswordUsecase) *PasswordUsecase {
	mock := &PasswordUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
