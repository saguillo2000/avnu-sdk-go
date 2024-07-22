# TypedDataMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Values** | **[]map[string]interface{}** |  | 
**IsEmpty** | **bool** |  | 
**Size** | **int32** |  | 
**Entries** | [**[]TypedDataMessageEntriesInner**](TypedDataMessageEntriesInner.md) |  | 
**Keys** | **[]string** |  | 

## Methods

### NewTypedDataMessage

`func NewTypedDataMessage(values []map[string]interface{}, isEmpty bool, size int32, entries []TypedDataMessageEntriesInner, keys []string, ) *TypedDataMessage`

NewTypedDataMessage instantiates a new TypedDataMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypedDataMessageWithDefaults

`func NewTypedDataMessageWithDefaults() *TypedDataMessage`

NewTypedDataMessageWithDefaults instantiates a new TypedDataMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *TypedDataMessage) GetContent() map[string]map[string]interface{}`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *TypedDataMessage) GetContentOk() (*map[string]map[string]interface{}, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *TypedDataMessage) SetContent(v map[string]map[string]interface{})`

SetContent sets Content field to given value.

### HasContent

`func (o *TypedDataMessage) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetValues

`func (o *TypedDataMessage) GetValues() []map[string]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *TypedDataMessage) GetValuesOk() (*[]map[string]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *TypedDataMessage) SetValues(v []map[string]interface{})`

SetValues sets Values field to given value.


### GetIsEmpty

`func (o *TypedDataMessage) GetIsEmpty() bool`

GetIsEmpty returns the IsEmpty field if non-nil, zero value otherwise.

### GetIsEmptyOk

`func (o *TypedDataMessage) GetIsEmptyOk() (*bool, bool)`

GetIsEmptyOk returns a tuple with the IsEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEmpty

`func (o *TypedDataMessage) SetIsEmpty(v bool)`

SetIsEmpty sets IsEmpty field to given value.


### GetSize

`func (o *TypedDataMessage) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *TypedDataMessage) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *TypedDataMessage) SetSize(v int32)`

SetSize sets Size field to given value.


### GetEntries

`func (o *TypedDataMessage) GetEntries() []TypedDataMessageEntriesInner`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *TypedDataMessage) GetEntriesOk() (*[]TypedDataMessageEntriesInner, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *TypedDataMessage) SetEntries(v []TypedDataMessageEntriesInner)`

SetEntries sets Entries field to given value.


### GetKeys

`func (o *TypedDataMessage) GetKeys() []string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *TypedDataMessage) GetKeysOk() (*[]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *TypedDataMessage) SetKeys(v []string)`

SetKeys sets Keys field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


