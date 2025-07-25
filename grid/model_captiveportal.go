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

// checks if the Captiveportal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Captiveportal{}

// Captiveportal struct for Captiveportal
type Captiveportal struct {
	// The reference to the object.
	Ref *string `json:"_ref,omitempty"`
	// The authentication server group assigned to this captive portal.
	AuthnServerGroup *string `json:"authn_server_group,omitempty"`
	// The company name that appears in the guest registration page.
	CompanyName *string `json:"company_name,omitempty"`
	// Determines if authentication failures are logged to syslog or not.
	EnableSyslogAuthFailure *bool `json:"enable_syslog_auth_failure,omitempty"`
	// Determines if successful authentications are logged to syslog or not.
	EnableSyslogAuthSuccess *bool `json:"enable_syslog_auth_success,omitempty"`
	// The type of user to be enabled for the captive portal.
	EnableUserType *string `json:"enable_user_type,omitempty"`
	// The encryption the captive portal uses.
	Encryption *string `json:"encryption,omitempty"`
	// The list of files associated with the captive portal.
	Files []CaptiveportalFiles `json:"files,omitempty"`
	// The name of the custom field that you are adding to the guest registration page.
	GuestCustomField1Name *string `json:"guest_custom_field1_name,omitempty"`
	// Determines if the custom field is required or not.
	GuestCustomField1Required *bool `json:"guest_custom_field1_required,omitempty"`
	// The name of the custom field that you are adding to the guest registration page.
	GuestCustomField2Name *string `json:"guest_custom_field2_name,omitempty"`
	// Determines if the custom field is required or not.
	GuestCustomField2Required *bool `json:"guest_custom_field2_required,omitempty"`
	// The name of the custom field that you are adding to the guest registration page.
	GuestCustomField3Name *string `json:"guest_custom_field3_name,omitempty"`
	// Determines if the custom field is required or not.
	GuestCustomField3Required *bool `json:"guest_custom_field3_required,omitempty"`
	// The name of the custom field that you are adding to the guest registration page.
	GuestCustomField4Name *string `json:"guest_custom_field4_name,omitempty"`
	// Determines if the custom field is required or not.
	GuestCustomField4Required *bool `json:"guest_custom_field4_required,omitempty"`
	// Determines if the email address of the guest is required or not.
	GuestEmailRequired *bool `json:"guest_email_required,omitempty"`
	// Determines if the first name of the guest is required or not.
	GuestFirstNameRequired *bool `json:"guest_first_name_required,omitempty"`
	// Determines if the last name of the guest is required or not.
	GuestLastNameRequired *bool `json:"guest_last_name_required,omitempty"`
	// Determines if the middle name of the guest is required or not.
	GuestMiddleNameRequired *bool `json:"guest_middle_name_required,omitempty"`
	// Determines if the phone number of the guest is required or not.
	GuestPhoneRequired *bool `json:"guest_phone_required,omitempty"`
	// The helpdesk message that appears in the guest registration page.
	HelpdeskMessage *string `json:"helpdesk_message,omitempty"`
	// Determines the IP address on which the captive portal listens. Valid if listen address type is 'IP'.
	ListenAddressIp *string `json:"listen_address_ip,omitempty"`
	// Determines the type of the IP address on which the captive portal listens.
	ListenAddressType *string `json:"listen_address_type,omitempty"`
	// The hostname of the Grid member that hosts the captive portal.
	Name *string `json:"name,omitempty"`
	// The network view of the captive portal.
	NetworkView *string `json:"network_view,omitempty"`
	// The TCP port used by the Captive Portal service. The port is required when the Captive Portal service is enabled. Valid values are between 1 and 63999. Please note that setting the port number to 80 or 443 might impact performance.
	Port *int64 `json:"port,omitempty"`
	// Determines if the captive portal service is enabled or not.
	ServiceEnabled *bool `json:"service_enabled,omitempty"`
	// The syslog level at which authentication failures are logged.
	SyslogAuthFailureLevel *string `json:"syslog_auth_failure_level,omitempty"`
	// The syslog level at which successful authentications are logged.
	SyslogAuthSuccessLevel *string `json:"syslog_auth_success_level,omitempty"`
	// The welcome message that appears in the guest registration page.
	WelcomeMessage *string `json:"welcome_message,omitempty"`
}

// NewCaptiveportal instantiates a new Captiveportal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaptiveportal() *Captiveportal {
	this := Captiveportal{}
	return &this
}

