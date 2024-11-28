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

// checks if the TaxDataGrandTotalRatesInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxDataGrandTotalRatesInterface{}

// TaxDataGrandTotalRatesInterface Interface GrandTotalRatesInterface
type TaxDataGrandTotalRatesInterface struct {
	// Tax percentage value
	Percent string `json:"percent"`
	// Rate title
	Title                string `json:"title"`
	AdditionalProperties map[string]interface{}
}

type _TaxDataGrandTotalRatesInterface TaxDataGrandTotalRatesInterface

// NewTaxDataGrandTotalRatesInterface instantiates a new TaxDataGrandTotalRatesInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxDataGrandTotalRatesInterface(percent string, title string) *TaxDataGrandTotalRatesInterface {
	this := TaxDataGrandTotalRatesInterface{}
	this.Percent = percent
	this.Title = title
	return &this
}

// NewTaxDataGrandTotalRatesInterfaceWithDefaults instantiates a new TaxDataGrandTotalRatesInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxDataGrandTotalRatesInterfaceWithDefaults() *TaxDataGrandTotalRatesInterface {
	this := TaxDataGrandTotalRatesInterface{}
	return &this
}

// GetPercent returns the Percent field value
func (o *TaxDataGrandTotalRatesInterface) GetPercent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Percent
}

// GetPercentOk returns a tuple with the Percent field value
// and a boolean to check if the value has been set.
func (o *TaxDataGrandTotalRatesInterface) GetPercentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Percent, true
}

// SetPercent sets field value
func (o *TaxDataGrandTotalRatesInterface) SetPercent(v string) {
	o.Percent = v
}

// GetTitle returns the Title field value
func (o *TaxDataGrandTotalRatesInterface) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *TaxDataGrandTotalRatesInterface) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *TaxDataGrandTotalRatesInterface) SetTitle(v string) {
	o.Title = v
}

func (o TaxDataGrandTotalRatesInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaxDataGrandTotalRatesInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["percent"] = o.Percent
	toSerialize["title"] = o.Title

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TaxDataGrandTotalRatesInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"percent",
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

	varTaxDataGrandTotalRatesInterface := _TaxDataGrandTotalRatesInterface{}

	err = json.Unmarshal(data, &varTaxDataGrandTotalRatesInterface)

	if err != nil {
		return err
	}

	*o = TaxDataGrandTotalRatesInterface(varTaxDataGrandTotalRatesInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "percent")
		delete(additionalProperties, "title")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *TaxDataGrandTotalRatesInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *TaxDataGrandTotalRatesInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableTaxDataGrandTotalRatesInterface struct {
	value *TaxDataGrandTotalRatesInterface
	isSet bool
}

func (v NullableTaxDataGrandTotalRatesInterface) Get() *TaxDataGrandTotalRatesInterface {
	return v.value
}

func (v *NullableTaxDataGrandTotalRatesInterface) Set(val *TaxDataGrandTotalRatesInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxDataGrandTotalRatesInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxDataGrandTotalRatesInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxDataGrandTotalRatesInterface(val *TaxDataGrandTotalRatesInterface) *NullableTaxDataGrandTotalRatesInterface {
	return &NullableTaxDataGrandTotalRatesInterface{value: val, isSet: true}
}

func (v NullableTaxDataGrandTotalRatesInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxDataGrandTotalRatesInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}