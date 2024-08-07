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

// checks if the BuildSwapRequestDtoV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuildSwapRequestDtoV2{}

// BuildSwapRequestDtoV2 struct for BuildSwapRequestDtoV2
type BuildSwapRequestDtoV2 struct {
	// The unique id of the quote
	QuoteId string `json:"quoteId"`
	// The address which will fill the quote. Not required if then taker address was provided during the quote request
	TakerAddress *string `json:"takerAddress,omitempty"`
	// The maximum acceptable slippage of the buyAmount amount. Default value is 5%. This value is ignored if slippage is not applicable to the selected quote
	Slippage float32 `json:"slippage"`
	// If true, the response will contains the approve call (if necessary)
	IncludeApprove bool `json:"includeApprove"`
}

type _BuildSwapRequestDtoV2 BuildSwapRequestDtoV2

// NewBuildSwapRequestDtoV2 instantiates a new BuildSwapRequestDtoV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildSwapRequestDtoV2(quoteId string, slippage float32, includeApprove bool) *BuildSwapRequestDtoV2 {
	this := BuildSwapRequestDtoV2{}
	this.QuoteId = quoteId
	this.Slippage = slippage
	this.IncludeApprove = includeApprove
	return &this
}

// NewBuildSwapRequestDtoV2WithDefaults instantiates a new BuildSwapRequestDtoV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildSwapRequestDtoV2WithDefaults() *BuildSwapRequestDtoV2 {
	this := BuildSwapRequestDtoV2{}
	return &this
}

// GetQuoteId returns the QuoteId field value
func (o *BuildSwapRequestDtoV2) GetQuoteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QuoteId
}

// GetQuoteIdOk returns a tuple with the QuoteId field value
// and a boolean to check if the value has been set.
func (o *BuildSwapRequestDtoV2) GetQuoteIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QuoteId, true
}

// SetQuoteId sets field value
func (o *BuildSwapRequestDtoV2) SetQuoteId(v string) {
	o.QuoteId = v
}

// GetTakerAddress returns the TakerAddress field value if set, zero value otherwise.
func (o *BuildSwapRequestDtoV2) GetTakerAddress() string {
	if o == nil || IsNil(o.TakerAddress) {
		var ret string
		return ret
	}
	return *o.TakerAddress
}

// GetTakerAddressOk returns a tuple with the TakerAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildSwapRequestDtoV2) GetTakerAddressOk() (*string, bool) {
	if o == nil || IsNil(o.TakerAddress) {
		return nil, false
	}
	return o.TakerAddress, true
}

// HasTakerAddress returns a boolean if a field has been set.
func (o *BuildSwapRequestDtoV2) HasTakerAddress() bool {
	if o != nil && !IsNil(o.TakerAddress) {
		return true
	}

	return false
}

// SetTakerAddress gets a reference to the given string and assigns it to the TakerAddress field.
func (o *BuildSwapRequestDtoV2) SetTakerAddress(v string) {
	o.TakerAddress = &v
}

// GetSlippage returns the Slippage field value
func (o *BuildSwapRequestDtoV2) GetSlippage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Slippage
}

// GetSlippageOk returns a tuple with the Slippage field value
// and a boolean to check if the value has been set.
func (o *BuildSwapRequestDtoV2) GetSlippageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slippage, true
}

// SetSlippage sets field value
func (o *BuildSwapRequestDtoV2) SetSlippage(v float32) {
	o.Slippage = v
}

// GetIncludeApprove returns the IncludeApprove field value
func (o *BuildSwapRequestDtoV2) GetIncludeApprove() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IncludeApprove
}

// GetIncludeApproveOk returns a tuple with the IncludeApprove field value
// and a boolean to check if the value has been set.
func (o *BuildSwapRequestDtoV2) GetIncludeApproveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncludeApprove, true
}

// SetIncludeApprove sets field value
func (o *BuildSwapRequestDtoV2) SetIncludeApprove(v bool) {
	o.IncludeApprove = v
}

func (o BuildSwapRequestDtoV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuildSwapRequestDtoV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["quoteId"] = o.QuoteId
	if !IsNil(o.TakerAddress) {
		toSerialize["takerAddress"] = o.TakerAddress
	}
	toSerialize["slippage"] = o.Slippage
	toSerialize["includeApprove"] = o.IncludeApprove
	return toSerialize, nil
}

func (o *BuildSwapRequestDtoV2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"quoteId",
		"slippage",
		"includeApprove",
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

	varBuildSwapRequestDtoV2 := _BuildSwapRequestDtoV2{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBuildSwapRequestDtoV2)

	if err != nil {
		return err
	}

	*o = BuildSwapRequestDtoV2(varBuildSwapRequestDtoV2)

	return err
}

type NullableBuildSwapRequestDtoV2 struct {
	value *BuildSwapRequestDtoV2
	isSet bool
}

func (v NullableBuildSwapRequestDtoV2) Get() *BuildSwapRequestDtoV2 {
	return v.value
}

func (v *NullableBuildSwapRequestDtoV2) Set(val *BuildSwapRequestDtoV2) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildSwapRequestDtoV2) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildSwapRequestDtoV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildSwapRequestDtoV2(val *BuildSwapRequestDtoV2) *NullableBuildSwapRequestDtoV2 {
	return &NullableBuildSwapRequestDtoV2{value: val, isSet: true}
}

func (v NullableBuildSwapRequestDtoV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildSwapRequestDtoV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


