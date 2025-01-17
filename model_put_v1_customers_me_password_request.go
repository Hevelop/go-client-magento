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

// checks if the PutV1CustomersMePasswordRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutV1CustomersMePasswordRequest{}

// PutV1CustomersMePasswordRequest struct for PutV1CustomersMePasswordRequest
type PutV1CustomersMePasswordRequest struct {
	CurrentPassword      string `json:"currentPassword"`
	NewPassword          string `json:"newPassword"`
	AdditionalProperties map[string]interface{}
}

type _PutV1CustomersMePasswordRequest PutV1CustomersMePasswordRequest

// NewPutV1CustomersMePasswordRequest instantiates a new PutV1CustomersMePasswordRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutV1CustomersMePasswordRequest(currentPassword string, newPassword string) *PutV1CustomersMePasswordRequest {
	this := PutV1CustomersMePasswordRequest{}
	this.CurrentPassword = currentPassword
	this.NewPassword = newPassword
	return &this
}

// NewPutV1CustomersMePasswordRequestWithDefaults instantiates a new PutV1CustomersMePasswordRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutV1CustomersMePasswordRequestWithDefaults() *PutV1CustomersMePasswordRequest {
	this := PutV1CustomersMePasswordRequest{}
	return &this
}

// GetCurrentPassword returns the CurrentPassword field value
func (o *PutV1CustomersMePasswordRequest) GetCurrentPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrentPassword
}

// GetCurrentPasswordOk returns a tuple with the CurrentPassword field value
// and a boolean to check if the value has been set.
func (o *PutV1CustomersMePasswordRequest) GetCurrentPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentPassword, true
}

// SetCurrentPassword sets field value
func (o *PutV1CustomersMePasswordRequest) SetCurrentPassword(v string) {
	o.CurrentPassword = v
}

// GetNewPassword returns the NewPassword field value
func (o *PutV1CustomersMePasswordRequest) GetNewPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NewPassword
}

// GetNewPasswordOk returns a tuple with the NewPassword field value
// and a boolean to check if the value has been set.
func (o *PutV1CustomersMePasswordRequest) GetNewPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewPassword, true
}

// SetNewPassword sets field value
func (o *PutV1CustomersMePasswordRequest) SetNewPassword(v string) {
	o.NewPassword = v
}

func (o PutV1CustomersMePasswordRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutV1CustomersMePasswordRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["currentPassword"] = o.CurrentPassword
	toSerialize["newPassword"] = o.NewPassword

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PutV1CustomersMePasswordRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"currentPassword",
		"newPassword",
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

	varPutV1CustomersMePasswordRequest := _PutV1CustomersMePasswordRequest{}

	err = json.Unmarshal(data, &varPutV1CustomersMePasswordRequest)

	if err != nil {
		return err
	}

	*o = PutV1CustomersMePasswordRequest(varPutV1CustomersMePasswordRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "currentPassword")
		delete(additionalProperties, "newPassword")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PutV1CustomersMePasswordRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *PutV1CustomersMePasswordRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullablePutV1CustomersMePasswordRequest struct {
	value *PutV1CustomersMePasswordRequest
	isSet bool
}

func (v NullablePutV1CustomersMePasswordRequest) Get() *PutV1CustomersMePasswordRequest {
	return v.value
}

func (v *NullablePutV1CustomersMePasswordRequest) Set(val *PutV1CustomersMePasswordRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutV1CustomersMePasswordRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutV1CustomersMePasswordRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutV1CustomersMePasswordRequest(val *PutV1CustomersMePasswordRequest) *NullablePutV1CustomersMePasswordRequest {
	return &NullablePutV1CustomersMePasswordRequest{value: val, isSet: true}
}

func (v NullablePutV1CustomersMePasswordRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutV1CustomersMePasswordRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
