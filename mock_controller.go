// Code generated by mockery v2.14.0. DO NOT EDIT.

package server

import mock "github.com/stretchr/testify/mock"

// MockController is an autogenerated mock type for the Controller type
type MockController struct {
	mock.Mock
}

// Mount provides a mock function with given fields: r
func (_m *MockController) Mount(r *Router) {
	_m.Called(r)
}

type mockConstructorTestingTNewMockController interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockController creates a new instance of MockController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockController(t mockConstructorTestingTNewMockController) *MockController {
	mock := &MockController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
