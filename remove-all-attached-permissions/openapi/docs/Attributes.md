# Attributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MfaAuthenticated** | **string** |  | 
**CreationDate** | **time.Time** |  | 

## Methods

### NewAttributes

`func NewAttributes(mfaAuthenticated string, creationDate time.Time, ) *Attributes`

NewAttributes instantiates a new Attributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributesWithDefaults

`func NewAttributesWithDefaults() *Attributes`

NewAttributesWithDefaults instantiates a new Attributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMfaAuthenticated

`func (o *Attributes) GetMfaAuthenticated() string`

GetMfaAuthenticated returns the MfaAuthenticated field if non-nil, zero value otherwise.

### GetMfaAuthenticatedOk

`func (o *Attributes) GetMfaAuthenticatedOk() (*string, bool)`

GetMfaAuthenticatedOk returns a tuple with the MfaAuthenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaAuthenticated

`func (o *Attributes) SetMfaAuthenticated(v string)`

SetMfaAuthenticated sets MfaAuthenticated field to given value.


### GetCreationDate

`func (o *Attributes) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *Attributes) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *Attributes) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