// NewCaptiveportalWithDefaults instantiates a new Captiveportal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaptiveportalWithDefaults() *Captiveportal {
	this := Captiveportal{}
	return &this
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *Captiveportal) GetRef() string {
	if o == nil || IsNil(o.Ref) {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Captiveportal) GetRefOk() (*string, bool) {
	if o == nil || IsNil(o.Ref) {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *Captiveportal) HasRef() bool {
	if o != nil && !IsNil(o.Ref) {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *Captiveportal) SetRef(v string) {
	o.Ref = &v
}

// GetAuthnServerGroup returns the AuthnServerGroup field value if set, zero value otherwise.
func (o *Captiveportal) GetAuthnServerGroup() string {
	if o == nil || IsNil(o.AuthnServerGroup) {
		var ret string
		return ret
	}
	return *o.AuthnServerGroup
}

// GetAuthnServerGroupOk returns a tuple with the AuthnServerGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Captiveportal) GetAuthnServerGroupOk() (*string, bool) {
	if o == nil || IsNil(o.AuthnServerGroup) {
		return nil, false
	}
	return o.AuthnServerGroup, true
}

// HasAuthnServerGroup returns a boolean if a field has been set.
func (o *Captiveportal) HasAuthnServerGroup() bool {
	if o != nil && !IsNil(o.AuthnServerGroup) {
		return true
	}

	return false
}

// SetAuthnServerGroup gets a reference to the given string and assigns it to the AuthnServerGroup field.
func (o *Captiveportal) SetAuthnServerGroup(v string) {
	o.AuthnServerGroup = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *Captiveportal) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Captiveportal) GetCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyName) {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *Captiveportal) HasCompanyName() bool {
	if o != nil && !IsNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *Captiveportal) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetEnableSyslogAuthFailure returns the EnableSyslogAuthFailure field value if set, zero value otherwise.
func (o *Captiveportal) GetEnableSyslogAuthFailure() bool {
	if o == nil || IsNil(o.EnableSyslogAuthFailure) {
		var ret bool
		return ret
	}
	return *o.EnableSyslogAuthFailure
}

// GetEnableSyslogAuthFailureOk returns a tuple with the EnableSyslogAuthFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Captiveportal) GetEnableSyslogAuthFailureOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableSyslogAuthFailure) {
		return nil, false
	}
	return o.EnableSyslogAuthFailure, true
}

// HasEnableSyslogAuthFailure returns a boolean if a field has been set.
func (o *Captiveportal) HasEnableSyslogAuthFailure() bool {
	if o != nil && !IsNil(o.EnableSyslogAuthFailure) {
		return true
	}

	return false
}

// SetEnableSyslogAuthFailure gets a reference to the given bool and assigns it to the EnableSyslogAuthFailure field.
func (o *Captiveportal) SetEnableSyslogAuthFailure(v bool) {
	o.EnableSyslogAuthFailure = &v
}

// GetEnableSyslogAuthSuccess returns the EnableSyslogAuthSuccess field value if set, zero value otherwise.
func (o *Captiveportal) GetEnableSyslogAuthSuccess() bool {
	if o == nil || IsNil(o.EnableSyslogAuthSuccess) {
		var ret bool
		return ret
	}
	return *o.EnableSyslogAuthSuccess
}

// GetEnableSyslogAuthSuccessOk returns a tuple with the EnableSyslogAuthSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Captiveportal) GetEnableSyslogAuthSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableSyslogAuthSuccess) {
		return nil, false
	}
	return o.EnableSyslogAuthSuccess, true
}

// HasEnableSyslogAuthSuccess returns a boolean if a field has been set.
func (o *Captiveportal) HasEnableSyslogAuthSuccess() bool {
	if o != nil && !IsNil(o.EnableSyslogAuthSuccess) {
		return true
	}

	return false
}

// SetEnableSyslogAuthSuccess gets a reference to the given bool and assigns it to the EnableSyslogAuthSuccess field.
func (o *Captiveportal) SetEnableSyslogAuthSuccess(v bool) {
	o.EnableSyslogAuthSuccess = &v
}

// GetEnableUserType returns the EnableUserType field value if set, zero value otherwise.
func (o *Captiveportal) GetEnableUserType() string {
	if o == nil || IsNil(o.EnableUserType) {
		var ret string
		return ret
	}
	return *o.EnableUserType
}

// GetEnableUserTypeOk returns a tuple with the EnableUserType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Captiveportal) GetEnableUserTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EnableUserType) {
		return nil, false
	}
	return o.EnableUserType, true
}

// HasEnableUserType returns a boolean if a field has been set.
func (o *Captiveportal) HasEnableUserType() bool {
	if o != nil && !IsNil(o.EnableUserType) {
		return true
	}

	return false
}

// SetEnableUserType gets a reference to the given string and assigns it to the EnableUserType field.
func (o *Captiveportal) SetEnableUserType(v string) {
	o.EnableUserType = &v
}

