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

// checks if the QuoteDataPaymentMethodInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuoteDataPaymentMethodInterface{}

// QuoteDataPaymentMethodInterface Interface PaymentMethodInterface
type QuoteDataPaymentMethodInterface struct {
	// Payment method code
	Code string `json:"code"`
	// Payment method title
	Title                string `json:"title"`
	AdditionalProperties map[string]interface{}
}

type _QuoteDataPaymentMethodInterface QuoteDataPaymentMethodInterface

// NewQuoteDataPaymentMethodInterface instantiates a new QuoteDataPaymentMethodInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteDataPaymentMethodInterface(code string, title string) *QuoteDataPaymentMethodInterface {
	this := QuoteDataPaymentMethodInterface{}
	this.Code = code
	this.Title = title
	return &this
}

// NewQuoteDataPaymentMethodInterfaceWithDefaults instantiates a new QuoteDataPaymentMethodInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteDataPaymentMethodInterfaceWithDefaults() *QuoteDataPaymentMethodInterface {
	this := QuoteDataPaymentMethodInterface{}
	return &this
}

// GetCode returns the Code field value
func (o *QuoteDataPaymentMethodInterface) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *QuoteDataPaymentMethodInterface) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *QuoteDataPaymentMethodInterface) SetCode(v string) {
	o.Code = v
}

// GetTitle returns the Title field value
func (o *QuoteDataPaymentMethodInterface) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *QuoteDataPaymentMethodInterface) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *QuoteDataPaymentMethodInterface) SetTitle(v string) {
	o.Title = v
}

func (o QuoteDataPaymentMethodInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuoteDataPaymentMethodInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["title"] = o.Title

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuoteDataPaymentMethodInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"title",
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

	varQuoteDataPaymentMethodInterface := _QuoteDataPaymentMethodInterface{}

	err = json.Unmarshal(data, &varQuoteDataPaymentMethodInterface)

	if err != nil {
		return err
	}

	*o = QuoteDataPaymentMethodInterface(varQuoteDataPaymentMethodInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "title")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *QuoteDataPaymentMethodInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *QuoteDataPaymentMethodInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableQuoteDataPaymentMethodInterface struct {
	value *QuoteDataPaymentMethodInterface
	isSet bool
}

func (v NullableQuoteDataPaymentMethodInterface) Get() *QuoteDataPaymentMethodInterface {
	return v.value
}

func (v *NullableQuoteDataPaymentMethodInterface) Set(val *QuoteDataPaymentMethodInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteDataPaymentMethodInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteDataPaymentMethodInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteDataPaymentMethodInterface(val *QuoteDataPaymentMethodInterface) *NullableQuoteDataPaymentMethodInterface {
	return &NullableQuoteDataPaymentMethodInterface{value: val, isSet: true}
}

func (v NullableQuoteDataPaymentMethodInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteDataPaymentMethodInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}