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

// checks if the SalesDataShippingAssignmentInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SalesDataShippingAssignmentInterface{}

// SalesDataShippingAssignmentInterface Interface ShippingAssignmentInterface
type SalesDataShippingAssignmentInterface struct {
	Shipping SalesDataShippingInterface `json:"shipping"`
	// Order items of shipping assignment
	Items []SalesDataOrderItemInterface `json:"items"`
	// Stock id
	StockId *int32 `json:"stock_id,omitempty"`
	// ExtensionInterface class for @see \\Magento\\Sales\\Api\\Data\\ShippingAssignmentInterface
	ExtensionAttributes  map[string]interface{} `json:"extension_attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SalesDataShippingAssignmentInterface SalesDataShippingAssignmentInterface

// NewSalesDataShippingAssignmentInterface instantiates a new SalesDataShippingAssignmentInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSalesDataShippingAssignmentInterface(shipping SalesDataShippingInterface, items []SalesDataOrderItemInterface) *SalesDataShippingAssignmentInterface {
	this := SalesDataShippingAssignmentInterface{}
	this.Shipping = shipping
	this.Items = items
	return &this
}

// NewSalesDataShippingAssignmentInterfaceWithDefaults instantiates a new SalesDataShippingAssignmentInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSalesDataShippingAssignmentInterfaceWithDefaults() *SalesDataShippingAssignmentInterface {
	this := SalesDataShippingAssignmentInterface{}
	return &this
}

// GetShipping returns the Shipping field value
func (o *SalesDataShippingAssignmentInterface) GetShipping() SalesDataShippingInterface {
	if o == nil {
		var ret SalesDataShippingInterface
		return ret
	}

	return o.Shipping
}

// GetShippingOk returns a tuple with the Shipping field value
// and a boolean to check if the value has been set.
func (o *SalesDataShippingAssignmentInterface) GetShippingOk() (*SalesDataShippingInterface, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Shipping, true
}

// SetShipping sets field value
func (o *SalesDataShippingAssignmentInterface) SetShipping(v SalesDataShippingInterface) {
	o.Shipping = v
}

// GetItems returns the Items field value
func (o *SalesDataShippingAssignmentInterface) GetItems() []SalesDataOrderItemInterface {
	if o == nil {
		var ret []SalesDataOrderItemInterface
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *SalesDataShippingAssignmentInterface) GetItemsOk() ([]SalesDataOrderItemInterface, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *SalesDataShippingAssignmentInterface) SetItems(v []SalesDataOrderItemInterface) {
	o.Items = v
}

// GetStockId returns the StockId field value if set, zero value otherwise.
func (o *SalesDataShippingAssignmentInterface) GetStockId() int32 {
	if o == nil || IsNil(o.StockId) {
		var ret int32
		return ret
	}
	return *o.StockId
}

// GetStockIdOk returns a tuple with the StockId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataShippingAssignmentInterface) GetStockIdOk() (*int32, bool) {
	if o == nil || IsNil(o.StockId) {
		return nil, false
	}
	return o.StockId, true
}

// HasStockId returns a boolean if a field has been set.
func (o *SalesDataShippingAssignmentInterface) HasStockId() bool {
	if o != nil && !IsNil(o.StockId) {
		return true
	}

	return false
}

// SetStockId gets a reference to the given int32 and assigns it to the StockId field.
func (o *SalesDataShippingAssignmentInterface) SetStockId(v int32) {
	o.StockId = &v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *SalesDataShippingAssignmentInterface) GetExtensionAttributes() map[string]interface{} {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataShippingAssignmentInterface) GetExtensionAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return map[string]interface{}{}, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *SalesDataShippingAssignmentInterface) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given map[string]interface{} and assigns it to the ExtensionAttributes field.
func (o *SalesDataShippingAssignmentInterface) SetExtensionAttributes(v map[string]interface{}) {
	o.ExtensionAttributes = v
}

func (o SalesDataShippingAssignmentInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SalesDataShippingAssignmentInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["shipping"] = o.Shipping
	toSerialize["items"] = o.Items
	if !IsNil(o.StockId) {
		toSerialize["stock_id"] = o.StockId
	}
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extension_attributes"] = o.ExtensionAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SalesDataShippingAssignmentInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"shipping",
		"items",
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

	varSalesDataShippingAssignmentInterface := _SalesDataShippingAssignmentInterface{}

	err = json.Unmarshal(data, &varSalesDataShippingAssignmentInterface)

	if err != nil {
		return err
	}

	*o = SalesDataShippingAssignmentInterface(varSalesDataShippingAssignmentInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "shipping")
		delete(additionalProperties, "items")
		delete(additionalProperties, "stock_id")
		delete(additionalProperties, "extension_attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *SalesDataShippingAssignmentInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *SalesDataShippingAssignmentInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableSalesDataShippingAssignmentInterface struct {
	value *SalesDataShippingAssignmentInterface
	isSet bool
}

func (v NullableSalesDataShippingAssignmentInterface) Get() *SalesDataShippingAssignmentInterface {
	return v.value
}

func (v *NullableSalesDataShippingAssignmentInterface) Set(val *SalesDataShippingAssignmentInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableSalesDataShippingAssignmentInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableSalesDataShippingAssignmentInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSalesDataShippingAssignmentInterface(val *SalesDataShippingAssignmentInterface) *NullableSalesDataShippingAssignmentInterface {
	return &NullableSalesDataShippingAssignmentInterface{value: val, isSet: true}
}

func (v NullableSalesDataShippingAssignmentInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSalesDataShippingAssignmentInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}