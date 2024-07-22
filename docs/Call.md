# Call

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractAddress** | **string** | Contract&#39;s address | 
**Entrypoint** | **string** | Entrypoint (string format) | 
**Calldata** | **[]string** | calldata (list of felt) | 

## Methods

### NewCall

`func NewCall(contractAddress string, entrypoint string, calldata []string, ) *Call`

NewCall instantiates a new Call object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallWithDefaults

`func NewCallWithDefaults() *Call`

NewCallWithDefaults instantiates a new Call object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractAddress

`func (o *Call) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *Call) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *Call) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.


### GetEntrypoint

`func (o *Call) GetEntrypoint() string`

GetEntrypoint returns the Entrypoint field if non-nil, zero value otherwise.

### GetEntrypointOk

`func (o *Call) GetEntrypointOk() (*string, bool)`

GetEntrypointOk returns a tuple with the Entrypoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrypoint

`func (o *Call) SetEntrypoint(v string)`

SetEntrypoint sets Entrypoint field to given value.


### GetCalldata

`func (o *Call) GetCalldata() []string`

GetCalldata returns the Calldata field if non-nil, zero value otherwise.

### GetCalldataOk

`func (o *Call) GetCalldataOk() (*[]string, bool)`

GetCalldataOk returns a tuple with the Calldata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalldata

`func (o *Call) SetCalldata(v []string)`

SetCalldata sets Calldata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


