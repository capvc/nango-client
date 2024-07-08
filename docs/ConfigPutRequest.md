# ConfigPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderConfigKey** | **string** | The integration ID that you created on Nango. | 
**Provider** | **string** | The Nango API Configuration. | 
**OauthClientId** | **string** | The ID of your OAuth app (registed with the external API). | 
**OauthClientSecret** | **string** | The secret of your OAuth app (registed with the external API). | 
**OauthScopes** | Pointer to **string** | Comma separated list of scopes. | [optional] 

## Methods

### NewConfigPutRequest

`func NewConfigPutRequest(providerConfigKey string, provider string, oauthClientId string, oauthClientSecret string, ) *ConfigPutRequest`

NewConfigPutRequest instantiates a new ConfigPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigPutRequestWithDefaults

`func NewConfigPutRequestWithDefaults() *ConfigPutRequest`

NewConfigPutRequestWithDefaults instantiates a new ConfigPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderConfigKey

`func (o *ConfigPutRequest) GetProviderConfigKey() string`

GetProviderConfigKey returns the ProviderConfigKey field if non-nil, zero value otherwise.

### GetProviderConfigKeyOk

`func (o *ConfigPutRequest) GetProviderConfigKeyOk() (*string, bool)`

GetProviderConfigKeyOk returns a tuple with the ProviderConfigKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderConfigKey

`func (o *ConfigPutRequest) SetProviderConfigKey(v string)`

SetProviderConfigKey sets ProviderConfigKey field to given value.


### GetProvider

`func (o *ConfigPutRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ConfigPutRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ConfigPutRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetOauthClientId

`func (o *ConfigPutRequest) GetOauthClientId() string`

GetOauthClientId returns the OauthClientId field if non-nil, zero value otherwise.

### GetOauthClientIdOk

`func (o *ConfigPutRequest) GetOauthClientIdOk() (*string, bool)`

GetOauthClientIdOk returns a tuple with the OauthClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthClientId

`func (o *ConfigPutRequest) SetOauthClientId(v string)`

SetOauthClientId sets OauthClientId field to given value.


### GetOauthClientSecret

`func (o *ConfigPutRequest) GetOauthClientSecret() string`

GetOauthClientSecret returns the OauthClientSecret field if non-nil, zero value otherwise.

### GetOauthClientSecretOk

`func (o *ConfigPutRequest) GetOauthClientSecretOk() (*string, bool)`

GetOauthClientSecretOk returns a tuple with the OauthClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthClientSecret

`func (o *ConfigPutRequest) SetOauthClientSecret(v string)`

SetOauthClientSecret sets OauthClientSecret field to given value.


### GetOauthScopes

`func (o *ConfigPutRequest) GetOauthScopes() string`

GetOauthScopes returns the OauthScopes field if non-nil, zero value otherwise.

### GetOauthScopesOk

`func (o *ConfigPutRequest) GetOauthScopesOk() (*string, bool)`

GetOauthScopesOk returns a tuple with the OauthScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthScopes

`func (o *ConfigPutRequest) SetOauthScopes(v string)`

SetOauthScopes sets OauthScopes field to given value.

### HasOauthScopes

`func (o *ConfigPutRequest) HasOauthScopes() bool`

HasOauthScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


