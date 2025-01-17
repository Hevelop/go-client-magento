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

// checks if the SalesDataShipmentItemCreationInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SalesDataShipmentItemCreationInterface{}

// SalesDataShipmentItemCreationInterface Input argument for shipment item creation Interface ShipmentItemCreationInterface
type SalesDataShipmentItemCreationInterface struct {
	// ExtensionInterface class for @see \\Magento\\Sales\\Api\\Data\\ShipmentItemCreationInterface
	ExtensionAttributes map[string]interface{} `json:"extension_attributes,omitempty"`
	// Order item ID.
	OrderItemId int32 `json:"order_item_id"`
	// Quantity.
	Qty                  float32 `json:"qty"`
	AdditionalProperties map[string]interface{}
}

type _SalesDataShipmentItemCreationInterface SalesDataShipmentItemCreationInterface

// NewSalesDataShipmentItemCreationInterface instantiates a new SalesDataShipmentItemCreationInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSalesDataShipmentItemCreationInterface(orderItemId int32, qty float32) *SalesDataShipmentItemCreationInterface {
	this := SalesDataShipmentItemCreationInterface{}
	this.OrderItemId = orderItemId
	this.Qty = qty
	return &this
}

// NewSalesDataShipmentItemCreationInterfaceWithDefaults instantiates a new SalesDataShipmentItemCreationInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSalesDataShipmentItemCreationInterfaceWithDefaults() *SalesDataShipmentItemCreationInterface {
	this := SalesDataShipmentItemCreationInterface{}
	return &this
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *SalesDataShipmentItemCreationInterface) GetExtensionAttributes() map[string]interface{} {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataShipmentItemCreationInterface) GetExtensionAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return map[string]interface{}{}, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *SalesDataShipmentItemCreationInterface) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given map[string]interface{} and assigns it to the ExtensionAttributes field.
func (o *SalesDataShipmentItemCreationInterface) SetExtensionAttributes(v map[string]interface{}) {
	o.ExtensionAttributes = v
}

// GetOrderItemId returns the OrderItemId field value
func (o *SalesDataShipmentItemCreationInterface) GetOrderItemId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OrderItemId
}

// GetOrderItemIdOk returns a tuple with the OrderItemId field value
// and a boolean to check if the value has been set.
func (o *SalesDataShipmentItemCreationInterface) GetOrderItemIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderItemId, true
}

// SetOrderItemId sets field value
func (o *SalesDataShipmentItemCreationInterface) SetOrderItemId(v int32) {
	o.OrderItemId = v
}

// GetQty returns the Qty field value
func (o *SalesDataShipmentItemCreationInterface) GetQty() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Qty
}

// GetQtyOk returns a tuple with the Qty field value
// and a boolean to check if the value has been set.
func (o *SalesDataShipmentItemCreationInterface) GetQtyOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Qty, true
}

// SetQty sets field value
func (o *SalesDataShipmentItemCreationInterface) SetQty(v float32) {
	o.Qty = v
}

func (o SalesDataShipmentItemCreationInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SalesDataShipmentItemCreationInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extension_attributes"] = o.ExtensionAttributes
	}
	toSerialize["order_item_id"] = o.OrderItemId
	toSerialize["qty"] = o.Qty

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SalesDataShipmentItemCreationInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"order_item_id",
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

	varSalesDataShipmentItemCreationInterface := _SalesDataShipmentItemCreationInterface{}

	err = json.Unmarshal(data, &varSalesDataShipmentItemCreationInterface)

	if err != nil {
		return err
	}

	*o = SalesDataShipmentItemCreationInterface(varSalesDataShipmentItemCreationInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "extension_attributes")
		delete(additionalProperties, "order_item_id")
		delete(additionalProperties, "qty")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *SalesDataShipmentItemCreationInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *SalesDataShipmentItemCreationInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableSalesDataShipmentItemCreationInterface struct {
	value *SalesDataShipmentItemCreationInterface
	isSet bool
}

func (v NullableSalesDataShipmentItemCreationInterface) Get() *SalesDataShipmentItemCreationInterface {
	return v.value
}

func (v *NullableSalesDataShipmentItemCreationInterface) Set(val *SalesDataShipmentItemCreationInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableSalesDataShipmentItemCreationInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableSalesDataShipmentItemCreationInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSalesDataShipmentItemCreationInterface(val *SalesDataShipmentItemCreationInterface) *NullableSalesDataShipmentItemCreationInterface {
	return &NullableSalesDataShipmentItemCreationInterface{value: val, isSet: true}
}

func (v NullableSalesDataShipmentItemCreationInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSalesDataShipmentItemCreationInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
