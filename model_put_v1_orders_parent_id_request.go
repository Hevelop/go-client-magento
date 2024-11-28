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

// checks if the PutV1OrdersParentIdRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutV1OrdersParentIdRequest{}

// PutV1OrdersParentIdRequest struct for PutV1OrdersParentIdRequest
type PutV1OrdersParentIdRequest struct {
	Entity               SalesDataOrderAddressInterface `json:"entity"`
	AdditionalProperties map[string]interface{}
}

type _PutV1OrdersParentIdRequest PutV1OrdersParentIdRequest

// NewPutV1OrdersParentIdRequest instantiates a new PutV1OrdersParentIdRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutV1OrdersParentIdRequest(entity SalesDataOrderAddressInterface) *PutV1OrdersParentIdRequest {
	this := PutV1OrdersParentIdRequest{}
	this.Entity = entity
	return &this
}

// NewPutV1OrdersParentIdRequestWithDefaults instantiates a new PutV1OrdersParentIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutV1OrdersParentIdRequestWithDefaults() *PutV1OrdersParentIdRequest {
	this := PutV1OrdersParentIdRequest{}
	return &this
}

// GetEntity returns the Entity field value
func (o *PutV1OrdersParentIdRequest) GetEntity() SalesDataOrderAddressInterface {
	if o == nil {
		var ret SalesDataOrderAddressInterface
		return ret
	}

	return o.Entity
}

// GetEntityOk returns a tuple with the Entity field value
// and a boolean to check if the value has been set.
func (o *PutV1OrdersParentIdRequest) GetEntityOk() (*SalesDataOrderAddressInterface, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Entity, true
}

// SetEntity sets field value
func (o *PutV1OrdersParentIdRequest) SetEntity(v SalesDataOrderAddressInterface) {
	o.Entity = v
}

func (o PutV1OrdersParentIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutV1OrdersParentIdRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entity"] = o.Entity

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PutV1OrdersParentIdRequest) UnmarshalJSON(data []byte) (err error) {
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

	varPutV1OrdersParentIdRequest := _PutV1OrdersParentIdRequest{}

	err = json.Unmarshal(data, &varPutV1OrdersParentIdRequest)

	if err != nil {
		return err
	}

	*o = PutV1OrdersParentIdRequest(varPutV1OrdersParentIdRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "entity")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PutV1OrdersParentIdRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *PutV1OrdersParentIdRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullablePutV1OrdersParentIdRequest struct {
	value *PutV1OrdersParentIdRequest
	isSet bool
}

func (v NullablePutV1OrdersParentIdRequest) Get() *PutV1OrdersParentIdRequest {
	return v.value
}

func (v *NullablePutV1OrdersParentIdRequest) Set(val *PutV1OrdersParentIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutV1OrdersParentIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutV1OrdersParentIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutV1OrdersParentIdRequest(val *PutV1OrdersParentIdRequest) *NullablePutV1OrdersParentIdRequest {
	return &NullablePutV1OrdersParentIdRequest{value: val, isSet: true}
}

func (v NullablePutV1OrdersParentIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutV1OrdersParentIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}