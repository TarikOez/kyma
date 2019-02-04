// Code generated by mockery v1.0.0
package automock

import mock "github.com/stretchr/testify/mock"
import shared "github.com/kyma-project/kyma/components/ui-api-layer/internal/domain/shared"

// ServiceCatalogRetriever is an autogenerated mock type for the ServiceCatalogRetriever type
type ServiceCatalogRetriever struct {
	mock.Mock
}

// ServiceBinding provides a mock function with given fields:
func (_m *ServiceCatalogRetriever) ServiceBinding() shared.ServiceBindingFinderLister {
	ret := _m.Called()

	var r0 shared.ServiceBindingFinderLister
	if rf, ok := ret.Get(0).(func() shared.ServiceBindingFinderLister); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(shared.ServiceBindingFinderLister)
		}
	}

	return r0
}