// GetEncryption returns the Encryption field value if set, zero value otherwise.
func (o *Captiveportal) GetEncryption() string {
	if o == nil || IsNil(o.Encryption) {
		var ret string
		return ret
	}
	return *o.Encryption
}

// GetEncryptionOk returns a tuple with the Encryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Captiveportal) GetEncryptionOk() (*string, bool) {
	if o == nil || IsNil(o.Encryption) {
		return nil, false
	}
	return o.Encryption, true
}

// HasEncryption returns a boolean if a field has been set.
func (o *Captiveportal) HasEncryption() bool {
	if o != nil && !IsNil(o.Encryption) {
		return true
	}

	return false
}

// SetEncryption gets a reference to the given string and assigns it to the Encryption field.
func (o *Captiveportal) SetEncryption(v string) {
	o.Encryption = &v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *Captiveportal) GetFiles() []CaptiveportalFiles {
	if o == nil || IsNil(o.Files) {
		var ret []CaptiveportalFiles
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Captiveportal) GetFilesOk() ([]CaptiveportalFiles, bool) {
	if o == nil || IsNil(o.Files) {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *Captiveportal) HasFiles() bool {
	if o != nil && !IsNil(o.Files) {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []CaptiveportalFiles and assigns it to the Files field.
func (o *Captiveportal) SetFiles(v []CaptiveportalFiles) {
	o.Files = v
}

// GetGuestCustomField1Name returns the GuestCustomField1Name field value if set, zero value otherwise.
func (o *Captiveportal) GetGuestCustomField1Name() string {
	if o == nil || IsNil(o.GuestCustomField1Name) {
		var ret string
		return ret
	}
	return *o.GuestCustomField1Name
}

// GetGuestCustomField1NameOk returns a tuple with the GuestCustomField1Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Captiveportal) GetGuestCustomField1NameOk() (*string, bool) {
	if o == nil || IsNil(o.GuestCustomField1Name) {
		return nil, false
	}
	return o.GuestCustomField1Name, true
}

// HasGuestCustomField1Name returns a boolean if a field has been set.
func (o *Captiveportal) HasGuestCustomField1Name() bool {
	if o != nil && !IsNil(o.GuestCustomField1Name) {
		return true
	}

	return false
}

// SetGuestCustomField1Name gets a reference to the given string and assigns it to the GuestCustomField1Name field.
func (o *Captiveportal) SetGuestCustomField1Name(v string) {
	o.GuestCustomField1Name = &v
}

// GetGuestCustomField1Required returns the GuestCustomField1Required field value if set, zero value otherwise.
func (o *Captiveportal) GetGuestCustomField1Required() bool {
	if o == nil || IsNil(o.GuestCustomField1Required) {
		var ret bool
		return ret
	}
	return *o.GuestCustomField1Required
}

// GetGuestCustomField1RequiredOk returns a tuple with the GuestCustomField1Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Captiveportal) GetGuestCustomField1RequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.GuestCustomField1Required) {
		return nil, false
	}
	return o.GuestCustomField1Required, true
}

// HasGuestCustomField1Required returns a boolean if a field has been set.
func (o *Captiveportal) HasGuestCustomField1Required() bool {
	if o != nil && !IsNil(o.GuestCustomField1Required) {
		return true
	}

	return false
}

// SetGuestCustomField1Required gets a reference to the given bool and assigns it to the GuestCustomField1Required field.
func (o *Captiveportal) SetGuestCustomField1Required(v bool) {
	o.GuestCustomField1Required = &v
}

// GetGuestCustomField2Name returns the GuestCustomField2Name field value if set, zero value otherwise.
func (o *Captiveportal) GetGuestCustomField2Name() string {
	if o == nil || IsNil(o.GuestCustomField2Name) {
		var ret string
		return ret
	}
	return *o.GuestCustomField2Name
}

// GetGuestCustomField2NameOk returns a tuple with the GuestCustomField2Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Captiveportal) GetGuestCustomField2NameOk() (*string, bool) {
	if o == nil || IsNil(o.GuestCustomField2Name) {
		return nil, false
	}
	return o.GuestCustomField2Name, true
}

// HasGuestCustomField2Name returns a boolean if a field has been set.
func (o *Captiveportal) HasGuestCustomField2Name() bool {
	if o != nil && !IsNil(o.GuestCustomField2Name) {
		return true
	}

	return false
}

// SetGuestCustomField2Name gets a reference to the given string and assigns it to the GuestCustomField2Name field.
func (o *Captiveportal) SetGuestCustomField2Name(v string) {
	o.GuestCustomField2Name = &v
}

