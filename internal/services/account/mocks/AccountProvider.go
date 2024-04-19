// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/karmaplush/simple-diet-tracker/internal/domain/models"
	mock "github.com/stretchr/testify/mock"
)

// AccountProvider is an autogenerated mock type for the AccountProvider type
type AccountProvider struct {
	mock.Mock
}

// AccountById provides a mock function with given fields: ctx, accountId
func (_m *AccountProvider) AccountById(ctx context.Context, accountId int64) (models.Account, error) {
	ret := _m.Called(ctx, accountId)

	if len(ret) == 0 {
		panic("no return value specified for AccountById")
	}

	var r0 models.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (models.Account, error)); ok {
		return rf(ctx, accountId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) models.Account); ok {
		r0 = rf(ctx, accountId)
	} else {
		r0 = ret.Get(0).(models.Account)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, accountId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AccountByUserId provides a mock function with given fields: ctx, userId
func (_m *AccountProvider) AccountByUserId(ctx context.Context, userId int64) (models.Account, error) {
	ret := _m.Called(ctx, userId)

	if len(ret) == 0 {
		panic("no return value specified for AccountByUserId")
	}

	var r0 models.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (models.Account, error)); ok {
		return rf(ctx, userId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) models.Account); ok {
		r0 = rf(ctx, userId)
	} else {
		r0 = ret.Get(0).(models.Account)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAccountProvider creates a new instance of AccountProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAccountProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *AccountProvider {
	mock := &AccountProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}