# ConnectionMetadataPost201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | Pointer to [**ConnectionMetadataPost201ResponseConnectionId**](ConnectionMetadataPost201ResponseConnectionId.md) |  | [optional] 
**ProviderConfigKey** | Pointer to **string** | The integration ID used to create the connection (aka Unique Key). | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewConnectionMetadataPost201Response

`func NewConnectionMetadataPost201Response() *ConnectionMetadataPost201Response`

NewConnectionMetadataPost201Response instantiates a new ConnectionMetadataPost201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionMetadataPost201ResponseWithDefaults

`func NewConnectionMetadataPost201ResponseWithDefaults() *ConnectionMetadataPost201Response`

NewConnectionMetadataPost201ResponseWithDefaults instantiates a new ConnectionMetadataPost201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *ConnectionMetadataPost201Response) GetConnectionId() ConnectionMetadataPost201ResponseConnectionId`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *ConnectionMetadataPost201Response) GetConnectionIdOk() (*ConnectionMetadataPost201ResponseConnectionId, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *ConnectionMetadataPost201Response) SetConnectionId(v ConnectionMetadataPost201ResponseConnectionId)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *ConnectionMetadataPost201Response) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetProviderConfigKey

`func (o *ConnectionMetadataPost201Response) GetProviderConfigKey() string`

GetProviderConfigKey returns the ProviderConfigKey field if non-nil, zero value otherwise.

### GetProviderConfigKeyOk

`func (o *ConnectionMetadataPost201Response) GetProviderConfigKeyOk() (*string, bool)`

GetProviderConfigKeyOk returns a tuple with the ProviderConfigKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderConfigKey

`func (o *ConnectionMetadataPost201Response) SetProviderConfigKey(v string)`

SetProviderConfigKey sets ProviderConfigKey field to given value.

### HasProviderConfigKey

`func (o *ConnectionMetadataPost201Response) HasProviderConfigKey() bool`

HasProviderConfigKey returns a boolean if a field has been set.

### GetMetadata

`func (o *ConnectionMetadataPost201Response) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ConnectionMetadataPost201Response) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ConnectionMetadataPost201Response) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ConnectionMetadataPost201Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


