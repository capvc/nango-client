# ConnectionPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | **string** | The connection ID used to create the connection. | 
**ProviderConfigKey** | **string** | The integration ID that you created on Nango. | 
**AccessToken** | Pointer to **string** | (OAuth 2, required) Existing access token. | [optional] 
**RefreshToken** | Pointer to **string** | (OAuth 2, optional) Pass the refresh token if you have it. | [optional] 
**ExpiresAt** | Pointer to **string** | (OAuth 2, optional) Safer and preferred. | [optional] 
**ExpiresIn** | Pointer to **int32** | (OAuth 2, optional) In seconds. | [optional] 
**NoExpiration** | Pointer to **bool** | (OAuth2, optional) If the provider gives access tokens that don&#39;t expire, pass in &#x60;true&#x60; to avoid an import validation error. | [optional] 
**OauthClientIdOverride** | Pointer to **string** | (OAuth2, optional) Override the integration&#39;s OAuth client id | [optional] 
**OauthClientSecretOverride** | Pointer to **string** | (OAuth2, optional) Override the integration&#39;s OAuth client secret | [optional] 
**OauthToken** | Pointer to **string** | (OAuth 1, required) The client token to be attached to the connection. | [optional] 
**OauthTokenSecret** | Pointer to **string** | (OAuth 1, required) The client token secret to be attached to the connection. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | (OAuth, required for some APIs) Metadata to be attached to the connection. | [optional] 
**ConnectionConfig** | Pointer to **map[string]interface{}** | (OAuth, required for some APIs) Additional configuration to be attached to the connection. | [optional] 
**Username** | Pointer to **string** | (Basic, required) username to be attached to the connection. | [optional] 
**Password** | Pointer to **string** | (Basic, required) password to be attached to the connection. | [optional] 
**ApiKey** | Pointer to **string** | (API Key, required) API key to be attached to the connection. | [optional] 

## Methods

### NewConnectionPostRequest

`func NewConnectionPostRequest(connectionId string, providerConfigKey string, ) *ConnectionPostRequest`

NewConnectionPostRequest instantiates a new ConnectionPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionPostRequestWithDefaults

`func NewConnectionPostRequestWithDefaults() *ConnectionPostRequest`

NewConnectionPostRequestWithDefaults instantiates a new ConnectionPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *ConnectionPostRequest) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *ConnectionPostRequest) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *ConnectionPostRequest) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetProviderConfigKey

`func (o *ConnectionPostRequest) GetProviderConfigKey() string`

GetProviderConfigKey returns the ProviderConfigKey field if non-nil, zero value otherwise.

### GetProviderConfigKeyOk

`func (o *ConnectionPostRequest) GetProviderConfigKeyOk() (*string, bool)`

GetProviderConfigKeyOk returns a tuple with the ProviderConfigKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderConfigKey

`func (o *ConnectionPostRequest) SetProviderConfigKey(v string)`

SetProviderConfigKey sets ProviderConfigKey field to given value.


### GetAccessToken

`func (o *ConnectionPostRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *ConnectionPostRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *ConnectionPostRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *ConnectionPostRequest) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetRefreshToken

`func (o *ConnectionPostRequest) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *ConnectionPostRequest) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *ConnectionPostRequest) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *ConnectionPostRequest) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ConnectionPostRequest) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ConnectionPostRequest) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ConnectionPostRequest) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ConnectionPostRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetExpiresIn

