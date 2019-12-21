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

// ReadPublicIpRangesResponse struct for ReadPublicIpRangesResponse
type ReadPublicIpRangesResponse struct {
	// The list of public IPv4 addresses used in the Region, in CIDR notation.
	PublicIps       *[]string        `json:"PublicIps,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// GetPublicIps returns the PublicIps field value if set, zero value otherwise.
func (o *ReadPublicIpRangesResponse) GetPublicIps() []string {
	if o == nil || o.PublicIps == nil {
		var ret []string
		return ret
	}
	return *o.PublicIps
}

// GetPublicIpsOk returns a tuple with the PublicIps field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ReadPublicIpRangesResponse) GetPublicIpsOk() ([]string, bool) {
	if o == nil || o.PublicIps == nil {
		var ret []string
		return ret, false
	}
	return *o.PublicIps, true
}

// HasPublicIps returns a boolean if a field has been set.
func (o *ReadPublicIpRangesResponse) HasPublicIps() bool {
	if o != nil && o.PublicIps != nil {
		return true
	}

	return false
}

// SetPublicIps gets a reference to the given []string and assigns it to the PublicIps field.
func (o *ReadPublicIpRangesResponse) SetPublicIps(v []string) {
	o.PublicIps = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadPublicIpRangesResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ReadPublicIpRangesResponse) GetResponseContextOk() (ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret, false
	}
	return *o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadPublicIpRangesResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadPublicIpRangesResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

type NullableReadPublicIpRangesResponse struct {
	Value        ReadPublicIpRangesResponse
	ExplicitNull bool
}

func (v NullableReadPublicIpRangesResponse) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableReadPublicIpRangesResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
