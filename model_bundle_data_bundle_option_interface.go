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

// checks if the BundleDataBundleOptionInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BundleDataBundleOptionInterface{}

// BundleDataBundleOptionInterface Interface BundleOptionInterface
type BundleDataBundleOptionInterface struct {
	// Bundle option id.
	OptionId int32 `json:"option_id"`
	// Bundle option quantity.
	OptionQty int32 `json:"option_qty"`
	// Bundle option selection ids.
	OptionSelections []int32 `json:"option_selections"`
	// ExtensionInterface class for @see \\Magento\\Bundle\\Api\\Data\\BundleOptionInterface
	ExtensionAttributes  map[string]interface{} `json:"extension_attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BundleDataBundleOptionInterface BundleDataBundleOptionInterface

// NewBundleDataBundleOptionInterface instantiates a new BundleDataBundleOptionInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleDataBundleOptionInterface(optionId int32, optionQty int32, optionSelections []int32) *BundleDataBundleOptionInterface {
	this := BundleDataBundleOptionInterface{}
	this.OptionId = optionId
	this.OptionQty = optionQty
	this.OptionSelections = optionSelections
	return &this
}

// NewBundleDataBundleOptionInterfaceWithDefaults instantiates a new BundleDataBundleOptionInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleDataBundleOptionInterfaceWithDefaults() *BundleDataBundleOptionInterface {
	this := BundleDataBundleOptionInterface{}
	return &this
}

// GetOptionId returns the OptionId field value
func (o *BundleDataBundleOptionInterface) GetOptionId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OptionId
}

// GetOptionIdOk returns a tuple with the OptionId field value
// and a boolean to check if the value has been set.
func (o *BundleDataBundleOptionInterface) GetOptionIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OptionId, true
}

// SetOptionId sets field value
func (o *BundleDataBundleOptionInterface) SetOptionId(v int32) {
	o.OptionId = v
}

// GetOptionQty returns the OptionQty field value
func (o *BundleDataBundleOptionInterface) GetOptionQty() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OptionQty
}

// GetOptionQtyOk returns a tuple with the OptionQty field value
// and a boolean to check if the value has been set.
func (o *BundleDataBundleOptionInterface) GetOptionQtyOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OptionQty, true
}

// SetOptionQty sets field value
func (o *BundleDataBundleOptionInterface) SetOptionQty(v int32) {
	o.OptionQty = v
}

// GetOptionSelections returns the OptionSelections field value
func (o *BundleDataBundleOptionInterface) GetOptionSelections() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.OptionSelections
}

// GetOptionSelectionsOk returns a tuple with the OptionSelections field value
// and a boolean to check if the value has been set.
func (o *BundleDataBundleOptionInterface) GetOptionSelectionsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OptionSelections, true
}

// SetOptionSelections sets field value
func (o *BundleDataBundleOptionInterface) SetOptionSelections(v []int32) {
	o.OptionSelections = v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *BundleDataBundleOptionInterface) GetExtensionAttributes() map[string]interface{} {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleDataBundleOptionInterface) GetExtensionAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return map[string]interface{}{}, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *BundleDataBundleOptionInterface) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given map[string]interface{} and assigns it to the ExtensionAttributes field.
func (o *BundleDataBundleOptionInterface) SetExtensionAttributes(v map[string]interface{}) {
	o.ExtensionAttributes = v
}

func (o BundleDataBundleOptionInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BundleDataBundleOptionInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["option_id"] = o.OptionId
	toSerialize["option_qty"] = o.OptionQty
	toSerialize["option_selections"] = o.OptionSelections
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extension_attributes"] = o.ExtensionAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BundleDataBundleOptionInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"option_id",
		"option_qty",
		"option_selections",
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

	varBundleDataBundleOptionInterface := _BundleDataBundleOptionInterface{}

	err = json.Unmarshal(data, &varBundleDataBundleOptionInterface)

	if err != nil {
		return err
	}

	*o = BundleDataBundleOptionInterface(varBundleDataBundleOptionInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "option_id")
		delete(additionalProperties, "option_qty")
		delete(additionalProperties, "option_selections")
		delete(additionalProperties, "extension_attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *BundleDataBundleOptionInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *BundleDataBundleOptionInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableBundleDataBundleOptionInterface struct {
	value *BundleDataBundleOptionInterface
	isSet bool
}

func (v NullableBundleDataBundleOptionInterface) Get() *BundleDataBundleOptionInterface {
	return v.value
}

func (v *NullableBundleDataBundleOptionInterface) Set(val *BundleDataBundleOptionInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleDataBundleOptionInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleDataBundleOptionInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleDataBundleOptionInterface(val *BundleDataBundleOptionInterface) *NullableBundleDataBundleOptionInterface {
	return &NullableBundleDataBundleOptionInterface{value: val, isSet: true}
}

func (v NullableBundleDataBundleOptionInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleDataBundleOptionInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}