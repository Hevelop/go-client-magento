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

// checks if the PostV1ProductsAttributesetsAttributesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostV1ProductsAttributesetsAttributesRequest{}

// PostV1ProductsAttributesetsAttributesRequest struct for PostV1ProductsAttributesetsAttributesRequest
type PostV1ProductsAttributesetsAttributesRequest struct {
	AttributeSetId       int32  `json:"attributeSetId"`
	AttributeGroupId     int32  `json:"attributeGroupId"`
	AttributeCode        string `json:"attributeCode"`
	SortOrder            int32  `json:"sortOrder"`
	AdditionalProperties map[string]interface{}
}

type _PostV1ProductsAttributesetsAttributesRequest PostV1ProductsAttributesetsAttributesRequest

// NewPostV1ProductsAttributesetsAttributesRequest instantiates a new PostV1ProductsAttributesetsAttributesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostV1ProductsAttributesetsAttributesRequest(attributeSetId int32, attributeGroupId int32, attributeCode string, sortOrder int32) *PostV1ProductsAttributesetsAttributesRequest {
	this := PostV1ProductsAttributesetsAttributesRequest{}
	this.AttributeSetId = attributeSetId
	this.AttributeGroupId = attributeGroupId
	this.AttributeCode = attributeCode
	this.SortOrder = sortOrder
	return &this
}

// NewPostV1ProductsAttributesetsAttributesRequestWithDefaults instantiates a new PostV1ProductsAttributesetsAttributesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostV1ProductsAttributesetsAttributesRequestWithDefaults() *PostV1ProductsAttributesetsAttributesRequest {
	this := PostV1ProductsAttributesetsAttributesRequest{}
	return &this
}

// GetAttributeSetId returns the AttributeSetId field value
func (o *PostV1ProductsAttributesetsAttributesRequest) GetAttributeSetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AttributeSetId
}

// GetAttributeSetIdOk returns a tuple with the AttributeSetId field value
// and a boolean to check if the value has been set.
func (o *PostV1ProductsAttributesetsAttributesRequest) GetAttributeSetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeSetId, true
}

// SetAttributeSetId sets field value
func (o *PostV1ProductsAttributesetsAttributesRequest) SetAttributeSetId(v int32) {
	o.AttributeSetId = v
}

// GetAttributeGroupId returns the AttributeGroupId field value
func (o *PostV1ProductsAttributesetsAttributesRequest) GetAttributeGroupId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AttributeGroupId
}

// GetAttributeGroupIdOk returns a tuple with the AttributeGroupId field value
// and a boolean to check if the value has been set.
func (o *PostV1ProductsAttributesetsAttributesRequest) GetAttributeGroupIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeGroupId, true
}

// SetAttributeGroupId sets field value
func (o *PostV1ProductsAttributesetsAttributesRequest) SetAttributeGroupId(v int32) {
	o.AttributeGroupId = v
}

// GetAttributeCode returns the AttributeCode field value
func (o *PostV1ProductsAttributesetsAttributesRequest) GetAttributeCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttributeCode
}

// GetAttributeCodeOk returns a tuple with the AttributeCode field value
// and a boolean to check if the value has been set.
func (o *PostV1ProductsAttributesetsAttributesRequest) GetAttributeCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeCode, true
}

// SetAttributeCode sets field value
func (o *PostV1ProductsAttributesetsAttributesRequest) SetAttributeCode(v string) {
	o.AttributeCode = v
}

// GetSortOrder returns the SortOrder field value
func (o *PostV1ProductsAttributesetsAttributesRequest) GetSortOrder() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value
// and a boolean to check if the value has been set.
func (o *PostV1ProductsAttributesetsAttributesRequest) GetSortOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SortOrder, true
}

// SetSortOrder sets field value
func (o *PostV1ProductsAttributesetsAttributesRequest) SetSortOrder(v int32) {
	o.SortOrder = v
}

func (o PostV1ProductsAttributesetsAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostV1ProductsAttributesetsAttributesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["attributeSetId"] = o.AttributeSetId
	toSerialize["attributeGroupId"] = o.AttributeGroupId
	toSerialize["attributeCode"] = o.AttributeCode
	toSerialize["sortOrder"] = o.SortOrder

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostV1ProductsAttributesetsAttributesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"attributeSetId",
		"attributeGroupId",
		"attributeCode",
		"sortOrder",
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

	varPostV1ProductsAttributesetsAttributesRequest := _PostV1ProductsAttributesetsAttributesRequest{}

	err = json.Unmarshal(data, &varPostV1ProductsAttributesetsAttributesRequest)

	if err != nil {
		return err
	}

	*o = PostV1ProductsAttributesetsAttributesRequest(varPostV1ProductsAttributesetsAttributesRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "attributeSetId")
		delete(additionalProperties, "attributeGroupId")
		delete(additionalProperties, "attributeCode")
		delete(additionalProperties, "sortOrder")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PostV1ProductsAttributesetsAttributesRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *PostV1ProductsAttributesetsAttributesRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullablePostV1ProductsAttributesetsAttributesRequest struct {
	value *PostV1ProductsAttributesetsAttributesRequest
	isSet bool
}

func (v NullablePostV1ProductsAttributesetsAttributesRequest) Get() *PostV1ProductsAttributesetsAttributesRequest {
	return v.value
}

func (v *NullablePostV1ProductsAttributesetsAttributesRequest) Set(val *PostV1ProductsAttributesetsAttributesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostV1ProductsAttributesetsAttributesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostV1ProductsAttributesetsAttributesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostV1ProductsAttributesetsAttributesRequest(val *PostV1ProductsAttributesetsAttributesRequest) *NullablePostV1ProductsAttributesetsAttributesRequest {
	return &NullablePostV1ProductsAttributesetsAttributesRequest{value: val, isSet: true}
}

func (v NullablePostV1ProductsAttributesetsAttributesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostV1ProductsAttributesetsAttributesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}