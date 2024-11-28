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

// checks if the PostV1CartsMineShippinginformationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostV1CartsMineShippinginformationRequest{}

// PostV1CartsMineShippinginformationRequest struct for PostV1CartsMineShippinginformationRequest
type PostV1CartsMineShippinginformationRequest struct {
	AddressInformation   CheckoutDataShippingInformationInterface `json:"addressInformation"`
	AdditionalProperties map[string]interface{}
}

type _PostV1CartsMineShippinginformationRequest PostV1CartsMineShippinginformationRequest

// NewPostV1CartsMineShippinginformationRequest instantiates a new PostV1CartsMineShippinginformationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostV1CartsMineShippinginformationRequest(addressInformation CheckoutDataShippingInformationInterface) *PostV1CartsMineShippinginformationRequest {
	this := PostV1CartsMineShippinginformationRequest{}
	this.AddressInformation = addressInformation
	return &this
}

// NewPostV1CartsMineShippinginformationRequestWithDefaults instantiates a new PostV1CartsMineShippinginformationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostV1CartsMineShippinginformationRequestWithDefaults() *PostV1CartsMineShippinginformationRequest {
	this := PostV1CartsMineShippinginformationRequest{}
	return &this
}

// GetAddressInformation returns the AddressInformation field value
func (o *PostV1CartsMineShippinginformationRequest) GetAddressInformation() CheckoutDataShippingInformationInterface {
	if o == nil {
		var ret CheckoutDataShippingInformationInterface
		return ret
	}

	return o.AddressInformation
}

// GetAddressInformationOk returns a tuple with the AddressInformation field value
// and a boolean to check if the value has been set.
func (o *PostV1CartsMineShippinginformationRequest) GetAddressInformationOk() (*CheckoutDataShippingInformationInterface, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddressInformation, true
}

// SetAddressInformation sets field value
func (o *PostV1CartsMineShippinginformationRequest) SetAddressInformation(v CheckoutDataShippingInformationInterface) {
	o.AddressInformation = v
}

func (o PostV1CartsMineShippinginformationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostV1CartsMineShippinginformationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["addressInformation"] = o.AddressInformation

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostV1CartsMineShippinginformationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"addressInformation",
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

	varPostV1CartsMineShippinginformationRequest := _PostV1CartsMineShippinginformationRequest{}

	err = json.Unmarshal(data, &varPostV1CartsMineShippinginformationRequest)

	if err != nil {
		return err
	}

	*o = PostV1CartsMineShippinginformationRequest(varPostV1CartsMineShippinginformationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "addressInformation")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PostV1CartsMineShippinginformationRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *PostV1CartsMineShippinginformationRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullablePostV1CartsMineShippinginformationRequest struct {
	value *PostV1CartsMineShippinginformationRequest
	isSet bool
}

func (v NullablePostV1CartsMineShippinginformationRequest) Get() *PostV1CartsMineShippinginformationRequest {
	return v.value
}

func (v *NullablePostV1CartsMineShippinginformationRequest) Set(val *PostV1CartsMineShippinginformationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostV1CartsMineShippinginformationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostV1CartsMineShippinginformationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostV1CartsMineShippinginformationRequest(val *PostV1CartsMineShippinginformationRequest) *NullablePostV1CartsMineShippinginformationRequest {
	return &NullablePostV1CartsMineShippinginformationRequest{value: val, isSet: true}
}

func (v NullablePostV1CartsMineShippinginformationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostV1CartsMineShippinginformationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
