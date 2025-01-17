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

// checks if the InventoryCatalogApiDataPartialInventoryTransferItemInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryCatalogApiDataPartialInventoryTransferItemInterface{}

// InventoryCatalogApiDataPartialInventoryTransferItemInterface Specifies item and quantity for partial inventory transfer.
type InventoryCatalogApiDataPartialInventoryTransferItemInterface struct {
	Sku                  string  `json:"sku"`
	Qty                  float32 `json:"qty"`
	AdditionalProperties map[string]interface{}
}

type _InventoryCatalogApiDataPartialInventoryTransferItemInterface InventoryCatalogApiDataPartialInventoryTransferItemInterface

// NewInventoryCatalogApiDataPartialInventoryTransferItemInterface instantiates a new InventoryCatalogApiDataPartialInventoryTransferItemInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryCatalogApiDataPartialInventoryTransferItemInterface(sku string, qty float32) *InventoryCatalogApiDataPartialInventoryTransferItemInterface {
	this := InventoryCatalogApiDataPartialInventoryTransferItemInterface{}
	this.Sku = sku
	this.Qty = qty
	return &this
}

// NewInventoryCatalogApiDataPartialInventoryTransferItemInterfaceWithDefaults instantiates a new InventoryCatalogApiDataPartialInventoryTransferItemInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryCatalogApiDataPartialInventoryTransferItemInterfaceWithDefaults() *InventoryCatalogApiDataPartialInventoryTransferItemInterface {
	this := InventoryCatalogApiDataPartialInventoryTransferItemInterface{}
	return &this
}

// GetSku returns the Sku field value
func (o *InventoryCatalogApiDataPartialInventoryTransferItemInterface) GetSku() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sku
}

// GetSkuOk returns a tuple with the Sku field value
// and a boolean to check if the value has been set.
func (o *InventoryCatalogApiDataPartialInventoryTransferItemInterface) GetSkuOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sku, true
}

// SetSku sets field value
func (o *InventoryCatalogApiDataPartialInventoryTransferItemInterface) SetSku(v string) {
	o.Sku = v
}

// GetQty returns the Qty field value
func (o *InventoryCatalogApiDataPartialInventoryTransferItemInterface) GetQty() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Qty
}

// GetQtyOk returns a tuple with the Qty field value
// and a boolean to check if the value has been set.
func (o *InventoryCatalogApiDataPartialInventoryTransferItemInterface) GetQtyOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Qty, true
}

// SetQty sets field value
func (o *InventoryCatalogApiDataPartialInventoryTransferItemInterface) SetQty(v float32) {
	o.Qty = v
}

func (o InventoryCatalogApiDataPartialInventoryTransferItemInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventoryCatalogApiDataPartialInventoryTransferItemInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sku"] = o.Sku
	toSerialize["qty"] = o.Qty

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InventoryCatalogApiDataPartialInventoryTransferItemInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sku",
		"qty",
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

	varInventoryCatalogApiDataPartialInventoryTransferItemInterface := _InventoryCatalogApiDataPartialInventoryTransferItemInterface{}

	err = json.Unmarshal(data, &varInventoryCatalogApiDataPartialInventoryTransferItemInterface)

	if err != nil {
		return err
	}

	*o = InventoryCatalogApiDataPartialInventoryTransferItemInterface(varInventoryCatalogApiDataPartialInventoryTransferItemInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "sku")
		delete(additionalProperties, "qty")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *InventoryCatalogApiDataPartialInventoryTransferItemInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *InventoryCatalogApiDataPartialInventoryTransferItemInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableInventoryCatalogApiDataPartialInventoryTransferItemInterface struct {
	value *InventoryCatalogApiDataPartialInventoryTransferItemInterface
	isSet bool
}

func (v NullableInventoryCatalogApiDataPartialInventoryTransferItemInterface) Get() *InventoryCatalogApiDataPartialInventoryTransferItemInterface {
	return v.value
}

func (v *NullableInventoryCatalogApiDataPartialInventoryTransferItemInterface) Set(val *InventoryCatalogApiDataPartialInventoryTransferItemInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryCatalogApiDataPartialInventoryTransferItemInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryCatalogApiDataPartialInventoryTransferItemInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryCatalogApiDataPartialInventoryTransferItemInterface(val *InventoryCatalogApiDataPartialInventoryTransferItemInterface) *NullableInventoryCatalogApiDataPartialInventoryTransferItemInterface {
	return &NullableInventoryCatalogApiDataPartialInventoryTransferItemInterface{value: val, isSet: true}
}

func (v NullableInventoryCatalogApiDataPartialInventoryTransferItemInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryCatalogApiDataPartialInventoryTransferItemInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
