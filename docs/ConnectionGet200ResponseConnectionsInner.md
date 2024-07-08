# ConnectionGet200ResponseConnectionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The internal Nango ID used for the connection. | 
**ConnectionId** | **string** | The connection ID used to create the connection. | 
**Provider** | **string** | The Nango API Configuration. | 
**ProviderConfigKey** | **string** | The integration ID used to create the connection (aka Unique Key, Provider Config Key). | 
**Created** | **string** | Connection creation date. | 
**Metadata** | Pointer to **map[string]interface{}** | Custom metadata attached to the connection | [optional] 

## Methods

### NewConnectionGet200ResponseConnectionsInner

`func NewConnectionGet200ResponseConnectionsInner(id int32, connectionId string, provider string, providerConfigKey string, created string, ) *ConnectionGet200ResponseConnectionsInner`

NewConnectionGet200ResponseConnectionsInner instantiates a new ConnectionGet200ResponseConnectionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionGet200ResponseConnectionsInnerWithDefaults

`func NewConnectionGet200ResponseConnectionsInnerWithDefaults() *ConnectionGet200ResponseConnectionsInner`

NewConnectionGet200ResponseConnectionsInnerWithDefaults instantiates a new ConnectionGet200ResponseConnectionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConnectionGet200ResponseConnectionsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectionGet200ResponseConnectionsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectionGet200ResponseConnectionsInner) SetId(v int32)`

SetId sets Id field to given value.


### GetConnectionId

`func (o *ConnectionGet200ResponseConnectionsInner) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *ConnectionGet200ResponseConnectionsInner) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *ConnectionGet200ResponseConnectionsInner) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetProvider

`func (o *ConnectionGet200ResponseConnectionsInner) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ConnectionGet200ResponseConnectionsInner) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ConnectionGet200ResponseConnectionsInner) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetProviderConfigKey

`func (o *ConnectionGet200ResponseConnectionsInner) GetProviderConfigKey() string`

GetProviderConfigKey returns the ProviderConfigKey field if non-nil, zero value otherwise.

### GetProviderConfigKeyOk

`func (o *ConnectionGet200ResponseConnectionsInner) GetProviderConfigKeyOk() (*string, bool)`

GetProviderConfigKeyOk returns a tuple with the ProviderConfigKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderConfigKey

`func (o *ConnectionGet200ResponseConnectionsInner) SetProviderConfigKey(v string)`

SetProviderConfigKey sets ProviderConfigKey field to given value.


### GetCreated

`func (o *ConnectionGet200ResponseConnectionsInner) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ConnectionGet200ResponseConnectionsInner) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ConnectionGet200ResponseConnectionsInner) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetMetadata

`func (o *ConnectionGet200ResponseConnectionsInner) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ConnectionGet200ResponseConnectionsInner) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ConnectionGet200ResponseConnectionsInner) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ConnectionGet200ResponseConnectionsInner) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


