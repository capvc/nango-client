# ConfigPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderConfigKey** | **string** | A unique integration ID, which you will use in the other API calls to reference this integration. | 
**Provider** | **string** | The Nango API Configuration. | 
**OauthClientId** | Pointer to **string** | The ID of your OAuth app (registed with the external API). Required for OAuth APIs. | [optional] 
**OauthClientSecret** | Pointer to **string** | The secret of your OAuth app (registed with the external API). Required for OAuth APIs. | [optional] 
**OauthScopes** | Pointer to **string** | Comma separated list of scopes. | [optional] 

## Methods

### NewConfigPostRequest

`func NewConfigPostRequest(providerConfigKey string, provider string, ) *ConfigPostRequest`

NewConfigPostRequest instantiates a new ConfigPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigPostRequestWithDefaults

`func NewConfigPostRequestWithDefaults() *ConfigPostRequest`

NewConfigPostRequestWithDefaults instantiates a new ConfigPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderConfigKey

`func (o *ConfigPostRequest) GetProviderConfigKey() string`

GetProviderConfigKey returns the ProviderConfigKey field if non-nil, zero value otherwise.

### GetProviderConfigKeyOk

`func (o *ConfigPostRequest) GetProviderConfigKeyOk() (*string, bool)`

GetProviderConfigKeyOk returns a tuple with the ProviderConfigKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderConfigKey

`func (o *ConfigPostRequest) SetProviderConfigKey(v string)`

SetProviderConfigKey sets ProviderConfigKey field to given value.


### GetProvider

`func (o *ConfigPostRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ConfigPostRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ConfigPostRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetOauthClientId

`func (o *ConfigPostRequest) GetOauthClientId() string`

GetOauthClientId returns the OauthClientId field if non-nil, zero value otherwise.

### GetOauthClientIdOk

`func (o *ConfigPostRequest) GetOauthClientIdOk() (*string, bool)`

GetOauthClientIdOk returns a tuple with the OauthClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthClientId

`func (o *ConfigPostRequest) SetOauthClientId(v string)`

SetOauthClientId sets OauthClientId field to given value.

### HasOauthClientId

`func (o *ConfigPostRequest) HasOauthClientId() bool`

HasOauthClientId returns a boolean if a field has been set.

### GetOauthClientSecret

`func (o *ConfigPostRequest) GetOauthClientSecret() string`

GetOauthClientSecret returns the OauthClientSecret field if non-nil, zero value otherwise.

### GetOauthClientSecretOk

`func (o *ConfigPostRequest) GetOauthClientSecretOk() (*string, bool)`

GetOauthClientSecretOk returns a tuple with the OauthClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthClientSecret

`func (o *ConfigPostRequest) SetOauthClientSecret(v string)`

SetOauthClientSecret sets OauthClientSecret field to given value.

### HasOauthClientSecret

`func (o *ConfigPostRequest) HasOauthClientSecret() bool`

HasOauthClientSecret returns a boolean if a field has been set.

### GetOauthScopes

`func (o *ConfigPostRequest) GetOauthScopes() string`

GetOauthScopes returns the OauthScopes field if non-nil, zero value otherwise.

### GetOauthScopesOk

`func (o *ConfigPostRequest) GetOauthScopesOk() (*string, bool)`

GetOauthScopesOk returns a tuple with the OauthScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthScopes

`func (o *ConfigPostRequest) SetOauthScopes(v string)`

SetOauthScopes sets OauthScopes field to given value.

### HasOauthScopes

`func (o *ConfigPostRequest) HasOauthScopes() bool`

HasOauthScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


