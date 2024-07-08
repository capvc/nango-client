# ActionTriggerPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionName** | **string** | The name of the action to trigger. | 
**Input** | Pointer to **map[string]interface{}** | The necessary input for your action&#39;s &#39;runAction&#39; function. | [optional] 

## Methods

### NewActionTriggerPostRequest

`func NewActionTriggerPostRequest(actionName string, ) *ActionTriggerPostRequest`

NewActionTriggerPostRequest instantiates a new ActionTriggerPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionTriggerPostRequestWithDefaults

`func NewActionTriggerPostRequestWithDefaults() *ActionTriggerPostRequest`

NewActionTriggerPostRequestWithDefaults instantiates a new ActionTriggerPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionName

`func (o *ActionTriggerPostRequest) GetActionName() string`

GetActionName returns the ActionName field if non-nil, zero value otherwise.

### GetActionNameOk

`func (o *ActionTriggerPostRequest) GetActionNameOk() (*string, bool)`

GetActionNameOk returns a tuple with the ActionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionName

`func (o *ActionTriggerPostRequest) SetActionName(v string)`

SetActionName sets ActionName field to given value.


### GetInput

`func (o *ActionTriggerPostRequest) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ActionTriggerPostRequest) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ActionTriggerPostRequest) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *ActionTriggerPostRequest) HasInput() bool`

HasInput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


