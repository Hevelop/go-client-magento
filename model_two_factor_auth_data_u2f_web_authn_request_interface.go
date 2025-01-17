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

// checks if the TwoFactorAuthDataU2fWebAuthnRequestInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TwoFactorAuthDataU2fWebAuthnRequestInterface{}

// TwoFactorAuthDataU2fWebAuthnRequestInterface Represents a WebAuthn dataset
type TwoFactorAuthDataU2fWebAuthnRequestInterface struct {
	// The needed data to initiate a WebAuthn registration ceremony
	CredentialRequestOptionsJson string `json:"credential_request_options_json"`
	// ExtensionInterface class for @see \\Magento\\TwoFactorAuth\\Api\\Data\\U2fWebAuthnRequestInterface
	ExtensionAttributes  map[string]interface{} `json:"extension_attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TwoFactorAuthDataU2fWebAuthnRequestInterface TwoFactorAuthDataU2fWebAuthnRequestInterface

// NewTwoFactorAuthDataU2fWebAuthnRequestInterface instantiates a new TwoFactorAuthDataU2fWebAuthnRequestInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTwoFactorAuthDataU2fWebAuthnRequestInterface(credentialRequestOptionsJson string) *TwoFactorAuthDataU2fWebAuthnRequestInterface {
	this := TwoFactorAuthDataU2fWebAuthnRequestInterface{}
	this.CredentialRequestOptionsJson = credentialRequestOptionsJson
	return &this
}

// NewTwoFactorAuthDataU2fWebAuthnRequestInterfaceWithDefaults instantiates a new TwoFactorAuthDataU2fWebAuthnRequestInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTwoFactorAuthDataU2fWebAuthnRequestInterfaceWithDefaults() *TwoFactorAuthDataU2fWebAuthnRequestInterface {
	this := TwoFactorAuthDataU2fWebAuthnRequestInterface{}
	return &this
}

// GetCredentialRequestOptionsJson returns the CredentialRequestOptionsJson field value
func (o *TwoFactorAuthDataU2fWebAuthnRequestInterface) GetCredentialRequestOptionsJson() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CredentialRequestOptionsJson
}

// GetCredentialRequestOptionsJsonOk returns a tuple with the CredentialRequestOptionsJson field value
// and a boolean to check if the value has been set.
func (o *TwoFactorAuthDataU2fWebAuthnRequestInterface) GetCredentialRequestOptionsJsonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CredentialRequestOptionsJson, true
}

// SetCredentialRequestOptionsJson sets field value
func (o *TwoFactorAuthDataU2fWebAuthnRequestInterface) SetCredentialRequestOptionsJson(v string) {
	o.CredentialRequestOptionsJson = v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *TwoFactorAuthDataU2fWebAuthnRequestInterface) GetExtensionAttributes() map[string]interface{} {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwoFactorAuthDataU2fWebAuthnRequestInterface) GetExtensionAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return map[string]interface{}{}, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *TwoFactorAuthDataU2fWebAuthnRequestInterface) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given map[string]interface{} and assigns it to the ExtensionAttributes field.
func (o *TwoFactorAuthDataU2fWebAuthnRequestInterface) SetExtensionAttributes(v map[string]interface{}) {
	o.ExtensionAttributes = v
}

func (o TwoFactorAuthDataU2fWebAuthnRequestInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TwoFactorAuthDataU2fWebAuthnRequestInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["credential_request_options_json"] = o.CredentialRequestOptionsJson
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extension_attributes"] = o.ExtensionAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TwoFactorAuthDataU2fWebAuthnRequestInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"credential_request_options_json",
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

	varTwoFactorAuthDataU2fWebAuthnRequestInterface := _TwoFactorAuthDataU2fWebAuthnRequestInterface{}

	err = json.Unmarshal(data, &varTwoFactorAuthDataU2fWebAuthnRequestInterface)

	if err != nil {
		return err
	}

	*o = TwoFactorAuthDataU2fWebAuthnRequestInterface(varTwoFactorAuthDataU2fWebAuthnRequestInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "credential_request_options_json")
		delete(additionalProperties, "extension_attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *TwoFactorAuthDataU2fWebAuthnRequestInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *TwoFactorAuthDataU2fWebAuthnRequestInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableTwoFactorAuthDataU2fWebAuthnRequestInterface struct {
	value *TwoFactorAuthDataU2fWebAuthnRequestInterface
	isSet bool
}

func (v NullableTwoFactorAuthDataU2fWebAuthnRequestInterface) Get() *TwoFactorAuthDataU2fWebAuthnRequestInterface {
	return v.value
}

func (v *NullableTwoFactorAuthDataU2fWebAuthnRequestInterface) Set(val *TwoFactorAuthDataU2fWebAuthnRequestInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableTwoFactorAuthDataU2fWebAuthnRequestInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableTwoFactorAuthDataU2fWebAuthnRequestInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTwoFactorAuthDataU2fWebAuthnRequestInterface(val *TwoFactorAuthDataU2fWebAuthnRequestInterface) *NullableTwoFactorAuthDataU2fWebAuthnRequestInterface {
	return &NullableTwoFactorAuthDataU2fWebAuthnRequestInterface{value: val, isSet: true}
}

func (v NullableTwoFactorAuthDataU2fWebAuthnRequestInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTwoFactorAuthDataU2fWebAuthnRequestInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
