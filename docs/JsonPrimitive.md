# JsonPrimitive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** |  | 
**IsString** | **bool** |  | 

## Methods

### NewJsonPrimitive

`func NewJsonPrimitive(content string, isString bool, ) *JsonPrimitive`

NewJsonPrimitive instantiates a new JsonPrimitive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonPrimitiveWithDefaults

`func NewJsonPrimitiveWithDefaults() *JsonPrimitive`

NewJsonPrimitiveWithDefaults instantiates a new JsonPrimitive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *JsonPrimitive) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *JsonPrimitive) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *JsonPrimitive) SetContent(v string)`

SetContent sets Content field to given value.


### GetIsString

`func (o *JsonPrimitive) GetIsString() bool`

GetIsString returns the IsString field if non-nil, zero value otherwise.

### GetIsStringOk

`func (o *JsonPrimitive) GetIsStringOk() (*bool, bool)`

GetIsStringOk returns a tuple with the IsString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsString

`func (o *JsonPrimitive) SetIsString(v bool)`

SetIsString sets IsString field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


