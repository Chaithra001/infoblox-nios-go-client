# Approvalworkflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**ApprovalGroup** | Pointer to **string** | The approval administration group. | [optional] 
**ApprovalNotifyTo** | Pointer to **string** | The destination for approval task notifications. | [optional] 
**ApprovedNotifyTo** | Pointer to **string** | The destination for approved task notifications. | [optional] 
**ApproverComment** | Pointer to **string** | The requirement for the comment when an approver approves a submitted task. | [optional] 
**EnableApprovalNotify** | Pointer to **bool** | Determines whether approval task notifications are enabled. | [optional] 
**EnableApprovedNotify** | Pointer to **bool** | Determines whether approved task notifications are enabled. | [optional] 
**EnableFailedNotify** | Pointer to **bool** | Determines whether failed task notifications are enabled. | [optional] 
**EnableNotifyGroup** | Pointer to **bool** | Determines whether e-mail notifications to admin group&#39;s e-mail address are enabled. | [optional] 
**EnableNotifyUser** | Pointer to **bool** | Determines whether e-mail notifications to an admin member&#39;s e-mail address are enabled. | [optional] 
**EnableRejectedNotify** | Pointer to **bool** | Determines whether rejected task notifications are enabled. | [optional] 
**EnableRescheduledNotify** | Pointer to **bool** | Determines whether rescheduled task notifications are enabled. | [optional] 
**EnableSucceededNotify** | Pointer to **bool** | Determines whether succeeded task notifications are enabled. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**FailedNotifyTo** | Pointer to **string** | The destination for failed task notifications. | [optional] 
**RejectedNotifyTo** | Pointer to **string** | The destination for rejected task notifications. | [optional] 
**RescheduledNotifyTo** | Pointer to **string** | The destination for rescheduled task notifications. | [optional] 
**SubmitterComment** | Pointer to **string** | The requirement for the comment when a submitter submits a task for approval. | [optional] 
**SubmitterGroup** | Pointer to **string** | The submitter admininstration group. | [optional] 
**SucceededNotifyTo** | Pointer to **string** | The destination for succeeded task notifications. | [optional] 
**TicketNumber** | Pointer to **string** | The requirement for the ticket number when a submitter submits a task for approval. | [optional] 

## Methods

### NewApprovalworkflow

`func NewApprovalworkflow() *Approvalworkflow`

NewApprovalworkflow instantiates a new Approvalworkflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalworkflowWithDefaults

`func NewApprovalworkflowWithDefaults() *Approvalworkflow`

NewApprovalworkflowWithDefaults instantiates a new Approvalworkflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Approvalworkflow) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Approvalworkflow) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Approvalworkflow) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Approvalworkflow) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetApprovalGroup

`func (o *Approvalworkflow) GetApprovalGroup() string`

GetApprovalGroup returns the ApprovalGroup field if non-nil, zero value otherwise.

### GetApprovalGroupOk

`func (o *Approvalworkflow) GetApprovalGroupOk() (*string, bool)`

GetApprovalGroupOk returns a tuple with the ApprovalGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalGroup

`func (o *Approvalworkflow) SetApprovalGroup(v string)`

SetApprovalGroup sets ApprovalGroup field to given value.

### HasApprovalGroup

`func (o *Approvalworkflow) HasApprovalGroup() bool`

HasApprovalGroup returns a boolean if a field has been set.

### GetApprovalNotifyTo

`func (o *Approvalworkflow) GetApprovalNotifyTo() string`

GetApprovalNotifyTo returns the ApprovalNotifyTo field if non-nil, zero value otherwise.

### GetApprovalNotifyToOk

`func (o *Approvalworkflow) GetApprovalNotifyToOk() (*string, bool)`

GetApprovalNotifyToOk returns a tuple with the ApprovalNotifyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalNotifyTo

`func (o *Approvalworkflow) SetApprovalNotifyTo(v string)`

SetApprovalNotifyTo sets ApprovalNotifyTo field to given value.

### HasApprovalNotifyTo

`func (o *Approvalworkflow) HasApprovalNotifyTo() bool`

HasApprovalNotifyTo returns a boolean if a field has been set.

### GetApprovedNotifyTo

`func (o *Approvalworkflow) GetApprovedNotifyTo() string`

