# ScriptsConfigGet200ResponseActionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Models** | Pointer to [**[]ScriptsConfigGet200ResponseActionsInnerModelsInner**](ScriptsConfigGet200ResponseActionsInnerModelsInner.md) |  | [optional] 
**LastDeployed** | Pointer to **string** |  | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**PreBuilt** | Pointer to **bool** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 
**Input** | Pointer to **map[string]interface{}** |  | [optional] 
**Returns** | Pointer to **map[string]interface{}** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Scopes** | Pointer to **[]string** |  | [optional] 
**Endpoints** | Pointer to **[]map[string]interface{}** |  | [optional] 
**NangoYamlVersion** | Pointer to **string** |  | [optional] 

## Methods

### NewScriptsConfigGet200ResponseActionsInner

`func NewScriptsConfigGet200ResponseActionsInner() *ScriptsConfigGet200ResponseActionsInner`

NewScriptsConfigGet200ResponseActionsInner instantiates a new ScriptsConfigGet200ResponseActionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptsConfigGet200ResponseActionsInnerWithDefaults

`func NewScriptsConfigGet200ResponseActionsInnerWithDefaults() *ScriptsConfigGet200ResponseActionsInner`

NewScriptsConfigGet200ResponseActionsInnerWithDefaults instantiates a new ScriptsConfigGet200ResponseActionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ScriptsConfigGet200ResponseActionsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScriptsConfigGet200ResponseActionsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScriptsConfigGet200ResponseActionsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ScriptsConfigGet200ResponseActionsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ScriptsConfigGet200ResponseActionsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScriptsConfigGet200ResponseActionsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScriptsConfigGet200ResponseActionsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ScriptsConfigGet200ResponseActionsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetModels

`func (o *ScriptsConfigGet200ResponseActionsInner) GetModels() []ScriptsConfigGet200ResponseActionsInnerModelsInner`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *ScriptsConfigGet200ResponseActionsInner) GetModelsOk() (*[]ScriptsConfigGet200ResponseActionsInnerModelsInner, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *ScriptsConfigGet200ResponseActionsInner) SetModels(v []ScriptsConfigGet200ResponseActionsInnerModelsInner)`

SetModels sets Models field to given value.

### HasModels

`func (o *ScriptsConfigGet200ResponseActionsInner) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetLastDeployed

`func (o *ScriptsConfigGet200ResponseActionsInner) GetLastDeployed() string`

GetLastDeployed returns the LastDeployed field if non-nil, zero value otherwise.

### GetLastDeployedOk

`func (o *ScriptsConfigGet200ResponseActionsInner) GetLastDeployedOk() (*string, bool)`

GetLastDeployedOk returns a tuple with the LastDeployed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeployed

`func (o *ScriptsConfigGet200ResponseActionsInner) SetLastDeployed(v string)`

SetLastDeployed sets LastDeployed field to given value.

### HasLastDeployed

`func (o *ScriptsConfigGet200ResponseActionsInner) HasLastDeployed() bool`

HasLastDeployed returns a boolean if a field has been set.

### GetIsPublic

`func (o *ScriptsConfigGet200ResponseActionsInner) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *ScriptsConfigGet200ResponseActionsInner) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *ScriptsConfigGet200ResponseActionsInner) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *ScriptsConfigGet200ResponseActionsInner) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetPreBuilt

`func (o *ScriptsConfigGet200ResponseActionsInner) GetPreBuilt() bool`

GetPreBuilt returns the PreBuilt field if non-nil, zero value otherwise.

### GetPreBuiltOk

`func (o *ScriptsConfigGet200ResponseActionsInner) GetPreBuiltOk() (*bool, bool)`

GetPreBuiltOk returns a tuple with the PreBuilt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreBuilt

`func (o *ScriptsConfigGet200ResponseActionsInner) SetPreBuilt(v bool)`

SetPreBuilt sets PreBuilt field to given value.

### HasPreBuilt

`func (o *ScriptsConfigGet200ResponseActionsInner) HasPreBuilt() bool`

HasPreBuilt returns a boolean if a field has been set.

### GetVersion

`func (o *ScriptsConfigGet200ResponseActionsInner) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ScriptsConfigGet200ResponseActionsInner) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ScriptsConfigGet200ResponseActionsInner) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ScriptsConfigGet200ResponseActionsInner) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAttributes

`func (o *ScriptsConfigGet200ResponseActionsInner) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ScriptsConfigGet200ResponseActionsInner) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ScriptsConfigGet200ResponseActionsInner) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ScriptsConfigGet200ResponseActionsInner) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetInput

`func (o *ScriptsConfigGet200ResponseActionsInner) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ScriptsConfigGet200ResponseActionsInner) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ScriptsConfigGet200ResponseActionsInner) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *ScriptsConfigGet200ResponseActionsInner) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetReturns

`func (o *ScriptsConfigGet200ResponseActionsInner) GetReturns() map[string]interface{}`

GetReturns returns the Returns field if non-nil, zero value otherwise.

### GetReturnsOk

`func (o *ScriptsConfigGet200ResponseActionsInner) GetReturnsOk() (*map[string]interface{}, bool)`

GetReturnsOk returns a tuple with the Returns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturns

`func (o *ScriptsConfigGet200ResponseActionsInner) SetReturns(v map[string]interface{})`

SetReturns sets Returns field to given value.

### HasReturns

`func (o *ScriptsConfigGet200ResponseActionsInner) HasReturns() bool`

HasReturns returns a boolean if a field has been set.

### GetDescription

`func (o *ScriptsConfigGet200ResponseActionsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScriptsConfigGet200ResponseActionsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScriptsConfigGet200ResponseActionsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScriptsConfigGet200ResponseActionsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetScopes

`func (o *ScriptsConfigGet200ResponseActionsInner) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ScriptsConfigGet200ResponseActionsInner) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ScriptsConfigGet200ResponseActionsInner) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ScriptsConfigGet200ResponseActionsInner) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetEndpoints

`func (o *ScriptsConfigGet200ResponseActionsInner) GetEndpoints() []map[string]interface{}`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *ScriptsConfigGet200ResponseActionsInner) GetEndpointsOk() (*[]map[string]interface{}, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *ScriptsConfigGet200ResponseActionsInner) SetEndpoints(v []map[string]interface{})`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *ScriptsConfigGet200ResponseActionsInner) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetNangoYamlVersion

`func (o *ScriptsConfigGet200ResponseActionsInner) GetNangoYamlVersion() string`

GetNangoYamlVersion returns the NangoYamlVersion field if non-nil, zero value otherwise.

### GetNangoYamlVersionOk

`func (o *ScriptsConfigGet200ResponseActionsInner) GetNangoYamlVersionOk() (*string, bool)`

GetNangoYamlVersionOk returns a tuple with the NangoYamlVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNangoYamlVersion

`func (o *ScriptsConfigGet200ResponseActionsInner) SetNangoYamlVersion(v string)`

SetNangoYamlVersion sets NangoYamlVersion field to given value.

### HasNangoYamlVersion

`func (o *ScriptsConfigGet200ResponseActionsInner) HasNangoYamlVersion() bool`

HasNangoYamlVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


