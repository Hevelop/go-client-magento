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

// checks if the TaxDataGrandTotalDetailsInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxDataGrandTotalDetailsInterface{}

// TaxDataGrandTotalDetailsInterface Interface GrandTotalDetailsInterface
type TaxDataGrandTotalDetailsInterface struct {
	// Tax amount value
	Amount float32 `json:"amount"`
	// Tax rates info
	Rates []TaxDataGrandTotalRatesInterface `json:"rates"`
	// Group identifier
	GroupId              int32 `json:"group_id"`
	AdditionalProperties map[string]interface{}
}

type _TaxDataGrandTotalDetailsInterface TaxDataGrandTotalDetailsInterface

// NewTaxDataGrandTotalDetailsInterface instantiates a new TaxDataGrandTotalDetailsInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxDataGrandTotalDetailsInterface(amount float32, rates []TaxDataGrandTotalRatesInterface, groupId int32) *TaxDataGrandTotalDetailsInterface {
	this := TaxDataGrandTotalDetailsInterface{}
	this.Amount = amount
	this.Rates = rates
	this.GroupId = groupId
	return &this
}

// NewTaxDataGrandTotalDetailsInterfaceWithDefaults instantiates a new TaxDataGrandTotalDetailsInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxDataGrandTotalDetailsInterfaceWithDefaults() *TaxDataGrandTotalDetailsInterface {
	this := TaxDataGrandTotalDetailsInterface{}
	return &this
}

// GetAmount returns the Amount field value
func (o *TaxDataGrandTotalDetailsInterface) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TaxDataGrandTotalDetailsInterface) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TaxDataGrandTotalDetailsInterface) SetAmount(v float32) {
	o.Amount = v
}

// GetRates returns the Rates field value
func (o *TaxDataGrandTotalDetailsInterface) GetRates() []TaxDataGrandTotalRatesInterface {
	if o == nil {
		var ret []TaxDataGrandTotalRatesInterface
		return ret
	}

	return o.Rates
}

// GetRatesOk returns a tuple with the Rates field value
// and a boolean to check if the value has been set.
func (o *TaxDataGrandTotalDetailsInterface) GetRatesOk() ([]TaxDataGrandTotalRatesInterface, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rates, true
}

// SetRates sets field value
func (o *TaxDataGrandTotalDetailsInterface) SetRates(v []TaxDataGrandTotalRatesInterface) {
	o.Rates = v
}

// GetGroupId returns the GroupId field value
func (o *TaxDataGrandTotalDetailsInterface) GetGroupId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *TaxDataGrandTotalDetailsInterface) GetGroupIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *TaxDataGrandTotalDetailsInterface) SetGroupId(v int32) {
	o.GroupId = v
}

func (o TaxDataGrandTotalDetailsInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaxDataGrandTotalDetailsInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["rates"] = o.Rates
	toSerialize["group_id"] = o.GroupId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TaxDataGrandTotalDetailsInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
		"rates",
		"group_id",
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

	varTaxDataGrandTotalDetailsInterface := _TaxDataGrandTotalDetailsInterface{}

	err = json.Unmarshal(data, &varTaxDataGrandTotalDetailsInterface)

	if err != nil {
		return err
	}

	*o = TaxDataGrandTotalDetailsInterface(varTaxDataGrandTotalDetailsInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "rates")
		delete(additionalProperties, "group_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *TaxDataGrandTotalDetailsInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *TaxDataGrandTotalDetailsInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableTaxDataGrandTotalDetailsInterface struct {
	value *TaxDataGrandTotalDetailsInterface
	isSet bool
}

func (v NullableTaxDataGrandTotalDetailsInterface) Get() *TaxDataGrandTotalDetailsInterface {
	return v.value
}

func (v *NullableTaxDataGrandTotalDetailsInterface) Set(val *TaxDataGrandTotalDetailsInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxDataGrandTotalDetailsInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxDataGrandTotalDetailsInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxDataGrandTotalDetailsInterface(val *TaxDataGrandTotalDetailsInterface) *NullableTaxDataGrandTotalDetailsInterface {
	return &NullableTaxDataGrandTotalDetailsInterface{value: val, isSet: true}
}

func (v NullableTaxDataGrandTotalDetailsInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxDataGrandTotalDetailsInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
