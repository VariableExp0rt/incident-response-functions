# AWSAPICallViaCloudTrail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestParameters** | [**NullableRequestParameters**](RequestParameters.md) |  | 
**UserIdentity** | [**UserIdentity**](UserIdentity.md) |  | 
**EventID** | **string** |  | 
**AwsRegion** | **string** |  | 
**EventVersion** | **string** |  | 
**ResponseElements** | **map[string]interface{}** |  | 
**SourceIPAddress** | **string** |  | 
**EventSource** | **string** |  | 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**ErrorCode** | Pointer to **string** |  | [optional] 
**UserAgent** | **string** |  | 
**EventType** | **string** |  | 
**ApiVersion** | Pointer to **string** |  | [optional] 
**RequestID** | **string** |  | 
**EventTime** | **time.Time** |  | 
**EventName** | **string** |  | 

## Methods

### NewAWSAPICallViaCloudTrail

`func NewAWSAPICallViaCloudTrail(requestParameters NullableRequestParameters, userIdentity UserIdentity, eventID string, awsRegion string, eventVersion string, responseElements map[string]interface{}, sourceIPAddress string, eventSource string, userAgent string, eventType string, requestID string, eventTime time.Time, eventName string, ) *AWSAPICallViaCloudTrail`

NewAWSAPICallViaCloudTrail instantiates a new AWSAPICallViaCloudTrail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSAPICallViaCloudTrailWithDefaults

`func NewAWSAPICallViaCloudTrailWithDefaults() *AWSAPICallViaCloudTrail`

NewAWSAPICallViaCloudTrailWithDefaults instantiates a new AWSAPICallViaCloudTrail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestParameters

`func (o *AWSAPICallViaCloudTrail) GetRequestParameters() RequestParameters`

GetRequestParameters returns the RequestParameters field if non-nil, zero value otherwise.

### GetRequestParametersOk

`func (o *AWSAPICallViaCloudTrail) GetRequestParametersOk() (*RequestParameters, bool)`

GetRequestParametersOk returns a tuple with the RequestParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestParameters

`func (o *AWSAPICallViaCloudTrail) SetRequestParameters(v RequestParameters)`

SetRequestParameters sets RequestParameters field to given value.


### SetRequestParametersNil

`func (o *AWSAPICallViaCloudTrail) SetRequestParametersNil(b bool)`

 SetRequestParametersNil sets the value for RequestParameters to be an explicit nil

### UnsetRequestParameters
`func (o *AWSAPICallViaCloudTrail) UnsetRequestParameters()`

UnsetRequestParameters ensures that no value is present for RequestParameters, not even an explicit nil
### GetUserIdentity

`func (o *AWSAPICallViaCloudTrail) GetUserIdentity() UserIdentity`

GetUserIdentity returns the UserIdentity field if non-nil, zero value otherwise.

### GetUserIdentityOk

`func (o *AWSAPICallViaCloudTrail) GetUserIdentityOk() (*UserIdentity, bool)`

GetUserIdentityOk returns a tuple with the UserIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdentity

`func (o *AWSAPICallViaCloudTrail) SetUserIdentity(v UserIdentity)`

SetUserIdentity sets UserIdentity field to given value.


### GetEventID

`func (o *AWSAPICallViaCloudTrail) GetEventID() string`

GetEventID returns the EventID field if non-nil, zero value otherwise.

### GetEventIDOk

`func (o *AWSAPICallViaCloudTrail) GetEventIDOk() (*string, bool)`

GetEventIDOk returns a tuple with the EventID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventID

`func (o *AWSAPICallViaCloudTrail) SetEventID(v string)`

SetEventID sets EventID field to given value.


### GetAwsRegion