// GetGuestCustomField2Required returns the GuestCustomField2Required field value if set, zero value otherwise.
func (o *Captiveportal) GetGuestCustomField2Required() bool {
	if o == nil || IsNil(o.GuestCustomField2Required) {
		var ret bool
		return ret
	}
	return *o.GuestCustomField2Required
}

// GetGuestCustomField2RequiredOk returns a tuple with the GuestCustomField2Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Captiveportal) GetGuestCustomField2RequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.GuestCustomField2Required) {
		return nil, false
	}
	return o.GuestCustomField2Required, true
}

// HasGuestCustomField2Required returns a boolean if a field has been set.
func (o *Captiveportal) HasGuestCustomField2Required() bool {
	if o != nil && !IsNil(o.GuestCustomField2Required) {
		return true
	}

	return false
}

// SetGuestCustomField2Required gets a reference to the given bool and assigns it to the GuestCustomField2Required field.
func (o *Captiveportal) SetGuestCustomField2Required(v bool) {
	o.GuestCustomField2Required = &v
}

// GetGuestCustomField3Name returns the GuestCustomField3Name field value if set, zero value otherwise.
func (o *Captiveportal) GetGuestCustomField3Name() string {
	if o == nil || IsNil(o.GuestCustomField3Name) {
		var ret string
		return ret
	}
	return *o.GuestCustomField3Name
}

// GetGuestCustomField3NameOk returns a tuple with the GuestCustomField3Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Captiveportal) GetGuestCustomField3NameOk() (*string, bool) {
	if o == nil || IsNil(o.GuestCustomField3Name) {
		return nil, false
	}
	return o.GuestCustomField3Name, true
}

// HasGuestCustomField3Name returns a boolean if a field has been set.
func (o *Captiveportal) HasGuestCustomField3Name() bool {
	if o != nil && !IsNil(o.GuestCustomField3Name) {
		return true
	}

	return false
}

// SetGuestCustomField3Name gets a reference to the given string and assigns it to the GuestCustomField3Name field.
func (o *Captiveportal) SetGuestCustomField3Name(v string) {
	o.GuestCustomField3Name = &v
}

// GetGuestCustomField3Required returns the GuestCustomField3Required field value if set, zero value otherwise.
func (o *Captiveportal) GetGuestCustomField3Required() bool {
	if o == nil || IsNil(o.GuestCustomField3Required) {
		var ret bool
		return ret
	}
	return *o.GuestCustomField3Required
}

// GetGuestCustomField3RequiredOk returns a tuple with the GuestCustomField3Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Captiveportal) GetGuestCustomField3RequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.GuestCustomField3Required) {
		return nil, false
	}
	return o.GuestCustomField3Required, true
}

// HasGuestCustomField3Required returns a boolean if a field has been set.
func (o *Captiveportal) HasGuestCustomField3Required() bool {
	if o != nil && !IsNil(o.GuestCustomField3Required) {
		return true
	}

	return false
}

// SetGuestCustomField3Required gets a reference to the given bool and assigns it to the GuestCustomField3Required field.
func (o *Captiveportal) SetGuestCustomField3Required(v bool) {
	o.GuestCustomField3Required = &v
}

// GetGuestCustomField4Name returns the GuestCustomField4Name field value if set, zero value otherwise.
func (o *Captiveportal) GetGuestCustomField4Name() string {
	if o == nil || IsNil(o.GuestCustomField4Name) {
		var ret string
		return ret
	}
	return *o.GuestCustomField4Name
}

// GetGuestCustomField4NameOk returns a tuple with the GuestCustomField4Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Captiveportal) GetGuestCustomField4NameOk() (*string, bool) {
	if o == nil || IsNil(o.GuestCustomField4Name) {
		return nil, false
	}
	return o.GuestCustomField4Name, true
}

// HasGuestCustomField4Name returns a boolean if a field has been set.
func (o *Captiveportal) HasGuestCustomField4Name() bool {
	if o != nil && !IsNil(o.GuestCustomField4Name) {
		return true
	}

	return false
}

// SetGuestCustomField4Name gets a reference to the given string and assigns it to the GuestCustomField4Name field.
func (o *Captiveportal) SetGuestCustomField4Name(v string) {
	o.GuestCustomField4Name = &v
}

// GetGuestCustomField4Required returns the GuestCustomField4Required field value if set, zero value otherwise.
func (o *Captiveportal) GetGuestCustomField4Required() bool {
	if o == nil || IsNil(o.GuestCustomField4Required) {
		var ret bool
		return ret
	}
	return *o.GuestCustomField4Required
}

// GetGuestCustomField4RequiredOk returns a tuple with the GuestCustomField4Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Captiveportal) GetGuestCustomField4RequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.GuestCustomField4Required) {
		return nil, false
	}
	return o.GuestCustomField4Required, true
}

