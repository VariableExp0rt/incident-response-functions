# SessionIssuer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  | 
**PrincipalId** | **string** |  | 
**Type** | **string** |  | 
**Arn** | **string** |  | 
**UserName** | **string** |  | 

## Methods

### NewSessionIssuer

`func NewSessionIssuer(accountId string, principalId string, type_ string, arn string, userName string, ) *SessionIssuer`

NewSessionIssuer instantiates a new SessionIssuer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionIssuerWithDefaults

`func NewSessionIssuerWithDefaults() *SessionIssuer`

NewSessionIssuerWithDefaults instantiates a new SessionIssuer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *SessionIssuer) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SessionIssuer) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SessionIssuer) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetPrincipalId

`func (o *SessionIssuer) GetPrincipalId() string`

GetPrincipalId returns the PrincipalId field if non-nil, zero value otherwise.

### GetPrincipalIdOk

`func (o *SessionIssuer) GetPrincipalIdOk() (*string, bool)`

GetPrincipalIdOk returns a tuple with the PrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalId

`func (o *SessionIssuer) SetPrincipalId(v string)`

SetPrincipalId sets PrincipalId field to given value.


### GetType

`func (o *SessionIssuer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SessionIssuer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SessionIssuer) SetType(v string)`

SetType sets Type field to given value.


### GetArn

`func (o *SessionIssuer) GetArn() string`

GetArn returns the Arn field if non-nil, zero value otherwise.

### GetArnOk

`func (o *SessionIssuer) GetArnOk() (*string, bool)`

GetArnOk returns a tuple with the Arn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArn

`func (o *SessionIssuer) SetArn(v string)`

SetArn sets Arn field to given value.


### GetUserName

`func (o *SessionIssuer) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *SessionIssuer) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *SessionIssuer) SetUserName(v string)`

SetUserName sets UserName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


