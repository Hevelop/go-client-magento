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

// checks if the InventorySourceSelectionApiDataSourceSelectionItemInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventorySourceSelectionApiDataSourceSelectionItemInterface{}

// InventorySourceSelectionApiDataSourceSelectionItemInterface Represents source selection result for the specific source and SKU
type InventorySourceSelectionApiDataSourceSelectionItemInterface struct {
	// Source code
	SourceCode string `json:"source_code"`
	// Item SKU
	Sku string `json:"sku"`
	// Quantity which will be deducted for this source
	QtyToDeduct float32 `json:"qty_to_deduct"`
	// Available quantity for this source
	QtyAvailable float32 `json:"qty_available"`
	// ExtensionInterface class for @see \\Magento\\InventorySourceSelectionApi\\Api\\Data\\SourceSelectionItemInterface
	ExtensionAttributes  map[string]interface{} `json:"extension_attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InventorySourceSelectionApiDataSourceSelectionItemInterface InventorySourceSelectionApiDataSourceSelectionItemInterface

// NewInventorySourceSelectionApiDataSourceSelectionItemInterface instantiates a new InventorySourceSelectionApiDataSourceSelectionItemInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventorySourceSelectionApiDataSourceSelectionItemInterface(sourceCode string, sku string, qtyToDeduct float32, qtyAvailable float32) *InventorySourceSelectionApiDataSourceSelectionItemInterface {
	this := InventorySourceSelectionApiDataSourceSelectionItemInterface{}
	this.SourceCode = sourceCode
	this.Sku = sku
	this.QtyToDeduct = qtyToDeduct
	this.QtyAvailable = qtyAvailable
	return &this
}

// NewInventorySourceSelectionApiDataSourceSelectionItemInterfaceWithDefaults instantiates a new InventorySourceSelectionApiDataSourceSelectionItemInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventorySourceSelectionApiDataSourceSelectionItemInterfaceWithDefaults() *InventorySourceSelectionApiDataSourceSelectionItemInterface {
	this := InventorySourceSelectionApiDataSourceSelectionItemInterface{}
	return &this
}

// GetSourceCode returns the SourceCode field value
func (o *InventorySourceSelectionApiDataSourceSelectionItemInterface) GetSourceCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceCode
}

// GetSourceCodeOk returns a tuple with the SourceCode field value
// and a boolean to check if the value has been set.
func (o *InventorySourceSelectionApiDataSourceSelectionItemInterface) GetSourceCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceCode, true
}

// SetSourceCode sets field value
func (o *InventorySourceSelectionApiDataSourceSelectionItemInterface) SetSourceCode(v string) {
	o.SourceCode = v
}

// GetSku returns the Sku field value
func (o *InventorySourceSelectionApiDataSourceSelectionItemInterface) GetSku() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sku
}

// GetSkuOk returns a tuple with the Sku field value
// and a boolean to check if the value has been set.
func (o *InventorySourceSelectionApiDataSourceSelectionItemInterface) GetSkuOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sku, true
}

// SetSku sets field value
func (o *InventorySourceSelectionApiDataSourceSelectionItemInterface) SetSku(v string) {
	o.Sku = v
}

// GetQtyToDeduct returns the QtyToDeduct field value
func (o *InventorySourceSelectionApiDataSourceSelectionItemInterface) GetQtyToDeduct() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.QtyToDeduct
}

// GetQtyToDeductOk returns a tuple with the QtyToDeduct field value
// and a boolean to check if the value has been set.
func (o *InventorySourceSelectionApiDataSourceSelectionItemInterface) GetQtyToDeductOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QtyToDeduct, true
}

// SetQtyToDeduct sets field value
func (o *InventorySourceSelectionApiDataSourceSelectionItemInterface) SetQtyToDeduct(v float32) {
	o.QtyToDeduct = v
}

// GetQtyAvailable returns the QtyAvailable field value
func (o *InventorySourceSelectionApiDataSourceSelectionItemInterface) GetQtyAvailable() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.QtyAvailable
}

// GetQtyAvailableOk returns a tuple with the QtyAvailable field value
// and a boolean to check if the value has been set.
func (o *InventorySourceSelectionApiDataSourceSelectionItemInterface) GetQtyAvailableOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QtyAvailable, true
}

// SetQtyAvailable sets field value
func (o *InventorySourceSelectionApiDataSourceSelectionItemInterface) SetQtyAvailable(v float32) {
	o.QtyAvailable = v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *InventorySourceSelectionApiDataSourceSelectionItemInterface) GetExtensionAttributes() map[string]interface{} {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventorySourceSelectionApiDataSourceSelectionItemInterface) GetExtensionAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return map[string]interface{}{}, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *InventorySourceSelectionApiDataSourceSelectionItemInterface) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given map[string]interface{} and assigns it to the ExtensionAttributes field.
func (o *InventorySourceSelectionApiDataSourceSelectionItemInterface) SetExtensionAttributes(v map[string]interface{}) {
	o.ExtensionAttributes = v
}

func (o InventorySourceSelectionApiDataSourceSelectionItemInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventorySourceSelectionApiDataSourceSelectionItemInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source_code"] = o.SourceCode
	toSerialize["sku"] = o.Sku
	toSerialize["qty_to_deduct"] = o.QtyToDeduct
	toSerialize["qty_available"] = o.QtyAvailable
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extension_attributes"] = o.ExtensionAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InventorySourceSelectionApiDataSourceSelectionItemInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"source_code",
		"sku",
		"qty_to_deduct",
		"qty_available",
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

	varInventorySourceSelectionApiDataSourceSelectionItemInterface := _InventorySourceSelectionApiDataSourceSelectionItemInterface{}

	err = json.Unmarshal(data, &varInventorySourceSelectionApiDataSourceSelectionItemInterface)

	if err != nil {
		return err
	}

	*o = InventorySourceSelectionApiDataSourceSelectionItemInterface(varInventorySourceSelectionApiDataSourceSelectionItemInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "source_code")
		delete(additionalProperties, "sku")
		delete(additionalProperties, "qty_to_deduct")
		delete(additionalProperties, "qty_available")
		delete(additionalProperties, "extension_attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *InventorySourceSelectionApiDataSourceSelectionItemInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *InventorySourceSelectionApiDataSourceSelectionItemInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableInventorySourceSelectionApiDataSourceSelectionItemInterface struct {
	value *InventorySourceSelectionApiDataSourceSelectionItemInterface
	isSet bool
}

func (v NullableInventorySourceSelectionApiDataSourceSelectionItemInterface) Get() *InventorySourceSelectionApiDataSourceSelectionItemInterface {
	return v.value
}

func (v *NullableInventorySourceSelectionApiDataSourceSelectionItemInterface) Set(val *InventorySourceSelectionApiDataSourceSelectionItemInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableInventorySourceSelectionApiDataSourceSelectionItemInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableInventorySourceSelectionApiDataSourceSelectionItemInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventorySourceSelectionApiDataSourceSelectionItemInterface(val *InventorySourceSelectionApiDataSourceSelectionItemInterface) *NullableInventorySourceSelectionApiDataSourceSelectionItemInterface {
	return &NullableInventorySourceSelectionApiDataSourceSelectionItemInterface{value: val, isSet: true}
}

func (v NullableInventorySourceSelectionApiDataSourceSelectionItemInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventorySourceSelectionApiDataSourceSelectionItemInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
