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

// FiltersSubnet One or more filters.
type FiltersSubnet struct {
	// The number of available IPs.
	AvailableIpsCounts *[]int64 `json:"AvailableIpsCounts,omitempty"`
	// The IP ranges in the Subnets, in CIDR notation (for example, 10.0.0.0/16).
	IpRanges *[]string `json:"IpRanges,omitempty"`
	// The IDs of the Nets in which the Subnets are.
	NetIds *[]string `json:"NetIds,omitempty"`
	// The states of the Subnets (`pending` \\| `available`).
	States *[]string `json:"States,omitempty"`
	// The IDs of the Subnets.
	SubnetIds *[]string `json:"SubnetIds,omitempty"`
	// The names of the Subregions in which the Subnets are located.
	SubregionNames *[]string `json:"SubregionNames,omitempty"`
}

// GetAvailableIpsCounts returns the AvailableIpsCounts field value if set, zero value otherwise.
func (o *FiltersSubnet) GetAvailableIpsCounts() []int64 {
	if o == nil || o.AvailableIpsCounts == nil {
		var ret []int64
		return ret
	}
	return *o.AvailableIpsCounts
}

// GetAvailableIpsCountsOk returns a tuple with the AvailableIpsCounts field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSubnet) GetAvailableIpsCountsOk() ([]int64, bool) {
	if o == nil || o.AvailableIpsCounts == nil {
		var ret []int64
		return ret, false
	}
	return *o.AvailableIpsCounts, true
}

// HasAvailableIpsCounts returns a boolean if a field has been set.
func (o *FiltersSubnet) HasAvailableIpsCounts() bool {
	if o != nil && o.AvailableIpsCounts != nil {
		return true
	}

	return false
}

// SetAvailableIpsCounts gets a reference to the given []int64 and assigns it to the AvailableIpsCounts field.
func (o *FiltersSubnet) SetAvailableIpsCounts(v []int64) {
	o.AvailableIpsCounts = &v
}

// GetIpRanges returns the IpRanges field value if set, zero value otherwise.
func (o *FiltersSubnet) GetIpRanges() []string {
	if o == nil || o.IpRanges == nil {
		var ret []string
		return ret
	}
	return *o.IpRanges
}

// GetIpRangesOk returns a tuple with the IpRanges field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSubnet) GetIpRangesOk() ([]string, bool) {
	if o == nil || o.IpRanges == nil {
		var ret []string
		return ret, false
	}
	return *o.IpRanges, true
}

// HasIpRanges returns a boolean if a field has been set.
func (o *FiltersSubnet) HasIpRanges() bool {
	if o != nil && o.IpRanges != nil {
		return true
	}

	return false
}

// SetIpRanges gets a reference to the given []string and assigns it to the IpRanges field.
func (o *FiltersSubnet) SetIpRanges(v []string) {
	o.IpRanges = &v
}

// GetNetIds returns the NetIds field value if set, zero value otherwise.
func (o *FiltersSubnet) GetNetIds() []string {
	if o == nil || o.NetIds == nil {
		var ret []string
		return ret
	}
	return *o.NetIds
}

// GetNetIdsOk returns a tuple with the NetIds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSubnet) GetNetIdsOk() ([]string, bool) {
	if o == nil || o.NetIds == nil {
		var ret []string
		return ret, false
	}
	return *o.NetIds, true
}

// HasNetIds returns a boolean if a field has been set.
func (o *FiltersSubnet) HasNetIds() bool {
	if o != nil && o.NetIds != nil {
		return true
	}

	return false
}

// SetNetIds gets a reference to the given []string and assigns it to the NetIds field.
func (o *FiltersSubnet) SetNetIds(v []string) {
	o.NetIds = &v
}

// GetStates returns the States field value if set, zero value otherwise.
func (o *FiltersSubnet) GetStates() []string {
	if o == nil || o.States == nil {
		var ret []string
		return ret
	}
	return *o.States
}

// GetStatesOk returns a tuple with the States field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSubnet) GetStatesOk() ([]string, bool) {
	if o == nil || o.States == nil {
		var ret []string
		return ret, false
	}
	return *o.States, true
}

// HasStates returns a boolean if a field has been set.
func (o *FiltersSubnet) HasStates() bool {
	if o != nil && o.States != nil {
		return true
	}

	return false
}

// SetStates gets a reference to the given []string and assigns it to the States field.
func (o *FiltersSubnet) SetStates(v []string) {
	o.States = &v
}

// GetSubnetIds returns the SubnetIds field value if set, zero value otherwise.
func (o *FiltersSubnet) GetSubnetIds() []string {
	if o == nil || o.SubnetIds == nil {
		var ret []string
		return ret
	}
	return *o.SubnetIds
}

// GetSubnetIdsOk returns a tuple with the SubnetIds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSubnet) GetSubnetIdsOk() ([]string, bool) {
	if o == nil || o.SubnetIds == nil {
		var ret []string
		return ret, false
	}
	return *o.SubnetIds, true
}

// HasSubnetIds returns a boolean if a field has been set.
func (o *FiltersSubnet) HasSubnetIds() bool {
	if o != nil && o.SubnetIds != nil {
		return true
	}

	return false
}

// SetSubnetIds gets a reference to the given []string and assigns it to the SubnetIds field.
func (o *FiltersSubnet) SetSubnetIds(v []string) {
	o.SubnetIds = &v
}

// GetSubregionNames returns the SubregionNames field value if set, zero value otherwise.
func (o *FiltersSubnet) GetSubregionNames() []string {
	if o == nil || o.SubregionNames == nil {
		var ret []string
		return ret
	}
	return *o.SubregionNames
}

// GetSubregionNamesOk returns a tuple with the SubregionNames field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSubnet) GetSubregionNamesOk() ([]string, bool) {
	if o == nil || o.SubregionNames == nil {
		var ret []string
		return ret, false
	}
	return *o.SubregionNames, true
}

// HasSubregionNames returns a boolean if a field has been set.
func (o *FiltersSubnet) HasSubregionNames() bool {
	if o != nil && o.SubregionNames != nil {
		return true
	}

	return false
}

// SetSubregionNames gets a reference to the given []string and assigns it to the SubregionNames field.
func (o *FiltersSubnet) SetSubregionNames(v []string) {
	o.SubregionNames = &v
}

type NullableFiltersSubnet struct {
	Value        FiltersSubnet
	ExplicitNull bool
}

func (v NullableFiltersSubnet) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableFiltersSubnet) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
