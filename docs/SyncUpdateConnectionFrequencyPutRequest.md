# SyncUpdateConnectionFrequencyPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderConfigKey** | **string** | The ID of the integration you established within Nango | 
**ConnectionId** | **string** | The ID of the connection | 
**SyncName** | **string** | The name of the sync you want to update | 
**Frequency** | **string** | The frequency you want to set (ex: &#39;every hour&#39;). Set null to revert to the default frequency. Uses the https://github.com/vercel/ms notations. Min frequency: 5 minutes. | 

## Methods

### NewSyncUpdateConnectionFrequencyPutRequest

`func NewSyncUpdateConnectionFrequencyPutRequest(providerConfigKey string, connectionId string, syncName string, frequency string, ) *SyncUpdateConnectionFrequencyPutRequest`

NewSyncUpdateConnectionFrequencyPutRequest instantiates a new SyncUpdateConnectionFrequencyPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncUpdateConnectionFrequencyPutRequestWithDefaults

`func NewSyncUpdateConnectionFrequencyPutRequestWithDefaults() *SyncUpdateConnectionFrequencyPutRequest`

NewSyncUpdateConnectionFrequencyPutRequestWithDefaults instantiates a new SyncUpdateConnectionFrequencyPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderConfigKey

`func (o *SyncUpdateConnectionFrequencyPutRequest) GetProviderConfigKey() string`

GetProviderConfigKey returns the ProviderConfigKey field if non-nil, zero value otherwise.

### GetProviderConfigKeyOk

`func (o *SyncUpdateConnectionFrequencyPutRequest) GetProviderConfigKeyOk() (*string, bool)`

GetProviderConfigKeyOk returns a tuple with the ProviderConfigKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderConfigKey

`func (o *SyncUpdateConnectionFrequencyPutRequest) SetProviderConfigKey(v string)`

SetProviderConfigKey sets ProviderConfigKey field to given value.


### GetConnectionId

`func (o *SyncUpdateConnectionFrequencyPutRequest) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *SyncUpdateConnectionFrequencyPutRequest) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *SyncUpdateConnectionFrequencyPutRequest) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetSyncName

`func (o *SyncUpdateConnectionFrequencyPutRequest) GetSyncName() string`

GetSyncName returns the SyncName field if non-nil, zero value otherwise.

### GetSyncNameOk

`func (o *SyncUpdateConnectionFrequencyPutRequest) GetSyncNameOk() (*string, bool)`

GetSyncNameOk returns a tuple with the SyncName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncName

`func (o *SyncUpdateConnectionFrequencyPutRequest) SetSyncName(v string)`

SetSyncName sets SyncName field to given value.


### GetFrequency

`func (o *SyncUpdateConnectionFrequencyPutRequest) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *SyncUpdateConnectionFrequencyPutRequest) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *SyncUpdateConnectionFrequencyPutRequest) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


