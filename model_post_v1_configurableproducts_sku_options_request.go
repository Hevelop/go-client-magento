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

// checks if the PostV1ConfigurableproductsSkuOptionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostV1ConfigurableproductsSkuOptionsRequest{}

// PostV1ConfigurableproductsSkuOptionsRequest struct for PostV1ConfigurableproductsSkuOptionsRequest
type PostV1ConfigurableproductsSkuOptionsRequest struct {
	Option               ConfigurableProductDataOptionInterface `json:"option"`
	AdditionalProperties map[string]interface{}
}

type _PostV1ConfigurableproductsSkuOptionsRequest PostV1ConfigurableproductsSkuOptionsRequest

// NewPostV1ConfigurableproductsSkuOptionsRequest instantiates a new PostV1ConfigurableproductsSkuOptionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostV1ConfigurableproductsSkuOptionsRequest(option ConfigurableProductDataOptionInterface) *PostV1ConfigurableproductsSkuOptionsRequest {
	this := PostV1ConfigurableproductsSkuOptionsRequest{}
	this.Option = option
	return &this
}

// NewPostV1ConfigurableproductsSkuOptionsRequestWithDefaults instantiates a new PostV1ConfigurableproductsSkuOptionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostV1ConfigurableproductsSkuOptionsRequestWithDefaults() *PostV1ConfigurableproductsSkuOptionsRequest {
	this := PostV1ConfigurableproductsSkuOptionsRequest{}
	return &this
}

// GetOption returns the Option field value
func (o *PostV1ConfigurableproductsSkuOptionsRequest) GetOption() ConfigurableProductDataOptionInterface {
	if o == nil {
		var ret ConfigurableProductDataOptionInterface
		return ret
	}

	return o.Option
}

// GetOptionOk returns a tuple with the Option field value
// and a boolean to check if the value has been set.
func (o *PostV1ConfigurableproductsSkuOptionsRequest) GetOptionOk() (*ConfigurableProductDataOptionInterface, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Option, true
}

// SetOption sets field value
func (o *PostV1ConfigurableproductsSkuOptionsRequest) SetOption(v ConfigurableProductDataOptionInterface) {
	o.Option = v
}

func (o PostV1ConfigurableproductsSkuOptionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostV1ConfigurableproductsSkuOptionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["option"] = o.Option

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostV1ConfigurableproductsSkuOptionsRequest) UnmarshalJSON(data []byte) (err error) {
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

	varPostV1ConfigurableproductsSkuOptionsRequest := _PostV1ConfigurableproductsSkuOptionsRequest{}

	err = json.Unmarshal(data, &varPostV1ConfigurableproductsSkuOptionsRequest)

	if err != nil {
		return err
	}

	*o = PostV1ConfigurableproductsSkuOptionsRequest(varPostV1ConfigurableproductsSkuOptionsRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "option")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PostV1ConfigurableproductsSkuOptionsRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *PostV1ConfigurableproductsSkuOptionsRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullablePostV1ConfigurableproductsSkuOptionsRequest struct {
	value *PostV1ConfigurableproductsSkuOptionsRequest
	isSet bool
}

func (v NullablePostV1ConfigurableproductsSkuOptionsRequest) Get() *PostV1ConfigurableproductsSkuOptionsRequest {
	return v.value
}

func (v *NullablePostV1ConfigurableproductsSkuOptionsRequest) Set(val *PostV1ConfigurableproductsSkuOptionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostV1ConfigurableproductsSkuOptionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostV1ConfigurableproductsSkuOptionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostV1ConfigurableproductsSkuOptionsRequest(val *PostV1ConfigurableproductsSkuOptionsRequest) *NullablePostV1ConfigurableproductsSkuOptionsRequest {
	return &NullablePostV1ConfigurableproductsSkuOptionsRequest{value: val, isSet: true}
}

func (v NullablePostV1ConfigurableproductsSkuOptionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostV1ConfigurableproductsSkuOptionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
