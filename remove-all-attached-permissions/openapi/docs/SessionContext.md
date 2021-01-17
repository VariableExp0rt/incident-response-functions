# SessionContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionIssuer** | [**SessionIssuer**](SessionIssuer.md) |  | 
**Attributes** | [**Attributes**](Attributes.md) |  | 

## Methods

### NewSessionContext

`func NewSessionContext(sessionIssuer SessionIssuer, attributes Attributes, ) *SessionContext`

NewSessionContext instantiates a new SessionContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionContextWithDefaults

`func NewSessionContextWithDefaults() *SessionContext`

NewSessionContextWithDefaults instantiates a new SessionContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionIssuer

`func (o *SessionContext) GetSessionIssuer() SessionIssuer`

GetSessionIssuer returns the SessionIssuer field if non-nil, zero value otherwise.

### GetSessionIssuerOk

`func (o *SessionContext) GetSessionIssuerOk() (*SessionIssuer, bool)`

GetSessionIssuerOk returns a tuple with the SessionIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionIssuer

`func (o *SessionContext) SetSessionIssuer(v SessionIssuer)`

SetSessionIssuer sets SessionIssuer field to given value.


### GetAttributes

`func (o *SessionContext) GetAttributes() Attributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SessionContext) GetAttributesOk() (*Attributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SessionContext) SetAttributes(v Attributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


