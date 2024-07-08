# SyncTriggerPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderConfigKey** | **string** | The ID of the integration that you established within Nango. | 
**ConnectionId** | Pointer to **string** | The ID of the connection. If omitted, the syncs will be triggered for all relevant connections. | [optional] 
**Syncs** | **[]string** | An array of sync names to trigger. If the array is empty, all syncs are triggered. | 
**FullResync** | Pointer to **bool** | Clear the records and reset the \&quot;lastSyncDate\&quot; associated with the sync before triggering a new sync job. | [optional] 

## Methods

### NewSyncTriggerPostRequest

`func NewSyncTriggerPostRequest(providerConfigKey string, syncs []string, ) *SyncTriggerPostRequest`

NewSyncTriggerPostRequest instantiates a new SyncTriggerPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncTriggerPostRequestWithDefaults

`func NewSyncTriggerPostRequestWithDefaults() *SyncTriggerPostRequest`

NewSyncTriggerPostRequestWithDefaults instantiates a new SyncTriggerPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderConfigKey

`func (o *SyncTriggerPostRequest) GetProviderConfigKey() string`

GetProviderConfigKey returns the ProviderConfigKey field if non-nil, zero value otherwise.

### GetProviderConfigKeyOk

`func (o *SyncTriggerPostRequest) GetProviderConfigKeyOk() (*string, bool)`

GetProviderConfigKeyOk returns a tuple with the ProviderConfigKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderConfigKey

`func (o *SyncTriggerPostRequest) SetProviderConfigKey(v string)`

SetProviderConfigKey sets ProviderConfigKey field to given value.


### GetConnectionId

`func (o *SyncTriggerPostRequest) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *SyncTriggerPostRequest) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *SyncTriggerPostRequest) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *SyncTriggerPostRequest) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetSyncs

`func (o *SyncTriggerPostRequest) GetSyncs() []string`

GetSyncs returns the Syncs field if non-nil, zero value otherwise.

### GetSyncsOk

`func (o *SyncTriggerPostRequest) GetSyncsOk() (*[]string, bool)`

GetSyncsOk returns a tuple with the Syncs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncs

`func (o *SyncTriggerPostRequest) SetSyncs(v []string)`

SetSyncs sets Syncs field to given value.


### GetFullResync

`func (o *SyncTriggerPostRequest) GetFullResync() bool`

GetFullResync returns the FullResync field if non-nil, zero value otherwise.

### GetFullResyncOk

`func (o *SyncTriggerPostRequest) GetFullResyncOk() (*bool, bool)`

GetFullResyncOk returns a tuple with the FullResync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullResync

`func (o *SyncTriggerPostRequest) SetFullResync(v bool)`

SetFullResync sets FullResync field to given value.

### HasFullResync

`func (o *SyncTriggerPostRequest) HasFullResync() bool`

HasFullResync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


