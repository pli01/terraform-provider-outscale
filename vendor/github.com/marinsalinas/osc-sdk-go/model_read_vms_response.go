/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 0.15
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package oscgo

import (
	"bytes"
	"encoding/json"
)

// ReadVmsResponse struct for ReadVmsResponse
type ReadVmsResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
	// Information about one or more VMs.
	Vms *[]Vm `json:"Vms,omitempty"`
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadVmsResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ReadVmsResponse) GetResponseContextOk() (ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret, false
	}
	return *o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadVmsResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadVmsResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

// GetVms returns the Vms field value if set, zero value otherwise.
func (o *ReadVmsResponse) GetVms() []Vm {
	if o == nil || o.Vms == nil {
		var ret []Vm
		return ret
	}
	return *o.Vms
}

// GetVmsOk returns a tuple with the Vms field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ReadVmsResponse) GetVmsOk() ([]Vm, bool) {
	if o == nil || o.Vms == nil {
		var ret []Vm
		return ret, false
	}
	return *o.Vms, true
}

// HasVms returns a boolean if a field has been set.
func (o *ReadVmsResponse) HasVms() bool {
	if o != nil && o.Vms != nil {
		return true
	}

	return false
}

// SetVms gets a reference to the given []Vm and assigns it to the Vms field.
func (o *ReadVmsResponse) SetVms(v []Vm) {
	o.Vms = &v
}

type NullableReadVmsResponse struct {
	Value        ReadVmsResponse
	ExplicitNull bool
}

func (v NullableReadVmsResponse) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableReadVmsResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
