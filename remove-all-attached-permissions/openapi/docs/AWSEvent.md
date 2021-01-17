# AWSEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | [**AWSAPICallViaCloudTrail**](AWSAPICallViaCloudTrail.md) |  | 
**DetailType** | **string** |  | 
**Resources** | **[]string** |  | 
**Id** | **string** |  | 
**Source** | **string** |  | 
**Time** | **time.Time** |  | 
**Region** | **string** |  | 
**Version** | **string** |  | 
**Account** | **string** |  | 

## Methods

### NewAWSEvent

`func NewAWSEvent(detail AWSAPICallViaCloudTrail, detailType string, resources []string, id string, source string, time time.Time, region string, version string, account string, ) *AWSEvent`

NewAWSEvent instantiates a new AWSEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSEventWithDefaults

`func NewAWSEventWithDefaults() *AWSEvent`

NewAWSEventWithDefaults instantiates a new AWSEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *AWSEvent) GetDetail() AWSAPICallViaCloudTrail`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *AWSEvent) GetDetailOk() (*AWSAPICallViaCloudTrail, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *AWSEvent) SetDetail(v AWSAPICallViaCloudTrail)`

SetDetail sets Detail field to given value.


### GetDetailType

`func (o *AWSEvent) GetDetailType() string`

GetDetailType returns the DetailType field if non-nil, zero value otherwise.

### GetDetailTypeOk

`func (o *AWSEvent) GetDetailTypeOk() (*string, bool)`

GetDetailTypeOk returns a tuple with the DetailType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailType

`func (o *AWSEvent) SetDetailType(v string)`

SetDetailType sets DetailType field to given value.


### GetResources

`func (o *AWSEvent) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *AWSEvent) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *AWSEvent) SetResources(v []string)`

SetResources sets Resources field to given value.


### GetId

`func (o *AWSEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AWSEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AWSEvent) SetId(v string)`

SetId sets Id field to given value.


### GetSource

`func (o *AWSEvent) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AWSEvent) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AWSEvent) SetSource(v string)`

SetSource sets Source field to given value.


### GetTime

`func (o *AWSEvent) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *AWSEvent) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *AWSEvent) SetTime(v time.Time)`

SetTime sets Time field to given value.


### GetRegion

`func (o *AWSEvent) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AWSEvent) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AWSEvent) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetVersion

`func (o *AWSEvent) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AWSEvent) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AWSEvent) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetAccount

`func (o *AWSEvent) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AWSEvent) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AWSEvent) SetAccount(v string)`

SetAccount sets Account field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


