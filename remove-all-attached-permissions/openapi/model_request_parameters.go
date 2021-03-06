/*
 * AWSAPICallViaCloudTrail
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// RequestParameters struct for RequestParameters
type RequestParameters struct {
	LogGroupName NullableString `json:"logGroupName"`
	LogStreamName NullableString `json:"logStreamName,omitempty"`
}

// NewRequestParameters instantiates a new RequestParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestParameters(logGroupName NullableString, ) *RequestParameters {
	this := RequestParameters{}
	this.LogGroupName = logGroupName
	return &this
}

// NewRequestParametersWithDefaults instantiates a new RequestParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestParametersWithDefaults() *RequestParameters {
	this := RequestParameters{}
	return &this
}

// GetLogGroupName returns the LogGroupName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RequestParameters) GetLogGroupName() string {
	if o == nil || o.LogGroupName.Get() == nil {
		var ret string
		return ret
	}

	return *o.LogGroupName.Get()
}

// GetLogGroupNameOk returns a tuple with the LogGroupName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestParameters) GetLogGroupNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LogGroupName.Get(), o.LogGroupName.IsSet()
}

// SetLogGroupName sets field value
func (o *RequestParameters) SetLogGroupName(v string) {
	o.LogGroupName.Set(&v)
}

// GetLogStreamName returns the LogStreamName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestParameters) GetLogStreamName() string {
	if o == nil || o.LogStreamName.Get() == nil {
		var ret string
		return ret
	}
	return *o.LogStreamName.Get()
}

// GetLogStreamNameOk returns a tuple with the LogStreamName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestParameters) GetLogStreamNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LogStreamName.Get(), o.LogStreamName.IsSet()
}

// HasLogStreamName returns a boolean if a field has been set.
func (o *RequestParameters) HasLogStreamName() bool {
	if o != nil && o.LogStreamName.IsSet() {
		return true
	}

	return false
}

// SetLogStreamName gets a reference to the given NullableString and assigns it to the LogStreamName field.
func (o *RequestParameters) SetLogStreamName(v string) {
	o.LogStreamName.Set(&v)
}
// SetLogStreamNameNil sets the value for LogStreamName to be an explicit nil
func (o *RequestParameters) SetLogStreamNameNil() {
	o.LogStreamName.Set(nil)
}

// UnsetLogStreamName ensures that no value is present for LogStreamName, not even an explicit nil
func (o *RequestParameters) UnsetLogStreamName() {
	o.LogStreamName.Unset()
}

func (o RequestParameters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["logGroupName"] = o.LogGroupName.Get()
	}
	if o.LogStreamName.IsSet() {
		toSerialize["logStreamName"] = o.LogStreamName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRequestParameters struct {
	value *RequestParameters
	isSet bool
}

func (v NullableRequestParameters) Get() *RequestParameters {
	return v.value
}

func (v *NullableRequestParameters) Set(val *RequestParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestParameters(val *RequestParameters) *NullableRequestParameters {
	return &NullableRequestParameters{value: val, isSet: true}
}

func (v NullableRequestParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


