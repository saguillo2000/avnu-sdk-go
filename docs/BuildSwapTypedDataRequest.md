# BuildSwapTypedDataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuoteId** | **string** | The unique id of the quote | 
**TakerAddress** | Pointer to **string** | The address which will fill the quote. Not required if then taker address was provided during the quote request | [optional] 
**Slippage** | **float32** | The maximum acceptable slippage of the buyAmount amount. Default value is 1%. This value is ignored if slippage is not applicable to the selected quote.Min value for gasless transaction is 0.5% | 
**IncludeApprove** | **bool** | If true, the typed data will contains the approve call | 
**GasTokenAddress** | **string** | The gas token&#39;s address the user wants to spend to execute the tx. | 
**MaxGasTokenAmount** | Pointer to **string** |  | [optional] 

## Methods

### NewBuildSwapTypedDataRequest

`func NewBuildSwapTypedDataRequest(quoteId string, slippage float32, includeApprove bool, gasTokenAddress string, ) *BuildSwapTypedDataRequest`

NewBuildSwapTypedDataRequest instantiates a new BuildSwapTypedDataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildSwapTypedDataRequestWithDefaults

`func NewBuildSwapTypedDataRequestWithDefaults() *BuildSwapTypedDataRequest`

NewBuildSwapTypedDataRequestWithDefaults instantiates a new BuildSwapTypedDataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuoteId

`func (o *BuildSwapTypedDataRequest) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *BuildSwapTypedDataRequest) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *BuildSwapTypedDataRequest) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.


### GetTakerAddress

`func (o *BuildSwapTypedDataRequest) GetTakerAddress() string`

GetTakerAddress returns the TakerAddress field if non-nil, zero value otherwise.

### GetTakerAddressOk

`func (o *BuildSwapTypedDataRequest) GetTakerAddressOk() (*string, bool)`

GetTakerAddressOk returns a tuple with the TakerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakerAddress

`func (o *BuildSwapTypedDataRequest) SetTakerAddress(v string)`

SetTakerAddress sets TakerAddress field to given value.

### HasTakerAddress

`func (o *BuildSwapTypedDataRequest) HasTakerAddress() bool`

HasTakerAddress returns a boolean if a field has been set.

### GetSlippage

`func (o *BuildSwapTypedDataRequest) GetSlippage() float32`

GetSlippage returns the Slippage field if non-nil, zero value otherwise.

### GetSlippageOk

`func (o *BuildSwapTypedDataRequest) GetSlippageOk() (*float32, bool)`

GetSlippageOk returns a tuple with the Slippage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlippage

`func (o *BuildSwapTypedDataRequest) SetSlippage(v float32)`

SetSlippage sets Slippage field to given value.


### GetIncludeApprove

`func (o *BuildSwapTypedDataRequest) GetIncludeApprove() bool`

GetIncludeApprove returns the IncludeApprove field if non-nil, zero value otherwise.

### GetIncludeApproveOk

`func (o *BuildSwapTypedDataRequest) GetIncludeApproveOk() (*bool, bool)`

GetIncludeApproveOk returns a tuple with the IncludeApprove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeApprove

`func (o *BuildSwapTypedDataRequest) SetIncludeApprove(v bool)`

SetIncludeApprove sets IncludeApprove field to given value.


### GetGasTokenAddress

`func (o *BuildSwapTypedDataRequest) GetGasTokenAddress() string`

GetGasTokenAddress returns the GasTokenAddress field if non-nil, zero value otherwise.

### GetGasTokenAddressOk

`func (o *BuildSwapTypedDataRequest) GetGasTokenAddressOk() (*string, bool)`

GetGasTokenAddressOk returns a tuple with the GasTokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasTokenAddress

`func (o *BuildSwapTypedDataRequest) SetGasTokenAddress(v string)`

SetGasTokenAddress sets GasTokenAddress field to given value.


### GetMaxGasTokenAmount

`func (o *BuildSwapTypedDataRequest) GetMaxGasTokenAmount() string`

GetMaxGasTokenAmount returns the MaxGasTokenAmount field if non-nil, zero value otherwise.

### GetMaxGasTokenAmountOk

`func (o *BuildSwapTypedDataRequest) GetMaxGasTokenAmountOk() (*string, bool)`

GetMaxGasTokenAmountOk returns a tuple with the MaxGasTokenAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxGasTokenAmount

`func (o *BuildSwapTypedDataRequest) SetMaxGasTokenAmount(v string)`

SetMaxGasTokenAmount sets MaxGasTokenAmount field to given value.

### HasMaxGasTokenAmount

`func (o *BuildSwapTypedDataRequest) HasMaxGasTokenAmount() bool`

HasMaxGasTokenAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


