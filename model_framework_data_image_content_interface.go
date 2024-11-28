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

// checks if the FrameworkDataImageContentInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FrameworkDataImageContentInterface{}

// FrameworkDataImageContentInterface Image Content data interface
type FrameworkDataImageContentInterface struct {
	// Media data (base64 encoded content)
	Base64EncodedData string `json:"base64_encoded_data"`
	// MIME type
	Type string `json:"type"`
	// Image name
	Name                 string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _FrameworkDataImageContentInterface FrameworkDataImageContentInterface

// NewFrameworkDataImageContentInterface instantiates a new FrameworkDataImageContentInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFrameworkDataImageContentInterface(base64EncodedData string, type_ string, name string) *FrameworkDataImageContentInterface {
	this := FrameworkDataImageContentInterface{}
	this.Base64EncodedData = base64EncodedData
	this.Type = type_
	this.Name = name
	return &this
}

// NewFrameworkDataImageContentInterfaceWithDefaults instantiates a new FrameworkDataImageContentInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFrameworkDataImageContentInterfaceWithDefaults() *FrameworkDataImageContentInterface {
	this := FrameworkDataImageContentInterface{}
	return &this
}

// GetBase64EncodedData returns the Base64EncodedData field value
func (o *FrameworkDataImageContentInterface) GetBase64EncodedData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Base64EncodedData
}

// GetBase64EncodedDataOk returns a tuple with the Base64EncodedData field value
// and a boolean to check if the value has been set.
func (o *FrameworkDataImageContentInterface) GetBase64EncodedDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Base64EncodedData, true
}

// SetBase64EncodedData sets field value
func (o *FrameworkDataImageContentInterface) SetBase64EncodedData(v string) {
	o.Base64EncodedData = v
}

// GetType returns the Type field value
func (o *FrameworkDataImageContentInterface) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FrameworkDataImageContentInterface) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FrameworkDataImageContentInterface) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value
func (o *FrameworkDataImageContentInterface) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FrameworkDataImageContentInterface) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FrameworkDataImageContentInterface) SetName(v string) {
	o.Name = v
}

func (o FrameworkDataImageContentInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FrameworkDataImageContentInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["base64_encoded_data"] = o.Base64EncodedData
	toSerialize["type"] = o.Type
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FrameworkDataImageContentInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"base64_encoded_data",
		"type",
		"name",
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

	varFrameworkDataImageContentInterface := _FrameworkDataImageContentInterface{}

	err = json.Unmarshal(data, &varFrameworkDataImageContentInterface)

	if err != nil {
		return err
	}

	*o = FrameworkDataImageContentInterface(varFrameworkDataImageContentInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "base64_encoded_data")
		delete(additionalProperties, "type")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *FrameworkDataImageContentInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *FrameworkDataImageContentInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableFrameworkDataImageContentInterface struct {
	value *FrameworkDataImageContentInterface
	isSet bool
}

func (v NullableFrameworkDataImageContentInterface) Get() *FrameworkDataImageContentInterface {
	return v.value
}

func (v *NullableFrameworkDataImageContentInterface) Set(val *FrameworkDataImageContentInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableFrameworkDataImageContentInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableFrameworkDataImageContentInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFrameworkDataImageContentInterface(val *FrameworkDataImageContentInterface) *NullableFrameworkDataImageContentInterface {
	return &NullableFrameworkDataImageContentInterface{value: val, isSet: true}
}

func (v NullableFrameworkDataImageContentInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFrameworkDataImageContentInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}