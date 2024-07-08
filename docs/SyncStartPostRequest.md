# SyncStartPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderConfigKey** | **string** | The ID of the integration that you established within Nango. | 
**ConnectionId** | Pointer to **string** | The ID of the connection. If omitted, the syncs will be started for all relevant connections. | [optional] 
**Syncs** | **[]string** | A list of sync names that you wish to start. | 

## Methods

### NewSyncStartPostRequest

`func NewSyncStartPostRequest(providerConfigKey string, syncs []string, ) *SyncStartPostRequest`

NewSyncStartPostRequest instantiates a new SyncStartPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncStartPostRequestWithDefaults

`func NewSyncStartPostRequestWithDefaults() *SyncStartPostRequest`

NewSyncStartPostRequestWithDefaults instantiates a new SyncStartPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderConfigKey

`func (o *SyncStartPostRequest) GetProviderConfigKey() string`

GetProviderConfigKey returns the ProviderConfigKey field if non-nil, zero value otherwise.

### GetProviderConfigKeyOk

`func (o *SyncStartPostRequest) GetProviderConfigKeyOk() (*string, bool)`

GetProviderConfigKeyOk returns a tuple with the ProviderConfigKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderConfigKey

`func (o *SyncStartPostRequest) SetProviderConfigKey(v string)`

SetProviderConfigKey sets ProviderConfigKey field to given value.


### GetConnectionId

`func (o *SyncStartPostRequest) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *SyncStartPostRequest) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *SyncStartPostRequest) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *SyncStartPostRequest) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetSyncs

`func (o *SyncStartPostRequest) GetSyncs() []string`

GetSyncs returns the Syncs field if non-nil, zero value otherwise.

### GetSyncsOk

`func (o *SyncStartPostRequest) GetSyncsOk() (*[]string, bool)`

GetSyncsOk returns a tuple with the Syncs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncs

`func (o *SyncStartPostRequest) SetSyncs(v []string)`

SetSyncs sets Syncs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


