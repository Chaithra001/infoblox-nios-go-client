# GetDtcMonitorPdpResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Comment** | Pointer to **string** | Comment for this DTC monitor; maximum 256 characters. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Interval** | Pointer to **int64** | The interval for TCP health check. | [optional] 
**Name** | Pointer to **string** | The display name for this DTC monitor. | [optional] 
**Port** | Pointer to **int64** | The port value for PDP requests. | [optional] 
**RetryDown** | Pointer to **int64** | The value of how many times the server should appear as down to be treated as dead after it was alive. | [optional] 
**RetryUp** | Pointer to **int64** | The value of how many times the server should appear as up to be treated as alive after it was dead. | [optional] 
**Timeout** | Pointer to **int64** | The timeout for TCP health check in seconds. | [optional] 
**Result** | Pointer to [**DtcMonitorPdp**](DtcMonitorPdp.md) |  | [optional] 

## Methods

### NewGetDtcMonitorPdpResponse

`func NewGetDtcMonitorPdpResponse() *GetDtcMonitorPdpResponse`

NewGetDtcMonitorPdpResponse instantiates a new GetDtcMonitorPdpResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDtcMonitorPdpResponseWithDefaults

`func NewGetDtcMonitorPdpResponseWithDefaults() *GetDtcMonitorPdpResponse`

NewGetDtcMonitorPdpResponseWithDefaults instantiates a new GetDtcMonitorPdpResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetDtcMonitorPdpResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetDtcMonitorPdpResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetDtcMonitorPdpResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetDtcMonitorPdpResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *GetDtcMonitorPdpResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetDtcMonitorPdpResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetDtcMonitorPdpResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetDtcMonitorPdpResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetDtcMonitorPdpResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetDtcMonitorPdpResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetDtcMonitorPdpResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetDtcMonitorPdpResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetInterval

`func (o *GetDtcMonitorPdpResponse) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *GetDtcMonitorPdpResponse) GetIntervalOk() (*int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *GetDtcMonitorPdpResponse) SetInterval(v int64)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *GetDtcMonitorPdpResponse) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetName

`func (o *GetDtcMonitorPdpResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetDtcMonitorPdpResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetDtcMonitorPdpResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetDtcMonitorPdpResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *GetDtcMonitorPdpResponse) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GetDtcMonitorPdpResponse) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GetDtcMonitorPdpResponse) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *GetDtcMonitorPdpResponse) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRetryDown

`func (o *GetDtcMonitorPdpResponse) GetRetryDown() int64`

GetRetryDown returns the RetryDown field if non-nil, zero value otherwise.

### GetRetryDownOk

`func (o *GetDtcMonitorPdpResponse) GetRetryDownOk() (*int64, bool)`

GetRetryDownOk returns a tuple with the RetryDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDown

`func (o *GetDtcMonitorPdpResponse) SetRetryDown(v int64)`

SetRetryDown sets RetryDown field to given value.

### HasRetryDown

`func (o *GetDtcMonitorPdpResponse) HasRetryDown() bool`

HasRetryDown returns a boolean if a field has been set.

### GetRetryUp

`func (o *GetDtcMonitorPdpResponse) GetRetryUp() int64`

GetRetryUp returns the RetryUp field if non-nil, zero value otherwise.

### GetRetryUpOk

`func (o *GetDtcMonitorPdpResponse) GetRetryUpOk() (*int64, bool)`

GetRetryUpOk returns a tuple with the RetryUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryUp

`func (o *GetDtcMonitorPdpResponse) SetRetryUp(v int64)`

SetRetryUp sets RetryUp field to given value.

### HasRetryUp

`func (o *GetDtcMonitorPdpResponse) HasRetryUp() bool`

HasRetryUp returns a boolean if a field has been set.

### GetTimeout

`func (o *GetDtcMonitorPdpResponse) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *GetDtcMonitorPdpResponse) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *GetDtcMonitorPdpResponse) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *GetDtcMonitorPdpResponse) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetResult

`func (o *GetDtcMonitorPdpResponse) GetResult() DtcMonitorPdp`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetDtcMonitorPdpResponse) GetResultOk() (*DtcMonitorPdp, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetDtcMonitorPdpResponse) SetResult(v DtcMonitorPdp)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetDtcMonitorPdpResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


