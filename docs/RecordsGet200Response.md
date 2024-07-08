# RecordsGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Records** | Pointer to [**[]RecordsGet200ResponseRecordsInner**](RecordsGet200ResponseRecordsInner.md) |  | [optional] 
**NextCursor** | Pointer to **string** | The base64-encoded cursor for pagination | [optional] 

## Methods

### NewRecordsGet200Response

`func NewRecordsGet200Response() *RecordsGet200Response`

NewRecordsGet200Response instantiates a new RecordsGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordsGet200ResponseWithDefaults

`func NewRecordsGet200ResponseWithDefaults() *RecordsGet200Response`

NewRecordsGet200ResponseWithDefaults instantiates a new RecordsGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecords

`func (o *RecordsGet200Response) GetRecords() []RecordsGet200ResponseRecordsInner`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *RecordsGet200Response) GetRecordsOk() (*[]RecordsGet200ResponseRecordsInner, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *RecordsGet200Response) SetRecords(v []RecordsGet200ResponseRecordsInner)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *RecordsGet200Response) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### GetNextCursor

`func (o *RecordsGet200Response) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *RecordsGet200Response) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *RecordsGet200Response) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *RecordsGet200Response) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


