# Quote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuoteId** | **string** | The unique id of the quote | 
**SellTokenAddress** | **string** | The token address user wants to sell | 
**SellAmount** | **string** | The amount of sellAmount that would be sold in this swap | 
**SellAmountInUsd** | **float64** | The amount of sellAmount that would be sold in this swap in usd | 
**BuyTokenAddress** | **string** | The token address user wants to buy | 
**BuyAmount** | **string** | The amount of buyToken that would be bought in this swap | 
**BuyAmountInUsd** | **float64** | The amount of buyToken that would be bought in this swap in usd | 
**BuyAmountWithoutFees** | **string** | The amount of buyToken without fees | 
**BuyAmountWithoutFeesInUsd** | **float64** | The amount of buyToken without fees in usd | 
**EstimatedAmount** | **bool** | If true, user should define a slippage when he executes the quote | 
**ChainId** | **string** | The chain&#39;s id | 
**BlockNumber** | Pointer to **string** | Defined when quote comes from a DEX | [optional] 
**Expiry** | Pointer to **float32** | Unix timestamp when quotes expires in seconds | [optional] 
**Routes** | [**[]Route**](Route.md) |  | 
**GasFees** | **string** | The estimated amount of gas fees | 
**GasFeesInUsd** | Pointer to **float64** | The estimated amount of gas fees in USD | [optional] 
**AvnuFees** | **string** | The actual fees taken by AVNU | 
**AvnuFeesInUsd** | **float64** | The actual fees taken by AVNU is usd | 
**AvnuFeesBps** | **string** | The fees in bps taken by AVNU | 
**IntegratorFees** | **string** | The actual fees taken by the integrator | 
**IntegratorFeesInUsd** | **float64** | The actual fees taken by the integrator in usd | 
**IntegratorFeesBps** | **string** | The fees in bps taken by the integrator | 
**PriceRatioUsd** | **float32** | Price ratio in usd and in bps | 
**LiquiditySource** | **string** | The type of liquidity source | 
**SellTokenPriceInUsd** | Pointer to **float64** | The amount in USD to buy 1 sellToken | [optional] 
**BuyTokenPriceInUsd** | Pointer to **float64** | The amount in USD to buy 1 buyToken | [optional] 
**Gasless** | [**Gasless**](Gasless.md) |  | 

## Methods

### NewQuote

`func NewQuote(quoteId string, sellTokenAddress string, sellAmount string, sellAmountInUsd float64, buyTokenAddress string, buyAmount string, buyAmountInUsd float64, buyAmountWithoutFees string, buyAmountWithoutFeesInUsd float64, estimatedAmount bool, chainId string, routes []Route, gasFees string, avnuFees string, avnuFeesInUsd float64, avnuFeesBps string, integratorFees string, integratorFeesInUsd float64, integratorFeesBps string, priceRatioUsd float32, liquiditySource string, gasless Gasless, ) *Quote`

NewQuote instantiates a new Quote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteWithDefaults

`func NewQuoteWithDefaults() *Quote`

NewQuoteWithDefaults instantiates a new Quote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuoteId

