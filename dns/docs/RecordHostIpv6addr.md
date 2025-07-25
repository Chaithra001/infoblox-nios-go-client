# RecordHostIpv6addr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AddressType** | Pointer to **string** | Type of the DHCP IPv6 Host Address object. | [optional] 
**ConfigureForDhcp** | Pointer to **bool** | Set this to True to enable the DHCP configuration for this IPv6 host address. | [optional] 
**DiscoverNowStatus** | Pointer to **string** | The discovery status of this IPv6 Host Address. | [optional] [readonly] 
**DiscoveredData** | Pointer to [**RecordHostIpv6addrDiscoveredData**](RecordHostIpv6addrDiscoveredData.md) |  | [optional] 
**DomainName** | Pointer to **string** | Use this method to set or retrieve the domain_name value of the DHCP IPv6 Host Address object. | [optional] 
**DomainNameServers** | Pointer to **[]string** | The IPv6 addresses of DNS recursive name servers to which the DHCP client can send name resolution requests. The DHCP server includes this information in the DNS Recursive Name Server option in Advertise, Rebind, Information-Request, and Reply messages. | [optional] 
**Duid** | Pointer to **string** | DHCPv6 Unique Identifier (DUID) of the address object. | [optional] 
**Host** | Pointer to **string** | The host to which the IPv6 host address belongs, in FQDN format. It is only present when the host address object is not returned as part of a host. | [optional] [readonly] 
**Ipv6addr** | Pointer to [**RecordHostIpv6addrIpv6addr**](RecordHostIpv6addrIpv6addr.md) |  | [optional] 
**FuncCall** | Pointer to [**FuncCall**](FuncCall.md) |  | [optional] 
**Ipv6prefix** | Pointer to **string** | The IPv6 Address prefix of the DHCP IPv6 Host Address object. | [optional] 
**Ipv6prefixBits** | Pointer to **int64** | Prefix bits of the DHCP IPv6 Host Address object. | [optional] 
**LastQueried** | Pointer to **int64** | The time of the last DNS query in Epoch seconds format. | [optional] [readonly] 
**LogicFilterRules** | Pointer to [**[]RecordHostIpv6addrLogicFilterRules**](RecordHostIpv6addrLogicFilterRules.md) | This field contains the logic filters to be applied on the this host address. This list corresponds to the match rules that are written to the dhcpd configuration file. | [optional] 
**Mac** | Pointer to **string** | The MAC address for this host address. | [optional] 
**MatchClient** | Pointer to **string** | The match_client value for this fixed address. Valid values are: \&quot;DUID\&quot;: The host IP address is leased to the matching DUID. \&quot;MAC_ADDRESS\&quot;: The host IP address is leased to the matching MAC address. | [optional] 
**MsAdUserData** | Pointer to [**RecordHostIpv6addrMsAdUserData**](RecordHostIpv6addrMsAdUserData.md) |  | [optional] 
**Network** | Pointer to **string** | The network of the host address, in FQDN/CIDR format. | [optional] [readonly] 
**NetworkView** | Pointer to **string** | The name of the network view in which the host address resides. | [optional] [readonly] 
**Options** | Pointer to [**[]RecordHostIpv6addrOptions**](RecordHostIpv6addrOptions.md) | An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object. | [optional] 
**PreferredLifetime** | Pointer to **int64** | Use this method to set or retrieve the preferred lifetime value of the DHCP IPv6 Host Address object. | [optional] 
**ReservedInterface** | Pointer to **string** | The reference to the reserved interface to which the device belongs. | [optional] 
**UseDomainName** | Pointer to **bool** | Use flag for: domain_name | [optional] 
**UseDomainNameServers** | Pointer to **bool** | Use flag for: domain_name_servers | [optional] 
**UseForEaInheritance** | Pointer to **bool** | Set this to True when using this host address for EA inheritance. | [optional] 
**UseLogicFilterRules** | Pointer to **bool** | Use flag for: logic_filter_rules | [optional] 
**UseOptions** | Pointer to **bool** | Use flag for: options | [optional] 
**UsePreferredLifetime** | Pointer to **bool** | Use flag for: preferred_lifetime | [optional] 
**UseValidLifetime** | Pointer to **bool** | Use flag for: valid_lifetime | [optional] 
**ValidLifetime** | Pointer to **int64** | Use this method to set or retrieve the valid lifetime value of the DHCP IPv6 Host Address object. | [optional] 

