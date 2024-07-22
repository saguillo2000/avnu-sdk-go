# SponsorActivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The sponsor&#39;s name | 
**SucceededTxCount** | **int64** | The number of succeeded transactions | 
**RevertedTxCount** | **int64** | The number of reverted transactions | 
**TxCount** | **int64** | The total number executed transactions | 
**SucceededGasFees** | **string** | The amount of ETH paid for succeeded transactions | 
**RevertedGasFees** | **string** | The amount of ETH paid for reverted transactions | 
**GasFees** | **string** | The amount of ETH paid for all transactions | 
**RemainingCredits** | **string** | The remaining credits. When zero, please contact us so you can recharge | 

## Methods

### NewSponsorActivity

`func NewSponsorActivity(name string, succeededTxCount int64, revertedTxCount int64, txCount int64, succeededGasFees string, revertedGasFees string, gasFees string, remainingCredits string, ) *SponsorActivity`

NewSponsorActivity instantiates a new SponsorActivity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSponsorActivityWithDefaults

`func NewSponsorActivityWithDefaults() *SponsorActivity`

NewSponsorActivityWithDefaults instantiates a new SponsorActivity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SponsorActivity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SponsorActivity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SponsorActivity) SetName(v string)`

SetName sets Name field to given value.


### GetSucceededTxCount

`func (o *SponsorActivity) GetSucceededTxCount() int64`

GetSucceededTxCount returns the SucceededTxCount field if non-nil, zero value otherwise.

### GetSucceededTxCountOk

`func (o *SponsorActivity) GetSucceededTxCountOk() (*int64, bool)`

GetSucceededTxCountOk returns a tuple with the SucceededTxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceededTxCount

`func (o *SponsorActivity) SetSucceededTxCount(v int64)`

SetSucceededTxCount sets SucceededTxCount field to given value.


### GetRevertedTxCount

`func (o *SponsorActivity) GetRevertedTxCount() int64`

GetRevertedTxCount returns the RevertedTxCount field if non-nil, zero value otherwise.

### GetRevertedTxCountOk

`func (o *SponsorActivity) GetRevertedTxCountOk() (*int64, bool)`

GetRevertedTxCountOk returns a tuple with the RevertedTxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevertedTxCount

`func (o *SponsorActivity) SetRevertedTxCount(v int64)`

SetRevertedTxCount sets RevertedTxCount field to given value.


### GetTxCount

`func (o *SponsorActivity) GetTxCount() int64`

GetTxCount returns the TxCount field if non-nil, zero value otherwise.

### GetTxCountOk

`func (o *SponsorActivity) GetTxCountOk() (*int64, bool)`

GetTxCountOk returns a tuple with the TxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxCount

`func (o *SponsorActivity) SetTxCount(v int64)`

SetTxCount sets TxCount field to given value.


### GetSucceededGasFees

`func (o *SponsorActivity) GetSucceededGasFees() string`

GetSucceededGasFees returns the SucceededGasFees field if non-nil, zero value otherwise.

### GetSucceededGasFeesOk

`func (o *SponsorActivity) GetSucceededGasFeesOk() (*string, bool)`

GetSucceededGasFeesOk returns a tuple with the SucceededGasFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceededGasFees

`func (o *SponsorActivity) SetSucceededGasFees(v string)`

SetSucceededGasFees sets SucceededGasFees field to given value.


### GetRevertedGasFees

`func (o *SponsorActivity) GetRevertedGasFees() string`

GetRevertedGasFees returns the RevertedGasFees field if non-nil, zero value otherwise.

### GetRevertedGasFeesOk

`func (o *SponsorActivity) GetRevertedGasFeesOk() (*string, bool)`

GetRevertedGasFeesOk returns a tuple with the RevertedGasFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevertedGasFees

`func (o *SponsorActivity) SetRevertedGasFees(v string)`

SetRevertedGasFees sets RevertedGasFees field to given value.


### GetGasFees

`func (o *SponsorActivity) GetGasFees() string`

GetGasFees returns the GasFees field if non-nil, zero value otherwise.

### GetGasFeesOk

`func (o *SponsorActivity) GetGasFeesOk() (*string, bool)`

GetGasFeesOk returns a tuple with the GasFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasFees

`func (o *SponsorActivity) SetGasFees(v string)`

SetGasFees sets GasFees field to given value.


### GetRemainingCredits

`func (o *SponsorActivity) GetRemainingCredits() string`

GetRemainingCredits returns the RemainingCredits field if non-nil, zero value otherwise.

### GetRemainingCreditsOk

`func (o *SponsorActivity) GetRemainingCreditsOk() (*string, bool)`

GetRemainingCreditsOk returns a tuple with the RemainingCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingCredits

`func (o *SponsorActivity) SetRemainingCredits(v string)`

SetRemainingCredits sets RemainingCredits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