// HasGuestCustomField4Required returns a boolean if a field has been set.
func (o *Captiveportal) HasGuestCustomField4Required() bool {
	if o != nil && !IsNil(o.GuestCustomField4Required) {
		return true
	}

	return false
}

// SetGuestCustomField4Required gets a reference to the given bool and assigns it to the GuestCustomField4Required field.
func (o *Captiveportal) SetGuestCustomField4Required(v bool) {
	o.GuestCustomField4Required = &v
}

// GetGuestEmailRequired returns the GuestEmailRequired field value if set, zero value otherwise.
func (o *Captiveportal) GetGuestEmailRequired() bool {
	if o == nil || IsNil(o.GuestEmailRequired) {
		var ret bool
		return ret
	}
	return *o.GuestEmailRequired
}

// GetGuestEmailRequiredOk returns a tuple with the GuestEmailRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Captiveportal) GetGuestEmailRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.GuestEmailRequired) {
		return nil, false
	}
	return o.GuestEmailRequired, true
}

// HasGuestEmailRequired returns a boolean if a field has been set.
func (o *Captiveportal) HasGuestEmailRequired() bool {
	if o != nil && !IsNil(o.GuestEmailRequired) {
		return true
	}

	return false
}

// SetGuestEmailRequired gets a reference to the given bool and assigns it to the GuestEmailRequired field.
func (o *Captiveportal) SetGuestEmailRequired(v bool) {
	o.GuestEmailRequired = &v
}

// GetGuestFirstNameRequired returns the GuestFirstNameRequired field value if set, zero value otherwise.
func (o *Captiveportal) GetGuestFirstNameRequired() bool {
	if o == nil || IsNil(o.GuestFirstNameRequired) {
		var ret bool
		return ret
	}
	return *o.GuestFirstNameRequired
}

// GetGuestFirstNameRequiredOk returns a tuple with the GuestFirstNameRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Captiveportal) GetGuestFirstNameRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.GuestFirstNameRequired) {
		return nil, false
	}
	return o.GuestFirstNameRequired, true
}

// HasGuestFirstNameRequired returns a boolean if a field has been set.
func (o *Captiveportal) HasGuestFirstNameRequired() bool {
	if o != nil && !IsNil(o.GuestFirstNameRequired) {
		return true
	}

	return false
}

// SetGuestFirstNameRequired gets a reference to the given bool and assigns it to the GuestFirstNameRequired field.
func (o *Captiveportal) SetGuestFirstNameRequired(v bool) {
	o.GuestFirstNameRequired = &v
}

// GetGuestLastNameRequired returns the GuestLastNameRequired field value if set, zero value otherwise.
func (o *Captiveportal) GetGuestLastNameRequired() bool {
	if o == nil || IsNil(o.GuestLastNameRequired) {
		var ret bool
		return ret
	}
	return *o.GuestLastNameRequired
}

// GetGuestLastNameRequiredOk returns a tuple with the GuestLastNameRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Captiveportal) GetGuestLastNameRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.GuestLastNameRequired) {
		return nil, false
	}
	return o.GuestLastNameRequired, true
}

// HasGuestLastNameRequired returns a boolean if a field has been set.
func (o *Captiveportal) HasGuestLastNameRequired() bool {
	if o != nil && !IsNil(o.GuestLastNameRequired) {
		return true
	}

	return false
}

// SetGuestLastNameRequired gets a reference to the given bool and assigns it to the GuestLastNameRequired field.
func (o *Captiveportal) SetGuestLastNameRequired(v bool) {
	o.GuestLastNameRequired = &v
}

// GetGuestMiddleNameRequired returns the GuestMiddleNameRequired field value if set, zero value otherwise.
func (o *Captiveportal) GetGuestMiddleNameRequired() bool {
	if o == nil || IsNil(o.GuestMiddleNameRequired) {
		var ret bool
		return ret
	}
	return *o.GuestMiddleNameRequired
}

// GetGuestMiddleNameRequiredOk returns a tuple with the GuestMiddleNameRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Captiveportal) GetGuestMiddleNameRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.GuestMiddleNameRequired) {
		return nil, false
	}
	return o.GuestMiddleNameRequired, true
}

// HasGuestMiddleNameRequired returns a boolean if a field has been set.
func (o *Captiveportal) HasGuestMiddleNameRequired() bool {
	if o != nil && !IsNil(o.GuestMiddleNameRequired) {
		return true
	}

	return false
}

// SetGuestMiddleNameRequired gets a reference to the given bool and assigns it to the GuestMiddleNameRequired field.
func (o *Captiveportal) SetGuestMiddleNameRequired(v bool) {
	o.GuestMiddleNameRequired = &v
}

