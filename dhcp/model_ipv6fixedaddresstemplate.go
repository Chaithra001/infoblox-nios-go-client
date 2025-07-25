/*
Infoblox DHCP API

OpenAPI specification for Infoblox NIOS WAPI DHCP objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dhcp

import (
	"encoding/json"
)

// checks if the Ipv6fixedaddresstemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Ipv6fixedaddresstemplate{}

// Ipv6fixedaddresstemplate struct for Ipv6fixedaddresstemplate
type Ipv6fixedaddresstemplate struct {
	// The reference to the object.
	Ref *string `json:"_ref,omitempty"`
	// A descriptive comment of an IPv6 fixed address template object.
	Comment *string `json:"comment,omitempty"`
	// Domain name of the IPv6 fixed address template object.
	DomainName *string `json:"domain_name,omitempty"`
	// The IPv6 addresses of DNS recursive name servers to which the DHCP client can send name resolution requests. The DHCP server includes this information in the DNS Recursive Name Server option in Advertise, Rebind, Information-Request, and Reply messages.
	DomainNameServers []string `json:"domain_name_servers,omitempty"`
	// Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}.
	ExtAttrs *map[string]ExtAttrs `json:"extattrs,omitempty"`
	// This field contains the logic filters to be applied to this IPv6 fixed address. This list corresponds to the match rules that are written to the DHCPv6 configuration file.
	LogicFilterRules []Ipv6fixedaddresstemplateLogicFilterRules `json:"logic_filter_rules,omitempty"`
	// Name of an IPv6 fixed address template object.
	Name *string `json:"name,omitempty"`
	// The number of IPv6 addresses for this fixed address.
	NumberOfAddresses *int64 `json:"number_of_addresses,omitempty"`
	// The start address offset for this IPv6 fixed address.
	Offset *int64 `json:"offset,omitempty"`
	// An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object.
	Options []Ipv6fixedaddresstemplateOptions `json:"options,omitempty"`
	// The preferred lifetime value for this DHCP IPv6 fixed address template object.
	PreferredLifetime *int64 `json:"preferred_lifetime,omitempty"`
	// Use flag for: domain_name
	UseDomainName *bool `json:"use_domain_name,omitempty"`
	// Use flag for: domain_name_servers
	UseDomainNameServers *bool `json:"use_domain_name_servers,omitempty"`
	// Use flag for: logic_filter_rules
	UseLogicFilterRules *bool `json:"use_logic_filter_rules,omitempty"`
	// Use flag for: options
	UseOptions *bool `json:"use_options,omitempty"`
	// Use flag for: preferred_lifetime
	UsePreferredLifetime *bool `json:"use_preferred_lifetime,omitempty"`
	// Use flag for: valid_lifetime
	UseValidLifetime *bool `json:"use_valid_lifetime,omitempty"`
	// The valid lifetime value for this DHCP IPv6 fixed address template object.
	ValidLifetime *int64 `json:"valid_lifetime,omitempty"`
}

// NewIpv6fixedaddresstemplate instantiates a new Ipv6fixedaddresstemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpv6fixedaddresstemplate() *Ipv6fixedaddresstemplate {
	this := Ipv6fixedaddresstemplate{}
	return &this
}

// NewIpv6fixedaddresstemplateWithDefaults instantiates a new Ipv6fixedaddresstemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpv6fixedaddresstemplateWithDefaults() *Ipv6fixedaddresstemplate {
	this := Ipv6fixedaddresstemplate{}
	return &this
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *Ipv6fixedaddresstemplate) GetRef() string {
	if o == nil || IsNil(o.Ref) {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ipv6fixedaddresstemplate) GetRefOk() (*string, bool) {
	if o == nil || IsNil(o.Ref) {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *Ipv6fixedaddresstemplate) HasRef() bool {
	if o != nil && !IsNil(o.Ref) {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *Ipv6fixedaddresstemplate) SetRef(v string) {
	o.Ref = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *Ipv6fixedaddresstemplate) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ipv6fixedaddresstemplate) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *Ipv6fixedaddresstemplate) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *Ipv6fixedaddresstemplate) SetComment(v string) {
	o.Comment = &v
}

// GetDomainName returns the DomainName field value if set, zero value otherwise.
func (o *Ipv6fixedaddresstemplate) GetDomainName() string {
	if o == nil || IsNil(o.DomainName) {
		var ret string
		return ret
	}
	return *o.DomainName
}

// GetDomainNameOk returns a tuple with the DomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ipv6fixedaddresstemplate) GetDomainNameOk() (*string, bool) {
	if o == nil || IsNil(o.DomainName) {
		return nil, false
	}
	return o.DomainName, true
}

// HasDomainName returns a boolean if a field has been set.
func (o *Ipv6fixedaddresstemplate) HasDomainName() bool {
	if o != nil && !IsNil(o.DomainName) {
		return true
	}

	return false
}

// SetDomainName gets a reference to the given string and assigns it to the DomainName field.
func (o *Ipv6fixedaddresstemplate) SetDomainName(v string) {
	o.DomainName = &v
}

// GetDomainNameServers returns the DomainNameServers field value if set, zero value otherwise.
func (o *Ipv6fixedaddresstemplate) GetDomainNameServers() []string {
	if o == nil || IsNil(o.DomainNameServers) {
		var ret []string
		return ret
	}
	return o.DomainNameServers
}

// GetDomainNameServersOk returns a tuple with the DomainNameServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ipv6fixedaddresstemplate) GetDomainNameServersOk() ([]string, bool) {
	if o == nil || IsNil(o.DomainNameServers) {
		return nil, false
	}
	return o.DomainNameServers, true
}

// HasDomainNameServers returns a boolean if a field has been set.
func (o *Ipv6fixedaddresstemplate) HasDomainNameServers() bool {
	if o != nil && !IsNil(o.DomainNameServers) {
		return true
	}

	return false
}

// SetDomainNameServers gets a reference to the given []string and assigns it to the DomainNameServers field.
func (o *Ipv6fixedaddresstemplate) SetDomainNameServers(v []string) {
	o.DomainNameServers = v
}

// GetExtAttrs returns the ExtAttrs field value if set, zero value otherwise.
func (o *Ipv6fixedaddresstemplate) GetExtAttrs() map[string]ExtAttrs {
	if o == nil || IsNil(o.ExtAttrs) {
		var ret map[string]ExtAttrs
		return ret
	}
	return *o.ExtAttrs
}

// GetExtAttrsOk returns a tuple with the ExtAttrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ipv6fixedaddresstemplate) GetExtAttrsOk() (*map[string]ExtAttrs, bool) {
	if o == nil || IsNil(o.ExtAttrs) {
		return nil, false
	}
	return o.ExtAttrs, true
}

// HasExtAttrs returns a boolean if a field has been set.
func (o *Ipv6fixedaddresstemplate) HasExtAttrs() bool {
	if o != nil && !IsNil(o.ExtAttrs) {
		return true
	}

	return false
}

// SetExtAttrs gets a reference to the given map[string]ExtAttrs and assigns it to the ExtAttrs field.
func (o *Ipv6fixedaddresstemplate) SetExtAttrs(v map[string]ExtAttrs) {
	o.ExtAttrs = &v
}

// GetLogicFilterRules returns the LogicFilterRules field value if set, zero value otherwise.
func (o *Ipv6fixedaddresstemplate) GetLogicFilterRules() []Ipv6fixedaddresstemplateLogicFilterRules {
	if o == nil || IsNil(o.LogicFilterRules) {
		var ret []Ipv6fixedaddresstemplateLogicFilterRules
		return ret
	}
	return o.LogicFilterRules
}

// GetLogicFilterRulesOk returns a tuple with the LogicFilterRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ipv6fixedaddresstemplate) GetLogicFilterRulesOk() ([]Ipv6fixedaddresstemplateLogicFilterRules, bool) {
	if o == nil || IsNil(o.LogicFilterRules) {
		return nil, false
	}
	return o.LogicFilterRules, true
}

// HasLogicFilterRules returns a boolean if a field has been set.
func (o *Ipv6fixedaddresstemplate) HasLogicFilterRules() bool {
	if o != nil && !IsNil(o.LogicFilterRules) {
		return true
	}

	return false
}

// SetLogicFilterRules gets a reference to the given []Ipv6fixedaddresstemplateLogicFilterRules and assigns it to the LogicFilterRules field.
func (o *Ipv6fixedaddresstemplate) SetLogicFilterRules(v []Ipv6fixedaddresstemplateLogicFilterRules) {
	o.LogicFilterRules = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Ipv6fixedaddresstemplate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ipv6fixedaddresstemplate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Ipv6fixedaddresstemplate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Ipv6fixedaddresstemplate) SetName(v string) {
	o.Name = &v
}

// GetNumberOfAddresses returns the NumberOfAddresses field value if set, zero value otherwise.
func (o *Ipv6fixedaddresstemplate) GetNumberOfAddresses() int64 {
	if o == nil || IsNil(o.NumberOfAddresses) {
		var ret int64
		return ret
	}
	return *o.NumberOfAddresses
}

// GetNumberOfAddressesOk returns a tuple with the NumberOfAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ipv6fixedaddresstemplate) GetNumberOfAddressesOk() (*int64, bool) {
	if o == nil || IsNil(o.NumberOfAddresses) {
		return nil, false
	}
	return o.NumberOfAddresses, true
}

// HasNumberOfAddresses returns a boolean if a field has been set.
func (o *Ipv6fixedaddresstemplate) HasNumberOfAddresses() bool {
	if o != nil && !IsNil(o.NumberOfAddresses) {
		return true
	}

	return false
}

// SetNumberOfAddresses gets a reference to the given int64 and assigns it to the NumberOfAddresses field.
func (o *Ipv6fixedaddresstemplate) SetNumberOfAddresses(v int64) {
	o.NumberOfAddresses = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *Ipv6fixedaddresstemplate) GetOffset() int64 {
	if o == nil || IsNil(o.Offset) {
		var ret int64
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ipv6fixedaddresstemplate) GetOffsetOk() (*int64, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *Ipv6fixedaddresstemplate) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int64 and assigns it to the Offset field.
func (o *Ipv6fixedaddresstemplate) SetOffset(v int64) {
	o.Offset = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *Ipv6fixedaddresstemplate) GetOptions() []Ipv6fixedaddresstemplateOptions {
	if o == nil || IsNil(o.Options) {
		var ret []Ipv6fixedaddresstemplateOptions
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ipv6fixedaddresstemplate) GetOptionsOk() ([]Ipv6fixedaddresstemplateOptions, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *Ipv6fixedaddresstemplate) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []Ipv6fixedaddresstemplateOptions and assigns it to the Options field.
func (o *Ipv6fixedaddresstemplate) SetOptions(v []Ipv6fixedaddresstemplateOptions) {
	o.Options = v
}

// GetPreferredLifetime returns the PreferredLifetime field value if set, zero value otherwise.
func (o *Ipv6fixedaddresstemplate) GetPreferredLifetime() int64 {
	if o == nil || IsNil(o.PreferredLifetime) {
		var ret int64
		return ret
	}
	return *o.PreferredLifetime
}

// GetPreferredLifetimeOk returns a tuple with the PreferredLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ipv6fixedaddresstemplate) GetPreferredLifetimeOk() (*int64, bool) {
	if o == nil || IsNil(o.PreferredLifetime) {
		return nil, false
	}
	return o.PreferredLifetime, true
}

// HasPreferredLifetime returns a boolean if a field has been set.
func (o *Ipv6fixedaddresstemplate) HasPreferredLifetime() bool {
	if o != nil && !IsNil(o.PreferredLifetime) {
		return true
	}

	return false
}

// SetPreferredLifetime gets a reference to the given int64 and assigns it to the PreferredLifetime field.
func (o *Ipv6fixedaddresstemplate) SetPreferredLifetime(v int64) {
	o.PreferredLifetime = &v
}

// GetUseDomainName returns the UseDomainName field value if set, zero value otherwise.
func (o *Ipv6fixedaddresstemplate) GetUseDomainName() bool {
	if o == nil || IsNil(o.UseDomainName) {
		var ret bool
		return ret
	}
	return *o.UseDomainName
}

// GetUseDomainNameOk returns a tuple with the UseDomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ipv6fixedaddresstemplate) GetUseDomainNameOk() (*bool, bool) {
	if o == nil || IsNil(o.UseDomainName) {
		return nil, false
	}
	return o.UseDomainName, true
}

// HasUseDomainName returns a boolean if a field has been set.
func (o *Ipv6fixedaddresstemplate) HasUseDomainName() bool {
	if o != nil && !IsNil(o.UseDomainName) {
		return true
	}

	return false
}

// SetUseDomainName gets a reference to the given bool and assigns it to the UseDomainName field.
func (o *Ipv6fixedaddresstemplate) SetUseDomainName(v bool) {
	o.UseDomainName = &v
}

// GetUseDomainNameServers returns the UseDomainNameServers field value if set, zero value otherwise.
func (o *Ipv6fixedaddresstemplate) GetUseDomainNameServers() bool {
	if o == nil || IsNil(o.UseDomainNameServers) {
		var ret bool
		return ret
	}
	return *o.UseDomainNameServers
}

// GetUseDomainNameServersOk returns a tuple with the UseDomainNameServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ipv6fixedaddresstemplate) GetUseDomainNameServersOk() (*bool, bool) {
	if o == nil || IsNil(o.UseDomainNameServers) {
		return nil, false
	}
	return o.UseDomainNameServers, true
}

// HasUseDomainNameServers returns a boolean if a field has been set.
func (o *Ipv6fixedaddresstemplate) HasUseDomainNameServers() bool {
	if o != nil && !IsNil(o.UseDomainNameServers) {
		return true
	}

	return false
}

// SetUseDomainNameServers gets a reference to the given bool and assigns it to the UseDomainNameServers field.
func (o *Ipv6fixedaddresstemplate) SetUseDomainNameServers(v bool) {
	o.UseDomainNameServers = &v
}

// GetUseLogicFilterRules returns the UseLogicFilterRules field value if set, zero value otherwise.
func (o *Ipv6fixedaddresstemplate) GetUseLogicFilterRules() bool {
	if o == nil || IsNil(o.UseLogicFilterRules) {
		var ret bool
		return ret
	}
	return *o.UseLogicFilterRules
}

// GetUseLogicFilterRulesOk returns a tuple with the UseLogicFilterRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ipv6fixedaddresstemplate) GetUseLogicFilterRulesOk() (*bool, bool) {
	if o == nil || IsNil(o.UseLogicFilterRules) {
		return nil, false
	}
	return o.UseLogicFilterRules, true
}

// HasUseLogicFilterRules returns a boolean if a field has been set.
func (o *Ipv6fixedaddresstemplate) HasUseLogicFilterRules() bool {
	if o != nil && !IsNil(o.UseLogicFilterRules) {
		return true
	}

	return false
}

// SetUseLogicFilterRules gets a reference to the given bool and assigns it to the UseLogicFilterRules field.
func (o *Ipv6fixedaddresstemplate) SetUseLogicFilterRules(v bool) {
	o.UseLogicFilterRules = &v
}

// GetUseOptions returns the UseOptions field value if set, zero value otherwise.
func (o *Ipv6fixedaddresstemplate) GetUseOptions() bool {
	if o == nil || IsNil(o.UseOptions) {
		var ret bool
		return ret
	}
	return *o.UseOptions
}

// GetUseOptionsOk returns a tuple with the UseOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ipv6fixedaddresstemplate) GetUseOptionsOk() (*bool, bool) {
	if o == nil || IsNil(o.UseOptions) {
		return nil, false
	}
	return o.UseOptions, true
}

// HasUseOptions returns a boolean if a field has been set.
func (o *Ipv6fixedaddresstemplate) HasUseOptions() bool {
	if o != nil && !IsNil(o.UseOptions) {
		return true
	}

	return false
}

// SetUseOptions gets a reference to the given bool and assigns it to the UseOptions field.
func (o *Ipv6fixedaddresstemplate) SetUseOptions(v bool) {
	o.UseOptions = &v
}

// GetUsePreferredLifetime returns the UsePreferredLifetime field value if set, zero value otherwise.
func (o *Ipv6fixedaddresstemplate) GetUsePreferredLifetime() bool {
	if o == nil || IsNil(o.UsePreferredLifetime) {
		var ret bool
		return ret
	}
	return *o.UsePreferredLifetime
}

// GetUsePreferredLifetimeOk returns a tuple with the UsePreferredLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ipv6fixedaddresstemplate) GetUsePreferredLifetimeOk() (*bool, bool) {
	if o == nil || IsNil(o.UsePreferredLifetime) {
		return nil, false
	}
	return o.UsePreferredLifetime, true
}

// HasUsePreferredLifetime returns a boolean if a field has been set.
func (o *Ipv6fixedaddresstemplate) HasUsePreferredLifetime() bool {
	if o != nil && !IsNil(o.UsePreferredLifetime) {
		return true
	}

	return false
}

// SetUsePreferredLifetime gets a reference to the given bool and assigns it to the UsePreferredLifetime field.
func (o *Ipv6fixedaddresstemplate) SetUsePreferredLifetime(v bool) {
	o.UsePreferredLifetime = &v
}

// GetUseValidLifetime returns the UseValidLifetime field value if set, zero value otherwise.
func (o *Ipv6fixedaddresstemplate) GetUseValidLifetime() bool {
	if o == nil || IsNil(o.UseValidLifetime) {
		var ret bool
		return ret
	}
	return *o.UseValidLifetime
}

// GetUseValidLifetimeOk returns a tuple with the UseValidLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ipv6fixedaddresstemplate) GetUseValidLifetimeOk() (*bool, bool) {
	if o == nil || IsNil(o.UseValidLifetime) {
		return nil, false
	}
	return o.UseValidLifetime, true
}

// HasUseValidLifetime returns a boolean if a field has been set.
func (o *Ipv6fixedaddresstemplate) HasUseValidLifetime() bool {
	if o != nil && !IsNil(o.UseValidLifetime) {
		return true
	}

	return false
}

// SetUseValidLifetime gets a reference to the given bool and assigns it to the UseValidLifetime field.
func (o *Ipv6fixedaddresstemplate) SetUseValidLifetime(v bool) {
	o.UseValidLifetime = &v
}

// GetValidLifetime returns the ValidLifetime field value if set, zero value otherwise.
func (o *Ipv6fixedaddresstemplate) GetValidLifetime() int64 {
	if o == nil || IsNil(o.ValidLifetime) {
		var ret int64
		return ret
	}
	return *o.ValidLifetime
}

// GetValidLifetimeOk returns a tuple with the ValidLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ipv6fixedaddresstemplate) GetValidLifetimeOk() (*int64, bool) {
	if o == nil || IsNil(o.ValidLifetime) {
		return nil, false
	}
	return o.ValidLifetime, true
}

// HasValidLifetime returns a boolean if a field has been set.
func (o *Ipv6fixedaddresstemplate) HasValidLifetime() bool {
	if o != nil && !IsNil(o.ValidLifetime) {
		return true
	}

	return false
}

// SetValidLifetime gets a reference to the given int64 and assigns it to the ValidLifetime field.
func (o *Ipv6fixedaddresstemplate) SetValidLifetime(v int64) {
	o.ValidLifetime = &v
}

func (o Ipv6fixedaddresstemplate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Ipv6fixedaddresstemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ref) {
		toSerialize["_ref"] = o.Ref
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.DomainName) {
		toSerialize["domain_name"] = o.DomainName
	}
	if !IsNil(o.DomainNameServers) {
		toSerialize["domain_name_servers"] = o.DomainNameServers
	}
	if !IsNil(o.ExtAttrs) {
		toSerialize["extattrs"] = o.ExtAttrs
	}
	if !IsNil(o.LogicFilterRules) {
		toSerialize["logic_filter_rules"] = o.LogicFilterRules
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NumberOfAddresses) {
		toSerialize["number_of_addresses"] = o.NumberOfAddresses
	}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if !IsNil(o.PreferredLifetime) {
		toSerialize["preferred_lifetime"] = o.PreferredLifetime
	}
	if !IsNil(o.UseDomainName) {
		toSerialize["use_domain_name"] = o.UseDomainName
	}
	if !IsNil(o.UseDomainNameServers) {
		toSerialize["use_domain_name_servers"] = o.UseDomainNameServers
	}
	if !IsNil(o.UseLogicFilterRules) {
		toSerialize["use_logic_filter_rules"] = o.UseLogicFilterRules
	}
	if !IsNil(o.UseOptions) {
		toSerialize["use_options"] = o.UseOptions
	}
	if !IsNil(o.UsePreferredLifetime) {
		toSerialize["use_preferred_lifetime"] = o.UsePreferredLifetime
	}
	if !IsNil(o.UseValidLifetime) {
		toSerialize["use_valid_lifetime"] = o.UseValidLifetime
	}
	if !IsNil(o.ValidLifetime) {
		toSerialize["valid_lifetime"] = o.ValidLifetime
	}
	return toSerialize, nil
}

type NullableIpv6fixedaddresstemplate struct {
	value *Ipv6fixedaddresstemplate
	isSet bool
}

func (v NullableIpv6fixedaddresstemplate) Get() *Ipv6fixedaddresstemplate {
	return v.value
}

func (v *NullableIpv6fixedaddresstemplate) Set(val *Ipv6fixedaddresstemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableIpv6fixedaddresstemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableIpv6fixedaddresstemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpv6fixedaddresstemplate(val *Ipv6fixedaddresstemplate) *NullableIpv6fixedaddresstemplate {
	return &NullableIpv6fixedaddresstemplate{value: val, isSet: true}
}

func (v NullableIpv6fixedaddresstemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpv6fixedaddresstemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