GetApprovedNotifyTo returns the ApprovedNotifyTo field if non-nil, zero value otherwise.

### GetApprovedNotifyToOk

`func (o *Approvalworkflow) GetApprovedNotifyToOk() (*string, bool)`

GetApprovedNotifyToOk returns a tuple with the ApprovedNotifyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedNotifyTo

`func (o *Approvalworkflow) SetApprovedNotifyTo(v string)`

SetApprovedNotifyTo sets ApprovedNotifyTo field to given value.

### HasApprovedNotifyTo

`func (o *Approvalworkflow) HasApprovedNotifyTo() bool`

HasApprovedNotifyTo returns a boolean if a field has been set.

### GetApproverComment

`func (o *Approvalworkflow) GetApproverComment() string`

GetApproverComment returns the ApproverComment field if non-nil, zero value otherwise.

### GetApproverCommentOk

`func (o *Approvalworkflow) GetApproverCommentOk() (*string, bool)`

GetApproverCommentOk returns a tuple with the ApproverComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverComment

`func (o *Approvalworkflow) SetApproverComment(v string)`

SetApproverComment sets ApproverComment field to given value.

### HasApproverComment

`func (o *Approvalworkflow) HasApproverComment() bool`

HasApproverComment returns a boolean if a field has been set.

### GetEnableApprovalNotify

`func (o *Approvalworkflow) GetEnableApprovalNotify() bool`

GetEnableApprovalNotify returns the EnableApprovalNotify field if non-nil, zero value otherwise.

### GetEnableApprovalNotifyOk

`func (o *Approvalworkflow) GetEnableApprovalNotifyOk() (*bool, bool)`

GetEnableApprovalNotifyOk returns a tuple with the EnableApprovalNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableApprovalNotify

`func (o *Approvalworkflow) SetEnableApprovalNotify(v bool)`

SetEnableApprovalNotify sets EnableApprovalNotify field to given value.

### HasEnableApprovalNotify

`func (o *Approvalworkflow) HasEnableApprovalNotify() bool`

HasEnableApprovalNotify returns a boolean if a field has been set.

### GetEnableApprovedNotify

`func (o *Approvalworkflow) GetEnableApprovedNotify() bool`

GetEnableApprovedNotify returns the EnableApprovedNotify field if non-nil, zero value otherwise.

### GetEnableApprovedNotifyOk

`func (o *Approvalworkflow) GetEnableApprovedNotifyOk() (*bool, bool)`

GetEnableApprovedNotifyOk returns a tuple with the EnableApprovedNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableApprovedNotify

`func (o *Approvalworkflow) SetEnableApprovedNotify(v bool)`

SetEnableApprovedNotify sets EnableApprovedNotify field to given value.

### HasEnableApprovedNotify

`func (o *Approvalworkflow) HasEnableApprovedNotify() bool`

HasEnableApprovedNotify returns a boolean if a field has been set.

### GetEnableFailedNotify

`func (o *Approvalworkflow) GetEnableFailedNotify() bool`

GetEnableFailedNotify returns the EnableFailedNotify field if non-nil, zero value otherwise.

### GetEnableFailedNotifyOk

`func (o *Approvalworkflow) GetEnableFailedNotifyOk() (*bool, bool)`

GetEnableFailedNotifyOk returns a tuple with the EnableFailedNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFailedNotify

`func (o *Approvalworkflow) SetEnableFailedNotify(v bool)`

SetEnableFailedNotify sets EnableFailedNotify field to given value.

### HasEnableFailedNotify

`func (o *Approvalworkflow) HasEnableFailedNotify() bool`

HasEnableFailedNotify returns a boolean if a field has been set.

### GetEnableNotifyGroup

`func (o *Approvalworkflow) GetEnableNotifyGroup() bool`

GetEnableNotifyGroup returns the EnableNotifyGroup field if non-nil, zero value otherwise.

### GetEnableNotifyGroupOk

`func (o *Approvalworkflow) GetEnableNotifyGroupOk() (*bool, bool)`

GetEnableNotifyGroupOk returns a tuple with the EnableNotifyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNotifyGroup

`func (o *Approvalworkflow) SetEnableNotifyGroup(v bool)`

SetEnableNotifyGroup sets EnableNotifyGroup field to given value.

### HasEnableNotifyGroup

