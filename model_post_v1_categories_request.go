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

// checks if the PostV1CategoriesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostV1CategoriesRequest{}

// PostV1CategoriesRequest struct for PostV1CategoriesRequest
type PostV1CategoriesRequest struct {
	Category             CatalogDataCategoryInterface `json:"category"`
	AdditionalProperties map[string]interface{}
}

type _PostV1CategoriesRequest PostV1CategoriesRequest

// NewPostV1CategoriesRequest instantiates a new PostV1CategoriesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostV1CategoriesRequest(category CatalogDataCategoryInterface) *PostV1CategoriesRequest {
	this := PostV1CategoriesRequest{}
	this.Category = category
	return &this
}

// NewPostV1CategoriesRequestWithDefaults instantiates a new PostV1CategoriesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostV1CategoriesRequestWithDefaults() *PostV1CategoriesRequest {
	this := PostV1CategoriesRequest{}
	return &this
}

// GetCategory returns the Category field value
func (o *PostV1CategoriesRequest) GetCategory() CatalogDataCategoryInterface {
	if o == nil {
		var ret CatalogDataCategoryInterface
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *PostV1CategoriesRequest) GetCategoryOk() (*CatalogDataCategoryInterface, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *PostV1CategoriesRequest) SetCategory(v CatalogDataCategoryInterface) {
	o.Category = v
}

func (o PostV1CategoriesRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostV1CategoriesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["category"] = o.Category

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostV1CategoriesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"category",
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

	varPostV1CategoriesRequest := _PostV1CategoriesRequest{}

	err = json.Unmarshal(data, &varPostV1CategoriesRequest)

	if err != nil {
		return err
	}

	*o = PostV1CategoriesRequest(varPostV1CategoriesRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "category")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PostV1CategoriesRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *PostV1CategoriesRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullablePostV1CategoriesRequest struct {
	value *PostV1CategoriesRequest
	isSet bool
}

func (v NullablePostV1CategoriesRequest) Get() *PostV1CategoriesRequest {
	return v.value
}

func (v *NullablePostV1CategoriesRequest) Set(val *PostV1CategoriesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostV1CategoriesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostV1CategoriesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostV1CategoriesRequest(val *PostV1CategoriesRequest) *NullablePostV1CategoriesRequest {
	return &NullablePostV1CategoriesRequest{value: val, isSet: true}
}

func (v NullablePostV1CategoriesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostV1CategoriesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
