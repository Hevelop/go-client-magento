/*
Commerce Admin REST endpoints - All inclusive

The schemas documented here are autogenerated from an instance of Adobe Commerce with B2B. Each schema represents a specific user role (Admin, Customer, and Guest) and determines which endpoints are accessible. Use the version switcher to select an Adobe Commerce version and corresponding API.  You can also <a href=\"https://developer.adobe.com/commerce/webapi/rest/quick-reference/generate-local\" target=\"_blank\">generate a local API reference</a> based on your own Adobe Commerce configuration, which allows you to see API documentation for your specific Adobe Commerce modules, third-party modules, and extension attributes.

API version: 2.4.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package magento

import (
	"encoding/json"
)

// checks if the TaxDataOrderTaxDetailsAppliedTaxExtensionInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxDataOrderTaxDetailsAppliedTaxExtensionInterface{}

// TaxDataOrderTaxDetailsAppliedTaxExtensionInterface ExtensionInterface class for @see \\Magento\\Tax\\Api\\Data\\OrderTaxDetailsAppliedTaxInterface
type TaxDataOrderTaxDetailsAppliedTaxExtensionInterface struct {
	Rates                []TaxDataAppliedTaxRateInterface `json:"rates,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TaxDataOrderTaxDetailsAppliedTaxExtensionInterface TaxDataOrderTaxDetailsAppliedTaxExtensionInterface

// NewTaxDataOrderTaxDetailsAppliedTaxExtensionInterface instantiates a new TaxDataOrderTaxDetailsAppliedTaxExtensionInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxDataOrderTaxDetailsAppliedTaxExtensionInterface() *TaxDataOrderTaxDetailsAppliedTaxExtensionInterface {
	this := TaxDataOrderTaxDetailsAppliedTaxExtensionInterface{}
	return &this
}

// NewTaxDataOrderTaxDetailsAppliedTaxExtensionInterfaceWithDefaults instantiates a new TaxDataOrderTaxDetailsAppliedTaxExtensionInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxDataOrderTaxDetailsAppliedTaxExtensionInterfaceWithDefaults() *TaxDataOrderTaxDetailsAppliedTaxExtensionInterface {
	this := TaxDataOrderTaxDetailsAppliedTaxExtensionInterface{}
	return &this
}

// GetRates returns the Rates field value if set, zero value otherwise.
func (o *TaxDataOrderTaxDetailsAppliedTaxExtensionInterface) GetRates() []TaxDataAppliedTaxRateInterface {
	if o == nil || IsNil(o.Rates) {
		var ret []TaxDataAppliedTaxRateInterface
		return ret
	}
	return o.Rates
}

// GetRatesOk returns a tuple with the Rates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxDataOrderTaxDetailsAppliedTaxExtensionInterface) GetRatesOk() ([]TaxDataAppliedTaxRateInterface, bool) {
	if o == nil || IsNil(o.Rates) {
		return nil, false
	}
	return o.Rates, true
}

// HasRates returns a boolean if a field has been set.
func (o *TaxDataOrderTaxDetailsAppliedTaxExtensionInterface) HasRates() bool {
	if o != nil && !IsNil(o.Rates) {
		return true
	}

	return false
}

// SetRates gets a reference to the given []TaxDataAppliedTaxRateInterface and assigns it to the Rates field.
func (o *TaxDataOrderTaxDetailsAppliedTaxExtensionInterface) SetRates(v []TaxDataAppliedTaxRateInterface) {
	o.Rates = v
}

func (o TaxDataOrderTaxDetailsAppliedTaxExtensionInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaxDataOrderTaxDetailsAppliedTaxExtensionInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rates) {
		toSerialize["rates"] = o.Rates
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TaxDataOrderTaxDetailsAppliedTaxExtensionInterface) UnmarshalJSON(data []byte) (err error) {
	varTaxDataOrderTaxDetailsAppliedTaxExtensionInterface := _TaxDataOrderTaxDetailsAppliedTaxExtensionInterface{}

	err = json.Unmarshal(data, &varTaxDataOrderTaxDetailsAppliedTaxExtensionInterface)

	if err != nil {
		return err
	}

	*o = TaxDataOrderTaxDetailsAppliedTaxExtensionInterface(varTaxDataOrderTaxDetailsAppliedTaxExtensionInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rates")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *TaxDataOrderTaxDetailsAppliedTaxExtensionInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *TaxDataOrderTaxDetailsAppliedTaxExtensionInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableTaxDataOrderTaxDetailsAppliedTaxExtensionInterface struct {
	value *TaxDataOrderTaxDetailsAppliedTaxExtensionInterface
	isSet bool
}

func (v NullableTaxDataOrderTaxDetailsAppliedTaxExtensionInterface) Get() *TaxDataOrderTaxDetailsAppliedTaxExtensionInterface {
	return v.value
}

func (v *NullableTaxDataOrderTaxDetailsAppliedTaxExtensionInterface) Set(val *TaxDataOrderTaxDetailsAppliedTaxExtensionInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxDataOrderTaxDetailsAppliedTaxExtensionInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxDataOrderTaxDetailsAppliedTaxExtensionInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxDataOrderTaxDetailsAppliedTaxExtensionInterface(val *TaxDataOrderTaxDetailsAppliedTaxExtensionInterface) *NullableTaxDataOrderTaxDetailsAppliedTaxExtensionInterface {
	return &NullableTaxDataOrderTaxDetailsAppliedTaxExtensionInterface{value: val, isSet: true}
}

func (v NullableTaxDataOrderTaxDetailsAppliedTaxExtensionInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxDataOrderTaxDetailsAppliedTaxExtensionInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