// GetGuestPhoneRequired returns the GuestPhoneRequired field value if set, zero value otherwise.
func (o *Captiveportal) GetGuestPhoneRequired() bool {
	if o == nil || IsNil(o.GuestPhoneRequired) {
		var ret bool
		return ret
	}
	return *o.GuestPhoneRequired
}

// GetGuestPhoneRequiredOk returns a tuple with the GuestPhoneRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Captiveportal) GetGuestPhoneRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.GuestPhoneRequired) {
		return nil, false
	}
	return o.GuestPhoneRequired, true
}

// HasGuestPhoneRequired returns a boolean if a field has been set.
func (o *Captiveportal) HasGuestPhoneRequired() bool {
	if o != nil && !IsNil(o.GuestPhoneRequired) {
		return true
	}

	return false
}

// SetGuestPhoneRequired gets a reference to the given bool and assigns it to the GuestPhoneRequired field.
func (o *Captiveportal) SetGuestPhoneRequired(v bool) {
	o.GuestPhoneRequired = &v
}

// GetHelpdeskMessage returns the HelpdeskMessage field value if set, zero value otherwise.
func (o *Captiveportal) GetHelpdeskMessage() string {
	if o == nil || IsNil(o.HelpdeskMessage) {
		var ret string
		return ret
	}
	return *o.HelpdeskMessage
}

// GetHelpdeskMessageOk returns a tuple with the HelpdeskMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Captiveportal) GetHelpdeskMessageOk() (*string, bool) {
	if o == nil || IsNil(o.HelpdeskMessage) {
		return nil, false
	}
	return o.HelpdeskMessage, true
}

// HasHelpdeskMessage returns a boolean if a field has been set.
func (o *Captiveportal) HasHelpdeskMessage() bool {
	if o != nil && !IsNil(o.HelpdeskMessage) {
		return true
	}

	return false
}

// SetHelpdeskMessage gets a reference to the given string and assigns it to the HelpdeskMessage field.
func (o *Captiveportal) SetHelpdeskMessage(v string) {
	o.HelpdeskMessage = &v
}

// GetListenAddressIp returns the ListenAddressIp field value if set, zero value otherwise.
func (o *Captiveportal) GetListenAddressIp() string {
	if o == nil || IsNil(o.ListenAddressIp) {
		var ret string
		return ret
	}
	return *o.ListenAddressIp
}

// GetListenAddressIpOk returns a tuple with the ListenAddressIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Captiveportal) GetListenAddressIpOk() (*string, bool) {
	if o == nil || IsNil(o.ListenAddressIp) {
		return nil, false
	}
	return o.ListenAddressIp, true
}

// HasListenAddressIp returns a boolean if a field has been set.
func (o *Captiveportal) HasListenAddressIp() bool {
	if o != nil && !IsNil(o.ListenAddressIp) {
		return true
	}

	return false
}

// SetListenAddressIp gets a reference to the given string and assigns it to the ListenAddressIp field.
func (o *Captiveportal) SetListenAddressIp(v string) {
	o.ListenAddressIp = &v
}

// GetListenAddressType returns the ListenAddressType field value if set, zero value otherwise.
func (o *Captiveportal) GetListenAddressType() string {
	if o == nil || IsNil(o.ListenAddressType) {
		var ret string
		return ret
	}
	return *o.ListenAddressType
}

// GetListenAddressTypeOk returns a tuple with the ListenAddressType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Captiveportal) GetListenAddressTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ListenAddressType) {
		return nil, false
	}
	return o.ListenAddressType, true
}

// HasListenAddressType returns a boolean if a field has been set.
func (o *Captiveportal) HasListenAddressType() bool {
	if o != nil && !IsNil(o.ListenAddressType) {
		return true
	}

	return false
}

// SetListenAddressType gets a reference to the given string and assigns it to the ListenAddressType field.
func (o *Captiveportal) SetListenAddressType(v string) {
	o.ListenAddressType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Captiveportal) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Captiveportal) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Captiveportal) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Captiveportal) SetName(v string) {
	o.Name = &v
}

// GetNetworkView returns the NetworkView field value if set, zero value otherwise.
func (o *Captiveportal) GetNetworkView() string {
	if o == nil || IsNil(o.NetworkView) {
		var ret string
		return ret
	}
	return *o.NetworkView
}

// GetNetworkViewOk returns a tuple with the NetworkView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Captiveportal) GetNetworkViewOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkView) {
		return nil, false
	}
	return o.NetworkView, true
}

// HasNetworkView returns a boolean if a field has been set.
func (o *Captiveportal) HasNetworkView() bool {
	if o != nil && !IsNil(o.NetworkView) {
		return true
	}

	return false
}

