# Ipv6networkDiscoveryBasicPollSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortScanning** | Pointer to **bool** | Determines whether port scanning is enabled or not. | [optional] 
**DeviceProfile** | Pointer to **bool** | Determines whether device profile is enabled or not. | [optional] 
**SnmpCollection** | Pointer to **bool** | Determines whether SNMP collection is enabled or not. | [optional] 
**CliCollection** | Pointer to **bool** | Determines whether CLI collection is enabled or not. | [optional] 
**NetbiosScanning** | Pointer to **bool** | Determines whether netbios scanning is enabled or not. | [optional] 
**CompletePingSweep** | Pointer to **bool** | Determines whether complete ping sweep is enabled or not. | [optional] 
**SmartSubnetPingSweep** | Pointer to **bool** | Determines whether smart subnet ping sweep is enabled or not. | [optional] 
**AutoArpRefreshBeforeSwitchPortPolling** | Pointer to **bool** | Determines whether auto ARP refresh before switch port polling is enabled or not. | [optional] 
**SwitchPortDataCollectionPolling** | Pointer to **string** | A switch port data collection polling mode. | [optional] 
**SwitchPortDataCollectionPollingSchedule** | Pointer to [**Ipv6networkdiscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule**](Ipv6networkdiscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule.md) |  | [optional] 
**SwitchPortDataCollectionPollingInterval** | Pointer to **int64** | Indicates the interval for switch port data collection polling. | [optional] 
**CredentialGroup** | Pointer to **string** | Credential group. | [optional] 
**PollingFrequencyModifier** | Pointer to **string** | Polling Frequency Modifier. | [optional] 
**UseGlobalPollingFrequencyModifier** | Pointer to **bool** | Use Global Polling Frequency Modifier. | [optional] 

## Methods

### NewIpv6networkDiscoveryBasicPollSettings

`func NewIpv6networkDiscoveryBasicPollSettings() *Ipv6networkDiscoveryBasicPollSettings`

NewIpv6networkDiscoveryBasicPollSettings instantiates a new Ipv6networkDiscoveryBasicPollSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpv6networkDiscoveryBasicPollSettingsWithDefaults

`func NewIpv6networkDiscoveryBasicPollSettingsWithDefaults() *Ipv6networkDiscoveryBasicPollSettings`

NewIpv6networkDiscoveryBasicPollSettingsWithDefaults instantiates a new Ipv6networkDiscoveryBasicPollSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortScanning

`func (o *Ipv6networkDiscoveryBasicPollSettings) GetPortScanning() bool`

GetPortScanning returns the PortScanning field if non-nil, zero value otherwise.

### GetPortScanningOk

`func (o *Ipv6networkDiscoveryBasicPollSettings) GetPortScanningOk() (*bool, bool)`

GetPortScanningOk returns a tuple with the PortScanning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortScanning

`func (o *Ipv6networkDiscoveryBasicPollSettings) SetPortScanning(v bool)`

SetPortScanning sets PortScanning field to given value.

### HasPortScanning

`func (o *Ipv6networkDiscoveryBasicPollSettings) HasPortScanning() bool`

HasPortScanning returns a boolean if a field has been set.

### GetDeviceProfile

`func (o *Ipv6networkDiscoveryBasicPollSettings) GetDeviceProfile() bool`

GetDeviceProfile returns the DeviceProfile field if non-nil, zero value otherwise.

### GetDeviceProfileOk

`func (o *Ipv6networkDiscoveryBasicPollSettings) GetDeviceProfileOk() (*bool, bool)`

GetDeviceProfileOk returns a tuple with the DeviceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceProfile

`func (o *Ipv6networkDiscoveryBasicPollSettings) SetDeviceProfile(v bool)`

SetDeviceProfile sets DeviceProfile field to given value.

### HasDeviceProfile

`func (o *Ipv6networkDiscoveryBasicPollSettings) HasDeviceProfile() bool`

HasDeviceProfile returns a boolean if a field has been set.

### GetSnmpCollection

`func (o *Ipv6networkDiscoveryBasicPollSettings) GetSnmpCollection() bool`

