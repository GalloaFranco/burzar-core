// Code generated by mockery v2.49.1. DO NOT EDIT.

package mocks

import (
	domain "github.com/GalloaFranco/burzar-core/internal/core/domain"
	mock "github.com/stretchr/testify/mock"
)

// CountryRisk is an autogenerated mock type for the CountryRisk type
type CountryRisk struct {
	mock.Mock
}

// Get provides a mock function with given fields:
func (_m *CountryRisk) Get() (*domain.CountryRisk, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *domain.CountryRisk
	var r1 error
	if rf, ok := ret.Get(0).(func() (*domain.CountryRisk, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *domain.CountryRisk); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.CountryRisk)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCountryRisk creates a new instance of CountryRisk. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCountryRisk(t interface {
	mock.TestingT
	Cleanup(func())
}) *CountryRisk {
	mock := &CountryRisk{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
