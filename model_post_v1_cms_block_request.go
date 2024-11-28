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

// checks if the PostV1CmsBlockRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostV1CmsBlockRequest{}

// PostV1CmsBlockRequest struct for PostV1CmsBlockRequest
type PostV1CmsBlockRequest struct {
	Block                CmsDataBlockInterface `json:"block"`
	AdditionalProperties map[string]interface{}
}

type _PostV1CmsBlockRequest PostV1CmsBlockRequest

// NewPostV1CmsBlockRequest instantiates a new PostV1CmsBlockRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostV1CmsBlockRequest(block CmsDataBlockInterface) *PostV1CmsBlockRequest {
	this := PostV1CmsBlockRequest{}
	this.Block = block
	return &this
}

// NewPostV1CmsBlockRequestWithDefaults instantiates a new PostV1CmsBlockRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostV1CmsBlockRequestWithDefaults() *PostV1CmsBlockRequest {
	this := PostV1CmsBlockRequest{}
	return &this
}

// GetBlock returns the Block field value
func (o *PostV1CmsBlockRequest) GetBlock() CmsDataBlockInterface {
	if o == nil {
		var ret CmsDataBlockInterface
		return ret
	}

	return o.Block
}

// GetBlockOk returns a tuple with the Block field value
// and a boolean to check if the value has been set.
func (o *PostV1CmsBlockRequest) GetBlockOk() (*CmsDataBlockInterface, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Block, true
}

// SetBlock sets field value
func (o *PostV1CmsBlockRequest) SetBlock(v CmsDataBlockInterface) {
	o.Block = v
}

func (o PostV1CmsBlockRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostV1CmsBlockRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["block"] = o.Block

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostV1CmsBlockRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"block",
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

	varPostV1CmsBlockRequest := _PostV1CmsBlockRequest{}

	err = json.Unmarshal(data, &varPostV1CmsBlockRequest)

	if err != nil {
		return err
	}

	*o = PostV1CmsBlockRequest(varPostV1CmsBlockRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "block")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PostV1CmsBlockRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *PostV1CmsBlockRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullablePostV1CmsBlockRequest struct {
	value *PostV1CmsBlockRequest
	isSet bool
}

func (v NullablePostV1CmsBlockRequest) Get() *PostV1CmsBlockRequest {
	return v.value
}

func (v *NullablePostV1CmsBlockRequest) Set(val *PostV1CmsBlockRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostV1CmsBlockRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostV1CmsBlockRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostV1CmsBlockRequest(val *PostV1CmsBlockRequest) *NullablePostV1CmsBlockRequest {
	return &NullablePostV1CmsBlockRequest{value: val, isSet: true}
}

func (v NullablePostV1CmsBlockRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostV1CmsBlockRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}