GetSnmpCollection returns the SnmpCollection field if non-nil, zero value otherwise.

### GetSnmpCollectionOk

`func (o *Ipv6networkDiscoveryBasicPollSettings) GetSnmpCollectionOk() (*bool, bool)`

GetSnmpCollectionOk returns a tuple with the SnmpCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpCollection

`func (o *Ipv6networkDiscoveryBasicPollSettings) SetSnmpCollection(v bool)`

SetSnmpCollection sets SnmpCollection field to given value.

### HasSnmpCollection

`func (o *Ipv6networkDiscoveryBasicPollSettings) HasSnmpCollection() bool`

HasSnmpCollection returns a boolean if a field has been set.

### GetCliCollection

`func (o *Ipv6networkDiscoveryBasicPollSettings) GetCliCollection() bool`

GetCliCollection returns the CliCollection field if non-nil, zero value otherwise.

### GetCliCollectionOk

`func (o *Ipv6networkDiscoveryBasicPollSettings) GetCliCollectionOk() (*bool, bool)`

GetCliCollectionOk returns a tuple with the CliCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliCollection

`func (o *Ipv6networkDiscoveryBasicPollSettings) SetCliCollection(v bool)`

SetCliCollection sets CliCollection field to given value.

### HasCliCollection

`func (o *Ipv6networkDiscoveryBasicPollSettings) HasCliCollection() bool`

HasCliCollection returns a boolean if a field has been set.

### GetNetbiosScanning

`func (o *Ipv6networkDiscoveryBasicPollSettings) GetNetbiosScanning() bool`

GetNetbiosScanning returns the NetbiosScanning field if non-nil, zero value otherwise.

### GetNetbiosScanningOk

`func (o *Ipv6networkDiscoveryBasicPollSettings) GetNetbiosScanningOk() (*bool, bool)`

GetNetbiosScanningOk returns a tuple with the NetbiosScanning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetbiosScanning

`func (o *Ipv6networkDiscoveryBasicPollSettings) SetNetbiosScanning(v bool)`

SetNetbiosScanning sets NetbiosScanning field to given value.

### HasNetbiosScanning

`func (o *Ipv6networkDiscoveryBasicPollSettings) HasNetbiosScanning() bool`

HasNetbiosScanning returns a boolean if a field has been set.

### GetCompletePingSweep

`func (o *Ipv6networkDiscoveryBasicPollSettings) GetCompletePingSweep() bool`

GetCompletePingSweep returns the CompletePingSweep field if non-nil, zero value otherwise.

### GetCompletePingSweepOk

`func (o *Ipv6networkDiscoveryBasicPollSettings) GetCompletePingSweepOk() (*bool, bool)`

GetCompletePingSweepOk returns a tuple with the CompletePingSweep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletePingSweep

`func (o *Ipv6networkDiscoveryBasicPollSettings) SetCompletePingSweep(v bool)`

SetCompletePingSweep sets CompletePingSweep field to given value.

### HasCompletePingSweep

`func (o *Ipv6networkDiscoveryBasicPollSettings) HasCompletePingSweep() bool`

HasCompletePingSweep returns a boolean if a field has been set.

### GetSmartSubnetPingSweep

`func (o *Ipv6networkDiscoveryBasicPollSettings) GetSmartSubnetPingSweep() bool`

GetSmartSubnetPingSweep returns the SmartSubnetPingSweep field if non-nil, zero value otherwise.

### GetSmartSubnetPingSweepOk

`func (o *Ipv6networkDiscoveryBasicPollSettings) GetSmartSubnetPingSweepOk() (*bool, bool)`

GetSmartSubnetPingSweepOk returns a tuple with the SmartSubnetPingSweep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartSubnetPingSweep

`func (o *Ipv6networkDiscoveryBasicPollSettings) SetSmartSubnetPingSweep(v bool)`

SetSmartSubnetPingSweep sets SmartSubnetPingSweep field to given value.

### HasSmartSubnetPingSweep

