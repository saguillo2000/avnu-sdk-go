# ErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | **[]string** |  | 
**RevertError** | Pointer to **string** |  | [optional] 

## Methods

### NewErrorResponse

`func NewErrorResponse(messages []string, ) *ErrorResponse`

NewErrorResponse instantiates a new ErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResponseWithDefaults

`func NewErrorResponseWithDefaults() *ErrorResponse`

NewErrorResponseWithDefaults instantiates a new ErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *ErrorResponse) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ErrorResponse) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ErrorResponse) SetMessages(v []string)`

SetMessages sets Messages field to given value.


### GetRevertError

`func (o *ErrorResponse) GetRevertError() string`

GetRevertError returns the RevertError field if non-nil, zero value otherwise.

### GetRevertErrorOk

`func (o *ErrorResponse) GetRevertErrorOk() (*string, bool)`

GetRevertErrorOk returns a tuple with the RevertError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevertError

`func (o *ErrorResponse) SetRevertError(v string)`

SetRevertError sets RevertError field to given value.

### HasRevertError

`func (o *ErrorResponse) HasRevertError() bool`

HasRevertError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


