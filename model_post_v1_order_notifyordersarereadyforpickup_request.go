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

// checks if the PostV1OrderNotifyordersarereadyforpickupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostV1OrderNotifyordersarereadyforpickupRequest{}

// PostV1OrderNotifyordersarereadyforpickupRequest struct for PostV1OrderNotifyordersarereadyforpickupRequest
type PostV1OrderNotifyordersarereadyforpickupRequest struct {
	OrderIds             []int32 `json:"orderIds"`
	AdditionalProperties map[string]interface{}
}

type _PostV1OrderNotifyordersarereadyforpickupRequest PostV1OrderNotifyordersarereadyforpickupRequest

// NewPostV1OrderNotifyordersarereadyforpickupRequest instantiates a new PostV1OrderNotifyordersarereadyforpickupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostV1OrderNotifyordersarereadyforpickupRequest(orderIds []int32) *PostV1OrderNotifyordersarereadyforpickupRequest {
	this := PostV1OrderNotifyordersarereadyforpickupRequest{}
	this.OrderIds = orderIds
	return &this
}

// NewPostV1OrderNotifyordersarereadyforpickupRequestWithDefaults instantiates a new PostV1OrderNotifyordersarereadyforpickupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostV1OrderNotifyordersarereadyforpickupRequestWithDefaults() *PostV1OrderNotifyordersarereadyforpickupRequest {
	this := PostV1OrderNotifyordersarereadyforpickupRequest{}
	return &this
}

// GetOrderIds returns the OrderIds field value
func (o *PostV1OrderNotifyordersarereadyforpickupRequest) GetOrderIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.OrderIds
}

// GetOrderIdsOk returns a tuple with the OrderIds field value
// and a boolean to check if the value has been set.
func (o *PostV1OrderNotifyordersarereadyforpickupRequest) GetOrderIdsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrderIds, true
}

// SetOrderIds sets field value
func (o *PostV1OrderNotifyordersarereadyforpickupRequest) SetOrderIds(v []int32) {
	o.OrderIds = v
}

func (o PostV1OrderNotifyordersarereadyforpickupRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostV1OrderNotifyordersarereadyforpickupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["orderIds"] = o.OrderIds

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostV1OrderNotifyordersarereadyforpickupRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"orderIds",
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

	varPostV1OrderNotifyordersarereadyforpickupRequest := _PostV1OrderNotifyordersarereadyforpickupRequest{}

	err = json.Unmarshal(data, &varPostV1OrderNotifyordersarereadyforpickupRequest)

	if err != nil {
		return err
	}

	*o = PostV1OrderNotifyordersarereadyforpickupRequest(varPostV1OrderNotifyordersarereadyforpickupRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orderIds")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PostV1OrderNotifyordersarereadyforpickupRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *PostV1OrderNotifyordersarereadyforpickupRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullablePostV1OrderNotifyordersarereadyforpickupRequest struct {
	value *PostV1OrderNotifyordersarereadyforpickupRequest
	isSet bool
}

func (v NullablePostV1OrderNotifyordersarereadyforpickupRequest) Get() *PostV1OrderNotifyordersarereadyforpickupRequest {
	return v.value
}

func (v *NullablePostV1OrderNotifyordersarereadyforpickupRequest) Set(val *PostV1OrderNotifyordersarereadyforpickupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostV1OrderNotifyordersarereadyforpickupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostV1OrderNotifyordersarereadyforpickupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostV1OrderNotifyordersarereadyforpickupRequest(val *PostV1OrderNotifyordersarereadyforpickupRequest) *NullablePostV1OrderNotifyordersarereadyforpickupRequest {
	return &NullablePostV1OrderNotifyordersarereadyforpickupRequest{value: val, isSet: true}
}

func (v NullablePostV1OrderNotifyordersarereadyforpickupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostV1OrderNotifyordersarereadyforpickupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