`func (o *Quote) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *Quote) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *Quote) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.


### GetSellTokenAddress

`func (o *Quote) GetSellTokenAddress() string`

GetSellTokenAddress returns the SellTokenAddress field if non-nil, zero value otherwise.

### GetSellTokenAddressOk

`func (o *Quote) GetSellTokenAddressOk() (*string, bool)`

GetSellTokenAddressOk returns a tuple with the SellTokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellTokenAddress

`func (o *Quote) SetSellTokenAddress(v string)`

SetSellTokenAddress sets SellTokenAddress field to given value.


### GetSellAmount

`func (o *Quote) GetSellAmount() string`

GetSellAmount returns the SellAmount field if non-nil, zero value otherwise.

### GetSellAmountOk

`func (o *Quote) GetSellAmountOk() (*string, bool)`

GetSellAmountOk returns a tuple with the SellAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellAmount

`func (o *Quote) SetSellAmount(v string)`

SetSellAmount sets SellAmount field to given value.


### GetSellAmountInUsd

`func (o *Quote) GetSellAmountInUsd() float64`

GetSellAmountInUsd returns the SellAmountInUsd field if non-nil, zero value otherwise.

### GetSellAmountInUsdOk

`func (o *Quote) GetSellAmountInUsdOk() (*float64, bool)`

GetSellAmountInUsdOk returns a tuple with the SellAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellAmountInUsd

`func (o *Quote) SetSellAmountInUsd(v float64)`

SetSellAmountInUsd sets SellAmountInUsd field to given value.


### GetBuyTokenAddress

`func (o *Quote) GetBuyTokenAddress() string`

GetBuyTokenAddress returns the BuyTokenAddress field if non-nil, zero value otherwise.

### GetBuyTokenAddressOk

`func (o *Quote) GetBuyTokenAddressOk() (*string, bool)`

GetBuyTokenAddressOk returns a tuple with the BuyTokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyTokenAddress

`func (o *Quote) SetBuyTokenAddress(v string)`

SetBuyTokenAddress sets BuyTokenAddress field to given value.


### GetBuyAmount

`func (o *Quote) GetBuyAmount() string`

GetBuyAmount returns the BuyAmount field if non-nil, zero value otherwise.

### GetBuyAmountOk

`func (o *Quote) GetBuyAmountOk() (*string, bool)`

GetBuyAmountOk returns a tuple with the BuyAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyAmount

`func (o *Quote) SetBuyAmount(v string)`

SetBuyAmount sets BuyAmount field to given value.


### GetBuyAmountInUsd

`func (o *Quote) GetBuyAmountInUsd() float64`

GetBuyAmountInUsd returns the BuyAmountInUsd field if non-nil, zero value otherwise.

### GetBuyAmountInUsdOk

`func (o *Quote) GetBuyAmountInUsdOk() (*float64, bool)`

GetBuyAmountInUsdOk returns a tuple with the BuyAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyAmountInUsd

`func (o *Quote) SetBuyAmountInUsd(v float64)`

SetBuyAmountInUsd sets BuyAmountInUsd field to given value.


### GetBuyAmountWithoutFees

`func (o *Quote) GetBuyAmountWithoutFees() string`

GetBuyAmountWithoutFees returns the BuyAmountWithoutFees field if non-nil, zero value otherwise.

### GetBuyAmountWithoutFeesOk

`func (o *Quote) GetBuyAmountWithoutFeesOk() (*string, bool)`

GetBuyAmountWithoutFeesOk returns a tuple with the BuyAmountWithoutFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyAmountWithoutFees

`func (o *Quote) SetBuyAmountWithoutFees(v string)`

SetBuyAmountWithoutFees sets BuyAmountWithoutFees field to given value.


### GetBuyAmountWithoutFeesInUsd

`func (o *Quote) GetBuyAmountWithoutFeesInUsd() float64`

GetBuyAmountWithoutFeesInUsd returns the BuyAmountWithoutFeesInUsd field if non-nil, zero value otherwise.

### GetBuyAmountWithoutFeesInUsdOk

`func (o *Quote) GetBuyAmountWithoutFeesInUsdOk() (*float64, bool)`

GetBuyAmountWithoutFeesInUsdOk returns a tuple with the BuyAmountWithoutFeesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyAmountWithoutFeesInUsd

`func (o *Quote) SetBuyAmountWithoutFeesInUsd(v float64)`

SetBuyAmountWithoutFeesInUsd sets BuyAmountWithoutFeesInUsd field to given value.


### GetEstimatedAmount

`func (o *Quote) GetEstimatedAmount() bool`

GetEstimatedAmount returns the EstimatedAmount field if non-nil, zero value otherwise.

### GetEstimatedAmountOk

`func (o *Quote) GetEstimatedAmountOk() (*bool, bool)`

GetEstimatedAmountOk returns a tuple with the EstimatedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedAmount

`func (o *Quote) SetEstimatedAmount(v bool)`

SetEstimatedAmount sets EstimatedAmount field to given value.


### GetChainId

`func (o *Quote) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *Quote) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *Quote) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetBlockNumber

`func (o *Quote) GetBlockNumber() string`

GetBlockNumber returns the BlockNumber field if non-nil, zero value otherwise.

