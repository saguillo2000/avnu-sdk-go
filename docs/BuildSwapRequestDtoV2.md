# BuildSwapRequestDtoV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuoteId** | **string** | The unique id of the quote | 
**TakerAddress** | Pointer to **string** | The address which will fill the quote. Not required if then taker address was provided during the quote request | [optional] 
**Slippage** | **float32** | The maximum acceptable slippage of the buyAmount amount. Default value is 5%. This value is ignored if slippage is not applicable to the selected quote | 
**IncludeApprove** | **bool** | If true, the response will contains the approve call (if necessary) | 

## Methods

### NewBuildSwapRequestDtoV2

`func NewBuildSwapRequestDtoV2(quoteId string, slippage float32, includeApprove bool, ) *BuildSwapRequestDtoV2`

NewBuildSwapRequestDtoV2 instantiates a new BuildSwapRequestDtoV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildSwapRequestDtoV2WithDefaults

`func NewBuildSwapRequestDtoV2WithDefaults() *BuildSwapRequestDtoV2`

NewBuildSwapRequestDtoV2WithDefaults instantiates a new BuildSwapRequestDtoV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuoteId

`func (o *BuildSwapRequestDtoV2) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *BuildSwapRequestDtoV2) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *BuildSwapRequestDtoV2) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.


### GetTakerAddress

`func (o *BuildSwapRequestDtoV2) GetTakerAddress() string`

GetTakerAddress returns the TakerAddress field if non-nil, zero value otherwise.

### GetTakerAddressOk

`func (o *BuildSwapRequestDtoV2) GetTakerAddressOk() (*string, bool)`

GetTakerAddressOk returns a tuple with the TakerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakerAddress

`func (o *BuildSwapRequestDtoV2) SetTakerAddress(v string)`

SetTakerAddress sets TakerAddress field to given value.

### HasTakerAddress

`func (o *BuildSwapRequestDtoV2) HasTakerAddress() bool`

HasTakerAddress returns a boolean if a field has been set.

### GetSlippage

`func (o *BuildSwapRequestDtoV2) GetSlippage() float32`

GetSlippage returns the Slippage field if non-nil, zero value otherwise.

### GetSlippageOk

`func (o *BuildSwapRequestDtoV2) GetSlippageOk() (*float32, bool)`

GetSlippageOk returns a tuple with the Slippage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlippage

`func (o *BuildSwapRequestDtoV2) SetSlippage(v float32)`

SetSlippage sets Slippage field to given value.


### GetIncludeApprove

`func (o *BuildSwapRequestDtoV2) GetIncludeApprove() bool`

GetIncludeApprove returns the IncludeApprove field if non-nil, zero value otherwise.

### GetIncludeApproveOk

`func (o *BuildSwapRequestDtoV2) GetIncludeApproveOk() (*bool, bool)`

GetIncludeApproveOk returns a tuple with the IncludeApprove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeApprove

`func (o *BuildSwapRequestDtoV2) SetIncludeApprove(v bool)`

SetIncludeApprove sets IncludeApprove field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