`func (o *Ipv6networkDiscoveryBasicPollSettings) HasSmartSubnetPingSweep() bool`

HasSmartSubnetPingSweep returns a boolean if a field has been set.

### GetAutoArpRefreshBeforeSwitchPortPolling

`func (o *Ipv6networkDiscoveryBasicPollSettings) GetAutoArpRefreshBeforeSwitchPortPolling() bool`

GetAutoArpRefreshBeforeSwitchPortPolling returns the AutoArpRefreshBeforeSwitchPortPolling field if non-nil, zero value otherwise.

### GetAutoArpRefreshBeforeSwitchPortPollingOk

`func (o *Ipv6networkDiscoveryBasicPollSettings) GetAutoArpRefreshBeforeSwitchPortPollingOk() (*bool, bool)`

GetAutoArpRefreshBeforeSwitchPortPollingOk returns a tuple with the AutoArpRefreshBeforeSwitchPortPolling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoArpRefreshBeforeSwitchPortPolling

`func (o *Ipv6networkDiscoveryBasicPollSettings) SetAutoArpRefreshBeforeSwitchPortPolling(v bool)`

SetAutoArpRefreshBeforeSwitchPortPolling sets AutoArpRefreshBeforeSwitchPortPolling field to given value.

### HasAutoArpRefreshBeforeSwitchPortPolling

`func (o *Ipv6networkDiscoveryBasicPollSettings) HasAutoArpRefreshBeforeSwitchPortPolling() bool`

HasAutoArpRefreshBeforeSwitchPortPolling returns a boolean if a field has been set.

### GetSwitchPortDataCollectionPolling

`func (o *Ipv6networkDiscoveryBasicPollSettings) GetSwitchPortDataCollectionPolling() string`

GetSwitchPortDataCollectionPolling returns the SwitchPortDataCollectionPolling field if non-nil, zero value otherwise.

### GetSwitchPortDataCollectionPollingOk

`func (o *Ipv6networkDiscoveryBasicPollSettings) GetSwitchPortDataCollectionPollingOk() (*string, bool)`

GetSwitchPortDataCollectionPollingOk returns a tuple with the SwitchPortDataCollectionPolling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchPortDataCollectionPolling

`func (o *Ipv6networkDiscoveryBasicPollSettings) SetSwitchPortDataCollectionPolling(v string)`

SetSwitchPortDataCollectionPolling sets SwitchPortDataCollectionPolling field to given value.

### HasSwitchPortDataCollectionPolling

`func (o *Ipv6networkDiscoveryBasicPollSettings) HasSwitchPortDataCollectionPolling() bool`

HasSwitchPortDataCollectionPolling returns a boolean if a field has been set.

### GetSwitchPortDataCollectionPollingSchedule

`func (o *Ipv6networkDiscoveryBasicPollSettings) GetSwitchPortDataCollectionPollingSchedule() Ipv6networkdiscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule`

GetSwitchPortDataCollectionPollingSchedule returns the SwitchPortDataCollectionPollingSchedule field if non-nil, zero value otherwise.

### GetSwitchPortDataCollectionPollingScheduleOk

`func (o *Ipv6networkDiscoveryBasicPollSettings) GetSwitchPortDataCollectionPollingScheduleOk() (*Ipv6networkdiscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule, bool)`

GetSwitchPortDataCollectionPollingScheduleOk returns a tuple with the SwitchPortDataCollectionPollingSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchPortDataCollectionPollingSchedule

`func (o *Ipv6networkDiscoveryBasicPollSettings) SetSwitchPortDataCollectionPollingSchedule(v Ipv6networkdiscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule)`

SetSwitchPortDataCollectionPollingSchedule sets SwitchPortDataCollectionPollingSchedule field to given value.

### HasSwitchPortDataCollectionPollingSchedule

`func (o *Ipv6networkDiscoveryBasicPollSettings) HasSwitchPortDataCollectionPollingSchedule() bool`

HasSwitchPortDataCollectionPollingSchedule returns a boolean if a field has been set.

### GetSwitchPortDataCollectionPollingInterval

