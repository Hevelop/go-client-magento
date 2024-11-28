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

// checks if the CompanyDataCompanyExtensionInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompanyDataCompanyExtensionInterface{}

// CompanyDataCompanyExtensionInterface ExtensionInterface class for @see \\Magento\\Company\\Api\\Data\\CompanyInterface
type CompanyDataCompanyExtensionInterface struct {
	ApplicablePaymentMethod   *int32                                          `json:"applicable_payment_method,omitempty"`
	AvailablePaymentMethods   *string                                         `json:"available_payment_methods,omitempty"`
	UseConfigSettings         *int32                                          `json:"use_config_settings,omitempty"`
	QuoteConfig               *NegotiableQuoteDataCompanyQuoteConfigInterface `json:"quote_config,omitempty"`
	IsPurchaseOrderEnabled    *bool                                           `json:"is_purchase_order_enabled,omitempty"`
	ApplicableShippingMethod  *int32                                          `json:"applicable_shipping_method,omitempty"`
	AvailableShippingMethods  *string                                         `json:"available_shipping_methods,omitempty"`
	UseConfigSettingsShipping *int32                                          `json:"use_config_settings_shipping,omitempty"`
	AdditionalProperties      map[string]interface{}
}

type _CompanyDataCompanyExtensionInterface CompanyDataCompanyExtensionInterface

// NewCompanyDataCompanyExtensionInterface instantiates a new CompanyDataCompanyExtensionInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyDataCompanyExtensionInterface() *CompanyDataCompanyExtensionInterface {
	this := CompanyDataCompanyExtensionInterface{}
	return &this
}

// NewCompanyDataCompanyExtensionInterfaceWithDefaults instantiates a new CompanyDataCompanyExtensionInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyDataCompanyExtensionInterfaceWithDefaults() *CompanyDataCompanyExtensionInterface {
	this := CompanyDataCompanyExtensionInterface{}
	return &this
}

// GetApplicablePaymentMethod returns the ApplicablePaymentMethod field value if set, zero value otherwise.
func (o *CompanyDataCompanyExtensionInterface) GetApplicablePaymentMethod() int32 {
	if o == nil || IsNil(o.ApplicablePaymentMethod) {
		var ret int32
		return ret
	}
	return *o.ApplicablePaymentMethod
}

// GetApplicablePaymentMethodOk returns a tuple with the ApplicablePaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyDataCompanyExtensionInterface) GetApplicablePaymentMethodOk() (*int32, bool) {
	if o == nil || IsNil(o.ApplicablePaymentMethod) {
		return nil, false
	}
	return o.ApplicablePaymentMethod, true
}

// HasApplicablePaymentMethod returns a boolean if a field has been set.
func (o *CompanyDataCompanyExtensionInterface) HasApplicablePaymentMethod() bool {
	if o != nil && !IsNil(o.ApplicablePaymentMethod) {
		return true
	}

	return false
}

// SetApplicablePaymentMethod gets a reference to the given int32 and assigns it to the ApplicablePaymentMethod field.
func (o *CompanyDataCompanyExtensionInterface) SetApplicablePaymentMethod(v int32) {
	o.ApplicablePaymentMethod = &v
}

// GetAvailablePaymentMethods returns the AvailablePaymentMethods field value if set, zero value otherwise.
func (o *CompanyDataCompanyExtensionInterface) GetAvailablePaymentMethods() string {
	if o == nil || IsNil(o.AvailablePaymentMethods) {
		var ret string
		return ret
	}
	return *o.AvailablePaymentMethods
}

// GetAvailablePaymentMethodsOk returns a tuple with the AvailablePaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyDataCompanyExtensionInterface) GetAvailablePaymentMethodsOk() (*string, bool) {
	if o == nil || IsNil(o.AvailablePaymentMethods) {
		return nil, false
	}
	return o.AvailablePaymentMethods, true
}

// HasAvailablePaymentMethods returns a boolean if a field has been set.
func (o *CompanyDataCompanyExtensionInterface) HasAvailablePaymentMethods() bool {
	if o != nil && !IsNil(o.AvailablePaymentMethods) {
		return true
	}

	return false
}

// SetAvailablePaymentMethods gets a reference to the given string and assigns it to the AvailablePaymentMethods field.
func (o *CompanyDataCompanyExtensionInterface) SetAvailablePaymentMethods(v string) {
	o.AvailablePaymentMethods = &v
}

// GetUseConfigSettings returns the UseConfigSettings field value if set, zero value otherwise.
func (o *CompanyDataCompanyExtensionInterface) GetUseConfigSettings() int32 {
	if o == nil || IsNil(o.UseConfigSettings) {
		var ret int32
		return ret
	}
	return *o.UseConfigSettings
}

// GetUseConfigSettingsOk returns a tuple with the UseConfigSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyDataCompanyExtensionInterface) GetUseConfigSettingsOk() (*int32, bool) {
	if o == nil || IsNil(o.UseConfigSettings) {
		return nil, false
	}
	return o.UseConfigSettings, true
}

