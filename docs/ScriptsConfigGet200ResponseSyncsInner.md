# ScriptsConfigGet200ResponseSyncsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Models** | Pointer to [**[]ScriptsConfigGet200ResponseSyncsInnerModelsInner**](ScriptsConfigGet200ResponseSyncsInnerModelsInner.md) |  | [optional] 
**SyncType** | Pointer to **string** |  | [optional] 
**Runs** | Pointer to **string** |  | [optional] 
**TrackDeletes** | Pointer to **bool** |  | [optional] 
**AutoStart** | Pointer to **bool** |  | [optional] 
**LastDeployed** | Pointer to **string** |  | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**PreBuilt** | Pointer to **bool** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 
**Input** | Pointer to **map[string]interface{}** |  | [optional] 
**Returns** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Scopes** | Pointer to **[]string** |  | [optional] 
**Endpoints** | Pointer to **[]map[string]interface{}** |  | [optional] 
**NangoYamlVersion** | Pointer to **string** |  | [optional] 
**WebhookSubscriptions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewScriptsConfigGet200ResponseSyncsInner

`func NewScriptsConfigGet200ResponseSyncsInner() *ScriptsConfigGet200ResponseSyncsInner`

NewScriptsConfigGet200ResponseSyncsInner instantiates a new ScriptsConfigGet200ResponseSyncsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptsConfigGet200ResponseSyncsInnerWithDefaults

`func NewScriptsConfigGet200ResponseSyncsInnerWithDefaults() *ScriptsConfigGet200ResponseSyncsInner`

NewScriptsConfigGet200ResponseSyncsInnerWithDefaults instantiates a new ScriptsConfigGet200ResponseSyncsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ScriptsConfigGet200ResponseSyncsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScriptsConfigGet200ResponseSyncsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScriptsConfigGet200ResponseSyncsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ScriptsConfigGet200ResponseSyncsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ScriptsConfigGet200ResponseSyncsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScriptsConfigGet200ResponseSyncsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScriptsConfigGet200ResponseSyncsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ScriptsConfigGet200ResponseSyncsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetModels

`func (o *ScriptsConfigGet200ResponseSyncsInner) GetModels() []ScriptsConfigGet200ResponseSyncsInnerModelsInner`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *ScriptsConfigGet200ResponseSyncsInner) GetModelsOk() (*[]ScriptsConfigGet200ResponseSyncsInnerModelsInner, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *ScriptsConfigGet200ResponseSyncsInner) SetModels(v []ScriptsConfigGet200ResponseSyncsInnerModelsInner)`

SetModels sets Models field to given value.

### HasModels

`func (o *ScriptsConfigGet200ResponseSyncsInner) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetSyncType

`func (o *ScriptsConfigGet200ResponseSyncsInner) GetSyncType() string`

GetSyncType returns the SyncType field if non-nil, zero value otherwise.

### GetSyncTypeOk

`func (o *ScriptsConfigGet200ResponseSyncsInner) GetSyncTypeOk() (*string, bool)`

GetSyncTypeOk returns a tuple with the SyncType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncType

`func (o *ScriptsConfigGet200ResponseSyncsInner) SetSyncType(v string)`

SetSyncType sets SyncType field to given value.

### HasSyncType

`func (o *ScriptsConfigGet200ResponseSyncsInner) HasSyncType() bool`

HasSyncType returns a boolean if a field has been set.

### GetRuns

`func (o *ScriptsConfigGet200ResponseSyncsInner) GetRuns() string`

GetRuns returns the Runs field if non-nil, zero value otherwise.

### GetRunsOk

`func (o *ScriptsConfigGet200ResponseSyncsInner) GetRunsOk() (*string, bool)`

GetRunsOk returns a tuple with the Runs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuns

`func (o *ScriptsConfigGet200ResponseSyncsInner) SetRuns(v string)`

SetRuns sets Runs field to given value.

### HasRuns

`func (o *ScriptsConfigGet200ResponseSyncsInner) HasRuns() bool`

HasRuns returns a boolean if a field has been set.

### GetTrackDeletes

`func (o *ScriptsConfigGet200ResponseSyncsInner) GetTrackDeletes() bool`

GetTrackDeletes returns the TrackDeletes field if non-nil, zero value otherwise.

### GetTrackDeletesOk

`func (o *ScriptsConfigGet200ResponseSyncsInner) GetTrackDeletesOk() (*bool, bool)`

GetTrackDeletesOk returns a tuple with the TrackDeletes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackDeletes

`func (o *ScriptsConfigGet200ResponseSyncsInner) SetTrackDeletes(v bool)`

SetTrackDeletes sets TrackDeletes field to given value.

### HasTrackDeletes

`func (o *ScriptsConfigGet200ResponseSyncsInner) HasTrackDeletes() bool`

HasTrackDeletes returns a boolean if a field has been set.

### GetAutoStart

`func (o *ScriptsConfigGet200ResponseSyncsInner) GetAutoStart() bool`

GetAutoStart returns the AutoStart field if non-nil, zero value otherwise.

### GetAutoStartOk

`func (o *ScriptsConfigGet200ResponseSyncsInner) GetAutoStartOk() (*bool, bool)`

GetAutoStartOk returns a tuple with the AutoStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoStart

`func (o *ScriptsConfigGet200ResponseSyncsInner) SetAutoStart(v bool)`

SetAutoStart sets AutoStart field to given value.

### HasAutoStart

`func (o *ScriptsConfigGet200ResponseSyncsInner) HasAutoStart() bool`

HasAutoStart returns a boolean if a field has been set.

### GetLastDeployed

`func (o *ScriptsConfigGet200ResponseSyncsInner) GetLastDeployed() string`

GetLastDeployed returns the LastDeployed field if non-nil, zero value otherwise.

