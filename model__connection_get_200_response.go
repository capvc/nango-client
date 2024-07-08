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

// checks if the ConnectionGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectionGet200Response{}

// ConnectionGet200Response struct for ConnectionGet200Response
type ConnectionGet200Response struct {
	Configs []ConnectionGet200ResponseConfigsInner `json:"configs"`
}

type _ConnectionGet200Response ConnectionGet200Response

// NewConnectionGet200Response instantiates a new ConnectionGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionGet200Response(configs []ConnectionGet200ResponseConfigsInner) *ConnectionGet200Response {
	this := ConnectionGet200Response{}
	this.Configs = configs
	return &this
}

// NewConnectionGet200ResponseWithDefaults instantiates a new ConnectionGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionGet200ResponseWithDefaults() *ConnectionGet200Response {
	this := ConnectionGet200Response{}
	return &this
}

// GetConfigs returns the Configs field value
func (o *ConnectionGet200Response) GetConfigs() []ConnectionGet200ResponseConfigsInner {
	if o == nil {
		var ret []ConnectionGet200ResponseConfigsInner
		return ret
	}

	return o.Configs
}

// GetConfigsOk returns a tuple with the Configs field value
// and a boolean to check if the value has been set.
func (o *ConnectionGet200Response) GetConfigsOk() ([]ConnectionGet200ResponseConfigsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Configs, true
}

// SetConfigs sets field value
func (o *ConnectionGet200Response) SetConfigs(v []ConnectionGet200ResponseConfigsInner) {
	o.Configs = v
}

func (o ConnectionGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectionGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["configs"] = o.Configs
	return toSerialize, nil
}

func (o *ConnectionGet200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"configs",
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

	varConnectionGet200Response := _ConnectionGet200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varConnectionGet200Response)

	if err != nil {
		return err
	}

	*o = ConnectionGet200Response(varConnectionGet200Response)

	return err
}

type NullableConnectionGet200Response struct {
	value *ConnectionGet200Response
	isSet bool
}

func (v NullableConnectionGet200Response) Get() *ConnectionGet200Response {
	return v.value
}

func (v *NullableConnectionGet200Response) Set(val *ConnectionGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionGet200Response(val *ConnectionGet200Response) *NullableConnectionGet200Response {
	return &NullableConnectionGet200Response{value: val, isSet: true}
}

func (v NullableConnectionGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


