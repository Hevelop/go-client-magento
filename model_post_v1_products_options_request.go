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

// checks if the PostV1ProductsOptionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostV1ProductsOptionsRequest{}

// PostV1ProductsOptionsRequest struct for PostV1ProductsOptionsRequest
type PostV1ProductsOptionsRequest struct {
	Option               CatalogDataProductCustomOptionInterface `json:"option"`
	AdditionalProperties map[string]interface{}
}

type _PostV1ProductsOptionsRequest PostV1ProductsOptionsRequest

// NewPostV1ProductsOptionsRequest instantiates a new PostV1ProductsOptionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostV1ProductsOptionsRequest(option CatalogDataProductCustomOptionInterface) *PostV1ProductsOptionsRequest {
	this := PostV1ProductsOptionsRequest{}
	this.Option = option
	return &this
}

// NewPostV1ProductsOptionsRequestWithDefaults instantiates a new PostV1ProductsOptionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostV1ProductsOptionsRequestWithDefaults() *PostV1ProductsOptionsRequest {
	this := PostV1ProductsOptionsRequest{}
	return &this
}

// GetOption returns the Option field value
func (o *PostV1ProductsOptionsRequest) GetOption() CatalogDataProductCustomOptionInterface {
	if o == nil {
		var ret CatalogDataProductCustomOptionInterface
		return ret
	}

	return o.Option
}

// GetOptionOk returns a tuple with the Option field value
// and a boolean to check if the value has been set.
func (o *PostV1ProductsOptionsRequest) GetOptionOk() (*CatalogDataProductCustomOptionInterface, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Option, true
}

// SetOption sets field value
func (o *PostV1ProductsOptionsRequest) SetOption(v CatalogDataProductCustomOptionInterface) {
	o.Option = v
}

func (o PostV1ProductsOptionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostV1ProductsOptionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["option"] = o.Option

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostV1ProductsOptionsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"option",
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

	varPostV1ProductsOptionsRequest := _PostV1ProductsOptionsRequest{}

	err = json.Unmarshal(data, &varPostV1ProductsOptionsRequest)

	if err != nil {
		return err
	}

	*o = PostV1ProductsOptionsRequest(varPostV1ProductsOptionsRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "option")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PostV1ProductsOptionsRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *PostV1ProductsOptionsRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullablePostV1ProductsOptionsRequest struct {
	value *PostV1ProductsOptionsRequest
	isSet bool
}

func (v NullablePostV1ProductsOptionsRequest) Get() *PostV1ProductsOptionsRequest {
	return v.value
}

func (v *NullablePostV1ProductsOptionsRequest) Set(val *PostV1ProductsOptionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostV1ProductsOptionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostV1ProductsOptionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostV1ProductsOptionsRequest(val *PostV1ProductsOptionsRequest) *NullablePostV1ProductsOptionsRequest {
	return &NullablePostV1ProductsOptionsRequest{value: val, isSet: true}
}

func (v NullablePostV1ProductsOptionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostV1ProductsOptionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
