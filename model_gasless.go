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

// checks if the Gasless type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Gasless{}

// Gasless struct for Gasless
type Gasless struct {
	// If true, the quote can be executed using gasless
	Active bool `json:"active"`
	GasTokenPrices []GasTokenPriceDto `json:"gasTokenPrices"`
}

type _Gasless Gasless

// NewGasless instantiates a new Gasless object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGasless(active bool, gasTokenPrices []GasTokenPriceDto) *Gasless {
	this := Gasless{}
	this.Active = active
	this.GasTokenPrices = gasTokenPrices
	return &this
}

// NewGaslessWithDefaults instantiates a new Gasless object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGaslessWithDefaults() *Gasless {
	this := Gasless{}
	return &this
}

// GetActive returns the Active field value
func (o *Gasless) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *Gasless) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *Gasless) SetActive(v bool) {
	o.Active = v
}

// GetGasTokenPrices returns the GasTokenPrices field value
func (o *Gasless) GetGasTokenPrices() []GasTokenPriceDto {
	if o == nil {
		var ret []GasTokenPriceDto
		return ret
	}

	return o.GasTokenPrices
}

// GetGasTokenPricesOk returns a tuple with the GasTokenPrices field value
// and a boolean to check if the value has been set.
func (o *Gasless) GetGasTokenPricesOk() ([]GasTokenPriceDto, bool) {
	if o == nil {
		return nil, false
	}
	return o.GasTokenPrices, true
}

// SetGasTokenPrices sets field value
func (o *Gasless) SetGasTokenPrices(v []GasTokenPriceDto) {
	o.GasTokenPrices = v
}

func (o Gasless) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Gasless) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["active"] = o.Active
	toSerialize["gasTokenPrices"] = o.GasTokenPrices
	return toSerialize, nil
}

func (o *Gasless) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"active",
		"gasTokenPrices",
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

	varGasless := _Gasless{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGasless)

	if err != nil {
		return err
	}

	*o = Gasless(varGasless)

	return err
}

type NullableGasless struct {
	value *Gasless
	isSet bool
}

func (v NullableGasless) Get() *Gasless {
	return v.value
}

func (v *NullableGasless) Set(val *Gasless) {
	v.value = val
	v.isSet = true
}

func (v NullableGasless) IsSet() bool {
	return v.isSet
}

func (v *NullableGasless) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGasless(val *Gasless) *NullableGasless {
	return &NullableGasless{value: val, isSet: true}
}

func (v NullableGasless) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGasless) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

