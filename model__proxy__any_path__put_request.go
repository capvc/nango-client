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

// checks if the ProxyAnyPathPutRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProxyAnyPathPutRequest{}

// ProxyAnyPathPutRequest struct for ProxyAnyPathPutRequest
type ProxyAnyPathPutRequest struct {
	ANY_BODY_PARAMS *string `json:"$ANY_BODY_PARAMS,omitempty"`
}

// NewProxyAnyPathPutRequest instantiates a new ProxyAnyPathPutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProxyAnyPathPutRequest() *ProxyAnyPathPutRequest {
	this := ProxyAnyPathPutRequest{}
	return &this
}

// NewProxyAnyPathPutRequestWithDefaults instantiates a new ProxyAnyPathPutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProxyAnyPathPutRequestWithDefaults() *ProxyAnyPathPutRequest {
	this := ProxyAnyPathPutRequest{}
	return &this
}

// GetANY_BODY_PARAMS returns the ANY_BODY_PARAMS field value if set, zero value otherwise.
func (o *ProxyAnyPathPutRequest) GetANY_BODY_PARAMS() string {
	if o == nil || IsNil(o.ANY_BODY_PARAMS) {
		var ret string
		return ret
	}
	return *o.ANY_BODY_PARAMS
}

// GetANY_BODY_PARAMSOk returns a tuple with the ANY_BODY_PARAMS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyAnyPathPutRequest) GetANY_BODY_PARAMSOk() (*string, bool) {
	if o == nil || IsNil(o.ANY_BODY_PARAMS) {
		return nil, false
	}
	return o.ANY_BODY_PARAMS, true
}

// HasANY_BODY_PARAMS returns a boolean if a field has been set.
func (o *ProxyAnyPathPutRequest) HasANY_BODY_PARAMS() bool {
	if o != nil && !IsNil(o.ANY_BODY_PARAMS) {
		return true
	}

	return false
}

// SetANY_BODY_PARAMS gets a reference to the given string and assigns it to the ANY_BODY_PARAMS field.
func (o *ProxyAnyPathPutRequest) SetANY_BODY_PARAMS(v string) {
	o.ANY_BODY_PARAMS = &v
}

func (o ProxyAnyPathPutRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProxyAnyPathPutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ANY_BODY_PARAMS) {
		toSerialize["$ANY_BODY_PARAMS"] = o.ANY_BODY_PARAMS
	}
	return toSerialize, nil
}

type NullableProxyAnyPathPutRequest struct {
	value *ProxyAnyPathPutRequest
	isSet bool
}

func (v NullableProxyAnyPathPutRequest) Get() *ProxyAnyPathPutRequest {
	return v.value
}

func (v *NullableProxyAnyPathPutRequest) Set(val *ProxyAnyPathPutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProxyAnyPathPutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProxyAnyPathPutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProxyAnyPathPutRequest(val *ProxyAnyPathPutRequest) *NullableProxyAnyPathPutRequest {
	return &NullableProxyAnyPathPutRequest{value: val, isSet: true}
}

func (v NullableProxyAnyPathPutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProxyAnyPathPutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


