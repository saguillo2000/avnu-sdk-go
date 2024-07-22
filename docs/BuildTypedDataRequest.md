# BuildTypedDataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserAddress** | **string** | The user&#39;s address | 
**Calls** | [**[]Call**](Call.md) | The list of calls that will be executed | 
**GasTokenAddress** | Pointer to **string** | Gas token address.If null, there is two options:1. the user must have a reward compatible with the calls. In this case, the reward&#39;s sponsor will pay the gas fees in ETH.2. the api-key header must be field. The api-key&#39;s owner will be charged for the consumed gas fees in ETH | [optional] 
**MaxGasTokenAmount** | Pointer to **string** | Max gas token amountIf null, there is two options:1. the user must have a reward compatible with the calls. In this case, the reward&#39;s sponsor will pay the gas fees in ETH.2. the api-key header must be field. The api-key&#39;s owner will be charged for the consumed gas fees in ETH | [optional] 
**AccountClassHash** | Pointer to **string** | Only set this field when the account is not deployed. When the accountClassHash is defined, the API will not check the gasless compatibility by calling the supportsInterface entrypoint but will instead look into an internal map. If the classHash is not supported by the API, please contact us so we can quickly add support. | [optional] 

## Methods

### NewBuildTypedDataRequest

`func NewBuildTypedDataRequest(userAddress string, calls []Call, ) *BuildTypedDataRequest`

NewBuildTypedDataRequest instantiates a new BuildTypedDataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildTypedDataRequestWithDefaults

`func NewBuildTypedDataRequestWithDefaults() *BuildTypedDataRequest`

NewBuildTypedDataRequestWithDefaults instantiates a new BuildTypedDataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserAddress

`func (o *BuildTypedDataRequest) GetUserAddress() string`

GetUserAddress returns the UserAddress field if non-nil, zero value otherwise.

### GetUserAddressOk

`func (o *BuildTypedDataRequest) GetUserAddressOk() (*string, bool)`

GetUserAddressOk returns a tuple with the UserAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAddress

`func (o *BuildTypedDataRequest) SetUserAddress(v string)`

SetUserAddress sets UserAddress field to given value.


### GetCalls

`func (o *BuildTypedDataRequest) GetCalls() []Call`

GetCalls returns the Calls field if non-nil, zero value otherwise.

### GetCallsOk

`func (o *BuildTypedDataRequest) GetCallsOk() (*[]Call, bool)`

GetCallsOk returns a tuple with the Calls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalls

`func (o *BuildTypedDataRequest) SetCalls(v []Call)`

SetCalls sets Calls field to given value.


### GetGasTokenAddress

`func (o *BuildTypedDataRequest) GetGasTokenAddress() string`

GetGasTokenAddress returns the GasTokenAddress field if non-nil, zero value otherwise.

### GetGasTokenAddressOk

`func (o *BuildTypedDataRequest) GetGasTokenAddressOk() (*string, bool)`

GetGasTokenAddressOk returns a tuple with the GasTokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasTokenAddress

`func (o *BuildTypedDataRequest) SetGasTokenAddress(v string)`

SetGasTokenAddress sets GasTokenAddress field to given value.

### HasGasTokenAddress

`func (o *BuildTypedDataRequest) HasGasTokenAddress() bool`

HasGasTokenAddress returns a boolean if a field has been set.

### GetMaxGasTokenAmount

`func (o *BuildTypedDataRequest) GetMaxGasTokenAmount() string`

GetMaxGasTokenAmount returns the MaxGasTokenAmount field if non-nil, zero value otherwise.

### GetMaxGasTokenAmountOk

`func (o *BuildTypedDataRequest) GetMaxGasTokenAmountOk() (*string, bool)`

GetMaxGasTokenAmountOk returns a tuple with the MaxGasTokenAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxGasTokenAmount

`func (o *BuildTypedDataRequest) SetMaxGasTokenAmount(v string)`

SetMaxGasTokenAmount sets MaxGasTokenAmount field to given value.

### HasMaxGasTokenAmount

`func (o *BuildTypedDataRequest) HasMaxGasTokenAmount() bool`

HasMaxGasTokenAmount returns a boolean if a field has been set.

### GetAccountClassHash

`func (o *BuildTypedDataRequest) GetAccountClassHash() string`

GetAccountClassHash returns the AccountClassHash field if non-nil, zero value otherwise.

### GetAccountClassHashOk

`func (o *BuildTypedDataRequest) GetAccountClassHashOk() (*string, bool)`

GetAccountClassHashOk returns a tuple with the AccountClassHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountClassHash

`func (o *BuildTypedDataRequest) SetAccountClassHash(v string)`

SetAccountClassHash sets AccountClassHash field to given value.

### HasAccountClassHash

`func (o *BuildTypedDataRequest) HasAccountClassHash() bool`

HasAccountClassHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


