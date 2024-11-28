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

// checks if the DirectoryDataExchangeRateInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DirectoryDataExchangeRateInterface{}

// DirectoryDataExchangeRateInterface Exchange Rate interface.
type DirectoryDataExchangeRateInterface struct {
	// The currency code associated with the exchange rate.
	CurrencyTo string `json:"currency_to"`
	// The exchange rate for the associated currency and the store's base currency.
	Rate float32 `json:"rate"`
	// ExtensionInterface class for @see \\Magento\\Directory\\Api\\Data\\ExchangeRateInterface
	ExtensionAttributes  map[string]interface{} `json:"extension_attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DirectoryDataExchangeRateInterface DirectoryDataExchangeRateInterface

// NewDirectoryDataExchangeRateInterface instantiates a new DirectoryDataExchangeRateInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDirectoryDataExchangeRateInterface(currencyTo string, rate float32) *DirectoryDataExchangeRateInterface {
	this := DirectoryDataExchangeRateInterface{}
	this.CurrencyTo = currencyTo
	this.Rate = rate
	return &this
}

// NewDirectoryDataExchangeRateInterfaceWithDefaults instantiates a new DirectoryDataExchangeRateInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDirectoryDataExchangeRateInterfaceWithDefaults() *DirectoryDataExchangeRateInterface {
	this := DirectoryDataExchangeRateInterface{}
	return &this
}

// GetCurrencyTo returns the CurrencyTo field value
func (o *DirectoryDataExchangeRateInterface) GetCurrencyTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyTo
}

// GetCurrencyToOk returns a tuple with the CurrencyTo field value
// and a boolean to check if the value has been set.
func (o *DirectoryDataExchangeRateInterface) GetCurrencyToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyTo, true
}

// SetCurrencyTo sets field value
func (o *DirectoryDataExchangeRateInterface) SetCurrencyTo(v string) {
	o.CurrencyTo = v
}

// GetRate returns the Rate field value
func (o *DirectoryDataExchangeRateInterface) GetRate() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Rate
}

// GetRateOk returns a tuple with the Rate field value
// and a boolean to check if the value has been set.
func (o *DirectoryDataExchangeRateInterface) GetRateOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rate, true
}

// SetRate sets field value
func (o *DirectoryDataExchangeRateInterface) SetRate(v float32) {
	o.Rate = v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *DirectoryDataExchangeRateInterface) GetExtensionAttributes() map[string]interface{} {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectoryDataExchangeRateInterface) GetExtensionAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return map[string]interface{}{}, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *DirectoryDataExchangeRateInterface) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given map[string]interface{} and assigns it to the ExtensionAttributes field.
func (o *DirectoryDataExchangeRateInterface) SetExtensionAttributes(v map[string]interface{}) {
	o.ExtensionAttributes = v
}

func (o DirectoryDataExchangeRateInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DirectoryDataExchangeRateInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["currency_to"] = o.CurrencyTo
	toSerialize["rate"] = o.Rate
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extension_attributes"] = o.ExtensionAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DirectoryDataExchangeRateInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"currency_to",
		"rate",
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

	varDirectoryDataExchangeRateInterface := _DirectoryDataExchangeRateInterface{}

	err = json.Unmarshal(data, &varDirectoryDataExchangeRateInterface)

	if err != nil {
		return err
	}

	*o = DirectoryDataExchangeRateInterface(varDirectoryDataExchangeRateInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "currency_to")
		delete(additionalProperties, "rate")
		delete(additionalProperties, "extension_attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *DirectoryDataExchangeRateInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *DirectoryDataExchangeRateInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableDirectoryDataExchangeRateInterface struct {
	value *DirectoryDataExchangeRateInterface
	isSet bool
}

func (v NullableDirectoryDataExchangeRateInterface) Get() *DirectoryDataExchangeRateInterface {
	return v.value
}

func (v *NullableDirectoryDataExchangeRateInterface) Set(val *DirectoryDataExchangeRateInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectoryDataExchangeRateInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectoryDataExchangeRateInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectoryDataExchangeRateInterface(val *DirectoryDataExchangeRateInterface) *NullableDirectoryDataExchangeRateInterface {
	return &NullableDirectoryDataExchangeRateInterface{value: val, isSet: true}
}

func (v NullableDirectoryDataExchangeRateInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDirectoryDataExchangeRateInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