`func (o *AWSAPICallViaCloudTrail) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *AWSAPICallViaCloudTrail) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *AWSAPICallViaCloudTrail) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.


### GetEventVersion

`func (o *AWSAPICallViaCloudTrail) GetEventVersion() string`

GetEventVersion returns the EventVersion field if non-nil, zero value otherwise.

### GetEventVersionOk

`func (o *AWSAPICallViaCloudTrail) GetEventVersionOk() (*string, bool)`

GetEventVersionOk returns a tuple with the EventVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventVersion

`func (o *AWSAPICallViaCloudTrail) SetEventVersion(v string)`

SetEventVersion sets EventVersion field to given value.


### GetResponseElements

`func (o *AWSAPICallViaCloudTrail) GetResponseElements() map[string]interface{}`

GetResponseElements returns the ResponseElements field if non-nil, zero value otherwise.

### GetResponseElementsOk

`func (o *AWSAPICallViaCloudTrail) GetResponseElementsOk() (*map[string]interface{}, bool)`

GetResponseElementsOk returns a tuple with the ResponseElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseElements

`func (o *AWSAPICallViaCloudTrail) SetResponseElements(v map[string]interface{})`

SetResponseElements sets ResponseElements field to given value.


### SetResponseElementsNil

`func (o *AWSAPICallViaCloudTrail) SetResponseElementsNil(b bool)`

 SetResponseElementsNil sets the value for ResponseElements to be an explicit nil

### UnsetResponseElements
`func (o *AWSAPICallViaCloudTrail) UnsetResponseElements()`

UnsetResponseElements ensures that no value is present for ResponseElements, not even an explicit nil
### GetSourceIPAddress

`func (o *AWSAPICallViaCloudTrail) GetSourceIPAddress() string`

GetSourceIPAddress returns the SourceIPAddress field if non-nil, zero value otherwise.

### GetSourceIPAddressOk

`func (o *AWSAPICallViaCloudTrail) GetSourceIPAddressOk() (*string, bool)`

GetSourceIPAddressOk returns a tuple with the SourceIPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIPAddress

`func (o *AWSAPICallViaCloudTrail) SetSourceIPAddress(v string)`

SetSourceIPAddress sets SourceIPAddress field to given value.


### GetEventSource

`func (o *AWSAPICallViaCloudTrail) GetEventSource() string`

GetEventSource returns the EventSource field if non-nil, zero value otherwise.

### GetEventSourceOk

`func (o *AWSAPICallViaCloudTrail) GetEventSourceOk() (*string, bool)`

GetEventSourceOk returns a tuple with the EventSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSource

`func (o *AWSAPICallViaCloudTrail) SetEventSource(v string)`

SetEventSource sets EventSource field to given value.


### GetErrorMessage

`func (o *AWSAPICallViaCloudTrail) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *AWSAPICallViaCloudTrail) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *AWSAPICallViaCloudTrail) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *AWSAPICallViaCloudTrail) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetErrorCode

`func (o *AWSAPICallViaCloudTrail) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *AWSAPICallViaCloudTrail) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *AWSAPICallViaCloudTrail) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *AWSAPICallViaCloudTrail) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetUserAgent

`func (o *AWSAPICallViaCloudTrail) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *AWSAPICallViaCloudTrail) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *AWSAPICallViaCloudTrail) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.


### GetEventType

`func (o *AWSAPICallViaCloudTrail) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *AWSAPICallViaCloudTrail) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *AWSAPICallViaCloudTrail) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetApiVersion

`func (o *AWSAPICallViaCloudTrail) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AWSAPICallViaCloudTrail) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AWSAPICallViaCloudTrail) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *AWSAPICallViaCloudTrail) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetRequestID

`func (o *AWSAPICallViaCloudTrail) GetRequestID() string`

GetRequestID returns the RequestID field if non-nil, zero value otherwise.

### GetRequestIDOk

`func (o *AWSAPICallViaCloudTrail) GetRequestIDOk() (*string, bool)`

GetRequestIDOk returns a tuple with the RequestID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestID

`func (o *AWSAPICallViaCloudTrail) SetRequestID(v string)`

SetRequestID sets RequestID field to given value.


### GetEventTime

`func (o *AWSAPICallViaCloudTrail) GetEventTime() time.Time`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *AWSAPICallViaCloudTrail) GetEventTimeOk() (*time.Time, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *AWSAPICallViaCloudTrail) SetEventTime(v time.Time)`

SetEventTime sets EventTime field to given value.


### GetEventName

`func (o *AWSAPICallViaCloudTrail) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *AWSAPICallViaCloudTrail) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *AWSAPICallViaCloudTrail) SetEventName(v string)`

SetEventName sets EventName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


