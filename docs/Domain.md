# Domain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | [**JsonPrimitive**](JsonPrimitive.md) |  | 
**Version** | [**JsonPrimitive**](JsonPrimitive.md) |  | 
**ChainId** | [**JsonPrimitive**](JsonPrimitive.md) |  | 
**Revision** | Pointer to **string** |  | [optional] 
**SeparatorNamelib** | **string** |  | 

## Methods

### NewDomain

`func NewDomain(name JsonPrimitive, version JsonPrimitive, chainId JsonPrimitive, separatorNamelib string, ) *Domain`

NewDomain instantiates a new Domain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainWithDefaults

`func NewDomainWithDefaults() *Domain`

NewDomainWithDefaults instantiates a new Domain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Domain) GetName() JsonPrimitive`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Domain) GetNameOk() (*JsonPrimitive, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Domain) SetName(v JsonPrimitive)`

SetName sets Name field to given value.


### GetVersion

`func (o *Domain) GetVersion() JsonPrimitive`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Domain) GetVersionOk() (*JsonPrimitive, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Domain) SetVersion(v JsonPrimitive)`

SetVersion sets Version field to given value.


### GetChainId

`func (o *Domain) GetChainId() JsonPrimitive`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *Domain) GetChainIdOk() (*JsonPrimitive, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *Domain) SetChainId(v JsonPrimitive)`

SetChainId sets ChainId field to given value.


### GetRevision

`func (o *Domain) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *Domain) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *Domain) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *Domain) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSeparatorNamelib

`func (o *Domain) GetSeparatorNamelib() string`

GetSeparatorNamelib returns the SeparatorNamelib field if non-nil, zero value otherwise.

### GetSeparatorNamelibOk

`func (o *Domain) GetSeparatorNamelibOk() (*string, bool)`

GetSeparatorNamelibOk returns a tuple with the SeparatorNamelib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeparatorNamelib

`func (o *Domain) SetSeparatorNamelib(v string)`

SetSeparatorNamelib sets SeparatorNamelib field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


