/*
AVNU API

REST API documentation for accessing liquidity on Layer 2.  AVNU is a decentralized exchange protocol enabling the fastest and the most efficient operations in DeFi for Layer 2 with better pricing, zero slippage, MEV-protection and gasless trading.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GasTokenPriceDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GasTokenPriceDto{}

// GasTokenPriceDto struct for GasTokenPriceDto
type GasTokenPriceDto struct {
	// The address of the gas token
	TokenAddress string `json:"tokenAddress"`
	// The estimated amount of gas token to pay the gas fees
	GasFeesInGasToken string `json:"gasFeesInGasToken"`
	// The estimated amount of gas token to pay the gas fees in usd
	GasFeesInUsd float64 `json:"gasFeesInUsd"`
}

type _GasTokenPriceDto GasTokenPriceDto

// NewGasTokenPriceDto instantiates a new GasTokenPriceDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGasTokenPriceDto(tokenAddress string, gasFeesInGasToken string, gasFeesInUsd float64) *GasTokenPriceDto {
	this := GasTokenPriceDto{}
	this.TokenAddress = tokenAddress
	this.GasFeesInGasToken = gasFeesInGasToken
	this.GasFeesInUsd = gasFeesInUsd
	return &this
}

// NewGasTokenPriceDtoWithDefaults instantiates a new GasTokenPriceDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGasTokenPriceDtoWithDefaults() *GasTokenPriceDto {
	this := GasTokenPriceDto{}
	return &this
}

// GetTokenAddress returns the TokenAddress field value
func (o *GasTokenPriceDto) GetTokenAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenAddress
}

// GetTokenAddressOk returns a tuple with the TokenAddress field value
// and a boolean to check if the value has been set.
func (o *GasTokenPriceDto) GetTokenAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenAddress, true
}

// SetTokenAddress sets field value
func (o *GasTokenPriceDto) SetTokenAddress(v string) {
	o.TokenAddress = v
}

// GetGasFeesInGasToken returns the GasFeesInGasToken field value
func (o *GasTokenPriceDto) GetGasFeesInGasToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GasFeesInGasToken
}

// GetGasFeesInGasTokenOk returns a tuple with the GasFeesInGasToken field value
// and a boolean to check if the value has been set.
func (o *GasTokenPriceDto) GetGasFeesInGasTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasFeesInGasToken, true
}

// SetGasFeesInGasToken sets field value
func (o *GasTokenPriceDto) SetGasFeesInGasToken(v string) {
	o.GasFeesInGasToken = v
}

// GetGasFeesInUsd returns the GasFeesInUsd field value
func (o *GasTokenPriceDto) GetGasFeesInUsd() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.GasFeesInUsd
}

// GetGasFeesInUsdOk returns a tuple with the GasFeesInUsd field value
// and a boolean to check if the value has been set.
func (o *GasTokenPriceDto) GetGasFeesInUsdOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasFeesInUsd, true
}

// SetGasFeesInUsd sets field value
func (o *GasTokenPriceDto) SetGasFeesInUsd(v float64) {
	o.GasFeesInUsd = v
}

func (o GasTokenPriceDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GasTokenPriceDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tokenAddress"] = o.TokenAddress
	toSerialize["gasFeesInGasToken"] = o.GasFeesInGasToken
	toSerialize["gasFeesInUsd"] = o.GasFeesInUsd
	return toSerialize, nil
}

func (o *GasTokenPriceDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tokenAddress",
		"gasFeesInGasToken",
		"gasFeesInUsd",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGasTokenPriceDto := _GasTokenPriceDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGasTokenPriceDto)

	if err != nil {
		return err
	}

	*o = GasTokenPriceDto(varGasTokenPriceDto)

	return err
}

type NullableGasTokenPriceDto struct {
	value *GasTokenPriceDto
	isSet bool
}

func (v NullableGasTokenPriceDto) Get() *GasTokenPriceDto {
	return v.value
}

func (v *NullableGasTokenPriceDto) Set(val *GasTokenPriceDto) {
	v.value = val
	v.isSet = true
}

func (v NullableGasTokenPriceDto) IsSet() bool {
	return v.isSet
}

func (v *NullableGasTokenPriceDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGasTokenPriceDto(val *GasTokenPriceDto) *NullableGasTokenPriceDto {
	return &NullableGasTokenPriceDto{value: val, isSet: true}
}

func (v NullableGasTokenPriceDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGasTokenPriceDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


