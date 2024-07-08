# ScriptsConfigGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderConfigKey** | Pointer to **string** |  | [optional] 
**Syncs** | Pointer to [**[]ScriptsConfigGet200ResponseSyncsInner**](ScriptsConfigGet200ResponseSyncsInner.md) |  | [optional] 
**Actions** | Pointer to [**[]ScriptsConfigGet200ResponseActionsInner**](ScriptsConfigGet200ResponseActionsInner.md) |  | [optional] 
**PostConnectionScripts** | Pointer to **[]string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 

## Methods

### NewScriptsConfigGet200Response

`func NewScriptsConfigGet200Response() *ScriptsConfigGet200Response`

NewScriptsConfigGet200Response instantiates a new ScriptsConfigGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptsConfigGet200ResponseWithDefaults

`func NewScriptsConfigGet200ResponseWithDefaults() *ScriptsConfigGet200Response`

NewScriptsConfigGet200ResponseWithDefaults instantiates a new ScriptsConfigGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderConfigKey

`func (o *ScriptsConfigGet200Response) GetProviderConfigKey() string`

GetProviderConfigKey returns the ProviderConfigKey field if non-nil, zero value otherwise.

### GetProviderConfigKeyOk

`func (o *ScriptsConfigGet200Response) GetProviderConfigKeyOk() (*string, bool)`

GetProviderConfigKeyOk returns a tuple with the ProviderConfigKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderConfigKey

`func (o *ScriptsConfigGet200Response) SetProviderConfigKey(v string)`

SetProviderConfigKey sets ProviderConfigKey field to given value.

### HasProviderConfigKey

`func (o *ScriptsConfigGet200Response) HasProviderConfigKey() bool`

HasProviderConfigKey returns a boolean if a field has been set.

### GetSyncs

`func (o *ScriptsConfigGet200Response) GetSyncs() []ScriptsConfigGet200ResponseSyncsInner`

GetSyncs returns the Syncs field if non-nil, zero value otherwise.

### GetSyncsOk

`func (o *ScriptsConfigGet200Response) GetSyncsOk() (*[]ScriptsConfigGet200ResponseSyncsInner, bool)`

GetSyncsOk returns a tuple with the Syncs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncs

`func (o *ScriptsConfigGet200Response) SetSyncs(v []ScriptsConfigGet200ResponseSyncsInner)`

SetSyncs sets Syncs field to given value.

### HasSyncs

`func (o *ScriptsConfigGet200Response) HasSyncs() bool`

HasSyncs returns a boolean if a field has been set.

### GetActions

`func (o *ScriptsConfigGet200Response) GetActions() []ScriptsConfigGet200ResponseActionsInner`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ScriptsConfigGet200Response) GetActionsOk() (*[]ScriptsConfigGet200ResponseActionsInner, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ScriptsConfigGet200Response) SetActions(v []ScriptsConfigGet200ResponseActionsInner)`

SetActions sets Actions field to given value.

### HasActions

`func (o *ScriptsConfigGet200Response) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetPostConnectionScripts

`func (o *ScriptsConfigGet200Response) GetPostConnectionScripts() []string`

GetPostConnectionScripts returns the PostConnectionScripts field if non-nil, zero value otherwise.

### GetPostConnectionScriptsOk

`func (o *ScriptsConfigGet200Response) GetPostConnectionScriptsOk() (*[]string, bool)`

GetPostConnectionScriptsOk returns a tuple with the PostConnectionScripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostConnectionScripts

`func (o *ScriptsConfigGet200Response) SetPostConnectionScripts(v []string)`

SetPostConnectionScripts sets PostConnectionScripts field to given value.

### HasPostConnectionScripts

`func (o *ScriptsConfigGet200Response) HasPostConnectionScripts() bool`

HasPostConnectionScripts returns a boolean if a field has been set.

### GetProvider

`func (o *ScriptsConfigGet200Response) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ScriptsConfigGet200Response) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ScriptsConfigGet200Response) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ScriptsConfigGet200Response) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


