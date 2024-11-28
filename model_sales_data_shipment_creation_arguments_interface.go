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

// checks if the SalesDataShipmentCreationArgumentsInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SalesDataShipmentCreationArgumentsInterface{}

// SalesDataShipmentCreationArgumentsInterface Interface for creation arguments for Shipment.
type SalesDataShipmentCreationArgumentsInterface struct {
	ExtensionAttributes  *SalesDataShipmentCreationArgumentsExtensionInterface `json:"extension_attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SalesDataShipmentCreationArgumentsInterface SalesDataShipmentCreationArgumentsInterface

// NewSalesDataShipmentCreationArgumentsInterface instantiates a new SalesDataShipmentCreationArgumentsInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSalesDataShipmentCreationArgumentsInterface() *SalesDataShipmentCreationArgumentsInterface {
	this := SalesDataShipmentCreationArgumentsInterface{}
	return &this
}

// NewSalesDataShipmentCreationArgumentsInterfaceWithDefaults instantiates a new SalesDataShipmentCreationArgumentsInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSalesDataShipmentCreationArgumentsInterfaceWithDefaults() *SalesDataShipmentCreationArgumentsInterface {
	this := SalesDataShipmentCreationArgumentsInterface{}
	return &this
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *SalesDataShipmentCreationArgumentsInterface) GetExtensionAttributes() SalesDataShipmentCreationArgumentsExtensionInterface {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret SalesDataShipmentCreationArgumentsExtensionInterface
		return ret
	}
	return *o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataShipmentCreationArgumentsInterface) GetExtensionAttributesOk() (*SalesDataShipmentCreationArgumentsExtensionInterface, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return nil, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *SalesDataShipmentCreationArgumentsInterface) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given SalesDataShipmentCreationArgumentsExtensionInterface and assigns it to the ExtensionAttributes field.
func (o *SalesDataShipmentCreationArgumentsInterface) SetExtensionAttributes(v SalesDataShipmentCreationArgumentsExtensionInterface) {
	o.ExtensionAttributes = &v
}

func (o SalesDataShipmentCreationArgumentsInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SalesDataShipmentCreationArgumentsInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extension_attributes"] = o.ExtensionAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SalesDataShipmentCreationArgumentsInterface) UnmarshalJSON(data []byte) (err error) {
	varSalesDataShipmentCreationArgumentsInterface := _SalesDataShipmentCreationArgumentsInterface{}

	err = json.Unmarshal(data, &varSalesDataShipmentCreationArgumentsInterface)

	if err != nil {
		return err
	}

	*o = SalesDataShipmentCreationArgumentsInterface(varSalesDataShipmentCreationArgumentsInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "extension_attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *SalesDataShipmentCreationArgumentsInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *SalesDataShipmentCreationArgumentsInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableSalesDataShipmentCreationArgumentsInterface struct {
	value *SalesDataShipmentCreationArgumentsInterface
	isSet bool
}

func (v NullableSalesDataShipmentCreationArgumentsInterface) Get() *SalesDataShipmentCreationArgumentsInterface {
	return v.value
}

func (v *NullableSalesDataShipmentCreationArgumentsInterface) Set(val *SalesDataShipmentCreationArgumentsInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableSalesDataShipmentCreationArgumentsInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableSalesDataShipmentCreationArgumentsInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSalesDataShipmentCreationArgumentsInterface(val *SalesDataShipmentCreationArgumentsInterface) *NullableSalesDataShipmentCreationArgumentsInterface {
	return &NullableSalesDataShipmentCreationArgumentsInterface{value: val, isSet: true}
}

func (v NullableSalesDataShipmentCreationArgumentsInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSalesDataShipmentCreationArgumentsInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
