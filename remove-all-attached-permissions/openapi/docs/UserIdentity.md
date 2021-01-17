# UserIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionContext** | [**SessionContext**](SessionContext.md) |  | 
**AccessKeyId** | **string** |  | 
**AccountId** | **string** |  | 
**PrincipalId** | **string** |  | 
**Type** | **string** |  | 
**Arn** | **string** |  | 
**InvokedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewUserIdentity

`func NewUserIdentity(sessionContext SessionContext, accessKeyId string, accountId string, principalId string, type_ string, arn string, ) *UserIdentity`

NewUserIdentity instantiates a new UserIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserIdentityWithDefaults

`func NewUserIdentityWithDefaults() *UserIdentity`

NewUserIdentityWithDefaults instantiates a new UserIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionContext

`func (o *UserIdentity) GetSessionContext() SessionContext`

GetSessionContext returns the SessionContext field if non-nil, zero value otherwise.

### GetSessionContextOk

`func (o *UserIdentity) GetSessionContextOk() (*SessionContext, bool)`

GetSessionContextOk returns a tuple with the SessionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionContext

`func (o *UserIdentity) SetSessionContext(v SessionContext)`

SetSessionContext sets SessionContext field to given value.


### GetAccessKeyId

`func (o *UserIdentity) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *UserIdentity) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *UserIdentity) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.


### GetAccountId

`func (o *UserIdentity) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UserIdentity) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UserIdentity) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetPrincipalId

`func (o *UserIdentity) GetPrincipalId() string`

GetPrincipalId returns the PrincipalId field if non-nil, zero value otherwise.

### GetPrincipalIdOk

`func (o *UserIdentity) GetPrincipalIdOk() (*string, bool)`

GetPrincipalIdOk returns a tuple with the PrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalId

`func (o *UserIdentity) SetPrincipalId(v string)`

SetPrincipalId sets PrincipalId field to given value.


### GetType

`func (o *UserIdentity) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserIdentity) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserIdentity) SetType(v string)`

SetType sets Type field to given value.


### GetArn

`func (o *UserIdentity) GetArn() string`

GetArn returns the Arn field if non-nil, zero value otherwise.

### GetArnOk

`func (o *UserIdentity) GetArnOk() (*string, bool)`

GetArnOk returns a tuple with the Arn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArn

`func (o *UserIdentity) SetArn(v string)`

SetArn sets Arn field to given value.


### GetInvokedBy

`func (o *UserIdentity) GetInvokedBy() string`

GetInvokedBy returns the InvokedBy field if non-nil, zero value otherwise.

### GetInvokedByOk

`func (o *UserIdentity) GetInvokedByOk() (*string, bool)`

GetInvokedByOk returns a tuple with the InvokedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokedBy

`func (o *UserIdentity) SetInvokedBy(v string)`

SetInvokedBy sets InvokedBy field to given value.

### HasInvokedBy

`func (o *UserIdentity) HasInvokedBy() bool`

HasInvokedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


