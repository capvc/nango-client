# SyncStatusGet200ResponseSyncsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier for the sync. | [optional] 
**Name** | Pointer to **string** | The name of the sync. | [optional] 
**Status** | Pointer to **string** | The current status of the sync. | [optional] 
**Type** | Pointer to **string** | The most recent sync type completed | [optional] 
**FinishedAt** | Pointer to **string** | ISO string of the most recently completed sync | [optional] 
**NextScheduledSyncAt** | Pointer to **string** | ISO string of the next scheduled sync time | [optional] 
**Frequency** | Pointer to **string** | The execution frequency of the sync | [optional] 
**LatestResult** | Pointer to **map[string]interface{}** | Additional information regarding the latest result of the sync. Contains a model name with added, updated and deleted records | [optional] 

## Methods

### NewSyncStatusGet200ResponseSyncsInner

`func NewSyncStatusGet200ResponseSyncsInner() *SyncStatusGet200ResponseSyncsInner`

NewSyncStatusGet200ResponseSyncsInner instantiates a new SyncStatusGet200ResponseSyncsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncStatusGet200ResponseSyncsInnerWithDefaults

`func NewSyncStatusGet200ResponseSyncsInnerWithDefaults() *SyncStatusGet200ResponseSyncsInner`

NewSyncStatusGet200ResponseSyncsInnerWithDefaults instantiates a new SyncStatusGet200ResponseSyncsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SyncStatusGet200ResponseSyncsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SyncStatusGet200ResponseSyncsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SyncStatusGet200ResponseSyncsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SyncStatusGet200ResponseSyncsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SyncStatusGet200ResponseSyncsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SyncStatusGet200ResponseSyncsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SyncStatusGet200ResponseSyncsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SyncStatusGet200ResponseSyncsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *SyncStatusGet200ResponseSyncsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SyncStatusGet200ResponseSyncsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SyncStatusGet200ResponseSyncsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SyncStatusGet200ResponseSyncsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *SyncStatusGet200ResponseSyncsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SyncStatusGet200ResponseSyncsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SyncStatusGet200ResponseSyncsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SyncStatusGet200ResponseSyncsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFinishedAt

`func (o *SyncStatusGet200ResponseSyncsInner) GetFinishedAt() string`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *SyncStatusGet200ResponseSyncsInner) GetFinishedAtOk() (*string, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *SyncStatusGet200ResponseSyncsInner) SetFinishedAt(v string)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *SyncStatusGet200ResponseSyncsInner) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetNextScheduledSyncAt

`func (o *SyncStatusGet200ResponseSyncsInner) GetNextScheduledSyncAt() string`

GetNextScheduledSyncAt returns the NextScheduledSyncAt field if non-nil, zero value otherwise.

### GetNextScheduledSyncAtOk

`func (o *SyncStatusGet200ResponseSyncsInner) GetNextScheduledSyncAtOk() (*string, bool)`

GetNextScheduledSyncAtOk returns a tuple with the NextScheduledSyncAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextScheduledSyncAt

`func (o *SyncStatusGet200ResponseSyncsInner) SetNextScheduledSyncAt(v string)`

SetNextScheduledSyncAt sets NextScheduledSyncAt field to given value.

### HasNextScheduledSyncAt

`func (o *SyncStatusGet200ResponseSyncsInner) HasNextScheduledSyncAt() bool`

HasNextScheduledSyncAt returns a boolean if a field has been set.

### GetFrequency

`func (o *SyncStatusGet200ResponseSyncsInner) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *SyncStatusGet200ResponseSyncsInner) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *SyncStatusGet200ResponseSyncsInner) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *SyncStatusGet200ResponseSyncsInner) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetLatestResult

`func (o *SyncStatusGet200ResponseSyncsInner) GetLatestResult() map[string]interface{}`

GetLatestResult returns the LatestResult field if non-nil, zero value otherwise.

### GetLatestResultOk

`func (o *SyncStatusGet200ResponseSyncsInner) GetLatestResultOk() (*map[string]interface{}, bool)`

GetLatestResultOk returns a tuple with the LatestResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestResult

`func (o *SyncStatusGet200ResponseSyncsInner) SetLatestResult(v map[string]interface{})`

SetLatestResult sets LatestResult field to given value.

### HasLatestResult

`func (o *SyncStatusGet200ResponseSyncsInner) HasLatestResult() bool`

HasLatestResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