## Methods

### NewRecordHostIpv6addr

`func NewRecordHostIpv6addr() *RecordHostIpv6addr`

NewRecordHostIpv6addr instantiates a new RecordHostIpv6addr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordHostIpv6addrWithDefaults

`func NewRecordHostIpv6addrWithDefaults() *RecordHostIpv6addr`

NewRecordHostIpv6addrWithDefaults instantiates a new RecordHostIpv6addr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *RecordHostIpv6addr) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *RecordHostIpv6addr) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *RecordHostIpv6addr) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *RecordHostIpv6addr) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAddressType

`func (o *RecordHostIpv6addr) GetAddressType() string`

GetAddressType returns the AddressType field if non-nil, zero value otherwise.

### GetAddressTypeOk

`func (o *RecordHostIpv6addr) GetAddressTypeOk() (*string, bool)`

GetAddressTypeOk returns a tuple with the AddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressType

`func (o *RecordHostIpv6addr) SetAddressType(v string)`

SetAddressType sets AddressType field to given value.

### HasAddressType

`func (o *RecordHostIpv6addr) HasAddressType() bool`

HasAddressType returns a boolean if a field has been set.

### GetConfigureForDhcp

`func (o *RecordHostIpv6addr) GetConfigureForDhcp() bool`

GetConfigureForDhcp returns the ConfigureForDhcp field if non-nil, zero value otherwise.

### GetConfigureForDhcpOk

`func (o *RecordHostIpv6addr) GetConfigureForDhcpOk() (*bool, bool)`

GetConfigureForDhcpOk returns a tuple with the ConfigureForDhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigureForDhcp

`func (o *RecordHostIpv6addr) SetConfigureForDhcp(v bool)`

SetConfigureForDhcp sets ConfigureForDhcp field to given value.

### HasConfigureForDhcp

`func (o *RecordHostIpv6addr) HasConfigureForDhcp() bool`

HasConfigureForDhcp returns a boolean if a field has been set.

### GetDiscoverNowStatus

`func (o *RecordHostIpv6addr) GetDiscoverNowStatus() string`

GetDiscoverNowStatus returns the DiscoverNowStatus field if non-nil, zero value otherwise.

### GetDiscoverNowStatusOk

`func (o *RecordHostIpv6addr) GetDiscoverNowStatusOk() (*string, bool)`

GetDiscoverNowStatusOk returns a tuple with the DiscoverNowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverNowStatus

`func (o *RecordHostIpv6addr) SetDiscoverNowStatus(v string)`

SetDiscoverNowStatus sets DiscoverNowStatus field to given value.

### HasDiscoverNowStatus

`func (o *RecordHostIpv6addr) HasDiscoverNowStatus() bool`

HasDiscoverNowStatus returns a boolean if a field has been set.

### GetDiscoveredData

`func (o *RecordHostIpv6addr) GetDiscoveredData() RecordHostIpv6addrDiscoveredData`

GetDiscoveredData returns the DiscoveredData field if non-nil, zero value otherwise.

### GetDiscoveredDataOk

`func (o *RecordHostIpv6addr) GetDiscoveredDataOk() (*RecordHostIpv6addrDiscoveredData, bool)`

GetDiscoveredDataOk returns a tuple with the DiscoveredData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredData

`func (o *RecordHostIpv6addr) SetDiscoveredData(v RecordHostIpv6addrDiscoveredData)`

SetDiscoveredData sets DiscoveredData field to given value.

### HasDiscoveredData

`func (o *RecordHostIpv6addr) HasDiscoveredData() bool`

