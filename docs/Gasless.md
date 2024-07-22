# Gasless

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** | If true, the quote can be executed using gasless | 
**GasTokenPrices** | [**[]GasTokenPriceDto**](GasTokenPriceDto.md) |  | 

## Methods

### NewGasless

`func NewGasless(active bool, gasTokenPrices []GasTokenPriceDto, ) *Gasless`

NewGasless instantiates a new Gasless object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGaslessWithDefaults

`func NewGaslessWithDefaults() *Gasless`

NewGaslessWithDefaults instantiates a new Gasless object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *Gasless) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Gasless) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Gasless) SetActive(v bool)`

SetActive sets Active field to given value.


### GetGasTokenPrices

`func (o *Gasless) GetGasTokenPrices() []GasTokenPriceDto`

GetGasTokenPrices returns the GasTokenPrices field if non-nil, zero value otherwise.

### GetGasTokenPricesOk

`func (o *Gasless) GetGasTokenPricesOk() (*[]GasTokenPriceDto, bool)`

GetGasTokenPricesOk returns a tuple with the GasTokenPrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasTokenPrices

`func (o *Gasless) SetGasTokenPrices(v []GasTokenPriceDto)`

SetGasTokenPrices sets GasTokenPrices field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


