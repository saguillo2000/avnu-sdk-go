# WhitelistedCall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractAddress** | **string** | The value can be &#39;*&#39; if all contracts are whitelisted or can be the contract address (hex format) | 
**Entrypoint** | **string** | The value can be &#39;*&#39; if all entrypoint are whitelisted or can be the entrypoint name (string format) | 

## Methods

### NewWhitelistedCall

`func NewWhitelistedCall(contractAddress string, entrypoint string, ) *WhitelistedCall`

NewWhitelistedCall instantiates a new WhitelistedCall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhitelistedCallWithDefaults

`func NewWhitelistedCallWithDefaults() *WhitelistedCall`

NewWhitelistedCallWithDefaults instantiates a new WhitelistedCall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractAddress

`func (o *WhitelistedCall) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *WhitelistedCall) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *WhitelistedCall) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.


### GetEntrypoint

`func (o *WhitelistedCall) GetEntrypoint() string`

GetEntrypoint returns the Entrypoint field if non-nil, zero value otherwise.

### GetEntrypointOk

`func (o *WhitelistedCall) GetEntrypointOk() (*string, bool)`

GetEntrypointOk returns a tuple with the Entrypoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrypoint

`func (o *WhitelistedCall) SetEntrypoint(v string)`

SetEntrypoint sets Entrypoint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


