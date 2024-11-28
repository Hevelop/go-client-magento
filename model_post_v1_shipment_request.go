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

// checks if the PostV1ShipmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostV1ShipmentRequest{}

// PostV1ShipmentRequest struct for PostV1ShipmentRequest
type PostV1ShipmentRequest struct {
	Entity               SalesDataShipmentInterface `json:"entity"`
	AdditionalProperties map[string]interface{}
}

type _PostV1ShipmentRequest PostV1ShipmentRequest

// NewPostV1ShipmentRequest instantiates a new PostV1ShipmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostV1ShipmentRequest(entity SalesDataShipmentInterface) *PostV1ShipmentRequest {
	this := PostV1ShipmentRequest{}
	this.Entity = entity
	return &this
}

// NewPostV1ShipmentRequestWithDefaults instantiates a new PostV1ShipmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostV1ShipmentRequestWithDefaults() *PostV1ShipmentRequest {
	this := PostV1ShipmentRequest{}
	return &this
}

// GetEntity returns the Entity field value
func (o *PostV1ShipmentRequest) GetEntity() SalesDataShipmentInterface {
	if o == nil {
		var ret SalesDataShipmentInterface
		return ret
	}

	return o.Entity
}

// GetEntityOk returns a tuple with the Entity field value
// and a boolean to check if the value has been set.
func (o *PostV1ShipmentRequest) GetEntityOk() (*SalesDataShipmentInterface, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Entity, true
}

// SetEntity sets field value
func (o *PostV1ShipmentRequest) SetEntity(v SalesDataShipmentInterface) {
	o.Entity = v
}

func (o PostV1ShipmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostV1ShipmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entity"] = o.Entity

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostV1ShipmentRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"entity",
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

	varPostV1ShipmentRequest := _PostV1ShipmentRequest{}

	err = json.Unmarshal(data, &varPostV1ShipmentRequest)

	if err != nil {
		return err
	}

	*o = PostV1ShipmentRequest(varPostV1ShipmentRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "entity")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PostV1ShipmentRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *PostV1ShipmentRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullablePostV1ShipmentRequest struct {
	value *PostV1ShipmentRequest
	isSet bool
}

func (v NullablePostV1ShipmentRequest) Get() *PostV1ShipmentRequest {
	return v.value
}

func (v *NullablePostV1ShipmentRequest) Set(val *PostV1ShipmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostV1ShipmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostV1ShipmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostV1ShipmentRequest(val *PostV1ShipmentRequest) *NullablePostV1ShipmentRequest {
	return &NullablePostV1ShipmentRequest{value: val, isSet: true}
}

func (v NullablePostV1ShipmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostV1ShipmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}