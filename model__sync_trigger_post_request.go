/*
Nango API

Nango API specs used to authorize & sync data with external APIs.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nango

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SyncTriggerPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SyncTriggerPostRequest{}

// SyncTriggerPostRequest struct for SyncTriggerPostRequest
type SyncTriggerPostRequest struct {
	// The ID of the integration that you established within Nango.
	ProviderConfigKey string `json:"provider_config_key"`
	// The ID of the connection. If omitted, the syncs will be triggered for all relevant connections.
	ConnectionId *string `json:"connection_id,omitempty"`
	// An array of sync names to trigger. If the array is empty, all syncs are triggered.
	Syncs []string `json:"syncs"`
	// Clear the records and reset the \"lastSyncDate\" associated with the sync before triggering a new sync job.
	FullResync *bool `json:"full_resync,omitempty"`
}

type _SyncTriggerPostRequest SyncTriggerPostRequest

// NewSyncTriggerPostRequest instantiates a new SyncTriggerPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyncTriggerPostRequest(providerConfigKey string, syncs []string) *SyncTriggerPostRequest {
	this := SyncTriggerPostRequest{}
	this.ProviderConfigKey = providerConfigKey
	this.Syncs = syncs
	return &this
}

// NewSyncTriggerPostRequestWithDefaults instantiates a new SyncTriggerPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyncTriggerPostRequestWithDefaults() *SyncTriggerPostRequest {
	this := SyncTriggerPostRequest{}
	return &this
}

// GetProviderConfigKey returns the ProviderConfigKey field value
func (o *SyncTriggerPostRequest) GetProviderConfigKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderConfigKey
}

// GetProviderConfigKeyOk returns a tuple with the ProviderConfigKey field value
// and a boolean to check if the value has been set.
func (o *SyncTriggerPostRequest) GetProviderConfigKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderConfigKey, true
}

// SetProviderConfigKey sets field value
func (o *SyncTriggerPostRequest) SetProviderConfigKey(v string) {
	o.ProviderConfigKey = v
}

// GetConnectionId returns the ConnectionId field value if set, zero value otherwise.
func (o *SyncTriggerPostRequest) GetConnectionId() string {
	if o == nil || IsNil(o.ConnectionId) {
		var ret string
		return ret
	}
	return *o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncTriggerPostRequest) GetConnectionIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionId) {
		return nil, false
	}
	return o.ConnectionId, true
}

// HasConnectionId returns a boolean if a field has been set.
func (o *SyncTriggerPostRequest) HasConnectionId() bool {
	if o != nil && !IsNil(o.ConnectionId) {
		return true
	}

	return false
}

// SetConnectionId gets a reference to the given string and assigns it to the ConnectionId field.
func (o *SyncTriggerPostRequest) SetConnectionId(v string) {
	o.ConnectionId = &v
}

// GetSyncs returns the Syncs field value
func (o *SyncTriggerPostRequest) GetSyncs() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Syncs
}

// GetSyncsOk returns a tuple with the Syncs field value
// and a boolean to check if the value has been set.
func (o *SyncTriggerPostRequest) GetSyncsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Syncs, true
}

// SetSyncs sets field value
func (o *SyncTriggerPostRequest) SetSyncs(v []string) {
	o.Syncs = v
}

// GetFullResync returns the FullResync field value if set, zero value otherwise.
func (o *SyncTriggerPostRequest) GetFullResync() bool {
	if o == nil || IsNil(o.FullResync) {
		var ret bool
		return ret
	}
	return *o.FullResync
}

// GetFullResyncOk returns a tuple with the FullResync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncTriggerPostRequest) GetFullResyncOk() (*bool, bool) {
	if o == nil || IsNil(o.FullResync) {
		return nil, false
	}
	return o.FullResync, true
}

// HasFullResync returns a boolean if a field has been set.
func (o *SyncTriggerPostRequest) HasFullResync() bool {
	if o != nil && !IsNil(o.FullResync) {
		return true
	}

	return false
}

// SetFullResync gets a reference to the given bool and assigns it to the FullResync field.
func (o *SyncTriggerPostRequest) SetFullResync(v bool) {
	o.FullResync = &v
}

func (o SyncTriggerPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SyncTriggerPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["provider_config_key"] = o.ProviderConfigKey
	if !IsNil(o.ConnectionId) {
		toSerialize["connection_id"] = o.ConnectionId
	}
	toSerialize["syncs"] = o.Syncs
	if !IsNil(o.FullResync) {
		toSerialize["full_resync"] = o.FullResync
	}
	return toSerialize, nil
}

func (o *SyncTriggerPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"provider_config_key",
		"syncs",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSyncTriggerPostRequest := _SyncTriggerPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSyncTriggerPostRequest)

	if err != nil {
		return err
	}

	*o = SyncTriggerPostRequest(varSyncTriggerPostRequest)

	return err
}

type NullableSyncTriggerPostRequest struct {
	value *SyncTriggerPostRequest
	isSet bool
}

func (v NullableSyncTriggerPostRequest) Get() *SyncTriggerPostRequest {
	return v.value
}

func (v *NullableSyncTriggerPostRequest) Set(val *SyncTriggerPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSyncTriggerPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSyncTriggerPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyncTriggerPostRequest(val *SyncTriggerPostRequest) *NullableSyncTriggerPostRequest {
	return &NullableSyncTriggerPostRequest{value: val, isSet: true}
}

func (v NullableSyncTriggerPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyncTriggerPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


