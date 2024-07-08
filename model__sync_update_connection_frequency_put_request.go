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

// checks if the SyncUpdateConnectionFrequencyPutRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SyncUpdateConnectionFrequencyPutRequest{}

// SyncUpdateConnectionFrequencyPutRequest struct for SyncUpdateConnectionFrequencyPutRequest
type SyncUpdateConnectionFrequencyPutRequest struct {
	// The ID of the integration you established within Nango
	ProviderConfigKey string `json:"provider_config_key"`
	// The ID of the connection
	ConnectionId string `json:"connection_id"`
	// The name of the sync you want to update
	SyncName string `json:"sync_name"`
	// The frequency you want to set (ex: 'every hour'). Set null to revert to the default frequency. Uses the https://github.com/vercel/ms notations. Min frequency: 5 minutes.
	Frequency string `json:"frequency"`
}

type _SyncUpdateConnectionFrequencyPutRequest SyncUpdateConnectionFrequencyPutRequest

// NewSyncUpdateConnectionFrequencyPutRequest instantiates a new SyncUpdateConnectionFrequencyPutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyncUpdateConnectionFrequencyPutRequest(providerConfigKey string, connectionId string, syncName string, frequency string) *SyncUpdateConnectionFrequencyPutRequest {
	this := SyncUpdateConnectionFrequencyPutRequest{}
	this.ProviderConfigKey = providerConfigKey
	this.ConnectionId = connectionId
	this.SyncName = syncName
	this.Frequency = frequency
	return &this
}

// NewSyncUpdateConnectionFrequencyPutRequestWithDefaults instantiates a new SyncUpdateConnectionFrequencyPutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyncUpdateConnectionFrequencyPutRequestWithDefaults() *SyncUpdateConnectionFrequencyPutRequest {
	this := SyncUpdateConnectionFrequencyPutRequest{}
	return &this
}

// GetProviderConfigKey returns the ProviderConfigKey field value
func (o *SyncUpdateConnectionFrequencyPutRequest) GetProviderConfigKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderConfigKey
}

// GetProviderConfigKeyOk returns a tuple with the ProviderConfigKey field value
// and a boolean to check if the value has been set.
func (o *SyncUpdateConnectionFrequencyPutRequest) GetProviderConfigKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderConfigKey, true
}

// SetProviderConfigKey sets field value
func (o *SyncUpdateConnectionFrequencyPutRequest) SetProviderConfigKey(v string) {
	o.ProviderConfigKey = v
}

// GetConnectionId returns the ConnectionId field value
func (o *SyncUpdateConnectionFrequencyPutRequest) GetConnectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value
// and a boolean to check if the value has been set.
func (o *SyncUpdateConnectionFrequencyPutRequest) GetConnectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionId, true
}

// SetConnectionId sets field value
func (o *SyncUpdateConnectionFrequencyPutRequest) SetConnectionId(v string) {
	o.ConnectionId = v
}

// GetSyncName returns the SyncName field value
func (o *SyncUpdateConnectionFrequencyPutRequest) GetSyncName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SyncName
}

// GetSyncNameOk returns a tuple with the SyncName field value
// and a boolean to check if the value has been set.
func (o *SyncUpdateConnectionFrequencyPutRequest) GetSyncNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SyncName, true
}

// SetSyncName sets field value
func (o *SyncUpdateConnectionFrequencyPutRequest) SetSyncName(v string) {
	o.SyncName = v
}

// GetFrequency returns the Frequency field value
func (o *SyncUpdateConnectionFrequencyPutRequest) GetFrequency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value
// and a boolean to check if the value has been set.
func (o *SyncUpdateConnectionFrequencyPutRequest) GetFrequencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Frequency, true
}

// SetFrequency sets field value
func (o *SyncUpdateConnectionFrequencyPutRequest) SetFrequency(v string) {
	o.Frequency = v
}

func (o SyncUpdateConnectionFrequencyPutRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SyncUpdateConnectionFrequencyPutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["provider_config_key"] = o.ProviderConfigKey
	toSerialize["connection_id"] = o.ConnectionId
	toSerialize["sync_name"] = o.SyncName
	toSerialize["frequency"] = o.Frequency
	return toSerialize, nil
}

func (o *SyncUpdateConnectionFrequencyPutRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"provider_config_key",
		"connection_id",
		"sync_name",
		"frequency",
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

	varSyncUpdateConnectionFrequencyPutRequest := _SyncUpdateConnectionFrequencyPutRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSyncUpdateConnectionFrequencyPutRequest)

	if err != nil {
		return err
	}

	*o = SyncUpdateConnectionFrequencyPutRequest(varSyncUpdateConnectionFrequencyPutRequest)

	return err
}

type NullableSyncUpdateConnectionFrequencyPutRequest struct {
	value *SyncUpdateConnectionFrequencyPutRequest
	isSet bool
}

func (v NullableSyncUpdateConnectionFrequencyPutRequest) Get() *SyncUpdateConnectionFrequencyPutRequest {
	return v.value
}

func (v *NullableSyncUpdateConnectionFrequencyPutRequest) Set(val *SyncUpdateConnectionFrequencyPutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSyncUpdateConnectionFrequencyPutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSyncUpdateConnectionFrequencyPutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyncUpdateConnectionFrequencyPutRequest(val *SyncUpdateConnectionFrequencyPutRequest) *NullableSyncUpdateConnectionFrequencyPutRequest {
	return &NullableSyncUpdateConnectionFrequencyPutRequest{value: val, isSet: true}
}

func (v NullableSyncUpdateConnectionFrequencyPutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyncUpdateConnectionFrequencyPutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


