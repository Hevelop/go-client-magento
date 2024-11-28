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

// checks if the InventoryApiDataSourceCarrierLinkInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryApiDataSourceCarrierLinkInterface{}

// InventoryApiDataSourceCarrierLinkInterface Represents relation between some physical storage and shipping method Used fully qualified namespaces in annotations for proper work of WebApi request parser
type InventoryApiDataSourceCarrierLinkInterface struct {
	// Carrier code
	CarrierCode *string `json:"carrier_code,omitempty"`
	// Position
	Position *int32 `json:"position,omitempty"`
	// ExtensionInterface class for @see \\Magento\\InventoryApi\\Api\\Data\\SourceCarrierLinkInterface
	ExtensionAttributes  map[string]interface{} `json:"extension_attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InventoryApiDataSourceCarrierLinkInterface InventoryApiDataSourceCarrierLinkInterface

// NewInventoryApiDataSourceCarrierLinkInterface instantiates a new InventoryApiDataSourceCarrierLinkInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryApiDataSourceCarrierLinkInterface() *InventoryApiDataSourceCarrierLinkInterface {
	this := InventoryApiDataSourceCarrierLinkInterface{}
	return &this
}

// NewInventoryApiDataSourceCarrierLinkInterfaceWithDefaults instantiates a new InventoryApiDataSourceCarrierLinkInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryApiDataSourceCarrierLinkInterfaceWithDefaults() *InventoryApiDataSourceCarrierLinkInterface {
	this := InventoryApiDataSourceCarrierLinkInterface{}
	return &this
}

// GetCarrierCode returns the CarrierCode field value if set, zero value otherwise.
func (o *InventoryApiDataSourceCarrierLinkInterface) GetCarrierCode() string {
	if o == nil || IsNil(o.CarrierCode) {
		var ret string
		return ret
	}
	return *o.CarrierCode
}

// GetCarrierCodeOk returns a tuple with the CarrierCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryApiDataSourceCarrierLinkInterface) GetCarrierCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CarrierCode) {
		return nil, false
	}
	return o.CarrierCode, true
}

// HasCarrierCode returns a boolean if a field has been set.
func (o *InventoryApiDataSourceCarrierLinkInterface) HasCarrierCode() bool {
	if o != nil && !IsNil(o.CarrierCode) {
		return true
	}

	return false
}

// SetCarrierCode gets a reference to the given string and assigns it to the CarrierCode field.
func (o *InventoryApiDataSourceCarrierLinkInterface) SetCarrierCode(v string) {
	o.CarrierCode = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *InventoryApiDataSourceCarrierLinkInterface) GetPosition() int32 {
	if o == nil || IsNil(o.Position) {
		var ret int32
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryApiDataSourceCarrierLinkInterface) GetPositionOk() (*int32, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *InventoryApiDataSourceCarrierLinkInterface) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int32 and assigns it to the Position field.
func (o *InventoryApiDataSourceCarrierLinkInterface) SetPosition(v int32) {
	o.Position = &v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *InventoryApiDataSourceCarrierLinkInterface) GetExtensionAttributes() map[string]interface{} {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryApiDataSourceCarrierLinkInterface) GetExtensionAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return map[string]interface{}{}, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *InventoryApiDataSourceCarrierLinkInterface) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given map[string]interface{} and assigns it to the ExtensionAttributes field.
func (o *InventoryApiDataSourceCarrierLinkInterface) SetExtensionAttributes(v map[string]interface{}) {
	o.ExtensionAttributes = v
}

func (o InventoryApiDataSourceCarrierLinkInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventoryApiDataSourceCarrierLinkInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CarrierCode) {
		toSerialize["carrier_code"] = o.CarrierCode
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extension_attributes"] = o.ExtensionAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InventoryApiDataSourceCarrierLinkInterface) UnmarshalJSON(data []byte) (err error) {
	varInventoryApiDataSourceCarrierLinkInterface := _InventoryApiDataSourceCarrierLinkInterface{}

	err = json.Unmarshal(data, &varInventoryApiDataSourceCarrierLinkInterface)

	if err != nil {
		return err
	}

	*o = InventoryApiDataSourceCarrierLinkInterface(varInventoryApiDataSourceCarrierLinkInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "carrier_code")
		delete(additionalProperties, "position")
		delete(additionalProperties, "extension_attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *InventoryApiDataSourceCarrierLinkInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *InventoryApiDataSourceCarrierLinkInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableInventoryApiDataSourceCarrierLinkInterface struct {
	value *InventoryApiDataSourceCarrierLinkInterface
	isSet bool
}

func (v NullableInventoryApiDataSourceCarrierLinkInterface) Get() *InventoryApiDataSourceCarrierLinkInterface {
	return v.value
}

func (v *NullableInventoryApiDataSourceCarrierLinkInterface) Set(val *InventoryApiDataSourceCarrierLinkInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryApiDataSourceCarrierLinkInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryApiDataSourceCarrierLinkInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryApiDataSourceCarrierLinkInterface(val *InventoryApiDataSourceCarrierLinkInterface) *NullableInventoryApiDataSourceCarrierLinkInterface {
	return &NullableInventoryApiDataSourceCarrierLinkInterface{value: val, isSet: true}
}

func (v NullableInventoryApiDataSourceCarrierLinkInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryApiDataSourceCarrierLinkInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
