# RequestParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogGroupName** | **NullableString** |  | 
**LogStreamName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewRequestParameters

`func NewRequestParameters(logGroupName NullableString, ) *RequestParameters`

NewRequestParameters instantiates a new RequestParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestParametersWithDefaults

`func NewRequestParametersWithDefaults() *RequestParameters`

NewRequestParametersWithDefaults instantiates a new RequestParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogGroupName

`func (o *RequestParameters) GetLogGroupName() string`

GetLogGroupName returns the LogGroupName field if non-nil, zero value otherwise.

### GetLogGroupNameOk

`func (o *RequestParameters) GetLogGroupNameOk() (*string, bool)`

GetLogGroupNameOk returns a tuple with the LogGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogGroupName

`func (o *RequestParameters) SetLogGroupName(v string)`

SetLogGroupName sets LogGroupName field to given value.


### SetLogGroupNameNil

`func (o *RequestParameters) SetLogGroupNameNil(b bool)`

 SetLogGroupNameNil sets the value for LogGroupName to be an explicit nil

### UnsetLogGroupName
`func (o *RequestParameters) UnsetLogGroupName()`

UnsetLogGroupName ensures that no value is present for LogGroupName, not even an explicit nil
### GetLogStreamName

`func (o *RequestParameters) GetLogStreamName() string`

GetLogStreamName returns the LogStreamName field if non-nil, zero value otherwise.

### GetLogStreamNameOk

`func (o *RequestParameters) GetLogStreamNameOk() (*string, bool)`

GetLogStreamNameOk returns a tuple with the LogStreamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogStreamName

`func (o *RequestParameters) SetLogStreamName(v string)`

SetLogStreamName sets LogStreamName field to given value.

### HasLogStreamName

`func (o *RequestParameters) HasLogStreamName() bool`

HasLogStreamName returns a boolean if a field has been set.

### SetLogStreamNameNil

`func (o *RequestParameters) SetLogStreamNameNil(b bool)`

 SetLogStreamNameNil sets the value for LogStreamName to be an explicit nil

### UnsetLogStreamName
`func (o *RequestParameters) UnsetLogStreamName()`

UnsetLogStreamName ensures that no value is present for LogStreamName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


