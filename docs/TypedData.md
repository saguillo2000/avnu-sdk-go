# TypedData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Types** | [**map[string][]Type**](array.md) |  | 
**PrimaryType** | **string** |  | 
**Domain** | [**Domain**](Domain.md) |  | 
**Message** | [**TypedDataMessage**](TypedDataMessage.md) |  | 

## Methods

### NewTypedData

`func NewTypedData(types map[string][]Type, primaryType string, domain Domain, message TypedDataMessage, ) *TypedData`

NewTypedData instantiates a new TypedData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypedDataWithDefaults

`func NewTypedDataWithDefaults() *TypedData`

NewTypedDataWithDefaults instantiates a new TypedData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypes

`func (o *TypedData) GetTypes() map[string][]Type`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *TypedData) GetTypesOk() (*map[string][]Type, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *TypedData) SetTypes(v map[string][]Type)`

SetTypes sets Types field to given value.


### GetPrimaryType

`func (o *TypedData) GetPrimaryType() string`

GetPrimaryType returns the PrimaryType field if non-nil, zero value otherwise.

### GetPrimaryTypeOk

`func (o *TypedData) GetPrimaryTypeOk() (*string, bool)`

GetPrimaryTypeOk returns a tuple with the PrimaryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryType

`func (o *TypedData) SetPrimaryType(v string)`

SetPrimaryType sets PrimaryType field to given value.


### GetDomain

`func (o *TypedData) GetDomain() Domain`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *TypedData) GetDomainOk() (*Domain, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *TypedData) SetDomain(v Domain)`

SetDomain sets Domain field to given value.


### GetMessage

`func (o *TypedData) GetMessage() TypedDataMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TypedData) GetMessageOk() (*TypedDataMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TypedData) SetMessage(v TypedDataMessage)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


