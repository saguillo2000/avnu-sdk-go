# RegisterReward

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The user&#39;s address | 
**Campaign** | **string** | The name of the company&#39;s campaign | 
**Protocol** | Pointer to **string** | The protocol where the reward can be used | [optional] 
**FreeTx** | **int64** | The number of free transaction | 
**ExpirationDate** | Pointer to **time.Time** | Reward&#39;s expiration date | [optional] 
**WhitelistedCalls** | [**[]WhitelistedCall**](WhitelistedCall.md) | The list of whitelisted calls | 

## Methods

### NewRegisterReward

`func NewRegisterReward(address string, campaign string, freeTx int64, whitelistedCalls []WhitelistedCall, ) *RegisterReward`

NewRegisterReward instantiates a new RegisterReward object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterRewardWithDefaults

`func NewRegisterRewardWithDefaults() *RegisterReward`

NewRegisterRewardWithDefaults instantiates a new RegisterReward object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *RegisterReward) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *RegisterReward) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *RegisterReward) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetCampaign

`func (o *RegisterReward) GetCampaign() string`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *RegisterReward) GetCampaignOk() (*string, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *RegisterReward) SetCampaign(v string)`

SetCampaign sets Campaign field to given value.


### GetProtocol

`func (o *RegisterReward) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *RegisterReward) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *RegisterReward) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *RegisterReward) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetFreeTx

`func (o *RegisterReward) GetFreeTx() int64`

GetFreeTx returns the FreeTx field if non-nil, zero value otherwise.

### GetFreeTxOk

`func (o *RegisterReward) GetFreeTxOk() (*int64, bool)`

GetFreeTxOk returns a tuple with the FreeTx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeTx

`func (o *RegisterReward) SetFreeTx(v int64)`

SetFreeTx sets FreeTx field to given value.


### GetExpirationDate

`func (o *RegisterReward) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *RegisterReward) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *RegisterReward) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *RegisterReward) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetWhitelistedCalls

`func (o *RegisterReward) GetWhitelistedCalls() []WhitelistedCall`

GetWhitelistedCalls returns the WhitelistedCalls field if non-nil, zero value otherwise.

### GetWhitelistedCallsOk

`func (o *RegisterReward) GetWhitelistedCallsOk() (*[]WhitelistedCall, bool)`

GetWhitelistedCallsOk returns a tuple with the WhitelistedCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelistedCalls

`func (o *RegisterReward) SetWhitelistedCalls(v []WhitelistedCall)`

SetWhitelistedCalls sets WhitelistedCalls field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


