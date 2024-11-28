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

// checks if the PostV1ImportJsonRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostV1ImportJsonRequest{}

// PostV1ImportJsonRequest struct for PostV1ImportJsonRequest
type PostV1ImportJsonRequest struct {
	Source               ImportJsonApiDataSourceDataInterface `json:"source"`
	AdditionalProperties map[string]interface{}
}

type _PostV1ImportJsonRequest PostV1ImportJsonRequest

// NewPostV1ImportJsonRequest instantiates a new PostV1ImportJsonRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostV1ImportJsonRequest(source ImportJsonApiDataSourceDataInterface) *PostV1ImportJsonRequest {
	this := PostV1ImportJsonRequest{}
	this.Source = source
	return &this
}

// NewPostV1ImportJsonRequestWithDefaults instantiates a new PostV1ImportJsonRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostV1ImportJsonRequestWithDefaults() *PostV1ImportJsonRequest {
	this := PostV1ImportJsonRequest{}
	return &this
}

// GetSource returns the Source field value
func (o *PostV1ImportJsonRequest) GetSource() ImportJsonApiDataSourceDataInterface {
	if o == nil {
		var ret ImportJsonApiDataSourceDataInterface
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *PostV1ImportJsonRequest) GetSourceOk() (*ImportJsonApiDataSourceDataInterface, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *PostV1ImportJsonRequest) SetSource(v ImportJsonApiDataSourceDataInterface) {
	o.Source = v
}

func (o PostV1ImportJsonRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostV1ImportJsonRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source"] = o.Source

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostV1ImportJsonRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"source",
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

	varPostV1ImportJsonRequest := _PostV1ImportJsonRequest{}

	err = json.Unmarshal(data, &varPostV1ImportJsonRequest)

	if err != nil {
		return err
	}

	*o = PostV1ImportJsonRequest(varPostV1ImportJsonRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "source")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PostV1ImportJsonRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *PostV1ImportJsonRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullablePostV1ImportJsonRequest struct {
	value *PostV1ImportJsonRequest
	isSet bool
}

func (v NullablePostV1ImportJsonRequest) Get() *PostV1ImportJsonRequest {
	return v.value
}

func (v *NullablePostV1ImportJsonRequest) Set(val *PostV1ImportJsonRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostV1ImportJsonRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostV1ImportJsonRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostV1ImportJsonRequest(val *PostV1ImportJsonRequest) *NullablePostV1ImportJsonRequest {
	return &NullablePostV1ImportJsonRequest{value: val, isSet: true}
}

func (v NullablePostV1ImportJsonRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostV1ImportJsonRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
