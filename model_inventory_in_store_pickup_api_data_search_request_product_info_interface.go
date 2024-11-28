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

// checks if the InventoryInStorePickupApiDataSearchRequestProductInfoInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryInStorePickupApiDataSearchRequestProductInfoInterface{}

// InventoryInStorePickupApiDataSearchRequestProductInfoInterface Product Info Data Transfer Object.
type InventoryInStorePickupApiDataSearchRequestProductInfoInterface struct {
	// Product SKU.
	Sku string `json:"sku"`
	// ExtensionInterface class for @see \\Magento\\InventoryInStorePickupApi\\Api\\Data\\SearchRequest\\ProductInfoInterface
	ExtensionAttributes  map[string]interface{} `json:"extension_attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InventoryInStorePickupApiDataSearchRequestProductInfoInterface InventoryInStorePickupApiDataSearchRequestProductInfoInterface

// NewInventoryInStorePickupApiDataSearchRequestProductInfoInterface instantiates a new InventoryInStorePickupApiDataSearchRequestProductInfoInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryInStorePickupApiDataSearchRequestProductInfoInterface(sku string) *InventoryInStorePickupApiDataSearchRequestProductInfoInterface {
	this := InventoryInStorePickupApiDataSearchRequestProductInfoInterface{}
	this.Sku = sku
	return &this
}

// NewInventoryInStorePickupApiDataSearchRequestProductInfoInterfaceWithDefaults instantiates a new InventoryInStorePickupApiDataSearchRequestProductInfoInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryInStorePickupApiDataSearchRequestProductInfoInterfaceWithDefaults() *InventoryInStorePickupApiDataSearchRequestProductInfoInterface {
	this := InventoryInStorePickupApiDataSearchRequestProductInfoInterface{}
	return &this
}

// GetSku returns the Sku field value
func (o *InventoryInStorePickupApiDataSearchRequestProductInfoInterface) GetSku() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sku
}

// GetSkuOk returns a tuple with the Sku field value
// and a boolean to check if the value has been set.
func (o *InventoryInStorePickupApiDataSearchRequestProductInfoInterface) GetSkuOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sku, true
}

// SetSku sets field value
func (o *InventoryInStorePickupApiDataSearchRequestProductInfoInterface) SetSku(v string) {
	o.Sku = v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *InventoryInStorePickupApiDataSearchRequestProductInfoInterface) GetExtensionAttributes() map[string]interface{} {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryInStorePickupApiDataSearchRequestProductInfoInterface) GetExtensionAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return map[string]interface{}{}, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *InventoryInStorePickupApiDataSearchRequestProductInfoInterface) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given map[string]interface{} and assigns it to the ExtensionAttributes field.
func (o *InventoryInStorePickupApiDataSearchRequestProductInfoInterface) SetExtensionAttributes(v map[string]interface{}) {
	o.ExtensionAttributes = v
}

func (o InventoryInStorePickupApiDataSearchRequestProductInfoInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventoryInStorePickupApiDataSearchRequestProductInfoInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sku"] = o.Sku
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extension_attributes"] = o.ExtensionAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InventoryInStorePickupApiDataSearchRequestProductInfoInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sku",
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

	varInventoryInStorePickupApiDataSearchRequestProductInfoInterface := _InventoryInStorePickupApiDataSearchRequestProductInfoInterface{}

	err = json.Unmarshal(data, &varInventoryInStorePickupApiDataSearchRequestProductInfoInterface)

	if err != nil {
		return err
	}

	*o = InventoryInStorePickupApiDataSearchRequestProductInfoInterface(varInventoryInStorePickupApiDataSearchRequestProductInfoInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "sku")
		delete(additionalProperties, "extension_attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *InventoryInStorePickupApiDataSearchRequestProductInfoInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *InventoryInStorePickupApiDataSearchRequestProductInfoInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableInventoryInStorePickupApiDataSearchRequestProductInfoInterface struct {
	value *InventoryInStorePickupApiDataSearchRequestProductInfoInterface
	isSet bool
}

func (v NullableInventoryInStorePickupApiDataSearchRequestProductInfoInterface) Get() *InventoryInStorePickupApiDataSearchRequestProductInfoInterface {
	return v.value
}

func (v *NullableInventoryInStorePickupApiDataSearchRequestProductInfoInterface) Set(val *InventoryInStorePickupApiDataSearchRequestProductInfoInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryInStorePickupApiDataSearchRequestProductInfoInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryInStorePickupApiDataSearchRequestProductInfoInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryInStorePickupApiDataSearchRequestProductInfoInterface(val *InventoryInStorePickupApiDataSearchRequestProductInfoInterface) *NullableInventoryInStorePickupApiDataSearchRequestProductInfoInterface {
	return &NullableInventoryInStorePickupApiDataSearchRequestProductInfoInterface{value: val, isSet: true}
}

func (v NullableInventoryInStorePickupApiDataSearchRequestProductInfoInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryInStorePickupApiDataSearchRequestProductInfoInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}