HasDiscoveredData returns a boolean if a field has been set.

### GetDomainName

`func (o *RecordHostIpv6addr) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *RecordHostIpv6addr) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *RecordHostIpv6addr) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *RecordHostIpv6addr) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetDomainNameServers

`func (o *RecordHostIpv6addr) GetDomainNameServers() []string`

GetDomainNameServers returns the DomainNameServers field if non-nil, zero value otherwise.

### GetDomainNameServersOk

`func (o *RecordHostIpv6addr) GetDomainNameServersOk() (*[]string, bool)`

GetDomainNameServersOk returns a tuple with the DomainNameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNameServers

`func (o *RecordHostIpv6addr) SetDomainNameServers(v []string)`

SetDomainNameServers sets DomainNameServers field to given value.

### HasDomainNameServers

`func (o *RecordHostIpv6addr) HasDomainNameServers() bool`

HasDomainNameServers returns a boolean if a field has been set.

### GetDuid

`func (o *RecordHostIpv6addr) GetDuid() string`

GetDuid returns the Duid field if non-nil, zero value otherwise.

### GetDuidOk

`func (o *RecordHostIpv6addr) GetDuidOk() (*string, bool)`

GetDuidOk returns a tuple with the Duid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuid

`func (o *RecordHostIpv6addr) SetDuid(v string)`

SetDuid sets Duid field to given value.

### HasDuid

`func (o *RecordHostIpv6addr) HasDuid() bool`

HasDuid returns a boolean if a field has been set.

### GetHost

`func (o *RecordHostIpv6addr) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *RecordHostIpv6addr) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *RecordHostIpv6addr) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *RecordHostIpv6addr) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetIpv6addr

`func (o *RecordHostIpv6addr) GetIpv6addr() RecordHostIpv6addrIpv6addr`

GetIpv6addr returns the Ipv6addr field if non-nil, zero value otherwise.

### GetIpv6addrOk

`func (o *RecordHostIpv6addr) GetIpv6addrOk() (*RecordHostIpv6addrIpv6addr, bool)`

GetIpv6addrOk returns a tuple with the Ipv6addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6addr

`func (o *RecordHostIpv6addr) SetIpv6addr(v RecordHostIpv6addrIpv6addr)`

SetIpv6addr sets Ipv6addr field to given value.

### HasIpv6addr

`func (o *RecordHostIpv6addr) HasIpv6addr() bool`

HasIpv6addr returns a boolean if a field has been set.

### GetFuncCall

`func (o *RecordHostIpv6addr) GetFuncCall() FuncCall`

GetFuncCall returns the FuncCall field if non-nil, zero value otherwise.

### GetFuncCallOk

`func (o *RecordHostIpv6addr) GetFuncCallOk() (*FuncCall, bool)`

GetFuncCallOk returns a tuple with the FuncCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuncCall

`func (o *RecordHostIpv6addr) SetFuncCall(v FuncCall)`

SetFuncCall sets FuncCall field to given value.

### HasFuncCall

`func (o *RecordHostIpv6addr) HasFuncCall() bool`

HasFuncCall returns a boolean if a field has been set.

### GetIpv6prefix

`func (o *RecordHostIpv6addr) GetIpv6prefix() string`

GetIpv6prefix returns the Ipv6prefix field if non-nil, zero value otherwise.

### GetIpv6prefixOk

`func (o *RecordHostIpv6addr) GetIpv6prefixOk() (*string, bool)`

GetIpv6prefixOk returns a tuple with the Ipv6prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6prefix

`func (o *RecordHostIpv6addr) SetIpv6prefix(v string)`

SetIpv6prefix sets Ipv6prefix field to given value.

### HasIpv6prefix

`func (o *RecordHostIpv6addr) HasIpv6prefix() bool`

HasIpv6prefix returns a boolean if a field has been set.

### GetIpv6prefixBits

`func (o *RecordHostIpv6addr) GetIpv6prefixBits() int64`

