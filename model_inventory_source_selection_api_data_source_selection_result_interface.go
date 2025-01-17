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

// checks if the InventorySourceSelectionApiDataSourceSelectionResultInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventorySourceSelectionApiDataSourceSelectionResultInterface{}

// InventorySourceSelectionApiDataSourceSelectionResultInterface Result of how we will deduct product qty from different Sources
type InventorySourceSelectionApiDataSourceSelectionResultInterface struct {
	SourceSelectionItems []InventorySourceSelectionApiDataSourceSelectionItemInterface `json:"source_selection_items"`
	Shippable            bool                                                          `json:"shippable"`
	// ExtensionInterface class for @see \\Magento\\InventorySourceSelectionApi\\Api\\Data\\SourceSelectionResultInterface
	ExtensionAttributes  map[string]interface{} `json:"extension_attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InventorySourceSelectionApiDataSourceSelectionResultInterface InventorySourceSelectionApiDataSourceSelectionResultInterface

// NewInventorySourceSelectionApiDataSourceSelectionResultInterface instantiates a new InventorySourceSelectionApiDataSourceSelectionResultInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventorySourceSelectionApiDataSourceSelectionResultInterface(sourceSelectionItems []InventorySourceSelectionApiDataSourceSelectionItemInterface, shippable bool) *InventorySourceSelectionApiDataSourceSelectionResultInterface {
	this := InventorySourceSelectionApiDataSourceSelectionResultInterface{}
	this.SourceSelectionItems = sourceSelectionItems
	this.Shippable = shippable
	return &this
}

// NewInventorySourceSelectionApiDataSourceSelectionResultInterfaceWithDefaults instantiates a new InventorySourceSelectionApiDataSourceSelectionResultInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventorySourceSelectionApiDataSourceSelectionResultInterfaceWithDefaults() *InventorySourceSelectionApiDataSourceSelectionResultInterface {
	this := InventorySourceSelectionApiDataSourceSelectionResultInterface{}
	return &this
}

// GetSourceSelectionItems returns the SourceSelectionItems field value
func (o *InventorySourceSelectionApiDataSourceSelectionResultInterface) GetSourceSelectionItems() []InventorySourceSelectionApiDataSourceSelectionItemInterface {
	if o == nil {
		var ret []InventorySourceSelectionApiDataSourceSelectionItemInterface
		return ret
	}

	return o.SourceSelectionItems
}

// GetSourceSelectionItemsOk returns a tuple with the SourceSelectionItems field value
// and a boolean to check if the value has been set.
func (o *InventorySourceSelectionApiDataSourceSelectionResultInterface) GetSourceSelectionItemsOk() ([]InventorySourceSelectionApiDataSourceSelectionItemInterface, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceSelectionItems, true
}

// SetSourceSelectionItems sets field value
func (o *InventorySourceSelectionApiDataSourceSelectionResultInterface) SetSourceSelectionItems(v []InventorySourceSelectionApiDataSourceSelectionItemInterface) {
	o.SourceSelectionItems = v
}

// GetShippable returns the Shippable field value
func (o *InventorySourceSelectionApiDataSourceSelectionResultInterface) GetShippable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Shippable
}

// GetShippableOk returns a tuple with the Shippable field value
// and a boolean to check if the value has been set.
func (o *InventorySourceSelectionApiDataSourceSelectionResultInterface) GetShippableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Shippable, true
}

// SetShippable sets field value
func (o *InventorySourceSelectionApiDataSourceSelectionResultInterface) SetShippable(v bool) {
	o.Shippable = v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *InventorySourceSelectionApiDataSourceSelectionResultInterface) GetExtensionAttributes() map[string]interface{} {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventorySourceSelectionApiDataSourceSelectionResultInterface) GetExtensionAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return map[string]interface{}{}, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *InventorySourceSelectionApiDataSourceSelectionResultInterface) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given map[string]interface{} and assigns it to the ExtensionAttributes field.
func (o *InventorySourceSelectionApiDataSourceSelectionResultInterface) SetExtensionAttributes(v map[string]interface{}) {
	o.ExtensionAttributes = v
}

func (o InventorySourceSelectionApiDataSourceSelectionResultInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventorySourceSelectionApiDataSourceSelectionResultInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source_selection_items"] = o.SourceSelectionItems
	toSerialize["shippable"] = o.Shippable
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extension_attributes"] = o.ExtensionAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InventorySourceSelectionApiDataSourceSelectionResultInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"source_selection_items",
		"shippable",
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

	varInventorySourceSelectionApiDataSourceSelectionResultInterface := _InventorySourceSelectionApiDataSourceSelectionResultInterface{}

	err = json.Unmarshal(data, &varInventorySourceSelectionApiDataSourceSelectionResultInterface)

	if err != nil {
		return err
	}

	*o = InventorySourceSelectionApiDataSourceSelectionResultInterface(varInventorySourceSelectionApiDataSourceSelectionResultInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "source_selection_items")
		delete(additionalProperties, "shippable")
		delete(additionalProperties, "extension_attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *InventorySourceSelectionApiDataSourceSelectionResultInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *InventorySourceSelectionApiDataSourceSelectionResultInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableInventorySourceSelectionApiDataSourceSelectionResultInterface struct {
	value *InventorySourceSelectionApiDataSourceSelectionResultInterface
	isSet bool
}

func (v NullableInventorySourceSelectionApiDataSourceSelectionResultInterface) Get() *InventorySourceSelectionApiDataSourceSelectionResultInterface {
	return v.value
}

func (v *NullableInventorySourceSelectionApiDataSourceSelectionResultInterface) Set(val *InventorySourceSelectionApiDataSourceSelectionResultInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableInventorySourceSelectionApiDataSourceSelectionResultInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableInventorySourceSelectionApiDataSourceSelectionResultInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventorySourceSelectionApiDataSourceSelectionResultInterface(val *InventorySourceSelectionApiDataSourceSelectionResultInterface) *NullableInventorySourceSelectionApiDataSourceSelectionResultInterface {
	return &NullableInventorySourceSelectionApiDataSourceSelectionResultInterface{value: val, isSet: true}
}

func (v NullableInventorySourceSelectionApiDataSourceSelectionResultInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventorySourceSelectionApiDataSourceSelectionResultInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