`func (o *Approvalworkflow) HasEnableNotifyGroup() bool`

HasEnableNotifyGroup returns a boolean if a field has been set.

### GetEnableNotifyUser

`func (o *Approvalworkflow) GetEnableNotifyUser() bool`

GetEnableNotifyUser returns the EnableNotifyUser field if non-nil, zero value otherwise.

### GetEnableNotifyUserOk

`func (o *Approvalworkflow) GetEnableNotifyUserOk() (*bool, bool)`

GetEnableNotifyUserOk returns a tuple with the EnableNotifyUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNotifyUser

`func (o *Approvalworkflow) SetEnableNotifyUser(v bool)`

SetEnableNotifyUser sets EnableNotifyUser field to given value.

### HasEnableNotifyUser

`func (o *Approvalworkflow) HasEnableNotifyUser() bool`

HasEnableNotifyUser returns a boolean if a field has been set.

### GetEnableRejectedNotify

`func (o *Approvalworkflow) GetEnableRejectedNotify() bool`

GetEnableRejectedNotify returns the EnableRejectedNotify field if non-nil, zero value otherwise.

### GetEnableRejectedNotifyOk

`func (o *Approvalworkflow) GetEnableRejectedNotifyOk() (*bool, bool)`

GetEnableRejectedNotifyOk returns a tuple with the EnableRejectedNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRejectedNotify

`func (o *Approvalworkflow) SetEnableRejectedNotify(v bool)`

SetEnableRejectedNotify sets EnableRejectedNotify field to given value.

### HasEnableRejectedNotify

`func (o *Approvalworkflow) HasEnableRejectedNotify() bool`

HasEnableRejectedNotify returns a boolean if a field has been set.

### GetEnableRescheduledNotify

`func (o *Approvalworkflow) GetEnableRescheduledNotify() bool`

GetEnableRescheduledNotify returns the EnableRescheduledNotify field if non-nil, zero value otherwise.

### GetEnableRescheduledNotifyOk

`func (o *Approvalworkflow) GetEnableRescheduledNotifyOk() (*bool, bool)`

GetEnableRescheduledNotifyOk returns a tuple with the EnableRescheduledNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRescheduledNotify

`func (o *Approvalworkflow) SetEnableRescheduledNotify(v bool)`

SetEnableRescheduledNotify sets EnableRescheduledNotify field to given value.

### HasEnableRescheduledNotify

`func (o *Approvalworkflow) HasEnableRescheduledNotify() bool`

HasEnableRescheduledNotify returns a boolean if a field has been set.

### GetEnableSucceededNotify

`func (o *Approvalworkflow) GetEnableSucceededNotify() bool`

GetEnableSucceededNotify returns the EnableSucceededNotify field if non-nil, zero value otherwise.

### GetEnableSucceededNotifyOk

`func (o *Approvalworkflow) GetEnableSucceededNotifyOk() (*bool, bool)`

GetEnableSucceededNotifyOk returns a tuple with the EnableSucceededNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSucceededNotify

`func (o *Approvalworkflow) SetEnableSucceededNotify(v bool)`

SetEnableSucceededNotify sets EnableSucceededNotify field to given value.

### HasEnableSucceededNotify

`func (o *Approvalworkflow) HasEnableSucceededNotify() bool`

HasEnableSucceededNotify returns a boolean if a field has been set.

### GetExtAttrs