GetIpv6prefixBits returns the Ipv6prefixBits field if non-nil, zero value otherwise.

### GetIpv6prefixBitsOk

`func (o *RecordHostIpv6addr) GetIpv6prefixBitsOk() (*int64, bool)`

GetIpv6prefixBitsOk returns a tuple with the Ipv6prefixBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6prefixBits

`func (o *RecordHostIpv6addr) SetIpv6prefixBits(v int64)`

SetIpv6prefixBits sets Ipv6prefixBits field to given value.

### HasIpv6prefixBits

`func (o *RecordHostIpv6addr) HasIpv6prefixBits() bool`

HasIpv6prefixBits returns a boolean if a field has been set.

### GetLastQueried

`func (o *RecordHostIpv6addr) GetLastQueried() int64`

GetLastQueried returns the LastQueried field if non-nil, zero value otherwise.

### GetLastQueriedOk

`func (o *RecordHostIpv6addr) GetLastQueriedOk() (*int64, bool)`

GetLastQueriedOk returns a tuple with the LastQueried field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQueried

`func (o *RecordHostIpv6addr) SetLastQueried(v int64)`

SetLastQueried sets LastQueried field to given value.

### HasLastQueried

`func (o *RecordHostIpv6addr) HasLastQueried() bool`

HasLastQueried returns a boolean if a field has been set.

### GetLogicFilterRules

`func (o *RecordHostIpv6addr) GetLogicFilterRules() []RecordHostIpv6addrLogicFilterRules`

GetLogicFilterRules returns the LogicFilterRules field if non-nil, zero value otherwise.

### GetLogicFilterRulesOk

`func (o *RecordHostIpv6addr) GetLogicFilterRulesOk() (*[]RecordHostIpv6addrLogicFilterRules, bool)`

GetLogicFilterRulesOk returns a tuple with the LogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicFilterRules

`func (o *RecordHostIpv6addr) SetLogicFilterRules(v []RecordHostIpv6addrLogicFilterRules)`

SetLogicFilterRules sets LogicFilterRules field to given value.

### HasLogicFilterRules

`func (o *RecordHostIpv6addr) HasLogicFilterRules() bool`

HasLogicFilterRules returns a boolean if a field has been set.

### GetMac

`func (o *RecordHostIpv6addr) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *RecordHostIpv6addr) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *RecordHostIpv6addr) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *RecordHostIpv6addr) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMatchClient

`func (o *RecordHostIpv6addr) GetMatchClient() string`

GetMatchClient returns the MatchClient field if non-nil, zero value otherwise.

### GetMatchClientOk

`func (o *RecordHostIpv6addr) GetMatchClientOk() (*string, bool)`

GetMatchClientOk returns a tuple with the MatchClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchClient

`func (o *RecordHostIpv6addr) SetMatchClient(v string)`

SetMatchClient sets MatchClient field to given value.

### HasMatchClient

`func (o *RecordHostIpv6addr) HasMatchClient() bool`

HasMatchClient returns a boolean if a field has been set.

### GetMsAdUserData

`func (o *RecordHostIpv6addr) GetMsAdUserData() RecordHostIpv6addrMsAdUserData`

GetMsAdUserData returns the MsAdUserData field if non-nil, zero value otherwise.

### GetMsAdUserDataOk

`func (o *RecordHostIpv6addr) GetMsAdUserDataOk() (*RecordHostIpv6addrMsAdUserData, bool)`

GetMsAdUserDataOk returns a tuple with the MsAdUserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAdUserData

`func (o *RecordHostIpv6addr) SetMsAdUserData(v RecordHostIpv6addrMsAdUserData)`

SetMsAdUserData sets MsAdUserData field to given value.

### HasMsAdUserData

`func (o *RecordHostIpv6addr) HasMsAdUserData() bool`

HasMsAdUserData returns a boolean if a field has been set.

### GetNetwork

`func (o *RecordHostIpv6addr) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *RecordHostIpv6addr) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *RecordHostIpv6addr) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *RecordHostIpv6addr) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNetworkView

