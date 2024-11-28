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

// checks if the TwoFactorAuthDataAuthyDeviceInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TwoFactorAuthDataAuthyDeviceInterface{}

// TwoFactorAuthDataAuthyDeviceInterface Authy device data interface
type TwoFactorAuthDataAuthyDeviceInterface struct {
	// The country
	Country string `json:"country"`
	// The phone number
	PhoneNumber string `json:"phone_number"`
	// The method to authenticate with
	Method string `json:"method"`
	// ExtensionInterface class for @see \\Magento\\TwoFactorAuth\\Api\\Data\\AuthyDeviceInterface
	ExtensionAttributes  map[string]interface{} `json:"extension_attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TwoFactorAuthDataAuthyDeviceInterface TwoFactorAuthDataAuthyDeviceInterface

// NewTwoFactorAuthDataAuthyDeviceInterface instantiates a new TwoFactorAuthDataAuthyDeviceInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTwoFactorAuthDataAuthyDeviceInterface(country string, phoneNumber string, method string) *TwoFactorAuthDataAuthyDeviceInterface {
	this := TwoFactorAuthDataAuthyDeviceInterface{}
	this.Country = country
	this.PhoneNumber = phoneNumber
	this.Method = method
	return &this
}

// NewTwoFactorAuthDataAuthyDeviceInterfaceWithDefaults instantiates a new TwoFactorAuthDataAuthyDeviceInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTwoFactorAuthDataAuthyDeviceInterfaceWithDefaults() *TwoFactorAuthDataAuthyDeviceInterface {
	this := TwoFactorAuthDataAuthyDeviceInterface{}
	return &this
}

// GetCountry returns the Country field value
func (o *TwoFactorAuthDataAuthyDeviceInterface) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *TwoFactorAuthDataAuthyDeviceInterface) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *TwoFactorAuthDataAuthyDeviceInterface) SetCountry(v string) {
	o.Country = v
}

// GetPhoneNumber returns the PhoneNumber field value
func (o *TwoFactorAuthDataAuthyDeviceInterface) GetPhoneNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value
// and a boolean to check if the value has been set.
func (o *TwoFactorAuthDataAuthyDeviceInterface) GetPhoneNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PhoneNumber, true
}

// SetPhoneNumber sets field value
func (o *TwoFactorAuthDataAuthyDeviceInterface) SetPhoneNumber(v string) {
	o.PhoneNumber = v
}

// GetMethod returns the Method field value
func (o *TwoFactorAuthDataAuthyDeviceInterface) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *TwoFactorAuthDataAuthyDeviceInterface) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *TwoFactorAuthDataAuthyDeviceInterface) SetMethod(v string) {
	o.Method = v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *TwoFactorAuthDataAuthyDeviceInterface) GetExtensionAttributes() map[string]interface{} {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwoFactorAuthDataAuthyDeviceInterface) GetExtensionAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return map[string]interface{}{}, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *TwoFactorAuthDataAuthyDeviceInterface) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given map[string]interface{} and assigns it to the ExtensionAttributes field.
func (o *TwoFactorAuthDataAuthyDeviceInterface) SetExtensionAttributes(v map[string]interface{}) {
	o.ExtensionAttributes = v
}

func (o TwoFactorAuthDataAuthyDeviceInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TwoFactorAuthDataAuthyDeviceInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["country"] = o.Country
	toSerialize["phone_number"] = o.PhoneNumber
	toSerialize["method"] = o.Method
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extension_attributes"] = o.ExtensionAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TwoFactorAuthDataAuthyDeviceInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"country",
		"phone_number",
		"method",
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

	varTwoFactorAuthDataAuthyDeviceInterface := _TwoFactorAuthDataAuthyDeviceInterface{}

	err = json.Unmarshal(data, &varTwoFactorAuthDataAuthyDeviceInterface)

	if err != nil {
		return err
	}

	*o = TwoFactorAuthDataAuthyDeviceInterface(varTwoFactorAuthDataAuthyDeviceInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "country")
		delete(additionalProperties, "phone_number")
		delete(additionalProperties, "method")
		delete(additionalProperties, "extension_attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *TwoFactorAuthDataAuthyDeviceInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *TwoFactorAuthDataAuthyDeviceInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableTwoFactorAuthDataAuthyDeviceInterface struct {
	value *TwoFactorAuthDataAuthyDeviceInterface
	isSet bool
}

func (v NullableTwoFactorAuthDataAuthyDeviceInterface) Get() *TwoFactorAuthDataAuthyDeviceInterface {
	return v.value
}

func (v *NullableTwoFactorAuthDataAuthyDeviceInterface) Set(val *TwoFactorAuthDataAuthyDeviceInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableTwoFactorAuthDataAuthyDeviceInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableTwoFactorAuthDataAuthyDeviceInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTwoFactorAuthDataAuthyDeviceInterface(val *TwoFactorAuthDataAuthyDeviceInterface) *NullableTwoFactorAuthDataAuthyDeviceInterface {
	return &NullableTwoFactorAuthDataAuthyDeviceInterface{value: val, isSet: true}
}

func (v NullableTwoFactorAuthDataAuthyDeviceInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTwoFactorAuthDataAuthyDeviceInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}