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

// checks if the PaymentServicesPaypalDataPaymentCardBinDetailsInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentServicesPaypalDataPaymentCardBinDetailsInterface{}

// PaymentServicesPaypalDataPaymentCardBinDetailsInterface
type PaymentServicesPaypalDataPaymentCardBinDetailsInterface struct {
	// Bin
	Bin                  string `json:"bin"`
	AdditionalProperties map[string]interface{}
}

type _PaymentServicesPaypalDataPaymentCardBinDetailsInterface PaymentServicesPaypalDataPaymentCardBinDetailsInterface

// NewPaymentServicesPaypalDataPaymentCardBinDetailsInterface instantiates a new PaymentServicesPaypalDataPaymentCardBinDetailsInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentServicesPaypalDataPaymentCardBinDetailsInterface(bin string) *PaymentServicesPaypalDataPaymentCardBinDetailsInterface {
	this := PaymentServicesPaypalDataPaymentCardBinDetailsInterface{}
	this.Bin = bin
	return &this
}

// NewPaymentServicesPaypalDataPaymentCardBinDetailsInterfaceWithDefaults instantiates a new PaymentServicesPaypalDataPaymentCardBinDetailsInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentServicesPaypalDataPaymentCardBinDetailsInterfaceWithDefaults() *PaymentServicesPaypalDataPaymentCardBinDetailsInterface {
	this := PaymentServicesPaypalDataPaymentCardBinDetailsInterface{}
	return &this
}

// GetBin returns the Bin field value
func (o *PaymentServicesPaypalDataPaymentCardBinDetailsInterface) GetBin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bin
}

// GetBinOk returns a tuple with the Bin field value
// and a boolean to check if the value has been set.
func (o *PaymentServicesPaypalDataPaymentCardBinDetailsInterface) GetBinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bin, true
}

// SetBin sets field value
func (o *PaymentServicesPaypalDataPaymentCardBinDetailsInterface) SetBin(v string) {
	o.Bin = v
}

func (o PaymentServicesPaypalDataPaymentCardBinDetailsInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentServicesPaypalDataPaymentCardBinDetailsInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bin"] = o.Bin

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaymentServicesPaypalDataPaymentCardBinDetailsInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bin",
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

	varPaymentServicesPaypalDataPaymentCardBinDetailsInterface := _PaymentServicesPaypalDataPaymentCardBinDetailsInterface{}

	err = json.Unmarshal(data, &varPaymentServicesPaypalDataPaymentCardBinDetailsInterface)

	if err != nil {
		return err
	}

	*o = PaymentServicesPaypalDataPaymentCardBinDetailsInterface(varPaymentServicesPaypalDataPaymentCardBinDetailsInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "bin")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PaymentServicesPaypalDataPaymentCardBinDetailsInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *PaymentServicesPaypalDataPaymentCardBinDetailsInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullablePaymentServicesPaypalDataPaymentCardBinDetailsInterface struct {
	value *PaymentServicesPaypalDataPaymentCardBinDetailsInterface
	isSet bool
}

func (v NullablePaymentServicesPaypalDataPaymentCardBinDetailsInterface) Get() *PaymentServicesPaypalDataPaymentCardBinDetailsInterface {
	return v.value
}

func (v *NullablePaymentServicesPaypalDataPaymentCardBinDetailsInterface) Set(val *PaymentServicesPaypalDataPaymentCardBinDetailsInterface) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentServicesPaypalDataPaymentCardBinDetailsInterface) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentServicesPaypalDataPaymentCardBinDetailsInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentServicesPaypalDataPaymentCardBinDetailsInterface(val *PaymentServicesPaypalDataPaymentCardBinDetailsInterface) *NullablePaymentServicesPaypalDataPaymentCardBinDetailsInterface {
	return &NullablePaymentServicesPaypalDataPaymentCardBinDetailsInterface{value: val, isSet: true}
}

func (v NullablePaymentServicesPaypalDataPaymentCardBinDetailsInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentServicesPaypalDataPaymentCardBinDetailsInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
