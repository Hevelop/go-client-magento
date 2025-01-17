/*
Commerce Admin REST endpoints - All inclusive

The schemas documented here are autogenerated from an instance of Adobe Commerce with B2B. Each schema represents a specific user role (Admin, Customer, and Guest) and determines which endpoints are accessible. Use the version switcher to select an Adobe Commerce version and corresponding API.  You can also <a href=\"https://developer.adobe.com/commerce/webapi/rest/quick-reference/generate-local\" target=\"_blank\">generate a local API reference</a> based on your own Adobe Commerce configuration, which allows you to see API documentation for your specific Adobe Commerce modules, third-party modules, and extension attributes.

API version: 2.4.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package magento

import (
	"encoding/json"
)

// checks if the SalesDataShipmentPackageCreationInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SalesDataShipmentPackageCreationInterface{}

// SalesDataShipmentPackageCreationInterface Shipment package interface. A shipment is a delivery package that contains products. A shipment document accompanies the shipment. This document lists the products and their quantities in the delivery package.
type SalesDataShipmentPackageCreationInterface struct {
	// ExtensionInterface class for @see \\Magento\\Sales\\Api\\Data\\ShipmentPackageCreationInterface
	ExtensionAttributes  map[string]interface{} `json:"extension_attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SalesDataShipmentPackageCreationInterface SalesDataShipmentPackageCreationInterface

// NewSalesDataShipmentPackageCreationInterface instantiates a new SalesDataShipmentPackageCreationInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSalesDataShipmentPackageCreationInterface() *SalesDataShipmentPackageCreationInterface {
	this := SalesDataShipmentPackageCreationInterface{}
	return &this
}

// NewSalesDataShipmentPackageCreationInterfaceWithDefaults instantiates a new SalesDataShipmentPackageCreationInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSalesDataShipmentPackageCreationInterfaceWithDefaults() *SalesDataShipmentPackageCreationInterface {
	this := SalesDataShipmentPackageCreationInterface{}
	return &this
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *SalesDataShipmentPackageCreationInterface) GetExtensionAttributes() map[string]interface{} {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataShipmentPackageCreationInterface) GetExtensionAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return map[string]interface{}{}, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *SalesDataShipmentPackageCreationInterface) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given map[string]interface{} and assigns it to the ExtensionAttributes field.
func (o *SalesDataShipmentPackageCreationInterface) SetExtensionAttributes(v map[string]interface{}) {
	o.ExtensionAttributes = v
}

func (o SalesDataShipmentPackageCreationInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SalesDataShipmentPackageCreationInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extension_attributes"] = o.ExtensionAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SalesDataShipmentPackageCreationInterface) UnmarshalJSON(data []byte) (err error) {
	varSalesDataShipmentPackageCreationInterface := _SalesDataShipmentPackageCreationInterface{}

	err = json.Unmarshal(data, &varSalesDataShipmentPackageCreationInterface)

	if err != nil {
		return err
	}

	*o = SalesDataShipmentPackageCreationInterface(varSalesDataShipmentPackageCreationInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "extension_attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *SalesDataShipmentPackageCreationInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *SalesDataShipmentPackageCreationInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableSalesDataShipmentPackageCreationInterface struct {
	value *SalesDataShipmentPackageCreationInterface
	isSet bool
}

func (v NullableSalesDataShipmentPackageCreationInterface) Get() *SalesDataShipmentPackageCreationInterface {
	return v.value
}

func (v *NullableSalesDataShipmentPackageCreationInterface) Set(val *SalesDataShipmentPackageCreationInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableSalesDataShipmentPackageCreationInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableSalesDataShipmentPackageCreationInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSalesDataShipmentPackageCreationInterface(val *SalesDataShipmentPackageCreationInterface) *NullableSalesDataShipmentPackageCreationInterface {
	return &NullableSalesDataShipmentPackageCreationInterface{value: val, isSet: true}
}

func (v NullableSalesDataShipmentPackageCreationInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSalesDataShipmentPackageCreationInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