### GetLastDeployedOk

`func (o *ScriptsConfigGet200ResponseSyncsInner) GetLastDeployedOk() (*string, bool)`

GetLastDeployedOk returns a tuple with the LastDeployed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeployed

`func (o *ScriptsConfigGet200ResponseSyncsInner) SetLastDeployed(v string)`

SetLastDeployed sets LastDeployed field to given value.

### HasLastDeployed

`func (o *ScriptsConfigGet200ResponseSyncsInner) HasLastDeployed() bool`

HasLastDeployed returns a boolean if a field has been set.

### GetIsPublic

`func (o *ScriptsConfigGet200ResponseSyncsInner) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *ScriptsConfigGet200ResponseSyncsInner) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *ScriptsConfigGet200ResponseSyncsInner) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *ScriptsConfigGet200ResponseSyncsInner) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetPreBuilt

`func (o *ScriptsConfigGet200ResponseSyncsInner) GetPreBuilt() bool`

GetPreBuilt returns the PreBuilt field if non-nil, zero value otherwise.

### GetPreBuiltOk

`func (o *ScriptsConfigGet200ResponseSyncsInner) GetPreBuiltOk() (*bool, bool)`

GetPreBuiltOk returns a tuple with the PreBuilt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreBuilt

`func (o *ScriptsConfigGet200ResponseSyncsInner) SetPreBuilt(v bool)`

SetPreBuilt sets PreBuilt field to given value.

### HasPreBuilt

`func (o *ScriptsConfigGet200ResponseSyncsInner) HasPreBuilt() bool`

HasPreBuilt returns a boolean if a field has been set.

### GetVersion

`func (o *ScriptsConfigGet200ResponseSyncsInner) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ScriptsConfigGet200ResponseSyncsInner) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ScriptsConfigGet200ResponseSyncsInner) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ScriptsConfigGet200ResponseSyncsInner) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAttributes

`func (o *ScriptsConfigGet200ResponseSyncsInner) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ScriptsConfigGet200ResponseSyncsInner) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ScriptsConfigGet200ResponseSyncsInner) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ScriptsConfigGet200ResponseSyncsInner) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetInput

`func (o *ScriptsConfigGet200ResponseSyncsInner) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ScriptsConfigGet200ResponseSyncsInner) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ScriptsConfigGet200ResponseSyncsInner) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *ScriptsConfigGet200ResponseSyncsInner) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetReturns

`func (o *ScriptsConfigGet200ResponseSyncsInner) GetReturns() []string`

GetReturns returns the Returns field if non-nil, zero value otherwise.

### GetReturnsOk

`func (o *ScriptsConfigGet200ResponseSyncsInner) GetReturnsOk() (*[]string, bool)`

GetReturnsOk returns a tuple with the Returns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturns

`func (o *ScriptsConfigGet200ResponseSyncsInner) SetReturns(v []string)`

SetReturns sets Returns field to given value.

### HasReturns

`func (o *ScriptsConfigGet200ResponseSyncsInner) HasReturns() bool`

HasReturns returns a boolean if a field has been set.

### GetDescription

`func (o *ScriptsConfigGet200ResponseSyncsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScriptsConfigGet200ResponseSyncsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScriptsConfigGet200ResponseSyncsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScriptsConfigGet200ResponseSyncsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetScopes

`func (o *ScriptsConfigGet200ResponseSyncsInner) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ScriptsConfigGet200ResponseSyncsInner) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ScriptsConfigGet200ResponseSyncsInner) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ScriptsConfigGet200ResponseSyncsInner) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetEndpoints

`func (o *ScriptsConfigGet200ResponseSyncsInner) GetEndpoints() []map[string]interface{}`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *ScriptsConfigGet200ResponseSyncsInner) GetEndpointsOk() (*[]map[string]interface{}, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *ScriptsConfigGet200ResponseSyncsInner) SetEndpoints(v []map[string]interface{})`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *ScriptsConfigGet200ResponseSyncsInner) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetNangoYamlVersion

`func (o *ScriptsConfigGet200ResponseSyncsInner) GetNangoYamlVersion() string`

GetNangoYamlVersion returns the NangoYamlVersion field if non-nil, zero value otherwise.

### GetNangoYamlVersionOk

`func (o *ScriptsConfigGet200ResponseSyncsInner) GetNangoYamlVersionOk() (*string, bool)`

GetNangoYamlVersionOk returns a tuple with the NangoYamlVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNangoYamlVersion

`func (o *ScriptsConfigGet200ResponseSyncsInner) SetNangoYamlVersion(v string)`

SetNangoYamlVersion sets NangoYamlVersion field to given value.

### HasNangoYamlVersion

`func (o *ScriptsConfigGet200ResponseSyncsInner) HasNangoYamlVersion() bool`

HasNangoYamlVersion returns a boolean if a field has been set.

### GetWebhookSubscriptions

`func (o *ScriptsConfigGet200ResponseSyncsInner) GetWebhookSubscriptions() []string`

GetWebhookSubscriptions returns the WebhookSubscriptions field if non-nil, zero value otherwise.

### GetWebhookSubscriptionsOk

`func (o *ScriptsConfigGet200ResponseSyncsInner) GetWebhookSubscriptionsOk() (*[]string, bool)`

GetWebhookSubscriptionsOk returns a tuple with the WebhookSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookSubscriptions

`func (o *ScriptsConfigGet200ResponseSyncsInner) SetWebhookSubscriptions(v []string)`

SetWebhookSubscriptions sets WebhookSubscriptions field to given value.

### HasWebhookSubscriptions

`func (o *ScriptsConfigGet200ResponseSyncsInner) HasWebhookSubscriptions() bool`

HasWebhookSubscriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


