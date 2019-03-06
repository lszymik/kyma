// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import mock "github.com/stretchr/testify/mock"

// CSRProvider is an autogenerated mock type for the CSRProvider type
type CSRProvider struct {
	mock.Mock
}

// CreateCSR provides a mock function with given fields: plainSubject
func (_m *CSRProvider) CreateCSR(plainSubject string) (string, error) {
	ret := _m.Called(plainSubject)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(plainSubject)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(plainSubject)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
