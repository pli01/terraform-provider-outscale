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

// LinkPublicIp Information about the EIP association.
type LinkPublicIp struct {
	// (Required in a Net) The ID representing the association of the EIP with the VM or the NIC.
	LinkPublicIpId *string `json:"LinkPublicIpId,omitempty"`
	// The name of the public DNS.
	PublicDnsName *string `json:"PublicDnsName,omitempty"`
	// The External IP address (EIP) associated with the NIC.
	PublicIp *string `json:"PublicIp,omitempty"`
	// The account ID of the owner of the EIP.
	PublicIpAccountId *string `json:"PublicIpAccountId,omitempty"`
	// The allocation ID of the EIP.
	PublicIpId *string `json:"PublicIpId,omitempty"`
}

// GetLinkPublicIpId returns the LinkPublicIpId field value if set, zero value otherwise.
func (o *LinkPublicIp) GetLinkPublicIpId() string {
	if o == nil || o.LinkPublicIpId == nil {
		var ret string
		return ret
	}
	return *o.LinkPublicIpId
}

// GetLinkPublicIpIdOk returns a tuple with the LinkPublicIpId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LinkPublicIp) GetLinkPublicIpIdOk() (string, bool) {
	if o == nil || o.LinkPublicIpId == nil {
		var ret string
		return ret, false
	}
	return *o.LinkPublicIpId, true
}

// HasLinkPublicIpId returns a boolean if a field has been set.
func (o *LinkPublicIp) HasLinkPublicIpId() bool {
	if o != nil && o.LinkPublicIpId != nil {
		return true
	}

	return false
}

// SetLinkPublicIpId gets a reference to the given string and assigns it to the LinkPublicIpId field.
func (o *LinkPublicIp) SetLinkPublicIpId(v string) {
	o.LinkPublicIpId = &v
}

// GetPublicDnsName returns the PublicDnsName field value if set, zero value otherwise.
func (o *LinkPublicIp) GetPublicDnsName() string {
	if o == nil || o.PublicDnsName == nil {
		var ret string
		return ret
	}
	return *o.PublicDnsName
}

// GetPublicDnsNameOk returns a tuple with the PublicDnsName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LinkPublicIp) GetPublicDnsNameOk() (string, bool) {
	if o == nil || o.PublicDnsName == nil {
		var ret string
		return ret, false
	}
	return *o.PublicDnsName, true
}

// HasPublicDnsName returns a boolean if a field has been set.
func (o *LinkPublicIp) HasPublicDnsName() bool {
	if o != nil && o.PublicDnsName != nil {
		return true
	}

	return false
}

// SetPublicDnsName gets a reference to the given string and assigns it to the PublicDnsName field.
func (o *LinkPublicIp) SetPublicDnsName(v string) {
	o.PublicDnsName = &v
}

// GetPublicIp returns the PublicIp field value if set, zero value otherwise.
func (o *LinkPublicIp) GetPublicIp() string {
	if o == nil || o.PublicIp == nil {
		var ret string
		return ret
	}
	return *o.PublicIp
}

// GetPublicIpOk returns a tuple with the PublicIp field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LinkPublicIp) GetPublicIpOk() (string, bool) {
	if o == nil || o.PublicIp == nil {
		var ret string
		return ret, false
	}
	return *o.PublicIp, true
}

// HasPublicIp returns a boolean if a field has been set.
func (o *LinkPublicIp) HasPublicIp() bool {
	if o != nil && o.PublicIp != nil {
		return true
	}

	return false
}

// SetPublicIp gets a reference to the given string and assigns it to the PublicIp field.
func (o *LinkPublicIp) SetPublicIp(v string) {
	o.PublicIp = &v
}

// GetPublicIpAccountId returns the PublicIpAccountId field value if set, zero value otherwise.
func (o *LinkPublicIp) GetPublicIpAccountId() string {
	if o == nil || o.PublicIpAccountId == nil {
		var ret string
		return ret
	}
	return *o.PublicIpAccountId
}

// GetPublicIpAccountIdOk returns a tuple with the PublicIpAccountId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LinkPublicIp) GetPublicIpAccountIdOk() (string, bool) {
	if o == nil || o.PublicIpAccountId == nil {
		var ret string
		return ret, false
	}
	return *o.PublicIpAccountId, true
}

// HasPublicIpAccountId returns a boolean if a field has been set.
func (o *LinkPublicIp) HasPublicIpAccountId() bool {
	if o != nil && o.PublicIpAccountId != nil {
		return true
	}

	return false
}

// SetPublicIpAccountId gets a reference to the given string and assigns it to the PublicIpAccountId field.
func (o *LinkPublicIp) SetPublicIpAccountId(v string) {
	o.PublicIpAccountId = &v
}

// GetPublicIpId returns the PublicIpId field value if set, zero value otherwise.
func (o *LinkPublicIp) GetPublicIpId() string {
	if o == nil || o.PublicIpId == nil {
		var ret string
		return ret
	}
	return *o.PublicIpId
}

// GetPublicIpIdOk returns a tuple with the PublicIpId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LinkPublicIp) GetPublicIpIdOk() (string, bool) {
	if o == nil || o.PublicIpId == nil {
		var ret string
		return ret, false
	}
	return *o.PublicIpId, true
}

// HasPublicIpId returns a boolean if a field has been set.
func (o *LinkPublicIp) HasPublicIpId() bool {
	if o != nil && o.PublicIpId != nil {
		return true
	}

	return false
}

// SetPublicIpId gets a reference to the given string and assigns it to the PublicIpId field.
func (o *LinkPublicIp) SetPublicIpId(v string) {
	o.PublicIpId = &v
}

type NullableLinkPublicIp struct {
	Value        LinkPublicIp
	ExplicitNull bool
}

func (v NullableLinkPublicIp) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableLinkPublicIp) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
