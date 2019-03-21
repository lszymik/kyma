// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import mock "github.com/stretchr/testify/mock"

import storage "github.com/kyma-project/kyma/components/console-backend-service/internal/domain/content/storage"

// AsyncApiSpecGetter is an autogenerated mock type for the AsyncApiSpecGetter type
type AsyncApiSpecGetter struct {
	mock.Mock
}

// Find provides a mock function with given fields: kind, id
func (_m *AsyncApiSpecGetter) Find(kind string, id string) (*storage.AsyncApiSpec, error) {
	ret := _m.Called(kind, id)

	var r0 *storage.AsyncApiSpec
	if rf, ok := ret.Get(0).(func(string, string) *storage.AsyncApiSpec); ok {
		r0 = rf(kind, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.AsyncApiSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(kind, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
