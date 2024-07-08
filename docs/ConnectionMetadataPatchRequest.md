# ConnectionMetadataPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | [**ConnectionMetadataPatchRequestConnectionId**](ConnectionMetadataPatchRequestConnectionId.md) |  | 
**ProviderConfigKey** | **string** | The integration ID used to create the connection (aka Unique Key). | 
**Metadata** | **map[string]interface{}** |  | 

## Methods

### NewConnectionMetadataPatchRequest

`func NewConnectionMetadataPatchRequest(connectionId ConnectionMetadataPatchRequestConnectionId, providerConfigKey string, metadata map[string]interface{}, ) *ConnectionMetadataPatchRequest`

NewConnectionMetadataPatchRequest instantiates a new ConnectionMetadataPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionMetadataPatchRequestWithDefaults

`func NewConnectionMetadataPatchRequestWithDefaults() *ConnectionMetadataPatchRequest`

NewConnectionMetadataPatchRequestWithDefaults instantiates a new ConnectionMetadataPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *ConnectionMetadataPatchRequest) GetConnectionId() ConnectionMetadataPatchRequestConnectionId`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *ConnectionMetadataPatchRequest) GetConnectionIdOk() (*ConnectionMetadataPatchRequestConnectionId, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *ConnectionMetadataPatchRequest) SetConnectionId(v ConnectionMetadataPatchRequestConnectionId)`

SetConnectionId sets ConnectionId field to given value.


### GetProviderConfigKey

`func (o *ConnectionMetadataPatchRequest) GetProviderConfigKey() string`

GetProviderConfigKey returns the ProviderConfigKey field if non-nil, zero value otherwise.

### GetProviderConfigKeyOk

`func (o *ConnectionMetadataPatchRequest) GetProviderConfigKeyOk() (*string, bool)`

GetProviderConfigKeyOk returns a tuple with the ProviderConfigKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderConfigKey

`func (o *ConnectionMetadataPatchRequest) SetProviderConfigKey(v string)`

SetProviderConfigKey sets ProviderConfigKey field to given value.


### GetMetadata

`func (o *ConnectionMetadataPatchRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ConnectionMetadataPatchRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ConnectionMetadataPatchRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


