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

// checks if the PostV1SharedCatalogRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostV1SharedCatalogRequest{}

// PostV1SharedCatalogRequest struct for PostV1SharedCatalogRequest
type PostV1SharedCatalogRequest struct {
	SharedCatalog        SharedCatalogDataSharedCatalogInterface `json:"sharedCatalog"`
	AdditionalProperties map[string]interface{}
}

type _PostV1SharedCatalogRequest PostV1SharedCatalogRequest

// NewPostV1SharedCatalogRequest instantiates a new PostV1SharedCatalogRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostV1SharedCatalogRequest(sharedCatalog SharedCatalogDataSharedCatalogInterface) *PostV1SharedCatalogRequest {
	this := PostV1SharedCatalogRequest{}
	this.SharedCatalog = sharedCatalog
	return &this
}

// NewPostV1SharedCatalogRequestWithDefaults instantiates a new PostV1SharedCatalogRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostV1SharedCatalogRequestWithDefaults() *PostV1SharedCatalogRequest {
	this := PostV1SharedCatalogRequest{}
	return &this
}

// GetSharedCatalog returns the SharedCatalog field value
func (o *PostV1SharedCatalogRequest) GetSharedCatalog() SharedCatalogDataSharedCatalogInterface {
	if o == nil {
		var ret SharedCatalogDataSharedCatalogInterface
		return ret
	}

	return o.SharedCatalog
}

// GetSharedCatalogOk returns a tuple with the SharedCatalog field value
// and a boolean to check if the value has been set.
func (o *PostV1SharedCatalogRequest) GetSharedCatalogOk() (*SharedCatalogDataSharedCatalogInterface, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SharedCatalog, true
}

// SetSharedCatalog sets field value
func (o *PostV1SharedCatalogRequest) SetSharedCatalog(v SharedCatalogDataSharedCatalogInterface) {
	o.SharedCatalog = v
}

func (o PostV1SharedCatalogRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostV1SharedCatalogRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sharedCatalog"] = o.SharedCatalog

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostV1SharedCatalogRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sharedCatalog",
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

	varPostV1SharedCatalogRequest := _PostV1SharedCatalogRequest{}

	err = json.Unmarshal(data, &varPostV1SharedCatalogRequest)

	if err != nil {
		return err
	}

	*o = PostV1SharedCatalogRequest(varPostV1SharedCatalogRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "sharedCatalog")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PostV1SharedCatalogRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *PostV1SharedCatalogRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullablePostV1SharedCatalogRequest struct {
	value *PostV1SharedCatalogRequest
	isSet bool
}

func (v NullablePostV1SharedCatalogRequest) Get() *PostV1SharedCatalogRequest {
	return v.value
}

func (v *NullablePostV1SharedCatalogRequest) Set(val *PostV1SharedCatalogRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostV1SharedCatalogRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostV1SharedCatalogRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostV1SharedCatalogRequest(val *PostV1SharedCatalogRequest) *NullablePostV1SharedCatalogRequest {
	return &NullablePostV1SharedCatalogRequest{value: val, isSet: true}
}

func (v NullablePostV1SharedCatalogRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostV1SharedCatalogRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}