// SetNetworkView gets a reference to the given string and assigns it to the NetworkView field.
func (o *Captiveportal) SetNetworkView(v string) {
	o.NetworkView = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *Captiveportal) GetPort() int64 {
	if o == nil || IsNil(o.Port) {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Captiveportal) GetPortOk() (*int64, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *Captiveportal) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *Captiveportal) SetPort(v int64) {
	o.Port = &v
}

// GetServiceEnabled returns the ServiceEnabled field value if set, zero value otherwise.
func (o *Captiveportal) GetServiceEnabled() bool {
	if o == nil || IsNil(o.ServiceEnabled) {
		var ret bool
		return ret
	}
	return *o.ServiceEnabled
}

// GetServiceEnabledOk returns a tuple with the ServiceEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Captiveportal) GetServiceEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ServiceEnabled) {
		return nil, false
	}
	return o.ServiceEnabled, true
}

// HasServiceEnabled returns a boolean if a field has been set.
func (o *Captiveportal) HasServiceEnabled() bool {
	if o != nil && !IsNil(o.ServiceEnabled) {
		return true
	}

	return false
}

// SetServiceEnabled gets a reference to the given bool and assigns it to the ServiceEnabled field.
func (o *Captiveportal) SetServiceEnabled(v bool) {
	o.ServiceEnabled = &v
}

// GetSyslogAuthFailureLevel returns the SyslogAuthFailureLevel field value if set, zero value otherwise.
func (o *Captiveportal) GetSyslogAuthFailureLevel() string {
	if o == nil || IsNil(o.SyslogAuthFailureLevel) {
		var ret string
		return ret
	}
	return *o.SyslogAuthFailureLevel
}

// GetSyslogAuthFailureLevelOk returns a tuple with the SyslogAuthFailureLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Captiveportal) GetSyslogAuthFailureLevelOk() (*string, bool) {
	if o == nil || IsNil(o.SyslogAuthFailureLevel) {
		return nil, false
	}
	return o.SyslogAuthFailureLevel, true
}

// HasSyslogAuthFailureLevel returns a boolean if a field has been set.
func (o *Captiveportal) HasSyslogAuthFailureLevel() bool {
	if o != nil && !IsNil(o.SyslogAuthFailureLevel) {
		return true
	}

	return false
}

// SetSyslogAuthFailureLevel gets a reference to the given string and assigns it to the SyslogAuthFailureLevel field.
func (o *Captiveportal) SetSyslogAuthFailureLevel(v string) {
	o.SyslogAuthFailureLevel = &v
}

// GetSyslogAuthSuccessLevel returns the SyslogAuthSuccessLevel field value if set, zero value otherwise.
func (o *Captiveportal) GetSyslogAuthSuccessLevel() string {
	if o == nil || IsNil(o.SyslogAuthSuccessLevel) {
		var ret string
		return ret
	}
	return *o.SyslogAuthSuccessLevel
}

// GetSyslogAuthSuccessLevelOk returns a tuple with the SyslogAuthSuccessLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Captiveportal) GetSyslogAuthSuccessLevelOk() (*string, bool) {
	if o == nil || IsNil(o.SyslogAuthSuccessLevel) {
		return nil, false
	}
	return o.SyslogAuthSuccessLevel, true
}

// HasSyslogAuthSuccessLevel returns a boolean if a field has been set.
func (o *Captiveportal) HasSyslogAuthSuccessLevel() bool {
	if o != nil && !IsNil(o.SyslogAuthSuccessLevel) {
		return true
	}

	return false
}

// SetSyslogAuthSuccessLevel gets a reference to the given string and assigns it to the SyslogAuthSuccessLevel field.
func (o *Captiveportal) SetSyslogAuthSuccessLevel(v string) {
	o.SyslogAuthSuccessLevel = &v
}

// GetWelcomeMessage returns the WelcomeMessage field value if set, zero value otherwise.
func (o *Captiveportal) GetWelcomeMessage() string {
	if o == nil || IsNil(o.WelcomeMessage) {
		var ret string
		return ret
	}
	return *o.WelcomeMessage
}

// GetWelcomeMessageOk returns a tuple with the WelcomeMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Captiveportal) GetWelcomeMessageOk() (*string, bool) {
	if o == nil || IsNil(o.WelcomeMessage) {
		return nil, false
	}
	return o.WelcomeMessage, true
}

// HasWelcomeMessage returns a boolean if a field has been set.
func (o *Captiveportal) HasWelcomeMessage() bool {
	if o != nil && !IsNil(o.WelcomeMessage) {
		return true
	}

	return false
}

