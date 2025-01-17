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

// checks if the PostV1EavAttributesetsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostV1EavAttributesetsRequest{}

// PostV1EavAttributesetsRequest struct for PostV1EavAttributesetsRequest
type PostV1EavAttributesetsRequest struct {
	EntityTypeCode       string                       `json:"entityTypeCode"`
	AttributeSet         EavDataAttributeSetInterface `json:"attributeSet"`
	SkeletonId           int32                        `json:"skeletonId"`
	AdditionalProperties map[string]interface{}
}

type _PostV1EavAttributesetsRequest PostV1EavAttributesetsRequest

// NewPostV1EavAttributesetsRequest instantiates a new PostV1EavAttributesetsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostV1EavAttributesetsRequest(entityTypeCode string, attributeSet EavDataAttributeSetInterface, skeletonId int32) *PostV1EavAttributesetsRequest {
	this := PostV1EavAttributesetsRequest{}
	this.EntityTypeCode = entityTypeCode
	this.AttributeSet = attributeSet
	this.SkeletonId = skeletonId
	return &this
}

// NewPostV1EavAttributesetsRequestWithDefaults instantiates a new PostV1EavAttributesetsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostV1EavAttributesetsRequestWithDefaults() *PostV1EavAttributesetsRequest {
	this := PostV1EavAttributesetsRequest{}
	return &this
}

// GetEntityTypeCode returns the EntityTypeCode field value
func (o *PostV1EavAttributesetsRequest) GetEntityTypeCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityTypeCode
}

// GetEntityTypeCodeOk returns a tuple with the EntityTypeCode field value
// and a boolean to check if the value has been set.
func (o *PostV1EavAttributesetsRequest) GetEntityTypeCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityTypeCode, true
}

// SetEntityTypeCode sets field value
func (o *PostV1EavAttributesetsRequest) SetEntityTypeCode(v string) {
	o.EntityTypeCode = v
}

// GetAttributeSet returns the AttributeSet field value
func (o *PostV1EavAttributesetsRequest) GetAttributeSet() EavDataAttributeSetInterface {
	if o == nil {
		var ret EavDataAttributeSetInterface
		return ret
	}

	return o.AttributeSet
}

// GetAttributeSetOk returns a tuple with the AttributeSet field value
// and a boolean to check if the value has been set.
func (o *PostV1EavAttributesetsRequest) GetAttributeSetOk() (*EavDataAttributeSetInterface, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeSet, true
}

// SetAttributeSet sets field value
func (o *PostV1EavAttributesetsRequest) SetAttributeSet(v EavDataAttributeSetInterface) {
	o.AttributeSet = v
}

// GetSkeletonId returns the SkeletonId field value
func (o *PostV1EavAttributesetsRequest) GetSkeletonId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SkeletonId
}

// GetSkeletonIdOk returns a tuple with the SkeletonId field value
// and a boolean to check if the value has been set.
func (o *PostV1EavAttributesetsRequest) GetSkeletonIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SkeletonId, true
}

// SetSkeletonId sets field value
func (o *PostV1EavAttributesetsRequest) SetSkeletonId(v int32) {
	o.SkeletonId = v
}

func (o PostV1EavAttributesetsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostV1EavAttributesetsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entityTypeCode"] = o.EntityTypeCode
	toSerialize["attributeSet"] = o.AttributeSet
	toSerialize["skeletonId"] = o.SkeletonId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostV1EavAttributesetsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"entityTypeCode",
		"attributeSet",
		"skeletonId",
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

	varPostV1EavAttributesetsRequest := _PostV1EavAttributesetsRequest{}

	err = json.Unmarshal(data, &varPostV1EavAttributesetsRequest)

	if err != nil {
		return err
	}

	*o = PostV1EavAttributesetsRequest(varPostV1EavAttributesetsRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "entityTypeCode")
		delete(additionalProperties, "attributeSet")
		delete(additionalProperties, "skeletonId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PostV1EavAttributesetsRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *PostV1EavAttributesetsRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullablePostV1EavAttributesetsRequest struct {
	value *PostV1EavAttributesetsRequest
	isSet bool
}

func (v NullablePostV1EavAttributesetsRequest) Get() *PostV1EavAttributesetsRequest {
	return v.value
}

func (v *NullablePostV1EavAttributesetsRequest) Set(val *PostV1EavAttributesetsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostV1EavAttributesetsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostV1EavAttributesetsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostV1EavAttributesetsRequest(val *PostV1EavAttributesetsRequest) *NullablePostV1EavAttributesetsRequest {
	return &NullablePostV1EavAttributesetsRequest{value: val, isSet: true}
}

func (v NullablePostV1EavAttributesetsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostV1EavAttributesetsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
