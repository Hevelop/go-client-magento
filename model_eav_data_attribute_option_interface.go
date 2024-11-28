/*
Commerce Admin REST endpoints - All inclusive

The schemas documented here are autogenerated from an instance of Adobe Commerce with B2B. Each schema represents a specific user role (Admin, Customer, and Guest) and determines which endpoints are accessible. Use the version switcher to select an Adobe Commerce version and corresponding API.  You can also <a href=\"https://developer.adobe.com/commerce/webapi/rest/quick-reference/generate-local\" target=\"_blank\">generate a local API reference</a> based on your own Adobe Commerce configuration, which allows you to see API documentation for your specific Adobe Commerce modules, third-party modules, and extension attributes.

API version: 2.4.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package magento

import (
	"encoding/json"
	"fmt"
)

// checks if the EavDataAttributeOptionInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EavDataAttributeOptionInterface{}

// EavDataAttributeOptionInterface Created from:
type EavDataAttributeOptionInterface struct {
	// Option label
	Label string `json:"label"`
	// Option value
	Value string `json:"value"`
	// Option order
	SortOrder *int32 `json:"sort_order,omitempty"`
	// Default
	IsDefault *bool `json:"is_default,omitempty"`
	// Option label for store scopes
	StoreLabels          []EavDataAttributeOptionLabelInterface `json:"store_labels,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EavDataAttributeOptionInterface EavDataAttributeOptionInterface

// NewEavDataAttributeOptionInterface instantiates a new EavDataAttributeOptionInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEavDataAttributeOptionInterface(label string, value string) *EavDataAttributeOptionInterface {
	this := EavDataAttributeOptionInterface{}
	this.Label = label
	this.Value = value
	return &this
}

// NewEavDataAttributeOptionInterfaceWithDefaults instantiates a new EavDataAttributeOptionInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEavDataAttributeOptionInterfaceWithDefaults() *EavDataAttributeOptionInterface {
	this := EavDataAttributeOptionInterface{}
	return &this
}

// GetLabel returns the Label field value
func (o *EavDataAttributeOptionInterface) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *EavDataAttributeOptionInterface) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *EavDataAttributeOptionInterface) SetLabel(v string) {
	o.Label = v
}

// GetValue returns the Value field value
func (o *EavDataAttributeOptionInterface) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *EavDataAttributeOptionInterface) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *EavDataAttributeOptionInterface) SetValue(v string) {
	o.Value = v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *EavDataAttributeOptionInterface) GetSortOrder() int32 {
	if o == nil || IsNil(o.SortOrder) {
		var ret int32
		return ret
	}
	return *o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EavDataAttributeOptionInterface) GetSortOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.SortOrder) {
		return nil, false
	}
	return o.SortOrder, true
}

// HasSortOrder returns a boolean if a field has been set.
func (o *EavDataAttributeOptionInterface) HasSortOrder() bool {
	if o != nil && !IsNil(o.SortOrder) {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given int32 and assigns it to the SortOrder field.
func (o *EavDataAttributeOptionInterface) SetSortOrder(v int32) {
	o.SortOrder = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *EavDataAttributeOptionInterface) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EavDataAttributeOptionInterface) GetIsDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *EavDataAttributeOptionInterface) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *EavDataAttributeOptionInterface) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetStoreLabels returns the StoreLabels field value if set, zero value otherwise.
func (o *EavDataAttributeOptionInterface) GetStoreLabels() []EavDataAttributeOptionLabelInterface {
	if o == nil || IsNil(o.StoreLabels) {
		var ret []EavDataAttributeOptionLabelInterface
		return ret
	}
	return o.StoreLabels
}

// GetStoreLabelsOk returns a tuple with the StoreLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EavDataAttributeOptionInterface) GetStoreLabelsOk() ([]EavDataAttributeOptionLabelInterface, bool) {
	if o == nil || IsNil(o.StoreLabels) {
		return nil, false
	}
	return o.StoreLabels, true
}

// HasStoreLabels returns a boolean if a field has been set.
func (o *EavDataAttributeOptionInterface) HasStoreLabels() bool {
	if o != nil && !IsNil(o.StoreLabels) {
		return true
	}

	return false
}

// SetStoreLabels gets a reference to the given []EavDataAttributeOptionLabelInterface and assigns it to the StoreLabels field.
func (o *EavDataAttributeOptionInterface) SetStoreLabels(v []EavDataAttributeOptionLabelInterface) {
	o.StoreLabels = v
}

func (o EavDataAttributeOptionInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EavDataAttributeOptionInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["label"] = o.Label
	toSerialize["value"] = o.Value
	if !IsNil(o.SortOrder) {
		toSerialize["sort_order"] = o.SortOrder
	}
	if !IsNil(o.IsDefault) {
		toSerialize["is_default"] = o.IsDefault
	}
	if !IsNil(o.StoreLabels) {
		toSerialize["store_labels"] = o.StoreLabels
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EavDataAttributeOptionInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"label",
		"value",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varEavDataAttributeOptionInterface := _EavDataAttributeOptionInterface{}

	err = json.Unmarshal(data, &varEavDataAttributeOptionInterface)

	if err != nil {
		return err
	}

	*o = EavDataAttributeOptionInterface(varEavDataAttributeOptionInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "label")
		delete(additionalProperties, "value")
		delete(additionalProperties, "sort_order")
		delete(additionalProperties, "is_default")
		delete(additionalProperties, "store_labels")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *EavDataAttributeOptionInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *EavDataAttributeOptionInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableEavDataAttributeOptionInterface struct {
	value *EavDataAttributeOptionInterface
	isSet bool
}

func (v NullableEavDataAttributeOptionInterface) Get() *EavDataAttributeOptionInterface {
	return v.value
}

func (v *NullableEavDataAttributeOptionInterface) Set(val *EavDataAttributeOptionInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableEavDataAttributeOptionInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableEavDataAttributeOptionInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEavDataAttributeOptionInterface(val *EavDataAttributeOptionInterface) *NullableEavDataAttributeOptionInterface {
	return &NullableEavDataAttributeOptionInterface{value: val, isSet: true}
}

func (v NullableEavDataAttributeOptionInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEavDataAttributeOptionInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
