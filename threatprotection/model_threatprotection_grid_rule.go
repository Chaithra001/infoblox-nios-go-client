/*
Infoblox THREATPROTECTION API

OpenAPI specification for Infoblox NIOS WAPI THREATPROTECTION objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package threatprotection

import (
	"encoding/json"
)

// checks if the ThreatprotectionGridRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThreatprotectionGridRule{}

// ThreatprotectionGridRule struct for ThreatprotectionGridRule
type ThreatprotectionGridRule struct {
	// The reference to the object.
	Ref *string `json:"_ref,omitempty"`
	// The list of allowed actions of the custom rule.
	AllowedActions []string `json:"allowed_actions,omitempty"`
	// The rule category the custom rule assigned to.
	Category *string `json:"category,omitempty"`
	// The human readable comment for the custom rule.
	Comment *string                         `json:"comment,omitempty"`
	Config  *ThreatprotectionGridRuleConfig `json:"config,omitempty"`
	// The description of the custom rule.
	Description *string `json:"description,omitempty"`
	// Determines if the custom rule is disabled.
	Disabled *bool `json:"disabled,omitempty"`
	// Determines if factory reset is enabled for the custom rule.
	IsFactoryResetEnabled *bool `json:"is_factory_reset_enabled,omitempty"`
	// The name of the rule custom rule concatenated with its rule config parameters.
	Name *string `json:"name,omitempty"`
	// The version of the ruleset the custom rule assigned to.
	Ruleset *string `json:"ruleset,omitempty"`
	// The Rule ID.
	Sid *int64 `json:"sid,omitempty"`
	// The threat protection rule template used to create this rule.
	Template *string `json:"template,omitempty"`
	// The type of the custom rule.
	Type *string `json:"type,omitempty"`
}

// NewThreatprotectionGridRule instantiates a new ThreatprotectionGridRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreatprotectionGridRule() *ThreatprotectionGridRule {
	this := ThreatprotectionGridRule{}
	return &this
}

// NewThreatprotectionGridRuleWithDefaults instantiates a new ThreatprotectionGridRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreatprotectionGridRuleWithDefaults() *ThreatprotectionGridRule {
	this := ThreatprotectionGridRule{}
	return &this
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *ThreatprotectionGridRule) GetRef() string {
	if o == nil || IsNil(o.Ref) {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatprotectionGridRule) GetRefOk() (*string, bool) {
	if o == nil || IsNil(o.Ref) {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *ThreatprotectionGridRule) HasRef() bool {
	if o != nil && !IsNil(o.Ref) {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *ThreatprotectionGridRule) SetRef(v string) {
	o.Ref = &v
}

// GetAllowedActions returns the AllowedActions field value if set, zero value otherwise.
func (o *ThreatprotectionGridRule) GetAllowedActions() []string {
	if o == nil || IsNil(o.AllowedActions) {
		var ret []string
		return ret
	}
	return o.AllowedActions
}

// GetAllowedActionsOk returns a tuple with the AllowedActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatprotectionGridRule) GetAllowedActionsOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedActions) {
		return nil, false
	}
	return o.AllowedActions, true
}

// HasAllowedActions returns a boolean if a field has been set.
func (o *ThreatprotectionGridRule) HasAllowedActions() bool {
	if o != nil && !IsNil(o.AllowedActions) {
		return true
	}

	return false
}

// SetAllowedActions gets a reference to the given []string and assigns it to the AllowedActions field.
func (o *ThreatprotectionGridRule) SetAllowedActions(v []string) {
	o.AllowedActions = v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *ThreatprotectionGridRule) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatprotectionGridRule) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *ThreatprotectionGridRule) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *ThreatprotectionGridRule) SetCategory(v string) {
	o.Category = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ThreatprotectionGridRule) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatprotectionGridRule) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ThreatprotectionGridRule) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ThreatprotectionGridRule) SetComment(v string) {
	o.Comment = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *ThreatprotectionGridRule) GetConfig() ThreatprotectionGridRuleConfig {
	if o == nil || IsNil(o.Config) {
		var ret ThreatprotectionGridRuleConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatprotectionGridRule) GetConfigOk() (*ThreatprotectionGridRuleConfig, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *ThreatprotectionGridRule) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given ThreatprotectionGridRuleConfig and assigns it to the Config field.
func (o *ThreatprotectionGridRule) SetConfig(v ThreatprotectionGridRuleConfig) {
	o.Config = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ThreatprotectionGridRule) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatprotectionGridRule) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ThreatprotectionGridRule) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ThreatprotectionGridRule) SetDescription(v string) {
	o.Description = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *ThreatprotectionGridRule) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatprotectionGridRule) GetDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *ThreatprotectionGridRule) HasDisabled() bool {
	if o != nil && !IsNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *ThreatprotectionGridRule) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetIsFactoryResetEnabled returns the IsFactoryResetEnabled field value if set, zero value otherwise.
func (o *ThreatprotectionGridRule) GetIsFactoryResetEnabled() bool {
	if o == nil || IsNil(o.IsFactoryResetEnabled) {
		var ret bool
		return ret
	}
	return *o.IsFactoryResetEnabled
}

// GetIsFactoryResetEnabledOk returns a tuple with the IsFactoryResetEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatprotectionGridRule) GetIsFactoryResetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFactoryResetEnabled) {
		return nil, false
	}
	return o.IsFactoryResetEnabled, true
}

// HasIsFactoryResetEnabled returns a boolean if a field has been set.
func (o *ThreatprotectionGridRule) HasIsFactoryResetEnabled() bool {
	if o != nil && !IsNil(o.IsFactoryResetEnabled) {
		return true
	}

	return false
}

// SetIsFactoryResetEnabled gets a reference to the given bool and assigns it to the IsFactoryResetEnabled field.
func (o *ThreatprotectionGridRule) SetIsFactoryResetEnabled(v bool) {
	o.IsFactoryResetEnabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ThreatprotectionGridRule) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatprotectionGridRule) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ThreatprotectionGridRule) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ThreatprotectionGridRule) SetName(v string) {
	o.Name = &v
}

// GetRuleset returns the Ruleset field value if set, zero value otherwise.
func (o *ThreatprotectionGridRule) GetRuleset() string {
	if o == nil || IsNil(o.Ruleset) {
		var ret string
		return ret
	}
	return *o.Ruleset
}

// GetRulesetOk returns a tuple with the Ruleset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatprotectionGridRule) GetRulesetOk() (*string, bool) {
	if o == nil || IsNil(o.Ruleset) {
		return nil, false
	}
	return o.Ruleset, true
}

// HasRuleset returns a boolean if a field has been set.
func (o *ThreatprotectionGridRule) HasRuleset() bool {
	if o != nil && !IsNil(o.Ruleset) {
		return true
	}

	return false
}

// SetRuleset gets a reference to the given string and assigns it to the Ruleset field.
func (o *ThreatprotectionGridRule) SetRuleset(v string) {
	o.Ruleset = &v
}

// GetSid returns the Sid field value if set, zero value otherwise.
func (o *ThreatprotectionGridRule) GetSid() int64 {
	if o == nil || IsNil(o.Sid) {
		var ret int64
		return ret
	}
	return *o.Sid
}

// GetSidOk returns a tuple with the Sid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatprotectionGridRule) GetSidOk() (*int64, bool) {
	if o == nil || IsNil(o.Sid) {
		return nil, false
	}
	return o.Sid, true
}

// HasSid returns a boolean if a field has been set.
func (o *ThreatprotectionGridRule) HasSid() bool {
	if o != nil && !IsNil(o.Sid) {
		return true
	}

	return false
}

// SetSid gets a reference to the given int64 and assigns it to the Sid field.
func (o *ThreatprotectionGridRule) SetSid(v int64) {
	o.Sid = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *ThreatprotectionGridRule) GetTemplate() string {
	if o == nil || IsNil(o.Template) {
		var ret string
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatprotectionGridRule) GetTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.Template) {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *ThreatprotectionGridRule) HasTemplate() bool {
	if o != nil && !IsNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given string and assigns it to the Template field.
func (o *ThreatprotectionGridRule) SetTemplate(v string) {
	o.Template = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ThreatprotectionGridRule) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatprotectionGridRule) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ThreatprotectionGridRule) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ThreatprotectionGridRule) SetType(v string) {
	o.Type = &v
}

func (o ThreatprotectionGridRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThreatprotectionGridRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ref) {
		toSerialize["_ref"] = o.Ref
	}
	if !IsNil(o.AllowedActions) {
		toSerialize["allowed_actions"] = o.AllowedActions
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	if !IsNil(o.IsFactoryResetEnabled) {
		toSerialize["is_factory_reset_enabled"] = o.IsFactoryResetEnabled
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Ruleset) {
		toSerialize["ruleset"] = o.Ruleset
	}
	if !IsNil(o.Sid) {
		toSerialize["sid"] = o.Sid
	}
	if !IsNil(o.Template) {
		toSerialize["template"] = o.Template
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableThreatprotectionGridRule struct {
	value *ThreatprotectionGridRule
	isSet bool
}

func (v NullableThreatprotectionGridRule) Get() *ThreatprotectionGridRule {
	return v.value
}

func (v *NullableThreatprotectionGridRule) Set(val *ThreatprotectionGridRule) {
	v.value = val
	v.isSet = true
}

func (v NullableThreatprotectionGridRule) IsSet() bool {
	return v.isSet
}

func (v *NullableThreatprotectionGridRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreatprotectionGridRule(val *ThreatprotectionGridRule) *NullableThreatprotectionGridRule {
	return &NullableThreatprotectionGridRule{value: val, isSet: true}
}

func (v NullableThreatprotectionGridRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreatprotectionGridRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
