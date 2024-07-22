/*
AVNU API

REST API documentation for accessing liquidity on Layer 2.  AVNU is a decentralized exchange protocol enabling the fastest and the most efficient operations in DeFi for Layer 2 with better pricing, zero slippage, MEV-protection and gasless trading.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// checks if the TypedDataMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TypedDataMessage{}

// TypedDataMessage struct for TypedDataMessage
type TypedDataMessage struct {
	Content *map[string]map[string]interface{} `json:"content,omitempty"`
	Values []map[string]interface{} `json:"values"`
	IsEmpty bool `json:"isEmpty"`
	Size int32 `json:"size"`
	Entries []TypedDataMessageEntriesInner `json:"entries"`
	Keys []string `json:"keys"`
	AdditionalProperties map[string]interface{}
}

type _TypedDataMessage TypedDataMessage

// NewTypedDataMessage instantiates a new TypedDataMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTypedDataMessage(values []map[string]interface{}, isEmpty bool, size int32, entries []TypedDataMessageEntriesInner, keys []string) *TypedDataMessage {
	this := TypedDataMessage{}
	this.Values = values
	this.IsEmpty = isEmpty
	this.Size = size
	this.Entries = entries
	this.Keys = keys
	return &this
}

// NewTypedDataMessageWithDefaults instantiates a new TypedDataMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTypedDataMessageWithDefaults() *TypedDataMessage {
	this := TypedDataMessage{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *TypedDataMessage) GetContent() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Content) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypedDataMessage) GetContentOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *TypedDataMessage) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given map[string]map[string]interface{} and assigns it to the Content field.
func (o *TypedDataMessage) SetContent(v map[string]map[string]interface{}) {
	o.Content = &v
}

// GetValues returns the Values field value
func (o *TypedDataMessage) GetValues() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *TypedDataMessage) GetValuesOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *TypedDataMessage) SetValues(v []map[string]interface{}) {
	o.Values = v
}

// GetIsEmpty returns the IsEmpty field value
func (o *TypedDataMessage) GetIsEmpty() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsEmpty
}

// GetIsEmptyOk returns a tuple with the IsEmpty field value
// and a boolean to check if the value has been set.
func (o *TypedDataMessage) GetIsEmptyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsEmpty, true
}

// SetIsEmpty sets field value
func (o *TypedDataMessage) SetIsEmpty(v bool) {
	o.IsEmpty = v
}

// GetSize returns the Size field value
func (o *TypedDataMessage) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *TypedDataMessage) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *TypedDataMessage) SetSize(v int32) {
	o.Size = v
}

// GetEntries returns the Entries field value
func (o *TypedDataMessage) GetEntries() []TypedDataMessageEntriesInner {
	if o == nil {
		var ret []TypedDataMessageEntriesInner
		return ret
	}

	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value
// and a boolean to check if the value has been set.
func (o *TypedDataMessage) GetEntriesOk() ([]TypedDataMessageEntriesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Entries, true
}

// SetEntries sets field value
func (o *TypedDataMessage) SetEntries(v []TypedDataMessageEntriesInner) {
	o.Entries = v
}

// GetKeys returns the Keys field value
func (o *TypedDataMessage) GetKeys() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Keys
}

// GetKeysOk returns a tuple with the Keys field value
// and a boolean to check if the value has been set.
func (o *TypedDataMessage) GetKeysOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Keys, true
}

// SetKeys sets field value
func (o *TypedDataMessage) SetKeys(v []string) {
	o.Keys = v
}

func (o TypedDataMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TypedDataMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	toSerialize["values"] = o.Values
	toSerialize["isEmpty"] = o.IsEmpty
	toSerialize["size"] = o.Size
	toSerialize["entries"] = o.Entries
	toSerialize["keys"] = o.Keys

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TypedDataMessage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"values",
		"isEmpty",
		"size",
		"entries",
		"keys",
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

	varTypedDataMessage := _TypedDataMessage{}

	err = json.Unmarshal(data, &varTypedDataMessage)

	if err != nil {
		return err
	}

	*o = TypedDataMessage(varTypedDataMessage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "content")
		delete(additionalProperties, "values")
		delete(additionalProperties, "isEmpty")
		delete(additionalProperties, "size")
		delete(additionalProperties, "entries")
		delete(additionalProperties, "keys")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTypedDataMessage struct {
	value *TypedDataMessage
	isSet bool
}

func (v NullableTypedDataMessage) Get() *TypedDataMessage {
	return v.value
}

func (v *NullableTypedDataMessage) Set(val *TypedDataMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableTypedDataMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableTypedDataMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypedDataMessage(val *TypedDataMessage) *NullableTypedDataMessage {
	return &NullableTypedDataMessage{value: val, isSet: true}
}

func (v NullableTypedDataMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypedDataMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

