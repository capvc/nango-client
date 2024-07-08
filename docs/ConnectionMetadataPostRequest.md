# ConnectionMetadataPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | [**ConnectionMetadataPostRequestConnectionId**](ConnectionMetadataPostRequestConnectionId.md) |  | 
**ProviderConfigKey** | **string** | The integration ID used to create the connection (aka Unique Key). | 
**Metadata** | **map[string]interface{}** |  | 

## Methods

### NewConnectionMetadataPostRequest

`func NewConnectionMetadataPostRequest(connectionId ConnectionMetadataPostRequestConnectionId, providerConfigKey string, metadata map[string]interface{}, ) *ConnectionMetadataPostRequest`

NewConnectionMetadataPostRequest instantiates a new ConnectionMetadataPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionMetadataPostRequestWithDefaults

`func NewConnectionMetadataPostRequestWithDefaults() *ConnectionMetadataPostRequest`

NewConnectionMetadataPostRequestWithDefaults instantiates a new ConnectionMetadataPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *ConnectionMetadataPostRequest) GetConnectionId() ConnectionMetadataPostRequestConnectionId`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *ConnectionMetadataPostRequest) GetConnectionIdOk() (*ConnectionMetadataPostRequestConnectionId, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *ConnectionMetadataPostRequest) SetConnectionId(v ConnectionMetadataPostRequestConnectionId)`

SetConnectionId sets ConnectionId field to given value.


### GetProviderConfigKey

`func (o *ConnectionMetadataPostRequest) GetProviderConfigKey() string`

GetProviderConfigKey returns the ProviderConfigKey field if non-nil, zero value otherwise.

### GetProviderConfigKeyOk

`func (o *ConnectionMetadataPostRequest) GetProviderConfigKeyOk() (*string, bool)`

GetProviderConfigKeyOk returns a tuple with the ProviderConfigKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderConfigKey

`func (o *ConnectionMetadataPostRequest) SetProviderConfigKey(v string)`

SetProviderConfigKey sets ProviderConfigKey field to given value.


### GetMetadata

`func (o *ConnectionMetadataPostRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ConnectionMetadataPostRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ConnectionMetadataPostRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


