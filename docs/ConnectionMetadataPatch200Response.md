# ConnectionMetadataPatch200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | Pointer to [**ConnectionMetadataPatch200ResponseConnectionId**](ConnectionMetadataPatch200ResponseConnectionId.md) |  | [optional] 
**ProviderConfigKey** | Pointer to **string** | The integration ID used to create the connection (aka Unique Key). | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewConnectionMetadataPatch200Response

`func NewConnectionMetadataPatch200Response() *ConnectionMetadataPatch200Response`

NewConnectionMetadataPatch200Response instantiates a new ConnectionMetadataPatch200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionMetadataPatch200ResponseWithDefaults

`func NewConnectionMetadataPatch200ResponseWithDefaults() *ConnectionMetadataPatch200Response`

NewConnectionMetadataPatch200ResponseWithDefaults instantiates a new ConnectionMetadataPatch200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *ConnectionMetadataPatch200Response) GetConnectionId() ConnectionMetadataPatch200ResponseConnectionId`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *ConnectionMetadataPatch200Response) GetConnectionIdOk() (*ConnectionMetadataPatch200ResponseConnectionId, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *ConnectionMetadataPatch200Response) SetConnectionId(v ConnectionMetadataPatch200ResponseConnectionId)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *ConnectionMetadataPatch200Response) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetProviderConfigKey

`func (o *ConnectionMetadataPatch200Response) GetProviderConfigKey() string`

GetProviderConfigKey returns the ProviderConfigKey field if non-nil, zero value otherwise.

### GetProviderConfigKeyOk

`func (o *ConnectionMetadataPatch200Response) GetProviderConfigKeyOk() (*string, bool)`

GetProviderConfigKeyOk returns a tuple with the ProviderConfigKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderConfigKey

`func (o *ConnectionMetadataPatch200Response) SetProviderConfigKey(v string)`

SetProviderConfigKey sets ProviderConfigKey field to given value.

### HasProviderConfigKey

`func (o *ConnectionMetadataPatch200Response) HasProviderConfigKey() bool`

HasProviderConfigKey returns a boolean if a field has been set.

### GetMetadata

`func (o *ConnectionMetadataPatch200Response) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ConnectionMetadataPatch200Response) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ConnectionMetadataPatch200Response) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ConnectionMetadataPatch200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


