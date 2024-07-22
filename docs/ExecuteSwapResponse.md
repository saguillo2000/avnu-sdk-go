# ExecuteSwapResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionHash** | **string** | The hash of the transaction | 
**GasTokenAddress** | Pointer to **string** | The address of the gas token used to pay gas fees | [optional] 
**GasTokenAmount** | Pointer to **string** | The amount of gas token used to pay gas fees | [optional] 

## Methods

### NewExecuteSwapResponse

`func NewExecuteSwapResponse(transactionHash string, ) *ExecuteSwapResponse`

NewExecuteSwapResponse instantiates a new ExecuteSwapResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecuteSwapResponseWithDefaults

`func NewExecuteSwapResponseWithDefaults() *ExecuteSwapResponse`

NewExecuteSwapResponseWithDefaults instantiates a new ExecuteSwapResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionHash

`func (o *ExecuteSwapResponse) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *ExecuteSwapResponse) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *ExecuteSwapResponse) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.


### GetGasTokenAddress

`func (o *ExecuteSwapResponse) GetGasTokenAddress() string`

GetGasTokenAddress returns the GasTokenAddress field if non-nil, zero value otherwise.

### GetGasTokenAddressOk

`func (o *ExecuteSwapResponse) GetGasTokenAddressOk() (*string, bool)`

GetGasTokenAddressOk returns a tuple with the GasTokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasTokenAddress

`func (o *ExecuteSwapResponse) SetGasTokenAddress(v string)`

SetGasTokenAddress sets GasTokenAddress field to given value.

### HasGasTokenAddress

`func (o *ExecuteSwapResponse) HasGasTokenAddress() bool`

HasGasTokenAddress returns a boolean if a field has been set.

### GetGasTokenAmount

`func (o *ExecuteSwapResponse) GetGasTokenAmount() string`

GetGasTokenAmount returns the GasTokenAmount field if non-nil, zero value otherwise.

### GetGasTokenAmountOk

`func (o *ExecuteSwapResponse) GetGasTokenAmountOk() (*string, bool)`

GetGasTokenAmountOk returns a tuple with the GasTokenAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasTokenAmount

`func (o *ExecuteSwapResponse) SetGasTokenAmount(v string)`

SetGasTokenAmount sets GasTokenAmount field to given value.

### HasGasTokenAmount

`func (o *ExecuteSwapResponse) HasGasTokenAmount() bool`

HasGasTokenAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


