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

// checks if the DirectoryDataCurrencyInformationInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DirectoryDataCurrencyInformationInterface{}

// DirectoryDataCurrencyInformationInterface Currency Information interface.
type DirectoryDataCurrencyInformationInterface struct {
	// The base currency code for the store.
	BaseCurrencyCode string `json:"base_currency_code"`
	// The currency symbol of the base currency for the store.
	BaseCurrencySymbol string `json:"base_currency_symbol"`
	// The default display currency code for the store.
	DefaultDisplayCurrencyCode string `json:"default_display_currency_code"`
	// The currency symbol of the default display currency for the store.
	DefaultDisplayCurrencySymbol string `json:"default_display_currency_symbol"`
	// The list of allowed currency codes for the store.
	AvailableCurrencyCodes []string `json:"available_currency_codes"`
	// The list of exchange rate information for the store.
	ExchangeRates []DirectoryDataExchangeRateInterface `json:"exchange_rates"`
	// ExtensionInterface class for @see \\Magento\\Directory\\Api\\Data\\CurrencyInformationInterface
	ExtensionAttributes  map[string]interface{} `json:"extension_attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DirectoryDataCurrencyInformationInterface DirectoryDataCurrencyInformationInterface

// NewDirectoryDataCurrencyInformationInterface instantiates a new DirectoryDataCurrencyInformationInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDirectoryDataCurrencyInformationInterface(baseCurrencyCode string, baseCurrencySymbol string, defaultDisplayCurrencyCode string, defaultDisplayCurrencySymbol string, availableCurrencyCodes []string, exchangeRates []DirectoryDataExchangeRateInterface) *DirectoryDataCurrencyInformationInterface {
	this := DirectoryDataCurrencyInformationInterface{}
	this.BaseCurrencyCode = baseCurrencyCode
	this.BaseCurrencySymbol = baseCurrencySymbol
	this.DefaultDisplayCurrencyCode = defaultDisplayCurrencyCode
	this.DefaultDisplayCurrencySymbol = defaultDisplayCurrencySymbol
	this.AvailableCurrencyCodes = availableCurrencyCodes
	this.ExchangeRates = exchangeRates
	return &this
}

// NewDirectoryDataCurrencyInformationInterfaceWithDefaults instantiates a new DirectoryDataCurrencyInformationInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDirectoryDataCurrencyInformationInterfaceWithDefaults() *DirectoryDataCurrencyInformationInterface {
	this := DirectoryDataCurrencyInformationInterface{}
	return &this
}

// GetBaseCurrencyCode returns the BaseCurrencyCode field value
func (o *DirectoryDataCurrencyInformationInterface) GetBaseCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseCurrencyCode
}

// GetBaseCurrencyCodeOk returns a tuple with the BaseCurrencyCode field value
// and a boolean to check if the value has been set.
func (o *DirectoryDataCurrencyInformationInterface) GetBaseCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseCurrencyCode, true
}

// SetBaseCurrencyCode sets field value
func (o *DirectoryDataCurrencyInformationInterface) SetBaseCurrencyCode(v string) {
	o.BaseCurrencyCode = v
}

// GetBaseCurrencySymbol returns the BaseCurrencySymbol field value
func (o *DirectoryDataCurrencyInformationInterface) GetBaseCurrencySymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseCurrencySymbol
}

// GetBaseCurrencySymbolOk returns a tuple with the BaseCurrencySymbol field value
// and a boolean to check if the value has been set.
func (o *DirectoryDataCurrencyInformationInterface) GetBaseCurrencySymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseCurrencySymbol, true
}

// SetBaseCurrencySymbol sets field value
func (o *DirectoryDataCurrencyInformationInterface) SetBaseCurrencySymbol(v string) {
	o.BaseCurrencySymbol = v
}

// GetDefaultDisplayCurrencyCode returns the DefaultDisplayCurrencyCode field value
func (o *DirectoryDataCurrencyInformationInterface) GetDefaultDisplayCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultDisplayCurrencyCode
}

// GetDefaultDisplayCurrencyCodeOk returns a tuple with the DefaultDisplayCurrencyCode field value
// and a boolean to check if the value has been set.
func (o *DirectoryDataCurrencyInformationInterface) GetDefaultDisplayCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultDisplayCurrencyCode, true
}

// SetDefaultDisplayCurrencyCode sets field value
func (o *DirectoryDataCurrencyInformationInterface) SetDefaultDisplayCurrencyCode(v string) {
	o.DefaultDisplayCurrencyCode = v
}

// GetDefaultDisplayCurrencySymbol returns the DefaultDisplayCurrencySymbol field value
func (o *DirectoryDataCurrencyInformationInterface) GetDefaultDisplayCurrencySymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultDisplayCurrencySymbol
}

// GetDefaultDisplayCurrencySymbolOk returns a tuple with the DefaultDisplayCurrencySymbol field value
// and a boolean to check if the value has been set.
func (o *DirectoryDataCurrencyInformationInterface) GetDefaultDisplayCurrencySymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultDisplayCurrencySymbol, true
}

// SetDefaultDisplayCurrencySymbol sets field value
func (o *DirectoryDataCurrencyInformationInterface) SetDefaultDisplayCurrencySymbol(v string) {
	o.DefaultDisplayCurrencySymbol = v
}

// GetAvailableCurrencyCodes returns the AvailableCurrencyCodes field value
func (o *DirectoryDataCurrencyInformationInterface) GetAvailableCurrencyCodes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AvailableCurrencyCodes
}

// GetAvailableCurrencyCodesOk returns a tuple with the AvailableCurrencyCodes field value
// and a boolean to check if the value has been set.
func (o *DirectoryDataCurrencyInformationInterface) GetAvailableCurrencyCodesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailableCurrencyCodes, true
}

// SetAvailableCurrencyCodes sets field value
func (o *DirectoryDataCurrencyInformationInterface) SetAvailableCurrencyCodes(v []string) {
	o.AvailableCurrencyCodes = v
}

// GetExchangeRates returns the ExchangeRates field value
func (o *DirectoryDataCurrencyInformationInterface) GetExchangeRates() []DirectoryDataExchangeRateInterface {
	if o == nil {
		var ret []DirectoryDataExchangeRateInterface
		return ret
	}

	return o.ExchangeRates
}

// GetExchangeRatesOk returns a tuple with the ExchangeRates field value
// and a boolean to check if the value has been set.
func (o *DirectoryDataCurrencyInformationInterface) GetExchangeRatesOk() ([]DirectoryDataExchangeRateInterface, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExchangeRates, true
}

// SetExchangeRates sets field value
func (o *DirectoryDataCurrencyInformationInterface) SetExchangeRates(v []DirectoryDataExchangeRateInterface) {
	o.ExchangeRates = v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *DirectoryDataCurrencyInformationInterface) GetExtensionAttributes() map[string]interface{} {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectoryDataCurrencyInformationInterface) GetExtensionAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return map[string]interface{}{}, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *DirectoryDataCurrencyInformationInterface) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given map[string]interface{} and assigns it to the ExtensionAttributes field.
func (o *DirectoryDataCurrencyInformationInterface) SetExtensionAttributes(v map[string]interface{}) {
	o.ExtensionAttributes = v
}

func (o DirectoryDataCurrencyInformationInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DirectoryDataCurrencyInformationInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["base_currency_code"] = o.BaseCurrencyCode
	toSerialize["base_currency_symbol"] = o.BaseCurrencySymbol
	toSerialize["default_display_currency_code"] = o.DefaultDisplayCurrencyCode
	toSerialize["default_display_currency_symbol"] = o.DefaultDisplayCurrencySymbol
	toSerialize["available_currency_codes"] = o.AvailableCurrencyCodes
	toSerialize["exchange_rates"] = o.ExchangeRates
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extension_attributes"] = o.ExtensionAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DirectoryDataCurrencyInformationInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"base_currency_code",
		"base_currency_symbol",
		"default_display_currency_code",
		"default_display_currency_symbol",
		"available_currency_codes",
		"exchange_rates",
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

	varDirectoryDataCurrencyInformationInterface := _DirectoryDataCurrencyInformationInterface{}

	err = json.Unmarshal(data, &varDirectoryDataCurrencyInformationInterface)

	if err != nil {
		return err
	}

	*o = DirectoryDataCurrencyInformationInterface(varDirectoryDataCurrencyInformationInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "base_currency_code")
		delete(additionalProperties, "base_currency_symbol")
		delete(additionalProperties, "default_display_currency_code")
		delete(additionalProperties, "default_display_currency_symbol")
		delete(additionalProperties, "available_currency_codes")
		delete(additionalProperties, "exchange_rates")
		delete(additionalProperties, "extension_attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *DirectoryDataCurrencyInformationInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *DirectoryDataCurrencyInformationInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableDirectoryDataCurrencyInformationInterface struct {
	value *DirectoryDataCurrencyInformationInterface
	isSet bool
}

func (v NullableDirectoryDataCurrencyInformationInterface) Get() *DirectoryDataCurrencyInformationInterface {
	return v.value
}

func (v *NullableDirectoryDataCurrencyInformationInterface) Set(val *DirectoryDataCurrencyInformationInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectoryDataCurrencyInformationInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectoryDataCurrencyInformationInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectoryDataCurrencyInformationInterface(val *DirectoryDataCurrencyInformationInterface) *NullableDirectoryDataCurrencyInformationInterface {
	return &NullableDirectoryDataCurrencyInformationInterface{value: val, isSet: true}
}

func (v NullableDirectoryDataCurrencyInformationInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDirectoryDataCurrencyInformationInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
