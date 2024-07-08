# ConfigProviderConfigKeyGet200ResponseConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UniqueKey** | **string** | The integration ID that you created in Nango. | 
**Provider** | **string** | The Nango API Configuration. | 
**Actions** | Pointer to [**[]ConfigProviderConfigKeyGet200ResponseConfigActionsInner**](ConfigProviderConfigKeyGet200ResponseConfigActionsInner.md) |  | [optional] 
**Syncs** | Pointer to [**[]ConfigProviderConfigKeyGet200ResponseConfigSyncsInner**](ConfigProviderConfigKeyGet200ResponseConfigSyncsInner.md) |  | [optional] 

## Methods

### NewConfigProviderConfigKeyGet200ResponseConfig

`func NewConfigProviderConfigKeyGet200ResponseConfig(uniqueKey string, provider string, ) *ConfigProviderConfigKeyGet200ResponseConfig`

NewConfigProviderConfigKeyGet200ResponseConfig instantiates a new ConfigProviderConfigKeyGet200ResponseConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigProviderConfigKeyGet200ResponseConfigWithDefaults

`func NewConfigProviderConfigKeyGet200ResponseConfigWithDefaults() *ConfigProviderConfigKeyGet200ResponseConfig`

NewConfigProviderConfigKeyGet200ResponseConfigWithDefaults instantiates a new ConfigProviderConfigKeyGet200ResponseConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUniqueKey

`func (o *ConfigProviderConfigKeyGet200ResponseConfig) GetUniqueKey() string`

GetUniqueKey returns the UniqueKey field if non-nil, zero value otherwise.

### GetUniqueKeyOk

`func (o *ConfigProviderConfigKeyGet200ResponseConfig) GetUniqueKeyOk() (*string, bool)`

GetUniqueKeyOk returns a tuple with the UniqueKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueKey

`func (o *ConfigProviderConfigKeyGet200ResponseConfig) SetUniqueKey(v string)`

SetUniqueKey sets UniqueKey field to given value.


### GetProvider

`func (o *ConfigProviderConfigKeyGet200ResponseConfig) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ConfigProviderConfigKeyGet200ResponseConfig) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ConfigProviderConfigKeyGet200ResponseConfig) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetActions

`func (o *ConfigProviderConfigKeyGet200ResponseConfig) GetActions() []ConfigProviderConfigKeyGet200ResponseConfigActionsInner`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ConfigProviderConfigKeyGet200ResponseConfig) GetActionsOk() (*[]ConfigProviderConfigKeyGet200ResponseConfigActionsInner, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ConfigProviderConfigKeyGet200ResponseConfig) SetActions(v []ConfigProviderConfigKeyGet200ResponseConfigActionsInner)`

SetActions sets Actions field to given value.

### HasActions

`func (o *ConfigProviderConfigKeyGet200ResponseConfig) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetSyncs

`func (o *ConfigProviderConfigKeyGet200ResponseConfig) GetSyncs() []ConfigProviderConfigKeyGet200ResponseConfigSyncsInner`

GetSyncs returns the Syncs field if non-nil, zero value otherwise.

### GetSyncsOk

`func (o *ConfigProviderConfigKeyGet200ResponseConfig) GetSyncsOk() (*[]ConfigProviderConfigKeyGet200ResponseConfigSyncsInner, bool)`

GetSyncsOk returns a tuple with the Syncs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncs

`func (o *ConfigProviderConfigKeyGet200ResponseConfig) SetSyncs(v []ConfigProviderConfigKeyGet200ResponseConfigSyncsInner)`

SetSyncs sets Syncs field to given value.

### HasSyncs

`func (o *ConfigProviderConfigKeyGet200ResponseConfig) HasSyncs() bool`

HasSyncs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


