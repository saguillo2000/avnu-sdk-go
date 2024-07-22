# ExecuteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserAddress** | **string** | The user&#39;s address | 
**TypedData** | **string** | The typed data that the user signed | 
**Signature** | **[]string** | The user&#39;s typed data signature | 
**DeploymentData** | Pointer to [**DeploymentDataDto**](DeploymentDataDto.md) |  | [optional] 

## Methods

### NewExecuteRequest

`func NewExecuteRequest(userAddress string, typedData string, signature []string, ) *ExecuteRequest`

NewExecuteRequest instantiates a new ExecuteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecuteRequestWithDefaults

`func NewExecuteRequestWithDefaults() *ExecuteRequest`

NewExecuteRequestWithDefaults instantiates a new ExecuteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserAddress

`func (o *ExecuteRequest) GetUserAddress() string`

GetUserAddress returns the UserAddress field if non-nil, zero value otherwise.

### GetUserAddressOk

`func (o *ExecuteRequest) GetUserAddressOk() (*string, bool)`

GetUserAddressOk returns a tuple with the UserAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAddress

`func (o *ExecuteRequest) SetUserAddress(v string)`

SetUserAddress sets UserAddress field to given value.


### GetTypedData

`func (o *ExecuteRequest) GetTypedData() string`

GetTypedData returns the TypedData field if non-nil, zero value otherwise.

### GetTypedDataOk

`func (o *ExecuteRequest) GetTypedDataOk() (*string, bool)`

GetTypedDataOk returns a tuple with the TypedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypedData

`func (o *ExecuteRequest) SetTypedData(v string)`

SetTypedData sets TypedData field to given value.


### GetSignature

`func (o *ExecuteRequest) GetSignature() []string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *ExecuteRequest) GetSignatureOk() (*[]string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *ExecuteRequest) SetSignature(v []string)`

SetSignature sets Signature field to given value.


### GetDeploymentData

`func (o *ExecuteRequest) GetDeploymentData() DeploymentDataDto`

GetDeploymentData returns the DeploymentData field if non-nil, zero value otherwise.

### GetDeploymentDataOk

`func (o *ExecuteRequest) GetDeploymentDataOk() (*DeploymentDataDto, bool)`

GetDeploymentDataOk returns a tuple with the DeploymentData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentData

`func (o *ExecuteRequest) SetDeploymentData(v DeploymentDataDto)`

SetDeploymentData sets DeploymentData field to given value.

### HasDeploymentData

`func (o *ExecuteRequest) HasDeploymentData() bool`

HasDeploymentData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


