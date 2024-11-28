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

// checks if the InventorySalesApiDataProductSalabilityErrorInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventorySalesApiDataProductSalabilityErrorInterface{}

// InventorySalesApiDataProductSalabilityErrorInterface
type InventorySalesApiDataProductSalabilityErrorInterface struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	// ExtensionInterface class for @see \\Magento\\InventorySalesApi\\Api\\Data\\ProductSalabilityErrorInterface
	ExtensionAttributes  map[string]interface{} `json:"extension_attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InventorySalesApiDataProductSalabilityErrorInterface InventorySalesApiDataProductSalabilityErrorInterface

// NewInventorySalesApiDataProductSalabilityErrorInterface instantiates a new InventorySalesApiDataProductSalabilityErrorInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventorySalesApiDataProductSalabilityErrorInterface(code string, message string) *InventorySalesApiDataProductSalabilityErrorInterface {
	this := InventorySalesApiDataProductSalabilityErrorInterface{}
	this.Code = code
	this.Message = message
	return &this
}

// NewInventorySalesApiDataProductSalabilityErrorInterfaceWithDefaults instantiates a new InventorySalesApiDataProductSalabilityErrorInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventorySalesApiDataProductSalabilityErrorInterfaceWithDefaults() *InventorySalesApiDataProductSalabilityErrorInterface {
	this := InventorySalesApiDataProductSalabilityErrorInterface{}
	return &this
}

// GetCode returns the Code field value
func (o *InventorySalesApiDataProductSalabilityErrorInterface) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *InventorySalesApiDataProductSalabilityErrorInterface) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *InventorySalesApiDataProductSalabilityErrorInterface) SetCode(v string) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *InventorySalesApiDataProductSalabilityErrorInterface) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *InventorySalesApiDataProductSalabilityErrorInterface) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *InventorySalesApiDataProductSalabilityErrorInterface) SetMessage(v string) {
	o.Message = v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *InventorySalesApiDataProductSalabilityErrorInterface) GetExtensionAttributes() map[string]interface{} {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventorySalesApiDataProductSalabilityErrorInterface) GetExtensionAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return map[string]interface{}{}, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *InventorySalesApiDataProductSalabilityErrorInterface) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given map[string]interface{} and assigns it to the ExtensionAttributes field.
func (o *InventorySalesApiDataProductSalabilityErrorInterface) SetExtensionAttributes(v map[string]interface{}) {
	o.ExtensionAttributes = v
}

func (o InventorySalesApiDataProductSalabilityErrorInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventorySalesApiDataProductSalabilityErrorInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extension_attributes"] = o.ExtensionAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InventorySalesApiDataProductSalabilityErrorInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"message",
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

	varInventorySalesApiDataProductSalabilityErrorInterface := _InventorySalesApiDataProductSalabilityErrorInterface{}

	err = json.Unmarshal(data, &varInventorySalesApiDataProductSalabilityErrorInterface)

	if err != nil {
		return err
	}

	*o = InventorySalesApiDataProductSalabilityErrorInterface(varInventorySalesApiDataProductSalabilityErrorInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "message")
		delete(additionalProperties, "extension_attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *InventorySalesApiDataProductSalabilityErrorInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *InventorySalesApiDataProductSalabilityErrorInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableInventorySalesApiDataProductSalabilityErrorInterface struct {
	value *InventorySalesApiDataProductSalabilityErrorInterface
	isSet bool
}

func (v NullableInventorySalesApiDataProductSalabilityErrorInterface) Get() *InventorySalesApiDataProductSalabilityErrorInterface {
	return v.value
}

func (v *NullableInventorySalesApiDataProductSalabilityErrorInterface) Set(val *InventorySalesApiDataProductSalabilityErrorInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableInventorySalesApiDataProductSalabilityErrorInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableInventorySalesApiDataProductSalabilityErrorInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventorySalesApiDataProductSalabilityErrorInterface(val *InventorySalesApiDataProductSalabilityErrorInterface) *NullableInventorySalesApiDataProductSalabilityErrorInterface {
	return &NullableInventorySalesApiDataProductSalabilityErrorInterface{value: val, isSet: true}
}

func (v NullableInventorySalesApiDataProductSalabilityErrorInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventorySalesApiDataProductSalabilityErrorInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}