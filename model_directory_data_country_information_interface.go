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

// checks if the DirectoryDataCountryInformationInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DirectoryDataCountryInformationInterface{}

// DirectoryDataCountryInformationInterface Country Information interface.
type DirectoryDataCountryInformationInterface struct {
	// The country id for the store.
	Id string `json:"id"`
	// The country 2 letter abbreviation for the store.
	TwoLetterAbbreviation string `json:"two_letter_abbreviation"`
	// The country 3 letter abbreviation for the store.
	ThreeLetterAbbreviation string `json:"three_letter_abbreviation"`
	// The country full name (in store locale) for the store.
	FullNameLocale string `json:"full_name_locale"`
	// The country full name (in English) for the store.
	FullNameEnglish string `json:"full_name_english"`
	// The available regions for the store.
	AvailableRegions []DirectoryDataRegionInformationInterface `json:"available_regions,omitempty"`
	// ExtensionInterface class for @see \\Magento\\Directory\\Api\\Data\\CountryInformationInterface
	ExtensionAttributes  map[string]interface{} `json:"extension_attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DirectoryDataCountryInformationInterface DirectoryDataCountryInformationInterface

// NewDirectoryDataCountryInformationInterface instantiates a new DirectoryDataCountryInformationInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDirectoryDataCountryInformationInterface(id string, twoLetterAbbreviation string, threeLetterAbbreviation string, fullNameLocale string, fullNameEnglish string) *DirectoryDataCountryInformationInterface {
	this := DirectoryDataCountryInformationInterface{}
	this.Id = id
	this.TwoLetterAbbreviation = twoLetterAbbreviation
	this.ThreeLetterAbbreviation = threeLetterAbbreviation
	this.FullNameLocale = fullNameLocale
	this.FullNameEnglish = fullNameEnglish
	return &this
}

// NewDirectoryDataCountryInformationInterfaceWithDefaults instantiates a new DirectoryDataCountryInformationInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDirectoryDataCountryInformationInterfaceWithDefaults() *DirectoryDataCountryInformationInterface {
	this := DirectoryDataCountryInformationInterface{}
	return &this
}

// GetId returns the Id field value
func (o *DirectoryDataCountryInformationInterface) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DirectoryDataCountryInformationInterface) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DirectoryDataCountryInformationInterface) SetId(v string) {
	o.Id = v
}

// GetTwoLetterAbbreviation returns the TwoLetterAbbreviation field value
func (o *DirectoryDataCountryInformationInterface) GetTwoLetterAbbreviation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TwoLetterAbbreviation
}

// GetTwoLetterAbbreviationOk returns a tuple with the TwoLetterAbbreviation field value
// and a boolean to check if the value has been set.
func (o *DirectoryDataCountryInformationInterface) GetTwoLetterAbbreviationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TwoLetterAbbreviation, true
}

// SetTwoLetterAbbreviation sets field value
func (o *DirectoryDataCountryInformationInterface) SetTwoLetterAbbreviation(v string) {
	o.TwoLetterAbbreviation = v
}

// GetThreeLetterAbbreviation returns the ThreeLetterAbbreviation field value
func (o *DirectoryDataCountryInformationInterface) GetThreeLetterAbbreviation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThreeLetterAbbreviation
}

// GetThreeLetterAbbreviationOk returns a tuple with the ThreeLetterAbbreviation field value
// and a boolean to check if the value has been set.
func (o *DirectoryDataCountryInformationInterface) GetThreeLetterAbbreviationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThreeLetterAbbreviation, true
}

// SetThreeLetterAbbreviation sets field value
func (o *DirectoryDataCountryInformationInterface) SetThreeLetterAbbreviation(v string) {
	o.ThreeLetterAbbreviation = v
}

// GetFullNameLocale returns the FullNameLocale field value
func (o *DirectoryDataCountryInformationInterface) GetFullNameLocale() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullNameLocale
}

// GetFullNameLocaleOk returns a tuple with the FullNameLocale field value
// and a boolean to check if the value has been set.
func (o *DirectoryDataCountryInformationInterface) GetFullNameLocaleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullNameLocale, true
}

// SetFullNameLocale sets field value
func (o *DirectoryDataCountryInformationInterface) SetFullNameLocale(v string) {
	o.FullNameLocale = v
}

// GetFullNameEnglish returns the FullNameEnglish field value
func (o *DirectoryDataCountryInformationInterface) GetFullNameEnglish() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullNameEnglish
}

// GetFullNameEnglishOk returns a tuple with the FullNameEnglish field value
// and a boolean to check if the value has been set.
func (o *DirectoryDataCountryInformationInterface) GetFullNameEnglishOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullNameEnglish, true
}

// SetFullNameEnglish sets field value
func (o *DirectoryDataCountryInformationInterface) SetFullNameEnglish(v string) {
	o.FullNameEnglish = v
}

// GetAvailableRegions returns the AvailableRegions field value if set, zero value otherwise.
func (o *DirectoryDataCountryInformationInterface) GetAvailableRegions() []DirectoryDataRegionInformationInterface {
	if o == nil || IsNil(o.AvailableRegions) {
		var ret []DirectoryDataRegionInformationInterface
		return ret
	}
	return o.AvailableRegions
}

// GetAvailableRegionsOk returns a tuple with the AvailableRegions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectoryDataCountryInformationInterface) GetAvailableRegionsOk() ([]DirectoryDataRegionInformationInterface, bool) {
	if o == nil || IsNil(o.AvailableRegions) {
		return nil, false
	}
	return o.AvailableRegions, true
}

// HasAvailableRegions returns a boolean if a field has been set.
func (o *DirectoryDataCountryInformationInterface) HasAvailableRegions() bool {
	if o != nil && !IsNil(o.AvailableRegions) {
		return true
	}

	return false
}

// SetAvailableRegions gets a reference to the given []DirectoryDataRegionInformationInterface and assigns it to the AvailableRegions field.
func (o *DirectoryDataCountryInformationInterface) SetAvailableRegions(v []DirectoryDataRegionInformationInterface) {
	o.AvailableRegions = v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *DirectoryDataCountryInformationInterface) GetExtensionAttributes() map[string]interface{} {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectoryDataCountryInformationInterface) GetExtensionAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return map[string]interface{}{}, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *DirectoryDataCountryInformationInterface) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given map[string]interface{} and assigns it to the ExtensionAttributes field.
func (o *DirectoryDataCountryInformationInterface) SetExtensionAttributes(v map[string]interface{}) {
	o.ExtensionAttributes = v
}

func (o DirectoryDataCountryInformationInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DirectoryDataCountryInformationInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["two_letter_abbreviation"] = o.TwoLetterAbbreviation
	toSerialize["three_letter_abbreviation"] = o.ThreeLetterAbbreviation
	toSerialize["full_name_locale"] = o.FullNameLocale
	toSerialize["full_name_english"] = o.FullNameEnglish
	if !IsNil(o.AvailableRegions) {
		toSerialize["available_regions"] = o.AvailableRegions
	}
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extension_attributes"] = o.ExtensionAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DirectoryDataCountryInformationInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"two_letter_abbreviation",
		"three_letter_abbreviation",
		"full_name_locale",
		"full_name_english",
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

	varDirectoryDataCountryInformationInterface := _DirectoryDataCountryInformationInterface{}

	err = json.Unmarshal(data, &varDirectoryDataCountryInformationInterface)

	if err != nil {
		return err
	}

	*o = DirectoryDataCountryInformationInterface(varDirectoryDataCountryInformationInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "two_letter_abbreviation")
		delete(additionalProperties, "three_letter_abbreviation")
		delete(additionalProperties, "full_name_locale")
		delete(additionalProperties, "full_name_english")
		delete(additionalProperties, "available_regions")
		delete(additionalProperties, "extension_attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *DirectoryDataCountryInformationInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *DirectoryDataCountryInformationInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableDirectoryDataCountryInformationInterface struct {
	value *DirectoryDataCountryInformationInterface
	isSet bool
}

func (v NullableDirectoryDataCountryInformationInterface) Get() *DirectoryDataCountryInformationInterface {
	return v.value
}

func (v *NullableDirectoryDataCountryInformationInterface) Set(val *DirectoryDataCountryInformationInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectoryDataCountryInformationInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectoryDataCountryInformationInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectoryDataCountryInformationInterface(val *DirectoryDataCountryInformationInterface) *NullableDirectoryDataCountryInformationInterface {
	return &NullableDirectoryDataCountryInformationInterface{value: val, isSet: true}
}

func (v NullableDirectoryDataCountryInformationInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDirectoryDataCountryInformationInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
