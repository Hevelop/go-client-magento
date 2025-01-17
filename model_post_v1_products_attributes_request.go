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

// checks if the PostV1ProductsAttributesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostV1ProductsAttributesRequest{}

// PostV1ProductsAttributesRequest struct for PostV1ProductsAttributesRequest
type PostV1ProductsAttributesRequest struct {
	Attribute            CatalogDataProductAttributeInterface `json:"attribute"`
	AdditionalProperties map[string]interface{}
}

type _PostV1ProductsAttributesRequest PostV1ProductsAttributesRequest

// NewPostV1ProductsAttributesRequest instantiates a new PostV1ProductsAttributesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostV1ProductsAttributesRequest(attribute CatalogDataProductAttributeInterface) *PostV1ProductsAttributesRequest {
	this := PostV1ProductsAttributesRequest{}
	this.Attribute = attribute
	return &this
}

// NewPostV1ProductsAttributesRequestWithDefaults instantiates a new PostV1ProductsAttributesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostV1ProductsAttributesRequestWithDefaults() *PostV1ProductsAttributesRequest {
	this := PostV1ProductsAttributesRequest{}
	return &this
}

// GetAttribute returns the Attribute field value
func (o *PostV1ProductsAttributesRequest) GetAttribute() CatalogDataProductAttributeInterface {
	if o == nil {
		var ret CatalogDataProductAttributeInterface
		return ret
	}

	return o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value
// and a boolean to check if the value has been set.
func (o *PostV1ProductsAttributesRequest) GetAttributeOk() (*CatalogDataProductAttributeInterface, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attribute, true
}

// SetAttribute sets field value
func (o *PostV1ProductsAttributesRequest) SetAttribute(v CatalogDataProductAttributeInterface) {
	o.Attribute = v
}

func (o PostV1ProductsAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostV1ProductsAttributesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["attribute"] = o.Attribute

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostV1ProductsAttributesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"attribute",
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

	varPostV1ProductsAttributesRequest := _PostV1ProductsAttributesRequest{}

	err = json.Unmarshal(data, &varPostV1ProductsAttributesRequest)

	if err != nil {
		return err
	}

	*o = PostV1ProductsAttributesRequest(varPostV1ProductsAttributesRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "attribute")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PostV1ProductsAttributesRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *PostV1ProductsAttributesRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullablePostV1ProductsAttributesRequest struct {
	value *PostV1ProductsAttributesRequest
	isSet bool
}

func (v NullablePostV1ProductsAttributesRequest) Get() *PostV1ProductsAttributesRequest {
	return v.value
}

func (v *NullablePostV1ProductsAttributesRequest) Set(val *PostV1ProductsAttributesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostV1ProductsAttributesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostV1ProductsAttributesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostV1ProductsAttributesRequest(val *PostV1ProductsAttributesRequest) *NullablePostV1ProductsAttributesRequest {
	return &NullablePostV1ProductsAttributesRequest{value: val, isSet: true}
}

func (v NullablePostV1ProductsAttributesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostV1ProductsAttributesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
