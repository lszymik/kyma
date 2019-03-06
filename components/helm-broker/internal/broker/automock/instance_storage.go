package automock

import "github.com/kyma-project/kyma/components/helm-broker/internal"
import "github.com/stretchr/testify/mock"

// InstanceStorage is an autogenerated mock type for the InstanceStorage type
type InstanceStorage struct {
	mock.Mock
}

// Get provides a mock function with given fields: id
func (_m *InstanceStorage) Get(id internal.InstanceID) (*internal.Instance, error) {
	ret := _m.Called(id)

	var r0 *internal.Instance
	if rf, ok := ret.Get(0).(func(internal.InstanceID) *internal.Instance); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internal.Instance)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(internal.InstanceID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *InstanceStorage) GetAll() ([]*internal.Instance, error) {
	ret := _m.Called()

	var r0 []*internal.Instance
	if rf, ok := ret.Get(0).(func() []*internal.Instance); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*internal.Instance)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: i
func (_m *InstanceStorage) Insert(i *internal.Instance) error {
	ret := _m.Called(i)

	var r0 error
	if rf, ok := ret.Get(0).(func(*internal.Instance) error); ok {
		r0 = rf(i)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Remove provides a mock function with given fields: id
func (_m *InstanceStorage) Remove(id internal.InstanceID) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(internal.InstanceID) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