`func (o *RecordHostIpv6addr) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *RecordHostIpv6addr) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *RecordHostIpv6addr) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *RecordHostIpv6addr) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetOptions

`func (o *RecordHostIpv6addr) GetOptions() []RecordHostIpv6addrOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *RecordHostIpv6addr) GetOptionsOk() (*[]RecordHostIpv6addrOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *RecordHostIpv6addr) SetOptions(v []RecordHostIpv6addrOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *RecordHostIpv6addr) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPreferredLifetime

`func (o *RecordHostIpv6addr) GetPreferredLifetime() int64`

GetPreferredLifetime returns the PreferredLifetime field if non-nil, zero value otherwise.

### GetPreferredLifetimeOk

`func (o *RecordHostIpv6addr) GetPreferredLifetimeOk() (*int64, bool)`

GetPreferredLifetimeOk returns a tuple with the PreferredLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLifetime

`func (o *RecordHostIpv6addr) SetPreferredLifetime(v int64)`

SetPreferredLifetime sets PreferredLifetime field to given value.

### HasPreferredLifetime

`func (o *RecordHostIpv6addr) HasPreferredLifetime() bool`

HasPreferredLifetime returns a boolean if a field has been set.

### GetReservedInterface

`func (o *RecordHostIpv6addr) GetReservedInterface() string`

GetReservedInterface returns the ReservedInterface field if non-nil, zero value otherwise.

### GetReservedInterfaceOk

`func (o *RecordHostIpv6addr) GetReservedInterfaceOk() (*string, bool)`

GetReservedInterfaceOk returns a tuple with the ReservedInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedInterface

`func (o *RecordHostIpv6addr) SetReservedInterface(v string)`

SetReservedInterface sets ReservedInterface field to given value.

### HasReservedInterface

`func (o *RecordHostIpv6addr) HasReservedInterface() bool`

HasReservedInterface returns a boolean if a field has been set.

### GetUseDomainName

`func (o *RecordHostIpv6addr) GetUseDomainName() bool`

GetUseDomainName returns the UseDomainName field if non-nil, zero value otherwise.

### GetUseDomainNameOk

`func (o *RecordHostIpv6addr) GetUseDomainNameOk() (*bool, bool)`

GetUseDomainNameOk returns a tuple with the UseDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDomainName

`func (o *RecordHostIpv6addr) SetUseDomainName(v bool)`

SetUseDomainName sets UseDomainName field to given value.

### HasUseDomainName

`func (o *RecordHostIpv6addr) HasUseDomainName() bool`

HasUseDomainName returns a boolean if a field has been set.

### GetUseDomainNameServers

`func (o *RecordHostIpv6addr) GetUseDomainNameServers() bool`

GetUseDomainNameServers returns the UseDomainNameServers field if non-nil, zero value otherwise.

### GetUseDomainNameServersOk

`func (o *RecordHostIpv6addr) GetUseDomainNameServersOk() (*bool, bool)`

GetUseDomainNameServersOk returns a tuple with the UseDomainNameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDomainNameServers

`func (o *RecordHostIpv6addr) SetUseDomainNameServers(v bool)`

SetUseDomainNameServers sets UseDomainNameServers field to given value.

### HasUseDomainNameServers

`func (o *RecordHostIpv6addr) HasUseDomainNameServers() bool`

HasUseDomainNameServers returns a boolean if a field has been set.

### GetUseForEaInheritance

`func (o *RecordHostIpv6addr) GetUseForEaInheritance() bool`

GetUseForEaInheritance returns the UseForEaInheritance field if non-nil, zero value otherwise.

### GetUseForEaInheritanceOk

`func (o *RecordHostIpv6addr) GetUseForEaInheritanceOk() (*bool, bool)`

GetUseForEaInheritanceOk returns a tuple with the UseForEaInheritance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseForEaInheritance

`func (o *RecordHostIpv6addr) SetUseForEaInheritance(v bool)`

