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

// checks if the StoreDataStoreConfigInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StoreDataStoreConfigInterface{}

// StoreDataStoreConfigInterface Interface for store config
type StoreDataStoreConfigInterface struct {
	// Store id
	Id int32 `json:"id"`
	// Store code
	Code string `json:"code"`
	// Website id of the store
	WebsiteId int32 `json:"website_id"`
	// Store locale
	Locale string `json:"locale"`
	// Base currency code
	BaseCurrencyCode string `json:"base_currency_code"`
	// Default display currency code
	DefaultDisplayCurrencyCode string `json:"default_display_currency_code"`
	// Timezone of the store
	Timezone string `json:"timezone"`
	// The unit of weight
	WeightUnit string `json:"weight_unit"`
	// Base URL for the store
	BaseUrl string `json:"base_url"`
	// Base link URL for the store
	BaseLinkUrl string `json:"base_link_url"`
	// Base static URL for the store
	BaseStaticUrl string `json:"base_static_url"`
	// Base media URL for the store
	BaseMediaUrl string `json:"base_media_url"`
	// Secure base URL for the store
	SecureBaseUrl string `json:"secure_base_url"`
	// Secure base link URL for the store
	SecureBaseLinkUrl string `json:"secure_base_link_url"`
	// Secure base static URL for the store
	SecureBaseStaticUrl string `json:"secure_base_static_url"`
	// Secure base media URL for the store
	SecureBaseMediaUrl string `json:"secure_base_media_url"`
	// ExtensionInterface class for @see \\Magento\\Store\\Api\\Data\\StoreConfigInterface
	ExtensionAttributes  map[string]interface{} `json:"extension_attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StoreDataStoreConfigInterface StoreDataStoreConfigInterface

// NewStoreDataStoreConfigInterface instantiates a new StoreDataStoreConfigInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoreDataStoreConfigInterface(id int32, code string, websiteId int32, locale string, baseCurrencyCode string, defaultDisplayCurrencyCode string, timezone string, weightUnit string, baseUrl string, baseLinkUrl string, baseStaticUrl string, baseMediaUrl string, secureBaseUrl string, secureBaseLinkUrl string, secureBaseStaticUrl string, secureBaseMediaUrl string) *StoreDataStoreConfigInterface {
	this := StoreDataStoreConfigInterface{}
	this.Id = id
	this.Code = code
	this.WebsiteId = websiteId
	this.Locale = locale
	this.BaseCurrencyCode = baseCurrencyCode
	this.DefaultDisplayCurrencyCode = defaultDisplayCurrencyCode
	this.Timezone = timezone
	this.WeightUnit = weightUnit
	this.BaseUrl = baseUrl
	this.BaseLinkUrl = baseLinkUrl
	this.BaseStaticUrl = baseStaticUrl
	this.BaseMediaUrl = baseMediaUrl
	this.SecureBaseUrl = secureBaseUrl
	this.SecureBaseLinkUrl = secureBaseLinkUrl
	this.SecureBaseStaticUrl = secureBaseStaticUrl
	this.SecureBaseMediaUrl = secureBaseMediaUrl
	return &this
}

// NewStoreDataStoreConfigInterfaceWithDefaults instantiates a new StoreDataStoreConfigInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoreDataStoreConfigInterfaceWithDefaults() *StoreDataStoreConfigInterface {
	this := StoreDataStoreConfigInterface{}
	return &this
}

// GetId returns the Id field value
func (o *StoreDataStoreConfigInterface) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StoreDataStoreConfigInterface) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *StoreDataStoreConfigInterface) SetId(v int32) {
	o.Id = v
}

// GetCode returns the Code field value
func (o *StoreDataStoreConfigInterface) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *StoreDataStoreConfigInterface) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *StoreDataStoreConfigInterface) SetCode(v string) {
	o.Code = v
}

// GetWebsiteId returns the WebsiteId field value
func (o *StoreDataStoreConfigInterface) GetWebsiteId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.WebsiteId
}

// GetWebsiteIdOk returns a tuple with the WebsiteId field value
// and a boolean to check if the value has been set.
func (o *StoreDataStoreConfigInterface) GetWebsiteIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WebsiteId, true
}

// SetWebsiteId sets field value
func (o *StoreDataStoreConfigInterface) SetWebsiteId(v int32) {
	o.WebsiteId = v
}

// GetLocale returns the Locale field value
func (o *StoreDataStoreConfigInterface) GetLocale() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value
// and a boolean to check if the value has been set.
func (o *StoreDataStoreConfigInterface) GetLocaleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locale, true
}

// SetLocale sets field value
func (o *StoreDataStoreConfigInterface) SetLocale(v string) {
	o.Locale = v
}

// GetBaseCurrencyCode returns the BaseCurrencyCode field value
func (o *StoreDataStoreConfigInterface) GetBaseCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseCurrencyCode
}

// GetBaseCurrencyCodeOk returns a tuple with the BaseCurrencyCode field value
// and a boolean to check if the value has been set.
func (o *StoreDataStoreConfigInterface) GetBaseCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseCurrencyCode, true
}

// SetBaseCurrencyCode sets field value
func (o *StoreDataStoreConfigInterface) SetBaseCurrencyCode(v string) {
	o.BaseCurrencyCode = v
}

// GetDefaultDisplayCurrencyCode returns the DefaultDisplayCurrencyCode field value
func (o *StoreDataStoreConfigInterface) GetDefaultDisplayCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultDisplayCurrencyCode
}

// GetDefaultDisplayCurrencyCodeOk returns a tuple with the DefaultDisplayCurrencyCode field value
// and a boolean to check if the value has been set.
func (o *StoreDataStoreConfigInterface) GetDefaultDisplayCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultDisplayCurrencyCode, true
}

// SetDefaultDisplayCurrencyCode sets field value
func (o *StoreDataStoreConfigInterface) SetDefaultDisplayCurrencyCode(v string) {
	o.DefaultDisplayCurrencyCode = v
}

// GetTimezone returns the Timezone field value
func (o *StoreDataStoreConfigInterface) GetTimezone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value
// and a boolean to check if the value has been set.
func (o *StoreDataStoreConfigInterface) GetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timezone, true
}

// SetTimezone sets field value
func (o *StoreDataStoreConfigInterface) SetTimezone(v string) {
	o.Timezone = v
}

// GetWeightUnit returns the WeightUnit field value
func (o *StoreDataStoreConfigInterface) GetWeightUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WeightUnit
}

// GetWeightUnitOk returns a tuple with the WeightUnit field value
// and a boolean to check if the value has been set.
func (o *StoreDataStoreConfigInterface) GetWeightUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WeightUnit, true
}

// SetWeightUnit sets field value
func (o *StoreDataStoreConfigInterface) SetWeightUnit(v string) {
	o.WeightUnit = v
}

// GetBaseUrl returns the BaseUrl field value
func (o *StoreDataStoreConfigInterface) GetBaseUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseUrl
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value
// and a boolean to check if the value has been set.
func (o *StoreDataStoreConfigInterface) GetBaseUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseUrl, true
}

// SetBaseUrl sets field value
func (o *StoreDataStoreConfigInterface) SetBaseUrl(v string) {
	o.BaseUrl = v
}

// GetBaseLinkUrl returns the BaseLinkUrl field value
func (o *StoreDataStoreConfigInterface) GetBaseLinkUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseLinkUrl
}

// GetBaseLinkUrlOk returns a tuple with the BaseLinkUrl field value
// and a boolean to check if the value has been set.
func (o *StoreDataStoreConfigInterface) GetBaseLinkUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseLinkUrl, true
}

// SetBaseLinkUrl sets field value
func (o *StoreDataStoreConfigInterface) SetBaseLinkUrl(v string) {
	o.BaseLinkUrl = v
}

// GetBaseStaticUrl returns the BaseStaticUrl field value
func (o *StoreDataStoreConfigInterface) GetBaseStaticUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseStaticUrl
}

// GetBaseStaticUrlOk returns a tuple with the BaseStaticUrl field value
// and a boolean to check if the value has been set.
func (o *StoreDataStoreConfigInterface) GetBaseStaticUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseStaticUrl, true
}

// SetBaseStaticUrl sets field value
func (o *StoreDataStoreConfigInterface) SetBaseStaticUrl(v string) {
	o.BaseStaticUrl = v
}

// GetBaseMediaUrl returns the BaseMediaUrl field value
func (o *StoreDataStoreConfigInterface) GetBaseMediaUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseMediaUrl
}

// GetBaseMediaUrlOk returns a tuple with the BaseMediaUrl field value
// and a boolean to check if the value has been set.
func (o *StoreDataStoreConfigInterface) GetBaseMediaUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseMediaUrl, true
}

// SetBaseMediaUrl sets field value
func (o *StoreDataStoreConfigInterface) SetBaseMediaUrl(v string) {
	o.BaseMediaUrl = v
}

// GetSecureBaseUrl returns the SecureBaseUrl field value
func (o *StoreDataStoreConfigInterface) GetSecureBaseUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecureBaseUrl
}

// GetSecureBaseUrlOk returns a tuple with the SecureBaseUrl field value
// and a boolean to check if the value has been set.
func (o *StoreDataStoreConfigInterface) GetSecureBaseUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecureBaseUrl, true
}

// SetSecureBaseUrl sets field value
func (o *StoreDataStoreConfigInterface) SetSecureBaseUrl(v string) {
	o.SecureBaseUrl = v
}

// GetSecureBaseLinkUrl returns the SecureBaseLinkUrl field value
func (o *StoreDataStoreConfigInterface) GetSecureBaseLinkUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecureBaseLinkUrl
}

// GetSecureBaseLinkUrlOk returns a tuple with the SecureBaseLinkUrl field value
// and a boolean to check if the value has been set.
func (o *StoreDataStoreConfigInterface) GetSecureBaseLinkUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecureBaseLinkUrl, true
}

// SetSecureBaseLinkUrl sets field value
func (o *StoreDataStoreConfigInterface) SetSecureBaseLinkUrl(v string) {
	o.SecureBaseLinkUrl = v
}

// GetSecureBaseStaticUrl returns the SecureBaseStaticUrl field value
func (o *StoreDataStoreConfigInterface) GetSecureBaseStaticUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecureBaseStaticUrl
}

// GetSecureBaseStaticUrlOk returns a tuple with the SecureBaseStaticUrl field value
// and a boolean to check if the value has been set.
func (o *StoreDataStoreConfigInterface) GetSecureBaseStaticUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecureBaseStaticUrl, true
}

// SetSecureBaseStaticUrl sets field value
func (o *StoreDataStoreConfigInterface) SetSecureBaseStaticUrl(v string) {
	o.SecureBaseStaticUrl = v
}

// GetSecureBaseMediaUrl returns the SecureBaseMediaUrl field value
func (o *StoreDataStoreConfigInterface) GetSecureBaseMediaUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecureBaseMediaUrl
}

// GetSecureBaseMediaUrlOk returns a tuple with the SecureBaseMediaUrl field value
// and a boolean to check if the value has been set.
func (o *StoreDataStoreConfigInterface) GetSecureBaseMediaUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecureBaseMediaUrl, true
}

// SetSecureBaseMediaUrl sets field value
func (o *StoreDataStoreConfigInterface) SetSecureBaseMediaUrl(v string) {
	o.SecureBaseMediaUrl = v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *StoreDataStoreConfigInterface) GetExtensionAttributes() map[string]interface{} {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreDataStoreConfigInterface) GetExtensionAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return map[string]interface{}{}, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *StoreDataStoreConfigInterface) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given map[string]interface{} and assigns it to the ExtensionAttributes field.
func (o *StoreDataStoreConfigInterface) SetExtensionAttributes(v map[string]interface{}) {
	o.ExtensionAttributes = v
}

func (o StoreDataStoreConfigInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StoreDataStoreConfigInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["code"] = o.Code
	toSerialize["website_id"] = o.WebsiteId
	toSerialize["locale"] = o.Locale
	toSerialize["base_currency_code"] = o.BaseCurrencyCode
	toSerialize["default_display_currency_code"] = o.DefaultDisplayCurrencyCode
	toSerialize["timezone"] = o.Timezone
	toSerialize["weight_unit"] = o.WeightUnit
	toSerialize["base_url"] = o.BaseUrl
	toSerialize["base_link_url"] = o.BaseLinkUrl
	toSerialize["base_static_url"] = o.BaseStaticUrl
	toSerialize["base_media_url"] = o.BaseMediaUrl
	toSerialize["secure_base_url"] = o.SecureBaseUrl
	toSerialize["secure_base_link_url"] = o.SecureBaseLinkUrl
	toSerialize["secure_base_static_url"] = o.SecureBaseStaticUrl
	toSerialize["secure_base_media_url"] = o.SecureBaseMediaUrl
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extension_attributes"] = o.ExtensionAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StoreDataStoreConfigInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"code",
		"website_id",
		"locale",
		"base_currency_code",
		"default_display_currency_code",
		"timezone",
		"weight_unit",
		"base_url",
		"base_link_url",
		"base_static_url",
		"base_media_url",
		"secure_base_url",
		"secure_base_link_url",
		"secure_base_static_url",
		"secure_base_media_url",
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

	varStoreDataStoreConfigInterface := _StoreDataStoreConfigInterface{}

	err = json.Unmarshal(data, &varStoreDataStoreConfigInterface)

	if err != nil {
		return err
	}

	*o = StoreDataStoreConfigInterface(varStoreDataStoreConfigInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "code")
		delete(additionalProperties, "website_id")
		delete(additionalProperties, "locale")
		delete(additionalProperties, "base_currency_code")
		delete(additionalProperties, "default_display_currency_code")
		delete(additionalProperties, "timezone")
		delete(additionalProperties, "weight_unit")
		delete(additionalProperties, "base_url")
		delete(additionalProperties, "base_link_url")
		delete(additionalProperties, "base_static_url")
		delete(additionalProperties, "base_media_url")
		delete(additionalProperties, "secure_base_url")
		delete(additionalProperties, "secure_base_link_url")
		delete(additionalProperties, "secure_base_static_url")
		delete(additionalProperties, "secure_base_media_url")
		delete(additionalProperties, "extension_attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *StoreDataStoreConfigInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *StoreDataStoreConfigInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableStoreDataStoreConfigInterface struct {
	value *StoreDataStoreConfigInterface
	isSet bool
}

func (v NullableStoreDataStoreConfigInterface) Get() *StoreDataStoreConfigInterface {
	return v.value
}

func (v *NullableStoreDataStoreConfigInterface) Set(val *StoreDataStoreConfigInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableStoreDataStoreConfigInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableStoreDataStoreConfigInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoreDataStoreConfigInterface(val *StoreDataStoreConfigInterface) *NullableStoreDataStoreConfigInterface {
	return &NullableStoreDataStoreConfigInterface{value: val, isSet: true}
}

func (v NullableStoreDataStoreConfigInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoreDataStoreConfigInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
