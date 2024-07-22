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

// checks if the Type type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Type{}

// Type struct for Type
type Type struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type _Type Type

// NewType instantiates a new Type object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewType(name string, type_ string) *Type {
	this := Type{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewTypeWithDefaults instantiates a new Type object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTypeWithDefaults() *Type {
	this := Type{}
	return &this
}

// GetName returns the Name field value
func (o *Type) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Type) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Type) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *Type) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Type) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Type) SetType(v string) {
	o.Type = v
}

func (o Type) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Type) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *Type) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"type",
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

	varType := _Type{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varType)

	if err != nil {
		return err
	}

	*o = Type(varType)

	return err
}

type NullableType struct {
	value *Type
	isSet bool
}

func (v NullableType) Get() *Type {
	return v.value
}

func (v *NullableType) Set(val *Type) {
	v.value = val
	v.isSet = true
}

func (v NullableType) IsSet() bool {
	return v.isSet
}

func (v *NullableType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableType(val *Type) *NullableType {
	return &NullableType{value: val, isSet: true}
}

func (v NullableType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


