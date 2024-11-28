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

// checks if the PutV1CustomersMeActivateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutV1CustomersMeActivateRequest{}

// PutV1CustomersMeActivateRequest struct for PutV1CustomersMeActivateRequest
type PutV1CustomersMeActivateRequest struct {
	ConfirmationKey      string `json:"confirmationKey"`
	AdditionalProperties map[string]interface{}
}

type _PutV1CustomersMeActivateRequest PutV1CustomersMeActivateRequest

// NewPutV1CustomersMeActivateRequest instantiates a new PutV1CustomersMeActivateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutV1CustomersMeActivateRequest(confirmationKey string) *PutV1CustomersMeActivateRequest {
	this := PutV1CustomersMeActivateRequest{}
	this.ConfirmationKey = confirmationKey
	return &this
}

// NewPutV1CustomersMeActivateRequestWithDefaults instantiates a new PutV1CustomersMeActivateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutV1CustomersMeActivateRequestWithDefaults() *PutV1CustomersMeActivateRequest {
	this := PutV1CustomersMeActivateRequest{}
	return &this
}

// GetConfirmationKey returns the ConfirmationKey field value
func (o *PutV1CustomersMeActivateRequest) GetConfirmationKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfirmationKey
}

// GetConfirmationKeyOk returns a tuple with the ConfirmationKey field value
// and a boolean to check if the value has been set.
func (o *PutV1CustomersMeActivateRequest) GetConfirmationKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfirmationKey, true
}

// SetConfirmationKey sets field value
func (o *PutV1CustomersMeActivateRequest) SetConfirmationKey(v string) {
	o.ConfirmationKey = v
}

func (o PutV1CustomersMeActivateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutV1CustomersMeActivateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["confirmationKey"] = o.ConfirmationKey

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PutV1CustomersMeActivateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"confirmationKey",
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

	varPutV1CustomersMeActivateRequest := _PutV1CustomersMeActivateRequest{}

	err = json.Unmarshal(data, &varPutV1CustomersMeActivateRequest)

	if err != nil {
		return err
	}

	*o = PutV1CustomersMeActivateRequest(varPutV1CustomersMeActivateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "confirmationKey")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PutV1CustomersMeActivateRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *PutV1CustomersMeActivateRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullablePutV1CustomersMeActivateRequest struct {
	value *PutV1CustomersMeActivateRequest
	isSet bool
}

func (v NullablePutV1CustomersMeActivateRequest) Get() *PutV1CustomersMeActivateRequest {
	return v.value
}

func (v *NullablePutV1CustomersMeActivateRequest) Set(val *PutV1CustomersMeActivateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutV1CustomersMeActivateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutV1CustomersMeActivateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutV1CustomersMeActivateRequest(val *PutV1CustomersMeActivateRequest) *NullablePutV1CustomersMeActivateRequest {
	return &NullablePutV1CustomersMeActivateRequest{value: val, isSet: true}
}

func (v NullablePutV1CustomersMeActivateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutV1CustomersMeActivateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}