// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// IngressControllerCaptureHTTPHeadersApplyConfiguration represents a declarative configuration of the IngressControllerCaptureHTTPHeaders type for use
// with apply.
type IngressControllerCaptureHTTPHeadersApplyConfiguration struct {
	Request  []IngressControllerCaptureHTTPHeaderApplyConfiguration `json:"request,omitempty"`
	Response []IngressControllerCaptureHTTPHeaderApplyConfiguration `json:"response,omitempty"`
}

// IngressControllerCaptureHTTPHeadersApplyConfiguration constructs a declarative configuration of the IngressControllerCaptureHTTPHeaders type for use with
// apply.
func IngressControllerCaptureHTTPHeaders() *IngressControllerCaptureHTTPHeadersApplyConfiguration {
	return &IngressControllerCaptureHTTPHeadersApplyConfiguration{}
}

// WithRequest adds the given value to the Request field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Request field.
func (b *IngressControllerCaptureHTTPHeadersApplyConfiguration) WithRequest(values ...*IngressControllerCaptureHTTPHeaderApplyConfiguration) *IngressControllerCaptureHTTPHeadersApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithRequest")
		}
		b.Request = append(b.Request, *values[i])
	}
	return b
}

// WithResponse adds the given value to the Response field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Response field.
func (b *IngressControllerCaptureHTTPHeadersApplyConfiguration) WithResponse(values ...*IngressControllerCaptureHTTPHeaderApplyConfiguration) *IngressControllerCaptureHTTPHeadersApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithResponse")
		}
		b.Response = append(b.Response, *values[i])
	}
	return b
}