// SetWelcomeMessage gets a reference to the given string and assigns it to the WelcomeMessage field.
func (o *Captiveportal) SetWelcomeMessage(v string) {
	o.WelcomeMessage = &v
}

func (o Captiveportal) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Captiveportal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ref) {
		toSerialize["_ref"] = o.Ref
	}
	if !IsNil(o.AuthnServerGroup) {
		toSerialize["authn_server_group"] = o.AuthnServerGroup
	}
	if !IsNil(o.CompanyName) {
		toSerialize["company_name"] = o.CompanyName
	}
	if !IsNil(o.EnableSyslogAuthFailure) {
		toSerialize["enable_syslog_auth_failure"] = o.EnableSyslogAuthFailure
	}
	if !IsNil(o.EnableSyslogAuthSuccess) {
		toSerialize["enable_syslog_auth_success"] = o.EnableSyslogAuthSuccess
	}
	if !IsNil(o.EnableUserType) {
		toSerialize["enable_user_type"] = o.EnableUserType
	}
	if !IsNil(o.Encryption) {
		toSerialize["encryption"] = o.Encryption
	}
	if !IsNil(o.Files) {
		toSerialize["files"] = o.Files
	}
	if !IsNil(o.GuestCustomField1Name) {
		toSerialize["guest_custom_field1_name"] = o.GuestCustomField1Name
	}
	if !IsNil(o.GuestCustomField1Required) {
		toSerialize["guest_custom_field1_required"] = o.GuestCustomField1Required
	}
	if !IsNil(o.GuestCustomField2Name) {
		toSerialize["guest_custom_field2_name"] = o.GuestCustomField2Name
	}
	if !IsNil(o.GuestCustomField2Required) {
		toSerialize["guest_custom_field2_required"] = o.GuestCustomField2Required
	}
	if !IsNil(o.GuestCustomField3Name) {
		toSerialize["guest_custom_field3_name"] = o.GuestCustomField3Name
	}
	if !IsNil(o.GuestCustomField3Required) {
		toSerialize["guest_custom_field3_required"] = o.GuestCustomField3Required
	}
	if !IsNil(o.GuestCustomField4Name) {
		toSerialize["guest_custom_field4_name"] = o.GuestCustomField4Name
	}
	if !IsNil(o.GuestCustomField4Required) {
		toSerialize["guest_custom_field4_required"] = o.GuestCustomField4Required
	}
	if !IsNil(o.GuestEmailRequired) {
		toSerialize["guest_email_required"] = o.GuestEmailRequired
	}
	if !IsNil(o.GuestFirstNameRequired) {
		toSerialize["guest_first_name_required"] = o.GuestFirstNameRequired
	}
	if !IsNil(o.GuestLastNameRequired) {
		toSerialize["guest_last_name_required"] = o.GuestLastNameRequired
	}
	if !IsNil(o.GuestMiddleNameRequired) {
		toSerialize["guest_middle_name_required"] = o.GuestMiddleNameRequired
	}
	if !IsNil(o.GuestPhoneRequired) {
		toSerialize["guest_phone_required"] = o.GuestPhoneRequired
	}
	if !IsNil(o.HelpdeskMessage) {
		toSerialize["helpdesk_message"] = o.HelpdeskMessage
	}
	if !IsNil(o.ListenAddressIp) {
		toSerialize["listen_address_ip"] = o.ListenAddressIp
	}
	if !IsNil(o.ListenAddressType) {
		toSerialize["listen_address_type"] = o.ListenAddressType
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NetworkView) {
		toSerialize["network_view"] = o.NetworkView
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.ServiceEnabled) {
		toSerialize["service_enabled"] = o.ServiceEnabled
	}
	if !IsNil(o.SyslogAuthFailureLevel) {
		toSerialize["syslog_auth_failure_level"] = o.SyslogAuthFailureLevel
	}
	if !IsNil(o.SyslogAuthSuccessLevel) {
		toSerialize["syslog_auth_success_level"] = o.SyslogAuthSuccessLevel
	}
	if !IsNil(o.WelcomeMessage) {
		toSerialize["welcome_message"] = o.WelcomeMessage
	}
	return toSerialize, nil
}

type NullableCaptiveportal struct {
	value *Captiveportal
	isSet bool
}

func (v NullableCaptiveportal) Get() *Captiveportal {
	return v.value
}

func (v *NullableCaptiveportal) Set(val *Captiveportal) {
	v.value = val
	v.isSet = true
}

func (v NullableCaptiveportal) IsSet() bool {
	return v.isSet
}

func (v *NullableCaptiveportal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaptiveportal(val *Captiveportal) *NullableCaptiveportal {
	return &NullableCaptiveportal{value: val, isSet: true}
}

func (v NullableCaptiveportal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaptiveportal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