// HasUseConfigSettings returns a boolean if a field has been set.
func (o *CompanyDataCompanyExtensionInterface) HasUseConfigSettings() bool {
	if o != nil && !IsNil(o.UseConfigSettings) {
		return true
	}

	return false
}

// SetUseConfigSettings gets a reference to the given int32 and assigns it to the UseConfigSettings field.
func (o *CompanyDataCompanyExtensionInterface) SetUseConfigSettings(v int32) {
	o.UseConfigSettings = &v
}

// GetQuoteConfig returns the QuoteConfig field value if set, zero value otherwise.
func (o *CompanyDataCompanyExtensionInterface) GetQuoteConfig() NegotiableQuoteDataCompanyQuoteConfigInterface {
	if o == nil || IsNil(o.QuoteConfig) {
		var ret NegotiableQuoteDataCompanyQuoteConfigInterface
		return ret
	}
	return *o.QuoteConfig
}

// GetQuoteConfigOk returns a tuple with the QuoteConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyDataCompanyExtensionInterface) GetQuoteConfigOk() (*NegotiableQuoteDataCompanyQuoteConfigInterface, bool) {
	if o == nil || IsNil(o.QuoteConfig) {
		return nil, false
	}
	return o.QuoteConfig, true
}

// HasQuoteConfig returns a boolean if a field has been set.
func (o *CompanyDataCompanyExtensionInterface) HasQuoteConfig() bool {
	if o != nil && !IsNil(o.QuoteConfig) {
		return true
	}

	return false
}

// SetQuoteConfig gets a reference to the given NegotiableQuoteDataCompanyQuoteConfigInterface and assigns it to the QuoteConfig field.
func (o *CompanyDataCompanyExtensionInterface) SetQuoteConfig(v NegotiableQuoteDataCompanyQuoteConfigInterface) {
	o.QuoteConfig = &v
}

// GetIsPurchaseOrderEnabled returns the IsPurchaseOrderEnabled field value if set, zero value otherwise.
func (o *CompanyDataCompanyExtensionInterface) GetIsPurchaseOrderEnabled() bool {
	if o == nil || IsNil(o.IsPurchaseOrderEnabled) {
		var ret bool
		return ret
	}
	return *o.IsPurchaseOrderEnabled
}

// GetIsPurchaseOrderEnabledOk returns a tuple with the IsPurchaseOrderEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyDataCompanyExtensionInterface) GetIsPurchaseOrderEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPurchaseOrderEnabled) {
		return nil, false
	}
	return o.IsPurchaseOrderEnabled, true
}

// HasIsPurchaseOrderEnabled returns a boolean if a field has been set.
func (o *CompanyDataCompanyExtensionInterface) HasIsPurchaseOrderEnabled() bool {
	if o != nil && !IsNil(o.IsPurchaseOrderEnabled) {
		return true
	}

	return false
}

// SetIsPurchaseOrderEnabled gets a reference to the given bool and assigns it to the IsPurchaseOrderEnabled field.
func (o *CompanyDataCompanyExtensionInterface) SetIsPurchaseOrderEnabled(v bool) {
	o.IsPurchaseOrderEnabled = &v
}

// GetApplicableShippingMethod returns the ApplicableShippingMethod field value if set, zero value otherwise.
func (o *CompanyDataCompanyExtensionInterface) GetApplicableShippingMethod() int32 {
	if o == nil || IsNil(o.ApplicableShippingMethod) {
		var ret int32
		return ret
	}
	return *o.ApplicableShippingMethod
}

// GetApplicableShippingMethodOk returns a tuple with the ApplicableShippingMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyDataCompanyExtensionInterface) GetApplicableShippingMethodOk() (*int32, bool) {
	if o == nil || IsNil(o.ApplicableShippingMethod) {
		return nil, false
	}
	return o.ApplicableShippingMethod, true
}

// HasApplicableShippingMethod returns a boolean if a field has been set.
func (o *CompanyDataCompanyExtensionInterface) HasApplicableShippingMethod() bool {
	if o != nil && !IsNil(o.ApplicableShippingMethod) {
		return true
	}

	return false
}

// SetApplicableShippingMethod gets a reference to the given int32 and assigns it to the ApplicableShippingMethod field.
func (o *CompanyDataCompanyExtensionInterface) SetApplicableShippingMethod(v int32) {
	o.ApplicableShippingMethod = &v
}

// GetAvailableShippingMethods returns the AvailableShippingMethods field value if set, zero value otherwise.
func (o *CompanyDataCompanyExtensionInterface) GetAvailableShippingMethods() string {
	if o == nil || IsNil(o.AvailableShippingMethods) {
		var ret string
		return ret
	}
	return *o.AvailableShippingMethods
}