`func (o *ConnectionPostRequest) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *ConnectionPostRequest) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *ConnectionPostRequest) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *ConnectionPostRequest) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetNoExpiration

`func (o *ConnectionPostRequest) GetNoExpiration() bool`

GetNoExpiration returns the NoExpiration field if non-nil, zero value otherwise.

### GetNoExpirationOk

`func (o *ConnectionPostRequest) GetNoExpirationOk() (*bool, bool)`

GetNoExpirationOk returns a tuple with the NoExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoExpiration

`func (o *ConnectionPostRequest) SetNoExpiration(v bool)`

SetNoExpiration sets NoExpiration field to given value.

### HasNoExpiration

`func (o *ConnectionPostRequest) HasNoExpiration() bool`

HasNoExpiration returns a boolean if a field has been set.

### GetOauthClientIdOverride

`func (o *ConnectionPostRequest) GetOauthClientIdOverride() string`

GetOauthClientIdOverride returns the OauthClientIdOverride field if non-nil, zero value otherwise.

### GetOauthClientIdOverrideOk

`func (o *ConnectionPostRequest) GetOauthClientIdOverrideOk() (*string, bool)`

GetOauthClientIdOverrideOk returns a tuple with the OauthClientIdOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthClientIdOverride

`func (o *ConnectionPostRequest) SetOauthClientIdOverride(v string)`

SetOauthClientIdOverride sets OauthClientIdOverride field to given value.

### HasOauthClientIdOverride

`func (o *ConnectionPostRequest) HasOauthClientIdOverride() bool`

HasOauthClientIdOverride returns a boolean if a field has been set.

### GetOauthClientSecretOverride

`func (o *ConnectionPostRequest) GetOauthClientSecretOverride() string`

GetOauthClientSecretOverride returns the OauthClientSecretOverride field if non-nil, zero value otherwise.

### GetOauthClientSecretOverrideOk

`func (o *ConnectionPostRequest) GetOauthClientSecretOverrideOk() (*string, bool)`

GetOauthClientSecretOverrideOk returns a tuple with the OauthClientSecretOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthClientSecretOverride

`func (o *ConnectionPostRequest) SetOauthClientSecretOverride(v string)`

SetOauthClientSecretOverride sets OauthClientSecretOverride field to given value.

### HasOauthClientSecretOverride

`func (o *ConnectionPostRequest) HasOauthClientSecretOverride() bool`

HasOauthClientSecretOverride returns a boolean if a field has been set.

### GetOauthToken

`func (o *ConnectionPostRequest) GetOauthToken() string`

GetOauthToken returns the OauthToken field if non-nil, zero value otherwise.

### GetOauthTokenOk

`func (o *ConnectionPostRequest) GetOauthTokenOk() (*string, bool)`

GetOauthTokenOk returns a tuple with the OauthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthToken

`func (o *ConnectionPostRequest) SetOauthToken(v string)`

SetOauthToken sets OauthToken field to given value.

### HasOauthToken

`func (o *ConnectionPostRequest) HasOauthToken() bool`

HasOauthToken returns a boolean if a field has been set.

### GetOauthTokenSecret

`func (o *ConnectionPostRequest) GetOauthTokenSecret() string`

GetOauthTokenSecret returns the OauthTokenSecret field if non-nil, zero value otherwise.

### GetOauthTokenSecretOk

`func (o *ConnectionPostRequest) GetOauthTokenSecretOk() (*string, bool)`

GetOauthTokenSecretOk returns a tuple with the OauthTokenSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthTokenSecret

`func (o *ConnectionPostRequest) SetOauthTokenSecret(v string)`

SetOauthTokenSecret sets OauthTokenSecret field to given value.

### HasOauthTokenSecret

`func (o *ConnectionPostRequest) HasOauthTokenSecret() bool`

HasOauthTokenSecret returns a boolean if a field has been set.

### GetMetadata

`func (o *ConnectionPostRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ConnectionPostRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ConnectionPostRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ConnectionPostRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetConnectionConfig

`func (o *ConnectionPostRequest) GetConnectionConfig() map[string]interface{}`

GetConnectionConfig returns the ConnectionConfig field if non-nil, zero value otherwise.

### GetConnectionConfigOk

`func (o *ConnectionPostRequest) GetConnectionConfigOk() (*map[string]interface{}, bool)`

GetConnectionConfigOk returns a tuple with the ConnectionConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionConfig

`func (o *ConnectionPostRequest) SetConnectionConfig(v map[string]interface{})`

SetConnectionConfig sets ConnectionConfig field to given value.

### HasConnectionConfig

`func (o *ConnectionPostRequest) HasConnectionConfig() bool`

HasConnectionConfig returns a boolean if a field has been set.

### GetUsername

`func (o *ConnectionPostRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ConnectionPostRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ConnectionPostRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ConnectionPostRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *ConnectionPostRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ConnectionPostRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ConnectionPostRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ConnectionPostRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetApiKey

`func (o *ConnectionPostRequest) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *ConnectionPostRequest) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *ConnectionPostRequest) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *ConnectionPostRequest) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


