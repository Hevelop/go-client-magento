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

// checks if the PutV1HierarchyMoveIdRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutV1HierarchyMoveIdRequest{}

// PutV1HierarchyMoveIdRequest struct for PutV1HierarchyMoveIdRequest
type PutV1HierarchyMoveIdRequest struct {
	NewParentId          int32 `json:"newParentId"`
	AdditionalProperties map[string]interface{}
}

type _PutV1HierarchyMoveIdRequest PutV1HierarchyMoveIdRequest

// NewPutV1HierarchyMoveIdRequest instantiates a new PutV1HierarchyMoveIdRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutV1HierarchyMoveIdRequest(newParentId int32) *PutV1HierarchyMoveIdRequest {
	this := PutV1HierarchyMoveIdRequest{}
	this.NewParentId = newParentId
	return &this
}

// NewPutV1HierarchyMoveIdRequestWithDefaults instantiates a new PutV1HierarchyMoveIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutV1HierarchyMoveIdRequestWithDefaults() *PutV1HierarchyMoveIdRequest {
	this := PutV1HierarchyMoveIdRequest{}
	return &this
}

// GetNewParentId returns the NewParentId field value
func (o *PutV1HierarchyMoveIdRequest) GetNewParentId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NewParentId
}

// GetNewParentIdOk returns a tuple with the NewParentId field value
// and a boolean to check if the value has been set.
func (o *PutV1HierarchyMoveIdRequest) GetNewParentIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewParentId, true
}

// SetNewParentId sets field value
func (o *PutV1HierarchyMoveIdRequest) SetNewParentId(v int32) {
	o.NewParentId = v
}

func (o PutV1HierarchyMoveIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutV1HierarchyMoveIdRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["newParentId"] = o.NewParentId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PutV1HierarchyMoveIdRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"newParentId",
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

	varPutV1HierarchyMoveIdRequest := _PutV1HierarchyMoveIdRequest{}

	err = json.Unmarshal(data, &varPutV1HierarchyMoveIdRequest)

	if err != nil {
		return err
	}

	*o = PutV1HierarchyMoveIdRequest(varPutV1HierarchyMoveIdRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "newParentId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PutV1HierarchyMoveIdRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *PutV1HierarchyMoveIdRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullablePutV1HierarchyMoveIdRequest struct {
	value *PutV1HierarchyMoveIdRequest
	isSet bool
}

func (v NullablePutV1HierarchyMoveIdRequest) Get() *PutV1HierarchyMoveIdRequest {
	return v.value
}

func (v *NullablePutV1HierarchyMoveIdRequest) Set(val *PutV1HierarchyMoveIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutV1HierarchyMoveIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutV1HierarchyMoveIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutV1HierarchyMoveIdRequest(val *PutV1HierarchyMoveIdRequest) *NullablePutV1HierarchyMoveIdRequest {
	return &NullablePutV1HierarchyMoveIdRequest{value: val, isSet: true}
}

func (v NullablePutV1HierarchyMoveIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutV1HierarchyMoveIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
