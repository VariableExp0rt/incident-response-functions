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
	"time"
)

// AWSEvent struct for AWSEvent
type AWSEvent struct {
	Detail AWSAPICallViaCloudTrail `json:"detail"`
	DetailType string `json:"detail-type"`
	Resources []string `json:"resources"`
	Id string `json:"id"`
	Source string `json:"source"`
	Time time.Time `json:"time"`
	Region string `json:"region"`
	Version string `json:"version"`
	Account string `json:"account"`
}

// NewAWSEvent instantiates a new AWSEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAWSEvent(detail AWSAPICallViaCloudTrail, detailType string, resources []string, id string, source string, time time.Time, region string, version string, account string, ) *AWSEvent {
	this := AWSEvent{}
	this.Detail = detail
	this.DetailType = detailType
	this.Resources = resources
	this.Id = id
	this.Source = source
	this.Time = time
	this.Region = region
	this.Version = version
	this.Account = account
	return &this
}

// NewAWSEventWithDefaults instantiates a new AWSEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAWSEventWithDefaults() *AWSEvent {
	this := AWSEvent{}
	return &this
}

// GetDetail returns the Detail field value
func (o *AWSEvent) GetDetail() AWSAPICallViaCloudTrail {
	if o == nil  {
		var ret AWSAPICallViaCloudTrail
		return ret
	}

	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
func (o *AWSEvent) GetDetailOk() (*AWSAPICallViaCloudTrail, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Detail, true
}

// SetDetail sets field value
func (o *AWSEvent) SetDetail(v AWSAPICallViaCloudTrail) {
	o.Detail = v
}

// GetDetailType returns the DetailType field value
func (o *AWSEvent) GetDetailType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.DetailType
}

// GetDetailTypeOk returns a tuple with the DetailType field value
// and a boolean to check if the value has been set.
func (o *AWSEvent) GetDetailTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DetailType, true
}

// SetDetailType sets field value
func (o *AWSEvent) SetDetailType(v string) {
	o.DetailType = v
}

// GetResources returns the Resources field value
func (o *AWSEvent) GetResources() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *AWSEvent) GetResourcesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Resources, true
}

// SetResources sets field value
func (o *AWSEvent) SetResources(v []string) {
	o.Resources = v
}

// GetId returns the Id field value
func (o *AWSEvent) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AWSEvent) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AWSEvent) SetId(v string) {
	o.Id = v
}

// GetSource returns the Source field value
func (o *AWSEvent) GetSource() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *AWSEvent) GetSourceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *AWSEvent) SetSource(v string) {
	o.Source = v
}

// GetTime returns the Time field value
func (o *AWSEvent) GetTime() time.Time {
	if o == nil  {
		var ret time.Time
		return ret
	}

	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *AWSEvent) GetTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value
func (o *AWSEvent) SetTime(v time.Time) {
	o.Time = v
}

// GetRegion returns the Region field value
func (o *AWSEvent) GetRegion() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *AWSEvent) GetRegionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *AWSEvent) SetRegion(v string) {
	o.Region = v
}

// GetVersion returns the Version field value
func (o *AWSEvent) GetVersion() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *AWSEvent) GetVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *AWSEvent) SetVersion(v string) {
	o.Version = v
}

// GetAccount returns the Account field value
func (o *AWSEvent) GetAccount() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *AWSEvent) GetAccountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *AWSEvent) SetAccount(v string) {
	o.Account = v
}

func (o AWSEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["detail"] = o.Detail
	}
	if true {
		toSerialize["detail-type"] = o.DetailType
	}
	if true {
		toSerialize["resources"] = o.Resources
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["time"] = o.Time
	}
	if true {
		toSerialize["region"] = o.Region
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["account"] = o.Account
	}
	return json.Marshal(toSerialize)
}

type NullableAWSEvent struct {
	value *AWSEvent
	isSet bool
}

func (v NullableAWSEvent) Get() *AWSEvent {
	return v.value
}

func (v *NullableAWSEvent) Set(val *AWSEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableAWSEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableAWSEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAWSEvent(val *AWSEvent) *NullableAWSEvent {
	return &NullableAWSEvent{value: val, isSet: true}
}

func (v NullableAWSEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAWSEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


