// Code generated by mockery v1.0.0
package storage

import mock "github.com/stretchr/testify/mock"

// mockStoreGetter is an autogenerated mock type for the storeGetter type
type mockStoreGetter struct {
	mock.Mock
}

// ApiSpec provides a mock function with given fields: id
func (_m *mockStoreGetter) ApiSpec(id string) (*ApiSpec, bool, error) {
	ret := _m.Called(id)

	var r0 *ApiSpec
	if rf, ok := ret.Get(0).(func(string) *ApiSpec); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ApiSpec)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(id)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// AsyncApiSpec provides a mock function with given fields: id
func (_m *mockStoreGetter) AsyncApiSpec(id string) (*AsyncApiSpec, bool, error) {
	ret := _m.Called(id)

	var r0 *AsyncApiSpec
	if rf, ok := ret.Get(0).(func(string) *AsyncApiSpec); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*AsyncApiSpec)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(id)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Content provides a mock function with given fields: id
func (_m *mockStoreGetter) Content(id string) (*Content, bool, error) {
	ret := _m.Called(id)

	var r0 *Content
	if rf, ok := ret.Get(0).(func(string) *Content); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Content)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(id)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NotificationChannel provides a mock function with given fields: stop
func (_m *mockStoreGetter) NotificationChannel(stop <-chan struct{}) <-chan notification {
	ret := _m.Called(stop)

	var r0 <-chan notification
	if rf, ok := ret.Get(0).(func(<-chan struct{}) <-chan notification); ok {
		r0 = rf(stop)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan notification)
		}
	}

	return r0
}
