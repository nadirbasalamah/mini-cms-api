// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	users "mini-cms-api/businesses/users"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetUserByEmail provides a mock function with given fields: userReq
func (_m *Repository) GetUserByEmail(userReq *users.Domain) (users.Domain, error) {
	ret := _m.Called(userReq)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByEmail")
	}

	var r0 users.Domain
	var r1 error
	if rf, ok := ret.Get(0).(func(*users.Domain) (users.Domain, error)); ok {
		return rf(userReq)
	}
	if rf, ok := ret.Get(0).(func(*users.Domain) users.Domain); ok {
		r0 = rf(userReq)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	if rf, ok := ret.Get(1).(func(*users.Domain) error); ok {
		r1 = rf(userReq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserInfo provides a mock function with given fields: id
func (_m *Repository) GetUserInfo(id string) (users.Domain, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetUserInfo")
	}

	var r0 users.Domain
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (users.Domain, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) users.Domain); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: userReq
func (_m *Repository) Register(userReq *users.Domain) (users.Domain, error) {
	ret := _m.Called(userReq)

	if len(ret) == 0 {
		panic("no return value specified for Register")
	}

	var r0 users.Domain
	var r1 error
	if rf, ok := ret.Get(0).(func(*users.Domain) (users.Domain, error)); ok {
		return rf(userReq)
	}
	if rf, ok := ret.Get(0).(func(*users.Domain) users.Domain); ok {
		r0 = rf(userReq)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	if rf, ok := ret.Get(1).(func(*users.Domain) error); ok {
		r1 = rf(userReq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
