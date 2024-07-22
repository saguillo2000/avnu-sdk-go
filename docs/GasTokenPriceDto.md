# GasTokenPriceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenAddress** | **string** | The address of the gas token | 
**GasFeesInGasToken** | **string** | The estimated amount of gas token to pay the gas fees | 
**GasFeesInUsd** | **float64** | The estimated amount of gas token to pay the gas fees in usd | 

## Methods

### NewGasTokenPriceDto

`func NewGasTokenPriceDto(tokenAddress string, gasFeesInGasToken string, gasFeesInUsd float64, ) *GasTokenPriceDto`

NewGasTokenPriceDto instantiates a new GasTokenPriceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGasTokenPriceDtoWithDefaults

`func NewGasTokenPriceDtoWithDefaults() *GasTokenPriceDto`

NewGasTokenPriceDtoWithDefaults instantiates a new GasTokenPriceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenAddress

`func (o *GasTokenPriceDto) GetTokenAddress() string`

GetTokenAddress returns the TokenAddress field if non-nil, zero value otherwise.

### GetTokenAddressOk

`func (o *GasTokenPriceDto) GetTokenAddressOk() (*string, bool)`

GetTokenAddressOk returns a tuple with the TokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAddress

`func (o *GasTokenPriceDto) SetTokenAddress(v string)`

SetTokenAddress sets TokenAddress field to given value.


### GetGasFeesInGasToken

`func (o *GasTokenPriceDto) GetGasFeesInGasToken() string`

GetGasFeesInGasToken returns the GasFeesInGasToken field if non-nil, zero value otherwise.

### GetGasFeesInGasTokenOk

`func (o *GasTokenPriceDto) GetGasFeesInGasTokenOk() (*string, bool)`

GetGasFeesInGasTokenOk returns a tuple with the GasFeesInGasToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasFeesInGasToken

`func (o *GasTokenPriceDto) SetGasFeesInGasToken(v string)`

SetGasFeesInGasToken sets GasFeesInGasToken field to given value.


### GetGasFeesInUsd

`func (o *GasTokenPriceDto) GetGasFeesInUsd() float64`

GetGasFeesInUsd returns the GasFeesInUsd field if non-nil, zero value otherwise.

### GetGasFeesInUsdOk

`func (o *GasTokenPriceDto) GetGasFeesInUsdOk() (*float64, bool)`

GetGasFeesInUsdOk returns a tuple with the GasFeesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasFeesInUsd

`func (o *GasTokenPriceDto) SetGasFeesInUsd(v float64)`

SetGasFeesInUsd sets GasFeesInUsd field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


