# SyncPausePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderConfigKey** | **string** | The ID of the integration that you established within Nango. | 
**ConnectionId** | Pointer to **string** | The ID of the connection. If omitted, the syncs will be paused for all relevant connections. | [optional] 
**Syncs** | **[]string** | A list of sync names that you wish to pause. | 

## Methods

### NewSyncPausePostRequest

`func NewSyncPausePostRequest(providerConfigKey string, syncs []string, ) *SyncPausePostRequest`

NewSyncPausePostRequest instantiates a new SyncPausePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncPausePostRequestWithDefaults

`func NewSyncPausePostRequestWithDefaults() *SyncPausePostRequest`

NewSyncPausePostRequestWithDefaults instantiates a new SyncPausePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderConfigKey

`func (o *SyncPausePostRequest) GetProviderConfigKey() string`

GetProviderConfigKey returns the ProviderConfigKey field if non-nil, zero value otherwise.

### GetProviderConfigKeyOk

`func (o *SyncPausePostRequest) GetProviderConfigKeyOk() (*string, bool)`

GetProviderConfigKeyOk returns a tuple with the ProviderConfigKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderConfigKey

`func (o *SyncPausePostRequest) SetProviderConfigKey(v string)`

SetProviderConfigKey sets ProviderConfigKey field to given value.


### GetConnectionId

`func (o *SyncPausePostRequest) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *SyncPausePostRequest) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *SyncPausePostRequest) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *SyncPausePostRequest) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetSyncs

`func (o *SyncPausePostRequest) GetSyncs() []string`

GetSyncs returns the Syncs field if non-nil, zero value otherwise.

### GetSyncsOk

`func (o *SyncPausePostRequest) GetSyncsOk() (*[]string, bool)`

GetSyncsOk returns a tuple with the Syncs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncs

`func (o *SyncPausePostRequest) SetSyncs(v []string)`

SetSyncs sets Syncs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


