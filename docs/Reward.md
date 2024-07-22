# Reward

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **time.Time** | Reward&#39;s creation date | 
**Address** | **string** | The user&#39;s address | 
**Sponsor** | **string** | The company that will pay the gas fees | 
**Campaign** | **string** | The name of the company&#39;s campaign | 
**Protocol** | Pointer to **string** | The protocol where the reward can be used | [optional] 
**FreeTx** | **int64** | The number of free transaction | 
**RemainingTx** | **int64** | The number of remaining transactions | 
**ExpirationDate** | Pointer to **time.Time** | Reward&#39;s expiration date | [optional] 
**WhitelistedCalls** | [**[]WhitelistedCall**](WhitelistedCall.md) | The list of whitelisted calls | 

## Methods

### NewReward

`func NewReward(date time.Time, address string, sponsor string, campaign string, freeTx int64, remainingTx int64, whitelistedCalls []WhitelistedCall, ) *Reward`

NewReward instantiates a new Reward object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardWithDefaults

`func NewRewardWithDefaults() *Reward`

NewRewardWithDefaults instantiates a new Reward object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *Reward) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Reward) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Reward) SetDate(v time.Time)`

SetDate sets Date field to given value.


### GetAddress

`func (o *Reward) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Reward) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Reward) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetSponsor

`func (o *Reward) GetSponsor() string`

GetSponsor returns the Sponsor field if non-nil, zero value otherwise.

### GetSponsorOk

`func (o *Reward) GetSponsorOk() (*string, bool)`

GetSponsorOk returns a tuple with the Sponsor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsor

`func (o *Reward) SetSponsor(v string)`

SetSponsor sets Sponsor field to given value.


### GetCampaign

`func (o *Reward) GetCampaign() string`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *Reward) GetCampaignOk() (*string, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *Reward) SetCampaign(v string)`

SetCampaign sets Campaign field to given value.


### GetProtocol

`func (o *Reward) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *Reward) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *Reward) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *Reward) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetFreeTx

`func (o *Reward) GetFreeTx() int64`

GetFreeTx returns the FreeTx field if non-nil, zero value otherwise.

### GetFreeTxOk

`func (o *Reward) GetFreeTxOk() (*int64, bool)`

GetFreeTxOk returns a tuple with the FreeTx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeTx

`func (o *Reward) SetFreeTx(v int64)`

SetFreeTx sets FreeTx field to given value.


### GetRemainingTx

`func (o *Reward) GetRemainingTx() int64`

GetRemainingTx returns the RemainingTx field if non-nil, zero value otherwise.

### GetRemainingTxOk

`func (o *Reward) GetRemainingTxOk() (*int64, bool)`

GetRemainingTxOk returns a tuple with the RemainingTx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingTx

`func (o *Reward) SetRemainingTx(v int64)`

SetRemainingTx sets RemainingTx field to given value.


### GetExpirationDate

`func (o *Reward) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *Reward) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *Reward) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *Reward) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetWhitelistedCalls

`func (o *Reward) GetWhitelistedCalls() []WhitelistedCall`

GetWhitelistedCalls returns the WhitelistedCalls field if non-nil, zero value otherwise.

### GetWhitelistedCallsOk

`func (o *Reward) GetWhitelistedCallsOk() (*[]WhitelistedCall, bool)`

GetWhitelistedCallsOk returns a tuple with the WhitelistedCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelistedCalls

`func (o *Reward) SetWhitelistedCalls(v []WhitelistedCall)`

SetWhitelistedCalls sets WhitelistedCalls field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


