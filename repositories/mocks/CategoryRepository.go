// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	models "mini-cms-api/models"

	mock "github.com/stretchr/testify/mock"
)

// CategoryRepository is an autogenerated mock type for the CategoryRepository type
type CategoryRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: categoryReq
func (_m *CategoryRepository) Create(categoryReq models.CategoryRequest) (models.Category, error) {
	ret := _m.Called(categoryReq)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 models.Category
	var r1 error
	if rf, ok := ret.Get(0).(func(models.CategoryRequest) (models.Category, error)); ok {
		return rf(categoryReq)
	}
	if rf, ok := ret.Get(0).(func(models.CategoryRequest) models.Category); ok {
		r0 = rf(categoryReq)
	} else {
		r0 = ret.Get(0).(models.Category)
	}

	if rf, ok := ret.Get(1).(func(models.CategoryRequest) error); ok {
		r1 = rf(categoryReq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *CategoryRepository) Delete(id string) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *CategoryRepository) GetAll() ([]models.Category, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []models.Category
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]models.Category, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []models.Category); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Category)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: id
func (_m *CategoryRepository) GetByID(id string) (models.Category, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 models.Category
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (models.Category, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) models.Category); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Category)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: categoryReq, id
func (_m *CategoryRepository) Update(categoryReq models.CategoryRequest, id string) (models.Category, error) {
	ret := _m.Called(categoryReq, id)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 models.Category
	var r1 error
	if rf, ok := ret.Get(0).(func(models.CategoryRequest, string) (models.Category, error)); ok {
		return rf(categoryReq, id)
	}
	if rf, ok := ret.Get(0).(func(models.CategoryRequest, string) models.Category); ok {
		r0 = rf(categoryReq, id)
	} else {
		r0 = ret.Get(0).(models.Category)
	}

	if rf, ok := ret.Get(1).(func(models.CategoryRequest, string) error); ok {
		r1 = rf(categoryReq, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCategoryRepository creates a new instance of CategoryRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCategoryRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *CategoryRepository {
	mock := &CategoryRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}