`func (o *Ipv6networkDiscoveryBasicPollSettings) GetSwitchPortDataCollectionPollingInterval() int64`

GetSwitchPortDataCollectionPollingInterval returns the SwitchPortDataCollectionPollingInterval field if non-nil, zero value otherwise.

### GetSwitchPortDataCollectionPollingIntervalOk

`func (o *Ipv6networkDiscoveryBasicPollSettings) GetSwitchPortDataCollectionPollingIntervalOk() (*int64, bool)`

GetSwitchPortDataCollectionPollingIntervalOk returns a tuple with the SwitchPortDataCollectionPollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchPortDataCollectionPollingInterval

`func (o *Ipv6networkDiscoveryBasicPollSettings) SetSwitchPortDataCollectionPollingInterval(v int64)`

SetSwitchPortDataCollectionPollingInterval sets SwitchPortDataCollectionPollingInterval field to given value.

### HasSwitchPortDataCollectionPollingInterval

`func (o *Ipv6networkDiscoveryBasicPollSettings) HasSwitchPortDataCollectionPollingInterval() bool`

HasSwitchPortDataCollectionPollingInterval returns a boolean if a field has been set.

### GetCredentialGroup

`func (o *Ipv6networkDiscoveryBasicPollSettings) GetCredentialGroup() string`

GetCredentialGroup returns the CredentialGroup field if non-nil, zero value otherwise.

### GetCredentialGroupOk

`func (o *Ipv6networkDiscoveryBasicPollSettings) GetCredentialGroupOk() (*string, bool)`

GetCredentialGroupOk returns a tuple with the CredentialGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialGroup

`func (o *Ipv6networkDiscoveryBasicPollSettings) SetCredentialGroup(v string)`

SetCredentialGroup sets CredentialGroup field to given value.

### HasCredentialGroup

`func (o *Ipv6networkDiscoveryBasicPollSettings) HasCredentialGroup() bool`

HasCredentialGroup returns a boolean if a field has been set.

### GetPollingFrequencyModifier

`func (o *Ipv6networkDiscoveryBasicPollSettings) GetPollingFrequencyModifier() string`

GetPollingFrequencyModifier returns the PollingFrequencyModifier field if non-nil, zero value otherwise.

### GetPollingFrequencyModifierOk

`func (o *Ipv6networkDiscoveryBasicPollSettings) GetPollingFrequencyModifierOk() (*string, bool)`

GetPollingFrequencyModifierOk returns a tuple with the PollingFrequencyModifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingFrequencyModifier

`func (o *Ipv6networkDiscoveryBasicPollSettings) SetPollingFrequencyModifier(v string)`

SetPollingFrequencyModifier sets PollingFrequencyModifier field to given value.

### HasPollingFrequencyModifier

`func (o *Ipv6networkDiscoveryBasicPollSettings) HasPollingFrequencyModifier() bool`

HasPollingFrequencyModifier returns a boolean if a field has been set.

### GetUseGlobalPollingFrequencyModifier

`func (o *Ipv6networkDiscoveryBasicPollSettings) GetUseGlobalPollingFrequencyModifier() bool`

GetUseGlobalPollingFrequencyModifier returns the UseGlobalPollingFrequencyModifier field if non-nil, zero value otherwise.

### GetUseGlobalPollingFrequencyModifierOk

`func (o *Ipv6networkDiscoveryBasicPollSettings) GetUseGlobalPollingFrequencyModifierOk() (*bool, bool)`

GetUseGlobalPollingFrequencyModifierOk returns a tuple with the UseGlobalPollingFrequencyModifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGlobalPollingFrequencyModifier

`func (o *Ipv6networkDiscoveryBasicPollSettings) SetUseGlobalPollingFrequencyModifier(v bool)`

SetUseGlobalPollingFrequencyModifier sets UseGlobalPollingFrequencyModifier field to given value.

### HasUseGlobalPollingFrequencyModifier

`func (o *Ipv6networkDiscoveryBasicPollSettings) HasUseGlobalPollingFrequencyModifier() bool`

HasUseGlobalPollingFrequencyModifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