// GetAvailableShippingMethodsOk returns a tuple with the AvailableShippingMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyDataCompanyExtensionInterface) GetAvailableShippingMethodsOk() (*string, bool) {
	if o == nil || IsNil(o.AvailableShippingMethods) {
		return nil, false
	}
	return o.AvailableShippingMethods, true
}

// HasAvailableShippingMethods returns a boolean if a field has been set.
func (o *CompanyDataCompanyExtensionInterface) HasAvailableShippingMethods() bool {
	if o != nil && !IsNil(o.AvailableShippingMethods) {
		return true
	}

	return false
}

// SetAvailableShippingMethods gets a reference to the given string and assigns it to the AvailableShippingMethods field.
func (o *CompanyDataCompanyExtensionInterface) SetAvailableShippingMethods(v string) {
	o.AvailableShippingMethods = &v
}

// GetUseConfigSettingsShipping returns the UseConfigSettingsShipping field value if set, zero value otherwise.
func (o *CompanyDataCompanyExtensionInterface) GetUseConfigSettingsShipping() int32 {
	if o == nil || IsNil(o.UseConfigSettingsShipping) {
		var ret int32
		return ret
	}
	return *o.UseConfigSettingsShipping
}

// GetUseConfigSettingsShippingOk returns a tuple with the UseConfigSettingsShipping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyDataCompanyExtensionInterface) GetUseConfigSettingsShippingOk() (*int32, bool) {
	if o == nil || IsNil(o.UseConfigSettingsShipping) {
		return nil, false
	}
	return o.UseConfigSettingsShipping, true
}

// HasUseConfigSettingsShipping returns a boolean if a field has been set.
func (o *CompanyDataCompanyExtensionInterface) HasUseConfigSettingsShipping() bool {
	if o != nil && !IsNil(o.UseConfigSettingsShipping) {
		return true
	}

	return false
}

// SetUseConfigSettingsShipping gets a reference to the given int32 and assigns it to the UseConfigSettingsShipping field.
func (o *CompanyDataCompanyExtensionInterface) SetUseConfigSettingsShipping(v int32) {
	o.UseConfigSettingsShipping = &v
}

func (o CompanyDataCompanyExtensionInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompanyDataCompanyExtensionInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplicablePaymentMethod) {
		toSerialize["applicable_payment_method"] = o.ApplicablePaymentMethod
	}
	if !IsNil(o.AvailablePaymentMethods) {
		toSerialize["available_payment_methods"] = o.AvailablePaymentMethods
	}
	if !IsNil(o.UseConfigSettings) {
		toSerialize["use_config_settings"] = o.UseConfigSettings
	}
	if !IsNil(o.QuoteConfig) {
		toSerialize["quote_config"] = o.QuoteConfig
	}
	if !IsNil(o.IsPurchaseOrderEnabled) {
		toSerialize["is_purchase_order_enabled"] = o.IsPurchaseOrderEnabled
	}
	if !IsNil(o.ApplicableShippingMethod) {
		toSerialize["applicable_shipping_method"] = o.ApplicableShippingMethod
	}
	if !IsNil(o.AvailableShippingMethods) {
		toSerialize["available_shipping_methods"] = o.AvailableShippingMethods
	}
	if !IsNil(o.UseConfigSettingsShipping) {
		toSerialize["use_config_settings_shipping"] = o.UseConfigSettingsShipping
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CompanyDataCompanyExtensionInterface) UnmarshalJSON(data []byte) (err error) {
	varCompanyDataCompanyExtensionInterface := _CompanyDataCompanyExtensionInterface{}

	err = json.Unmarshal(data, &varCompanyDataCompanyExtensionInterface)

	if err != nil {
		return err
	}

	*o = CompanyDataCompanyExtensionInterface(varCompanyDataCompanyExtensionInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "applicable_payment_method")
		delete(additionalProperties, "available_payment_methods")
		delete(additionalProperties, "use_config_settings")
		delete(additionalProperties, "quote_config")
		delete(additionalProperties, "is_purchase_order_enabled")
		delete(additionalProperties, "applicable_shipping_method")
		delete(additionalProperties, "available_shipping_methods")
		delete(additionalProperties, "use_config_settings_shipping")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *CompanyDataCompanyExtensionInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *CompanyDataCompanyExtensionInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableCompanyDataCompanyExtensionInterface struct {
	value *CompanyDataCompanyExtensionInterface
	isSet bool
}

func (v NullableCompanyDataCompanyExtensionInterface) Get() *CompanyDataCompanyExtensionInterface {
	return v.value
}

func (v *NullableCompanyDataCompanyExtensionInterface) Set(val *CompanyDataCompanyExtensionInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyDataCompanyExtensionInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyDataCompanyExtensionInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyDataCompanyExtensionInterface(val *CompanyDataCompanyExtensionInterface) *NullableCompanyDataCompanyExtensionInterface {
	return &NullableCompanyDataCompanyExtensionInterface{value: val, isSet: true}
}

func (v NullableCompanyDataCompanyExtensionInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyDataCompanyExtensionInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