### GetBlockNumberOk

`func (o *Quote) GetBlockNumberOk() (*string, bool)`

GetBlockNumberOk returns a tuple with the BlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNumber

`func (o *Quote) SetBlockNumber(v string)`

SetBlockNumber sets BlockNumber field to given value.

### HasBlockNumber

`func (o *Quote) HasBlockNumber() bool`

HasBlockNumber returns a boolean if a field has been set.

### GetExpiry

`func (o *Quote) GetExpiry() float32`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *Quote) GetExpiryOk() (*float32, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *Quote) SetExpiry(v float32)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *Quote) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetRoutes

`func (o *Quote) GetRoutes() []Route`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *Quote) GetRoutesOk() (*[]Route, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *Quote) SetRoutes(v []Route)`

SetRoutes sets Routes field to given value.


### GetGasFees

`func (o *Quote) GetGasFees() string`

GetGasFees returns the GasFees field if non-nil, zero value otherwise.

### GetGasFeesOk

`func (o *Quote) GetGasFeesOk() (*string, bool)`

GetGasFeesOk returns a tuple with the GasFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasFees

`func (o *Quote) SetGasFees(v string)`

SetGasFees sets GasFees field to given value.


### GetGasFeesInUsd

`func (o *Quote) GetGasFeesInUsd() float64`

GetGasFeesInUsd returns the GasFeesInUsd field if non-nil, zero value otherwise.

### GetGasFeesInUsdOk

`func (o *Quote) GetGasFeesInUsdOk() (*float64, bool)`

GetGasFeesInUsdOk returns a tuple with the GasFeesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasFeesInUsd

`func (o *Quote) SetGasFeesInUsd(v float64)`

SetGasFeesInUsd sets GasFeesInUsd field to given value.

### HasGasFeesInUsd

`func (o *Quote) HasGasFeesInUsd() bool`

HasGasFeesInUsd returns a boolean if a field has been set.

### GetAvnuFees

`func (o *Quote) GetAvnuFees() string`

GetAvnuFees returns the AvnuFees field if non-nil, zero value otherwise.

### GetAvnuFeesOk

`func (o *Quote) GetAvnuFeesOk() (*string, bool)`

GetAvnuFeesOk returns a tuple with the AvnuFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvnuFees

`func (o *Quote) SetAvnuFees(v string)`

SetAvnuFees sets AvnuFees field to given value.


### GetAvnuFeesInUsd

`func (o *Quote) GetAvnuFeesInUsd() float64`

GetAvnuFeesInUsd returns the AvnuFeesInUsd field if non-nil, zero value otherwise.

### GetAvnuFeesInUsdOk

`func (o *Quote) GetAvnuFeesInUsdOk() (*float64, bool)`

GetAvnuFeesInUsdOk returns a tuple with the AvnuFeesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvnuFeesInUsd

`func (o *Quote) SetAvnuFeesInUsd(v float64)`

SetAvnuFeesInUsd sets AvnuFeesInUsd field to given value.


### GetAvnuFeesBps

`func (o *Quote) GetAvnuFeesBps() string`

GetAvnuFeesBps returns the AvnuFeesBps field if non-nil, zero value otherwise.

### GetAvnuFeesBpsOk

`func (o *Quote) GetAvnuFeesBpsOk() (*string, bool)`

GetAvnuFeesBpsOk returns a tuple with the AvnuFeesBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvnuFeesBps

`func (o *Quote) SetAvnuFeesBps(v string)`

SetAvnuFeesBps sets AvnuFeesBps field to given value.


### GetIntegratorFees

`func (o *Quote) GetIntegratorFees() string`

GetIntegratorFees returns the IntegratorFees field if non-nil, zero value otherwise.

### GetIntegratorFeesOk

`func (o *Quote) GetIntegratorFeesOk() (*string, bool)`

GetIntegratorFeesOk returns a tuple with the IntegratorFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegratorFees

`func (o *Quote) SetIntegratorFees(v string)`

SetIntegratorFees sets IntegratorFees field to given value.


### GetIntegratorFeesInUsd

`func (o *Quote) GetIntegratorFeesInUsd() float64`

GetIntegratorFeesInUsd returns the IntegratorFeesInUsd field if non-nil, zero value otherwise.

### GetIntegratorFeesInUsdOk

`func (o *Quote) GetIntegratorFeesInUsdOk() (*float64, bool)`

GetIntegratorFeesInUsdOk returns a tuple with the IntegratorFeesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegratorFeesInUsd

`func (o *Quote) SetIntegratorFeesInUsd(v float64)`

SetIntegratorFeesInUsd sets IntegratorFeesInUsd field to given value.


### GetIntegratorFeesBps

`func (o *Quote) GetIntegratorFeesBps() string`

GetIntegratorFeesBps returns the IntegratorFeesBps field if non-nil, zero value otherwise.

### GetIntegratorFeesBpsOk

`func (o *Quote) GetIntegratorFeesBpsOk() (*string, bool)`

GetIntegratorFeesBpsOk returns a tuple with the IntegratorFeesBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegratorFeesBps

`func (o *Quote) SetIntegratorFeesBps(v string)`

SetIntegratorFeesBps sets IntegratorFeesBps field to given value.


### GetPriceRatioUsd

`func (o *Quote) GetPriceRatioUsd() float32`

GetPriceRatioUsd returns the PriceRatioUsd field if non-nil, zero value otherwise.

### GetPriceRatioUsdOk

`func (o *Quote) GetPriceRatioUsdOk() (*float32, bool)`

GetPriceRatioUsdOk returns a tuple with the PriceRatioUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceRatioUsd

`func (o *Quote) SetPriceRatioUsd(v float32)`

SetPriceRatioUsd sets PriceRatioUsd field to given value.


### GetLiquiditySource

`func (o *Quote) GetLiquiditySource() string`

GetLiquiditySource returns the LiquiditySource field if non-nil, zero value otherwise.

### GetLiquiditySourceOk

`func (o *Quote) GetLiquiditySourceOk() (*string, bool)`

GetLiquiditySourceOk returns a tuple with the LiquiditySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquiditySource

`func (o *Quote) SetLiquiditySource(v string)`

SetLiquiditySource sets LiquiditySource field to given value.


### GetSellTokenPriceInUsd

`func (o *Quote) GetSellTokenPriceInUsd() float64`

GetSellTokenPriceInUsd returns the SellTokenPriceInUsd field if non-nil, zero value otherwise.

### GetSellTokenPriceInUsdOk

`func (o *Quote) GetSellTokenPriceInUsdOk() (*float64, bool)`

GetSellTokenPriceInUsdOk returns a tuple with the SellTokenPriceInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellTokenPriceInUsd

`func (o *Quote) SetSellTokenPriceInUsd(v float64)`

SetSellTokenPriceInUsd sets SellTokenPriceInUsd field to given value.

### HasSellTokenPriceInUsd

`func (o *Quote) HasSellTokenPriceInUsd() bool`

HasSellTokenPriceInUsd returns a boolean if a field has been set.

### GetBuyTokenPriceInUsd

`func (o *Quote) GetBuyTokenPriceInUsd() float64`

GetBuyTokenPriceInUsd returns the BuyTokenPriceInUsd field if non-nil, zero value otherwise.

### GetBuyTokenPriceInUsdOk

`func (o *Quote) GetBuyTokenPriceInUsdOk() (*float64, bool)`

GetBuyTokenPriceInUsdOk returns a tuple with the BuyTokenPriceInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyTokenPriceInUsd

`func (o *Quote) SetBuyTokenPriceInUsd(v float64)`

SetBuyTokenPriceInUsd sets BuyTokenPriceInUsd field to given value.

### HasBuyTokenPriceInUsd

`func (o *Quote) HasBuyTokenPriceInUsd() bool`

HasBuyTokenPriceInUsd returns a boolean if a field has been set.

### GetGasless

`func (o *Quote) GetGasless() Gasless`

GetGasless returns the Gasless field if non-nil, zero value otherwise.

### GetGaslessOk

`func (o *Quote) GetGaslessOk() (*Gasless, bool)`

GetGaslessOk returns a tuple with the Gasless field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasless

`func (o *Quote) SetGasless(v Gasless)`

SetGasless sets Gasless field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