SetUseForEaInheritance sets UseForEaInheritance field to given value.

### HasUseForEaInheritance

`func (o *RecordHostIpv6addr) HasUseForEaInheritance() bool`

HasUseForEaInheritance returns a boolean if a field has been set.

### GetUseLogicFilterRules

`func (o *RecordHostIpv6addr) GetUseLogicFilterRules() bool`

GetUseLogicFilterRules returns the UseLogicFilterRules field if non-nil, zero value otherwise.

### GetUseLogicFilterRulesOk

`func (o *RecordHostIpv6addr) GetUseLogicFilterRulesOk() (*bool, bool)`

GetUseLogicFilterRulesOk returns a tuple with the UseLogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLogicFilterRules

`func (o *RecordHostIpv6addr) SetUseLogicFilterRules(v bool)`

SetUseLogicFilterRules sets UseLogicFilterRules field to given value.

### HasUseLogicFilterRules

`func (o *RecordHostIpv6addr) HasUseLogicFilterRules() bool`

HasUseLogicFilterRules returns a boolean if a field has been set.

### GetUseOptions

`func (o *RecordHostIpv6addr) GetUseOptions() bool`

GetUseOptions returns the UseOptions field if non-nil, zero value otherwise.

### GetUseOptionsOk

`func (o *RecordHostIpv6addr) GetUseOptionsOk() (*bool, bool)`

GetUseOptionsOk returns a tuple with the UseOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOptions

`func (o *RecordHostIpv6addr) SetUseOptions(v bool)`

SetUseOptions sets UseOptions field to given value.

### HasUseOptions

`func (o *RecordHostIpv6addr) HasUseOptions() bool`

HasUseOptions returns a boolean if a field has been set.

### GetUsePreferredLifetime

`func (o *RecordHostIpv6addr) GetUsePreferredLifetime() bool`

GetUsePreferredLifetime returns the UsePreferredLifetime field if non-nil, zero value otherwise.

### GetUsePreferredLifetimeOk

`func (o *RecordHostIpv6addr) GetUsePreferredLifetimeOk() (*bool, bool)`

GetUsePreferredLifetimeOk returns a tuple with the UsePreferredLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePreferredLifetime

`func (o *RecordHostIpv6addr) SetUsePreferredLifetime(v bool)`

SetUsePreferredLifetime sets UsePreferredLifetime field to given value.

### HasUsePreferredLifetime

`func (o *RecordHostIpv6addr) HasUsePreferredLifetime() bool`

HasUsePreferredLifetime returns a boolean if a field has been set.

### GetUseValidLifetime

`func (o *RecordHostIpv6addr) GetUseValidLifetime() bool`

GetUseValidLifetime returns the UseValidLifetime field if non-nil, zero value otherwise.

### GetUseValidLifetimeOk

`func (o *RecordHostIpv6addr) GetUseValidLifetimeOk() (*bool, bool)`

GetUseValidLifetimeOk returns a tuple with the UseValidLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseValidLifetime

`func (o *RecordHostIpv6addr) SetUseValidLifetime(v bool)`

SetUseValidLifetime sets UseValidLifetime field to given value.

### HasUseValidLifetime

`func (o *RecordHostIpv6addr) HasUseValidLifetime() bool`

HasUseValidLifetime returns a boolean if a field has been set.

### GetValidLifetime

`func (o *RecordHostIpv6addr) GetValidLifetime() int64`

GetValidLifetime returns the ValidLifetime field if non-nil, zero value otherwise.

### GetValidLifetimeOk

`func (o *RecordHostIpv6addr) GetValidLifetimeOk() (*int64, bool)`

GetValidLifetimeOk returns a tuple with the ValidLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidLifetime

`func (o *RecordHostIpv6addr) SetValidLifetime(v int64)`

SetValidLifetime sets ValidLifetime field to given value.

### HasValidLifetime

`func (o *RecordHostIpv6addr) HasValidLifetime() bool`

HasValidLifetime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


