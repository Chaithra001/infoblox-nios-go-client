/*
Infoblox DNS API

OpenAPI specification for Infoblox NIOS WAPI DNS objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"encoding/json"
)

// checks if the RecordMx type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecordMx{}

// RecordMx struct for RecordMx
type RecordMx struct {
	// The reference to the object.
	Ref                *string                     `json:"_ref,omitempty"`
	AwsRte53RecordInfo *RecordMxAwsRte53RecordInfo `json:"aws_rte53_record_info,omitempty"`
	CloudInfo          *RecordMxCloudInfo          `json:"cloud_info,omitempty"`
	// Comment for the record; maximum 256 characters.
	Comment *string `json:"comment,omitempty"`
	// The time of the record creation in Epoch seconds format.
	CreationTime *int64 `json:"creation_time,omitempty"`
	// The record creator. Note that changing creator from or to 'SYSTEM' value is not allowed.
	Creator *string `json:"creator,omitempty"`
	// The GSS-TSIG principal that owns this record.
	DdnsPrincipal *string `json:"ddns_principal,omitempty"`
	// Determines if the DDNS updates for this record are allowed or not.
	DdnsProtected *bool `json:"ddns_protected,omitempty"`
	// Determines if the record is disabled or not. False means that the record is enabled.
	Disable *bool `json:"disable,omitempty"`
	// The Mail exchanger name in punycode format.
	DnsMailExchanger *string `json:"dns_mail_exchanger,omitempty"`
	// The name for a MX record in punycode format.
	DnsName *string `json:"dns_name,omitempty"`
	// Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}.
	ExtAttrs *map[string]ExtAttrs `json:"extattrs,omitempty"`
	// Determines if the reclamation is allowed for the record or not.
	ForbidReclamation *bool `json:"forbid_reclamation,omitempty"`
	// The time of the last DNS query in Epoch seconds format.
	LastQueried *int64 `json:"last_queried,omitempty"`
	// Mail exchanger name in FQDN format. This value can be in unicode format.
	MailExchanger *string `json:"mail_exchanger,omitempty"`
	// Name for the MX record in FQDN format. This value can be in unicode format.
	Name *string `json:"name,omitempty"`
	// Preference value, 0 to 65535 (inclusive) in 32-bit unsigned integer format.
	Preference *int64 `json:"preference,omitempty"`
	// Determines if the record is reclaimable or not.
	Reclaimable *bool `json:"reclaimable,omitempty"`
	// The name of the shared record group in which the record resides. This field exists only on db_objects if this record is a shared record.
	SharedRecordGroup *string `json:"shared_record_group,omitempty"`
	// The Time To Live (TTL) value for record. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached.
	Ttl *int64 `json:"ttl,omitempty"`
	// Use flag for: ttl
	UseTtl *bool `json:"use_ttl,omitempty"`
	// The name of the DNS view in which the record resides. Example: \"external\".
	View *string `json:"view,omitempty"`
	// The name of the zone in which the record resides. Example: \"zone.com\". If a view is not specified when searching by zone, the default view is used.
	Zone *string `json:"zone,omitempty"`
}

// NewRecordMx instantiates a new RecordMx object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecordMx() *RecordMx {
	this := RecordMx{}
	return &this
}

// NewRecordMxWithDefaults instantiates a new RecordMx object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecordMxWithDefaults() *RecordMx {
	this := RecordMx{}
	return &this
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *RecordMx) GetRef() string {
	if o == nil || IsNil(o.Ref) {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordMx) GetRefOk() (*string, bool) {
	if o == nil || IsNil(o.Ref) {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *RecordMx) HasRef() bool {
	if o != nil && !IsNil(o.Ref) {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *RecordMx) SetRef(v string) {
	o.Ref = &v
}

// GetAwsRte53RecordInfo returns the AwsRte53RecordInfo field value if set, zero value otherwise.
func (o *RecordMx) GetAwsRte53RecordInfo() RecordMxAwsRte53RecordInfo {
	if o == nil || IsNil(o.AwsRte53RecordInfo) {
		var ret RecordMxAwsRte53RecordInfo
		return ret
	}
	return *o.AwsRte53RecordInfo
}

// GetAwsRte53RecordInfoOk returns a tuple with the AwsRte53RecordInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordMx) GetAwsRte53RecordInfoOk() (*RecordMxAwsRte53RecordInfo, bool) {
	if o == nil || IsNil(o.AwsRte53RecordInfo) {
		return nil, false
	}
	return o.AwsRte53RecordInfo, true
}

// HasAwsRte53RecordInfo returns a boolean if a field has been set.
func (o *RecordMx) HasAwsRte53RecordInfo() bool {
	if o != nil && !IsNil(o.AwsRte53RecordInfo) {
		return true
	}

	return false
}

// SetAwsRte53RecordInfo gets a reference to the given RecordMxAwsRte53RecordInfo and assigns it to the AwsRte53RecordInfo field.
func (o *RecordMx) SetAwsRte53RecordInfo(v RecordMxAwsRte53RecordInfo) {
	o.AwsRte53RecordInfo = &v
}

// GetCloudInfo returns the CloudInfo field value if set, zero value otherwise.
func (o *RecordMx) GetCloudInfo() RecordMxCloudInfo {
	if o == nil || IsNil(o.CloudInfo) {
		var ret RecordMxCloudInfo
		return ret
	}
	return *o.CloudInfo
}

// GetCloudInfoOk returns a tuple with the CloudInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordMx) GetCloudInfoOk() (*RecordMxCloudInfo, bool) {
	if o == nil || IsNil(o.CloudInfo) {
		return nil, false
	}
	return o.CloudInfo, true
}

// HasCloudInfo returns a boolean if a field has been set.
func (o *RecordMx) HasCloudInfo() bool {
	if o != nil && !IsNil(o.CloudInfo) {
		return true
	}

	return false
}

// SetCloudInfo gets a reference to the given RecordMxCloudInfo and assigns it to the CloudInfo field.
func (o *RecordMx) SetCloudInfo(v RecordMxCloudInfo) {
	o.CloudInfo = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *RecordMx) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordMx) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *RecordMx) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *RecordMx) SetComment(v string) {
	o.Comment = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *RecordMx) GetCreationTime() int64 {
	if o == nil || IsNil(o.CreationTime) {
		var ret int64
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordMx) GetCreationTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CreationTime) {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *RecordMx) HasCreationTime() bool {
	if o != nil && !IsNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given int64 and assigns it to the CreationTime field.
func (o *RecordMx) SetCreationTime(v int64) {
	o.CreationTime = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *RecordMx) GetCreator() string {
	if o == nil || IsNil(o.Creator) {
		var ret string
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordMx) GetCreatorOk() (*string, bool) {
	if o == nil || IsNil(o.Creator) {
		return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *RecordMx) HasCreator() bool {
	if o != nil && !IsNil(o.Creator) {
		return true
	}

	return false
}

// SetCreator gets a reference to the given string and assigns it to the Creator field.
func (o *RecordMx) SetCreator(v string) {
	o.Creator = &v
}

// GetDdnsPrincipal returns the DdnsPrincipal field value if set, zero value otherwise.
func (o *RecordMx) GetDdnsPrincipal() string {
	if o == nil || IsNil(o.DdnsPrincipal) {
		var ret string
		return ret
	}
	return *o.DdnsPrincipal
}

// GetDdnsPrincipalOk returns a tuple with the DdnsPrincipal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordMx) GetDdnsPrincipalOk() (*string, bool) {
	if o == nil || IsNil(o.DdnsPrincipal) {
		return nil, false
	}
	return o.DdnsPrincipal, true
}

// HasDdnsPrincipal returns a boolean if a field has been set.
func (o *RecordMx) HasDdnsPrincipal() bool {
	if o != nil && !IsNil(o.DdnsPrincipal) {
		return true
	}

	return false
}

// SetDdnsPrincipal gets a reference to the given string and assigns it to the DdnsPrincipal field.
func (o *RecordMx) SetDdnsPrincipal(v string) {
	o.DdnsPrincipal = &v
}

// GetDdnsProtected returns the DdnsProtected field value if set, zero value otherwise.
func (o *RecordMx) GetDdnsProtected() bool {
	if o == nil || IsNil(o.DdnsProtected) {
		var ret bool
		return ret
	}
	return *o.DdnsProtected
}

// GetDdnsProtectedOk returns a tuple with the DdnsProtected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordMx) GetDdnsProtectedOk() (*bool, bool) {
	if o == nil || IsNil(o.DdnsProtected) {
		return nil, false
	}
	return o.DdnsProtected, true
}

// HasDdnsProtected returns a boolean if a field has been set.
func (o *RecordMx) HasDdnsProtected() bool {
	if o != nil && !IsNil(o.DdnsProtected) {
		return true
	}

	return false
}

// SetDdnsProtected gets a reference to the given bool and assigns it to the DdnsProtected field.
func (o *RecordMx) SetDdnsProtected(v bool) {
	o.DdnsProtected = &v
}

// GetDisable returns the Disable field value if set, zero value otherwise.
func (o *RecordMx) GetDisable() bool {
	if o == nil || IsNil(o.Disable) {
		var ret bool
		return ret
	}
	return *o.Disable
}

// GetDisableOk returns a tuple with the Disable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordMx) GetDisableOk() (*bool, bool) {
	if o == nil || IsNil(o.Disable) {
		return nil, false
	}
	return o.Disable, true
}

// HasDisable returns a boolean if a field has been set.
func (o *RecordMx) HasDisable() bool {
	if o != nil && !IsNil(o.Disable) {
		return true
	}

	return false
}

// SetDisable gets a reference to the given bool and assigns it to the Disable field.
func (o *RecordMx) SetDisable(v bool) {
	o.Disable = &v
}

// GetDnsMailExchanger returns the DnsMailExchanger field value if set, zero value otherwise.
func (o *RecordMx) GetDnsMailExchanger() string {
	if o == nil || IsNil(o.DnsMailExchanger) {
		var ret string
		return ret
	}
	return *o.DnsMailExchanger
}

// GetDnsMailExchangerOk returns a tuple with the DnsMailExchanger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordMx) GetDnsMailExchangerOk() (*string, bool) {
	if o == nil || IsNil(o.DnsMailExchanger) {
		return nil, false
	}
	return o.DnsMailExchanger, true
}

// HasDnsMailExchanger returns a boolean if a field has been set.
func (o *RecordMx) HasDnsMailExchanger() bool {
	if o != nil && !IsNil(o.DnsMailExchanger) {
		return true
	}

	return false
}

// SetDnsMailExchanger gets a reference to the given string and assigns it to the DnsMailExchanger field.
func (o *RecordMx) SetDnsMailExchanger(v string) {
	o.DnsMailExchanger = &v
}

// GetDnsName returns the DnsName field value if set, zero value otherwise.
func (o *RecordMx) GetDnsName() string {
	if o == nil || IsNil(o.DnsName) {
		var ret string
		return ret
	}
	return *o.DnsName
}

// GetDnsNameOk returns a tuple with the DnsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordMx) GetDnsNameOk() (*string, bool) {
	if o == nil || IsNil(o.DnsName) {
		return nil, false
	}
	return o.DnsName, true
}

// HasDnsName returns a boolean if a field has been set.
func (o *RecordMx) HasDnsName() bool {
	if o != nil && !IsNil(o.DnsName) {
		return true
	}

	return false
}

// SetDnsName gets a reference to the given string and assigns it to the DnsName field.
func (o *RecordMx) SetDnsName(v string) {
	o.DnsName = &v
}

// GetExtAttrs returns the ExtAttrs field value if set, zero value otherwise.
func (o *RecordMx) GetExtAttrs() map[string]ExtAttrs {
	if o == nil || IsNil(o.ExtAttrs) {
		var ret map[string]ExtAttrs
		return ret
	}
	return *o.ExtAttrs
}

// GetExtAttrsOk returns a tuple with the ExtAttrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordMx) GetExtAttrsOk() (*map[string]ExtAttrs, bool) {
	if o == nil || IsNil(o.ExtAttrs) {
		return nil, false
	}
	return o.ExtAttrs, true
}

// HasExtAttrs returns a boolean if a field has been set.
func (o *RecordMx) HasExtAttrs() bool {
	if o != nil && !IsNil(o.ExtAttrs) {
		return true
	}

	return false
}

// SetExtAttrs gets a reference to the given map[string]ExtAttrs and assigns it to the ExtAttrs field.
func (o *RecordMx) SetExtAttrs(v map[string]ExtAttrs) {
	o.ExtAttrs = &v
}

// GetForbidReclamation returns the ForbidReclamation field value if set, zero value otherwise.
func (o *RecordMx) GetForbidReclamation() bool {
	if o == nil || IsNil(o.ForbidReclamation) {
		var ret bool
		return ret
	}
	return *o.ForbidReclamation
}

// GetForbidReclamationOk returns a tuple with the ForbidReclamation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordMx) GetForbidReclamationOk() (*bool, bool) {
	if o == nil || IsNil(o.ForbidReclamation) {
		return nil, false
	}
	return o.ForbidReclamation, true
}

// HasForbidReclamation returns a boolean if a field has been set.
func (o *RecordMx) HasForbidReclamation() bool {
	if o != nil && !IsNil(o.ForbidReclamation) {
		return true
	}

	return false
}

// SetForbidReclamation gets a reference to the given bool and assigns it to the ForbidReclamation field.
func (o *RecordMx) SetForbidReclamation(v bool) {
	o.ForbidReclamation = &v
}

// GetLastQueried returns the LastQueried field value if set, zero value otherwise.
func (o *RecordMx) GetLastQueried() int64 {
	if o == nil || IsNil(o.LastQueried) {
		var ret int64
		return ret
	}
	return *o.LastQueried
}

// GetLastQueriedOk returns a tuple with the LastQueried field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordMx) GetLastQueriedOk() (*int64, bool) {
	if o == nil || IsNil(o.LastQueried) {
		return nil, false
	}
	return o.LastQueried, true
}

// HasLastQueried returns a boolean if a field has been set.
func (o *RecordMx) HasLastQueried() bool {
	if o != nil && !IsNil(o.LastQueried) {
		return true
	}

	return false
}

// SetLastQueried gets a reference to the given int64 and assigns it to the LastQueried field.
func (o *RecordMx) SetLastQueried(v int64) {
	o.LastQueried = &v
}

// GetMailExchanger returns the MailExchanger field value if set, zero value otherwise.
func (o *RecordMx) GetMailExchanger() string {
	if o == nil || IsNil(o.MailExchanger) {
		var ret string
		return ret
	}
	return *o.MailExchanger
}

// GetMailExchangerOk returns a tuple with the MailExchanger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordMx) GetMailExchangerOk() (*string, bool) {
	if o == nil || IsNil(o.MailExchanger) {
		return nil, false
	}
	return o.MailExchanger, true
}

// HasMailExchanger returns a boolean if a field has been set.
func (o *RecordMx) HasMailExchanger() bool {
	if o != nil && !IsNil(o.MailExchanger) {
		return true
	}

	return false
}

// SetMailExchanger gets a reference to the given string and assigns it to the MailExchanger field.
func (o *RecordMx) SetMailExchanger(v string) {
	o.MailExchanger = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RecordMx) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordMx) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RecordMx) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RecordMx) SetName(v string) {
	o.Name = &v
}

// GetPreference returns the Preference field value if set, zero value otherwise.
func (o *RecordMx) GetPreference() int64 {
	if o == nil || IsNil(o.Preference) {
		var ret int64
		return ret
	}
	return *o.Preference
}

// GetPreferenceOk returns a tuple with the Preference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordMx) GetPreferenceOk() (*int64, bool) {
	if o == nil || IsNil(o.Preference) {
		return nil, false
	}
	return o.Preference, true
}

// HasPreference returns a boolean if a field has been set.
func (o *RecordMx) HasPreference() bool {
	if o != nil && !IsNil(o.Preference) {
		return true
	}

	return false
}

// SetPreference gets a reference to the given int64 and assigns it to the Preference field.
func (o *RecordMx) SetPreference(v int64) {
	o.Preference = &v
}

// GetReclaimable returns the Reclaimable field value if set, zero value otherwise.
func (o *RecordMx) GetReclaimable() bool {
	if o == nil || IsNil(o.Reclaimable) {
		var ret bool
		return ret
	}
	return *o.Reclaimable
}

// GetReclaimableOk returns a tuple with the Reclaimable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordMx) GetReclaimableOk() (*bool, bool) {
	if o == nil || IsNil(o.Reclaimable) {
		return nil, false
	}
	return o.Reclaimable, true
}

// HasReclaimable returns a boolean if a field has been set.
func (o *RecordMx) HasReclaimable() bool {
	if o != nil && !IsNil(o.Reclaimable) {
		return true
	}

	return false
}

// SetReclaimable gets a reference to the given bool and assigns it to the Reclaimable field.
func (o *RecordMx) SetReclaimable(v bool) {
	o.Reclaimable = &v
}

// GetSharedRecordGroup returns the SharedRecordGroup field value if set, zero value otherwise.
func (o *RecordMx) GetSharedRecordGroup() string {
	if o == nil || IsNil(o.SharedRecordGroup) {
		var ret string
		return ret
	}
	return *o.SharedRecordGroup
}

// GetSharedRecordGroupOk returns a tuple with the SharedRecordGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordMx) GetSharedRecordGroupOk() (*string, bool) {
	if o == nil || IsNil(o.SharedRecordGroup) {
		return nil, false
	}
	return o.SharedRecordGroup, true
}

// HasSharedRecordGroup returns a boolean if a field has been set.
func (o *RecordMx) HasSharedRecordGroup() bool {
	if o != nil && !IsNil(o.SharedRecordGroup) {
		return true
	}

	return false
}

// SetSharedRecordGroup gets a reference to the given string and assigns it to the SharedRecordGroup field.
func (o *RecordMx) SetSharedRecordGroup(v string) {
	o.SharedRecordGroup = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *RecordMx) GetTtl() int64 {
	if o == nil || IsNil(o.Ttl) {
		var ret int64
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordMx) GetTtlOk() (*int64, bool) {
	if o == nil || IsNil(o.Ttl) {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *RecordMx) HasTtl() bool {
	if o != nil && !IsNil(o.Ttl) {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int64 and assigns it to the Ttl field.
func (o *RecordMx) SetTtl(v int64) {
	o.Ttl = &v
}

// GetUseTtl returns the UseTtl field value if set, zero value otherwise.
func (o *RecordMx) GetUseTtl() bool {
	if o == nil || IsNil(o.UseTtl) {
		var ret bool
		return ret
	}
	return *o.UseTtl
}

// GetUseTtlOk returns a tuple with the UseTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordMx) GetUseTtlOk() (*bool, bool) {
	if o == nil || IsNil(o.UseTtl) {
		return nil, false
	}
	return o.UseTtl, true
}

// HasUseTtl returns a boolean if a field has been set.
func (o *RecordMx) HasUseTtl() bool {
	if o != nil && !IsNil(o.UseTtl) {
		return true
	}

	return false
}

// SetUseTtl gets a reference to the given bool and assigns it to the UseTtl field.
func (o *RecordMx) SetUseTtl(v bool) {
	o.UseTtl = &v
}

// GetView returns the View field value if set, zero value otherwise.
func (o *RecordMx) GetView() string {
	if o == nil || IsNil(o.View) {
		var ret string
		return ret
	}
	return *o.View
}

// GetViewOk returns a tuple with the View field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordMx) GetViewOk() (*string, bool) {
	if o == nil || IsNil(o.View) {
		return nil, false
	}
	return o.View, true
}

// HasView returns a boolean if a field has been set.
func (o *RecordMx) HasView() bool {
	if o != nil && !IsNil(o.View) {
		return true
	}

	return false
}

// SetView gets a reference to the given string and assigns it to the View field.
func (o *RecordMx) SetView(v string) {
	o.View = &v
}

// GetZone returns the Zone field value if set, zero value otherwise.
func (o *RecordMx) GetZone() string {
	if o == nil || IsNil(o.Zone) {
		var ret string
		return ret
	}
	return *o.Zone
}

// GetZoneOk returns a tuple with the Zone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordMx) GetZoneOk() (*string, bool) {
	if o == nil || IsNil(o.Zone) {
		return nil, false
	}
	return o.Zone, true
}

// HasZone returns a boolean if a field has been set.
func (o *RecordMx) HasZone() bool {
	if o != nil && !IsNil(o.Zone) {
		return true
	}

	return false
}

// SetZone gets a reference to the given string and assigns it to the Zone field.
func (o *RecordMx) SetZone(v string) {
	o.Zone = &v
}

func (o RecordMx) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecordMx) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ref) {
		toSerialize["_ref"] = o.Ref
	}
	if !IsNil(o.AwsRte53RecordInfo) {
		toSerialize["aws_rte53_record_info"] = o.AwsRte53RecordInfo
	}
	if !IsNil(o.CloudInfo) {
		toSerialize["cloud_info"] = o.CloudInfo
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.CreationTime) {
		toSerialize["creation_time"] = o.CreationTime
	}
	if !IsNil(o.Creator) {
		toSerialize["creator"] = o.Creator
	}
	if !IsNil(o.DdnsPrincipal) {
		toSerialize["ddns_principal"] = o.DdnsPrincipal
	}
	if !IsNil(o.DdnsProtected) {
		toSerialize["ddns_protected"] = o.DdnsProtected
	}
	if !IsNil(o.Disable) {
		toSerialize["disable"] = o.Disable
	}
	if !IsNil(o.DnsMailExchanger) {
		toSerialize["dns_mail_exchanger"] = o.DnsMailExchanger
	}
	if !IsNil(o.DnsName) {
		toSerialize["dns_name"] = o.DnsName
	}
	if !IsNil(o.ExtAttrs) {
		toSerialize["extattrs"] = o.ExtAttrs
	}
	if !IsNil(o.ForbidReclamation) {
		toSerialize["forbid_reclamation"] = o.ForbidReclamation
	}
	if !IsNil(o.LastQueried) {
		toSerialize["last_queried"] = o.LastQueried
	}
	if !IsNil(o.MailExchanger) {
		toSerialize["mail_exchanger"] = o.MailExchanger
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Preference) {
		toSerialize["preference"] = o.Preference
	}
	if !IsNil(o.Reclaimable) {
		toSerialize["reclaimable"] = o.Reclaimable
	}
	if !IsNil(o.SharedRecordGroup) {
		toSerialize["shared_record_group"] = o.SharedRecordGroup
	}
	if !IsNil(o.Ttl) {
		toSerialize["ttl"] = o.Ttl
	}
	if !IsNil(o.UseTtl) {
		toSerialize["use_ttl"] = o.UseTtl
	}
	if !IsNil(o.View) {
		toSerialize["view"] = o.View
	}
	if !IsNil(o.Zone) {
		toSerialize["zone"] = o.Zone
	}
	return toSerialize, nil
}

type NullableRecordMx struct {
	value *RecordMx
	isSet bool
}

func (v NullableRecordMx) Get() *RecordMx {
	return v.value
}

func (v *NullableRecordMx) Set(val *RecordMx) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordMx) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordMx) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordMx(val *RecordMx) *NullableRecordMx {
	return &NullableRecordMx{value: val, isSet: true}
}

func (v NullableRecordMx) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecordMx) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
