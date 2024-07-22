# ExecuteSwapRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuoteId** | **string** | The unique id of the quote | 
**Signature** | **[]string** |  | 

## Methods

### NewExecuteSwapRequest

`func NewExecuteSwapRequest(quoteId string, signature []string, ) *ExecuteSwapRequest`

NewExecuteSwapRequest instantiates a new ExecuteSwapRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecuteSwapRequestWithDefaults

`func NewExecuteSwapRequestWithDefaults() *ExecuteSwapRequest`

NewExecuteSwapRequestWithDefaults instantiates a new ExecuteSwapRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuoteId

`func (o *ExecuteSwapRequest) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *ExecuteSwapRequest) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *ExecuteSwapRequest) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.


### GetSignature

`func (o *ExecuteSwapRequest) GetSignature() []string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *ExecuteSwapRequest) GetSignatureOk() (*[]string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *ExecuteSwapRequest) SetSignature(v []string)`

SetSignature sets Signature field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


