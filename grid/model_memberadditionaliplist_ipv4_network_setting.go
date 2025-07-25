/*
Infoblox GRID API

OpenAPI specification for Infoblox NIOS WAPI GRID objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package grid

import (
	"encoding/json"
)

// checks if the MemberadditionaliplistIpv4NetworkSetting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MemberadditionaliplistIpv4NetworkSetting{}

// MemberadditionaliplistIpv4NetworkSetting struct for MemberadditionaliplistIpv4NetworkSetting
type MemberadditionaliplistIpv4NetworkSetting struct {
	// The IPv4 Address of the Grid Member.
	Address *string `json:"address,omitempty"`
	// The default gateway for the Grid Member.
	Gateway *string `json:"gateway,omitempty"`
	// The subnet mask for the Grid Member.
	SubnetMask *string `json:"subnet_mask,omitempty"`
	// The identifier for the VLAN. Valid values are from 1 to 4096.
	VlanId *int64 `json:"vlan_id,omitempty"`
	// Determines if the current address is the primary VLAN address or not.
	Primary *bool `json:"primary,omitempty"`
	// The DSCP (Differentiated Services Code Point) value determines relative priorities for the type of services on your network. The appliance implements QoS (Quality of Service) rules based on this configuration. Valid values are from 0 to 63.
	Dscp *int64 `json:"dscp,omitempty"`
	// LAN netmask only for GCP HA.
	LanSubnetMask *string `json:"lan_subnet_mask,omitempty"`
	// LAN gateway only for GCP HA.
	LanGateway *string `json:"lan_gateway,omitempty"`
	// Use flag for: dscp
	UseDscp              *bool `json:"use_dscp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MemberadditionaliplistIpv4NetworkSetting MemberadditionaliplistIpv4NetworkSetting

// NewMemberadditionaliplistIpv4NetworkSetting instantiates a new MemberadditionaliplistIpv4NetworkSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemberadditionaliplistIpv4NetworkSetting() *MemberadditionaliplistIpv4NetworkSetting {
	this := MemberadditionaliplistIpv4NetworkSetting{}
	return &this
}

// NewMemberadditionaliplistIpv4NetworkSettingWithDefaults instantiates a new MemberadditionaliplistIpv4NetworkSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemberadditionaliplistIpv4NetworkSettingWithDefaults() *MemberadditionaliplistIpv4NetworkSetting {
	this := MemberadditionaliplistIpv4NetworkSetting{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *MemberadditionaliplistIpv4NetworkSetting) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberadditionaliplistIpv4NetworkSetting) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *MemberadditionaliplistIpv4NetworkSetting) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *MemberadditionaliplistIpv4NetworkSetting) SetAddress(v string) {
	o.Address = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *MemberadditionaliplistIpv4NetworkSetting) GetGateway() string {
	if o == nil || IsNil(o.Gateway) {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberadditionaliplistIpv4NetworkSetting) GetGatewayOk() (*string, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *MemberadditionaliplistIpv4NetworkSetting) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *MemberadditionaliplistIpv4NetworkSetting) SetGateway(v string) {
	o.Gateway = &v
}

// GetSubnetMask returns the SubnetMask field value if set, zero value otherwise.
func (o *MemberadditionaliplistIpv4NetworkSetting) GetSubnetMask() string {
	if o == nil || IsNil(o.SubnetMask) {
		var ret string
		return ret
	}
	return *o.SubnetMask
}

// GetSubnetMaskOk returns a tuple with the SubnetMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberadditionaliplistIpv4NetworkSetting) GetSubnetMaskOk() (*string, bool) {
	if o == nil || IsNil(o.SubnetMask) {
		return nil, false
	}
	return o.SubnetMask, true
}

// HasSubnetMask returns a boolean if a field has been set.
func (o *MemberadditionaliplistIpv4NetworkSetting) HasSubnetMask() bool {
	if o != nil && !IsNil(o.SubnetMask) {
		return true
	}

	return false
}

// SetSubnetMask gets a reference to the given string and assigns it to the SubnetMask field.
func (o *MemberadditionaliplistIpv4NetworkSetting) SetSubnetMask(v string) {
	o.SubnetMask = &v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *MemberadditionaliplistIpv4NetworkSetting) GetVlanId() int64 {
	if o == nil || IsNil(o.VlanId) {
		var ret int64
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberadditionaliplistIpv4NetworkSetting) GetVlanIdOk() (*int64, bool) {
	if o == nil || IsNil(o.VlanId) {
		return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *MemberadditionaliplistIpv4NetworkSetting) HasVlanId() bool {
	if o != nil && !IsNil(o.VlanId) {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given int64 and assigns it to the VlanId field.
func (o *MemberadditionaliplistIpv4NetworkSetting) SetVlanId(v int64) {
	o.VlanId = &v
}

// GetPrimary returns the Primary field value if set, zero value otherwise.
func (o *MemberadditionaliplistIpv4NetworkSetting) GetPrimary() bool {
	if o == nil || IsNil(o.Primary) {
		var ret bool
		return ret
	}
	return *o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberadditionaliplistIpv4NetworkSetting) GetPrimaryOk() (*bool, bool) {
	if o == nil || IsNil(o.Primary) {
		return nil, false
	}
	return o.Primary, true
}

// HasPrimary returns a boolean if a field has been set.
func (o *MemberadditionaliplistIpv4NetworkSetting) HasPrimary() bool {
	if o != nil && !IsNil(o.Primary) {
		return true
	}

	return false
}

// SetPrimary gets a reference to the given bool and assigns it to the Primary field.
func (o *MemberadditionaliplistIpv4NetworkSetting) SetPrimary(v bool) {
	o.Primary = &v
}

// GetDscp returns the Dscp field value if set, zero value otherwise.
func (o *MemberadditionaliplistIpv4NetworkSetting) GetDscp() int64 {
	if o == nil || IsNil(o.Dscp) {
		var ret int64
		return ret
	}
	return *o.Dscp
}

// GetDscpOk returns a tuple with the Dscp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberadditionaliplistIpv4NetworkSetting) GetDscpOk() (*int64, bool) {
	if o == nil || IsNil(o.Dscp) {
		return nil, false
	}
	return o.Dscp, true
}

// HasDscp returns a boolean if a field has been set.
func (o *MemberadditionaliplistIpv4NetworkSetting) HasDscp() bool {
	if o != nil && !IsNil(o.Dscp) {
		return true
	}

	return false
}

// SetDscp gets a reference to the given int64 and assigns it to the Dscp field.
func (o *MemberadditionaliplistIpv4NetworkSetting) SetDscp(v int64) {
	o.Dscp = &v
}

// GetLanSubnetMask returns the LanSubnetMask field value if set, zero value otherwise.
func (o *MemberadditionaliplistIpv4NetworkSetting) GetLanSubnetMask() string {
	if o == nil || IsNil(o.LanSubnetMask) {
		var ret string
		return ret
	}
	return *o.LanSubnetMask
}

// GetLanSubnetMaskOk returns a tuple with the LanSubnetMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberadditionaliplistIpv4NetworkSetting) GetLanSubnetMaskOk() (*string, bool) {
	if o == nil || IsNil(o.LanSubnetMask) {
		return nil, false
	}
	return o.LanSubnetMask, true
}

// HasLanSubnetMask returns a boolean if a field has been set.
func (o *MemberadditionaliplistIpv4NetworkSetting) HasLanSubnetMask() bool {
	if o != nil && !IsNil(o.LanSubnetMask) {
		return true
	}

	return false
}

// SetLanSubnetMask gets a reference to the given string and assigns it to the LanSubnetMask field.
func (o *MemberadditionaliplistIpv4NetworkSetting) SetLanSubnetMask(v string) {
	o.LanSubnetMask = &v
}

// GetLanGateway returns the LanGateway field value if set, zero value otherwise.
func (o *MemberadditionaliplistIpv4NetworkSetting) GetLanGateway() string {
	if o == nil || IsNil(o.LanGateway) {
		var ret string
		return ret
	}
	return *o.LanGateway
}

// GetLanGatewayOk returns a tuple with the LanGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberadditionaliplistIpv4NetworkSetting) GetLanGatewayOk() (*string, bool) {
	if o == nil || IsNil(o.LanGateway) {
		return nil, false
	}
	return o.LanGateway, true
}

// HasLanGateway returns a boolean if a field has been set.
func (o *MemberadditionaliplistIpv4NetworkSetting) HasLanGateway() bool {
	if o != nil && !IsNil(o.LanGateway) {
		return true
	}

	return false
}

// SetLanGateway gets a reference to the given string and assigns it to the LanGateway field.
func (o *MemberadditionaliplistIpv4NetworkSetting) SetLanGateway(v string) {
	o.LanGateway = &v
}

// GetUseDscp returns the UseDscp field value if set, zero value otherwise.
func (o *MemberadditionaliplistIpv4NetworkSetting) GetUseDscp() bool {
	if o == nil || IsNil(o.UseDscp) {
		var ret bool
		return ret
	}
	return *o.UseDscp
}

// GetUseDscpOk returns a tuple with the UseDscp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberadditionaliplistIpv4NetworkSetting) GetUseDscpOk() (*bool, bool) {
	if o == nil || IsNil(o.UseDscp) {
		return nil, false
	}
	return o.UseDscp, true
}

// HasUseDscp returns a boolean if a field has been set.
func (o *MemberadditionaliplistIpv4NetworkSetting) HasUseDscp() bool {
	if o != nil && !IsNil(o.UseDscp) {
		return true
	}

	return false
}

// SetUseDscp gets a reference to the given bool and assigns it to the UseDscp field.
func (o *MemberadditionaliplistIpv4NetworkSetting) SetUseDscp(v bool) {
	o.UseDscp = &v
}

func (o MemberadditionaliplistIpv4NetworkSetting) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MemberadditionaliplistIpv4NetworkSetting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	if !IsNil(o.SubnetMask) {
		toSerialize["subnet_mask"] = o.SubnetMask
	}
	if !IsNil(o.VlanId) {
		toSerialize["vlan_id"] = o.VlanId
	}
	if !IsNil(o.Primary) {
		toSerialize["primary"] = o.Primary
	}
	if !IsNil(o.Dscp) {
		toSerialize["dscp"] = o.Dscp
	}
	if !IsNil(o.LanSubnetMask) {
		toSerialize["lan_subnet_mask"] = o.LanSubnetMask
	}
	if !IsNil(o.LanGateway) {
		toSerialize["lan_gateway"] = o.LanGateway
	}
	if !IsNil(o.UseDscp) {
		toSerialize["use_dscp"] = o.UseDscp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MemberadditionaliplistIpv4NetworkSetting) UnmarshalJSON(data []byte) (err error) {
	varMemberadditionaliplistIpv4NetworkSetting := _MemberadditionaliplistIpv4NetworkSetting{}

	err = json.Unmarshal(data, &varMemberadditionaliplistIpv4NetworkSetting)

	if err != nil {
		return err
	}

	*o = MemberadditionaliplistIpv4NetworkSetting(varMemberadditionaliplistIpv4NetworkSetting)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "gateway")
		delete(additionalProperties, "subnet_mask")
		delete(additionalProperties, "vlan_id")
		delete(additionalProperties, "primary")
		delete(additionalProperties, "dscp")
		delete(additionalProperties, "lan_subnet_mask")
		delete(additionalProperties, "lan_gateway")
		delete(additionalProperties, "use_dscp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMemberadditionaliplistIpv4NetworkSetting struct {
	value *MemberadditionaliplistIpv4NetworkSetting
	isSet bool
}

func (v NullableMemberadditionaliplistIpv4NetworkSetting) Get() *MemberadditionaliplistIpv4NetworkSetting {
	return v.value
}

func (v *NullableMemberadditionaliplistIpv4NetworkSetting) Set(val *MemberadditionaliplistIpv4NetworkSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableMemberadditionaliplistIpv4NetworkSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableMemberadditionaliplistIpv4NetworkSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemberadditionaliplistIpv4NetworkSetting(val *MemberadditionaliplistIpv4NetworkSetting) *NullableMemberadditionaliplistIpv4NetworkSetting {
	return &NullableMemberadditionaliplistIpv4NetworkSetting{value: val, isSet: true}
}

func (v NullableMemberadditionaliplistIpv4NetworkSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemberadditionaliplistIpv4NetworkSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
