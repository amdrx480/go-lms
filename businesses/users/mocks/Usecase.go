// Code generated by mockery v2.53.0. DO NOT EDIT.

package mocks

import (
	context "context"

	users "github.com/amdrx480/go-lms/businesses/users"
	mock "github.com/stretchr/testify/mock"
)

// UseCase is an autogenerated mock type for the UseCase type
type UseCase struct {
	mock.Mock
}

// GetUserProfile provides a mock function with given fields: ctx
func (_m *UseCase) GetUserProfile(ctx context.Context) (users.Domain, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetUserProfile")
	}

	var r0 users.Domain
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (users.Domain, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) users.Domain); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: ctx, userDomain
func (_m *UseCase) Login(ctx context.Context, userDomain *users.Domain) (string, error) {
	ret := _m.Called(ctx, userDomain)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *users.Domain) (string, error)); ok {
		return rf(ctx, userDomain)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *users.Domain) string); ok {
		r0 = rf(ctx, userDomain)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *users.Domain) error); ok {
		r1 = rf(ctx, userDomain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: ctx, userDomain
func (_m *UseCase) Register(ctx context.Context, userDomain *users.Domain) (users.Domain, error) {
	ret := _m.Called(ctx, userDomain)

	if len(ret) == 0 {
		panic("no return value specified for Register")
	}

	var r0 users.Domain
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *users.Domain) (users.Domain, error)); ok {
		return rf(ctx, userDomain)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *users.Domain) users.Domain); ok {
		r0 = rf(ctx, userDomain)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *users.Domain) error); ok {
		r1 = rf(ctx, userDomain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUseCase creates a new instance of UseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *UseCase {
	mock := &UseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