`func (o *Approvalworkflow) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *Approvalworkflow) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *Approvalworkflow) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *Approvalworkflow) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetFailedNotifyTo

`func (o *Approvalworkflow) GetFailedNotifyTo() string`

GetFailedNotifyTo returns the FailedNotifyTo field if non-nil, zero value otherwise.

### GetFailedNotifyToOk

`func (o *Approvalworkflow) GetFailedNotifyToOk() (*string, bool)`

GetFailedNotifyToOk returns a tuple with the FailedNotifyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedNotifyTo

`func (o *Approvalworkflow) SetFailedNotifyTo(v string)`

SetFailedNotifyTo sets FailedNotifyTo field to given value.

### HasFailedNotifyTo

`func (o *Approvalworkflow) HasFailedNotifyTo() bool`

HasFailedNotifyTo returns a boolean if a field has been set.

### GetRejectedNotifyTo

`func (o *Approvalworkflow) GetRejectedNotifyTo() string`

GetRejectedNotifyTo returns the RejectedNotifyTo field if non-nil, zero value otherwise.

### GetRejectedNotifyToOk

`func (o *Approvalworkflow) GetRejectedNotifyToOk() (*string, bool)`

GetRejectedNotifyToOk returns a tuple with the RejectedNotifyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedNotifyTo

`func (o *Approvalworkflow) SetRejectedNotifyTo(v string)`

SetRejectedNotifyTo sets RejectedNotifyTo field to given value.

### HasRejectedNotifyTo

`func (o *Approvalworkflow) HasRejectedNotifyTo() bool`

HasRejectedNotifyTo returns a boolean if a field has been set.

### GetRescheduledNotifyTo

`func (o *Approvalworkflow) GetRescheduledNotifyTo() string`

GetRescheduledNotifyTo returns the RescheduledNotifyTo field if non-nil, zero value otherwise.

### GetRescheduledNotifyToOk

`func (o *Approvalworkflow) GetRescheduledNotifyToOk() (*string, bool)`

GetRescheduledNotifyToOk returns a tuple with the RescheduledNotifyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRescheduledNotifyTo

`func (o *Approvalworkflow) SetRescheduledNotifyTo(v string)`

SetRescheduledNotifyTo sets RescheduledNotifyTo field to given value.

### HasRescheduledNotifyTo

`func (o *Approvalworkflow) HasRescheduledNotifyTo() bool`

HasRescheduledNotifyTo returns a boolean if a field has been set.

### GetSubmitterComment

`func (o *Approvalworkflow) GetSubmitterComment() string`

GetSubmitterComment returns the SubmitterComment field if non-nil, zero value otherwise.

### GetSubmitterCommentOk

`func (o *Approvalworkflow) GetSubmitterCommentOk() (*string, bool)`

GetSubmitterCommentOk returns a tuple with the SubmitterComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitterComment

`func (o *Approvalworkflow) SetSubmitterComment(v string)`

SetSubmitterComment sets SubmitterComment field to given value.

### HasSubmitterComment

`func (o *Approvalworkflow) HasSubmitterComment() bool`

HasSubmitterComment returns a boolean if a field has been set.

### GetSubmitterGroup

`func (o *Approvalworkflow) GetSubmitterGroup() string`

GetSubmitterGroup returns the SubmitterGroup field if non-nil, zero value otherwise.

### GetSubmitterGroupOk

`func (o *Approvalworkflow) GetSubmitterGroupOk() (*string, bool)`

GetSubmitterGroupOk returns a tuple with the SubmitterGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitterGroup

`func (o *Approvalworkflow) SetSubmitterGroup(v string)`

SetSubmitterGroup sets SubmitterGroup field to given value.

### HasSubmitterGroup

`func (o *Approvalworkflow) HasSubmitterGroup() bool`

HasSubmitterGroup returns a boolean if a field has been set.

### GetSucceededNotifyTo

`func (o *Approvalworkflow) GetSucceededNotifyTo() string`

GetSucceededNotifyTo returns the SucceededNotifyTo field if non-nil, zero value otherwise.

### GetSucceededNotifyToOk

`func (o *Approvalworkflow) GetSucceededNotifyToOk() (*string, bool)`

GetSucceededNotifyToOk returns a tuple with the SucceededNotifyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceededNotifyTo

`func (o *Approvalworkflow) SetSucceededNotifyTo(v string)`

SetSucceededNotifyTo sets SucceededNotifyTo field to given value.

### HasSucceededNotifyTo

`func (o *Approvalworkflow) HasSucceededNotifyTo() bool`

HasSucceededNotifyTo returns a boolean if a field has been set.

### GetTicketNumber

`func (o *Approvalworkflow) GetTicketNumber() string`

GetTicketNumber returns the TicketNumber field if non-nil, zero value otherwise.

### GetTicketNumberOk

`func (o *Approvalworkflow) GetTicketNumberOk() (*string, bool)`

GetTicketNumberOk returns a tuple with the TicketNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketNumber

`func (o *Approvalworkflow) SetTicketNumber(v string)`

SetTicketNumber sets TicketNumber field to given value.

### HasTicketNumber

`func (o *Approvalworkflow) HasTicketNumber() bool`

HasTicketNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


