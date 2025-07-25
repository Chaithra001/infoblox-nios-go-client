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

// checks if the MemberThreatprotectionNatRules type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MemberThreatprotectionNatRules{}

// MemberThreatprotectionNatRules struct for MemberThreatprotectionNatRules
type MemberThreatprotectionNatRules struct {
	// The rule type for the threat protection NAT mapping rule.
	RuleType *string `json:"rule_type,omitempty"`
	// The IP address for the threat protection NAT mapping rule.
	Address *string `json:"address,omitempty"`
	// The network address for the threat protection NAT mapping rule.
	Network *string `json:"network,omitempty"`
	// The network CIDR for the threat protection NAT mapping rule.
	Cidr *int64 `json:"cidr,omitempty"`
	// The start address for the range of the threat protection NAT mapping rule.
	StartAddress *string `json:"start_address,omitempty"`
	// The end address for the range of the threat protection NAT mapping rule.
	EndAddress *string `json:"end_address,omitempty"`
	// The NAT port configuration for the threat protection NAT mapping rule.
	NatPorts             []MemberthreatprotectionnatrulesNatPorts `json:"nat_ports,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MemberThreatprotectionNatRules MemberThreatprotectionNatRules

// NewMemberThreatprotectionNatRules instantiates a new MemberThreatprotectionNatRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemberThreatprotectionNatRules() *MemberThreatprotectionNatRules {
	this := MemberThreatprotectionNatRules{}
	return &this
}

// NewMemberThreatprotectionNatRulesWithDefaults instantiates a new MemberThreatprotectionNatRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemberThreatprotectionNatRulesWithDefaults() *MemberThreatprotectionNatRules {
	this := MemberThreatprotectionNatRules{}
	return &this
}

// GetRuleType returns the RuleType field value if set, zero value otherwise.
func (o *MemberThreatprotectionNatRules) GetRuleType() string {
	if o == nil || IsNil(o.RuleType) {
		var ret string
		return ret
	}
	return *o.RuleType
}

// GetRuleTypeOk returns a tuple with the RuleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberThreatprotectionNatRules) GetRuleTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RuleType) {
		return nil, false
	}
	return o.RuleType, true
}

// HasRuleType returns a boolean if a field has been set.
func (o *MemberThreatprotectionNatRules) HasRuleType() bool {
	if o != nil && !IsNil(o.RuleType) {
		return true
	}

	return false
}

// SetRuleType gets a reference to the given string and assigns it to the RuleType field.
func (o *MemberThreatprotectionNatRules) SetRuleType(v string) {
	o.RuleType = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *MemberThreatprotectionNatRules) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberThreatprotectionNatRules) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *MemberThreatprotectionNatRules) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *MemberThreatprotectionNatRules) SetAddress(v string) {
	o.Address = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *MemberThreatprotectionNatRules) GetNetwork() string {
	if o == nil || IsNil(o.Network) {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberThreatprotectionNatRules) GetNetworkOk() (*string, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *MemberThreatprotectionNatRules) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *MemberThreatprotectionNatRules) SetNetwork(v string) {
	o.Network = &v
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *MemberThreatprotectionNatRules) GetCidr() int64 {
	if o == nil || IsNil(o.Cidr) {
		var ret int64
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberThreatprotectionNatRules) GetCidrOk() (*int64, bool) {
	if o == nil || IsNil(o.Cidr) {
		return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *MemberThreatprotectionNatRules) HasCidr() bool {
	if o != nil && !IsNil(o.Cidr) {
		return true
	}

	return false
}

// SetCidr gets a reference to the given int64 and assigns it to the Cidr field.
func (o *MemberThreatprotectionNatRules) SetCidr(v int64) {
	o.Cidr = &v
}

// GetStartAddress returns the StartAddress field value if set, zero value otherwise.
func (o *MemberThreatprotectionNatRules) GetStartAddress() string {
	if o == nil || IsNil(o.StartAddress) {
		var ret string
		return ret
	}
	return *o.StartAddress
}

// GetStartAddressOk returns a tuple with the StartAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberThreatprotectionNatRules) GetStartAddressOk() (*string, bool) {
	if o == nil || IsNil(o.StartAddress) {
		return nil, false
	}
	return o.StartAddress, true
}

// HasStartAddress returns a boolean if a field has been set.
func (o *MemberThreatprotectionNatRules) HasStartAddress() bool {
	if o != nil && !IsNil(o.StartAddress) {
		return true
	}

	return false
}

// SetStartAddress gets a reference to the given string and assigns it to the StartAddress field.
func (o *MemberThreatprotectionNatRules) SetStartAddress(v string) {
	o.StartAddress = &v
}

// GetEndAddress returns the EndAddress field value if set, zero value otherwise.
func (o *MemberThreatprotectionNatRules) GetEndAddress() string {
	if o == nil || IsNil(o.EndAddress) {
		var ret string
		return ret
	}
	return *o.EndAddress
}

// GetEndAddressOk returns a tuple with the EndAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberThreatprotectionNatRules) GetEndAddressOk() (*string, bool) {
	if o == nil || IsNil(o.EndAddress) {
		return nil, false
	}
	return o.EndAddress, true
}

// HasEndAddress returns a boolean if a field has been set.
func (o *MemberThreatprotectionNatRules) HasEndAddress() bool {
	if o != nil && !IsNil(o.EndAddress) {
		return true
	}

	return false
}

// SetEndAddress gets a reference to the given string and assigns it to the EndAddress field.
func (o *MemberThreatprotectionNatRules) SetEndAddress(v string) {
	o.EndAddress = &v
}

// GetNatPorts returns the NatPorts field value if set, zero value otherwise.
func (o *MemberThreatprotectionNatRules) GetNatPorts() []MemberthreatprotectionnatrulesNatPorts {
	if o == nil || IsNil(o.NatPorts) {
		var ret []MemberthreatprotectionnatrulesNatPorts
		return ret
	}
	return o.NatPorts
}

// GetNatPortsOk returns a tuple with the NatPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberThreatprotectionNatRules) GetNatPortsOk() ([]MemberthreatprotectionnatrulesNatPorts, bool) {
	if o == nil || IsNil(o.NatPorts) {
		return nil, false
	}
	return o.NatPorts, true
}

// HasNatPorts returns a boolean if a field has been set.
func (o *MemberThreatprotectionNatRules) HasNatPorts() bool {
	if o != nil && !IsNil(o.NatPorts) {
		return true
	}

	return false
}

// SetNatPorts gets a reference to the given []MemberthreatprotectionnatrulesNatPorts and assigns it to the NatPorts field.
func (o *MemberThreatprotectionNatRules) SetNatPorts(v []MemberthreatprotectionnatrulesNatPorts) {
	o.NatPorts = v
}

func (o MemberThreatprotectionNatRules) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MemberThreatprotectionNatRules) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RuleType) {
		toSerialize["rule_type"] = o.RuleType
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.Cidr) {
		toSerialize["cidr"] = o.Cidr
	}
	if !IsNil(o.StartAddress) {
		toSerialize["start_address"] = o.StartAddress
	}
	if !IsNil(o.EndAddress) {
		toSerialize["end_address"] = o.EndAddress
	}
	if !IsNil(o.NatPorts) {
		toSerialize["nat_ports"] = o.NatPorts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MemberThreatprotectionNatRules) UnmarshalJSON(data []byte) (err error) {
	varMemberThreatprotectionNatRules := _MemberThreatprotectionNatRules{}

	err = json.Unmarshal(data, &varMemberThreatprotectionNatRules)

	if err != nil {
		return err
	}

	*o = MemberThreatprotectionNatRules(varMemberThreatprotectionNatRules)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rule_type")
		delete(additionalProperties, "address")
		delete(additionalProperties, "network")
		delete(additionalProperties, "cidr")
		delete(additionalProperties, "start_address")
		delete(additionalProperties, "end_address")
		delete(additionalProperties, "nat_ports")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMemberThreatprotectionNatRules struct {
	value *MemberThreatprotectionNatRules
	isSet bool
}

func (v NullableMemberThreatprotectionNatRules) Get() *MemberThreatprotectionNatRules {
	return v.value
}

func (v *NullableMemberThreatprotectionNatRules) Set(val *MemberThreatprotectionNatRules) {
	v.value = val
	v.isSet = true
}

func (v NullableMemberThreatprotectionNatRules) IsSet() bool {
	return v.isSet
}

func (v *NullableMemberThreatprotectionNatRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemberThreatprotectionNatRules(val *MemberThreatprotectionNatRules) *NullableMemberThreatprotectionNatRules {
	return &NullableMemberThreatprotectionNatRules{value: val, isSet: true}
}

func (v NullableMemberThreatprotectionNatRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemberThreatprotectionNatRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
