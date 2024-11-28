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

// checks if the PaymentServicesPaypalDataPaymentSourceDetailsInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentServicesPaypalDataPaymentSourceDetailsInterface{}

// PaymentServicesPaypalDataPaymentSourceDetailsInterface
type PaymentServicesPaypalDataPaymentSourceDetailsInterface struct {
	Card                 PaymentServicesPaypalDataPaymentCardDetailsInterface `json:"card"`
	AdditionalProperties map[string]interface{}
}

type _PaymentServicesPaypalDataPaymentSourceDetailsInterface PaymentServicesPaypalDataPaymentSourceDetailsInterface

// NewPaymentServicesPaypalDataPaymentSourceDetailsInterface instantiates a new PaymentServicesPaypalDataPaymentSourceDetailsInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentServicesPaypalDataPaymentSourceDetailsInterface(card PaymentServicesPaypalDataPaymentCardDetailsInterface) *PaymentServicesPaypalDataPaymentSourceDetailsInterface {
	this := PaymentServicesPaypalDataPaymentSourceDetailsInterface{}
	this.Card = card
	return &this
}

// NewPaymentServicesPaypalDataPaymentSourceDetailsInterfaceWithDefaults instantiates a new PaymentServicesPaypalDataPaymentSourceDetailsInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentServicesPaypalDataPaymentSourceDetailsInterfaceWithDefaults() *PaymentServicesPaypalDataPaymentSourceDetailsInterface {
	this := PaymentServicesPaypalDataPaymentSourceDetailsInterface{}
	return &this
}

// GetCard returns the Card field value
func (o *PaymentServicesPaypalDataPaymentSourceDetailsInterface) GetCard() PaymentServicesPaypalDataPaymentCardDetailsInterface {
	if o == nil {
		var ret PaymentServicesPaypalDataPaymentCardDetailsInterface
		return ret
	}

	return o.Card
}

// GetCardOk returns a tuple with the Card field value
// and a boolean to check if the value has been set.
func (o *PaymentServicesPaypalDataPaymentSourceDetailsInterface) GetCardOk() (*PaymentServicesPaypalDataPaymentCardDetailsInterface, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Card, true
}

// SetCard sets field value
func (o *PaymentServicesPaypalDataPaymentSourceDetailsInterface) SetCard(v PaymentServicesPaypalDataPaymentCardDetailsInterface) {
	o.Card = v
}

func (o PaymentServicesPaypalDataPaymentSourceDetailsInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentServicesPaypalDataPaymentSourceDetailsInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["card"] = o.Card

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaymentServicesPaypalDataPaymentSourceDetailsInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"card",
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

	varPaymentServicesPaypalDataPaymentSourceDetailsInterface := _PaymentServicesPaypalDataPaymentSourceDetailsInterface{}

	err = json.Unmarshal(data, &varPaymentServicesPaypalDataPaymentSourceDetailsInterface)

	if err != nil {
		return err
	}

	*o = PaymentServicesPaypalDataPaymentSourceDetailsInterface(varPaymentServicesPaypalDataPaymentSourceDetailsInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "card")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PaymentServicesPaypalDataPaymentSourceDetailsInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *PaymentServicesPaypalDataPaymentSourceDetailsInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullablePaymentServicesPaypalDataPaymentSourceDetailsInterface struct {
	value *PaymentServicesPaypalDataPaymentSourceDetailsInterface
	isSet bool
}

func (v NullablePaymentServicesPaypalDataPaymentSourceDetailsInterface) Get() *PaymentServicesPaypalDataPaymentSourceDetailsInterface {
	return v.value
}

func (v *NullablePaymentServicesPaypalDataPaymentSourceDetailsInterface) Set(val *PaymentServicesPaypalDataPaymentSourceDetailsInterface) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentServicesPaypalDataPaymentSourceDetailsInterface) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentServicesPaypalDataPaymentSourceDetailsInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentServicesPaypalDataPaymentSourceDetailsInterface(val *PaymentServicesPaypalDataPaymentSourceDetailsInterface) *NullablePaymentServicesPaypalDataPaymentSourceDetailsInterface {
	return &NullablePaymentServicesPaypalDataPaymentSourceDetailsInterface{value: val, isSet: true}
}

func (v NullablePaymentServicesPaypalDataPaymentSourceDetailsInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentServicesPaypalDataPaymentSourceDetailsInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}