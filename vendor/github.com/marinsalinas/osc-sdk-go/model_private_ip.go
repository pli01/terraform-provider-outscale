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

// PrivateIp Information about the private IP.
type PrivateIp struct {
	// If `true`, the IP address is the primary private IP address of the NIC.
	IsPrimary    *bool         `json:"IsPrimary,omitempty"`
	LinkPublicIp *LinkPublicIp `json:"LinkPublicIp,omitempty"`
	// The name of the private DNS.
	PrivateDnsName *string `json:"PrivateDnsName,omitempty"`
	// The private IP address of the NIC.
	PrivateIp *string `json:"PrivateIp,omitempty"`
}

// GetIsPrimary returns the IsPrimary field value if set, zero value otherwise.
func (o *PrivateIp) GetIsPrimary() bool {
	if o == nil || o.IsPrimary == nil {
		var ret bool
		return ret
	}
	return *o.IsPrimary
}

// GetIsPrimaryOk returns a tuple with the IsPrimary field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *PrivateIp) GetIsPrimaryOk() (bool, bool) {
	if o == nil || o.IsPrimary == nil {
		var ret bool
		return ret, false
	}
	return *o.IsPrimary, true
}

// HasIsPrimary returns a boolean if a field has been set.
func (o *PrivateIp) HasIsPrimary() bool {
	if o != nil && o.IsPrimary != nil {
		return true
	}

	return false
}

// SetIsPrimary gets a reference to the given bool and assigns it to the IsPrimary field.
func (o *PrivateIp) SetIsPrimary(v bool) {
	o.IsPrimary = &v
}

// GetLinkPublicIp returns the LinkPublicIp field value if set, zero value otherwise.
func (o *PrivateIp) GetLinkPublicIp() LinkPublicIp {
	if o == nil || o.LinkPublicIp == nil {
		var ret LinkPublicIp
		return ret
	}
	return *o.LinkPublicIp
}

// GetLinkPublicIpOk returns a tuple with the LinkPublicIp field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *PrivateIp) GetLinkPublicIpOk() (LinkPublicIp, bool) {
	if o == nil || o.LinkPublicIp == nil {
		var ret LinkPublicIp
		return ret, false
	}
	return *o.LinkPublicIp, true
}

// HasLinkPublicIp returns a boolean if a field has been set.
func (o *PrivateIp) HasLinkPublicIp() bool {
	if o != nil && o.LinkPublicIp != nil {
		return true
	}

	return false
}

// SetLinkPublicIp gets a reference to the given LinkPublicIp and assigns it to the LinkPublicIp field.
func (o *PrivateIp) SetLinkPublicIp(v LinkPublicIp) {
	o.LinkPublicIp = &v
}

// GetPrivateDnsName returns the PrivateDnsName field value if set, zero value otherwise.
func (o *PrivateIp) GetPrivateDnsName() string {
	if o == nil || o.PrivateDnsName == nil {
		var ret string
		return ret
	}
	return *o.PrivateDnsName
}

// GetPrivateDnsNameOk returns a tuple with the PrivateDnsName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *PrivateIp) GetPrivateDnsNameOk() (string, bool) {
	if o == nil || o.PrivateDnsName == nil {
		var ret string
		return ret, false
	}
	return *o.PrivateDnsName, true
}

// HasPrivateDnsName returns a boolean if a field has been set.
func (o *PrivateIp) HasPrivateDnsName() bool {
	if o != nil && o.PrivateDnsName != nil {
		return true
	}

	return false
}

// SetPrivateDnsName gets a reference to the given string and assigns it to the PrivateDnsName field.
func (o *PrivateIp) SetPrivateDnsName(v string) {
	o.PrivateDnsName = &v
}

// GetPrivateIp returns the PrivateIp field value if set, zero value otherwise.
func (o *PrivateIp) GetPrivateIp() string {
	if o == nil || o.PrivateIp == nil {
		var ret string
		return ret
	}
	return *o.PrivateIp
}

// GetPrivateIpOk returns a tuple with the PrivateIp field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *PrivateIp) GetPrivateIpOk() (string, bool) {
	if o == nil || o.PrivateIp == nil {
		var ret string
		return ret, false
	}
	return *o.PrivateIp, true
}

// HasPrivateIp returns a boolean if a field has been set.
func (o *PrivateIp) HasPrivateIp() bool {
	if o != nil && o.PrivateIp != nil {
		return true
	}

	return false
}

// SetPrivateIp gets a reference to the given string and assigns it to the PrivateIp field.
func (o *PrivateIp) SetPrivateIp(v string) {
	o.PrivateIp = &v
}

type NullablePrivateIp struct {
	Value        PrivateIp
	ExplicitNull bool
}

func (v NullablePrivateIp) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullablePrivateIp) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
