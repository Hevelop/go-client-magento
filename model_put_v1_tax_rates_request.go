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

// checks if the PutV1TaxRatesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutV1TaxRatesRequest{}

// PutV1TaxRatesRequest struct for PutV1TaxRatesRequest
type PutV1TaxRatesRequest struct {
	TaxRate              TaxDataTaxRateInterface `json:"taxRate"`
	AdditionalProperties map[string]interface{}
}

type _PutV1TaxRatesRequest PutV1TaxRatesRequest

// NewPutV1TaxRatesRequest instantiates a new PutV1TaxRatesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutV1TaxRatesRequest(taxRate TaxDataTaxRateInterface) *PutV1TaxRatesRequest {
	this := PutV1TaxRatesRequest{}
	this.TaxRate = taxRate
	return &this
}

// NewPutV1TaxRatesRequestWithDefaults instantiates a new PutV1TaxRatesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutV1TaxRatesRequestWithDefaults() *PutV1TaxRatesRequest {
	this := PutV1TaxRatesRequest{}
	return &this
}

// GetTaxRate returns the TaxRate field value
func (o *PutV1TaxRatesRequest) GetTaxRate() TaxDataTaxRateInterface {
	if o == nil {
		var ret TaxDataTaxRateInterface
		return ret
	}

	return o.TaxRate
}

// GetTaxRateOk returns a tuple with the TaxRate field value
// and a boolean to check if the value has been set.
func (o *PutV1TaxRatesRequest) GetTaxRateOk() (*TaxDataTaxRateInterface, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaxRate, true
}

// SetTaxRate sets field value
func (o *PutV1TaxRatesRequest) SetTaxRate(v TaxDataTaxRateInterface) {
	o.TaxRate = v
}

func (o PutV1TaxRatesRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutV1TaxRatesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["taxRate"] = o.TaxRate

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PutV1TaxRatesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"taxRate",
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

	varPutV1TaxRatesRequest := _PutV1TaxRatesRequest{}

	err = json.Unmarshal(data, &varPutV1TaxRatesRequest)

	if err != nil {
		return err
	}

	*o = PutV1TaxRatesRequest(varPutV1TaxRatesRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "taxRate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PutV1TaxRatesRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *PutV1TaxRatesRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullablePutV1TaxRatesRequest struct {
	value *PutV1TaxRatesRequest
	isSet bool
}

func (v NullablePutV1TaxRatesRequest) Get() *PutV1TaxRatesRequest {
	return v.value
}

func (v *NullablePutV1TaxRatesRequest) Set(val *PutV1TaxRatesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutV1TaxRatesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutV1TaxRatesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutV1TaxRatesRequest(val *PutV1TaxRatesRequest) *NullablePutV1TaxRatesRequest {
	return &NullablePutV1TaxRatesRequest{value: val, isSet: true}
}

func (v NullablePutV1TaxRatesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutV1TaxRatesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}