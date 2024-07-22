# Route

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the source | 
**Address** | **string** | The address of the source | 
**Percent** | **float32** | The percentage distribution of sellToken. 1 is 100% | 
**SellTokenAddress** | **string** | The token address user wants to sell | 
**BuyTokenAddress** | **string** | The token address user wants to buy | 

## Methods

### NewRoute

`func NewRoute(name string, address string, percent float32, sellTokenAddress string, buyTokenAddress string, ) *Route`

NewRoute instantiates a new Route object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteWithDefaults

`func NewRouteWithDefaults() *Route`

NewRouteWithDefaults instantiates a new Route object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Route) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Route) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Route) SetName(v string)`

SetName sets Name field to given value.


### GetAddress

`func (o *Route) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Route) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Route) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetPercent

`func (o *Route) GetPercent() float32`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *Route) GetPercentOk() (*float32, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *Route) SetPercent(v float32)`

SetPercent sets Percent field to given value.


### GetSellTokenAddress

`func (o *Route) GetSellTokenAddress() string`

GetSellTokenAddress returns the SellTokenAddress field if non-nil, zero value otherwise.

### GetSellTokenAddressOk

`func (o *Route) GetSellTokenAddressOk() (*string, bool)`

GetSellTokenAddressOk returns a tuple with the SellTokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellTokenAddress

`func (o *Route) SetSellTokenAddress(v string)`

SetSellTokenAddress sets SellTokenAddress field to given value.


### GetBuyTokenAddress

`func (o *Route) GetBuyTokenAddress() string`

GetBuyTokenAddress returns the BuyTokenAddress field if non-nil, zero value otherwise.

### GetBuyTokenAddressOk

`func (o *Route) GetBuyTokenAddressOk() (*string, bool)`

GetBuyTokenAddressOk returns a tuple with the BuyTokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyTokenAddress

`func (o *Route) SetBuyTokenAddress(v string)`

SetBuyTokenAddress sets BuyTokenAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


