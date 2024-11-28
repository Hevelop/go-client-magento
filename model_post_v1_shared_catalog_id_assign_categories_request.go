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

// checks if the PostV1SharedCatalogIdAssignCategoriesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostV1SharedCatalogIdAssignCategoriesRequest{}

// PostV1SharedCatalogIdAssignCategoriesRequest struct for PostV1SharedCatalogIdAssignCategoriesRequest
type PostV1SharedCatalogIdAssignCategoriesRequest struct {
	Categories           []CatalogDataCategoryInterface `json:"categories"`
	AdditionalProperties map[string]interface{}
}

type _PostV1SharedCatalogIdAssignCategoriesRequest PostV1SharedCatalogIdAssignCategoriesRequest

// NewPostV1SharedCatalogIdAssignCategoriesRequest instantiates a new PostV1SharedCatalogIdAssignCategoriesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostV1SharedCatalogIdAssignCategoriesRequest(categories []CatalogDataCategoryInterface) *PostV1SharedCatalogIdAssignCategoriesRequest {
	this := PostV1SharedCatalogIdAssignCategoriesRequest{}
	this.Categories = categories
	return &this
}

// NewPostV1SharedCatalogIdAssignCategoriesRequestWithDefaults instantiates a new PostV1SharedCatalogIdAssignCategoriesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostV1SharedCatalogIdAssignCategoriesRequestWithDefaults() *PostV1SharedCatalogIdAssignCategoriesRequest {
	this := PostV1SharedCatalogIdAssignCategoriesRequest{}
	return &this
}

// GetCategories returns the Categories field value
func (o *PostV1SharedCatalogIdAssignCategoriesRequest) GetCategories() []CatalogDataCategoryInterface {
	if o == nil {
		var ret []CatalogDataCategoryInterface
		return ret
	}

	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value
// and a boolean to check if the value has been set.
func (o *PostV1SharedCatalogIdAssignCategoriesRequest) GetCategoriesOk() ([]CatalogDataCategoryInterface, bool) {
	if o == nil {
		return nil, false
	}
	return o.Categories, true
}

// SetCategories sets field value
func (o *PostV1SharedCatalogIdAssignCategoriesRequest) SetCategories(v []CatalogDataCategoryInterface) {
	o.Categories = v
}

func (o PostV1SharedCatalogIdAssignCategoriesRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostV1SharedCatalogIdAssignCategoriesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["categories"] = o.Categories

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostV1SharedCatalogIdAssignCategoriesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"categories",
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

	varPostV1SharedCatalogIdAssignCategoriesRequest := _PostV1SharedCatalogIdAssignCategoriesRequest{}

	err = json.Unmarshal(data, &varPostV1SharedCatalogIdAssignCategoriesRequest)

	if err != nil {
		return err
	}

	*o = PostV1SharedCatalogIdAssignCategoriesRequest(varPostV1SharedCatalogIdAssignCategoriesRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "categories")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PostV1SharedCatalogIdAssignCategoriesRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *PostV1SharedCatalogIdAssignCategoriesRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullablePostV1SharedCatalogIdAssignCategoriesRequest struct {
	value *PostV1SharedCatalogIdAssignCategoriesRequest
	isSet bool
}

func (v NullablePostV1SharedCatalogIdAssignCategoriesRequest) Get() *PostV1SharedCatalogIdAssignCategoriesRequest {
	return v.value
}

func (v *NullablePostV1SharedCatalogIdAssignCategoriesRequest) Set(val *PostV1SharedCatalogIdAssignCategoriesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostV1SharedCatalogIdAssignCategoriesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostV1SharedCatalogIdAssignCategoriesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostV1SharedCatalogIdAssignCategoriesRequest(val *PostV1SharedCatalogIdAssignCategoriesRequest) *NullablePostV1SharedCatalogIdAssignCategoriesRequest {
	return &NullablePostV1SharedCatalogIdAssignCategoriesRequest{value: val, isSet: true}
}

func (v NullablePostV1SharedCatalogIdAssignCategoriesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostV1SharedCatalogIdAssignCategoriesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}