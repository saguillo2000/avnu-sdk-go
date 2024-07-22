# BuildSwapResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChainId** | **string** | The chain&#39;s id | 
**ContractAddress** | **string** | The address of the contract to send call data to | 
**Entrypoint** | **string** | The entrypoint | 
**Calldata** | **[]string** | The call data required to be sent to the to contract address | 

## Methods

### NewBuildSwapResponse

`func NewBuildSwapResponse(chainId string, contractAddress string, entrypoint string, calldata []string, ) *BuildSwapResponse`

NewBuildSwapResponse instantiates a new BuildSwapResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildSwapResponseWithDefaults

`func NewBuildSwapResponseWithDefaults() *BuildSwapResponse`

NewBuildSwapResponseWithDefaults instantiates a new BuildSwapResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChainId

`func (o *BuildSwapResponse) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *BuildSwapResponse) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *BuildSwapResponse) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetContractAddress

`func (o *BuildSwapResponse) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *BuildSwapResponse) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *BuildSwapResponse) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.


### GetEntrypoint

`func (o *BuildSwapResponse) GetEntrypoint() string`

GetEntrypoint returns the Entrypoint field if non-nil, zero value otherwise.

### GetEntrypointOk

`func (o *BuildSwapResponse) GetEntrypointOk() (*string, bool)`

GetEntrypointOk returns a tuple with the Entrypoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrypoint

`func (o *BuildSwapResponse) SetEntrypoint(v string)`

SetEntrypoint sets Entrypoint field to given value.


### GetCalldata

`func (o *BuildSwapResponse) GetCalldata() []string`

GetCalldata returns the Calldata field if non-nil, zero value otherwise.

### GetCalldataOk

`func (o *BuildSwapResponse) GetCalldataOk() (*[]string, bool)`

GetCalldataOk returns a tuple with the Calldata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalldata

`func (o *BuildSwapResponse) SetCalldata(v []string)`

SetCalldata sets Calldata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


