/*
Nango API

Nango API specs used to authorize & sync data with external APIs.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nango

import (
	"encoding/json"
)

// checks if the ScriptsConfigGet200ResponseSyncsInnerModelsInnerFieldsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScriptsConfigGet200ResponseSyncsInnerModelsInnerFieldsInner{}

// ScriptsConfigGet200ResponseSyncsInnerModelsInnerFieldsInner struct for ScriptsConfigGet200ResponseSyncsInnerModelsInnerFieldsInner
type ScriptsConfigGet200ResponseSyncsInnerModelsInnerFieldsInner struct {
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewScriptsConfigGet200ResponseSyncsInnerModelsInnerFieldsInner instantiates a new ScriptsConfigGet200ResponseSyncsInnerModelsInnerFieldsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScriptsConfigGet200ResponseSyncsInnerModelsInnerFieldsInner() *ScriptsConfigGet200ResponseSyncsInnerModelsInnerFieldsInner {
	this := ScriptsConfigGet200ResponseSyncsInnerModelsInnerFieldsInner{}
	return &this
}

// NewScriptsConfigGet200ResponseSyncsInnerModelsInnerFieldsInnerWithDefaults instantiates a new ScriptsConfigGet200ResponseSyncsInnerModelsInnerFieldsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScriptsConfigGet200ResponseSyncsInnerModelsInnerFieldsInnerWithDefaults() *ScriptsConfigGet200ResponseSyncsInnerModelsInnerFieldsInner {
	this := ScriptsConfigGet200ResponseSyncsInnerModelsInnerFieldsInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ScriptsConfigGet200ResponseSyncsInnerModelsInnerFieldsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScriptsConfigGet200ResponseSyncsInnerModelsInnerFieldsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ScriptsConfigGet200ResponseSyncsInnerModelsInnerFieldsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ScriptsConfigGet200ResponseSyncsInnerModelsInnerFieldsInner) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ScriptsConfigGet200ResponseSyncsInnerModelsInnerFieldsInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScriptsConfigGet200ResponseSyncsInnerModelsInnerFieldsInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ScriptsConfigGet200ResponseSyncsInnerModelsInnerFieldsInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ScriptsConfigGet200ResponseSyncsInnerModelsInnerFieldsInner) SetType(v string) {
	o.Type = &v
}

func (o ScriptsConfigGet200ResponseSyncsInnerModelsInnerFieldsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScriptsConfigGet200ResponseSyncsInnerModelsInnerFieldsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableScriptsConfigGet200ResponseSyncsInnerModelsInnerFieldsInner struct {
	value *ScriptsConfigGet200ResponseSyncsInnerModelsInnerFieldsInner
	isSet bool
}

func (v NullableScriptsConfigGet200ResponseSyncsInnerModelsInnerFieldsInner) Get() *ScriptsConfigGet200ResponseSyncsInnerModelsInnerFieldsInner {
	return v.value
}

func (v *NullableScriptsConfigGet200ResponseSyncsInnerModelsInnerFieldsInner) Set(val *ScriptsConfigGet200ResponseSyncsInnerModelsInnerFieldsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableScriptsConfigGet200ResponseSyncsInnerModelsInnerFieldsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableScriptsConfigGet200ResponseSyncsInnerModelsInnerFieldsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScriptsConfigGet200ResponseSyncsInnerModelsInnerFieldsInner(val *ScriptsConfigGet200ResponseSyncsInnerModelsInnerFieldsInner) *NullableScriptsConfigGet200ResponseSyncsInnerModelsInnerFieldsInner {
	return &NullableScriptsConfigGet200ResponseSyncsInnerModelsInnerFieldsInner{value: val, isSet: true}
}

func (v NullableScriptsConfigGet200ResponseSyncsInnerModelsInnerFieldsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScriptsConfigGet200ResponseSyncsInnerModelsInnerFieldsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


